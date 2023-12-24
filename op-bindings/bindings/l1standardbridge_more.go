// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const L1StandardBridgeStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/L1/L1StandardBridge.sol:L1StandardBridge\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"src/L1/L1StandardBridge.sol:L1StandardBridge\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"src/L1/L1StandardBridge.sol:L1StandardBridge\",\"label\":\"spacer_0_2_20\",\"offset\":2,\"slot\":\"0\",\"type\":\"t_address\"},{\"astId\":1003,\"contract\":\"src/L1/L1StandardBridge.sol:L1StandardBridge\",\"label\":\"spacer_1_0_20\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_address\"},{\"astId\":1004,\"contract\":\"src/L1/L1StandardBridge.sol:L1StandardBridge\",\"label\":\"deposits\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_mapping(t_address,t_mapping(t_address,t_uint256))\"},{\"astId\":1005,\"contract\":\"src/L1/L1StandardBridge.sol:L1StandardBridge\",\"label\":\"messenger\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_contract(CrossDomainMessenger)1009\"},{\"astId\":1006,\"contract\":\"src/L1/L1StandardBridge.sol:L1StandardBridge\",\"label\":\"otherBridge\",\"offset\":0,\"slot\":\"4\",\"type\":\"t_contract(StandardBridge)1010\"},{\"astId\":1007,\"contract\":\"src/L1/L1StandardBridge.sol:L1StandardBridge\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"5\",\"type\":\"t_array(t_uint256)45_storage\"},{\"astId\":1008,\"contract\":\"src/L1/L1StandardBridge.sol:L1StandardBridge\",\"label\":\"superchainConfig\",\"offset\":0,\"slot\":\"50\",\"type\":\"t_contract(SuperchainConfig)1011\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)45_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[45]\",\"numberOfBytes\":\"1440\",\"base\":\"t_uint256\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_contract(CrossDomainMessenger)1009\":{\"encoding\":\"inplace\",\"label\":\"contract CrossDomainMessenger\",\"numberOfBytes\":\"20\"},\"t_contract(StandardBridge)1010\":{\"encoding\":\"inplace\",\"label\":\"contract StandardBridge\",\"numberOfBytes\":\"20\"},\"t_contract(SuperchainConfig)1011\":{\"encoding\":\"inplace\",\"label\":\"contract SuperchainConfig\",\"numberOfBytes\":\"20\"},\"t_mapping(t_address,t_mapping(t_address,t_uint256))\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e mapping(address =\u003e uint256))\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_mapping(t_address,t_uint256)\"},\"t_mapping(t_address,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_uint256\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var L1StandardBridgeStorageLayout = new(solc.StorageLayout)

