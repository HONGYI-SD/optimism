// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const PreimageOracleStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/cannon/PreimageOracle.sol:PreimageOracle\",\"label\":\"preimageLengths\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_mapping(t_bytes32,t_uint256)\"},{\"astId\":1001,\"contract\":\"src/cannon/PreimageOracle.sol:PreimageOracle\",\"label\":\"preimageParts\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_mapping(t_bytes32,t_mapping(t_uint256,t_bytes32))\"},{\"astId\":1002,\"contract\":\"src/cannon/PreimageOracle.sol:PreimageOracle\",\"label\":\"preimagePartOk\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_mapping(t_bytes32,t_mapping(t_uint256,t_bool))\"},{\"astId\":1003,\"contract\":\"src/cannon/PreimageOracle.sol:PreimageOracle\",\"label\":\"zeroHashes\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_array(t_bytes32)16_storage\"},{\"astId\":1004,\"contract\":\"src/cannon/PreimageOracle.sol:PreimageOracle\",\"label\":\"proposals\",\"offset\":0,\"slot\":\"19\",\"type\":\"t_array(t_struct(LargePreimageProposalKeys)1009_storage)dyn_storage\"},{\"astId\":1005,\"contract\":\"src/cannon/PreimageOracle.sol:PreimageOracle\",\"label\":\"proposalBranches\",\"offset\":0,\"slot\":\"20\",\"type\":\"t_mapping(t_address,t_mapping(t_uint256,t_array(t_bytes32)16_storage))\"},{\"astId\":1006,\"contract\":\"src/cannon/PreimageOracle.sol:PreimageOracle\",\"label\":\"proposalMetadata\",\"offset\":0,\"slot\":\"21\",\"type\":\"t_mapping(t_address,t_mapping(t_uint256,t_userDefinedValueType(LPPMetaData)1010))\"},{\"astId\":1007,\"contract\":\"src/cannon/PreimageOracle.sol:PreimageOracle\",\"label\":\"proposalParts\",\"offset\":0,\"slot\":\"22\",\"type\":\"t_mapping(t_address,t_mapping(t_uint256,t_bytes32))\"},{\"astId\":1008,\"contract\":\"src/cannon/PreimageOracle.sol:PreimageOracle\",\"label\":\"proposalBlocks\",\"offset\":0,\"slot\":\"23\",\"type\":\"t_mapping(t_address,t_mapping(t_uint256,t_array(t_uint64)dyn_storage))\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_bytes32)16_storage\":{\"encoding\":\"inplace\",\"label\":\"bytes32[16]\",\"numberOfBytes\":\"512\",\"base\":\"t_bytes32\"},\"t_array(t_struct(LargePreimageProposalKeys)1009_storage)dyn_storage\":{\"encoding\":\"dynamic_array\",\"label\":\"struct PreimageOracle.LargePreimageProposalKeys[]\",\"numberOfBytes\":\"32\",\"base\":\"t_struct(LargePreimageProposalKeys)1009_storage\"},\"t_array(t_uint64)dyn_storage\":{\"encoding\":\"dynamic_array\",\"label\":\"uint64[]\",\"numberOfBytes\":\"32\",\"base\":\"t_uint64\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_mapping(t_address,t_mapping(t_uint256,t_array(t_bytes32)16_storage))\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e mapping(uint256 =\u003e bytes32[16]))\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_mapping(t_uint256,t_array(t_bytes32)16_storage)\"},\"t_mapping(t_address,t_mapping(t_uint256,t_array(t_uint64)dyn_storage))\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e mapping(uint256 =\u003e uint64[]))\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_mapping(t_uint256,t_array(t_uint64)dyn_storage)\"},\"t_mapping(t_address,t_mapping(t_uint256,t_bytes32))\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e mapping(uint256 =\u003e bytes32))\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_mapping(t_uint256,t_bytes32)\"},\"t_mapping(t_address,t_mapping(t_uint256,t_userDefinedValueType(LPPMetaData)1010))\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e mapping(uint256 =\u003e LPPMetaData))\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_mapping(t_uint256,t_userDefinedValueType(LPPMetaData)1010)\"},\"t_mapping(t_bytes32,t_mapping(t_uint256,t_bool))\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e mapping(uint256 =\u003e bool))\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_mapping(t_uint256,t_bool)\"},\"t_mapping(t_bytes32,t_mapping(t_uint256,t_bytes32))\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e mapping(uint256 =\u003e bytes32))\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_mapping(t_uint256,t_bytes32)\"},\"t_mapping(t_bytes32,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_uint256\"},\"t_mapping(t_uint256,t_array(t_bytes32)16_storage)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint256 =\u003e bytes32[16])\",\"numberOfBytes\":\"32\",\"key\":\"t_uint256\",\"value\":\"t_array(t_bytes32)16_storage\"},\"t_mapping(t_uint256,t_array(t_uint64)dyn_storage)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint256 =\u003e uint64[])\",\"numberOfBytes\":\"32\",\"key\":\"t_uint256\",\"value\":\"t_array(t_uint64)dyn_storage\"},\"t_mapping(t_uint256,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint256 =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_uint256\",\"value\":\"t_bool\"},\"t_mapping(t_uint256,t_bytes32)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint256 =\u003e bytes32)\",\"numberOfBytes\":\"32\",\"key\":\"t_uint256\",\"value\":\"t_bytes32\"},\"t_mapping(t_uint256,t_userDefinedValueType(LPPMetaData)1010)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint256 =\u003e LPPMetaData)\",\"numberOfBytes\":\"32\",\"key\":\"t_uint256\",\"value\":\"t_userDefinedValueType(LPPMetaData)1010\"},\"t_struct(LargePreimageProposalKeys)1009_storage\":{\"encoding\":\"inplace\",\"label\":\"struct PreimageOracle.LargePreimageProposalKeys\",\"numberOfBytes\":\"64\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint64\":{\"encoding\":\"inplace\",\"label\":\"uint64\",\"numberOfBytes\":\"8\"},\"t_userDefinedValueType(LPPMetaData)1010\":{\"encoding\":\"inplace\",\"label\":\"LPPMetaData\",\"numberOfBytes\":\"32\"}}}"

