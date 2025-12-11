// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package trades

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// TradesTrade is an auto generated low-level Go binding around an user-defined struct.
type TradesTrade struct {
	PlayerA   common.Address
	CardA     *big.Int
	AcceptedA bool
	PlayerB   common.Address
	CardB     *big.Int
	AcceptedB bool
	Complete  bool
}

// TradesMetaData contains all meta data concerning the Trades contract.
var TradesMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cardsContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tradeId\",\"type\":\"uint256\"}],\"name\":\"acceptOffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tradeId\",\"type\":\"uint256\"}],\"name\":\"cancelOffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cards\",\"outputs\":[{\"internalType\":\"contractICards\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cardA\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"playerB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cardB\",\"type\":\"uint256\"}],\"name\":\"createOffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tradeId\",\"type\":\"uint256\"}],\"name\":\"getTrade\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"playerA\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cardA\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"acceptedA\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"playerB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cardB\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"acceptedB\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"complete\",\"type\":\"bool\"}],\"internalType\":\"structTrades.Trade\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextTradeId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"trades\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"playerA\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cardA\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"acceptedA\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"playerB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cardB\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"acceptedB\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"complete\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// TradesABI is the input ABI used to generate the binding from.
// Deprecated: Use TradesMetaData.ABI instead.
var TradesABI = TradesMetaData.ABI

// Trades is an auto generated Go binding around an Ethereum contract.
type Trades struct {
	TradesCaller     // Read-only binding to the contract
	TradesTransactor // Write-only binding to the contract
	TradesFilterer   // Log filterer for contract events
}

