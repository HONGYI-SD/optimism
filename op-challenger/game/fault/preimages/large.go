package preimages

import (
	"bytes"
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"math/big"

	"github.com/ethereum-optimism/optimism/op-challenger/game/fault/contracts"
	"github.com/ethereum-optimism/optimism/op-challenger/game/fault/types"
	"github.com/ethereum-optimism/optimism/op-challenger/game/keccak/matrix"
	"github.com/ethereum-optimism/optimism/op-service/txmgr"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

var _ PreimageUploader = (*LargePreimageUploader)(nil)

// MaxLeafsPerChunk is the maximum number of leafs per chunk.
const MaxLeafsPerChunk = 300

// MaxChunkSize is the maximum size of a preimage chunk in bytes.
// Notice, the max chunk size must be a multiple of the leaf size.
// The max chunk size is roughly 0.04MB to avoid memory expansion.
const MaxChunkSize = MaxLeafsPerChunk * matrix.LeafSize

// Chunk is a contigous segment of preimage data.
type Chunk struct {
	// Input is the preimage data.
	Input []byte
	// Commitments are the keccak commitments for each leaf in the chunk.
	Commitments [][32]byte
	// Finalize indicates whether the chunk is the final chunk.
	Finalize bool
}

// LargePreimageUploader handles uploading large preimages by
// streaming the merkleized preimage to the PreimageOracle contract,
// tightly packed across multiple transactions.
type LargePreimageUploader struct {
	log log.Logger

	txMgr    txmgr.TxManager
	contract PreimageOracleContract
}

func NewLargePreimageUploader(logger log.Logger, txMgr txmgr.TxManager, contract PreimageOracleContract) *LargePreimageUploader {
	return &LargePreimageUploader{logger, txMgr, contract}
}

func NewContractLeaf(input []byte, index *big.Int, commitment common.Hash) contracts.Leaf {
	// Pad the input to the 136 byte leaf size.
	input = append(input, make([]byte, matrix.LeafSize-len(input))...)
	return contracts.Leaf{
		Input:           ([136]byte)(input),
		Index:           index,
		StateCommitment: commitment,
	}
}

func (p *LargePreimageUploader) UploadPreimage(ctx context.Context, parent uint64, data *types.PreimageOracleData) error {
	stateMatrix := matrix.NewStateMatrix()
	// SAFETY: The chunk size is a multiple of the leaf size (except for the last chunk/leaf).
	chunk := make([]byte, 0, MaxChunkSize)
	calls := []Chunk{}
	// SAFETY: The prestate and postate leaf are always set since the preimage data must be sufficiently large.
	var prestateLeaf contracts.Leaf
	var poststateLeaf contracts.Leaf
	commitments := make([][32]byte, 0, MaxLeafsPerChunk)
	in := bytes.NewReader(data.OracleData)
	for i := 0; ; i++ {
		// SAFETY: The preimage input is never mutated and the keccak state matrix handles the keccak padding rules.
		leaf, err := stateMatrix.AbsorbNextLeaf(in)
		stateCommitment := stateMatrix.StateCommitment()
		chunk = append(chunk, leaf...)
		commitments = append(commitments, stateCommitment)
		// SAFETY: The last leaf will always return an [io.EOF] error from [AbsorbNextLeaf].
		// SAFETY: Since the last leaf will always hit this branch, we can safely update the poststate leaf.
		if errors.Is(err, io.EOF) {
			calls = append(calls, Chunk{chunk, commitments[:], true})
			poststateLeaf = NewContractLeaf(leaf, big.NewInt(int64(i)), stateCommitment)
			break
		}
		if err != nil {
			return fmt.Errorf("failed to absorb leaf: %w", err)
		}

		// Only create a call if the chunk is full.
		if len(chunk) >= MaxChunkSize {
			calls = append(calls, Chunk{chunk, commitments[:], false})
			chunk = make([]byte, 0, MaxChunkSize)
			commitments = make([][32]byte, 0, MaxLeafsPerChunk)
		}

		// SAFETY: This is never the last leaf since the loop will break before this assignment.
		prestateLeaf = NewContractLeaf(leaf, big.NewInt(int64(i)), stateCommitment)
	}

	// TODO(client-pod#473): The UUID must be deterministic so the challenger can resume uploads.
	uuid, err := p.newUUID()
	if err != nil {
		return fmt.Errorf("failed to generate UUID: %w", err)
	}

	err = p.initLargePreimage(ctx, uuid, data.OracleOffset, uint32(len(data.OracleData)))
	if err != nil {
		return fmt.Errorf("failed to initialize large preimage with uuid: %s: %w", uuid, err)
	}

	err = p.addLargePreimageLeafs(ctx, uuid, calls)
	if err != nil {
		return fmt.Errorf("failed to add leaves to large preimage with uuid: %s: %w", uuid, err)
	}

	// TODO(client-pod#474): To generate the inclusion proofs, we need to construct the merkle tree for the preimage data.
	// TODO(client-pod#474): 1. Create the fully zeroed out merkle tree.
	// TODO(client-pod#474): 2. Take all of the input and state commitments and reconstruct the merkle tree.
	// TODO(client-pod#474): 3. Generate proofs of inclusion.
	prestateProof := contracts.MerkleProof{}
	poststateProof := contracts.MerkleProof{}

	return p.Squeeze(
		ctx,
		p.txMgr.From(),
		uuid,
		stateMatrix,
		prestateLeaf,
		prestateProof,
		poststateLeaf,
		poststateProof,
	)
}

func (p *LargePreimageUploader) newUUID() (*big.Int, error) {
	max := new(big.Int)
	max.Exp(big.NewInt(2), big.NewInt(130), nil).Sub(max, big.NewInt(1))
	return rand.Int(rand.Reader, max)
}

func (p *LargePreimageUploader) Squeeze(
	ctx context.Context,
	claimant common.Address,
	uuid *big.Int,
	stateMatrix *matrix.StateMatrix,
	preState contracts.Leaf,
	preStateProof contracts.MerkleProof,
	postState contracts.Leaf,
	postStateProof contracts.MerkleProof,
) error {
	// TODO(client-pod#474): Return the ErrChallengePeriodNotOver error if the challenge period is not over.
	//                       This allows the responder to retry the squeeze later.
	//                       Other errors should force the responder to stop retrying.
	//                       Nil errors should indicate the squeeze was successful.
	if err := p.contract.CallSqueeze(ctx, claimant, uuid, stateMatrix, preState, preStateProof, postState, postStateProof); err != nil {
		return fmt.Errorf("failed to call squeeze: %w", err)
	}
	tx, err := p.contract.Squeeze(claimant, uuid, stateMatrix, preState, preStateProof, postState, postStateProof)
	if err != nil {
		return fmt.Errorf("failed to create pre-image oracle tx: %w", err)
	}
	if err := p.sendTxAndWait(ctx, tx); err != nil {
		return fmt.Errorf("failed to populate pre-image oracle: %w", err)
	}
	return nil
}

// initLargePreimage initializes the large preimage proposal.
// This method *must* be called before adding any leaves.
func (p *LargePreimageUploader) initLargePreimage(ctx context.Context, uuid *big.Int, partOffset uint32, claimedSize uint32) error {
	candidate, err := p.contract.InitLargePreimage(uuid, partOffset, claimedSize)
	if err != nil {
		return fmt.Errorf("failed to create pre-image oracle tx: %w", err)
	}
	if err := p.sendTxAndWait(ctx, candidate); err != nil {
		return fmt.Errorf("failed to populate pre-image oracle: %w", err)
	}
	return nil
}

// addLargePreimageLeafs adds leafs to the large preimage proposal.
// This method **must** be called after calling [initLargePreimage].
// SAFETY: submits transactions in a [Queue] for latency while preserving submission order.
func (p *LargePreimageUploader) addLargePreimageLeafs(ctx context.Context, uuid *big.Int, chunks []Chunk) error {
	queue := txmgr.NewQueue[int](ctx, p.txMgr, 10)
	receiptChs := make([]chan txmgr.TxReceipt[int], len(chunks))
	for i, chunk := range chunks {
		tx, err := p.contract.AddLeaves(uuid, chunk.Input, chunk.Commitments, chunk.Finalize)
		if err != nil {
			return fmt.Errorf("failed to create pre-image oracle tx: %w", err)
		}
		receiptChs[i] = make(chan txmgr.TxReceipt[int], 1)
		queue.Send(i, tx, receiptChs[i])
	}
	for _, receiptCh := range receiptChs {
		receipt := <-receiptCh
		if receipt.Err != nil {
			return receipt.Err
		}
		if receipt.Receipt.Status == ethtypes.ReceiptStatusFailed {
			p.log.Error("LargePreimageUploader add leafs tx successfully published but reverted", "tx_hash", receipt.Receipt.TxHash)
		} else {
			p.log.Debug("LargePreimageUploader add leafs tx successfully published", "tx_hash", receipt.Receipt.TxHash)
		}
	}
	return nil
}

// sendTxAndWait sends a transaction through the [txmgr] and waits for a receipt.
// This sets the tx GasLimit to 0, performing gas estimation online through the [txmgr].
func (p *LargePreimageUploader) sendTxAndWait(ctx context.Context, candidate txmgr.TxCandidate) error {
	receipt, err := p.txMgr.Send(ctx, candidate)
	if err != nil {
		return err
	}
	if receipt.Status == ethtypes.ReceiptStatusFailed {
		p.log.Error("LargePreimageUploader tx successfully published but reverted", "tx_hash", receipt.TxHash)
	} else {
		p.log.Debug("LargePreimageUploader tx successfully published", "tx_hash", receipt.TxHash)
	}
	return nil
}