var PreimageOracleStorageLayout = new(solc.StorageLayout)

var PreimageOracleDeployedBin = "0x608060405234801561001057600080fd5b50600436106101a35760003560e01c80639d7e8769116100ee578063dd24f9bf11610097578063ec5efcbc11610071578063ec5efcbc1461043a578063f3f480d91461044d578063faf37bc714610473578063fef2b4ed1461048657600080fd5b8063dd24f9bf146103d9578063e03110e1146103ff578063e15926111461042757600080fd5b8063b4801e61116100c8578063b4801e61146103ab578063d18534b5146103be578063da35c664146103d157600080fd5b80639d7e87691461035a5780639f99ef821461036d578063b2e67ba81461038057600080fd5b806361238bde116101505780638542cf501161012a5780638542cf50146102ae578063882856ef146102ec5780639d53a6481461031857600080fd5b806361238bde146102455780636551927b146102705780637ac547671461029b57600080fd5b80633909af5c116101815780633909af5c146102155780634d52b4c91461022a57806352f0f3ad1461023257600080fd5b8063013cf08b146101a85780630359a563146101ec5780632055b36b1461020d575b600080fd5b6101bb6101b63660046127b4565b6104a6565b6040805173ffffffffffffffffffffffffffffffffffffffff90931683526020830191909152015b60405180910390f35b6101ff6101fa3660046127f6565b6104eb565b6040519081526020016101e3565b6101ff601081565b6102286102233660046129f1565b610623565b005b6101ff610871565b6101ff610240366004612add565b61088c565b6101ff610253366004612b18565b600160209081526000928352604080842090915290825290205481565b6101ff61027e3660046127f6565b601560209081526000928352604080842090915290825290205481565b6101ff6102a93660046127b4565b610961565b6102dc6102bc366004612b18565b600260209081526000928352604080842090915290825290205460ff1681565b60405190151581526020016101e3565b6102ff6102fa366004612b3a565b610978565b60405167ffffffffffffffff90911681526020016101e3565b6101ff6103263660046127f6565b73ffffffffffffffffffffffffffffffffffffffff9091166000908152601760209081526040808320938352929052205490565b610228610368366004612baf565b6109d2565b61022861037b366004612c3b565b610b73565b6101ff61038e3660046127f6565b601660209081526000928352604080842090915290825290205481565b6101ff6103b9366004612b3a565b6110ee565b6102286103cc3660046129f1565b611120565b6013546101ff565b7f00000000000000000000000000000000000000000000000000000000000000006101ff565b61041261040d366004612b18565b6114eb565b604080519283526020830191909152016101e3565b610228610435366004612ccc565b6115dc565b610228610448366004612d18565b6116e4565b7f00000000000000000000000000000000000000000000000000000000000000006101ff565b610228610481366004612db1565b61185e565b6101ff6104943660046127b4565b60006020819052908152604090205481565b601381815481106104b657600080fd5b60009182526020909120600290910201805460019091015473ffffffffffffffffffffffffffffffffffffffff909116915082565b73ffffffffffffffffffffffffffffffffffffffff82166000908152601560209081526040808320848452909152812054819061052e9060601c63ffffffff1690565b63ffffffff16905060005b601081101561061b57816001166001036105c15773ffffffffffffffffffffffffffffffffffffffff851660009081526014602090815260408083208784529091529020816010811061058e5761058e612ded565b01546040805160208101929092528101849052606001604051602081830303815290604052805190602001209250610602565b82600382601081106105d5576105d5612ded565b01546040805160208101939093528201526060016040516020818303038152906040528051906020012092505b60019190911c908061061381612e4b565b915050610539565b505092915050565b600061062f8a8a6104eb565b905061065286868360208b013561064d6106488d612e83565b611a71565b611ab1565b80156106705750610670838383602088013561064d6106488a612e83565b6106a6576040517f09bde33900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8660400135886040516020016106bc9190612f52565b6040516020818303038152906040528051906020012014610709576040517f1968a90200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b83602001358760200135600161071f9190612f90565b14610756576040517f9a3b119900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61079e886107648680612fa8565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611b1292505050565b6107a788611c6d565b8360400135886040516020016107bd9190612f52565b604051602081830303815290604052805190602001200361080a576040517f9843145b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050505073ffffffffffffffffffffffffffffffffffffffff9590951660009081526015602090815260408083209683529590529390932080547fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000016600117905550505050565b600161087f6010600261312f565b610889919061313b565b81565b60006108988686612509565b90506108a5836008612f90565b8211806108b25750602083115b156108e9576040517ffe25498700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000602081815260c085901b82526008959095528251828252600286526040808320858452875280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660019081179091558484528752808320948352938652838220558181529384905292205592915050565b6003816010811061097157600080fd5b0154905081565b601760205282600052604060002060205281600052604060002081815481106109a057600080fd5b906000526020600020906004918282040191900660080292509250509054906101000a900467ffffffffffffffff1681565b6000806000603088600037602060006030600060025afa806109fc5763f91129696000526004601cfd5b6000517effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f010000000000000000000000000000000000000000000000000000000000000017608081815260a08d905260c08c905260308b60e037603089609083013760008060c083600a5afa925082610a7e576309bde3396000526004601cfd5b6044359550600886018710610a9b5763fe2549876000526004601cfd5b8560c01b81528b600882015286810151935060308b8237603081019c909c52505060509099207effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f0500000000000000000000000000000000000000000000000000000000000000176000818152600260209081526040808320878452825280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600190811790915584845282528083209683529581528582209b909b5590815298899052509620959095555050505050565b333214610bac576040517fba092d1600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60608115610bc557610bbe86866125b6565b9050610bff565b85858080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509293505050505b3360009081526014602090815260408083208a845290915280822081516102008101928390529160109082845b815481526020019060010190808311610c2c5750503360009081526015602090815260408083208f8452909152902054939450610c6e925083915061263f9050565b63ffffffff16600003610cad576040517f87138d5c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610cb78160c01c90565b67ffffffffffffffff1615610cf8576040517f475a253500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000610d0a8260a01c63ffffffff1690565b67ffffffffffffffff1690506000610d288360401c63ffffffff1690565b63ffffffff169050600882108015610d3e575080155b15610dc5576000610d558460801c63ffffffff1690565b905060008160c01b6000528b356008528351905080601660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008f8152602001908152602001600020819055505050610e7a565b60088210158015610de3575080610ddd60088461313b565b92508210155b8015610df75750610df48982612f90565b82105b15610e7a576000610e08828461313b565b905089610e16826020612f90565b10158015610e22575086155b15610e59576040517ffe25498700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b3360009081526016602090815260408083208f84529091529020908b013590555b6000610e8c8460601c63ffffffff1690565b63ffffffff169050855160208701608882048a1415608883061715610eb9576307b1daf16000526004601cfd5b60405160c8810160405260005b83811015610f69578083018051835260208101516020840152604081015160408401526060810151606084015260808101516080840152508460888301526088810460051b8d013560a883015260c882206001860195508560005b610200811015610f5e576001821615610f3e5782818d0152610f5e565b8b81015160009081526020938452604090209260019290921c9101610f21565b505050608801610ec6565b50505050600160106002610f7d919061312f565b610f87919061313b565b811115610fc0576040517f6229572300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b3360009081526014602090815260408083208f84529091529020610fe69086601061272a565b503360009081526017602090815260408083208f845282528220805460018101825590835291206004820401805460039092166008026101000a67ffffffffffffffff81810219909316439093160291909117905561109a611048838c612f90565b60401b7fffffffffffffffffffffffffffffffffffffffff00000000ffffffffffffffff606084901b167fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff8716171790565b935086156110c55777ffffffffffffffffffffffffffffffffffffffffffffffff84164260c01b1793505b50503360009081526015602090815260408083209c83529b905299909920555050505050505050565b6014602052826000526040600020602052816000526040600020816010811061111657600080fd5b0154925083915050565b73ffffffffffffffffffffffffffffffffffffffff891660009081526015602090815260408083208b845290915290205467ffffffffffffffff811615611193576040517fc334f06900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000006111be8260c01c90565b6111d29067ffffffffffffffff164261313b565b11611209576040517f55d4cbf900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006112158b8b6104eb565b905061122e87878360208c013561064d6106488e612e83565b801561124c575061124c848483602089013561064d6106488b612e83565b611282576040517f09bde33900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8760400135896040516020016112989190612f52565b60405160208183030381529060405280519060200120146112e5576040517f1968a90200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8460200135886020013560016112fb9190612f90565b14158061132d575060016113158360601c63ffffffff1690565b61131f9190613152565b63ffffffff16856020013514155b15611364576040517f9a3b119900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006113768360801c63ffffffff1690565b63ffffffff1690508061138f8460401c63ffffffff1690565b63ffffffff16146113cc576040517f7b1dafd100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6113da8a6107648880612fa8565b6113e38a611c6d565b60006113ee8b61264b565b905060006114028560a01c63ffffffff1690565b67ffffffffffffffff169050600160026000848152602001908152602001600020600083815260200190815260200160002060006101000a81548160ff021916908315150217905550601660008f73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008e8152602001908152602001600020546001600084815260200190815260200160002060008381526020019081526020016000208190555082600080848152602001908152602001600020819055505050505050505050505050505050565b6000828152600260209081526040808320848452909152812054819060ff16611574576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f7072652d696d616765206d757374206578697374000000000000000000000000604482015260640160405180910390fd5b5060008381526020818152604090912054611590816008612f90565b61159b856020612f90565b106115b957836115ac826008612f90565b6115b6919061313b565b91505b506000938452600160209081526040808620948652939052919092205492909150565b6044356000806008830186106115fa5763fe2549876000526004601cfd5b60c083901b6080526088838682378087017ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80151908490207effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f02000000000000000000000000000000000000000000000000000000000000001760008181526002602090815260408083208b8452825280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600190811790915584845282528083209a83529981528982209390935590815290819052959095209190915550505050565b60006116f086866104eb565b9050611709838383602088013561064d6106488a612e83565b61173f576040517f09bde33900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60208401351561177b576040517f9a3b119900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611783612768565b611791816107648780612fa8565b61179a81611c6d565b8460400135816040516020016117b09190612f52565b60405160208183030381529060405280519060200120036117fd576040517f9843145b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050505073ffffffffffffffffffffffffffffffffffffffff9290921660009081526015602090815260408083209383529290522080547fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000016600117905550565b333214611897576040517fba092d1600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6118a2816008613177565b63ffffffff168263ffffffff16106118e6576040517ffe25498700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000008163ffffffff161015611946576040517f7b1dafd100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b336000818152601560209081526040808320878452825280832080547fffffffffffffffff0000000000000000ffffffffffffffffffffffffffffffff1660a09790971b7fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff169690961760809590951b9490941790945582518084019093529082529181019283526013805460018101825592525160029091027f66de8ffda797e3de9c05e8fc57b3bf0ec28a930d40b0d285d93c06501cf6a0908101805473ffffffffffffffffffffffffffffffffffffffff9093167fffffffffffffffffffffffff00000000000000000000000000000000000000009093169290921790915590517f66de8ffda797e3de9c05e8fc57b3bf0ec28a930d40b0d285d93c06501cf6a09190910155565b6000816000015182602001518360400151604051602001611a949392919061319f565b604051602081830303815290604052805190602001209050919050565b60008160005b6010811015611b05578060051b880135600186831c1660018114611aea5760008481526020839052604090209350611afb565b600082815260208590526040902093505b5050600101611ab7565b5090931495945050505050565b6088815114611b2057600080fd5b6020810160208301611ba1565b8260031b8201518060001a8160011a60081b178160021a60101b8260031a60181b17178160041a60201b8260051a60281b178260061a60301b8360071a60381b1717179050611b9b81611b86868560059190911b015190565b1867ffffffffffffffff16600586901b840152565b50505050565b611bad60008383611b2d565b611bb960018383611b2d565b611bc560028383611b2d565b611bd160038383611b2d565b611bdd60048383611b2d565b611be960058383611b2d565b611bf560068383611b2d565b611c0160078383611b2d565b611c0d60088383611b2d565b611c1960098383611b2d565b611c25600a8383611b2d565b611c31600b8383611b2d565b611c3d600c8383611b2d565b611c49600d8383611b2d565b611c55600e8383611b2d565b611c61600f8383611b2d565b611b9b60108383611b2d565b6040805178010000000000008082800000000000808a8000000080008000602082015279808b00000000800000018000000080008081800000000000800991810191909152788a00000000000000880000000080008009000000008000000a60608201527b8000808b800000000000008b8000000000008089800000000000800360808201527f80000000000080028000000000000080000000000000800a800000008000000a60a08201527f800000008000808180000000000080800000000080000001800000008000800860c082015260009060e001604051602081830303815290604052905060208201602082016123e9565b6102808101516101e082015161014083015160a0840151845118189118186102a082015161020083015161016084015160c0850151602086015118189118186102c083015161022084015161018085015160e0860151604087015118189118186102e08401516102408501516101a0860151610100870151606088015118189118186103008501516102608601516101c0870151610120880151608089015118189118188084603f1c611e208660011b67ffffffffffffffff1690565b18188584603f1c611e3b8660011b67ffffffffffffffff1690565b18188584603f1c611e568660011b67ffffffffffffffff1690565b181895508483603f1c611e738560011b67ffffffffffffffff1690565b181894508387603f1c611e908960011b67ffffffffffffffff1690565b60208b01518b51861867ffffffffffffffff168c5291189190911897508118600181901b603f9190911c18935060c08801518118601481901c602c9190911b1867ffffffffffffffff1660208901526101208801518718602c81901c60149190911b1867ffffffffffffffff1660c08901526102c08801518618600381901c603d9190911b1867ffffffffffffffff166101208901526101c08801518718601981901c60279190911b1867ffffffffffffffff166102c08901526102808801518218602e81901c60129190911b1867ffffffffffffffff166101c089015260408801518618600281901c603e9190911b1867ffffffffffffffff166102808901526101808801518618601581901c602b9190911b1867ffffffffffffffff1660408901526101a08801518518602781901c60199190911b1867ffffffffffffffff166101808901526102608801518718603881901c60089190911b1867ffffffffffffffff166101a08901526102e08801518518600881901c60389190911b1867ffffffffffffffff166102608901526101e08801518218601781901c60299190911b1867ffffffffffffffff166102e089015260808801518718602581901c601b9190911b1867ffffffffffffffff166101e08901526103008801518718603281901c600e9190911b1867ffffffffffffffff1660808901526102a08801518118603e81901c60029190911b1867ffffffffffffffff166103008901526101008801518518600981901c60379190911b1867ffffffffffffffff166102a08901526102008801518118601381901c602d9190911b1867ffffffffffffffff1661010089015260a08801518218601c81901c60249190911b1867ffffffffffffffff1661020089015260608801518518602481901c601c9190911b1867ffffffffffffffff1660a08901526102408801518518602b81901c60159190911b1867ffffffffffffffff1660608901526102208801518618603181901c600f9190911b1867ffffffffffffffff166102408901526101608801518118603681901c600a9190911b1867ffffffffffffffff166102208901525060e08701518518603a81901c60069190911b1867ffffffffffffffff166101608801526101408701518118603d81901c60039190911b1867ffffffffffffffff1660e0880152505067ffffffffffffffff81166101408601525050505050565b61221081611d63565b805160208201805160408401805160608601805160808801805167ffffffffffffffff871986168a188116808c528619851689188216909952831982169095188516909552841988169091188316909152941990921618811690925260a08301805160c0808601805160e0880180516101008a0180516101208c018051861985168a188d16909a528319821686188c16909652801989169092188a169092528619861618881690529219909216909218841690526101408401805161016086018051610180880180516101a08a0180516101c08c0180518619851689188d169099528319821686188c16909652801988169092188a169092528519851618881690529119909116909118841690526101e08401805161020086018051610220880180516102408a0180516102608c0180518619851689188d169099528319821686188c16909652801988169092188a16909252851985161888169052911990911690911884169052610280840180516102a0860180516102c0880180516102e08a0180516103008c0180518619851689188d169099528319821686188c16909652801988169092188a16909252851985161888169052911990911690911884169052600386901b850151901c9081189091168252611b9b565b6123f560008284612207565b61240160018284612207565b61240d60028284612207565b61241960038284612207565b61242560048284612207565b61243160058284612207565b61243d60068284612207565b61244960078284612207565b61245560088284612207565b61246160098284612207565b61246d600a8284612207565b612479600b8284612207565b612485600c8284612207565b612491600d8284612207565b61249d600e8284612207565b6124a9600f8284612207565b6124b560108284612207565b6124c160118284612207565b6124cd60128284612207565b6124d960138284612207565b6124e560148284612207565b6124f160158284612207565b6124fd60168284612207565b611b9b60178284612207565b7f01000000000000000000000000000000000000000000000000000000000000007effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8316176125af818360408051600093845233602052918152606090922091527effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f01000000000000000000000000000000000000000000000000000000000000001790565b9392505050565b6060604051905081602082018181018286833760888306808015612614576088829003850160808582017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01536001845160001a1784538652612626565b60018353608060878401536088850186525b5050505050601f19603f82510116810160405292915050565b60801c63ffffffff1690565b60006126ce565b66ff00ff00ff00ff8160081c1667ff00ff00ff00ff0061267c8360081b67ffffffffffffffff1690565b1617905065ffff0000ffff8160101c1667ffff0000ffff00006126a98360101b67ffffffffffffffff1690565b1617905060008160201c6126c78360201b67ffffffffffffffff1690565b1792915050565b608082015160208301906126e690612652565b612652565b60408201516126f490612652565b60401b1761270c6126e160018460059190911b015190565b825160809190911b9061271e90612652565b60c01b17179392505050565b8260108101928215612758579160200282015b8281111561275857825182559160200191906001019061273d565b50612764929150612780565b5090565b604051806020016040528061277b612795565b905290565b5b808211156127645760008155600101612781565b6040518061032001604052806019906020820280368337509192915050565b6000602082840312156127c657600080fd5b5035919050565b803573ffffffffffffffffffffffffffffffffffffffff811681146127f157600080fd5b919050565b6000806040838503121561280957600080fd5b612812836127cd565b946020939093013593505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051610320810167ffffffffffffffff8111828210171561287357612873612820565b60405290565b6040516060810167ffffffffffffffff8111828210171561287357612873612820565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156128e3576128e3612820565b604052919050565b60006103208083850312156128ff57600080fd5b604051602080820167ffffffffffffffff838210818311171561292457612924612820565b8160405283955087601f88011261293a57600080fd5b61294261284f565b948701949150818886111561295657600080fd5b875b8681101561297e57803583811681146129715760008081fd5b8452928401928401612958565b50909352509295945050505050565b60006060828403121561299f57600080fd5b50919050565b60008083601f8401126129b757600080fd5b50813567ffffffffffffffff8111156129cf57600080fd5b6020830191508360208260051b85010111156129ea57600080fd5b9250929050565b60008060008060008060008060006103e08a8c031215612a1057600080fd5b612a198a6127cd565b985060208a01359750612a2f8b60408c016128eb565b96506103608a013567ffffffffffffffff80821115612a4d57600080fd5b612a598d838e0161298d565b97506103808c0135915080821115612a7057600080fd5b612a7c8d838e016129a5565b90975095506103a08c0135915080821115612a9657600080fd5b612aa28d838e0161298d565b94506103c08c0135915080821115612ab957600080fd5b50612ac68c828d016129a5565b915080935050809150509295985092959850929598565b600080600080600060a08688031215612af557600080fd5b505083359560208501359550604085013594606081013594506080013592509050565b60008060408385031215612b2b57600080fd5b50508035926020909101359150565b600080600060608486031215612b4f57600080fd5b612b58846127cd565b95602085013595506040909401359392505050565b60008083601f840112612b7f57600080fd5b50813567ffffffffffffffff811115612b9757600080fd5b6020830191508360208285010111156129ea57600080fd5b600080600080600080600060a0888a031215612bca57600080fd5b8735965060208801359550604088013567ffffffffffffffff80821115612bf057600080fd5b612bfc8b838c01612b6d565b909750955060608a0135915080821115612c1557600080fd5b50612c228a828b01612b6d565b989b979a50959894979596608090950135949350505050565b60008060008060008060808789031215612c5457600080fd5b86359550602087013567ffffffffffffffff80821115612c7357600080fd5b612c7f8a838b01612b6d565b90975095506040890135915080821115612c9857600080fd5b50612ca589828a016129a5565b90945092505060608701358015158114612cbe57600080fd5b809150509295509295509295565b600080600060408486031215612ce157600080fd5b83359250602084013567ffffffffffffffff811115612cff57600080fd5b612d0b86828701612b6d565b9497909650939450505050565b600080600080600060808688031215612d3057600080fd5b612d39866127cd565b945060208601359350604086013567ffffffffffffffff80821115612d5d57600080fd5b612d6989838a0161298d565b94506060880135915080821115612d7f57600080fd5b50612d8c888289016129a5565b969995985093965092949392505050565b803563ffffffff811681146127f157600080fd5b600080600060608486031215612dc657600080fd5b83359250612dd660208501612d9d565b9150612de460408501612d9d565b90509250925092565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203612e7c57612e7c612e1c565b5060010190565b600060608236031215612e9557600080fd5b612e9d612879565b823567ffffffffffffffff80821115612eb557600080fd5b9084019036601f830112612ec857600080fd5b8135602082821115612edc57612edc612820565b612f0c817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160161289c565b92508183523681838601011115612f2257600080fd5b81818501828501376000918301810191909152908352848101359083015250604092830135928101929092525090565b81516103208201908260005b6019811015612f8757825167ffffffffffffffff16825260209283019290910190600101612f5e565b50505092915050565b60008219821115612fa357612fa3612e1c565b500190565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112612fdd57600080fd5b83018035915067ffffffffffffffff821115612ff857600080fd5b6020019150368190038213156129ea57600080fd5b600181815b8085111561306657817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0482111561304c5761304c612e1c565b8085161561305957918102915b93841c9390800290613012565b509250929050565b60008261307d57506001613129565b8161308a57506000613129565b81600181146130a057600281146130aa576130c6565b6001915050613129565b60ff8411156130bb576130bb612e1c565b50506001821b613129565b5060208310610133831016604e8410600b84101617156130e9575081810a613129565b6130f3838361300d565b807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0482111561312557613125612e1c565b0290505b92915050565b60006125af838361306e565b60008282101561314d5761314d612e1c565b500390565b600063ffffffff8381169083168181101561316f5761316f612e1c565b039392505050565b600063ffffffff80831681851680830382111561319657613196612e1c565b01949350505050565b6000845160005b818110156131c057602081880181015185830152016131a6565b818111156131cf576000828501525b509190910192835250602082015260400191905056fea164736f6c634300080f000a"


func init() {
	if err := json.Unmarshal([]byte(PreimageOracleStorageLayoutJSON), PreimageOracleStorageLayout); err != nil {
		panic(err)
	}

	layouts["PreimageOracle"] = PreimageOracleStorageLayout
	deployedBytecodes["PreimageOracle"] = PreimageOracleDeployedBin
	immutableReferences["PreimageOracle"] = true
}
