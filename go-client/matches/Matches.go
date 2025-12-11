// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package matches

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

// MatchesMatch is an auto generated low-level Go binding around an user-defined struct.
type MatchesMatch struct {
	PlayerA  common.Address
	CardA    *big.Int
	PlayerB  common.Address
	CardB    *big.Int
	Finished bool
	Winner   common.Address
}

// MatchesMetaData contains all meta data concerning the Matches contract.
var MatchesMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cardsAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"cards\",\"outputs\":[{\"internalType\":\"contractICards\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enqueue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"matchId\",\"type\":\"uint256\"}],\"name\":\"getMatch\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"playerA\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cardA\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"playerB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cardB\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"finished\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"}],\"internalType\":\"structMatches.Match\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isWaiting\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"matches\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"playerA\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cardA\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"playerB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cardB\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"finished\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextMatchId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"matchId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cardId\",\"type\":\"uint256\"}],\"name\":\"playCard\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"waitingPlayer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MatchesABI is the input ABI used to generate the binding from.
// Deprecated: Use MatchesMetaData.ABI instead.
var MatchesABI = MatchesMetaData.ABI

// Matches is an auto generated Go binding around an Ethereum contract.
type Matches struct {
	MatchesCaller     // Read-only binding to the contract
	MatchesTransactor // Write-only binding to the contract
	MatchesFilterer   // Log filterer for contract events
}

