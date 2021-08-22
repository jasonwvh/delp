// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package reviews

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
)

// ReviewsReview is an auto generated low-level Go binding around an user-defined struct.
type ReviewsReview struct {
	Id      uint8
	Content string
	City    string
}

// ReviewsMetaData contains all meta data concerning the Reviews contract.
var ReviewsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"city\",\"type\":\"string\"}],\"name\":\"ReviewCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_content\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_city\",\"type\":\"string\"}],\"name\":\"createReview\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_id\",\"type\":\"uint8\"}],\"name\":\"getReviewsById\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"city\",\"type\":\"string\"}],\"internalType\":\"structReviews.Review\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reviewCount\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"reviews\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"city\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040526000805460ff191690553480156200001b57600080fd5b50620000776040518060400160405280600d81526020016c04d79204669727374204461707609c1b8152506040518060400160405280600c81526020016b25bab0b63090263ab6b83ab960a11b8152506200007d60201b60201c565b62000305565b6000805460ff1690806200009183620002d6565b82546101009290920a60ff8181021990931691831602179091556040805160608101825260008054841680835260208084018981528486018990529183526001808252949092208351815460ff19169616959095178555518051929550620000ff9385019291019062000167565b50604082015180516200011d91600284019160209091019062000167565b50506000546040517f305c4f9f74e2beb8cd6c1d15a55b17a85544eab06b0dd9e3f70f3339d530ce0892506200015b9160ff1690859085906200025d565b60405180910390a15050565b828054620001759062000299565b90600052602060002090601f016020900481019282620001995760008555620001e4565b82601f10620001b457805160ff1916838001178555620001e4565b82800160010185558215620001e4579182015b82811115620001e4578251825591602001919060010190620001c7565b50620001f2929150620001f6565b5090565b5b80821115620001f25760008155600101620001f7565b6000815180845260005b81811015620002355760208185018101518683018201520162000217565b8181111562000248576000602083870101525b50601f01601f19169290920160200192915050565b60ff841681526060602082015260006200027b60608301856200020d565b82810360408401526200028f81856200020d565b9695505050505050565b600181811c90821680620002ae57607f821691505b60208210811415620002d057634e487b7160e01b600052602260045260246000fd5b50919050565b600060ff821660ff811415620002fc57634e487b7160e01b600052601160045260246000fd5b60010192915050565b6107a580620003156000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c8063471eff0414610051578063627c2d421461007c578063c5d666c314610091578063eff5924f146100b1575b600080fd5b61006461005f3660046105f2565b6100d0565b604051610073939291906106b8565b60405180910390f35b61008f61008a36600461058e565b610208565b005b6100a461009f3660046105f2565b6102ea565b6040516100739190610669565b6000546100be9060ff1681565b60405160ff9091168152602001610073565b60016020819052600091825260409091208054918101805460ff909316926100f7906106f0565b80601f0160208091040260200160405190810160405280929190818152602001828054610123906106f0565b80156101705780601f1061014557610100808354040283529160200191610170565b820191906000526020600020905b81548152906001019060200180831161015357829003601f168201915b505050505090806002018054610185906106f0565b80601f01602080910402602001604051908101604052809291908181526020018280546101b1906106f0565b80156101fe5780601f106101d3576101008083540402835291602001916101fe565b820191906000526020600020905b8154815290600101906020018083116101e157829003601f168201915b5050505050905083565b6000805460ff16908061021a8361072b565b82546101009290920a60ff8181021990931691831602179091556040805160608101825260008054841680835260208084018981528486018990529183526001808252949092208351815460ff1916961695909517855551805192955061028693850192910190610468565b50604082015180516102a2916002840191602090910190610468565b50506000546040517f305c4f9f74e2beb8cd6c1d15a55b17a85544eab06b0dd9e3f70f3339d530ce0892506102de9160ff1690859085906106b8565b60405180910390a15050565b6103116040518060600160405280600060ff16815260200160608152602001606081525090565b60ff80831660009081526001602081815260409283902083516060810190945280549094168352908301805484928401919061034c906106f0565b80601f0160208091040260200160405190810160405280929190818152602001828054610378906106f0565b80156103c55780601f1061039a576101008083540402835291602001916103c5565b820191906000526020600020905b8154815290600101906020018083116103a857829003601f168201915b505050505081526020016002820180546103de906106f0565b80601f016020809104026020016040519081016040528092919081815260200182805461040a906106f0565b80156104575780601f1061042c57610100808354040283529160200191610457565b820191906000526020600020905b81548152906001019060200180831161043a57829003601f168201915b505050505081525050915050919050565b828054610474906106f0565b90600052602060002090601f01602090048101928261049657600085556104dc565b82601f106104af57805160ff19168380011785556104dc565b828001600101855582156104dc579182015b828111156104dc5782518255916020019190600101906104c1565b506104e89291506104ec565b5090565b5b808211156104e857600081556001016104ed565b600082601f83011261051257600080fd5b813567ffffffffffffffff8082111561052d5761052d610759565b604051601f8301601f19908116603f0116810190828211818310171561055557610555610759565b8160405283815286602085880101111561056e57600080fd5b836020870160208301376000602085830101528094505050505092915050565b600080604083850312156105a157600080fd5b823567ffffffffffffffff808211156105b957600080fd5b6105c586838701610501565b935060208501359150808211156105db57600080fd5b506105e885828601610501565b9150509250929050565b60006020828403121561060457600080fd5b813560ff8116811461061557600080fd5b9392505050565b6000815180845260005b8181101561064257602081850181015186830182015201610626565b81811115610654576000602083870101525b50601f01601f19169290920160200192915050565b6020815260ff82511660208201526000602083015160606040840152610692608084018261061c565b90506040840151601f198483030160608501526106af828261061c565b95945050505050565b60ff841681526060602082015260006106d4606083018561061c565b82810360408401526106e6818561061c565b9695505050505050565b600181811c9082168061070457607f821691505b6020821081141561072557634e487b7160e01b600052602260045260246000fd5b50919050565b600060ff821660ff81141561075057634e487b7160e01b600052601160045260246000fd5b60010192915050565b634e487b7160e01b600052604160045260246000fdfea2646970667358221220d375d06c484a1f61330b223f2653c8d894fb62e4fa388773e1bed8577befb21e64736f6c63430008070033",
}