// TradesCaller is an auto generated read-only Go binding around an Ethereum contract.
type TradesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TradesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TradesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TradesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TradesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TradesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TradesSession struct {
	Contract     *Trades           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TradesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TradesCallerSession struct {
	Contract *TradesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TradesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TradesTransactorSession struct {
	Contract     *TradesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TradesRaw is an auto generated low-level Go binding around an Ethereum contract.
type TradesRaw struct {
	Contract *Trades // Generic contract binding to access the raw methods on
}

// TradesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TradesCallerRaw struct {
	Contract *TradesCaller // Generic read-only contract binding to access the raw methods on
}

// TradesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TradesTransactorRaw struct {
	Contract *TradesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTrades creates a new instance of Trades, bound to a specific deployed contract.
func NewTrades(address common.Address, backend bind.ContractBackend) (*Trades, error) {
	contract, err := bindTrades(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Trades{TradesCaller: TradesCaller{contract: contract}, TradesTransactor: TradesTransactor{contract: contract}, TradesFilterer: TradesFilterer{contract: contract}}, nil
}

// NewTradesCaller creates a new read-only instance of Trades, bound to a specific deployed contract.
func NewTradesCaller(address common.Address, caller bind.ContractCaller) (*TradesCaller, error) {
	contract, err := bindTrades(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TradesCaller{contract: contract}, nil
}

// NewTradesTransactor creates a new write-only instance of Trades, bound to a specific deployed contract.
func NewTradesTransactor(address common.Address, transactor bind.ContractTransactor) (*TradesTransactor, error) {
	contract, err := bindTrades(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TradesTransactor{contract: contract}, nil
}

// NewTradesFilterer creates a new log filterer instance of Trades, bound to a specific deployed contract.
func NewTradesFilterer(address common.Address, filterer bind.ContractFilterer) (*TradesFilterer, error) {
	contract, err := bindTrades(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TradesFilterer{contract: contract}, nil
}

// bindTrades binds a generic wrapper to an already deployed contract.
func bindTrades(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TradesMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Trades *TradesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Trades.Contract.TradesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Trades *TradesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trades.Contract.TradesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Trades *TradesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Trades.Contract.TradesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Trades *TradesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Trades.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Trades *TradesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trades.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Trades *TradesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Trades.Contract.contract.Transact(opts, method, params...)
}

// Cards is a free data retrieval call binding the contract method 0x58a4903f.
//
// Solidity: function cards() view returns(address)
func (_Trades *TradesCaller) Cards(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Trades.contract.Call(opts, &out, "cards")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Cards is a free data retrieval call binding the contract method 0x58a4903f.
//
// Solidity: function cards() view returns(address)
func (_Trades *TradesSession) Cards() (common.Address, error) {
	return _Trades.Contract.Cards(&_Trades.CallOpts)
}

// Cards is a free data retrieval call binding the contract method 0x58a4903f.
//
// Solidity: function cards() view returns(address)
func (_Trades *TradesCallerSession) Cards() (common.Address, error) {
	return _Trades.Contract.Cards(&_Trades.CallOpts)
}

// GetTrade is a free data retrieval call binding the contract method 0x2db25e05.
//
// Solidity: function getTrade(uint256 tradeId) view returns((address,uint256,bool,address,uint256,bool,bool))
func (_Trades *TradesCaller) GetTrade(opts *bind.CallOpts, tradeId *big.Int) (TradesTrade, error) {
	var out []interface{}
	err := _Trades.contract.Call(opts, &out, "getTrade", tradeId)

	if err != nil {
		return *new(TradesTrade), err
	}

	out0 := *abi.ConvertType(out[0], new(TradesTrade)).(*TradesTrade)

	return out0, err

}

// GetTrade is a free data retrieval call binding the contract method 0x2db25e05.
//
// Solidity: function getTrade(uint256 tradeId) view returns((address,uint256,bool,address,uint256,bool,bool))
func (_Trades *TradesSession) GetTrade(tradeId *big.Int) (TradesTrade, error) {
	return _Trades.Contract.GetTrade(&_Trades.CallOpts, tradeId)
}

// GetTrade is a free data retrieval call binding the contract method 0x2db25e05.
//
// Solidity: function getTrade(uint256 tradeId) view returns((address,uint256,bool,address,uint256,bool,bool))
func (_Trades *TradesCallerSession) GetTrade(tradeId *big.Int) (TradesTrade, error) {
	return _Trades.Contract.GetTrade(&_Trades.CallOpts, tradeId)
}

// NextTradeId is a free data retrieval call binding the contract method 0x813ad083.
//
// Solidity: function nextTradeId() view returns(uint256)
func (_Trades *TradesCaller) NextTradeId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Trades.contract.Call(opts, &out, "nextTradeId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextTradeId is a free data retrieval call binding the contract method 0x813ad083.
//
// Solidity: function nextTradeId() view returns(uint256)
func (_Trades *TradesSession) NextTradeId() (*big.Int, error) {
	return _Trades.Contract.NextTradeId(&_Trades.CallOpts)
}

// NextTradeId is a free data retrieval call binding the contract method 0x813ad083.
//
// Solidity: function nextTradeId() view returns(uint256)
func (_Trades *TradesCallerSession) NextTradeId() (*big.Int, error) {
	return _Trades.Contract.NextTradeId(&_Trades.CallOpts)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address operator, address from, uint256[] ids, uint256[] values, bytes data) pure returns(bytes4)
func (_Trades *TradesCaller) OnERC1155BatchReceived(opts *bind.CallOpts, operator common.Address, from common.Address, ids []*big.Int, values []*big.Int, data []byte) ([4]byte, error) {
	var out []interface{}
	err := _Trades.contract.Call(opts, &out, "onERC1155BatchReceived", operator, from, ids, values, data)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address operator, address from, uint256[] ids, uint256[] values, bytes data) pure returns(bytes4)
func (_Trades *TradesSession) OnERC1155BatchReceived(operator common.Address, from common.Address, ids []*big.Int, values []*big.Int, data []byte) ([4]byte, error) {
	return _Trades.Contract.OnERC1155BatchReceived(&_Trades.CallOpts, operator, from, ids, values, data)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address operator, address from, uint256[] ids, uint256[] values, bytes data) pure returns(bytes4)
func (_Trades *TradesCallerSession) OnERC1155BatchReceived(operator common.Address, from common.Address, ids []*big.Int, values []*big.Int, data []byte) ([4]byte, error) {
	return _Trades.Contract.OnERC1155BatchReceived(&_Trades.CallOpts, operator, from, ids, values, data)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address operator, address from, uint256 id, uint256 value, bytes data) pure returns(bytes4)
func (_Trades *TradesCaller) OnERC1155Received(opts *bind.CallOpts, operator common.Address, from common.Address, id *big.Int, value *big.Int, data []byte) ([4]byte, error) {
	var out []interface{}
	err := _Trades.contract.Call(opts, &out, "onERC1155Received", operator, from, id, value, data)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address operator, address from, uint256 id, uint256 value, bytes data) pure returns(bytes4)
func (_Trades *TradesSession) OnERC1155Received(operator common.Address, from common.Address, id *big.Int, value *big.Int, data []byte) ([4]byte, error) {
	return _Trades.Contract.OnERC1155Received(&_Trades.CallOpts, operator, from, id, value, data)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address operator, address from, uint256 id, uint256 value, bytes data) pure returns(bytes4)
func (_Trades *TradesCallerSession) OnERC1155Received(operator common.Address, from common.Address, id *big.Int, value *big.Int, data []byte) ([4]byte, error) {
	return _Trades.Contract.OnERC1155Received(&_Trades.CallOpts, operator, from, id, value, data)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Trades *TradesCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Trades.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Trades *TradesSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Trades.Contract.SupportsInterface(&_Trades.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Trades *TradesCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Trades.Contract.SupportsInterface(&_Trades.CallOpts, interfaceId)
}

// Trades is a free data retrieval call binding the contract method 0x1e6c598e.
//
// Solidity: function trades(uint256 ) view returns(address playerA, uint256 cardA, bool acceptedA, address playerB, uint256 cardB, bool acceptedB, bool complete)
func (_Trades *TradesCaller) Trades(opts *bind.CallOpts, arg0 *big.Int) (struct {
	PlayerA   common.Address
	CardA     *big.Int
	AcceptedA bool
	PlayerB   common.Address
	CardB     *big.Int
	AcceptedB bool
	Complete  bool
}, error) {
	var out []interface{}
	err := _Trades.contract.Call(opts, &out, "trades", arg0)

	outstruct := new(struct {
		PlayerA   common.Address
		CardA     *big.Int
		AcceptedA bool
		PlayerB   common.Address
		CardB     *big.Int
		AcceptedB bool
		Complete  bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PlayerA = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.CardA = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.AcceptedA = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.PlayerB = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.CardB = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.AcceptedB = *abi.ConvertType(out[5], new(bool)).(*bool)
	outstruct.Complete = *abi.ConvertType(out[6], new(bool)).(*bool)

	return *outstruct, err

}

// Trades is a free data retrieval call binding the contract method 0x1e6c598e.
//
// Solidity: function trades(uint256 ) view returns(address playerA, uint256 cardA, bool acceptedA, address playerB, uint256 cardB, bool acceptedB, bool complete)
func (_Trades *TradesSession) Trades(arg0 *big.Int) (struct {
	PlayerA   common.Address
	CardA     *big.Int
	AcceptedA bool
	PlayerB   common.Address
	CardB     *big.Int
	AcceptedB bool
	Complete  bool
}, error) {
	return _Trades.Contract.Trades(&_Trades.CallOpts, arg0)
}

// Trades is a free data retrieval call binding the contract method 0x1e6c598e.
//
// Solidity: function trades(uint256 ) view returns(address playerA, uint256 cardA, bool acceptedA, address playerB, uint256 cardB, bool acceptedB, bool complete)
func (_Trades *TradesCallerSession) Trades(arg0 *big.Int) (struct {
	PlayerA   common.Address
	CardA     *big.Int
	AcceptedA bool
	PlayerB   common.Address
	CardB     *big.Int
	AcceptedB bool
	Complete  bool
}, error) {
	return _Trades.Contract.Trades(&_Trades.CallOpts, arg0)
}

// AcceptOffer is a paid mutator transaction binding the contract method 0xc815729d.
//
// Solidity: function acceptOffer(uint256 tradeId) returns()
func (_Trades *TradesTransactor) AcceptOffer(opts *bind.TransactOpts, tradeId *big.Int) (*types.Transaction, error) {
	return _Trades.contract.Transact(opts, "acceptOffer", tradeId)
}

// AcceptOffer is a paid mutator transaction binding the contract method 0xc815729d.
//
// Solidity: function acceptOffer(uint256 tradeId) returns()
func (_Trades *TradesSession) AcceptOffer(tradeId *big.Int) (*types.Transaction, error) {
	return _Trades.Contract.AcceptOffer(&_Trades.TransactOpts, tradeId)
}

// AcceptOffer is a paid mutator transaction binding the contract method 0xc815729d.
//
// Solidity: function acceptOffer(uint256 tradeId) returns()
func (_Trades *TradesTransactorSession) AcceptOffer(tradeId *big.Int) (*types.Transaction, error) {
	return _Trades.Contract.AcceptOffer(&_Trades.TransactOpts, tradeId)
}

// CancelOffer is a paid mutator transaction binding the contract method 0xef706adf.
//
// Solidity: function cancelOffer(uint256 tradeId) returns()
func (_Trades *TradesTransactor) CancelOffer(opts *bind.TransactOpts, tradeId *big.Int) (*types.Transaction, error) {
	return _Trades.contract.Transact(opts, "cancelOffer", tradeId)
}

// CancelOffer is a paid mutator transaction binding the contract method 0xef706adf.
//
// Solidity: function cancelOffer(uint256 tradeId) returns()
func (_Trades *TradesSession) CancelOffer(tradeId *big.Int) (*types.Transaction, error) {
	return _Trades.Contract.CancelOffer(&_Trades.TransactOpts, tradeId)
}

// CancelOffer is a paid mutator transaction binding the contract method 0xef706adf.
//
// Solidity: function cancelOffer(uint256 tradeId) returns()
func (_Trades *TradesTransactorSession) CancelOffer(tradeId *big.Int) (*types.Transaction, error) {
	return _Trades.Contract.CancelOffer(&_Trades.TransactOpts, tradeId)
}

// CreateOffer is a paid mutator transaction binding the contract method 0x9b59d95b.
//
// Solidity: function createOffer(uint256 cardA, address playerB, uint256 cardB) returns()
func (_Trades *TradesTransactor) CreateOffer(opts *bind.TransactOpts, cardA *big.Int, playerB common.Address, cardB *big.Int) (*types.Transaction, error) {
	return _Trades.contract.Transact(opts, "createOffer", cardA, playerB, cardB)
}

// CreateOffer is a paid mutator transaction binding the contract method 0x9b59d95b.
//
// Solidity: function createOffer(uint256 cardA, address playerB, uint256 cardB) returns()
func (_Trades *TradesSession) CreateOffer(cardA *big.Int, playerB common.Address, cardB *big.Int) (*types.Transaction, error) {
	return _Trades.Contract.CreateOffer(&_Trades.TransactOpts, cardA, playerB, cardB)
}

// CreateOffer is a paid mutator transaction binding the contract method 0x9b59d95b.
//
// Solidity: function createOffer(uint256 cardA, address playerB, uint256 cardB) returns()
func (_Trades *TradesTransactorSession) CreateOffer(cardA *big.Int, playerB common.Address, cardB *big.Int) (*types.Transaction, error) {
	return _Trades.Contract.CreateOffer(&_Trades.TransactOpts, cardA, playerB, cardB)
}