var L1StandardBridgeDeployedBin = "0x6080604052600436106101795760003560e01c8063838b2520116100cb5780639a2ac6d51161007f578063c0c53b8b11610059578063c0c53b8b146104f1578063c89701a214610511578063e11013dd1461053e57600080fd5b80639a2ac6d5146104ab578063a9f9e675146104be578063b1a1a882146104de57600080fd5b80638f601f66116100b05780638f601f661461043a57806391c49bf8146103cf578063927ede2d1461048057600080fd5b8063838b2520146103fa578063870876231461041a57600080fd5b80633cb747bf1161012d57806358a997f61161010757806358a997f61461038a5780635c975abb146103aa5780637f46ddb2146103cf57600080fd5b80633cb747bf146102e7578063540abf731461031457806354fd4d501461033457600080fd5b80631532ec341161015e5780631532ec341461026a5780631635f5fd1461027d57806335e80ab31461029057600080fd5b80630166a07a1461023757806309fc88431461025757600080fd5b3661023257333b15610212576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f4100000000000000000060648201526084015b60405180910390fd5b610230333362030d406040518060200160405280600081525061056d565b005b600080fd5b34801561024357600080fd5b5061023061025236600461269e565b610580565b61023061026536600461274f565b61099a565b6102306102783660046127a2565b610a71565b61023061028b3660046127a2565b610a85565b34801561029c57600080fd5b506032546102bd9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b3480156102f357600080fd5b506003546102bd9073ffffffffffffffffffffffffffffffffffffffff1681565b34801561032057600080fd5b5061023061032f366004612815565b610f4e565b34801561034057600080fd5b5061037d6040518060400160405280600581526020017f322e312e3000000000000000000000000000000000000000000000000000000081525081565b6040516102de9190612902565b34801561039657600080fd5b506102306103a5366004612915565b610f93565b3480156103b657600080fd5b506103bf611067565b60405190151581526020016102de565b3480156103db57600080fd5b5060045473ffffffffffffffffffffffffffffffffffffffff166102bd565b34801561040657600080fd5b50610230610415366004612815565b611100565b34801561042657600080fd5b50610230610435366004612915565b611145565b34801561044657600080fd5b50610472610455366004612998565b600260209081526000928352604080842090915290825290205481565b6040519081526020016102de565b34801561048c57600080fd5b5060035473ffffffffffffffffffffffffffffffffffffffff166102bd565b6102306104b93660046129d1565b611219565b3480156104ca57600080fd5b506102306104d936600461269e565b61125b565b6102306104ec36600461274f565b61126a565b3480156104fd57600080fd5b5061023061050c366004612a34565b61133b565b34801561051d57600080fd5b506004546102bd9073ffffffffffffffffffffffffffffffffffffffff1681565b61023061054c3660046129d1565b611511565b73ffffffffffffffffffffffffffffffffffffffff163b151590565b61057a8484348585611554565b50505050565b60035473ffffffffffffffffffffffffffffffffffffffff1633148015610653575060048054600354604080517f6e296e45000000000000000000000000000000000000000000000000000000008152905173ffffffffffffffffffffffffffffffffffffffff938416949390921692636e296e459282820192602092908290030181865afa158015610617573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061063b9190612a7f565b73ffffffffffffffffffffffffffffffffffffffff16145b610705576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604160248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20746865206f7468657220627269646760648201527f6500000000000000000000000000000000000000000000000000000000000000608482015260a401610209565b61070d611067565b15610774576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f5374616e646172644272696467653a20706175736564000000000000000000006044820152606401610209565b61077d8761171e565b156108cb5761078c8787611780565b61083e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604a60248201527f5374616e646172644272696467653a2077726f6e672072656d6f746520746f6b60448201527f656e20666f72204f7074696d69736d204d696e7461626c65204552433230206c60648201527f6f63616c20746f6b656e00000000000000000000000000000000000000000000608482015260a401610209565b6040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8581166004830152602482018590528816906340c10f1990604401600060405180830381600087803b1580156108ae57600080fd5b505af11580156108c2573d6000803e3d6000fd5b5050505061094d565b73ffffffffffffffffffffffffffffffffffffffff8088166000908152600260209081526040808320938a1683529290522054610909908490612acb565b73ffffffffffffffffffffffffffffffffffffffff8089166000818152600260209081526040808320948c168352939052919091209190915561094d9085856118a0565b610991878787878787878080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061197492505050565b50505050505050565b333b15610a29576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f410000000000000000006064820152608401610209565b610a6c3333348686868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061155492505050565b505050565b610a7e8585858585610a85565b5050505050565b60035473ffffffffffffffffffffffffffffffffffffffff1633148015610b58575060048054600354604080517f6e296e45000000000000000000000000000000000000000000000000000000008152905173ffffffffffffffffffffffffffffffffffffffff938416949390921692636e296e459282820192602092908290030181865afa158015610b1c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b409190612a7f565b73ffffffffffffffffffffffffffffffffffffffff16145b610c0a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604160248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20746865206f7468657220627269646760648201527f6500000000000000000000000000000000000000000000000000000000000000608482015260a401610209565b610c12611067565b15610c79576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f5374616e646172644272696467653a20706175736564000000000000000000006044820152606401610209565b823414610d08576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603a60248201527f5374616e646172644272696467653a20616d6f756e742073656e7420646f657360448201527f206e6f74206d6174636820616d6f756e742072657175697265640000000000006064820152608401610209565b3073ffffffffffffffffffffffffffffffffffffffff851603610dad576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f5374616e646172644272696467653a2063616e6e6f742073656e6420746f207360448201527f656c6600000000000000000000000000000000000000000000000000000000006064820152608401610209565b60035473ffffffffffffffffffffffffffffffffffffffff90811690851603610e58576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602860248201527f5374616e646172644272696467653a2063616e6e6f742073656e6420746f206d60448201527f657373656e6765720000000000000000000000000000000000000000000000006064820152608401610209565b610e9a85858585858080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611a0292505050565b6000610eb7855a8660405180602001604052806000815250611a75565b905080610f46576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f5374616e646172644272696467653a20455448207472616e736665722066616960448201527f6c656400000000000000000000000000000000000000000000000000000000006064820152608401610209565b505050505050565b61099187873388888888888080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611a8f92505050565b333b15611022576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f410000000000000000006064820152608401610209565b610f4686863333888888888080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611dba92505050565b603254604080517f5c975abb000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff1691635c975abb9160048083019260209291908290030181865afa1580156110d7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110fb9190612ae2565b905090565b61099187873388888888888080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611dba92505050565b333b156111d4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f410000000000000000006064820152608401610209565b610f4686863333888888888080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611a8f92505050565b61057a33858585858080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061056d92505050565b61099187878787878787610580565b333b156112f9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f410000000000000000006064820152608401610209565b610a6c33338585858080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061056d92505050565b600054610100900460ff161580801561135b5750600054600160ff909116105b806113755750303b158015611375575060005460ff166001145b611401576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610209565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561145f57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b603280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff84161790556114a98484611dc9565b801561057a57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a150505050565b61057a3385348686868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061155492505050565b8234146115e3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603e60248201527f5374616e646172644272696467653a206272696467696e6720455448206d757360448201527f7420696e636c7564652073756666696369656e74204554482076616c756500006064820152608401610209565b6115ef85858584611eb3565b60035460045460405173ffffffffffffffffffffffffffffffffffffffff92831692633dbb202b9287929116907f1635f5fd0000000000000000000000000000000000000000000000000000000090611652908b908b9086908a90602401612b04565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009485161790525160e086901b90921682526116e592918890600401612b4d565b6000604051808303818588803b1580156116fe57600080fd5b505af1158015611712573d6000803e3d6000fd5b50505050505050505050565b600061174a827f1d1d8b6300000000000000000000000000000000000000000000000000000000611f26565b8061177a575061177a827fec4fc8e300000000000000000000000000000000000000000000000000000000611f26565b92915050565b60006117ac837f1d1d8b6300000000000000000000000000000000000000000000000000000000611f26565b15611855578273ffffffffffffffffffffffffffffffffffffffff1663c01e1bd66040518163ffffffff1660e01b8152600401602060405180830381865afa1580156117fc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118209190612a7f565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614905061177a565b8273ffffffffffffffffffffffffffffffffffffffff1663d6c0b2c46040518163ffffffff1660e01b8152600401602060405180830381865afa1580156117fc573d6000803e3d6000fd5b60405173ffffffffffffffffffffffffffffffffffffffff8316602482015260448101829052610a6c9084907fa9059cbb00000000000000000000000000000000000000000000000000000000906064015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152611f49565b8373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167f3ceee06c1e37648fcbb6ed52e17b3e1f275a1f8c7b22a84b2b84732431e046b38686866040516119ec93929190612b92565b60405180910390a4610f46868686868686612055565b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f2ac69ee804d9a7a0984249f508dfab7cb2534b465b6ce1580f99a38ba9c5e6318484604051611a61929190612bd0565b60405180910390a361057a848484846120dd565b600080600080845160208601878a8af19695505050505050565b611a988761171e565b15611be657611aa78787611780565b611b59576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604a60248201527f5374616e646172644272696467653a2077726f6e672072656d6f746520746f6b60448201527f656e20666f72204f7074696d69736d204d696e7461626c65204552433230206c60648201527f6f63616c20746f6b656e00000000000000000000000000000000000000000000608482015260a401610209565b6040517f9dc29fac00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff868116600483015260248201859052881690639dc29fac90604401600060405180830381600087803b158015611bc957600080fd5b505af1158015611bdd573d6000803e3d6000fd5b50505050611c7a565b611c0873ffffffffffffffffffffffffffffffffffffffff881686308661214a565b73ffffffffffffffffffffffffffffffffffffffff8088166000908152600260209081526040808320938a1683529290522054611c46908490612be9565b73ffffffffffffffffffffffffffffffffffffffff8089166000908152600260209081526040808320938b16835292905220555b611c888787878787866121a8565b60035460045460405173ffffffffffffffffffffffffffffffffffffffff92831692633dbb202b9216907f0166a07a0000000000000000000000000000000000000000000000000000000090611cec908b908d908c908c908c908b90602401612c01565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009485161790525160e085901b9092168252611d7f92918790600401612b4d565b600060405180830381600087803b158015611d9957600080fd5b505af1158015611dad573d6000803e3d6000fd5b5050505050505050505050565b61099187878787878787611a8f565b600054610100900460ff16611e60576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610209565b6003805473ffffffffffffffffffffffffffffffffffffffff9384167fffffffffffffffffffffffff00000000000000000000000000000000000000009182161790915560048054929093169116179055565b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f35d79ab81f2b2017e19afb5c5571778877782d7a8786f5907f93b0f4702f4f238484604051611f12929190612bd0565b60405180910390a361057a84848484612236565b6000611f3183612295565b8015611f425750611f4283836122f9565b9392505050565b6000611fab826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166123c89092919063ffffffff16565b805190915015610a6c5780806020019051810190611fc99190612ae2565b610a6c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610209565b8373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167fd59c65b35445225835c83f50b6ede06a7be047d22e357073e250d9af537518cd8686866040516120cd93929190612b92565b60405180910390a4505050505050565b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f31b2166ff604fc5672ea5df08a78081d2bc6d746cadce880747f3643d819e83d848460405161213c929190612bd0565b60405180910390a350505050565b60405173ffffffffffffffffffffffffffffffffffffffff8085166024830152831660448201526064810182905261057a9085907f23b872dd00000000000000000000000000000000000000000000000000000000906084016118f2565b8373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167f718594027abd4eaed59f95162563e0cc6d0e8d5b86b1c7be8b1b0ac3343d039686868660405161222093929190612b92565b60405180910390a4610f468686868686866123df565b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f2849b43074093a05396b6f2a937dee8565b15a48a7b3d4bffb732a5017380af5848460405161213c929190612bd0565b60006122c1827f01ffc9a7000000000000000000000000000000000000000000000000000000006122f9565b801561177a57506122f2827fffffffff000000000000000000000000000000000000000000000000000000006122f9565b1592915050565b604080517fffffffff000000000000000000000000000000000000000000000000000000008316602480830191909152825180830390910181526044909101909152602080820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f01ffc9a700000000000000000000000000000000000000000000000000000000178152825160009392849283928392918391908a617530fa92503d915060005190508280156123b1575060208210155b80156123bd5750600081115b979650505050505050565b60606123d78484600085612457565b949350505050565b8373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167f7ff126db8024424bbfd9826e8ab82ff59136289ea440b04b39a0df1b03b9cabf8686866040516120cd93929190612b92565b6060824710156124e9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610209565b73ffffffffffffffffffffffffffffffffffffffff85163b612567576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610209565b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040516125909190612c5c565b60006040518083038185875af1925050503d80600081146125cd576040519150601f19603f3d011682016040523d82523d6000602084013e6125d2565b606091505b50915091506123bd828286606083156125ec575081611f42565b8251156125fc5782518084602001fd5b816040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102099190612902565b73ffffffffffffffffffffffffffffffffffffffff8116811461265257600080fd5b50565b60008083601f84011261266757600080fd5b50813567ffffffffffffffff81111561267f57600080fd5b60208301915083602082850101111561269757600080fd5b9250929050565b600080600080600080600060c0888a0312156126b957600080fd5b87356126c481612630565b965060208801356126d481612630565b955060408801356126e481612630565b945060608801356126f481612630565b93506080880135925060a088013567ffffffffffffffff81111561271757600080fd5b6127238a828b01612655565b989b979a50959850939692959293505050565b803563ffffffff8116811461274a57600080fd5b919050565b60008060006040848603121561276457600080fd5b61276d84612736565b9250602084013567ffffffffffffffff81111561278957600080fd5b61279586828701612655565b9497909650939450505050565b6000806000806000608086880312156127ba57600080fd5b85356127c581612630565b945060208601356127d581612630565b935060408601359250606086013567ffffffffffffffff8111156127f857600080fd5b61280488828901612655565b969995985093965092949392505050565b600080600080600080600060c0888a03121561283057600080fd5b873561283b81612630565b9650602088013561284b81612630565b9550604088013561285b81612630565b94506060880135935061287060808901612736565b925060a088013567ffffffffffffffff81111561271757600080fd5b60005b838110156128a757818101518382015260200161288f565b8381111561057a5750506000910152565b600081518084526128d081602086016020860161288c565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000611f4260208301846128b8565b60008060008060008060a0878903121561292e57600080fd5b863561293981612630565b9550602087013561294981612630565b94506040870135935061295e60608801612736565b9250608087013567ffffffffffffffff81111561297a57600080fd5b61298689828a01612655565b979a9699509497509295939492505050565b600080604083850312156129ab57600080fd5b82356129b681612630565b915060208301356129c681612630565b809150509250929050565b600080600080606085870312156129e757600080fd5b84356129f281612630565b9350612a0060208601612736565b9250604085013567ffffffffffffffff811115612a1c57600080fd5b612a2887828801612655565b95989497509550505050565b600080600060608486031215612a4957600080fd5b8335612a5481612630565b92506020840135612a6481612630565b91506040840135612a7481612630565b809150509250925092565b600060208284031215612a9157600080fd5b8151611f4281612630565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082821015612add57612add612a9c565b500390565b600060208284031215612af457600080fd5b81518015158114611f4257600080fd5b600073ffffffffffffffffffffffffffffffffffffffff808716835280861660208401525083604083015260806060830152612b4360808301846128b8565b9695505050505050565b73ffffffffffffffffffffffffffffffffffffffff84168152606060208201526000612b7c60608301856128b8565b905063ffffffff83166040830152949350505050565b73ffffffffffffffffffffffffffffffffffffffff84168152826020820152606060408201526000612bc760608301846128b8565b95945050505050565b8281526040602082015260006123d760408301846128b8565b60008219821115612bfc57612bfc612a9c565b500190565b600073ffffffffffffffffffffffffffffffffffffffff80891683528088166020840152808716604084015280861660608401525083608083015260c060a0830152612c5060c08301846128b8565b98975050505050505050565b60008251612c6e81846020870161288c565b919091019291505056fea164736f6c634300080f000a"


func init() {
	if err := json.Unmarshal([]byte(L1StandardBridgeStorageLayoutJSON), L1StandardBridgeStorageLayout); err != nil {
		panic(err)
	}

	layouts["L1StandardBridge"] = L1StandardBridgeStorageLayout
	deployedBytecodes["L1StandardBridge"] = L1StandardBridgeDeployedBin
	immutableReferences["L1StandardBridge"] = false
}
