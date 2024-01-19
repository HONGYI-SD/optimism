package preimages

import (
	"context"
	"errors"
	"math/big"
	"testing"

	"github.com/ethereum-optimism/optimism/op-challenger/game/fault/contracts"
	"github.com/ethereum-optimism/optimism/op-challenger/game/fault/types"
	"github.com/ethereum-optimism/optimism/op-challenger/game/keccak/matrix"
	"github.com/ethereum-optimism/optimism/op-service/testlog"
	"github.com/ethereum-optimism/optimism/op-service/txmgr"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/stretchr/testify/require"
)

var (
	mockAddLeavesError = errors.New("mock add leaves error")
	mockSqueezeError   = errors.New("mock squeeze error")
)

func TestLargePreimageUploader_UploadPreimage_ContractFails(t *testing.T) {
	t.Run("InitFails", func(t *testing.T) {
		oracle, _, contract := newTestLargePreimageUploader(t)
		contract.initFails = true
		err := oracle.UploadPreimage(context.Background(), 0, &types.PreimageOracleData{})
		require.ErrorIs(t, err, mockInitLPPError)
		require.Equal(t, 1, contract.initCalls)
	})

	t.Run("AddLeavesFails", func(t *testing.T) {
		oracle, _, contract := newTestLargePreimageUploader(t)
		contract.addFails = true
		err := oracle.UploadPreimage(context.Background(), 0, &types.PreimageOracleData{})
		require.ErrorIs(t, err, mockAddLeavesError)
		require.Equal(t, 1, contract.addCalls)
	})

	t.Run("SqueezeFails", func(t *testing.T) {
		oracle, _, contract := newTestLargePreimageUploader(t)
		contract.squeezeFails = true
		err := oracle.UploadPreimage(context.Background(), 0, &types.PreimageOracleData{})
		require.ErrorIs(t, err, mockSqueezeError)
		require.Equal(t, 1, contract.squeezeCalls)
	})
}

func TestLargePreimageUploader_UploadPreimage_Succeeds(t *testing.T) {
	fullLeaf := make([]byte, matrix.LeafSize)
	for i := 0; i < matrix.LeafSize; i++ {
		fullLeaf[i] = byte(i)
	}
	chunk := make([]byte, 0, MaxChunkSize)
	for i := 0; i < MaxLeafsPerChunk; i++ {
		chunk = append(chunk, fullLeaf...)
	}
	tests := []struct {
		name          string
		input         []byte
		addCalls      int
		prestateLeaf  contracts.Leaf
		poststateLeaf contracts.Leaf
	}{
		{
			name:     "FullLeaf",
			input:    fullLeaf,
			addCalls: 1,
			prestateLeaf: contracts.Leaf{
				Input:           ([136]byte)(fullLeaf),
				Index:           big.NewInt(0),
				StateCommitment: common.HexToHash("9788a3b3bc36c482525b5890767be37130c997917bceca6e91a6c93359a4d1c6"),
			},
			poststateLeaf: contracts.Leaf{
				Input:           ([136]byte)(make([]byte, matrix.LeafSize)),
				Index:           big.NewInt(1),
				StateCommitment: common.HexToHash("78358b902b7774b314bcffdf0948746f18d6044086e76e3924d585dca3486c7d"),
			},
		},
		{
			name:     "MultipleLeaves",
			input:    append(fullLeaf, append(fullLeaf, fullLeaf...)...),
			addCalls: 1,
			prestateLeaf: contracts.Leaf{
				Input:           ([136]byte)(fullLeaf),
				Index:           big.NewInt(2),
				StateCommitment: common.HexToHash("e3deed8ab6f8bbcf3d4fe825d74f703b3f2fc2f5b0afaa2574926fcfd0d4c895"),
			},
			poststateLeaf: contracts.Leaf{
				Input:           ([136]byte)(make([]byte, matrix.LeafSize)),
				Index:           big.NewInt(3),
				StateCommitment: common.HexToHash("79115eeab1ff2eccf5baf3ea2dda13bc79c548ce906bdd16433a23089c679df2"),
			},
		},
		{
			name:     "MultipleLeavesUnaligned",
			input:    append(fullLeaf, append(fullLeaf, byte(9))...),
			addCalls: 1,
			prestateLeaf: contracts.Leaf{
				Input:           ([136]byte)(fullLeaf),
				Index:           big.NewInt(1),
				StateCommitment: common.HexToHash("b5ea400e375b2c1ce348f3cc4ad5b6ad28e1b36759ddd2aba155f0b1d476b015"),
			},
			poststateLeaf: contracts.Leaf{
				Input:           ([136]byte)(append([]byte{byte(9)}, make([]byte, matrix.LeafSize-1)...)),
				Index:           big.NewInt(2),
				StateCommitment: common.HexToHash("fa87e115dc4786e699bf80cc75d13ac1e2db0708c1418fc8cbc9800d17b5811a"),
			},
		},
		{
			name:     "MultipleChunks",
			input:    append(chunk, append(fullLeaf, fullLeaf...)...),
			addCalls: 2,
			prestateLeaf: contracts.Leaf{
				Input:           ([136]byte)(fullLeaf),
				Index:           big.NewInt(301),
				StateCommitment: common.HexToHash("4e9c55542478939feca4ff55ee98fbc632bb65a784a55b94536644bc87298ca4"),
			},
			poststateLeaf: contracts.Leaf{
				Input:           ([136]byte)(make([]byte, matrix.LeafSize)),
				Index:           big.NewInt(302),
				StateCommitment: common.HexToHash("775020bfcaa93700263d040a4eeec3c8c3cf09e178457d04044594beaaf5e20b"),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			oracle, _, contract := newTestLargePreimageUploader(t)
			data := types.PreimageOracleData{
				OracleData: test.input,
			}
			err := oracle.UploadPreimage(context.Background(), 0, &data)
			require.NoError(t, err)
			require.Equal(t, test.addCalls, contract.addCalls)
			// There must always be at least one init and squeeze call
			// for successful large preimage upload calls.
			require.Equal(t, 1, contract.initCalls)
			require.Equal(t, 1, contract.squeezeCalls)
			require.Equal(t, test.prestateLeaf, contract.squeezePrestate)
			require.Equal(t, test.poststateLeaf, contract.squeezePostate)
		})
	}

}