// MatchesCaller is an auto generated read-only Go binding around an Ethereum contract.
type MatchesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MatchesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MatchesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MatchesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MatchesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MatchesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MatchesSession struct {
	Contract     *Matches          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MatchesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MatchesCallerSession struct {
	Contract *MatchesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// MatchesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MatchesTransactorSession struct {
	Contract     *MatchesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MatchesRaw is an auto generated low-level Go binding around an Ethereum contract.
type MatchesRaw struct {
	Contract *Matches // Generic contract binding to access the raw methods on
}

// MatchesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MatchesCallerRaw struct {
	Contract *MatchesCaller // Generic read-only contract binding to access the raw methods on
}

// MatchesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MatchesTransactorRaw struct {
	Contract *MatchesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMatches creates a new instance of Matches, bound to a specific deployed contract.
func NewMatches(address common.Address, backend bind.ContractBackend) (*Matches, error) {
	contract, err := bindMatches(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Matches{MatchesCaller: MatchesCaller{contract: contract}, MatchesTransactor: MatchesTransactor{contract: contract}, MatchesFilterer: MatchesFilterer{contract: contract}}, nil
}

// NewMatchesCaller creates a new read-only instance of Matches, bound to a specific deployed contract.
func NewMatchesCaller(address common.Address, caller bind.ContractCaller) (*MatchesCaller, error) {
	contract, err := bindMatches(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MatchesCaller{contract: contract}, nil
}

// NewMatchesTransactor creates a new write-only instance of Matches, bound to a specific deployed contract.
func NewMatchesTransactor(address common.Address, transactor bind.ContractTransactor) (*MatchesTransactor, error) {
	contract, err := bindMatches(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MatchesTransactor{contract: contract}, nil
}

// NewMatchesFilterer creates a new log filterer instance of Matches, bound to a specific deployed contract.
func NewMatchesFilterer(address common.Address, filterer bind.ContractFilterer) (*MatchesFilterer, error) {
	contract, err := bindMatches(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MatchesFilterer{contract: contract}, nil
}

// bindMatches binds a generic wrapper to an already deployed contract.
func bindMatches(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MatchesMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Matches *MatchesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Matches.Contract.MatchesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Matches *MatchesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Matches.Contract.MatchesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Matches *MatchesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Matches.Contract.MatchesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Matches *MatchesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Matches.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Matches *MatchesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Matches.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Matches *MatchesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Matches.Contract.contract.Transact(opts, method, params...)
}

// Cards is a free data retrieval call binding the contract method 0x58a4903f.
//
// Solidity: function cards() view returns(address)
func (_Matches *MatchesCaller) Cards(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Matches.contract.Call(opts, &out, "cards")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Cards is a free data retrieval call binding the contract method 0x58a4903f.
//
// Solidity: function cards() view returns(address)
func (_Matches *MatchesSession) Cards() (common.Address, error) {
	return _Matches.Contract.Cards(&_Matches.CallOpts)
}

// Cards is a free data retrieval call binding the contract method 0x58a4903f.
//
// Solidity: function cards() view returns(address)
func (_Matches *MatchesCallerSession) Cards() (common.Address, error) {
	return _Matches.Contract.Cards(&_Matches.CallOpts)
}

// GetMatch is a free data retrieval call binding the contract method 0x3d092b3d.
//
// Solidity: function getMatch(uint256 matchId) view returns((address,uint256,address,uint256,bool,address))
func (_Matches *MatchesCaller) GetMatch(opts *bind.CallOpts, matchId *big.Int) (MatchesMatch, error) {
	var out []interface{}
	err := _Matches.contract.Call(opts, &out, "getMatch", matchId)

	if err != nil {
		return *new(MatchesMatch), err
	}

	out0 := *abi.ConvertType(out[0], new(MatchesMatch)).(*MatchesMatch)

	return out0, err

}

// GetMatch is a free data retrieval call binding the contract method 0x3d092b3d.
//
// Solidity: function getMatch(uint256 matchId) view returns((address,uint256,address,uint256,bool,address))
func (_Matches *MatchesSession) GetMatch(matchId *big.Int) (MatchesMatch, error) {
	return _Matches.Contract.GetMatch(&_Matches.CallOpts, matchId)
}

// GetMatch is a free data retrieval call binding the contract method 0x3d092b3d.
//
// Solidity: function getMatch(uint256 matchId) view returns((address,uint256,address,uint256,bool,address))
func (_Matches *MatchesCallerSession) GetMatch(matchId *big.Int) (MatchesMatch, error) {
	return _Matches.Contract.GetMatch(&_Matches.CallOpts, matchId)
}

// IsWaiting is a free data retrieval call binding the contract method 0x7fecd538.
//
// Solidity: function isWaiting() view returns(bool)
func (_Matches *MatchesCaller) IsWaiting(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Matches.contract.Call(opts, &out, "isWaiting")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWaiting is a free data retrieval call binding the contract method 0x7fecd538.
//
// Solidity: function isWaiting() view returns(bool)
func (_Matches *MatchesSession) IsWaiting() (bool, error) {
	return _Matches.Contract.IsWaiting(&_Matches.CallOpts)
}

// IsWaiting is a free data retrieval call binding the contract method 0x7fecd538.
//
// Solidity: function isWaiting() view returns(bool)
func (_Matches *MatchesCallerSession) IsWaiting() (bool, error) {
	return _Matches.Contract.IsWaiting(&_Matches.CallOpts)
}

// Matches is a free data retrieval call binding the contract method 0x4768d4ef.
//
// Solidity: function matches(uint256 ) view returns(address playerA, uint256 cardA, address playerB, uint256 cardB, bool finished, address winner)
func (_Matches *MatchesCaller) Matches(opts *bind.CallOpts, arg0 *big.Int) (struct {
	PlayerA  common.Address
	CardA    *big.Int
	PlayerB  common.Address
	CardB    *big.Int
	Finished bool
	Winner   common.Address
}, error) {
	var out []interface{}
	err := _Matches.contract.Call(opts, &out, "matches", arg0)

	outstruct := new(struct {
		PlayerA  common.Address
		CardA    *big.Int
		PlayerB  common.Address
		CardB    *big.Int
		Finished bool
		Winner   common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PlayerA = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.CardA = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.PlayerB = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.CardB = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Finished = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.Winner = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Matches is a free data retrieval call binding the contract method 0x4768d4ef.
//
// Solidity: function matches(uint256 ) view returns(address playerA, uint256 cardA, address playerB, uint256 cardB, bool finished, address winner)
func (_Matches *MatchesSession) Matches(arg0 *big.Int) (struct {
	PlayerA  common.Address
	CardA    *big.Int
	PlayerB  common.Address
	CardB    *big.Int
	Finished bool
	Winner   common.Address
}, error) {
	return _Matches.Contract.Matches(&_Matches.CallOpts, arg0)
}

// Matches is a free data retrieval call binding the contract method 0x4768d4ef.
//
// Solidity: function matches(uint256 ) view returns(address playerA, uint256 cardA, address playerB, uint256 cardB, bool finished, address winner)
func (_Matches *MatchesCallerSession) Matches(arg0 *big.Int) (struct {
	PlayerA  common.Address
	CardA    *big.Int
	PlayerB  common.Address
	CardB    *big.Int
	Finished bool
	Winner   common.Address
}, error) {
	return _Matches.Contract.Matches(&_Matches.CallOpts, arg0)
}

// NextMatchId is a free data retrieval call binding the contract method 0xc5adf7c9.
//
// Solidity: function nextMatchId() view returns(uint256)
func (_Matches *MatchesCaller) NextMatchId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Matches.contract.Call(opts, &out, "nextMatchId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextMatchId is a free data retrieval call binding the contract method 0xc5adf7c9.
//
// Solidity: function nextMatchId() view returns(uint256)
func (_Matches *MatchesSession) NextMatchId() (*big.Int, error) {
	return _Matches.Contract.NextMatchId(&_Matches.CallOpts)
}

// NextMatchId is a free data retrieval call binding the contract method 0xc5adf7c9.
//
// Solidity: function nextMatchId() view returns(uint256)
func (_Matches *MatchesCallerSession) NextMatchId() (*big.Int, error) {
	return _Matches.Contract.NextMatchId(&_Matches.CallOpts)
}

// WaitingPlayer is a free data retrieval call binding the contract method 0x9d369d43.
//
// Solidity: function waitingPlayer() view returns(address)
func (_Matches *MatchesCaller) WaitingPlayer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Matches.contract.Call(opts, &out, "waitingPlayer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WaitingPlayer is a free data retrieval call binding the contract method 0x9d369d43.
//
// Solidity: function waitingPlayer() view returns(address)
func (_Matches *MatchesSession) WaitingPlayer() (common.Address, error) {
	return _Matches.Contract.WaitingPlayer(&_Matches.CallOpts)
}

// WaitingPlayer is a free data retrieval call binding the contract method 0x9d369d43.
//
// Solidity: function waitingPlayer() view returns(address)
func (_Matches *MatchesCallerSession) WaitingPlayer() (common.Address, error) {
	return _Matches.Contract.WaitingPlayer(&_Matches.CallOpts)
}

// Enqueue is a paid mutator transaction binding the contract method 0x52b17866.
//
// Solidity: function enqueue() returns()
func (_Matches *MatchesTransactor) Enqueue(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Matches.contract.Transact(opts, "enqueue")
}

// Enqueue is a paid mutator transaction binding the contract method 0x52b17866.
//
// Solidity: function enqueue() returns()
func (_Matches *MatchesSession) Enqueue() (*types.Transaction, error) {
	return _Matches.Contract.Enqueue(&_Matches.TransactOpts)
}

// Enqueue is a paid mutator transaction binding the contract method 0x52b17866.
//
// Solidity: function enqueue() returns()
func (_Matches *MatchesTransactorSession) Enqueue() (*types.Transaction, error) {
	return _Matches.Contract.Enqueue(&_Matches.TransactOpts)
}

// PlayCard is a paid mutator transaction binding the contract method 0x1b53804a.
//
// Solidity: function playCard(uint256 matchId, uint256 cardId) returns()
func (_Matches *MatchesTransactor) PlayCard(opts *bind.TransactOpts, matchId *big.Int, cardId *big.Int) (*types.Transaction, error) {
	return _Matches.contract.Transact(opts, "playCard", matchId, cardId)
}

// PlayCard is a paid mutator transaction binding the contract method 0x1b53804a.
//
// Solidity: function playCard(uint256 matchId, uint256 cardId) returns()
func (_Matches *MatchesSession) PlayCard(matchId *big.Int, cardId *big.Int) (*types.Transaction, error) {
	return _Matches.Contract.PlayCard(&_Matches.TransactOpts, matchId, cardId)
}

// PlayCard is a paid mutator transaction binding the contract method 0x1b53804a.
//
// Solidity: function playCard(uint256 matchId, uint256 cardId) returns()
func (_Matches *MatchesTransactorSession) PlayCard(matchId *big.Int, cardId *big.Int) (*types.Transaction, error) {
	return _Matches.Contract.PlayCard(&_Matches.TransactOpts, matchId, cardId)
}