// ReviewsABI is the input ABI used to generate the binding from.
// Deprecated: Use ReviewsMetaData.ABI instead.
var ReviewsABI = ReviewsMetaData.ABI

// ReviewsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ReviewsMetaData.Bin instead.
var ReviewsBin = ReviewsMetaData.Bin

// DeployReviews deploys a new Ethereum contract, binding an instance of Reviews to it.
func DeployReviews(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Reviews, error) {
	parsed, err := ReviewsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ReviewsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Reviews{ReviewsCaller: ReviewsCaller{contract: contract}, ReviewsTransactor: ReviewsTransactor{contract: contract}, ReviewsFilterer: ReviewsFilterer{contract: contract}}, nil
}

// Reviews is an auto generated Go binding around an Ethereum contract.
type Reviews struct {
	ReviewsCaller     // Read-only binding to the contract
	ReviewsTransactor // Write-only binding to the contract
	ReviewsFilterer   // Log filterer for contract events
}

// ReviewsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReviewsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReviewsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReviewsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReviewsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReviewsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReviewsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReviewsSession struct {
	Contract     *Reviews          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReviewsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReviewsCallerSession struct {
	Contract *ReviewsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ReviewsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReviewsTransactorSession struct {
	Contract     *ReviewsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ReviewsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReviewsRaw struct {
	Contract *Reviews // Generic contract binding to access the raw methods on
}

// ReviewsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReviewsCallerRaw struct {
	Contract *ReviewsCaller // Generic read-only contract binding to access the raw methods on
}

// ReviewsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReviewsTransactorRaw struct {
	Contract *ReviewsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReviews creates a new instance of Reviews, bound to a specific deployed contract.
func NewReviews(address common.Address, backend bind.ContractBackend) (*Reviews, error) {
	contract, err := bindReviews(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Reviews{ReviewsCaller: ReviewsCaller{contract: contract}, ReviewsTransactor: ReviewsTransactor{contract: contract}, ReviewsFilterer: ReviewsFilterer{contract: contract}}, nil
}

// NewReviewsCaller creates a new read-only instance of Reviews, bound to a specific deployed contract.
func NewReviewsCaller(address common.Address, caller bind.ContractCaller) (*ReviewsCaller, error) {
	contract, err := bindReviews(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReviewsCaller{contract: contract}, nil
}

// NewReviewsTransactor creates a new write-only instance of Reviews, bound to a specific deployed contract.
func NewReviewsTransactor(address common.Address, transactor bind.ContractTransactor) (*ReviewsTransactor, error) {
	contract, err := bindReviews(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReviewsTransactor{contract: contract}, nil
}

// NewReviewsFilterer creates a new log filterer instance of Reviews, bound to a specific deployed contract.
func NewReviewsFilterer(address common.Address, filterer bind.ContractFilterer) (*ReviewsFilterer, error) {
	contract, err := bindReviews(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReviewsFilterer{contract: contract}, nil
}

// bindReviews binds a generic wrapper to an already deployed contract.
func bindReviews(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReviewsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Reviews *ReviewsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Reviews.Contract.ReviewsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Reviews *ReviewsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reviews.Contract.ReviewsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Reviews *ReviewsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Reviews.Contract.ReviewsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Reviews *ReviewsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Reviews.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Reviews *ReviewsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reviews.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Reviews *ReviewsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Reviews.Contract.contract.Transact(opts, method, params...)
}

// GetReviewsById is a free data retrieval call binding the contract method 0xc5d666c3.
//
// Solidity: function getReviewsById(uint8 _id) view returns((uint8,string,string))
func (_Reviews *ReviewsCaller) GetReviewsById(opts *bind.CallOpts, _id uint8) (ReviewsReview, error) {
	var out []interface{}
	err := _Reviews.contract.Call(opts, &out, "getReviewsById", _id)

	if err != nil {
		return *new(ReviewsReview), err
	}

	out0 := *abi.ConvertType(out[0], new(ReviewsReview)).(*ReviewsReview)

	return out0, err

}

// GetReviewsById is a free data retrieval call binding the contract method 0xc5d666c3.
//
// Solidity: function getReviewsById(uint8 _id) view returns((uint8,string,string))
func (_Reviews *ReviewsSession) GetReviewsById(_id uint8) (ReviewsReview, error) {
	return _Reviews.Contract.GetReviewsById(&_Reviews.CallOpts, _id)
}

// GetReviewsById is a free data retrieval call binding the contract method 0xc5d666c3.
//
// Solidity: function getReviewsById(uint8 _id) view returns((uint8,string,string))
func (_Reviews *ReviewsCallerSession) GetReviewsById(_id uint8) (ReviewsReview, error) {
	return _Reviews.Contract.GetReviewsById(&_Reviews.CallOpts, _id)
}

// ReviewCount is a free data retrieval call binding the contract method 0xeff5924f.
//
// Solidity: function reviewCount() view returns(uint8)
func (_Reviews *ReviewsCaller) ReviewCount(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Reviews.contract.Call(opts, &out, "reviewCount")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ReviewCount is a free data retrieval call binding the contract method 0xeff5924f.
//
// Solidity: function reviewCount() view returns(uint8)
func (_Reviews *ReviewsSession) ReviewCount() (uint8, error) {
	return _Reviews.Contract.ReviewCount(&_Reviews.CallOpts)
}

// ReviewCount is a free data retrieval call binding the contract method 0xeff5924f.
//
// Solidity: function reviewCount() view returns(uint8)
func (_Reviews *ReviewsCallerSession) ReviewCount() (uint8, error) {
	return _Reviews.Contract.ReviewCount(&_Reviews.CallOpts)
}

// Reviews is a free data retrieval call binding the contract method 0x471eff04.
//
// Solidity: function reviews(uint8 ) view returns(uint8 id, string content, string city)
func (_Reviews *ReviewsCaller) Reviews(opts *bind.CallOpts, arg0 uint8) (struct {
	Id      uint8
	Content string
	City    string
}, error) {
	var out []interface{}
	err := _Reviews.contract.Call(opts, &out, "reviews", arg0)

	outstruct := new(struct {
		Id      uint8
		Content string
		City    string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.Content = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.City = *abi.ConvertType(out[2], new(string)).(*string)

	return *outstruct, err

}

// Reviews is a free data retrieval call binding the contract method 0x471eff04.
//
// Solidity: function reviews(uint8 ) view returns(uint8 id, string content, string city)
func (_Reviews *ReviewsSession) Reviews(arg0 uint8) (struct {
	Id      uint8
	Content string
	City    string
}, error) {
	return _Reviews.Contract.Reviews(&_Reviews.CallOpts, arg0)
}

// Reviews is a free data retrieval call binding the contract method 0x471eff04.
//
// Solidity: function reviews(uint8 ) view returns(uint8 id, string content, string city)
func (_Reviews *ReviewsCallerSession) Reviews(arg0 uint8) (struct {
	Id      uint8
	Content string
	City    string
}, error) {
	return _Reviews.Contract.Reviews(&_Reviews.CallOpts, arg0)
}

// CreateReview is a paid mutator transaction binding the contract method 0x627c2d42.
//
// Solidity: function createReview(string _content, string _city) returns()
func (_Reviews *ReviewsTransactor) CreateReview(opts *bind.TransactOpts, _content string, _city string) (*types.Transaction, error) {
	return _Reviews.contract.Transact(opts, "createReview", _content, _city)
}

// CreateReview is a paid mutator transaction binding the contract method 0x627c2d42.
//
// Solidity: function createReview(string _content, string _city) returns()
func (_Reviews *ReviewsSession) CreateReview(_content string, _city string) (*types.Transaction, error) {
	return _Reviews.Contract.CreateReview(&_Reviews.TransactOpts, _content, _city)
}

// CreateReview is a paid mutator transaction binding the contract method 0x627c2d42.
//
// Solidity: function createReview(string _content, string _city) returns()
func (_Reviews *ReviewsTransactorSession) CreateReview(_content string, _city string) (*types.Transaction, error) {
	return _Reviews.Contract.CreateReview(&_Reviews.TransactOpts, _content, _city)
}

// ReviewsReviewCreatedIterator is returned from FilterReviewCreated and is used to iterate over the raw logs and unpacked data for ReviewCreated events raised by the Reviews contract.
type ReviewsReviewCreatedIterator struct {
	Event *ReviewsReviewCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ReviewsReviewCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReviewsReviewCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ReviewsReviewCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ReviewsReviewCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReviewsReviewCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReviewsReviewCreated represents a ReviewCreated event raised by the Reviews contract.
type ReviewsReviewCreated struct {
	Id      uint8
	Content string
	City    string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterReviewCreated is a free log retrieval operation binding the contract event 0x305c4f9f74e2beb8cd6c1d15a55b17a85544eab06b0dd9e3f70f3339d530ce08.
//
// Solidity: event ReviewCreated(uint8 id, string content, string city)
func (_Reviews *ReviewsFilterer) FilterReviewCreated(opts *bind.FilterOpts) (*ReviewsReviewCreatedIterator, error) {

	logs, sub, err := _Reviews.contract.FilterLogs(opts, "ReviewCreated")
	if err != nil {
		return nil, err
	}
	return &ReviewsReviewCreatedIterator{contract: _Reviews.contract, event: "ReviewCreated", logs: logs, sub: sub}, nil
}

// WatchReviewCreated is a free log subscription operation binding the contract event 0x305c4f9f74e2beb8cd6c1d15a55b17a85544eab06b0dd9e3f70f3339d530ce08.
//
// Solidity: event ReviewCreated(uint8 id, string content, string city)
func (_Reviews *ReviewsFilterer) WatchReviewCreated(opts *bind.WatchOpts, sink chan<- *ReviewsReviewCreated) (event.Subscription, error) {

	logs, sub, err := _Reviews.contract.WatchLogs(opts, "ReviewCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReviewsReviewCreated)
				if err := _Reviews.contract.UnpackLog(event, "ReviewCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReviewCreated is a log parse operation binding the contract event 0x305c4f9f74e2beb8cd6c1d15a55b17a85544eab06b0dd9e3f70f3339d530ce08.
//
// Solidity: event ReviewCreated(uint8 id, string content, string city)
func (_Reviews *ReviewsFilterer) ParseReviewCreated(log types.Log) (*ReviewsReviewCreated, error) {
	event := new(ReviewsReviewCreated)
	if err := _Reviews.contract.UnpackLog(event, "ReviewCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