func newTestLargePreimageUploader(t *testing.T) (*LargePreimageUploader, *mockTxMgr, *mockPreimageOracleContract) {
	logger := testlog.Logger(t, log.LvlError)
	txMgr := &mockTxMgr{}
	contract := &mockPreimageOracleContract{
		addData: make([]byte, 0),
	}
	return NewLargePreimageUploader(logger, txMgr, contract), txMgr, contract
}

type mockPreimageOracleContract struct {
	initCalls       int
	initFails       bool
	addCalls        int
	addFails        bool
	squeezeCalls    int
	squeezeFails    bool
	addData         []byte
	squeezePrestate contracts.Leaf
	squeezePostate  contracts.Leaf
}

func (s *mockPreimageOracleContract) InitLargePreimage(_ *big.Int, _ uint32, _ uint32) (txmgr.TxCandidate, error) {
	s.initCalls++
	if s.initFails {
		return txmgr.TxCandidate{}, mockInitLPPError
	}
	return txmgr.TxCandidate{}, nil
}
func (s *mockPreimageOracleContract) AddLeaves(_ *big.Int, input []byte, _ [][32]byte, _ bool) (txmgr.TxCandidate, error) {
	s.addCalls++
	s.addData = append(s.addData, input...)
	if s.addFails {
		return txmgr.TxCandidate{}, mockAddLeavesError
	}
	return txmgr.TxCandidate{}, nil
}
func (s *mockPreimageOracleContract) Squeeze(_ common.Address, _ *big.Int, _ *matrix.StateMatrix, prestate contracts.Leaf, _ contracts.MerkleProof, postate contracts.Leaf, _ contracts.MerkleProof) (txmgr.TxCandidate, error) {
	s.squeezeCalls++
	s.squeezePrestate = prestate
	s.squeezePostate = postate
	if s.squeezeFails {
		return txmgr.TxCandidate{}, mockSqueezeError
	}
	return txmgr.TxCandidate{}, nil
}
func (s *mockPreimageOracleContract) CallSqueeze(_ context.Context, _ common.Address, _ *big.Int, _ *matrix.StateMatrix, _ contracts.Leaf, _ contracts.MerkleProof, _ contracts.Leaf, _ contracts.MerkleProof) error {
	return nil
}
