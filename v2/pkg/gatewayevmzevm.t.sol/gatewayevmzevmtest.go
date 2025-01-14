// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gatewayevmzevm

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

// CallOptions is an auto generated low-level Go binding around an user-defined struct.
type CallOptions struct {
	GasLimit        *big.Int
	IsArbitraryCall bool
}

// RevertContext is an auto generated low-level Go binding around an user-defined struct.
type RevertContext struct {
	Sender        common.Address
	Asset         common.Address
	Amount        *big.Int
	RevertMessage []byte
}

// RevertOptions is an auto generated low-level Go binding around an user-defined struct.
type RevertOptions struct {
	RevertAddress    common.Address
	CallOnRevert     bool
	AbortAddress     common.Address
	RevertMessage    []byte
	OnRevertGasLimit *big.Int
}

// StdInvariantFuzzArtifactSelector is an auto generated low-level Go binding around an user-defined struct.
type StdInvariantFuzzArtifactSelector struct {
	Artifact  string
	Selectors [][4]byte
}

// StdInvariantFuzzInterface is an auto generated low-level Go binding around an user-defined struct.
type StdInvariantFuzzInterface struct {
	Addr      common.Address
	Artifacts []string
}

// StdInvariantFuzzSelector is an auto generated low-level Go binding around an user-defined struct.
type StdInvariantFuzzSelector struct {
	Addr      common.Address
	Selectors [][4]byte
}

// GatewayEVMZEVMTestMetaData contains all meta data concerning the GatewayEVMZEVMTest contract.
var GatewayEVMZEVMTestMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"IS_TEST\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"failed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setUp\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"targetArtifactSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifactSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzArtifactSelector[]\",\"components\":[{\"name\":\"artifact\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetInterfaces\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedInterfaces_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzInterface[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"artifacts\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"testCallReceiverEVMFromSenderZEVM\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testCallReceiverEVMFromZEVM\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallReceiverEVMFromSenderZEVM\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallReceiverEVMFromZEVM\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Called\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Called\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"zrc20\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposited\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"asset\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DepositedAndCalled\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"asset\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"payload\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Executed\",\"inputs\":[{\"name\":\"destination\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExecutedWithERC20\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedERC20\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"destination\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedNoParams\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedNonPayable\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strs\",\"type\":\"string[]\",\"indexed\":false,\"internalType\":\"string[]\"},{\"name\":\"nums\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"flag\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedOnCall\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedPayable\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"str\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"num\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"flag\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReceivedRevert\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Reverted\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedGatewayTSSAddress\",\"inputs\":[{\"name\":\"oldTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawn\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"zrc20\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"gasfee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"protocolFlatFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnAndCalled\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"zrc20\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"gasfee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"protocolFlatFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_address\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes32\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_int\",\"inputs\":[{\"name\":\"\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_address\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes32\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_string\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_string\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_uint\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"logs\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ApprovalFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerIsNotProtocol\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ConnectorInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CustodyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DepositFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExecutionFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedZetaSent\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"GasFeeTransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientERC20Amount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientETHAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientGasLimit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientZRC20Amount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientZetaAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTarget\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MessageSizeExceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAllowedToCallOnCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAllowedToCallOnRevert\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotWhitelistedInCustody\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyWZETAOrProtocol\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PayloadSizeExceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawalFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZRC20BurnFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZRC20DepositFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZRC20TransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x600c8054600160ff199182168117909255601f8054909116909117905560a06040526000608052602880546001600160a01b0319169055348015604157600080fd5b5061eadb806100516000396000f3fe608060405234801561001057600080fd5b506004361061011b5760003560e01c806385226c81116100b2578063b5508aa911610081578063d7a525fc11610066578063d7a525fc146101ec578063e20c9f71146101f4578063fa7626d4146101fc57600080fd5b8063b5508aa9146101cc578063ba414fa6146101d457600080fd5b806385226c8114610192578063916a17c6146101a75780639683c695146101bc578063b0464fdc146101c457600080fd5b80633f7286f4116100ee5780633f7286f414610165578063524744131461016d57806366d9a9a0146101755780636ff15ccc1461018a57600080fd5b80630a9254e4146101205780631ed7831c1461012a5780632ade3880146101485780633e5e3c231461015d575b600080fd5b610128610209565b005b6101326112fe565b60405161013f9190617bf1565b60405180910390f35b610150611360565b60405161013f9190617c8d565b6101326114a2565b610132611502565b610128611562565b61017d611dea565b60405161013f9190617df3565b610128611f6c565b61019a6127b9565b60405161013f9190617e91565b6101af612889565b60405161013f9190617f08565b610128612984565b6101af612faa565b61019a6130a5565b6101dc613175565b604051901515815260200161013f565b610128613249565b610132613986565b601f546101dc9060ff1681565b602580547fffffffffffffffffffffffff00000000000000000000000000000000000000009081163017909155602680548216611234179055602780548216615678179055602e805490911661432117905560405161026790617af7565b60408082526004908201527f746573740000000000000000000000000000000000000000000000000000000060608201526080602082018190526003908201527f54544b000000000000000000000000000000000000000000000000000000000060a082015260c001604051809103906000f0801580156102ec573d6000803e3d6000fd5b50602280547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039290921691909117905560405161033190617af7565b604080825260049082018190527f7a6574610000000000000000000000000000000000000000000000000000000060608301526080602083018190528201527f5a4554410000000000000000000000000000000000000000000000000000000060a082015260c001604051809103906000f0801580156103b5573d6000803e3d6000fd5b50602380547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03928316908117909155604080518082018252600e81527f4761746577617945564d2e736f6c0000000000000000000000000000000000006020820152602754602554925190851660248201526044810193909352921660648201526000916104a7916084015b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fc0c53b8b000000000000000000000000000000000000000000000000000000001790526139e6565b601f80547fffffffffffffffffffffff0000000000000000000000000000000000000000ff166101006001600160a01b0384811682029290921792839055604080518082018252601081527f4552433230437573746f64792e736f6c000000000000000000000000000000006020820152602754602554925193909504841660248401529383166044830152909116606482015291925061054a9160840161044a565b602080547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03838116919091178255604080518082018252601a81527f5a657461436f6e6e6563746f724e6f6e4e61746976652e736f6c00000000000093810193909352601f546023546027546025549351610100909304851660248401529084166044830152831660648201529116608482015291925061064e9160a40160408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167ff8c8765e000000000000000000000000000000000000000000000000000000001790526139e6565b602180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03838116919091179091556027546040517fc88a5e6d00000000000000000000000000000000000000000000000000000000815291166004820152670de0b6b3a76400006024820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c88a5e6d90604401600060405180830381600087803b15801561070057600080fd5b505af1158015610714573d6000803e3d6000fd5b50506025546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506306447d569150602401600060405180830381600087803b15801561078a57600080fd5b505af115801561079e573d6000803e3d6000fd5b5050601f546020546040517fae7a3a6f0000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015261010090920416925063ae7a3a6f9150602401600060405180830381600087803b15801561080957600080fd5b505af115801561081d573d6000803e3d6000fd5b5050601f546021546040517f10188aef0000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201526101009092041692506310188aef9150602401600060405180830381600087803b15801561088857600080fd5b505af115801561089c573d6000803e3d6000fd5b505050507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156108fe57600080fd5b505af1158015610912573d6000803e3d6000fd5b50506022546025546040517f40c10f190000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620f42406024820152911692506340c10f199150604401600060405180830381600087803b15801561098157600080fd5b505af1158015610995573d6000803e3d6000fd5b50506022546020546040517fa9059cbb0000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201526207a12060248201529116925063a9059cbb91506044016020604051808303816000875af1158015610a09573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a2d9190617f9f565b50604051610a3a90617b04565b604051809103906000f080158015610a56573d6000803e3d6000fd5b50602480547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03928316178155604080518082018252600f81527f476174657761795a45564d2e736f6c00000000000000000000000000000000006020820152602354602e5492519085169381019390935292166044820152610b3c919060640160408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f485cc955000000000000000000000000000000000000000000000000000000001790526139e6565b602980546001600160a01b03929092167fffffffffffffffffffffffff00000000000000000000000000000000000000009283168117909155602a80549092168117909155604051610b8d90617b11565b6001600160a01b039091168152602001604051809103906000f080158015610bb9573d6000803e3d6000fd5b50602b80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03929092169190911790556040517f06447d5600000000000000000000000000000000000000000000000000000000815273735b14bb79463307aacbed86daf3322b1e6226ab6004820181905290737109709ecfa91a80626ff3989d68f67f5b1dd12d906306447d5690602401600060405180830381600087803b158015610c6e57600080fd5b505af1158015610c82573d6000803e3d6000fd5b505050506000806000604051610c9790617b1e565b6001600160a01b03938416815291831660208301529091166040820152606001604051809103906000f080158015610cd3573d6000803e3d6000fd5b50602c80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03928316908117909155602a54604051601293600193600093849391921690610d2990617b2b565b610d3896959493929190617fc1565b604051809103906000f080158015610d54573d6000803e3d6000fd5b50602d80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03928316908117909155602c546040517fee2815ba0000000000000000000000000000000000000000000000000000000081526001600482015260248101929092529091169063ee2815ba90604401600060405180830381600087803b158015610deb57600080fd5b505af1158015610dff573d6000803e3d6000fd5b5050602c546040517fa7cb050700000000000000000000000000000000000000000000000000000000815260016004820181905260248201526001600160a01b03909116925063a7cb05079150604401600060405180830381600087803b158015610e6957600080fd5b505af1158015610e7d573d6000803e3d6000fd5b5050602d54602e546040517f47e7ef240000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620f42406024820152911692506347e7ef2491506044016020604051808303816000875af1158015610ef1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f159190617f9f565b50602d54602b546040517f47e7ef240000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620f424060248201529116906347e7ef24906044016020604051808303816000875af1158015610f86573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610faa9190617f9f565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561100957600080fd5b505af115801561101d573d6000803e3d6000fd5b5050602e546040517fca669fa70000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b15801561109357600080fd5b505af11580156110a7573d6000803e3d6000fd5b5050602d54602a546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152620f424060248201529116925063095ea7b391506044016020604051808303816000875af115801561111b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061113f9190617f9f565b506027546040517fc88a5e6d0000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152670de0b6b3a76400006024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c88a5e6d90604401600060405180830381600087803b1580156111c057600080fd5b505af11580156111d4573d6000803e3d6000fd5b5050604080518082018252600180825260006020928301819052602f829055603080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055835160a08101855261032180825281850193845281860190815285519485019095528184526060810184905260808101919091528051603180549351151574010000000000000000000000000000000000000000027fffffffffffffffffffffff0000000000000000000000000000000000000000009094166001600160a01b0392831617939093178355935160328054919095167fffffffffffffffffffffffff000000000000000000000000000000000000000091909116179093559193509091506033906112ec9082618181565b50608082015181600301559050505050565b6060601680548060200260200160405190810160405280929190818152602001828054801561135657602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611338575b5050505050905090565b6060601e805480602002602001604051908101604052809291908181526020016000905b8282101561149957600084815260208082206040805180820182526002870290920180546001600160a01b03168352600181018054835181870281018701909452808452939591948681019491929084015b828210156114825783829060005260206000200180546113f5906180e5565b80601f0160208091040260200160405190810160405280929190818152602001828054611421906180e5565b801561146e5780601f106114435761010080835404028352916020019161146e565b820191906000526020600020905b81548152906001019060200180831161145157829003601f168201915b5050505050815260200190600101906113d6565b505050508152505081526020019060010190611384565b50505050905090565b60606018805480602002602001604051908101604052809291908181526020018280548015611356576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311611338575050505050905090565b60606017805480602002602001604051908101604052809291908181526020018280548015611356576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311611338575050505050905090565b604080518082018252600681527f48656c6c6f2100000000000000000000000000000000000000000000000000006020820152602d54602b5492517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b0393841660048201529192602a92600192670de0b6b3a7640000926000929116906370a0823190602401602060405180830381865afa15801561160e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116329190618240565b6040519091506000907fe04d4f97000000000000000000000000000000000000000000000000000000009061166f90889088908890602401618259565b60408051601f19818403018152918152602080830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff00000000000000000000000000000000000000000000000000000000909516949094179093526024549051919350600092611708926001600160a01b03909216910160609190911b6bffffffffffffffffffffffff1916815260140190565b60408051601f1981840301815290829052602d5461173f926207a120916001600160a01b0316908690602f90603190602401618372565b60408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f7b15118b00000000000000000000000000000000000000000000000000000000179052602a5490517ff30c7ba3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f30c7ba3916117fc916001600160a01b03919091169060009086906004016183e9565b600060405180830381600087803b15801561181657600080fd5b505af115801561182a573d6000803e3d6000fd5b5050602e546040517fca669fa70000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b1580156118a057600080fd5b505af11580156118b4573d6000803e3d6000fd5b5050602b5460245460405160609190911b6bffffffffffffffffffffffff191660208201526001600160a01b039091169250630abd8905915060340160408051601f1981840301815290829052602d547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16835261194e926207a120916001600160a01b0316908d908d908d90600401618411565b600060405180830381600087803b15801561196857600080fd5b505af115801561197c573d6000803e3d6000fd5b5050601f546040517fc88a5e6d0000000000000000000000000000000000000000000000000000000081526101009091046001600160a01b0316600482015260248101879052737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c88a5e6d9150604401600060405180830381600087803b1580156119fd57600080fd5b505af1158015611a11573d6000803e3d6000fd5b5050602480546040517f81bad6f30000000000000000000000000000000000000000000000000000000081526001600482018190529281018390526044810183905260648101929092526001600160a01b03166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b158015611aa457600080fd5b505af1158015611ab8573d6000803e3d6000fd5b505050507f1f1ff1f5fb41346850b2f5c04e6c767e2f1c8a525c5c0c5cdb60cdf3ca5f62fa601f60019054906101000a90046001600160a01b031685898989604051611b08959493929190618466565b60405180910390a1601f546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526101009091046001600160a01b03166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b158015611ba257600080fd5b505af1158015611bb6573d6000803e3d6000fd5b50506024546040516001600160a01b0390911692507fcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f9150611bfb90879086906184a7565b60405180910390a26027546040517fca669fa70000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015611c7557600080fd5b505af1158015611c89573d6000803e3d6000fd5b5050601f546024546040517f38e225270000000000000000000000000000000000000000000000000000000081526001600160a01b03610100909304831694506338e2252793508892611ce69260289291169088906004016184c0565b60006040518083038185885af1158015611d04573d6000803e3d6000fd5b50505050506040513d6000823e601f3d908101601f19168201604052611d2d91908101906185ab565b50602d54602b546040517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260009291909116906370a0823190602401602060405180830381865afa158015611d98573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611dbc9190618240565b9050611de0816001611dd16207a1208861860f565b611ddb919061860f565b613a05565b5050505050505050565b6060601b805480602002602001604051908101604052809291908181526020016000905b828210156114995783829060005260206000209060020201604051806040016040529081600082018054611e41906180e5565b80601f0160208091040260200160405190810160405280929190818152602001828054611e6d906180e5565b8015611eba5780601f10611e8f57610100808354040283529160200191611eba565b820191906000526020600020905b815481529060010190602001808311611e9d57829003601f168201915b5050505050815260200160018201805480602002602001604051908101604052809291908181526020018280548015611f5457602002820191906000526020600020906000905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681526020019060040190602082600301049283019260010382029150808411611f015790505b50505050508152505081526020019060010190611e0e565b604080518082018252600681527f48656c6c6f21000000000000000000000000000000000000000000000000000060208201529051602a90600190670de0b6b3a7640000906000907fe04d4f970000000000000000000000000000000000000000000000000000000090611fe890879087908790602401618259565b60408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009490941693909317909252602a5491517f81bad6f30000000000000000000000000000000000000000000000000000000081526001600482018190526024820181905260448201819052606482018190526001600160a01b039093166084820152909250737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b1580156120da57600080fd5b505af11580156120ee573d6000803e3d6000fd5b5050602e5460245460405160609190911b6bffffffffffffffffffffffff19166020820152600093506001600160a01b0390911691507fd90f94752d2b12f364f4a2237ebe1aff24ba6127585376bf4935f6a7be17dd2a9060340160408051601f19818403018152828252602d547f4d8943bb000000000000000000000000000000000000000000000000000000008452915190926001600160a01b03909216916207a1209188918491634d8943bb916004808201926020929091908290030181865afa1580156121c3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121e79190618240565b60408051808201825260018082526020820152905161220f9695949392918c91603190618622565b60405180910390a3602e546040517fca669fa70000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561228957600080fd5b505af115801561229d573d6000803e3d6000fd5b5050602a546024546040805160609290921b6bffffffffffffffffffffffff1916602083015280518083036014018152602d5460748401835260016034850181815260549095015291517f7b15118b0000000000000000000000000000000000000000000000000000000081526001600160a01b039485169650637b15118b955061233b9491936207a1209392909216918991603190600401618692565b600060405180830381600087803b15801561235557600080fd5b505af1158015612369573d6000803e3d6000fd5b5050602d54602e546040517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b03918216600482015260009450911691506370a0823190602401602060405180830381865afa1580156123d5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906123f99190618240565b905061240c81611ddb846207a12061860f565b601f546040517fc88a5e6d0000000000000000000000000000000000000000000000000000000081526101009091046001600160a01b0316600482015260248101859052737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c88a5e6d90604401600060405180830381600087803b15801561248957600080fd5b505af115801561249d573d6000803e3d6000fd5b5050602480546040517f81bad6f30000000000000000000000000000000000000000000000000000000081526001600482018190529281018390526044810183905260648101929092526001600160a01b03166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b15801561253057600080fd5b505af1158015612544573d6000803e3d6000fd5b505050507f1f1ff1f5fb41346850b2f5c04e6c767e2f1c8a525c5c0c5cdb60cdf3ca5f62fa601f60019054906101000a90046001600160a01b031685898989604051612594959493929190618466565b60405180910390a1601f546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526101009091046001600160a01b03166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b15801561262e57600080fd5b505af1158015612642573d6000803e3d6000fd5b50506024546040516001600160a01b0390911692507fcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f915061268790879087906184a7565b60405180910390a26027546040517fca669fa70000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561270157600080fd5b505af1158015612715573d6000803e3d6000fd5b5050601f546024546040517f38e225270000000000000000000000000000000000000000000000000000000081526001600160a01b03610100909304831694506338e22527935088926127729260289291169089906004016184c0565b60006040518083038185885af1158015612790573d6000803e3d6000fd5b50505050506040513d6000823e601f3d908101601f19168201604052611de091908101906185ab565b6060601a805480602002602001604051908101604052809291908181526020016000905b828210156114995783829060005260206000200180546127fc906180e5565b80601f0160208091040260200160405190810160405280929190818152602001828054612828906180e5565b80156128755780601f1061284a57610100808354040283529160200191612875565b820191906000526020600020905b81548152906001019060200180831161285857829003601f168201915b5050505050815260200190600101906127dd565b6060601d805480602002602001604051908101604052809291908181526020016000905b828210156114995760008481526020908190206040805180820182526002860290920180546001600160a01b0316835260018101805483518187028101870190945280845293949193858301939283018282801561296c57602002820191906000526020600020906000905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190600401906020826003010492830192600103820291508084116129195790505b505050505081525050815260200190600101906128ad565b604080518082018252600681527f48656c6c6f21000000000000000000000000000000000000000000000000000060208201529051602a90600190670de0b6b3a7640000906000907fe04d4f970000000000000000000000000000000000000000000000000000000090612a0090879087908790602401618259565b60408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009490941693909317909252602e5491517fca669fa70000000000000000000000000000000000000000000000000000000081526001600160a01b0390921660048301529150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015612ad357600080fd5b505af1158015612ae7573d6000803e3d6000fd5b5050602a546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b158015612b7957600080fd5b505af1158015612b8d573d6000803e3d6000fd5b5050602d54602e5460245460405160609190911b6bffffffffffffffffffffffff191660208201526001600160a01b039283169450911691507f306ee13f48319a123b222c69908e44dcf91abffc20cacc502e3cf5a4ff23e0e49060340160408051601f198184030181528282018252600180845260208401529051612c179287916031906186e6565b60405180910390a3602a5460245460405160609190911b6bffffffffffffffffffffffff191660208201526001600160a01b03909116906306cb89839060340160408051601f19818403018152602d5483830183526001808552602085015291517fffffffff0000000000000000000000000000000000000000000000000000000060e086901b168152612cbe9391926001600160a01b0316918791603190600401618732565b600060405180830381600087803b158015612cd857600080fd5b505af1158015612cec573d6000803e3d6000fd5b5050601f546040517fc88a5e6d0000000000000000000000000000000000000000000000000000000081526101009091046001600160a01b0316600482015260248101859052737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c88a5e6d9150604401600060405180830381600087803b158015612d6d57600080fd5b505af1158015612d81573d6000803e3d6000fd5b5050601f546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526101009091046001600160a01b03166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b158015612e1757600080fd5b505af1158015612e2b573d6000803e3d6000fd5b50506024546040516001600160a01b0390911692507fcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f9150612e7090859085906184a7565b60405180910390a26027546040517fca669fa70000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015612eea57600080fd5b505af1158015612efe573d6000803e3d6000fd5b5050601f546024546040517f38e225270000000000000000000000000000000000000000000000000000000081526001600160a01b03610100909304831694506338e2252793508692612f5b9260289291169087906004016184c0565b60006040518083038185885af1158015612f79573d6000803e3d6000fd5b50505050506040513d6000823e601f3d908101601f19168201604052612fa291908101906185ab565b505050505050565b6060601c805480602002602001604051908101604052809291908181526020016000905b828210156114995760008481526020908190206040805180820182526002860290920180546001600160a01b0316835260018101805483518187028101870190945280845293949193858301939283018282801561308d57602002820191906000526020600020906000905b82829054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168152602001906004019060208260030104928301926001038202915080841161303a5790505b50505050508152505081526020019060010190612fce565b60606019805480602002602001604051908101604052809291908181526020016000905b828210156114995783829060005260206000200180546130e8906180e5565b80601f0160208091040260200160405190810160405280929190818152602001828054613114906180e5565b80156131615780601f1061313657610100808354040283529160200191613161565b820191906000526020600020905b81548152906001019060200180831161314457829003601f168201915b5050505050815260200190600101906130c9565b60085460009060ff161561318d575060085460ff1690565b6040517f667f9d70000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d600482018190527f6661696c65640000000000000000000000000000000000000000000000000000602483015260009163667f9d7090604401602060405180830381865afa15801561321e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906132429190618240565b1415905090565b604080518082018252600681527f48656c6c6f21000000000000000000000000000000000000000000000000000060208201529051602a90600190670de0b6b3a7640000906000907fe04d4f9700000000000000000000000000000000000000000000000000000000906132c590879087908790602401618259565b60408051601f19818403018152918152602080830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090951694909417909352602454905191935060009261335e926001600160a01b03909216910160609190911b6bffffffffffffffffffffffff1916815260140190565b60408051601f1981840301815290829052602d54613392926001600160a01b03909116908590602f9060319060240161879a565b60408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f06cb898300000000000000000000000000000000000000000000000000000000179052602a5490517ff30c7ba3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f30c7ba39161344f916001600160a01b03919091169060009086906004016183e9565b600060405180830381600087803b15801561346957600080fd5b505af115801561347d573d6000803e3d6000fd5b5050602e546040517fca669fa70000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063ca669fa79150602401600060405180830381600087803b1580156134f357600080fd5b505af1158015613507573d6000803e3d6000fd5b5050602b5460245460405160609190911b6bffffffffffffffffffffffff191660208201526001600160a01b039091169250637a34d8bb915060340160408051601f1981840301815290829052602d547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16835261359e926001600160a01b03909116908b908b908b906004016187eb565b600060405180830381600087803b1580156135b857600080fd5b505af11580156135cc573d6000803e3d6000fd5b5050601f546040517fc88a5e6d0000000000000000000000000000000000000000000000000000000081526101009091046001600160a01b0316600482015260248101869052737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c88a5e6d9150604401600060405180830381600087803b15801561364d57600080fd5b505af1158015613661573d6000803e3d6000fd5b5050602480546040517f81bad6f30000000000000000000000000000000000000000000000000000000081526001600482018190529281018390526044810183905260648101929092526001600160a01b03166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506381bad6f3915060a401600060405180830381600087803b1580156136f457600080fd5b505af1158015613708573d6000803e3d6000fd5b505050507f1f1ff1f5fb41346850b2f5c04e6c767e2f1c8a525c5c0c5cdb60cdf3ca5f62fa601f60019054906101000a90046001600160a01b031684888888604051613758959493929190618466565b60405180910390a1601f546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526101009091046001600160a01b03166084820152737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b1580156137f257600080fd5b505af1158015613806573d6000803e3d6000fd5b50506024546040516001600160a01b0390911692507fcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f915061384b90869086906184a7565b60405180910390a26027546040517fca669fa70000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156138c557600080fd5b505af11580156138d9573d6000803e3d6000fd5b5050601f546024546040517f38e225270000000000000000000000000000000000000000000000000000000081526001600160a01b03610100909304831694506338e22527935087926139369260289291169088906004016184c0565b60006040518083038185885af1158015613954573d6000803e3d6000fd5b50505050506040513d6000823e601f3d908101601f1916820160405261397d91908101906185ab565b50505050505050565b60606015805480602002602001604051908101604052809291908181526020018280548015611356576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311611338575050505050905090565b60006139f0617b38565b6139fb848483613a84565b9150505b92915050565b6040517f98296c540000000000000000000000000000000000000000000000000000000081526004810183905260248101829052737109709ecfa91a80626ff3989d68f67f5b1dd12d906398296c549060440160006040518083038186803b158015613a7057600080fd5b505afa158015612fa2573d6000803e3d6000fd5b600080613a918584613aff565b9050613af46040518060400160405280601d81526020017f4552433139363750726f78792e736f6c3a4552433139363750726f78790000008152508286604051602001613adf929190618839565b60405160208183030381529060405285613b0b565b9150505b9392505050565b6000613af88383613b39565b60c08101515160009015613b2f57613b2884848460c00151613b54565b9050613af8565b613b288484613cfa565b6000613b458383613de5565b613af883836020015184613b0b565b600080613b5f613df5565b90506000613b6d8683613ec8565b90506000613b84826060015183602001518561436e565b90506000613b9483838989614580565b90506000613ba1826153fd565b602081015181519192509060030b15613c1457898260400151604051602001613bcb92919061885b565b60408051601f19818403018152908290527f08c379a0000000000000000000000000000000000000000000000000000000008252613c0b916004016188dc565b60405180910390fd5b6000613c576040518060400160405280601581526020017f4465706c6f79656420746f20616464726573733a2000000000000000000000008152508360016155cc565b6040517fc6ce059d000000000000000000000000000000000000000000000000000000008152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c6ce059d90613caa9084906004016188dc565b602060405180830381865afa158015613cc7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613ceb91906188ef565b9b9a5050505050505050505050565b6040517f8d1cc9250000000000000000000000000000000000000000000000000000000081526000908190737109709ecfa91a80626ff3989d68f67f5b1dd12d90638d1cc92590613d4f9087906004016188dc565b600060405180830381865afa158015613d6c573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052613d9491908101906185ab565b90506000613dc28285604051602001613dae929190618918565b6040516020818303038152906040526157cc565b90506001600160a01b0381166139fb578484604051602001613bcb929190618947565b613df1828260006157df565b5050565b604080518082018252600381527f6f75740000000000000000000000000000000000000000000000000000000000602082015290517fd145736c000000000000000000000000000000000000000000000000000000008152606091737109709ecfa91a80626ff3989d68f67f5b1dd12d91829063d145736c90613e7c9084906004016189f2565b600060405180830381865afa158015613e99573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052613ec19190810190618a39565b9250505090565b613efa6040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b6000737109709ecfa91a80626ff3989d68f67f5b1dd12d9050613f456040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b613f4e856158e2565b60208201526000613f5e86615cc7565b90506000836001600160a01b031663d930a0e66040518163ffffffff1660e01b8152600401600060405180830381865afa158015613fa0573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052613fc89190810190618a39565b86838560200151604051602001613fe29493929190618a82565b60408051601f19818403018152908290527f60f9bb1100000000000000000000000000000000000000000000000000000000825291506000906001600160a01b038616906360f9bb119061403a9085906004016188dc565b600060405180830381865afa158015614057573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261407f9190810190618a39565b6040517fdb4235f60000000000000000000000000000000000000000000000000000000081529091506001600160a01b0386169063db4235f6906140c7908490600401618b86565b602060405180830381865afa1580156140e4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906141089190617f9f565b61411d5781604051602001613bcb9190618bd8565b6040517f49c4fac80000000000000000000000000000000000000000000000000000000081526001600160a01b038616906349c4fac890614162908490600401618c6a565b600060405180830381865afa15801561417f573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526141a79190810190618a39565b84526040517fdb4235f60000000000000000000000000000000000000000000000000000000081526001600160a01b0386169063db4235f6906141ee908490600401618cbc565b602060405180830381865afa15801561420b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061422f9190617f9f565b156142c4576040517f49c4fac80000000000000000000000000000000000000000000000000000000081526001600160a01b038616906349c4fac890614279908490600401618cbc565b600060405180830381865afa158015614296573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526142be9190810190618a39565b60408501525b846001600160a01b03166349c4fac88286600001516040516020016142e99190618d0e565b6040516020818303038152906040526040518363ffffffff1660e01b8152600401614315929190618d7a565b600060405180830381865afa158015614332573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261435a9190810190618a39565b606085015250608083015250949350505050565b60408051600480825260a0820190925260609160009190816020015b606081526020019060019003908161438a5790505090506040518060400160405280600481526020017f6772657000000000000000000000000000000000000000000000000000000000815250816000815181106143ea576143ea618d9f565b60200260200101819052506040518060400160405280600381526020017f2d726c00000000000000000000000000000000000000000000000000000000008152508160018151811061443e5761443e618d9f565b60200260200101819052508460405160200161445a9190618dce565b6040516020818303038152906040528160028151811061447c5761447c618d9f565b6020026020010181905250826040516020016144989190618e3a565b604051602081830303815290604052816003815181106144ba576144ba618d9f565b602002602001018190525060006144d0826153fd565b602080820151604080518082018252600581527f2e6a736f6e00000000000000000000000000000000000000000000000000000081850190815282518084018452600080825290860152825180840190935290518252928101929092529192506145619060408051808201825260008082526020918201528151808301909252845182528085019082015290615f4a565b6145765785604051602001613bcb9190618e7b565b9695505050505050565b60a0810151604080518082018252600080825260209182015281518083019092528251808352928101910152606090737109709ecfa91a80626ff3989d68f67f5b1dd12d90156145d0565b511590565b6147445782602001511561468c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605860248201527f54686520606c6963656e73655479706560206f7074696f6e2063616e6e6f742060448201527f62652075736564207768656e207468652060736b6970566572696679536f757260648201527f6365436f646560206f7074696f6e206973206074727565600000000000000000608482015260a401613c0b565b8260c0015115614744576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605360248201527f54686520606c6963656e73655479706560206f7074696f6e2063616e6e6f742060448201527f62652075736564207768656e207468652060736b69704c6963656e736554797060648201527f6560206f7074696f6e2069732060747275656000000000000000000000000000608482015260a401613c0b565b6040805160ff8082526120008201909252600091816020015b606081526020019060019003908161475d57905050905060006040518060400160405280600381526020017f6e707800000000000000000000000000000000000000000000000000000000008152508282806147b890618f0c565b935060ff16815181106147cd576147cd618d9f565b60200260200101819052506040518060400160405280600d81526020017f302e302e312d616c7068612e370000000000000000000000000000000000000081525060405160200161481e9190618f2b565b60405160208183030381529060405282828061483990618f0c565b935060ff168151811061484e5761484e618d9f565b60200260200101819052506040518060400160405280600681526020017f6465706c6f79000000000000000000000000000000000000000000000000000081525082828061489b90618f0c565b935060ff16815181106148b0576148b0618d9f565b60200260200101819052506040518060400160405280600e81526020017f2d2d636f6e74726163744e616d650000000000000000000000000000000000008152508282806148fd90618f0c565b935060ff168151811061491257614912618d9f565b6020026020010181905250876020015182828061492e90618f0c565b935060ff168151811061494357614943618d9f565b60200260200101819052506040518060400160405280600e81526020017f2d2d636f6e74726163745061746800000000000000000000000000000000000081525082828061499090618f0c565b935060ff16815181106149a5576149a5618d9f565b6020908102919091010152875182826149bd81618f0c565b935060ff16815181106149d2576149d2618d9f565b60200260200101819052506040518060400160405280600981526020017f2d2d636861696e49640000000000000000000000000000000000000000000000815250828280614a1f90618f0c565b935060ff1681518110614a3457614a34618d9f565b6020026020010181905250614a4846615fab565b8282614a5381618f0c565b935060ff1681518110614a6857614a68618d9f565b60200260200101819052506040518060400160405280600f81526020017f2d2d6275696c64496e666f46696c650000000000000000000000000000000000815250828280614ab590618f0c565b935060ff1681518110614aca57614aca618d9f565b602002602001018190525086828280614ae290618f0c565b935060ff1681518110614af757614af7618d9f565b6020908102919091010152855115614c1e5760408051808201909152601581527f2d2d636f6e7374727563746f7242797465636f6465000000000000000000000060208201528282614b4881618f0c565b935060ff1681518110614b5d57614b5d618d9f565b60209081029190910101526040517f71aad10d0000000000000000000000000000000000000000000000000000000081526001600160a01b038416906371aad10d90614bad9089906004016188dc565b600060405180830381865afa158015614bca573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052614bf29190810190618a39565b8282614bfd81618f0c565b935060ff1681518110614c1257614c12618d9f565b60200260200101819052505b846020015115614cee5760408051808201909152601281527f2d2d766572696679536f75726365436f6465000000000000000000000000000060208201528282614c6781618f0c565b935060ff1681518110614c7c57614c7c618d9f565b60200260200101819052506040518060400160405280600581526020017f66616c7365000000000000000000000000000000000000000000000000000000815250828280614cc990618f0c565b935060ff1681518110614cde57614cde618d9f565b6020026020010181905250614eb5565b614d266145cb8660a0015160408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b614db95760408051808201909152600d81527f2d2d6c6963656e7365547970650000000000000000000000000000000000000060208201528282614d6981618f0c565b935060ff1681518110614d7e57614d7e618d9f565b60200260200101819052508460a00151604051602001614d9e9190618dce565b604051602081830303815290604052828280614cc990618f0c565b8460c00151158015614dfc575060408089015181518083018352600080825260209182015282518084019093528151835290810190820152614dfa90511590565b155b15614eb55760408051808201909152600d81527f2d2d6c6963656e7365547970650000000000000000000000000000000000000060208201528282614e4081618f0c565b935060ff1681518110614e5557614e55618d9f565b6020026020010181905250614e698861604b565b604051602001614e799190618dce565b604051602081830303815290604052828280614e9490618f0c565b935060ff1681518110614ea957614ea9618d9f565b60200260200101819052505b60408086015181518083018352600080825260209182015282518084019093528151835290810190820152614ee990511590565b614f7e5760408051808201909152600b81527f2d2d72656c61796572496400000000000000000000000000000000000000000060208201528282614f2c81618f0c565b935060ff1681518110614f4157614f41618d9f565b60200260200101819052508460400151828280614f5d90618f0c565b935060ff1681518110614f7257614f72618d9f565b60200260200101819052505b60608501511561509f5760408051808201909152600681527f2d2d73616c74000000000000000000000000000000000000000000000000000060208201528282614fc781618f0c565b935060ff1681518110614fdc57614fdc618d9f565b602090810291909101015260608501516040517fb11a19e800000000000000000000000000000000000000000000000000000000815260048101919091526001600160a01b0384169063b11a19e890602401600060405180830381865afa15801561504b573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526150739190810190618a39565b828261507e81618f0c565b935060ff168151811061509357615093618d9f565b60200260200101819052505b60e085015151156151465760408051808201909152600a81527f2d2d6761734c696d697400000000000000000000000000000000000000000000602082015282826150e981618f0c565b935060ff16815181106150fe576150fe618d9f565b602002602001018190525061511a8560e0015160000151615fab565b828261512581618f0c565b935060ff168151811061513a5761513a618d9f565b60200260200101819052505b60e085015160200151156151f05760408051808201909152600a81527f2d2d6761735072696365000000000000000000000000000000000000000000006020820152828261519381618f0c565b935060ff16815181106151a8576151a8618d9f565b60200260200101819052506151c48560e0015160200151615fab565b82826151cf81618f0c565b935060ff16815181106151e4576151e4618d9f565b60200260200101819052505b60e0850151604001511561529a5760408051808201909152600e81527f2d2d6d61784665655065724761730000000000000000000000000000000000006020820152828261523d81618f0c565b935060ff168151811061525257615252618d9f565b602002602001018190525061526e8560e0015160400151615fab565b828261527981618f0c565b935060ff168151811061528e5761528e618d9f565b60200260200101819052505b60e085015160600151156153445760408051808201909152601681527f2d2d6d61785072696f7269747946656550657247617300000000000000000000602082015282826152e781618f0c565b935060ff16815181106152fc576152fc618d9f565b60200260200101819052506153188560e0015160600151615fab565b828261532381618f0c565b935060ff168151811061533857615338618d9f565b60200260200101819052505b60008160ff1667ffffffffffffffff811115615362576153626180b6565b60405190808252806020026020018201604052801561539557816020015b60608152602001906001900390816153805790505b50905060005b8260ff168160ff1610156153ee57838160ff16815181106153be576153be618d9f565b6020026020010151828260ff16815181106153db576153db618d9f565b602090810291909101015260010161539b565b5093505050505b949350505050565b6154246040518060600160405280600060030b815260200160608152602001606081525090565b60408051808201825260048082527f6261736800000000000000000000000000000000000000000000000000000000602083015291517fd145736c000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d92600091849163d145736c916154aa91869101618f96565b600060405180830381865afa1580156154c7573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526154ef9190810190618a39565b905060006154fd8683616b3a565b90506000846001600160a01b031663f45c1ce7836040518263ffffffff1660e01b815260040161552d9190617e91565b6000604051808303816000875af115801561554c573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526155749190810190618fdd565b805190915060030b1580159061558d5750602081015151155b801561559c5750604081015151155b1561457657816000815181106155b4576155b4618d9f565b6020026020010151604051602001613bcb9190619093565b606060006156018560408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b6040805180820182526000808252602091820152815180830190925286518252808701908201529091506156389082905b90616c8f565b156157955760006156b5826156af846156a961567b8a60408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b6040805180820182526000808252602091820152815180830190925282518252918201519181019190915290565b90616cb6565b90616d18565b604080518082018252600181527f0a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152909150615719908290616c8f565b1561578357604080518082018252600181527f0a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152615780905b8290616d9d565b90505b61578c81616dc3565b92505050613af8565b82156157ae578484604051602001613bcb92919061927f565b5050604080516020810190915260008152613af8565b509392505050565b6000808251602084016000f09392505050565b8160a00151156157ee57505050565b60006157fb848484616e2c565b90506000615808826153fd565b602081015181519192509060030b1580156158a45750604080518082018252600781527f5355434345535300000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526158a490604080518082018252600080825260209182015281518083019092528451825280850190820152615632565b156158b157505050505050565b604082015151156158d1578160400151604051602001613bcb9190619326565b80604051602001613bcb9190619384565b606060006159178360408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600481527f2e736f6c000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015290915061597c905b8290615f4a565b156159eb57604080518082018252600481527f2e736f6c0000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152613af8906159e69083906173c7565b616dc3565b604080518082018252600181527f3a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152615a4d905b8290617451565b600103615b1a57604080518082018252600181527f3a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152615ab390615779565b50604080518082018252600181527f3a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152613af8906159e6905b8390616d9d565b604080518082018252600581527f2e6a736f6e00000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152615b7990615975565b15615cb057604080518082018252600181527f2f00000000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820181905284518086019095529251845283015290615be19083906174eb565b905060008160018351615bf4919061860f565b81518110615c0457615c04618d9f565b60200260200101519050615ca76159e6615c7a6040518060400160405280600581526020017f2e6a736f6e00000000000000000000000000000000000000000000000000000081525060408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600080825260209182015281518083019092528551825280860190820152906173c7565b95945050505050565b82604051602001613bcb91906193ef565b50919050565b60606000615cfc8360408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600481527f2e736f6c0000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152909150615d5e90615975565b15615d6c57613af881616dc3565b604080518082018252600181527f3a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152615dcb90615a46565b600103615e3557604080518082018252600181527f3a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152613af8906159e690615b13565b604080518082018252600581527f2e6a736f6e00000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152615e9490615975565b15615cb057604080518082018252600181527f2f00000000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820181905284518086019095529251845283015290615efc9083906174eb565b9050600181511115615f38578060028251615f17919061860f565b81518110615f2757615f27618d9f565b602002602001015192505050919050565b5082604051602001613bcb91906193ef565b805182516000911115615f5f575060006139ff565b81518351602085015160009291615f75916194cd565b615f7f919061860f565b905082602001518103615f965760019150506139ff565b82516020840151819020912014905092915050565b60606000615fb883617590565b600101905060008167ffffffffffffffff811115615fd857615fd86180b6565b6040519080825280601f01601f191660200182016040528015616002576020820181803683370190505b5090508181016020015b600019017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a850494508461600c57509392505050565b604081810151815180830183526000808252602091820181905283518085018552835181529282018383015283518085018552600a81527f554e4c4943454e53454400000000000000000000000000000000000000000000818401908152855180870187528381528401929092528451808601909552518452908301526060916160d7905b8290617672565b1561611757505060408051808201909152600481527f4e6f6e65000000000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600981527f556e6c6963656e7365000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152616176906160d0565b156161b657505060408051808201909152600981527f556e6c6963656e736500000000000000000000000000000000000000000000006020820152919050565b604080518082018252600381527f4d4954000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152616215906160d0565b1561625557505060408051808201909152600381527f4d495400000000000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f47504c2d322e302d6f6e6c790000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526162b4906160d0565b806163195750604080518082018252601081527f47504c2d322e302d6f722d6c617465720000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152616319906160d0565b1561635957505060408051808201909152600981527f474e552047504c763200000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f47504c2d332e302d6f6e6c790000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526163b8906160d0565b8061641d5750604080518082018252601081527f47504c2d332e302d6f722d6c61746572000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261641d906160d0565b1561645d57505060408051808201909152600981527f474e552047504c763300000000000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4c47504c2d322e312d6f6e6c7900000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526164bc906160d0565b806165215750604080518082018252601181527f4c47504c2d322e312d6f722d6c6174657200000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152616521906160d0565b1561656157505060408051808201909152600c81527f474e55204c47504c76322e3100000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4c47504c2d332e302d6f6e6c7900000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526165c0906160d0565b806166255750604080518082018252601181527f4c47504c2d332e302d6f722d6c6174657200000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152616625906160d0565b1561666557505060408051808201909152600a81527f474e55204c47504c7633000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f4253442d322d436c617573650000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526166c4906160d0565b1561670457505060408051808201909152600c81527f4253442d322d436c6175736500000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f4253442d332d436c61757365000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152616763906160d0565b156167a357505060408051808201909152600c81527f4253442d332d436c6175736500000000000000000000000000000000000000006020820152919050565b604080518082018252600781527f4d504c2d322e300000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152616802906160d0565b1561684257505060408051808201909152600781527f4d504c2d322e30000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600781527f4f534c2d332e3000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526168a1906160d0565b156168e157505060408051808201909152600781527f4f534c2d332e30000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600a81527f4170616368652d322e300000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152616940906160d0565b1561698057505060408051808201909152600a81527f4170616368652d322e30000000000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4147504c2d332e302d6f6e6c7900000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526169df906160d0565b80616a445750604080518082018252601181527f4147504c2d332e302d6f722d6c6174657200000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152616a44906160d0565b15616a8457505060408051808201909152600a81527f474e55204147504c7633000000000000000000000000000000000000000000006020820152919050565b604080518082018252600881527f4255534c2d312e3100000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152616ae3906160d0565b15616b2357505060408051808201909152600781527f42534c20312e31000000000000000000000000000000000000000000000000006020820152919050565b60408084015184519151613bcb92906020016194e0565b60608060005b8451811015616bc55781858281518110616b5c57616b5c618d9f565b6020026020010151604051602001616b75929190618918565b604051602081830303815290604052915060018551616b94919061860f565b8114616bbd5781604051602001616bab9190619649565b60405160208183030381529060405291505b600101616b40565b5060408051600380825260808201909252600091816020015b6060815260200190600190039081616bde5790505090508381600081518110616c0957616c09618d9f565b60200260200101819052506040518060400160405280600281526020017f2d6300000000000000000000000000000000000000000000000000000000000081525081600181518110616c5d57616c5d618d9f565b60200260200101819052508181600281518110616c7c57616c7c618d9f565b6020908102919091010152949350505050565b6020808301518351835192840151600093616cad9291849190617686565b14159392505050565b60408051808201909152600080825260208201526000616ce88460000151856020015185600001518660200151617797565b9050836020015181616cfa919061860f565b84518590616d0990839061860f565b90525060208401525090919050565b6040805180820190915260008082526020820152815183511015616d3d5750816139ff565b6020808301519084015160019114616d645750815160208481015190840151829020919020145b8015616d9557825184518590616d7b90839061860f565b9052508251602085018051616d919083906194cd565b9052505b509192915050565b6040805180820190915260008082526020820152616dbc8383836178b7565b5092915050565b60606000826000015167ffffffffffffffff811115616de457616de46180b6565b6040519080825280601f01601f191660200182016040528015616e0e576020820181803683370190505b5090506000602082019050616dbc8185602001518660000151617962565b60606000616e38613df5565b6040805160ff808252612000820190925291925060009190816020015b6060815260200190600190039081616e5557905050905060006040518060400160405280600381526020017f6e70780000000000000000000000000000000000000000000000000000000000815250828280616eb090618f0c565b935060ff1681518110616ec557616ec5618d9f565b60200260200101819052506040518060400160405280600781526020017f5e312e33322e3300000000000000000000000000000000000000000000000000815250604051602001616f16919061968a565b604051602081830303815290604052828280616f3190618f0c565b935060ff1681518110616f4657616f46618d9f565b60200260200101819052506040518060400160405280600881526020017f76616c6964617465000000000000000000000000000000000000000000000000815250828280616f9390618f0c565b935060ff1681518110616fa857616fa8618d9f565b602002602001018190525082604051602001616fc49190618e3a565b604051602081830303815290604052828280616fdf90618f0c565b935060ff1681518110616ff457616ff4618d9f565b60200260200101819052506040518060400160405280600a81526020017f2d2d636f6e74726163740000000000000000000000000000000000000000000081525082828061704190618f0c565b935060ff168151811061705657617056618d9f565b602002602001018190525061706b87846179dc565b828261707681618f0c565b935060ff168151811061708b5761708b618d9f565b6020908102919091010152855151156171375760408051808201909152600b81527f2d2d7265666572656e6365000000000000000000000000000000000000000000602082015282826170dd81618f0c565b935060ff16815181106170f2576170f2618d9f565b602002602001018190525061710b8660000151846179dc565b828261711681618f0c565b935060ff168151811061712b5761712b618d9f565b60200260200101819052505b8560800151156171a55760408051808201909152601881527f2d2d756e73616665536b697053746f72616765436865636b00000000000000006020820152828261718081618f0c565b935060ff168151811061719557617195618d9f565b602002602001018190525061720b565b841561720b5760408051808201909152601281527f2d2d726571756972655265666572656e63650000000000000000000000000000602082015282826171ea81618f0c565b935060ff16815181106171ff576171ff618d9f565b60200260200101819052505b604086015151156172a75760408051808201909152600d81527f2d2d756e73616665416c6c6f77000000000000000000000000000000000000006020820152828261725581618f0c565b935060ff168151811061726a5761726a618d9f565b6020026020010181905250856040015182828061728690618f0c565b935060ff168151811061729b5761729b618d9f565b60200260200101819052505b8560600151156173115760408051808201909152601481527f2d2d756e73616665416c6c6f7752656e616d6573000000000000000000000000602082015282826172f081618f0c565b935060ff168151811061730557617305618d9f565b60200260200101819052505b60008160ff1667ffffffffffffffff81111561732f5761732f6180b6565b60405190808252806020026020018201604052801561736257816020015b606081526020019060019003908161734d5790505b50905060005b8260ff168160ff1610156173bb57838160ff168151811061738b5761738b618d9f565b6020026020010151828260ff16815181106173a8576173a8618d9f565b6020908102919091010152600101617368565b50979650505050505050565b60408051808201909152600080825260208201528151835110156173ec5750816139ff565b81518351602085015160009291617402916194cd565b61740c919061860f565b6020840151909150600190821461742d575082516020840151819020908220145b80156174485783518551869061744490839061860f565b9052505b50929392505050565b60008082600001516174758560000151866020015186600001518760200151617797565b61747f91906194cd565b90505b8351602085015161749391906194cd565b8111616dbc57816174a3816196cf565b92505082600001516174da8560200151836174be919061860f565b86516174ca919061860f565b8386600001518760200151617797565b6174e491906194cd565b9050617482565b606060006174f98484617451565b6175049060016194cd565b67ffffffffffffffff81111561751c5761751c6180b6565b60405190808252806020026020018201604052801561754f57816020015b606081526020019060019003908161753a5790505b50905060005b81518110156157c45761756b6159e68686616d9d565b82828151811061757d5761757d618d9f565b6020908102919091010152600101617555565b6000807a184f03e93ff9f4daa797ed6e38ed64bf6a1f01000000000000000083106175d9577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef81000000008310617605576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc10000831061762357662386f26fc10000830492506010015b6305f5e100831061763b576305f5e100830492506008015b612710831061764f57612710830492506004015b60648310617661576064830492506002015b600a83106139ff5760010192915050565b600061767e8383617a1c565b159392505050565b60008085841161778d576020841161773957600084156176d15760016176ad86602061860f565b6176b89060086196e9565b6176c39060026197e7565b6176cd919061860f565b1990505b83518116856176e089896194cd565b6176ea919061860f565b805190935082165b8181146177245787841161770c57879450505050506153f5565b83617716816197f3565b9450508284511690506176f2565b61772e87856194cd565b9450505050506153f5565b838320617746858861860f565b61775090876194cd565b91505b85821061778b578482208082036177785761776e86846194cd565b93505050506153f5565b61778360018461860f565b925050617753565b505b5092949350505050565b600083818685116178a2576020851161785157600085156177e35760016177bf87602061860f565b6177ca9060086196e9565b6177d59060026197e7565b6177df919061860f565b1990505b845181166000876177f48b8b6194cd565b6177fe919061860f565b855190915083165b8281146178435781861061782b5761781e8b8b6194cd565b96505050505050506153f5565b85617835816196cf565b965050838651169050617806565b8596505050505050506153f5565b508383206000905b617863868961860f565b82116178a05785832080820361787f57839450505050506153f5565b61788a6001856194cd565b9350508180617898906196cf565b925050617859565b505b6178ac87876194cd565b979650505050505050565b604080518082019091526000808252602082015260006178e98560000151866020015186600001518760200151617797565b602080870180519186019190915251909150617905908261860f565b83528451602086015161791891906194cd565b81036179275760008552617959565b8351835161793591906194cd565b8551869061794490839061860f565b905250835161795390826194cd565b60208601525b50909392505050565b6020811061799a57815183526179796020846194cd565b92506179866020836194cd565b915061799360208261860f565b9050617962565b60001981156179c95760016179b083602061860f565b6179bc906101006197e7565b6179c6919061860f565b90505b9151835183169219169190911790915250565b606060006179ea8484613ec8565b8051602080830151604051939450617a049390910161980a565b60405160208183030381529060405291505092915050565b8151815160009190811115617a2f575081515b6020808501519084015160005b83811015617ae85782518251808214617ab8576000196020871015617a9757600184617a6989602061860f565b617a7391906194cd565b617a7e9060086196e9565b617a899060026197e7565b617a93919061860f565b1990505b8181168382168181039114617ab55797506139ff9650505050505050565b50505b617ac36020866194cd565b9450617ad06020856194cd565b93505050602081617ae191906194cd565b9050617a3c565b50845186516145769190619862565b610c9f8061988383390190565b610f2a8061a52283390190565b610aa98061b44c83390190565b610b3f8061bef583390190565b6120728061ca3483390190565b6040518060e00160405280606081526020016060815260200160608152602001600015158152602001600015158152602001600015158152602001617b7b617b80565b905290565b60405180610100016040528060001515815260200160001515815260200160608152602001600080191681526020016060815260200160608152602001600015158152602001617b7b6040518060800160405280600081526020016000815260200160008152602001600081525090565b602080825282518282018190526000918401906040840190835b81811015617c325783516001600160a01b0316835260209384019390920191600101617c0b565b509095945050505050565b60005b83811015617c58578181015183820152602001617c40565b50506000910152565b60008151808452617c79816020860160208601617c3d565b601f01601f19169290920160200192915050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b82811015617d89577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452815180516001600160a01b03168652602090810151604082880181905281519088018190529101906060600582901b88018101919088019060005b81811015617d6f577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08a8503018352617d59848651617c61565b6020958601959094509290920191600101617d1f565b509197505050602094850194929092019150600101617cb5565b50929695505050505050565b600081518084526020840193506020830160005b82811015617de95781517fffffffff0000000000000000000000000000000000000000000000000000000016865260209586019590910190600101617da9565b5093949350505050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b82811015617d89577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08786030184528151805160408752617e5f6040880182617c61565b9050602082015191508681036020880152617e7a8183617d95565b965050506020938401939190910190600101617e1b565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b82811015617d89577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452617ef3858351617c61565b94506020938401939190910190600101617eb9565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b82811015617d89577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845281516001600160a01b0381511686526020810151905060406020870152617f896040870182617d95565b9550506020938401939190910190600101617f30565b600060208284031215617fb157600080fd5b81518015158114613af857600080fd5b610100815260056101008201527f544f4b454e000000000000000000000000000000000000000000000000000000610120820152610140602082015260036101408201527f544b4e000000000000000000000000000000000000000000000000000000000061016082015260006101808201905060ff881660408301528660608301526003861061807b577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8560808301528460a083015261809c60c08301856001600160a01b03169052565b6001600160a01b03831660e0830152979650505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600181811c908216806180f957607f821691505b602082108103615cc1577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b601f82111561817c57806000526020600020601f840160051c810160208510156181595750805b601f840160051c820191505b818110156181795760008155600101618165565b50505b505050565b815167ffffffffffffffff81111561819b5761819b6180b6565b6181af816181a984546180e5565b84618132565b6020601f8211600181146181e357600083156181cb5750848201515b600019600385901b1c1916600184901b178455618179565b600084815260208120601f198516915b8281101561821357878501518255602094850194600190920191016181f3565b50848210156182315786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b60006020828403121561825257600080fd5b5051919050565b60608152600061826c6060830186617c61565b602083019490945250901515604090910152919050565b600081546001600160a01b038116845260ff8160a01c1615156020850152506001600160a01b0360018301541660408401526002820160a06060850152600081546182cd816180e5565b8060a088015260018216600081146182ec57600181146183265761835a565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00831660c089015260c082151560051b890101935061835a565b84600052602060002060005b838110156183515781548a820160c00152600190910190602001618332565b890160c0019450505b50505060038401546080860152809250505092915050565b60e08152600061838560e0830189617c61565b62ffffff881660208401526001600160a01b038716604084015282810360608401526183b18187617c61565b85546080850152600186015460ff16151560a085015290505b82810360c08401526183dc8185618283565b9998505050505050505050565b6001600160a01b0384168152826020820152606060408201526000615ca76060830184617c61565b60c08152600061842460c0830189617c61565b8760208401526001600160a01b0387166040840152828103606084015261844b8187617c61565b6080840195909552505090151560a090910152949350505050565b6001600160a01b038616815284602082015260a06040820152600061848e60a0830186617c61565b6060830194909452509015156080909101529392505050565b8281526040602082015260006153f56040830184617c61565b6001600160a01b0384541681526001600160a01b0383166020820152606060408201526000615ca76060830184617c61565b6040516060810167ffffffffffffffff81118282101715618515576185156180b6565b60405290565b60008067ffffffffffffffff841115618536576185366180b6565b50604051601f19601f85018116603f0116810181811067ffffffffffffffff82111715618565576185656180b6565b60405283815290508082840185101561857d57600080fd5b6157c4846020830185617c3d565b600082601f83011261859c57600080fd5b613af88383516020850161851b565b6000602082840312156185bd57600080fd5b815167ffffffffffffffff8111156185d457600080fd5b6139fb8482850161858b565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b818103818111156139ff576139ff6185e0565b6101208152600061863761012083018b617c61565b6001600160a01b038a16602084015288604084015287606084015286608084015282810360a084015261866a8187617c61565b855160c08501526020860151151560e08501529050828103610100840152613ceb8185618283565b60e0815260006186a560e0830189617c61565b8760208401526001600160a01b038716604084015282810360608401526186cc8187617c61565b855160808501526020860151151560a085015290506183ca565b60a0815260006186f960a0830187617c61565b828103602084015261870b8187617c61565b85516040850152602086015115156060850152905082810360808401526178ac8185618283565b60c08152600061874560c0830188617c61565b6001600160a01b038716602084015282810360408401526187668187617c61565b8551606085015260208601511515608085015290505b82810360a084015261878e8185618283565b98975050505050505050565b60c0815260006187ad60c0830188617c61565b6001600160a01b038716602084015282810360408401526187ce8187617c61565b85546060850152600186015460ff1615156080850152905061877c565b60a0815260006187fe60a0830188617c61565b6001600160a01b0387166020840152828103604084015261881f8187617c61565b606084019590955250509015156080909101529392505050565b6001600160a01b03831681526040602082015260006153f56040830184617c61565b7f4661696c656420746f206465706c6f7920636f6e74726163742000000000000081526000835161889381601a850160208801617c3d565b7f3a20000000000000000000000000000000000000000000000000000000000000601a9184019182015283516188d081601c840160208801617c3d565b01601c01949350505050565b602081526000613af86020830184617c61565b60006020828403121561890157600080fd5b81516001600160a01b0381168114613af857600080fd5b6000835161892a818460208801617c3d565b83519083019061893e818360208801617c3d565b01949350505050565b7f4661696c656420746f206465706c6f7920636f6e74726163742000000000000081526000835161897f81601a850160208801617c3d565b7f207573696e6720636f6e7374727563746f722064617461202200000000000000601a9184019182015283516189bc816033840160208801617c3d565b7f220000000000000000000000000000000000000000000000000000000000000060339290910191820152603401949350505050565b60408152600b60408201527f464f554e4452595f4f55540000000000000000000000000000000000000000006060820152608060208201526000613af86080830184617c61565b600060208284031215618a4b57600080fd5b815167ffffffffffffffff811115618a6257600080fd5b8201601f81018413618a7357600080fd5b6139fb8482516020840161851b565b60008551618a94818460208a01617c3d565b7f2f000000000000000000000000000000000000000000000000000000000000009083019081528551618ace816001840160208a01617c3d565b7f2f00000000000000000000000000000000000000000000000000000000000000600192909101918201528451618b0c816002840160208901617c3d565b6001818301019150507f2f0000000000000000000000000000000000000000000000000000000000000060018201528351618b4e816002840160208801617c3d565b7f2e6a736f6e000000000000000000000000000000000000000000000000000000600292909101918201526007019695505050505050565b604081526000618b996040830184617c61565b8281036020840152600481527f2e6173740000000000000000000000000000000000000000000000000000000060208201526040810191505092915050565b7f436f756c64206e6f742066696e642041535420696e2061727469666163742000815260008251618c1081601f850160208701617c3d565b7f2e205365742060617374203d20747275656020696e20666f756e6472792e746f601f9390910192830152507f6d6c000000000000000000000000000000000000000000000000000000000000603f820152604101919050565b604081526000618c7d6040830184617c61565b8281036020840152601181527f2e6173742e6162736f6c7574655061746800000000000000000000000000000060208201526040810191505092915050565b604081526000618ccf6040830184617c61565b8281036020840152600c81527f2e6173742e6c6963656e7365000000000000000000000000000000000000000060208201526040810191505092915050565b7f2e6d657461646174612e736f75726365732e5b27000000000000000000000000815260008251618d46816014850160208701617c3d565b7f275d2e6b656363616b32353600000000000000000000000000000000000000006014939091019283015250602001919050565b604081526000618d8d6040830185617c61565b8281036020840152613af48185617c61565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f2200000000000000000000000000000000000000000000000000000000000000815260008251618e06816001850160208701617c3d565b7f22000000000000000000000000000000000000000000000000000000000000006001939091019283015250600201919050565b60008251618e4c818460208701617c3d565b7f2f6275696c642d696e666f000000000000000000000000000000000000000000920191825250600b01919050565b7f436f756c64206e6f742066696e64206275696c642d696e666f2066696c65207781527f697468206d61746368696e6720736f7572636520636f6465206861736820666f60208201527f7220636f6e747261637420000000000000000000000000000000000000000000604082015260008251618eff81604b850160208701617c3d565b91909101604b0192915050565b600060ff821660ff8103618f2257618f226185e0565b60010192915050565b7f406f70656e7a657070656c696e2f646566656e6465722d6465706c6f792d636c81527f69656e742d636c69400000000000000000000000000000000000000000000000602082015260008251618f89816029850160208701617c3d565b9190910160290192915050565b60408152601660408201527f4f50454e5a455050454c494e5f424153485f50415448000000000000000000006060820152608060208201526000613af86080830184617c61565b600060208284031215618fef57600080fd5b815167ffffffffffffffff81111561900657600080fd5b82016060818503121561901857600080fd5b6190206184f2565b81518060030b811461903157600080fd5b8152602082015167ffffffffffffffff81111561904d57600080fd5b6190598682850161858b565b602083015250604082015167ffffffffffffffff81111561907957600080fd5b6190858682850161858b565b604083015250949350505050565b7f4661696c656420746f2072756e206261736820636f6d6d616e6420776974682081527f22000000000000000000000000000000000000000000000000000000000000006020820152600082516190f1816021850160208701617c3d565b7f222e20496620796f7520617265207573696e672057696e646f77732c2073657460219390910192830152507f20746865204f50454e5a455050454c494e5f424153485f5041544820656e766960418201527f726f6e6d656e74207661726961626c6520746f207468652066756c6c7920717560618201527f616c69666965642070617468206f66207468652062617368206578656375746160818201527f626c652e20466f72206578616d706c652c20696620796f75206172652075736960a18201527f6e672047697420666f722057696e646f77732c206164642074686520666f6c6c60c18201527f6f77696e67206c696e6520696e20746865202e656e762066696c65206f66207960e18201527f6f75722070726f6a65637420287573696e6720666f727761726420736c6173686101018201527f6573293a0a4f50454e5a455050454c494e5f424153485f504154483d22433a2f6101218201527f50726f6772616d2046696c65732f4769742f62696e2f6261736822000000000061014182015261015c01919050565b7f4661696c656420746f2066696e64206c696e652077697468207072656669782081527f27000000000000000000000000000000000000000000000000000000000000006020820152600083516192dd816021850160208801617c3d565b7f2720696e206f75747075743a2000000000000000000000000000000000000000602191840191820152835161931a81602e840160208801617c3d565b01602e01949350505050565b7f4661696c656420746f2072756e2075706772616465207361666574792076616c81527f69646174696f6e3a200000000000000000000000000000000000000000000000602082015260008251618f89816029850160208701617c3d565b7f55706772616465207361666574792076616c69646174696f6e206661696c656481527f3a0a0000000000000000000000000000000000000000000000000000000000006020820152600082516193e2816022850160208701617c3d565b9190910160220192915050565b7f436f6e7472616374206e616d652000000000000000000000000000000000000081526000825161942781600e850160208701617c3d565b7f206d75737420626520696e2074686520666f726d6174204d79436f6e74726163600e9390910192830152507f742e736f6c3a4d79436f6e7472616374206f72204d79436f6e74726163742e73602e8201527f6f6c206f72206f75742f4d79436f6e74726163742e736f6c2f4d79436f6e7472604e8201527f6163742e6a736f6e000000000000000000000000000000000000000000000000606e820152607601919050565b808201808211156139ff576139ff6185e0565b7f53504458206c6963656e7365206964656e746966696572200000000000000000815260008351619518816018850160208801617c3d565b7f20696e2000000000000000000000000000000000000000000000000000000000601891840191820152835161955581601c840160208801617c3d565b7f20646f6573206e6f74206c6f6f6b206c696b65206120737570706f7274656420601c92909101918201527f6c6963656e736520666f7220626c6f636b206578706c6f726572207665726966603c8201527f69636174696f6e2e205573652074686520606c6963656e73655479706560206f605c8201527f7074696f6e20746f20737065636966792061206c6963656e736520747970652c607c8201527f206f7220736574207468652060736b69704c6963656e73655479706560206f70609c8201527f74696f6e20746f2060747275656020746f20736b69702e00000000000000000060bc82015260d301949350505050565b6000825161965b818460208701617c3d565b7f2000000000000000000000000000000000000000000000000000000000000000920191825250600101919050565b7f406f70656e7a657070656c696e2f75706772616465732d636f726540000000008152600082516196c281601c850160208701617c3d565b91909101601c0192915050565b600060001982036196e2576196e26185e0565b5060010190565b80820281158282048414176139ff576139ff6185e0565b6001815b600184111561973b5780850481111561971f5761971f6185e0565b600184161561972d57908102905b60019390931c928002619704565b935093915050565b600082619752575060016139ff565b8161975f575060006139ff565b8160018114619775576002811461977f5761979b565b60019150506139ff565b60ff841115619790576197906185e0565b50506001821b6139ff565b5060208310610133831016604e8410600b84101617156197be575081810a6139ff565b6197cb6000198484619700565b80600019048211156197df576197df6185e0565b029392505050565b6000613af88383619743565b600081619802576198026185e0565b506000190190565b6000835161981c818460208801617c3d565b7f3a000000000000000000000000000000000000000000000000000000000000009083019081528351619856816001840160208801617c3d565b01600101949350505050565b8181036000831280158383131683831282161715616dbc57616dbc6185e056fe608060405234801561001057600080fd5b50604051610c9f380380610c9f83398101604081905261002f9161010d565b8181600361003d83826101ff565b50600461004a82826101ff565b50505050506102bd565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261007b57600080fd5b81516001600160401b0381111561009457610094610054565b604051601f8201601f19908116603f011681016001600160401b03811182821017156100c2576100c2610054565b6040528181528382016020018510156100da57600080fd5b60005b828110156100f9576020818601810151838301820152016100dd565b506000918101602001919091529392505050565b6000806040838503121561012057600080fd5b82516001600160401b0381111561013657600080fd5b6101428582860161006a565b602085015190935090506001600160401b0381111561016057600080fd5b61016c8582860161006a565b9150509250929050565b600181811c9082168061018a57607f821691505b6020821081036101aa57634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156101fa57806000526020600020601f840160051c810160208510156101d75750805b601f840160051c820191505b818110156101f757600081556001016101e3565b50505b505050565b81516001600160401b0381111561021857610218610054565b61022c816102268454610176565b846101b0565b6020601f82116001811461026057600083156102485750848201515b600019600385901b1c1916600184901b1784556101f7565b600084815260208120601f198516915b828110156102905787850151825560209485019460019092019101610270565b50848210156102ae5786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b6109d3806102cc6000396000f3fe608060405234801561001057600080fd5b50600436106100be5760003560e01c806340c10f191161007657806395d89b411161005b57806395d89b4114610183578063a9059cbb1461018b578063dd62ed3e1461019e57600080fd5b806340c10f191461013857806370a082311461014d57600080fd5b806318160ddd116100a757806318160ddd1461010457806323b872dd14610116578063313ce5671461012957600080fd5b806306fdde03146100c3578063095ea7b3146100e1575b600080fd5b6100cb6101e4565b6040516100d891906107bf565b60405180910390f35b6100f46100ef366004610854565b610276565b60405190151581526020016100d8565b6002545b6040519081526020016100d8565b6100f461012436600461087e565b610290565b604051601281526020016100d8565b61014b610146366004610854565b6102b4565b005b61010861015b3660046108bb565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b6100cb6102c2565b6100f4610199366004610854565b6102d1565b6101086101ac3660046108dd565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b6060600380546101f390610910565b80601f016020809104026020016040519081016040528092919081815260200182805461021f90610910565b801561026c5780601f106102415761010080835404028352916020019161026c565b820191906000526020600020905b81548152906001019060200180831161024f57829003601f168201915b5050505050905090565b6000336102848185856102df565b60019150505b92915050565b60003361029e8582856102f1565b6102a98585856103c5565b506001949350505050565b6102be8282610470565b5050565b6060600480546101f390610910565b6000336102848185856103c5565b6102ec83838360016104cc565b505050565b73ffffffffffffffffffffffffffffffffffffffff8381166000908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146103bf57818110156103b0576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8416600482015260248101829052604481018390526064015b60405180910390fd5b6103bf848484840360006104cc565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8316610415576040517f96c6fd1e000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b73ffffffffffffffffffffffffffffffffffffffff8216610465576040517fec442f05000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b6102ec838383610614565b73ffffffffffffffffffffffffffffffffffffffff82166104c0576040517fec442f05000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b6102be60008383610614565b73ffffffffffffffffffffffffffffffffffffffff841661051c576040517fe602df05000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b73ffffffffffffffffffffffffffffffffffffffff831661056c576040517f94280d62000000000000000000000000000000000000000000000000000000008152600060048201526024016103a7565b73ffffffffffffffffffffffffffffffffffffffff808516600090815260016020908152604080832093871683529290522082905580156103bf578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258460405161060691815260200190565b60405180910390a350505050565b73ffffffffffffffffffffffffffffffffffffffff831661064c5780600260008282546106419190610963565b909155506106fe9050565b73ffffffffffffffffffffffffffffffffffffffff8316600090815260208190526040902054818110156106d2576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8516600482015260248101829052604481018390526064016103a7565b73ffffffffffffffffffffffffffffffffffffffff841660009081526020819052604090209082900390555b73ffffffffffffffffffffffffffffffffffffffff821661072757600280548290039055610753565b73ffffffffffffffffffffffffffffffffffffffff821660009081526020819052604090208054820190555b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516107b291815260200190565b60405180910390a3505050565b602081526000825180602084015260005b818110156107ed57602081860181015160408684010152016107d0565b5060006040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461084f57600080fd5b919050565b6000806040838503121561086757600080fd5b6108708361082b565b946020939093013593505050565b60008060006060848603121561089357600080fd5b61089c8461082b565b92506108aa6020850161082b565b929592945050506040919091013590565b6000602082840312156108cd57600080fd5b6108d68261082b565b9392505050565b600080604083850312156108f057600080fd5b6108f98361082b565b91506109076020840161082b565b90509250929050565b600181811c9082168061092457607f821691505b60208210810361095d577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b8082018082111561028a577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fdfea2646970667358221220837c7d9916de10b20cdb18567d8c7679613426bbd7b0b72548d8000a412f307b64736f6c634300081a00336080604052348015600f57600080fd5b506001600055610f06806100246000396000f3fe60806040526004361061006e5760003560e01c8063c51316911161004b578063c5131691146100d5578063c9028a36146100f5578063e04d4f9714610115578063f05b6abf1461012857005b8063357fc5a214610077578063676cc054146100975780636ed70169146100c057005b3661007557005b005b34801561008357600080fd5b50610075610092366004610709565b610148565b6100aa6100a5366004610745565b6101de565b6040516100b79190610840565b60405180910390f35b3480156100cc57600080fd5b5061007561023f565b3480156100e157600080fd5b506100756100f0366004610709565b610274565b34801561010157600080fd5b50610075610110366004610853565b61034f565b6100756101233660046109b3565b61038b565b34801561013457600080fd5b50610075610143366004610a9f565b6103cf565b610150610404565b61017273ffffffffffffffffffffffffffffffffffffffff8316338386610447565b604080513381526020810185905273ffffffffffffffffffffffffffffffffffffffff848116828401528316606082015290517f2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af609181900360800190a16101d96001600055565b505050565b60607fd80b62959d9a7e797f352e4015e65d345f402ea21972256fb0ba94f00a35250161020e6020860186610b89565b848460405161021f93929190610bed565b60405180910390a1506040805160208101909152600081525b9392505050565b6040513381527fbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a09060200160405180910390a1565b61027c610404565b6000610289600285610c26565b9050806000036102c5576040517f1f2a200500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6102e773ffffffffffffffffffffffffffffffffffffffff8416338484610447565b604080513381526020810183905273ffffffffffffffffffffffffffffffffffffffff858116828401528416606082015290517f2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af609181900360800190a1506101d96001600055565b7f689a5a5cb55e795ffe4cd8b419cd3bb0a3373974c54d25f64e734d7388b93e9b3382604051610380929190610c61565b60405180910390a150565b7f1f1ff1f5fb41346850b2f5c04e6c767e2f1c8a525c5c0c5cdb60cdf3ca5f62fa33348585856040516103c2959493929190610d53565b60405180910390a1505050565b7f74a53cd528a921fca7dbdee62f86819051d3cc98f214951f4238e8843f20b146338484846040516103c29493929190610ddd565b600260005403610440576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002600055565b6040805173ffffffffffffffffffffffffffffffffffffffff85811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd000000000000000000000000000000000000000000000000000000001790526104dc9085906104e2565b50505050565b600061050473ffffffffffffffffffffffffffffffffffffffff84168361057d565b905080516000141580156105295750808060200190518101906105279190610e97565b155b156101d9576040517f5274afe700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201526024015b60405180910390fd5b606061023883836000846000808573ffffffffffffffffffffffffffffffffffffffff1684866040516105b09190610eb4565b60006040518083038185875af1925050503d80600081146105ed576040519150601f19603f3d011682016040523d82523d6000602084013e6105f2565b606091505b509150915061060286838361060c565b9695505050505050565b6060826106215761061c8261069b565b610238565b8151158015610645575073ffffffffffffffffffffffffffffffffffffffff84163b155b15610694576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401610574565b5080610238565b8051156106ab5780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50565b803573ffffffffffffffffffffffffffffffffffffffff8116811461070457600080fd5b919050565b60008060006060848603121561071e57600080fd5b8335925061072e602085016106e0565b915061073c604085016106e0565b90509250925092565b6000806000838503604081121561075b57600080fd5b602081121561076957600080fd5b50839250602084013567ffffffffffffffff81111561078757600080fd5b8401601f8101861361079857600080fd5b803567ffffffffffffffff8111156107af57600080fd5b8660208284010111156107c157600080fd5b939660209190910195509293505050565b60005b838110156107ed5781810151838201526020016107d5565b50506000910152565b6000815180845261080e8160208601602086016107d2565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60208152600061023860208301846107f6565b60006020828403121561086557600080fd5b813567ffffffffffffffff81111561087c57600080fd5b82016080818503121561023857600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156109045761090461088e565b604052919050565b600082601f83011261091d57600080fd5b813567ffffffffffffffff8111156109375761093761088e565b61096860207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116016108bd565b81815284602083860101111561097d57600080fd5b816020850160208301376000918101602001919091529392505050565b80151581146106dd57600080fd5b80356107048161099a565b6000806000606084860312156109c857600080fd5b833567ffffffffffffffff8111156109df57600080fd5b6109eb8682870161090c565b935050602084013591506040840135610a038161099a565b809150509250925092565b600067ffffffffffffffff821115610a2857610a2861088e565b5060051b60200190565b600082601f830112610a4357600080fd5b8135610a56610a5182610a0e565b6108bd565b8082825260208201915060208360051b860101925085831115610a7857600080fd5b602085015b83811015610a95578035835260209283019201610a7d565b5095945050505050565b600080600060608486031215610ab457600080fd5b833567ffffffffffffffff811115610acb57600080fd5b8401601f81018613610adc57600080fd5b8035610aea610a5182610a0e565b8082825260208201915060208360051b850101925088831115610b0c57600080fd5b602084015b83811015610b4e57803567ffffffffffffffff811115610b3057600080fd5b610b3f8b60208389010161090c565b84525060209283019201610b11565b509550505050602084013567ffffffffffffffff811115610b6e57600080fd5b610b7a86828701610a32565b92505061073c604085016109a8565b600060208284031215610b9b57600080fd5b610238826106e0565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff84168152604060208201526000610c1d604083018486610ba4565b95945050505050565b600082610c5c577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b73ffffffffffffffffffffffffffffffffffffffff831681526040602082015273ffffffffffffffffffffffffffffffffffffffff610c9f836106e0565b16604082015273ffffffffffffffffffffffffffffffffffffffff610cc6602084016106e0565b166060820152600080604084013590508060808401525060608301357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112610d1257600080fd5b830160208101903567ffffffffffffffff811115610d2f57600080fd5b803603821315610d3e57600080fd5b608060a085015261060260c085018284610ba4565b73ffffffffffffffffffffffffffffffffffffffff8616815284602082015260a060408201526000610d8860a08301866107f6565b6060830194909452509015156080909101529392505050565b600081518084526020840193506020830160005b82811015610dd3578151865260209586019590910190600101610db5565b5093949350505050565b60006080820173ffffffffffffffffffffffffffffffffffffffff871683526080602084015280865180835260a08501915060a08160051b86010192506020880160005b82811015610e70577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60878603018452610e5b8583516107f6565b94506020938401939190910190600101610e21565b505050508281036040840152610e868186610da1565b915050610c1d606083018415159052565b600060208284031215610ea957600080fd5b81516102388161099a565b60008251610ec68184602087016107d2565b919091019291505056fea2646970667358221220e589ce064ae2868bfa054c8896cb2f4e7139370860040db8f5a8cf19a653059164736f6c634300081a00336080604052348015600f57600080fd5b50604051610aa9380380610aa9833981016040819052602c916050565b600080546001600160a01b0319166001600160a01b0392909216919091179055607e565b600060208284031215606157600080fd5b81516001600160a01b0381168114607757600080fd5b9392505050565b610a1c8061008d6000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80630abd890514610046578063116191b61461005b5780637a34d8bb146100a4575b600080fd5b610059610054366004610658565b6100b7565b005b60005461007b9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b6100596100b23660046106f8565b61032a565b60008383836040516024016100ce939291906107f4565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe04d4f970000000000000000000000000000000000000000000000000000000017905260005490915073ffffffffffffffffffffffffffffffffffffffff8087169163095ea7b3911661017589600261081e565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16815273ffffffffffffffffffffffffffffffffffffffff909216600483015260248201526044016020604051808303816000875af11580156101e5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610209919061085e565b61023f576040517f8164f84200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805160a081018252610321808252600160208084018290528385019290925283518083018552600080825260608501919091526080840181905284518086018652918252918101829052905492517f7b15118b0000000000000000000000000000000000000000000000000000000081529192909173ffffffffffffffffffffffffffffffffffffffff90911690637b15118b906102ed908c908c908c90899088908a906004016108f7565b600060405180830381600087803b15801561030757600080fd5b505af115801561031b573d6000803e3d6000fd5b50505050505050505050505050565b6000838383604051602401610341939291906107f4565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152602080830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe04d4f9700000000000000000000000000000000000000000000000000000000179052815160a081018352610321808252600182840181905282850191909152835180840185526000808252606084019190915260808301819052845180860186528281529384018190525493517f095ea7b300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9485166004820152602481019190915293945092909188169063095ea7b3906044016020604051808303816000875af115801561047c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104a0919061085e565b506000546040517f06cb898300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff909116906306cb8983906104ff908b908b90889087908990600401610972565b600060405180830381600087803b15801561051957600080fd5b505af115801561052d573d6000803e3d6000fd5b505050505050505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f83011261057b57600080fd5b81356020830160008067ffffffffffffffff84111561059c5761059c61053b565b506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f85018116603f0116810181811067ffffffffffffffff821117156105e9576105e961053b565b60405283815290508082840187101561060157600080fd5b838360208301376000602085830101528094505050505092915050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461064257600080fd5b919050565b801515811461065557600080fd5b50565b60008060008060008060c0878903121561067157600080fd5b863567ffffffffffffffff81111561068857600080fd5b61069489828a0161056a565b965050602087013594506106aa6040880161061e565b9350606087013567ffffffffffffffff8111156106c657600080fd5b6106d289828a0161056a565b9350506080870135915060a08701356106ea81610647565b809150509295509295509295565b600080600080600060a0868803121561071057600080fd5b853567ffffffffffffffff81111561072757600080fd5b6107338882890161056a565b9550506107426020870161061e565b9350604086013567ffffffffffffffff81111561075e57600080fd5b61076a8882890161056a565b93505060608601359150608086013561078281610647565b809150509295509295909350565b6000815180845260005b818110156107b65760208185018101518683018201520161079a565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b6060815260006108076060830186610790565b602083019490945250901515604090910152919050565b80820180821115610858577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b92915050565b60006020828403121561087057600080fd5b815161087b81610647565b9392505050565b73ffffffffffffffffffffffffffffffffffffffff815116825260208101511515602083015273ffffffffffffffffffffffffffffffffffffffff60408201511660408301526000606082015160a060608501526108e360a0850182610790565b608093840151949093019390935250919050565b60e08152600061090a60e0830189610790565b87602084015273ffffffffffffffffffffffffffffffffffffffff87166040840152828103606084015261093e8187610790565b855160808501526020860151151560a0850152905082810360c08401526109658185610882565b9998505050505050505050565b60c08152600061098560c0830188610790565b73ffffffffffffffffffffffffffffffffffffffff8716602084015282810360408401526109b38187610790565b85516060850152602086015115156080850152905082810360a08401526109da8185610882565b9897505050505050505056fea2646970667358221220b235f8065ec62b494fd6d715f8aaded6a4ba95478c5fde21640b6bc5fbfc444364736f6c634300081a0033608060405234801561001057600080fd5b50604051610b3f380380610b3f83398101604081905261002f916100b9565b600380546001600160a01b038086166001600160a01b0319928316179092556004805485841690831617905560058054928416929091169190911790556040517f80699e81136d69cb8367ad52a994e25c722a86da654b561d0c14b61a777e7ac590600090a15050506100fc565b80516001600160a01b03811681146100b457600080fd5b919050565b6000806000606084860312156100ce57600080fd5b6100d78461009d565b92506100e56020850161009d565b91506100f36040850161009d565b90509250925092565b610a348061010b6000396000f3fe608060405234801561001057600080fd5b50600436106100c95760003560e01c806397770dff11610081578063d7fd7afb1161005b578063d7fd7afb146101f2578063d936a01214610220578063ee2815ba1461024057600080fd5b806397770dff146101b9578063a7cb0507146101cc578063c63585cc146101df57600080fd5b8063513a9c05116100b2578063513a9c0514610143578063569541b914610179578063842da36d1461019957600080fd5b80630be15547146100ce5780633c669d551461012e575b600080fd5b6101046100dc36600461071e565b60016020526000908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b61014161013c366004610760565b610253565b005b61010461015136600461071e565b60026020526000908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b6003546101049073ffffffffffffffffffffffffffffffffffffffff1681565b6005546101049073ffffffffffffffffffffffffffffffffffffffff1681565b6101416101c73660046107fd565b6103a0565b6101416101da36600461081f565b610419565b6101046101ed366004610841565b610467565b61021261020036600461071e565b60006020819052908152604090205481565b604051908152602001610125565b6004546101049073ffffffffffffffffffffffffffffffffffffffff1681565b61014161024e366004610884565b61059c565b604080516080810182526000606082019081528152336020820152468183015290517fa9059cbb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff87811660048301526024820186905286169063a9059cbb906044016020604051808303816000875af11580156102e7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061030b91906108b0565b506040517fde43156e00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff87169063de43156e90610366908490899089908990899060040161091b565b600060405180830381600087803b15801561038057600080fd5b505af1158015610394573d6000803e3d6000fd5b50505050505050505050565b600380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fdba79d534382d1a8ae108e4c8ecb27c6ae42ab8b91d44eedf88bd329f3868d5e9060200160405180910390a150565b6000828152602081815260409182902083905581518481529081018390527f49f492222906ac486c3c1401fa545626df1f0c0e5a77a05597ea2ed66af9850d91015b60405180910390a15050565b60008060006104768585610620565b6040517fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606084811b8216602084015283901b166034820152919350915086906048016040516020818303038152906040528051906020012060405160200161055c9291907fff00000000000000000000000000000000000000000000000000000000000000815260609290921b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016600183015260158201527f96e8ac4277198ff8b6f785478aa9a39f403cb768dd02cbee326c3e7da348845f603582015260550190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101209695505050505050565b60008281526001602090815260409182902080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff85169081179091558251858152918201527fd1b36d30f6248e97c473b4d1348ca164a4ef6759022f54a58ec200326c39c45d910161045b565b6000808273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1603610688576040517fcb1e7cfe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16106106c25782846106c5565b83835b909250905073ffffffffffffffffffffffffffffffffffffffff8216610717576040517f78b507da00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b9250929050565b60006020828403121561073057600080fd5b5035919050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461075b57600080fd5b919050565b60008060008060006080868803121561077857600080fd5b61078186610737565b945061078f60208701610737565b935060408601359250606086013567ffffffffffffffff8111156107b257600080fd5b8601601f810188136107c357600080fd5b803567ffffffffffffffff8111156107da57600080fd5b8860208284010111156107ec57600080fd5b959894975092955050506020019190565b60006020828403121561080f57600080fd5b61081882610737565b9392505050565b6000806040838503121561083257600080fd5b50508035926020909101359150565b60008060006060848603121561085657600080fd5b61085f84610737565b925061086d60208501610737565b915061087b60408501610737565b90509250925092565b6000806040838503121561089757600080fd5b823591506108a760208401610737565b90509250929050565b6000602082840312156108c257600080fd5b8151801515811461081857600080fd5b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b60808152600086516060608084015280518060e085015260005b81811015610953576020818401810151610100878401015201610935565b5060008482016101000152602089015173ffffffffffffffffffffffffffffffffffffffff811660a0860152601f9091017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168401915050604088015160c084015273ffffffffffffffffffffffffffffffffffffffff871660208401528560408401526101008382030160608401526109f2610100820185876108d2565b9897505050505050505056fea26469706673582212206ea13208f1b08858bb63739a59e2bb25eb98e88ea5d76edf6c7c11c11e10a3e964736f6c634300081a003360c060405234801561001057600080fd5b5060405161207238038061207283398101604081905261002f916101f0565b6001600160a01b038216158061004c57506001600160a01b038116155b1561006a5760405163d92e233d60e01b815260040160405180910390fd5b60066100768982610342565b5060076100838882610342565b506008805460ff191660ff881617905560808590528360028111156100aa576100aa610400565b60a08160028111156100be576100be610400565b905250600192909255600080546001600160a01b039283166001600160a01b0319909116179055600880549190921661010002610100600160a81b0319909116179055506104169350505050565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261013357600080fd5b81516001600160401b0381111561014c5761014c61010c565b604051601f8201601f19908116603f011681016001600160401b038111828210171561017a5761017a61010c565b60405281815283820160200185101561019257600080fd5b60005b828110156101b157602081860181015183830182015201610195565b506000918101602001919091529392505050565b8051600381106101d457600080fd5b919050565b80516001600160a01b03811681146101d457600080fd5b600080600080600080600080610100898b03121561020d57600080fd5b88516001600160401b0381111561022357600080fd5b61022f8b828c01610122565b60208b015190995090506001600160401b0381111561024d57600080fd5b6102598b828c01610122565b975050604089015160ff8116811461027057600080fd5b60608a0151909650945061028660808a016101c5565b60a08a0151909450925061029c60c08a016101d9565b91506102aa60e08a016101d9565b90509295985092959890939650565b600181811c908216806102cd57607f821691505b6020821081036102ed57634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111561033d57806000526020600020601f840160051c8101602085101561031a5750805b601f840160051c820191505b8181101561033a5760008155600101610326565b50505b505050565b81516001600160401b0381111561035b5761035b61010c565b61036f8161036984546102b9565b846102f3565b6020601f8211600181146103a3576000831561038b5750848201515b600019600385901b1c1916600184901b17845561033a565b600084815260208120601f198516915b828110156103d357878501518255602094850194600190920191016103b3565b50848210156103f15786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b634e487b7160e01b600052602160045260246000fd5b60805160a051611c1b61045760003960006103440152600081816102f001528181610bdc01528181610ce201528181610efe01526110040152611c1b6000f3fe608060405234801561001057600080fd5b50600436106101b95760003560e01c806395d89b41116100f9578063ccc7759911610097578063eddeb12311610071578063eddeb12314610461578063f2441b3214610474578063f687d12a14610494578063fc5fecd5146104a757600080fd5b8063ccc77599146103d4578063d9eeebed146103e7578063dd62ed3e1461041b57600080fd5b8063b84c8246116100d3578063b84c824614610386578063c47f00271461039b578063c7012626146103ae578063c835d7cc146103c157600080fd5b806395d89b4114610337578063a3413d031461033f578063a9059cbb1461037357600080fd5b80633ce4a5bc116101665780634d8943bb116101405780634d8943bb146102ac57806370a08231146102b557806385e1f4d0146102eb5780638b851b951461031257600080fd5b80633ce4a5bc1461024657806342966c681461028657806347e7ef241461029957600080fd5b806318160ddd1161019757806318160ddd1461021657806323b872dd1461021e578063313ce5671461023157600080fd5b806306fdde03146101be578063091d2788146101dc578063095ea7b3146101f3575b600080fd5b6101c66104ba565b6040516101d39190611648565b60405180910390f35b6101e560015481565b6040519081526020016101d3565b610206610201366004611687565b61054c565b60405190151581526020016101d3565b6005546101e5565b61020661022c3660046116b3565b610563565b60085460405160ff90911681526020016101d3565b61026173735b14bb79463307aacbed86daf3322b1e6226ab81565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101d3565b6102066102943660046116f4565b6105fa565b6102066102a7366004611687565b61060e565b6101e560025481565b6101e56102c336600461170d565b73ffffffffffffffffffffffffffffffffffffffff1660009081526003602052604090205490565b6101e57f000000000000000000000000000000000000000000000000000000000000000081565b60085461026190610100900473ffffffffffffffffffffffffffffffffffffffff1681565b6101c6610767565b6103667f000000000000000000000000000000000000000000000000000000000000000081565b6040516101d3919061172a565b610206610381366004611687565b610776565b610399610394366004611832565b610783565b005b6103996103a9366004611832565b6107e0565b6102066103bc366004611883565b610839565b6103996103cf36600461170d565b610988565b6103996103e236600461170d565b610a9c565b6103ef610bb0565b6040805173ffffffffffffffffffffffffffffffffffffffff90931683526020830191909152016101d3565b6101e56104293660046118dc565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260046020908152604080832093909416825291909152205490565b61039961046f3660046116f4565b610dce565b6000546102619073ffffffffffffffffffffffffffffffffffffffff1681565b6103996104a23660046116f4565b610e50565b6103ef6104b53660046116f4565b610ed2565b6060600680546104c990611915565b80601f01602080910402602001604051908101604052809291908181526020018280546104f590611915565b80156105425780601f1061051757610100808354040283529160200191610542565b820191906000526020600020905b81548152906001019060200180831161052557829003601f168201915b5050505050905090565b60006105593384846110ee565b5060015b92915050565b60006105708484846111f7565b73ffffffffffffffffffffffffffffffffffffffff84166000908152600460209081526040808320338452909152902054828110156105db576040517f10bad14700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6105ef85336105ea8685611997565b6110ee565b506001949350505050565b600061060633836113b2565b506001919050565b60003373735b14bb79463307aacbed86daf3322b1e6226ab1480159061064c575060005473ffffffffffffffffffffffffffffffffffffffff163314155b80156106755750600854610100900473ffffffffffffffffffffffffffffffffffffffff163314155b156106ac576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6106b683836114f4565b6040517f735b14bb79463307aacbed86daf3322b1e6226ab000000000000000000000000602082015273ffffffffffffffffffffffffffffffffffffffff8416907f67fc7bdaed5b0ec550d8706b87d60568ab70c6b781263c70101d54cd1564aab390603401604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152908290526107569186906119aa565b60405180910390a250600192915050565b6060600780546104c990611915565b60006105593384846111f7565b3373735b14bb79463307aacbed86daf3322b1e6226ab146107d0576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60076107dc8282611a1b565b5050565b3373735b14bb79463307aacbed86daf3322b1e6226ab1461082d576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60066107dc8282611a1b565b6000806000610846610bb0565b6040517f23b872dd00000000000000000000000000000000000000000000000000000000815233600482015273735b14bb79463307aacbed86daf3322b1e6226ab602482015260448101829052919350915073ffffffffffffffffffffffffffffffffffffffff8316906323b872dd906064016020604051808303816000875af11580156108d8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108fc9190611b34565b610932576040517f0a7cd6d600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61093c33856113b2565b60025460405133917f9ffbffc04a397460ee1dbe8c9503e098090567d6b7f4b3c02a8617d800b6d9559161097591899189918791611b56565b60405180910390a2506001949350505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab146109d5576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116610a22576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fd55614e962c5fd6ece71614f6348d702468a997a394dd5e5c1677950226d97ae906020015b60405180910390a150565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610ae9576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116610b36576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600880547fffffffffffffffffffffff0000000000000000000000000000000000000000ff1661010073ffffffffffffffffffffffffffffffffffffffff8416908102919091179091556040519081527f88815d964e380677e86d817e7d65dea59cb7b4c3b5b7a0c8ec7ea4a74f90a38790602001610a91565b600080546040517f0be155470000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060048201528291829173ffffffffffffffffffffffffffffffffffffffff90911690630be1554790602401602060405180830381865afa158015610c43573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c679190611b85565b905073ffffffffffffffffffffffffffffffffffffffff8116610cb6576040517f78fff39600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080546040517fd7fd7afb0000000000000000000000000000000000000000000000000000000081527f0000000000000000000000000000000000000000000000000000000000000000600482015273ffffffffffffffffffffffffffffffffffffffff9091169063d7fd7afb90602401602060405180830381865afa158015610d45573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d699190611ba2565b905080600003610da5576040517fe661aed000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060025460015483610db89190611bbb565b610dc29190611bd2565b92959294509192505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610e1b576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60028190556040518181527fef13af88e424b5d15f49c77758542c1938b08b8b95b91ed0751f98ba99000d8f90602001610a91565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610e9d576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60018190556040518181527fff5788270f43bfc1ca41c503606d2594aa3023a1a7547de403a3e2f146a4a80a90602001610a91565b600080546040517f0be155470000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060048201528291829173ffffffffffffffffffffffffffffffffffffffff90911690630be1554790602401602060405180830381865afa158015610f65573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f899190611b85565b905073ffffffffffffffffffffffffffffffffffffffff8116610fd8576040517f78fff39600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080546040517fd7fd7afb0000000000000000000000000000000000000000000000000000000081527f0000000000000000000000000000000000000000000000000000000000000000600482015273ffffffffffffffffffffffffffffffffffffffff9091169063d7fd7afb90602401602060405180830381865afa158015611067573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061108b9190611ba2565b9050806000036110c7576040517fe661aed000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002546000906110d78784611bbb565b6110e19190611bd2565b9296929550919350505050565b73ffffffffffffffffffffffffffffffffffffffff831661113b576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216611188576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff83811660008181526004602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff8316611244576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216611291576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8316600090815260036020526040902054818110156112f1576040517ffe382aa700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6112fb8282611997565b73ffffffffffffffffffffffffffffffffffffffff808616600090815260036020526040808220939093559085168152908120805484929061133e908490611bd2565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516113a491815260200190565b60405180910390a350505050565b73ffffffffffffffffffffffffffffffffffffffff82166113ff576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff82166000908152600360205260409020548181101561145f576040517ffe382aa700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6114698282611997565b73ffffffffffffffffffffffffffffffffffffffff8416600090815260036020526040812091909155600580548492906114a4908490611997565b909155505060405182815260009073ffffffffffffffffffffffffffffffffffffffff8516907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef906020016111ea565b73ffffffffffffffffffffffffffffffffffffffff8216611541576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80600560008282546115539190611bd2565b909155505073ffffffffffffffffffffffffffffffffffffffff82166000908152600360205260408120805483929061158d908490611bd2565b909155505060405181815273ffffffffffffffffffffffffffffffffffffffff8316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b6000815180845260005b8181101561160a576020818501810151868301820152016115ee565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b60208152600061165b60208301846115e4565b9392505050565b73ffffffffffffffffffffffffffffffffffffffff8116811461168457600080fd5b50565b6000806040838503121561169a57600080fd5b82356116a581611662565b946020939093013593505050565b6000806000606084860312156116c857600080fd5b83356116d381611662565b925060208401356116e381611662565b929592945050506040919091013590565b60006020828403121561170657600080fd5b5035919050565b60006020828403121561171f57600080fd5b813561165b81611662565b6020810160038310611765577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b91905290565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60008067ffffffffffffffff8411156117b5576117b561176b565b506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f85018116603f0116810181811067ffffffffffffffff821117156118025761180261176b565b60405283815290508082840185101561181a57600080fd5b83836020830137600060208583010152509392505050565b60006020828403121561184457600080fd5b813567ffffffffffffffff81111561185b57600080fd5b8201601f8101841361186c57600080fd5b61187b8482356020840161179a565b949350505050565b6000806040838503121561189657600080fd5b823567ffffffffffffffff8111156118ad57600080fd5b8301601f810185136118be57600080fd5b6118cd8582356020840161179a565b95602094909401359450505050565b600080604083850312156118ef57600080fd5b82356118fa81611662565b9150602083013561190a81611662565b809150509250929050565b600181811c9082168061192957607f821691505b602082108103611962577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8181038181111561055d5761055d611968565b6040815260006119bd60408301856115e4565b90508260208301529392505050565b601f821115611a1657806000526020600020601f840160051c810160208510156119f35750805b601f840160051c820191505b81811015611a1357600081556001016119ff565b50505b505050565b815167ffffffffffffffff811115611a3557611a3561176b565b611a4981611a438454611915565b846119cc565b6020601f821160018114611a9b5760008315611a655750848201515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b178455611a13565b6000848152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08516915b82811015611ae95787850151825560209485019460019092019101611ac9565b5084821015611b2557868401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b60f8161c191681555b50505050600190811b01905550565b600060208284031215611b4657600080fd5b8151801515811461165b57600080fd5b608081526000611b6960808301876115e4565b6020830195909552506040810192909252606090910152919050565b600060208284031215611b9757600080fd5b815161165b81611662565b600060208284031215611bb457600080fd5b5051919050565b808202811582820484141761055d5761055d611968565b8082018082111561055d5761055d61196856fea26469706673582212208307d60e253f5034856b93df00e3e2f46b06f9765d906dbd93ee947935fc608764736f6c634300081a0033a26469706673582212207f656d2f4aeb4c5f3bb773139f4068741a9374dcea586b24785cefd9a7bd837364736f6c634300081a0033",
}

// GatewayEVMZEVMTestABI is the input ABI used to generate the binding from.
// Deprecated: Use GatewayEVMZEVMTestMetaData.ABI instead.
var GatewayEVMZEVMTestABI = GatewayEVMZEVMTestMetaData.ABI

// GatewayEVMZEVMTestBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GatewayEVMZEVMTestMetaData.Bin instead.
var GatewayEVMZEVMTestBin = GatewayEVMZEVMTestMetaData.Bin

// DeployGatewayEVMZEVMTest deploys a new Ethereum contract, binding an instance of GatewayEVMZEVMTest to it.
func DeployGatewayEVMZEVMTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GatewayEVMZEVMTest, error) {
	parsed, err := GatewayEVMZEVMTestMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GatewayEVMZEVMTestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GatewayEVMZEVMTest{GatewayEVMZEVMTestCaller: GatewayEVMZEVMTestCaller{contract: contract}, GatewayEVMZEVMTestTransactor: GatewayEVMZEVMTestTransactor{contract: contract}, GatewayEVMZEVMTestFilterer: GatewayEVMZEVMTestFilterer{contract: contract}}, nil
}

// GatewayEVMZEVMTest is an auto generated Go binding around an Ethereum contract.
type GatewayEVMZEVMTest struct {
	GatewayEVMZEVMTestCaller     // Read-only binding to the contract
	GatewayEVMZEVMTestTransactor // Write-only binding to the contract
	GatewayEVMZEVMTestFilterer   // Log filterer for contract events
}

// GatewayEVMZEVMTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type GatewayEVMZEVMTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayEVMZEVMTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GatewayEVMZEVMTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayEVMZEVMTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GatewayEVMZEVMTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayEVMZEVMTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GatewayEVMZEVMTestSession struct {
	Contract     *GatewayEVMZEVMTest // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// GatewayEVMZEVMTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GatewayEVMZEVMTestCallerSession struct {
	Contract *GatewayEVMZEVMTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// GatewayEVMZEVMTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GatewayEVMZEVMTestTransactorSession struct {
	Contract     *GatewayEVMZEVMTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// GatewayEVMZEVMTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type GatewayEVMZEVMTestRaw struct {
	Contract *GatewayEVMZEVMTest // Generic contract binding to access the raw methods on
}

// GatewayEVMZEVMTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GatewayEVMZEVMTestCallerRaw struct {
	Contract *GatewayEVMZEVMTestCaller // Generic read-only contract binding to access the raw methods on
}

// GatewayEVMZEVMTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GatewayEVMZEVMTestTransactorRaw struct {
	Contract *GatewayEVMZEVMTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGatewayEVMZEVMTest creates a new instance of GatewayEVMZEVMTest, bound to a specific deployed contract.
func NewGatewayEVMZEVMTest(address common.Address, backend bind.ContractBackend) (*GatewayEVMZEVMTest, error) {
	contract, err := bindGatewayEVMZEVMTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTest{GatewayEVMZEVMTestCaller: GatewayEVMZEVMTestCaller{contract: contract}, GatewayEVMZEVMTestTransactor: GatewayEVMZEVMTestTransactor{contract: contract}, GatewayEVMZEVMTestFilterer: GatewayEVMZEVMTestFilterer{contract: contract}}, nil
}

// NewGatewayEVMZEVMTestCaller creates a new read-only instance of GatewayEVMZEVMTest, bound to a specific deployed contract.
func NewGatewayEVMZEVMTestCaller(address common.Address, caller bind.ContractCaller) (*GatewayEVMZEVMTestCaller, error) {
	contract, err := bindGatewayEVMZEVMTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestCaller{contract: contract}, nil
}

// NewGatewayEVMZEVMTestTransactor creates a new write-only instance of GatewayEVMZEVMTest, bound to a specific deployed contract.
func NewGatewayEVMZEVMTestTransactor(address common.Address, transactor bind.ContractTransactor) (*GatewayEVMZEVMTestTransactor, error) {
	contract, err := bindGatewayEVMZEVMTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestTransactor{contract: contract}, nil
}

// NewGatewayEVMZEVMTestFilterer creates a new log filterer instance of GatewayEVMZEVMTest, bound to a specific deployed contract.
func NewGatewayEVMZEVMTestFilterer(address common.Address, filterer bind.ContractFilterer) (*GatewayEVMZEVMTestFilterer, error) {
	contract, err := bindGatewayEVMZEVMTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestFilterer{contract: contract}, nil
}

// bindGatewayEVMZEVMTest binds a generic wrapper to an already deployed contract.
func bindGatewayEVMZEVMTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GatewayEVMZEVMTestMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayEVMZEVMTest.Contract.GatewayEVMZEVMTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMZEVMTest.Contract.GatewayEVMZEVMTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayEVMZEVMTest.Contract.GatewayEVMZEVMTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayEVMZEVMTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMZEVMTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayEVMZEVMTest.Contract.contract.Transact(opts, method, params...)
}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestCaller) ISTEST(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GatewayEVMZEVMTest.contract.Call(opts, &out, "IS_TEST")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestSession) ISTEST() (bool, error) {
	return _GatewayEVMZEVMTest.Contract.ISTEST(&_GatewayEVMZEVMTest.CallOpts)
}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestCallerSession) ISTEST() (bool, error) {
	return _GatewayEVMZEVMTest.Contract.ISTEST(&_GatewayEVMZEVMTest.CallOpts)
}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestCaller) ExcludeArtifacts(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _GatewayEVMZEVMTest.contract.Call(opts, &out, "excludeArtifacts")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestSession) ExcludeArtifacts() ([]string, error) {
	return _GatewayEVMZEVMTest.Contract.ExcludeArtifacts(&_GatewayEVMZEVMTest.CallOpts)
}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestCallerSession) ExcludeArtifacts() ([]string, error) {
	return _GatewayEVMZEVMTest.Contract.ExcludeArtifacts(&_GatewayEVMZEVMTest.CallOpts)
}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestCaller) ExcludeContracts(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GatewayEVMZEVMTest.contract.Call(opts, &out, "excludeContracts")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestSession) ExcludeContracts() ([]common.Address, error) {
	return _GatewayEVMZEVMTest.Contract.ExcludeContracts(&_GatewayEVMZEVMTest.CallOpts)
}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestCallerSession) ExcludeContracts() ([]common.Address, error) {
	return _GatewayEVMZEVMTest.Contract.ExcludeContracts(&_GatewayEVMZEVMTest.CallOpts)
}

// ExcludeSelectors is a free data retrieval call binding the contract method 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestCaller) ExcludeSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzSelector, error) {
	var out []interface{}
	err := _GatewayEVMZEVMTest.contract.Call(opts, &out, "excludeSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzSelector)).(*[]StdInvariantFuzzSelector)

	return out0, err

}

// ExcludeSelectors is a free data retrieval call binding the contract method 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestSession) ExcludeSelectors() ([]StdInvariantFuzzSelector, error) {
	return _GatewayEVMZEVMTest.Contract.ExcludeSelectors(&_GatewayEVMZEVMTest.CallOpts)
}

// ExcludeSelectors is a free data retrieval call binding the contract method 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestCallerSession) ExcludeSelectors() ([]StdInvariantFuzzSelector, error) {
	return _GatewayEVMZEVMTest.Contract.ExcludeSelectors(&_GatewayEVMZEVMTest.CallOpts)
}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestCaller) ExcludeSenders(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GatewayEVMZEVMTest.contract.Call(opts, &out, "excludeSenders")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestSession) ExcludeSenders() ([]common.Address, error) {
	return _GatewayEVMZEVMTest.Contract.ExcludeSenders(&_GatewayEVMZEVMTest.CallOpts)
}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestCallerSession) ExcludeSenders() ([]common.Address, error) {
	return _GatewayEVMZEVMTest.Contract.ExcludeSenders(&_GatewayEVMZEVMTest.CallOpts)
}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestCaller) Failed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GatewayEVMZEVMTest.contract.Call(opts, &out, "failed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestSession) Failed() (bool, error) {
	return _GatewayEVMZEVMTest.Contract.Failed(&_GatewayEVMZEVMTest.CallOpts)
}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestCallerSession) Failed() (bool, error) {
	return _GatewayEVMZEVMTest.Contract.Failed(&_GatewayEVMZEVMTest.CallOpts)
}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestCaller) TargetArtifactSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzArtifactSelector, error) {
	var out []interface{}
	err := _GatewayEVMZEVMTest.contract.Call(opts, &out, "targetArtifactSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzArtifactSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzArtifactSelector)).(*[]StdInvariantFuzzArtifactSelector)

	return out0, err

}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestSession) TargetArtifactSelectors() ([]StdInvariantFuzzArtifactSelector, error) {
	return _GatewayEVMZEVMTest.Contract.TargetArtifactSelectors(&_GatewayEVMZEVMTest.CallOpts)
}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestCallerSession) TargetArtifactSelectors() ([]StdInvariantFuzzArtifactSelector, error) {
	return _GatewayEVMZEVMTest.Contract.TargetArtifactSelectors(&_GatewayEVMZEVMTest.CallOpts)
}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestCaller) TargetArtifacts(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _GatewayEVMZEVMTest.contract.Call(opts, &out, "targetArtifacts")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestSession) TargetArtifacts() ([]string, error) {
	return _GatewayEVMZEVMTest.Contract.TargetArtifacts(&_GatewayEVMZEVMTest.CallOpts)
}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestCallerSession) TargetArtifacts() ([]string, error) {
	return _GatewayEVMZEVMTest.Contract.TargetArtifacts(&_GatewayEVMZEVMTest.CallOpts)
}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestCaller) TargetContracts(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GatewayEVMZEVMTest.contract.Call(opts, &out, "targetContracts")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestSession) TargetContracts() ([]common.Address, error) {
	return _GatewayEVMZEVMTest.Contract.TargetContracts(&_GatewayEVMZEVMTest.CallOpts)
}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestCallerSession) TargetContracts() ([]common.Address, error) {
	return _GatewayEVMZEVMTest.Contract.TargetContracts(&_GatewayEVMZEVMTest.CallOpts)
}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestCaller) TargetInterfaces(opts *bind.CallOpts) ([]StdInvariantFuzzInterface, error) {
	var out []interface{}
	err := _GatewayEVMZEVMTest.contract.Call(opts, &out, "targetInterfaces")

	if err != nil {
		return *new([]StdInvariantFuzzInterface), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzInterface)).(*[]StdInvariantFuzzInterface)

	return out0, err

}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestSession) TargetInterfaces() ([]StdInvariantFuzzInterface, error) {
	return _GatewayEVMZEVMTest.Contract.TargetInterfaces(&_GatewayEVMZEVMTest.CallOpts)
}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestCallerSession) TargetInterfaces() ([]StdInvariantFuzzInterface, error) {
	return _GatewayEVMZEVMTest.Contract.TargetInterfaces(&_GatewayEVMZEVMTest.CallOpts)
}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestCaller) TargetSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzSelector, error) {
	var out []interface{}
	err := _GatewayEVMZEVMTest.contract.Call(opts, &out, "targetSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzSelector)).(*[]StdInvariantFuzzSelector)

	return out0, err

}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestSession) TargetSelectors() ([]StdInvariantFuzzSelector, error) {
	return _GatewayEVMZEVMTest.Contract.TargetSelectors(&_GatewayEVMZEVMTest.CallOpts)
}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestCallerSession) TargetSelectors() ([]StdInvariantFuzzSelector, error) {
	return _GatewayEVMZEVMTest.Contract.TargetSelectors(&_GatewayEVMZEVMTest.CallOpts)
}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestCaller) TargetSenders(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GatewayEVMZEVMTest.contract.Call(opts, &out, "targetSenders")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestSession) TargetSenders() ([]common.Address, error) {
	return _GatewayEVMZEVMTest.Contract.TargetSenders(&_GatewayEVMZEVMTest.CallOpts)
}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestCallerSession) TargetSenders() ([]common.Address, error) {
	return _GatewayEVMZEVMTest.Contract.TargetSenders(&_GatewayEVMZEVMTest.CallOpts)
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestTransactor) SetUp(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMZEVMTest.contract.Transact(opts, "setUp")
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestSession) SetUp() (*types.Transaction, error) {
	return _GatewayEVMZEVMTest.Contract.SetUp(&_GatewayEVMZEVMTest.TransactOpts)
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestTransactorSession) SetUp() (*types.Transaction, error) {
	return _GatewayEVMZEVMTest.Contract.SetUp(&_GatewayEVMZEVMTest.TransactOpts)
}

// TestCallReceiverEVMFromSenderZEVM is a paid mutator transaction binding the contract method 0xd7a525fc.
//
// Solidity: function testCallReceiverEVMFromSenderZEVM() returns()
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestTransactor) TestCallReceiverEVMFromSenderZEVM(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMZEVMTest.contract.Transact(opts, "testCallReceiverEVMFromSenderZEVM")
}

// TestCallReceiverEVMFromSenderZEVM is a paid mutator transaction binding the contract method 0xd7a525fc.
//
// Solidity: function testCallReceiverEVMFromSenderZEVM() returns()
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestSession) TestCallReceiverEVMFromSenderZEVM() (*types.Transaction, error) {
	return _GatewayEVMZEVMTest.Contract.TestCallReceiverEVMFromSenderZEVM(&_GatewayEVMZEVMTest.TransactOpts)
}

// TestCallReceiverEVMFromSenderZEVM is a paid mutator transaction binding the contract method 0xd7a525fc.
//
// Solidity: function testCallReceiverEVMFromSenderZEVM() returns()
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestTransactorSession) TestCallReceiverEVMFromSenderZEVM() (*types.Transaction, error) {
	return _GatewayEVMZEVMTest.Contract.TestCallReceiverEVMFromSenderZEVM(&_GatewayEVMZEVMTest.TransactOpts)
}

// TestCallReceiverEVMFromZEVM is a paid mutator transaction binding the contract method 0x9683c695.
//
// Solidity: function testCallReceiverEVMFromZEVM() returns()
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestTransactor) TestCallReceiverEVMFromZEVM(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMZEVMTest.contract.Transact(opts, "testCallReceiverEVMFromZEVM")
}

// TestCallReceiverEVMFromZEVM is a paid mutator transaction binding the contract method 0x9683c695.
//
// Solidity: function testCallReceiverEVMFromZEVM() returns()
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestSession) TestCallReceiverEVMFromZEVM() (*types.Transaction, error) {
	return _GatewayEVMZEVMTest.Contract.TestCallReceiverEVMFromZEVM(&_GatewayEVMZEVMTest.TransactOpts)
}

// TestCallReceiverEVMFromZEVM is a paid mutator transaction binding the contract method 0x9683c695.
//
// Solidity: function testCallReceiverEVMFromZEVM() returns()
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestTransactorSession) TestCallReceiverEVMFromZEVM() (*types.Transaction, error) {
	return _GatewayEVMZEVMTest.Contract.TestCallReceiverEVMFromZEVM(&_GatewayEVMZEVMTest.TransactOpts)
}

// TestWithdrawAndCallReceiverEVMFromSenderZEVM is a paid mutator transaction binding the contract method 0x52474413.
//
// Solidity: function testWithdrawAndCallReceiverEVMFromSenderZEVM() returns()
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestTransactor) TestWithdrawAndCallReceiverEVMFromSenderZEVM(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMZEVMTest.contract.Transact(opts, "testWithdrawAndCallReceiverEVMFromSenderZEVM")
}

// TestWithdrawAndCallReceiverEVMFromSenderZEVM is a paid mutator transaction binding the contract method 0x52474413.
//
// Solidity: function testWithdrawAndCallReceiverEVMFromSenderZEVM() returns()
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestSession) TestWithdrawAndCallReceiverEVMFromSenderZEVM() (*types.Transaction, error) {
	return _GatewayEVMZEVMTest.Contract.TestWithdrawAndCallReceiverEVMFromSenderZEVM(&_GatewayEVMZEVMTest.TransactOpts)
}

// TestWithdrawAndCallReceiverEVMFromSenderZEVM is a paid mutator transaction binding the contract method 0x52474413.
//
// Solidity: function testWithdrawAndCallReceiverEVMFromSenderZEVM() returns()
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestTransactorSession) TestWithdrawAndCallReceiverEVMFromSenderZEVM() (*types.Transaction, error) {
	return _GatewayEVMZEVMTest.Contract.TestWithdrawAndCallReceiverEVMFromSenderZEVM(&_GatewayEVMZEVMTest.TransactOpts)
}

// TestWithdrawAndCallReceiverEVMFromZEVM is a paid mutator transaction binding the contract method 0x6ff15ccc.
//
// Solidity: function testWithdrawAndCallReceiverEVMFromZEVM() returns()
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestTransactor) TestWithdrawAndCallReceiverEVMFromZEVM(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayEVMZEVMTest.contract.Transact(opts, "testWithdrawAndCallReceiverEVMFromZEVM")
}

// TestWithdrawAndCallReceiverEVMFromZEVM is a paid mutator transaction binding the contract method 0x6ff15ccc.
//
// Solidity: function testWithdrawAndCallReceiverEVMFromZEVM() returns()
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestSession) TestWithdrawAndCallReceiverEVMFromZEVM() (*types.Transaction, error) {
	return _GatewayEVMZEVMTest.Contract.TestWithdrawAndCallReceiverEVMFromZEVM(&_GatewayEVMZEVMTest.TransactOpts)
}

// TestWithdrawAndCallReceiverEVMFromZEVM is a paid mutator transaction binding the contract method 0x6ff15ccc.
//
// Solidity: function testWithdrawAndCallReceiverEVMFromZEVM() returns()
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestTransactorSession) TestWithdrawAndCallReceiverEVMFromZEVM() (*types.Transaction, error) {
	return _GatewayEVMZEVMTest.Contract.TestWithdrawAndCallReceiverEVMFromZEVM(&_GatewayEVMZEVMTest.TransactOpts)
}

// GatewayEVMZEVMTestCalledIterator is returned from FilterCalled and is used to iterate over the raw logs and unpacked data for Called events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestCalledIterator struct {
	Event *GatewayEVMZEVMTestCalled // Event containing the contract specifics and raw log

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
func (it *GatewayEVMZEVMTestCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestCalled)
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
		it.Event = new(GatewayEVMZEVMTestCalled)
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
func (it *GatewayEVMZEVMTestCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestCalled represents a Called event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestCalled struct {
	Sender        common.Address
	Receiver      common.Address
	Payload       []byte
	RevertOptions RevertOptions
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCalled is a free log retrieval operation binding the contract event 0xd34634f30f94a646fdf4ce7078f38fc5fa0d3f0b193658facea4e3e43330d974.
//
// Solidity: event Called(address indexed sender, address indexed receiver, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterCalled(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*GatewayEVMZEVMTestCalledIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "Called", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestCalledIterator{contract: _GatewayEVMZEVMTest.contract, event: "Called", logs: logs, sub: sub}, nil
}

// WatchCalled is a free log subscription operation binding the contract event 0xd34634f30f94a646fdf4ce7078f38fc5fa0d3f0b193658facea4e3e43330d974.
//
// Solidity: event Called(address indexed sender, address indexed receiver, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchCalled(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestCalled, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "Called", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestCalled)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "Called", log); err != nil {
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

// ParseCalled is a log parse operation binding the contract event 0xd34634f30f94a646fdf4ce7078f38fc5fa0d3f0b193658facea4e3e43330d974.
//
// Solidity: event Called(address indexed sender, address indexed receiver, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseCalled(log types.Log) (*GatewayEVMZEVMTestCalled, error) {
	event := new(GatewayEVMZEVMTestCalled)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "Called", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestCalled0Iterator is returned from FilterCalled0 and is used to iterate over the raw logs and unpacked data for Called0 events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestCalled0Iterator struct {
	Event *GatewayEVMZEVMTestCalled0 // Event containing the contract specifics and raw log

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
func (it *GatewayEVMZEVMTestCalled0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestCalled0)
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
		it.Event = new(GatewayEVMZEVMTestCalled0)
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
func (it *GatewayEVMZEVMTestCalled0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestCalled0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestCalled0 represents a Called0 event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestCalled0 struct {
	Sender        common.Address
	Zrc20         common.Address
	Receiver      []byte
	Message       []byte
	CallOptions   CallOptions
	RevertOptions RevertOptions
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCalled0 is a free log retrieval operation binding the contract event 0x306ee13f48319a123b222c69908e44dcf91abffc20cacc502e3cf5a4ff23e0e4.
//
// Solidity: event Called(address indexed sender, address indexed zrc20, bytes receiver, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterCalled0(opts *bind.FilterOpts, sender []common.Address, zrc20 []common.Address) (*GatewayEVMZEVMTestCalled0Iterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var zrc20Rule []interface{}
	for _, zrc20Item := range zrc20 {
		zrc20Rule = append(zrc20Rule, zrc20Item)
	}

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "Called0", senderRule, zrc20Rule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestCalled0Iterator{contract: _GatewayEVMZEVMTest.contract, event: "Called0", logs: logs, sub: sub}, nil
}

// WatchCalled0 is a free log subscription operation binding the contract event 0x306ee13f48319a123b222c69908e44dcf91abffc20cacc502e3cf5a4ff23e0e4.
//
// Solidity: event Called(address indexed sender, address indexed zrc20, bytes receiver, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchCalled0(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestCalled0, sender []common.Address, zrc20 []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var zrc20Rule []interface{}
	for _, zrc20Item := range zrc20 {
		zrc20Rule = append(zrc20Rule, zrc20Item)
	}

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "Called0", senderRule, zrc20Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestCalled0)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "Called0", log); err != nil {
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

// ParseCalled0 is a log parse operation binding the contract event 0x306ee13f48319a123b222c69908e44dcf91abffc20cacc502e3cf5a4ff23e0e4.
//
// Solidity: event Called(address indexed sender, address indexed zrc20, bytes receiver, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseCalled0(log types.Log) (*GatewayEVMZEVMTestCalled0, error) {
	event := new(GatewayEVMZEVMTestCalled0)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "Called0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestDepositedIterator struct {
	Event *GatewayEVMZEVMTestDeposited // Event containing the contract specifics and raw log

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
func (it *GatewayEVMZEVMTestDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestDeposited)
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
		it.Event = new(GatewayEVMZEVMTestDeposited)
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
func (it *GatewayEVMZEVMTestDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestDeposited represents a Deposited event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestDeposited struct {
	Sender        common.Address
	Receiver      common.Address
	Amount        *big.Int
	Asset         common.Address
	Payload       []byte
	RevertOptions RevertOptions
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0xc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c.
//
// Solidity: event Deposited(address indexed sender, address indexed receiver, uint256 amount, address asset, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterDeposited(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*GatewayEVMZEVMTestDepositedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "Deposited", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestDepositedIterator{contract: _GatewayEVMZEVMTest.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0xc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c.
//
// Solidity: event Deposited(address indexed sender, address indexed receiver, uint256 amount, address asset, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestDeposited, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "Deposited", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestDeposited)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0xc6f891b65320c682b217616a62b51f218fee95d5f0ba83e758ef9ab4ee8e975c.
//
// Solidity: event Deposited(address indexed sender, address indexed receiver, uint256 amount, address asset, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseDeposited(log types.Log) (*GatewayEVMZEVMTestDeposited, error) {
	event := new(GatewayEVMZEVMTestDeposited)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestDepositedAndCalledIterator is returned from FilterDepositedAndCalled and is used to iterate over the raw logs and unpacked data for DepositedAndCalled events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestDepositedAndCalledIterator struct {
	Event *GatewayEVMZEVMTestDepositedAndCalled // Event containing the contract specifics and raw log

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
func (it *GatewayEVMZEVMTestDepositedAndCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestDepositedAndCalled)
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
		it.Event = new(GatewayEVMZEVMTestDepositedAndCalled)
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
func (it *GatewayEVMZEVMTestDepositedAndCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestDepositedAndCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestDepositedAndCalled represents a DepositedAndCalled event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestDepositedAndCalled struct {
	Sender        common.Address
	Receiver      common.Address
	Amount        *big.Int
	Asset         common.Address
	Payload       []byte
	RevertOptions RevertOptions
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDepositedAndCalled is a free log retrieval operation binding the contract event 0xa795d4377323e4c2d4c346b8050a7dd504c4043be8884c81b8d9690706c8388f.
//
// Solidity: event DepositedAndCalled(address indexed sender, address indexed receiver, uint256 amount, address asset, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterDepositedAndCalled(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*GatewayEVMZEVMTestDepositedAndCalledIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "DepositedAndCalled", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestDepositedAndCalledIterator{contract: _GatewayEVMZEVMTest.contract, event: "DepositedAndCalled", logs: logs, sub: sub}, nil
}

// WatchDepositedAndCalled is a free log subscription operation binding the contract event 0xa795d4377323e4c2d4c346b8050a7dd504c4043be8884c81b8d9690706c8388f.
//
// Solidity: event DepositedAndCalled(address indexed sender, address indexed receiver, uint256 amount, address asset, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchDepositedAndCalled(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestDepositedAndCalled, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "DepositedAndCalled", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestDepositedAndCalled)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "DepositedAndCalled", log); err != nil {
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

// ParseDepositedAndCalled is a log parse operation binding the contract event 0xa795d4377323e4c2d4c346b8050a7dd504c4043be8884c81b8d9690706c8388f.
//
// Solidity: event DepositedAndCalled(address indexed sender, address indexed receiver, uint256 amount, address asset, bytes payload, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseDepositedAndCalled(log types.Log) (*GatewayEVMZEVMTestDepositedAndCalled, error) {
	event := new(GatewayEVMZEVMTestDepositedAndCalled)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "DepositedAndCalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestExecutedIterator is returned from FilterExecuted and is used to iterate over the raw logs and unpacked data for Executed events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestExecutedIterator struct {
	Event *GatewayEVMZEVMTestExecuted // Event containing the contract specifics and raw log

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
func (it *GatewayEVMZEVMTestExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestExecuted)
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
		it.Event = new(GatewayEVMZEVMTestExecuted)
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
func (it *GatewayEVMZEVMTestExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestExecuted represents a Executed event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestExecuted struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExecuted is a free log retrieval operation binding the contract event 0xcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f.
//
// Solidity: event Executed(address indexed destination, uint256 value, bytes data)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterExecuted(opts *bind.FilterOpts, destination []common.Address) (*GatewayEVMZEVMTestExecutedIterator, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "Executed", destinationRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestExecutedIterator{contract: _GatewayEVMZEVMTest.contract, event: "Executed", logs: logs, sub: sub}, nil
}

// WatchExecuted is a free log subscription operation binding the contract event 0xcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f.
//
// Solidity: event Executed(address indexed destination, uint256 value, bytes data)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchExecuted(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestExecuted, destination []common.Address) (event.Subscription, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "Executed", destinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestExecuted)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "Executed", log); err != nil {
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

// ParseExecuted is a log parse operation binding the contract event 0xcaf938de11c367272220bfd1d2baa99ca46665e7bc4d85f00adb51b90fe1fa9f.
//
// Solidity: event Executed(address indexed destination, uint256 value, bytes data)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseExecuted(log types.Log) (*GatewayEVMZEVMTestExecuted, error) {
	event := new(GatewayEVMZEVMTestExecuted)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "Executed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestExecutedWithERC20Iterator is returned from FilterExecutedWithERC20 and is used to iterate over the raw logs and unpacked data for ExecutedWithERC20 events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestExecutedWithERC20Iterator struct {
	Event *GatewayEVMZEVMTestExecutedWithERC20 // Event containing the contract specifics and raw log

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
func (it *GatewayEVMZEVMTestExecutedWithERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestExecutedWithERC20)
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
		it.Event = new(GatewayEVMZEVMTestExecutedWithERC20)
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
func (it *GatewayEVMZEVMTestExecutedWithERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestExecutedWithERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestExecutedWithERC20 represents a ExecutedWithERC20 event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestExecutedWithERC20 struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExecutedWithERC20 is a free log retrieval operation binding the contract event 0x29c40793bffd84cb810179f15d1ceec72bc7f0785514c668ba36645cf99b7382.
//
// Solidity: event ExecutedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterExecutedWithERC20(opts *bind.FilterOpts, token []common.Address, to []common.Address) (*GatewayEVMZEVMTestExecutedWithERC20Iterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "ExecutedWithERC20", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestExecutedWithERC20Iterator{contract: _GatewayEVMZEVMTest.contract, event: "ExecutedWithERC20", logs: logs, sub: sub}, nil
}

// WatchExecutedWithERC20 is a free log subscription operation binding the contract event 0x29c40793bffd84cb810179f15d1ceec72bc7f0785514c668ba36645cf99b7382.
//
// Solidity: event ExecutedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchExecutedWithERC20(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestExecutedWithERC20, token []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "ExecutedWithERC20", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestExecutedWithERC20)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "ExecutedWithERC20", log); err != nil {
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

// ParseExecutedWithERC20 is a log parse operation binding the contract event 0x29c40793bffd84cb810179f15d1ceec72bc7f0785514c668ba36645cf99b7382.
//
// Solidity: event ExecutedWithERC20(address indexed token, address indexed to, uint256 amount, bytes data)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseExecutedWithERC20(log types.Log) (*GatewayEVMZEVMTestExecutedWithERC20, error) {
	event := new(GatewayEVMZEVMTestExecutedWithERC20)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "ExecutedWithERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestReceivedERC20Iterator is returned from FilterReceivedERC20 and is used to iterate over the raw logs and unpacked data for ReceivedERC20 events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestReceivedERC20Iterator struct {
	Event *GatewayEVMZEVMTestReceivedERC20 // Event containing the contract specifics and raw log

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
func (it *GatewayEVMZEVMTestReceivedERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestReceivedERC20)
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
		it.Event = new(GatewayEVMZEVMTestReceivedERC20)
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
func (it *GatewayEVMZEVMTestReceivedERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestReceivedERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestReceivedERC20 represents a ReceivedERC20 event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestReceivedERC20 struct {
	Sender      common.Address
	Amount      *big.Int
	Token       common.Address
	Destination common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterReceivedERC20 is a free log retrieval operation binding the contract event 0x2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af60.
//
// Solidity: event ReceivedERC20(address sender, uint256 amount, address token, address destination)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterReceivedERC20(opts *bind.FilterOpts) (*GatewayEVMZEVMTestReceivedERC20Iterator, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "ReceivedERC20")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestReceivedERC20Iterator{contract: _GatewayEVMZEVMTest.contract, event: "ReceivedERC20", logs: logs, sub: sub}, nil
}

// WatchReceivedERC20 is a free log subscription operation binding the contract event 0x2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af60.
//
// Solidity: event ReceivedERC20(address sender, uint256 amount, address token, address destination)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchReceivedERC20(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestReceivedERC20) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "ReceivedERC20")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestReceivedERC20)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "ReceivedERC20", log); err != nil {
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

// ParseReceivedERC20 is a log parse operation binding the contract event 0x2b58128f24a9f59127cc5b5430d70542b22220f2d9adaa86e442b816ab98af60.
//
// Solidity: event ReceivedERC20(address sender, uint256 amount, address token, address destination)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseReceivedERC20(log types.Log) (*GatewayEVMZEVMTestReceivedERC20, error) {
	event := new(GatewayEVMZEVMTestReceivedERC20)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "ReceivedERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestReceivedNoParamsIterator is returned from FilterReceivedNoParams and is used to iterate over the raw logs and unpacked data for ReceivedNoParams events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestReceivedNoParamsIterator struct {
	Event *GatewayEVMZEVMTestReceivedNoParams // Event containing the contract specifics and raw log

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
func (it *GatewayEVMZEVMTestReceivedNoParamsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestReceivedNoParams)
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
		it.Event = new(GatewayEVMZEVMTestReceivedNoParams)
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
func (it *GatewayEVMZEVMTestReceivedNoParamsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestReceivedNoParamsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestReceivedNoParams represents a ReceivedNoParams event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestReceivedNoParams struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceivedNoParams is a free log retrieval operation binding the contract event 0xbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a0.
//
// Solidity: event ReceivedNoParams(address sender)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterReceivedNoParams(opts *bind.FilterOpts) (*GatewayEVMZEVMTestReceivedNoParamsIterator, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "ReceivedNoParams")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestReceivedNoParamsIterator{contract: _GatewayEVMZEVMTest.contract, event: "ReceivedNoParams", logs: logs, sub: sub}, nil
}

// WatchReceivedNoParams is a free log subscription operation binding the contract event 0xbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a0.
//
// Solidity: event ReceivedNoParams(address sender)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchReceivedNoParams(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestReceivedNoParams) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "ReceivedNoParams")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestReceivedNoParams)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "ReceivedNoParams", log); err != nil {
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

// ParseReceivedNoParams is a log parse operation binding the contract event 0xbcaadb46b82a48af60b608f58959ae6b8310d1b0a0d094c2e9ec3208ed39f2a0.
//
// Solidity: event ReceivedNoParams(address sender)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseReceivedNoParams(log types.Log) (*GatewayEVMZEVMTestReceivedNoParams, error) {
	event := new(GatewayEVMZEVMTestReceivedNoParams)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "ReceivedNoParams", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestReceivedNonPayableIterator is returned from FilterReceivedNonPayable and is used to iterate over the raw logs and unpacked data for ReceivedNonPayable events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestReceivedNonPayableIterator struct {
	Event *GatewayEVMZEVMTestReceivedNonPayable // Event containing the contract specifics and raw log

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
func (it *GatewayEVMZEVMTestReceivedNonPayableIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestReceivedNonPayable)
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
		it.Event = new(GatewayEVMZEVMTestReceivedNonPayable)
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
func (it *GatewayEVMZEVMTestReceivedNonPayableIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestReceivedNonPayableIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestReceivedNonPayable represents a ReceivedNonPayable event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestReceivedNonPayable struct {
	Sender common.Address
	Strs   []string
	Nums   []*big.Int
	Flag   bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceivedNonPayable is a free log retrieval operation binding the contract event 0x74a53cd528a921fca7dbdee62f86819051d3cc98f214951f4238e8843f20b146.
//
// Solidity: event ReceivedNonPayable(address sender, string[] strs, uint256[] nums, bool flag)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterReceivedNonPayable(opts *bind.FilterOpts) (*GatewayEVMZEVMTestReceivedNonPayableIterator, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "ReceivedNonPayable")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestReceivedNonPayableIterator{contract: _GatewayEVMZEVMTest.contract, event: "ReceivedNonPayable", logs: logs, sub: sub}, nil
}

// WatchReceivedNonPayable is a free log subscription operation binding the contract event 0x74a53cd528a921fca7dbdee62f86819051d3cc98f214951f4238e8843f20b146.
//
// Solidity: event ReceivedNonPayable(address sender, string[] strs, uint256[] nums, bool flag)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchReceivedNonPayable(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestReceivedNonPayable) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "ReceivedNonPayable")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestReceivedNonPayable)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "ReceivedNonPayable", log); err != nil {
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

// ParseReceivedNonPayable is a log parse operation binding the contract event 0x74a53cd528a921fca7dbdee62f86819051d3cc98f214951f4238e8843f20b146.
//
// Solidity: event ReceivedNonPayable(address sender, string[] strs, uint256[] nums, bool flag)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseReceivedNonPayable(log types.Log) (*GatewayEVMZEVMTestReceivedNonPayable, error) {
	event := new(GatewayEVMZEVMTestReceivedNonPayable)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "ReceivedNonPayable", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestReceivedOnCallIterator is returned from FilterReceivedOnCall and is used to iterate over the raw logs and unpacked data for ReceivedOnCall events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestReceivedOnCallIterator struct {
	Event *GatewayEVMZEVMTestReceivedOnCall // Event containing the contract specifics and raw log

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
func (it *GatewayEVMZEVMTestReceivedOnCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestReceivedOnCall)
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
		it.Event = new(GatewayEVMZEVMTestReceivedOnCall)
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
func (it *GatewayEVMZEVMTestReceivedOnCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestReceivedOnCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestReceivedOnCall represents a ReceivedOnCall event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestReceivedOnCall struct {
	Sender  common.Address
	Message []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterReceivedOnCall is a free log retrieval operation binding the contract event 0xd80b62959d9a7e797f352e4015e65d345f402ea21972256fb0ba94f00a352501.
//
// Solidity: event ReceivedOnCall(address sender, bytes message)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterReceivedOnCall(opts *bind.FilterOpts) (*GatewayEVMZEVMTestReceivedOnCallIterator, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "ReceivedOnCall")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestReceivedOnCallIterator{contract: _GatewayEVMZEVMTest.contract, event: "ReceivedOnCall", logs: logs, sub: sub}, nil
}

// WatchReceivedOnCall is a free log subscription operation binding the contract event 0xd80b62959d9a7e797f352e4015e65d345f402ea21972256fb0ba94f00a352501.
//
// Solidity: event ReceivedOnCall(address sender, bytes message)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchReceivedOnCall(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestReceivedOnCall) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "ReceivedOnCall")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestReceivedOnCall)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "ReceivedOnCall", log); err != nil {
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

// ParseReceivedOnCall is a log parse operation binding the contract event 0xd80b62959d9a7e797f352e4015e65d345f402ea21972256fb0ba94f00a352501.
//
// Solidity: event ReceivedOnCall(address sender, bytes message)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseReceivedOnCall(log types.Log) (*GatewayEVMZEVMTestReceivedOnCall, error) {
	event := new(GatewayEVMZEVMTestReceivedOnCall)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "ReceivedOnCall", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestReceivedPayableIterator is returned from FilterReceivedPayable and is used to iterate over the raw logs and unpacked data for ReceivedPayable events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestReceivedPayableIterator struct {
	Event *GatewayEVMZEVMTestReceivedPayable // Event containing the contract specifics and raw log

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
func (it *GatewayEVMZEVMTestReceivedPayableIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestReceivedPayable)
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
		it.Event = new(GatewayEVMZEVMTestReceivedPayable)
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
func (it *GatewayEVMZEVMTestReceivedPayableIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestReceivedPayableIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestReceivedPayable represents a ReceivedPayable event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestReceivedPayable struct {
	Sender common.Address
	Value  *big.Int
	Str    string
	Num    *big.Int
	Flag   bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceivedPayable is a free log retrieval operation binding the contract event 0x1f1ff1f5fb41346850b2f5c04e6c767e2f1c8a525c5c0c5cdb60cdf3ca5f62fa.
//
// Solidity: event ReceivedPayable(address sender, uint256 value, string str, uint256 num, bool flag)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterReceivedPayable(opts *bind.FilterOpts) (*GatewayEVMZEVMTestReceivedPayableIterator, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "ReceivedPayable")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestReceivedPayableIterator{contract: _GatewayEVMZEVMTest.contract, event: "ReceivedPayable", logs: logs, sub: sub}, nil
}

// WatchReceivedPayable is a free log subscription operation binding the contract event 0x1f1ff1f5fb41346850b2f5c04e6c767e2f1c8a525c5c0c5cdb60cdf3ca5f62fa.
//
// Solidity: event ReceivedPayable(address sender, uint256 value, string str, uint256 num, bool flag)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchReceivedPayable(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestReceivedPayable) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "ReceivedPayable")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestReceivedPayable)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "ReceivedPayable", log); err != nil {
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

// ParseReceivedPayable is a log parse operation binding the contract event 0x1f1ff1f5fb41346850b2f5c04e6c767e2f1c8a525c5c0c5cdb60cdf3ca5f62fa.
//
// Solidity: event ReceivedPayable(address sender, uint256 value, string str, uint256 num, bool flag)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseReceivedPayable(log types.Log) (*GatewayEVMZEVMTestReceivedPayable, error) {
	event := new(GatewayEVMZEVMTestReceivedPayable)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "ReceivedPayable", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestReceivedRevertIterator is returned from FilterReceivedRevert and is used to iterate over the raw logs and unpacked data for ReceivedRevert events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestReceivedRevertIterator struct {
	Event *GatewayEVMZEVMTestReceivedRevert // Event containing the contract specifics and raw log

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
func (it *GatewayEVMZEVMTestReceivedRevertIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestReceivedRevert)
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
		it.Event = new(GatewayEVMZEVMTestReceivedRevert)
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
func (it *GatewayEVMZEVMTestReceivedRevertIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestReceivedRevertIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestReceivedRevert represents a ReceivedRevert event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestReceivedRevert struct {
	Sender        common.Address
	RevertContext RevertContext
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterReceivedRevert is a free log retrieval operation binding the contract event 0x689a5a5cb55e795ffe4cd8b419cd3bb0a3373974c54d25f64e734d7388b93e9b.
//
// Solidity: event ReceivedRevert(address sender, (address,address,uint256,bytes) revertContext)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterReceivedRevert(opts *bind.FilterOpts) (*GatewayEVMZEVMTestReceivedRevertIterator, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "ReceivedRevert")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestReceivedRevertIterator{contract: _GatewayEVMZEVMTest.contract, event: "ReceivedRevert", logs: logs, sub: sub}, nil
}

// WatchReceivedRevert is a free log subscription operation binding the contract event 0x689a5a5cb55e795ffe4cd8b419cd3bb0a3373974c54d25f64e734d7388b93e9b.
//
// Solidity: event ReceivedRevert(address sender, (address,address,uint256,bytes) revertContext)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchReceivedRevert(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestReceivedRevert) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "ReceivedRevert")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestReceivedRevert)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "ReceivedRevert", log); err != nil {
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

// ParseReceivedRevert is a log parse operation binding the contract event 0x689a5a5cb55e795ffe4cd8b419cd3bb0a3373974c54d25f64e734d7388b93e9b.
//
// Solidity: event ReceivedRevert(address sender, (address,address,uint256,bytes) revertContext)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseReceivedRevert(log types.Log) (*GatewayEVMZEVMTestReceivedRevert, error) {
	event := new(GatewayEVMZEVMTestReceivedRevert)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "ReceivedRevert", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestRevertedIterator is returned from FilterReverted and is used to iterate over the raw logs and unpacked data for Reverted events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestRevertedIterator struct {
	Event *GatewayEVMZEVMTestReverted // Event containing the contract specifics and raw log

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
func (it *GatewayEVMZEVMTestRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestReverted)
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
		it.Event = new(GatewayEVMZEVMTestReverted)
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
func (it *GatewayEVMZEVMTestRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestReverted represents a Reverted event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestReverted struct {
	To            common.Address
	Token         common.Address
	Amount        *big.Int
	Data          []byte
	RevertContext RevertContext
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterReverted is a free log retrieval operation binding the contract event 0xde7603a6ed5d07c9f43597ccfe9043d15b66d3284f0de321f5cdf56329e6e035.
//
// Solidity: event Reverted(address indexed to, address indexed token, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterReverted(opts *bind.FilterOpts, to []common.Address, token []common.Address) (*GatewayEVMZEVMTestRevertedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "Reverted", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestRevertedIterator{contract: _GatewayEVMZEVMTest.contract, event: "Reverted", logs: logs, sub: sub}, nil
}

// WatchReverted is a free log subscription operation binding the contract event 0xde7603a6ed5d07c9f43597ccfe9043d15b66d3284f0de321f5cdf56329e6e035.
//
// Solidity: event Reverted(address indexed to, address indexed token, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchReverted(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestReverted, to []common.Address, token []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "Reverted", toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestReverted)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "Reverted", log); err != nil {
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

// ParseReverted is a log parse operation binding the contract event 0xde7603a6ed5d07c9f43597ccfe9043d15b66d3284f0de321f5cdf56329e6e035.
//
// Solidity: event Reverted(address indexed to, address indexed token, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseReverted(log types.Log) (*GatewayEVMZEVMTestReverted, error) {
	event := new(GatewayEVMZEVMTestReverted)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "Reverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestUpdatedGatewayTSSAddressIterator is returned from FilterUpdatedGatewayTSSAddress and is used to iterate over the raw logs and unpacked data for UpdatedGatewayTSSAddress events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestUpdatedGatewayTSSAddressIterator struct {
	Event *GatewayEVMZEVMTestUpdatedGatewayTSSAddress // Event containing the contract specifics and raw log

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
func (it *GatewayEVMZEVMTestUpdatedGatewayTSSAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestUpdatedGatewayTSSAddress)
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
		it.Event = new(GatewayEVMZEVMTestUpdatedGatewayTSSAddress)
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
func (it *GatewayEVMZEVMTestUpdatedGatewayTSSAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestUpdatedGatewayTSSAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestUpdatedGatewayTSSAddress represents a UpdatedGatewayTSSAddress event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestUpdatedGatewayTSSAddress struct {
	OldTSSAddress common.Address
	NewTSSAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdatedGatewayTSSAddress is a free log retrieval operation binding the contract event 0x3a7b8d6372645f474fe60c115a2ef21421306a3ed4664fa0023c461413c08579.
//
// Solidity: event UpdatedGatewayTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterUpdatedGatewayTSSAddress(opts *bind.FilterOpts) (*GatewayEVMZEVMTestUpdatedGatewayTSSAddressIterator, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "UpdatedGatewayTSSAddress")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestUpdatedGatewayTSSAddressIterator{contract: _GatewayEVMZEVMTest.contract, event: "UpdatedGatewayTSSAddress", logs: logs, sub: sub}, nil
}

// WatchUpdatedGatewayTSSAddress is a free log subscription operation binding the contract event 0x3a7b8d6372645f474fe60c115a2ef21421306a3ed4664fa0023c461413c08579.
//
// Solidity: event UpdatedGatewayTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchUpdatedGatewayTSSAddress(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestUpdatedGatewayTSSAddress) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "UpdatedGatewayTSSAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestUpdatedGatewayTSSAddress)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "UpdatedGatewayTSSAddress", log); err != nil {
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

// ParseUpdatedGatewayTSSAddress is a log parse operation binding the contract event 0x3a7b8d6372645f474fe60c115a2ef21421306a3ed4664fa0023c461413c08579.
//
// Solidity: event UpdatedGatewayTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseUpdatedGatewayTSSAddress(log types.Log) (*GatewayEVMZEVMTestUpdatedGatewayTSSAddress, error) {
	event := new(GatewayEVMZEVMTestUpdatedGatewayTSSAddress)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "UpdatedGatewayTSSAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestWithdrawnIterator struct {
	Event *GatewayEVMZEVMTestWithdrawn // Event containing the contract specifics and raw log

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
func (it *GatewayEVMZEVMTestWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestWithdrawn)
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
		it.Event = new(GatewayEVMZEVMTestWithdrawn)
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
func (it *GatewayEVMZEVMTestWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestWithdrawn represents a Withdrawn event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestWithdrawn struct {
	Sender          common.Address
	ChainId         *big.Int
	Receiver        []byte
	Zrc20           common.Address
	Value           *big.Int
	Gasfee          *big.Int
	ProtocolFlatFee *big.Int
	Message         []byte
	CallOptions     CallOptions
	RevertOptions   RevertOptions
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x07bf64173efd8f3dfb9e4eb3834bab9d5b85a3d89a1c6425797329de0668502c.
//
// Solidity: event Withdrawn(address indexed sender, uint256 indexed chainId, bytes receiver, address zrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterWithdrawn(opts *bind.FilterOpts, sender []common.Address, chainId []*big.Int) (*GatewayEVMZEVMTestWithdrawnIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "Withdrawn", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestWithdrawnIterator{contract: _GatewayEVMZEVMTest.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x07bf64173efd8f3dfb9e4eb3834bab9d5b85a3d89a1c6425797329de0668502c.
//
// Solidity: event Withdrawn(address indexed sender, uint256 indexed chainId, bytes receiver, address zrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestWithdrawn, sender []common.Address, chainId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "Withdrawn", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestWithdrawn)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x07bf64173efd8f3dfb9e4eb3834bab9d5b85a3d89a1c6425797329de0668502c.
//
// Solidity: event Withdrawn(address indexed sender, uint256 indexed chainId, bytes receiver, address zrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseWithdrawn(log types.Log) (*GatewayEVMZEVMTestWithdrawn, error) {
	event := new(GatewayEVMZEVMTestWithdrawn)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestWithdrawnAndCalledIterator is returned from FilterWithdrawnAndCalled and is used to iterate over the raw logs and unpacked data for WithdrawnAndCalled events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestWithdrawnAndCalledIterator struct {
	Event *GatewayEVMZEVMTestWithdrawnAndCalled // Event containing the contract specifics and raw log

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
func (it *GatewayEVMZEVMTestWithdrawnAndCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestWithdrawnAndCalled)
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
		it.Event = new(GatewayEVMZEVMTestWithdrawnAndCalled)
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
func (it *GatewayEVMZEVMTestWithdrawnAndCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestWithdrawnAndCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestWithdrawnAndCalled represents a WithdrawnAndCalled event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestWithdrawnAndCalled struct {
	Sender          common.Address
	ChainId         *big.Int
	Receiver        []byte
	Zrc20           common.Address
	Value           *big.Int
	Gasfee          *big.Int
	ProtocolFlatFee *big.Int
	Message         []byte
	CallOptions     CallOptions
	RevertOptions   RevertOptions
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnAndCalled is a free log retrieval operation binding the contract event 0xd90f94752d2b12f364f4a2237ebe1aff24ba6127585376bf4935f6a7be17dd2a.
//
// Solidity: event WithdrawnAndCalled(address indexed sender, uint256 indexed chainId, bytes receiver, address zrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterWithdrawnAndCalled(opts *bind.FilterOpts, sender []common.Address, chainId []*big.Int) (*GatewayEVMZEVMTestWithdrawnAndCalledIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "WithdrawnAndCalled", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestWithdrawnAndCalledIterator{contract: _GatewayEVMZEVMTest.contract, event: "WithdrawnAndCalled", logs: logs, sub: sub}, nil
}

// WatchWithdrawnAndCalled is a free log subscription operation binding the contract event 0xd90f94752d2b12f364f4a2237ebe1aff24ba6127585376bf4935f6a7be17dd2a.
//
// Solidity: event WithdrawnAndCalled(address indexed sender, uint256 indexed chainId, bytes receiver, address zrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchWithdrawnAndCalled(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestWithdrawnAndCalled, sender []common.Address, chainId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "WithdrawnAndCalled", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestWithdrawnAndCalled)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "WithdrawnAndCalled", log); err != nil {
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

// ParseWithdrawnAndCalled is a log parse operation binding the contract event 0xd90f94752d2b12f364f4a2237ebe1aff24ba6127585376bf4935f6a7be17dd2a.
//
// Solidity: event WithdrawnAndCalled(address indexed sender, uint256 indexed chainId, bytes receiver, address zrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseWithdrawnAndCalled(log types.Log) (*GatewayEVMZEVMTestWithdrawnAndCalled, error) {
	event := new(GatewayEVMZEVMTestWithdrawnAndCalled)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "WithdrawnAndCalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestLogIterator is returned from FilterLog and is used to iterate over the raw logs and unpacked data for Log events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogIterator struct {
	Event *GatewayEVMZEVMTestLog // Event containing the contract specifics and raw log

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
func (it *GatewayEVMZEVMTestLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestLog)
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
		it.Event = new(GatewayEVMZEVMTestLog)
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
func (it *GatewayEVMZEVMTestLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestLog represents a Log event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLog struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLog is a free log retrieval operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterLog(opts *bind.FilterOpts) (*GatewayEVMZEVMTestLogIterator, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "log")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestLogIterator{contract: _GatewayEVMZEVMTest.contract, event: "log", logs: logs, sub: sub}, nil
}

// WatchLog is a free log subscription operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchLog(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestLog) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "log")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestLog)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log", log); err != nil {
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

// ParseLog is a log parse operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseLog(log types.Log) (*GatewayEVMZEVMTestLog, error) {
	event := new(GatewayEVMZEVMTestLog)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestLogAddressIterator is returned from FilterLogAddress and is used to iterate over the raw logs and unpacked data for LogAddress events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogAddressIterator struct {
	Event *GatewayEVMZEVMTestLogAddress // Event containing the contract specifics and raw log

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
func (it *GatewayEVMZEVMTestLogAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestLogAddress)
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
		it.Event = new(GatewayEVMZEVMTestLogAddress)
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
func (it *GatewayEVMZEVMTestLogAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestLogAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestLogAddress represents a LogAddress event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogAddress struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogAddress is a free log retrieval operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterLogAddress(opts *bind.FilterOpts) (*GatewayEVMZEVMTestLogAddressIterator, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "log_address")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestLogAddressIterator{contract: _GatewayEVMZEVMTest.contract, event: "log_address", logs: logs, sub: sub}, nil
}

// WatchLogAddress is a free log subscription operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchLogAddress(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestLogAddress) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "log_address")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestLogAddress)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_address", log); err != nil {
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

// ParseLogAddress is a log parse operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseLogAddress(log types.Log) (*GatewayEVMZEVMTestLogAddress, error) {
	event := new(GatewayEVMZEVMTestLogAddress)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_address", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestLogArrayIterator is returned from FilterLogArray and is used to iterate over the raw logs and unpacked data for LogArray events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogArrayIterator struct {
	Event *GatewayEVMZEVMTestLogArray // Event containing the contract specifics and raw log

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
func (it *GatewayEVMZEVMTestLogArrayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestLogArray)
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
		it.Event = new(GatewayEVMZEVMTestLogArray)
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
func (it *GatewayEVMZEVMTestLogArrayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestLogArrayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestLogArray represents a LogArray event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogArray struct {
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray is a free log retrieval operation binding the contract event 0xfb102865d50addddf69da9b5aa1bced66c80cf869a5c8d0471a467e18ce9cab1.
//
// Solidity: event log_array(uint256[] val)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterLogArray(opts *bind.FilterOpts) (*GatewayEVMZEVMTestLogArrayIterator, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "log_array")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestLogArrayIterator{contract: _GatewayEVMZEVMTest.contract, event: "log_array", logs: logs, sub: sub}, nil
}

// WatchLogArray is a free log subscription operation binding the contract event 0xfb102865d50addddf69da9b5aa1bced66c80cf869a5c8d0471a467e18ce9cab1.
//
// Solidity: event log_array(uint256[] val)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchLogArray(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestLogArray) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "log_array")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestLogArray)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_array", log); err != nil {
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

// ParseLogArray is a log parse operation binding the contract event 0xfb102865d50addddf69da9b5aa1bced66c80cf869a5c8d0471a467e18ce9cab1.
//
// Solidity: event log_array(uint256[] val)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseLogArray(log types.Log) (*GatewayEVMZEVMTestLogArray, error) {
	event := new(GatewayEVMZEVMTestLogArray)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_array", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestLogArray0Iterator is returned from FilterLogArray0 and is used to iterate over the raw logs and unpacked data for LogArray0 events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogArray0Iterator struct {
	Event *GatewayEVMZEVMTestLogArray0 // Event containing the contract specifics and raw log

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
func (it *GatewayEVMZEVMTestLogArray0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestLogArray0)
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
		it.Event = new(GatewayEVMZEVMTestLogArray0)
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
func (it *GatewayEVMZEVMTestLogArray0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestLogArray0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestLogArray0 represents a LogArray0 event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogArray0 struct {
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray0 is a free log retrieval operation binding the contract event 0x890a82679b470f2bd82816ed9b161f97d8b967f37fa3647c21d5bf39749e2dd5.
//
// Solidity: event log_array(int256[] val)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterLogArray0(opts *bind.FilterOpts) (*GatewayEVMZEVMTestLogArray0Iterator, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "log_array0")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestLogArray0Iterator{contract: _GatewayEVMZEVMTest.contract, event: "log_array0", logs: logs, sub: sub}, nil
}

// WatchLogArray0 is a free log subscription operation binding the contract event 0x890a82679b470f2bd82816ed9b161f97d8b967f37fa3647c21d5bf39749e2dd5.
//
// Solidity: event log_array(int256[] val)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchLogArray0(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestLogArray0) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "log_array0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestLogArray0)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_array0", log); err != nil {
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

// ParseLogArray0 is a log parse operation binding the contract event 0x890a82679b470f2bd82816ed9b161f97d8b967f37fa3647c21d5bf39749e2dd5.
//
// Solidity: event log_array(int256[] val)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseLogArray0(log types.Log) (*GatewayEVMZEVMTestLogArray0, error) {
	event := new(GatewayEVMZEVMTestLogArray0)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_array0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestLogArray1Iterator is returned from FilterLogArray1 and is used to iterate over the raw logs and unpacked data for LogArray1 events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogArray1Iterator struct {
	Event *GatewayEVMZEVMTestLogArray1 // Event containing the contract specifics and raw log

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
func (it *GatewayEVMZEVMTestLogArray1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestLogArray1)
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
		it.Event = new(GatewayEVMZEVMTestLogArray1)
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
func (it *GatewayEVMZEVMTestLogArray1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestLogArray1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestLogArray1 represents a LogArray1 event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogArray1 struct {
	Val []common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray1 is a free log retrieval operation binding the contract event 0x40e1840f5769073d61bd01372d9b75baa9842d5629a0c99ff103be1178a8e9e2.
//
// Solidity: event log_array(address[] val)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterLogArray1(opts *bind.FilterOpts) (*GatewayEVMZEVMTestLogArray1Iterator, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "log_array1")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestLogArray1Iterator{contract: _GatewayEVMZEVMTest.contract, event: "log_array1", logs: logs, sub: sub}, nil
}

// WatchLogArray1 is a free log subscription operation binding the contract event 0x40e1840f5769073d61bd01372d9b75baa9842d5629a0c99ff103be1178a8e9e2.
//
// Solidity: event log_array(address[] val)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchLogArray1(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestLogArray1) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "log_array1")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestLogArray1)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_array1", log); err != nil {
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

// ParseLogArray1 is a log parse operation binding the contract event 0x40e1840f5769073d61bd01372d9b75baa9842d5629a0c99ff103be1178a8e9e2.
//
// Solidity: event log_array(address[] val)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseLogArray1(log types.Log) (*GatewayEVMZEVMTestLogArray1, error) {
	event := new(GatewayEVMZEVMTestLogArray1)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_array1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestLogBytesIterator is returned from FilterLogBytes and is used to iterate over the raw logs and unpacked data for LogBytes events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogBytesIterator struct {
	Event *GatewayEVMZEVMTestLogBytes // Event containing the contract specifics and raw log

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
func (it *GatewayEVMZEVMTestLogBytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestLogBytes)
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
		it.Event = new(GatewayEVMZEVMTestLogBytes)
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
func (it *GatewayEVMZEVMTestLogBytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestLogBytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestLogBytes represents a LogBytes event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogBytes struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogBytes is a free log retrieval operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterLogBytes(opts *bind.FilterOpts) (*GatewayEVMZEVMTestLogBytesIterator, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "log_bytes")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestLogBytesIterator{contract: _GatewayEVMZEVMTest.contract, event: "log_bytes", logs: logs, sub: sub}, nil
}

// WatchLogBytes is a free log subscription operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchLogBytes(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestLogBytes) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "log_bytes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestLogBytes)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_bytes", log); err != nil {
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

// ParseLogBytes is a log parse operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseLogBytes(log types.Log) (*GatewayEVMZEVMTestLogBytes, error) {
	event := new(GatewayEVMZEVMTestLogBytes)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_bytes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestLogBytes32Iterator is returned from FilterLogBytes32 and is used to iterate over the raw logs and unpacked data for LogBytes32 events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogBytes32Iterator struct {
	Event *GatewayEVMZEVMTestLogBytes32 // Event containing the contract specifics and raw log

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
func (it *GatewayEVMZEVMTestLogBytes32Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestLogBytes32)
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
		it.Event = new(GatewayEVMZEVMTestLogBytes32)
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
func (it *GatewayEVMZEVMTestLogBytes32Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestLogBytes32Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestLogBytes32 represents a LogBytes32 event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogBytes32 struct {
	Arg0 [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogBytes32 is a free log retrieval operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterLogBytes32(opts *bind.FilterOpts) (*GatewayEVMZEVMTestLogBytes32Iterator, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "log_bytes32")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestLogBytes32Iterator{contract: _GatewayEVMZEVMTest.contract, event: "log_bytes32", logs: logs, sub: sub}, nil
}

// WatchLogBytes32 is a free log subscription operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchLogBytes32(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestLogBytes32) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "log_bytes32")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestLogBytes32)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_bytes32", log); err != nil {
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

// ParseLogBytes32 is a log parse operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseLogBytes32(log types.Log) (*GatewayEVMZEVMTestLogBytes32, error) {
	event := new(GatewayEVMZEVMTestLogBytes32)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_bytes32", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestLogIntIterator is returned from FilterLogInt and is used to iterate over the raw logs and unpacked data for LogInt events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogIntIterator struct {
	Event *GatewayEVMZEVMTestLogInt // Event containing the contract specifics and raw log

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
func (it *GatewayEVMZEVMTestLogIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestLogInt)
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
		it.Event = new(GatewayEVMZEVMTestLogInt)
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
func (it *GatewayEVMZEVMTestLogIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestLogIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestLogInt represents a LogInt event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogInt struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogInt is a free log retrieval operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterLogInt(opts *bind.FilterOpts) (*GatewayEVMZEVMTestLogIntIterator, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "log_int")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestLogIntIterator{contract: _GatewayEVMZEVMTest.contract, event: "log_int", logs: logs, sub: sub}, nil
}

// WatchLogInt is a free log subscription operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchLogInt(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestLogInt) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "log_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestLogInt)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_int", log); err != nil {
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

// ParseLogInt is a log parse operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseLogInt(log types.Log) (*GatewayEVMZEVMTestLogInt, error) {
	event := new(GatewayEVMZEVMTestLogInt)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestLogNamedAddressIterator is returned from FilterLogNamedAddress and is used to iterate over the raw logs and unpacked data for LogNamedAddress events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogNamedAddressIterator struct {
	Event *GatewayEVMZEVMTestLogNamedAddress // Event containing the contract specifics and raw log

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
func (it *GatewayEVMZEVMTestLogNamedAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestLogNamedAddress)
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
		it.Event = new(GatewayEVMZEVMTestLogNamedAddress)
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
func (it *GatewayEVMZEVMTestLogNamedAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestLogNamedAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestLogNamedAddress represents a LogNamedAddress event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogNamedAddress struct {
	Key string
	Val common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedAddress is a free log retrieval operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterLogNamedAddress(opts *bind.FilterOpts) (*GatewayEVMZEVMTestLogNamedAddressIterator, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "log_named_address")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestLogNamedAddressIterator{contract: _GatewayEVMZEVMTest.contract, event: "log_named_address", logs: logs, sub: sub}, nil
}

// WatchLogNamedAddress is a free log subscription operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchLogNamedAddress(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestLogNamedAddress) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "log_named_address")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestLogNamedAddress)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_named_address", log); err != nil {
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

// ParseLogNamedAddress is a log parse operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseLogNamedAddress(log types.Log) (*GatewayEVMZEVMTestLogNamedAddress, error) {
	event := new(GatewayEVMZEVMTestLogNamedAddress)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_named_address", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestLogNamedArrayIterator is returned from FilterLogNamedArray and is used to iterate over the raw logs and unpacked data for LogNamedArray events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogNamedArrayIterator struct {
	Event *GatewayEVMZEVMTestLogNamedArray // Event containing the contract specifics and raw log

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
func (it *GatewayEVMZEVMTestLogNamedArrayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestLogNamedArray)
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
		it.Event = new(GatewayEVMZEVMTestLogNamedArray)
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
func (it *GatewayEVMZEVMTestLogNamedArrayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestLogNamedArrayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestLogNamedArray represents a LogNamedArray event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogNamedArray struct {
	Key string
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray is a free log retrieval operation binding the contract event 0x00aaa39c9ffb5f567a4534380c737075702e1f7f14107fc95328e3b56c0325fb.
//
// Solidity: event log_named_array(string key, uint256[] val)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterLogNamedArray(opts *bind.FilterOpts) (*GatewayEVMZEVMTestLogNamedArrayIterator, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "log_named_array")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestLogNamedArrayIterator{contract: _GatewayEVMZEVMTest.contract, event: "log_named_array", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray is a free log subscription operation binding the contract event 0x00aaa39c9ffb5f567a4534380c737075702e1f7f14107fc95328e3b56c0325fb.
//
// Solidity: event log_named_array(string key, uint256[] val)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchLogNamedArray(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestLogNamedArray) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "log_named_array")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestLogNamedArray)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_named_array", log); err != nil {
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

// ParseLogNamedArray is a log parse operation binding the contract event 0x00aaa39c9ffb5f567a4534380c737075702e1f7f14107fc95328e3b56c0325fb.
//
// Solidity: event log_named_array(string key, uint256[] val)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseLogNamedArray(log types.Log) (*GatewayEVMZEVMTestLogNamedArray, error) {
	event := new(GatewayEVMZEVMTestLogNamedArray)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_named_array", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestLogNamedArray0Iterator is returned from FilterLogNamedArray0 and is used to iterate over the raw logs and unpacked data for LogNamedArray0 events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogNamedArray0Iterator struct {
	Event *GatewayEVMZEVMTestLogNamedArray0 // Event containing the contract specifics and raw log

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
func (it *GatewayEVMZEVMTestLogNamedArray0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestLogNamedArray0)
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
		it.Event = new(GatewayEVMZEVMTestLogNamedArray0)
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
func (it *GatewayEVMZEVMTestLogNamedArray0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestLogNamedArray0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestLogNamedArray0 represents a LogNamedArray0 event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogNamedArray0 struct {
	Key string
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray0 is a free log retrieval operation binding the contract event 0xa73eda09662f46dde729be4611385ff34fe6c44fbbc6f7e17b042b59a3445b57.
//
// Solidity: event log_named_array(string key, int256[] val)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterLogNamedArray0(opts *bind.FilterOpts) (*GatewayEVMZEVMTestLogNamedArray0Iterator, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "log_named_array0")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestLogNamedArray0Iterator{contract: _GatewayEVMZEVMTest.contract, event: "log_named_array0", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray0 is a free log subscription operation binding the contract event 0xa73eda09662f46dde729be4611385ff34fe6c44fbbc6f7e17b042b59a3445b57.
//
// Solidity: event log_named_array(string key, int256[] val)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchLogNamedArray0(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestLogNamedArray0) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "log_named_array0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestLogNamedArray0)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_named_array0", log); err != nil {
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

// ParseLogNamedArray0 is a log parse operation binding the contract event 0xa73eda09662f46dde729be4611385ff34fe6c44fbbc6f7e17b042b59a3445b57.
//
// Solidity: event log_named_array(string key, int256[] val)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseLogNamedArray0(log types.Log) (*GatewayEVMZEVMTestLogNamedArray0, error) {
	event := new(GatewayEVMZEVMTestLogNamedArray0)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_named_array0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestLogNamedArray1Iterator is returned from FilterLogNamedArray1 and is used to iterate over the raw logs and unpacked data for LogNamedArray1 events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogNamedArray1Iterator struct {
	Event *GatewayEVMZEVMTestLogNamedArray1 // Event containing the contract specifics and raw log

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
func (it *GatewayEVMZEVMTestLogNamedArray1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestLogNamedArray1)
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
		it.Event = new(GatewayEVMZEVMTestLogNamedArray1)
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
func (it *GatewayEVMZEVMTestLogNamedArray1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestLogNamedArray1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestLogNamedArray1 represents a LogNamedArray1 event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogNamedArray1 struct {
	Key string
	Val []common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray1 is a free log retrieval operation binding the contract event 0x3bcfb2ae2e8d132dd1fce7cf278a9a19756a9fceabe470df3bdabb4bc577d1bd.
//
// Solidity: event log_named_array(string key, address[] val)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterLogNamedArray1(opts *bind.FilterOpts) (*GatewayEVMZEVMTestLogNamedArray1Iterator, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "log_named_array1")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestLogNamedArray1Iterator{contract: _GatewayEVMZEVMTest.contract, event: "log_named_array1", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray1 is a free log subscription operation binding the contract event 0x3bcfb2ae2e8d132dd1fce7cf278a9a19756a9fceabe470df3bdabb4bc577d1bd.
//
// Solidity: event log_named_array(string key, address[] val)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchLogNamedArray1(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestLogNamedArray1) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "log_named_array1")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestLogNamedArray1)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_named_array1", log); err != nil {
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

// ParseLogNamedArray1 is a log parse operation binding the contract event 0x3bcfb2ae2e8d132dd1fce7cf278a9a19756a9fceabe470df3bdabb4bc577d1bd.
//
// Solidity: event log_named_array(string key, address[] val)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseLogNamedArray1(log types.Log) (*GatewayEVMZEVMTestLogNamedArray1, error) {
	event := new(GatewayEVMZEVMTestLogNamedArray1)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_named_array1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestLogNamedBytesIterator is returned from FilterLogNamedBytes and is used to iterate over the raw logs and unpacked data for LogNamedBytes events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogNamedBytesIterator struct {
	Event *GatewayEVMZEVMTestLogNamedBytes // Event containing the contract specifics and raw log

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
func (it *GatewayEVMZEVMTestLogNamedBytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestLogNamedBytes)
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
		it.Event = new(GatewayEVMZEVMTestLogNamedBytes)
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
func (it *GatewayEVMZEVMTestLogNamedBytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestLogNamedBytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestLogNamedBytes represents a LogNamedBytes event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogNamedBytes struct {
	Key string
	Val []byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedBytes is a free log retrieval operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterLogNamedBytes(opts *bind.FilterOpts) (*GatewayEVMZEVMTestLogNamedBytesIterator, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "log_named_bytes")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestLogNamedBytesIterator{contract: _GatewayEVMZEVMTest.contract, event: "log_named_bytes", logs: logs, sub: sub}, nil
}

// WatchLogNamedBytes is a free log subscription operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchLogNamedBytes(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestLogNamedBytes) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "log_named_bytes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestLogNamedBytes)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_named_bytes", log); err != nil {
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

// ParseLogNamedBytes is a log parse operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseLogNamedBytes(log types.Log) (*GatewayEVMZEVMTestLogNamedBytes, error) {
	event := new(GatewayEVMZEVMTestLogNamedBytes)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_named_bytes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestLogNamedBytes32Iterator is returned from FilterLogNamedBytes32 and is used to iterate over the raw logs and unpacked data for LogNamedBytes32 events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogNamedBytes32Iterator struct {
	Event *GatewayEVMZEVMTestLogNamedBytes32 // Event containing the contract specifics and raw log

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
func (it *GatewayEVMZEVMTestLogNamedBytes32Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestLogNamedBytes32)
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
		it.Event = new(GatewayEVMZEVMTestLogNamedBytes32)
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
func (it *GatewayEVMZEVMTestLogNamedBytes32Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestLogNamedBytes32Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestLogNamedBytes32 represents a LogNamedBytes32 event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogNamedBytes32 struct {
	Key string
	Val [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedBytes32 is a free log retrieval operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterLogNamedBytes32(opts *bind.FilterOpts) (*GatewayEVMZEVMTestLogNamedBytes32Iterator, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "log_named_bytes32")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestLogNamedBytes32Iterator{contract: _GatewayEVMZEVMTest.contract, event: "log_named_bytes32", logs: logs, sub: sub}, nil
}

// WatchLogNamedBytes32 is a free log subscription operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchLogNamedBytes32(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestLogNamedBytes32) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "log_named_bytes32")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestLogNamedBytes32)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_named_bytes32", log); err != nil {
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

// ParseLogNamedBytes32 is a log parse operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseLogNamedBytes32(log types.Log) (*GatewayEVMZEVMTestLogNamedBytes32, error) {
	event := new(GatewayEVMZEVMTestLogNamedBytes32)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_named_bytes32", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestLogNamedDecimalIntIterator is returned from FilterLogNamedDecimalInt and is used to iterate over the raw logs and unpacked data for LogNamedDecimalInt events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogNamedDecimalIntIterator struct {
	Event *GatewayEVMZEVMTestLogNamedDecimalInt // Event containing the contract specifics and raw log

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
func (it *GatewayEVMZEVMTestLogNamedDecimalIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestLogNamedDecimalInt)
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
		it.Event = new(GatewayEVMZEVMTestLogNamedDecimalInt)
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
func (it *GatewayEVMZEVMTestLogNamedDecimalIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestLogNamedDecimalIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestLogNamedDecimalInt represents a LogNamedDecimalInt event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogNamedDecimalInt struct {
	Key      string
	Val      *big.Int
	Decimals *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogNamedDecimalInt is a free log retrieval operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterLogNamedDecimalInt(opts *bind.FilterOpts) (*GatewayEVMZEVMTestLogNamedDecimalIntIterator, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "log_named_decimal_int")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestLogNamedDecimalIntIterator{contract: _GatewayEVMZEVMTest.contract, event: "log_named_decimal_int", logs: logs, sub: sub}, nil
}

// WatchLogNamedDecimalInt is a free log subscription operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchLogNamedDecimalInt(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestLogNamedDecimalInt) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "log_named_decimal_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestLogNamedDecimalInt)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_named_decimal_int", log); err != nil {
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

// ParseLogNamedDecimalInt is a log parse operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseLogNamedDecimalInt(log types.Log) (*GatewayEVMZEVMTestLogNamedDecimalInt, error) {
	event := new(GatewayEVMZEVMTestLogNamedDecimalInt)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_named_decimal_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestLogNamedDecimalUintIterator is returned from FilterLogNamedDecimalUint and is used to iterate over the raw logs and unpacked data for LogNamedDecimalUint events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogNamedDecimalUintIterator struct {
	Event *GatewayEVMZEVMTestLogNamedDecimalUint // Event containing the contract specifics and raw log

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
func (it *GatewayEVMZEVMTestLogNamedDecimalUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestLogNamedDecimalUint)
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
		it.Event = new(GatewayEVMZEVMTestLogNamedDecimalUint)
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
func (it *GatewayEVMZEVMTestLogNamedDecimalUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestLogNamedDecimalUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestLogNamedDecimalUint represents a LogNamedDecimalUint event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogNamedDecimalUint struct {
	Key      string
	Val      *big.Int
	Decimals *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogNamedDecimalUint is a free log retrieval operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterLogNamedDecimalUint(opts *bind.FilterOpts) (*GatewayEVMZEVMTestLogNamedDecimalUintIterator, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "log_named_decimal_uint")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestLogNamedDecimalUintIterator{contract: _GatewayEVMZEVMTest.contract, event: "log_named_decimal_uint", logs: logs, sub: sub}, nil
}

// WatchLogNamedDecimalUint is a free log subscription operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchLogNamedDecimalUint(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestLogNamedDecimalUint) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "log_named_decimal_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestLogNamedDecimalUint)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_named_decimal_uint", log); err != nil {
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

// ParseLogNamedDecimalUint is a log parse operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseLogNamedDecimalUint(log types.Log) (*GatewayEVMZEVMTestLogNamedDecimalUint, error) {
	event := new(GatewayEVMZEVMTestLogNamedDecimalUint)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_named_decimal_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestLogNamedIntIterator is returned from FilterLogNamedInt and is used to iterate over the raw logs and unpacked data for LogNamedInt events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogNamedIntIterator struct {
	Event *GatewayEVMZEVMTestLogNamedInt // Event containing the contract specifics and raw log

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
func (it *GatewayEVMZEVMTestLogNamedIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestLogNamedInt)
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
		it.Event = new(GatewayEVMZEVMTestLogNamedInt)
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
func (it *GatewayEVMZEVMTestLogNamedIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestLogNamedIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestLogNamedInt represents a LogNamedInt event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogNamedInt struct {
	Key string
	Val *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedInt is a free log retrieval operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterLogNamedInt(opts *bind.FilterOpts) (*GatewayEVMZEVMTestLogNamedIntIterator, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "log_named_int")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestLogNamedIntIterator{contract: _GatewayEVMZEVMTest.contract, event: "log_named_int", logs: logs, sub: sub}, nil
}

// WatchLogNamedInt is a free log subscription operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchLogNamedInt(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestLogNamedInt) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "log_named_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestLogNamedInt)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_named_int", log); err != nil {
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

// ParseLogNamedInt is a log parse operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseLogNamedInt(log types.Log) (*GatewayEVMZEVMTestLogNamedInt, error) {
	event := new(GatewayEVMZEVMTestLogNamedInt)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_named_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestLogNamedStringIterator is returned from FilterLogNamedString and is used to iterate over the raw logs and unpacked data for LogNamedString events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogNamedStringIterator struct {
	Event *GatewayEVMZEVMTestLogNamedString // Event containing the contract specifics and raw log

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
func (it *GatewayEVMZEVMTestLogNamedStringIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestLogNamedString)
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
		it.Event = new(GatewayEVMZEVMTestLogNamedString)
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
func (it *GatewayEVMZEVMTestLogNamedStringIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestLogNamedStringIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestLogNamedString represents a LogNamedString event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogNamedString struct {
	Key string
	Val string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedString is a free log retrieval operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterLogNamedString(opts *bind.FilterOpts) (*GatewayEVMZEVMTestLogNamedStringIterator, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "log_named_string")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestLogNamedStringIterator{contract: _GatewayEVMZEVMTest.contract, event: "log_named_string", logs: logs, sub: sub}, nil
}

// WatchLogNamedString is a free log subscription operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchLogNamedString(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestLogNamedString) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "log_named_string")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestLogNamedString)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_named_string", log); err != nil {
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

// ParseLogNamedString is a log parse operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseLogNamedString(log types.Log) (*GatewayEVMZEVMTestLogNamedString, error) {
	event := new(GatewayEVMZEVMTestLogNamedString)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_named_string", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestLogNamedUintIterator is returned from FilterLogNamedUint and is used to iterate over the raw logs and unpacked data for LogNamedUint events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogNamedUintIterator struct {
	Event *GatewayEVMZEVMTestLogNamedUint // Event containing the contract specifics and raw log

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
func (it *GatewayEVMZEVMTestLogNamedUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestLogNamedUint)
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
		it.Event = new(GatewayEVMZEVMTestLogNamedUint)
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
func (it *GatewayEVMZEVMTestLogNamedUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestLogNamedUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestLogNamedUint represents a LogNamedUint event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogNamedUint struct {
	Key string
	Val *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedUint is a free log retrieval operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterLogNamedUint(opts *bind.FilterOpts) (*GatewayEVMZEVMTestLogNamedUintIterator, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "log_named_uint")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestLogNamedUintIterator{contract: _GatewayEVMZEVMTest.contract, event: "log_named_uint", logs: logs, sub: sub}, nil
}

// WatchLogNamedUint is a free log subscription operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchLogNamedUint(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestLogNamedUint) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "log_named_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestLogNamedUint)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_named_uint", log); err != nil {
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

// ParseLogNamedUint is a log parse operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseLogNamedUint(log types.Log) (*GatewayEVMZEVMTestLogNamedUint, error) {
	event := new(GatewayEVMZEVMTestLogNamedUint)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_named_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestLogStringIterator is returned from FilterLogString and is used to iterate over the raw logs and unpacked data for LogString events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogStringIterator struct {
	Event *GatewayEVMZEVMTestLogString // Event containing the contract specifics and raw log

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
func (it *GatewayEVMZEVMTestLogStringIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestLogString)
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
		it.Event = new(GatewayEVMZEVMTestLogString)
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
func (it *GatewayEVMZEVMTestLogStringIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestLogStringIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestLogString represents a LogString event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogString struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogString is a free log retrieval operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterLogString(opts *bind.FilterOpts) (*GatewayEVMZEVMTestLogStringIterator, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "log_string")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestLogStringIterator{contract: _GatewayEVMZEVMTest.contract, event: "log_string", logs: logs, sub: sub}, nil
}

// WatchLogString is a free log subscription operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchLogString(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestLogString) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "log_string")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestLogString)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_string", log); err != nil {
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

// ParseLogString is a log parse operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseLogString(log types.Log) (*GatewayEVMZEVMTestLogString, error) {
	event := new(GatewayEVMZEVMTestLogString)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_string", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestLogUintIterator is returned from FilterLogUint and is used to iterate over the raw logs and unpacked data for LogUint events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogUintIterator struct {
	Event *GatewayEVMZEVMTestLogUint // Event containing the contract specifics and raw log

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
func (it *GatewayEVMZEVMTestLogUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestLogUint)
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
		it.Event = new(GatewayEVMZEVMTestLogUint)
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
func (it *GatewayEVMZEVMTestLogUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestLogUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestLogUint represents a LogUint event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogUint struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogUint is a free log retrieval operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterLogUint(opts *bind.FilterOpts) (*GatewayEVMZEVMTestLogUintIterator, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "log_uint")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestLogUintIterator{contract: _GatewayEVMZEVMTest.contract, event: "log_uint", logs: logs, sub: sub}, nil
}

// WatchLogUint is a free log subscription operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchLogUint(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestLogUint) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "log_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestLogUint)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_uint", log); err != nil {
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

// ParseLogUint is a log parse operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseLogUint(log types.Log) (*GatewayEVMZEVMTestLogUint, error) {
	event := new(GatewayEVMZEVMTestLogUint)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "log_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayEVMZEVMTestLogsIterator is returned from FilterLogs and is used to iterate over the raw logs and unpacked data for Logs events raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogsIterator struct {
	Event *GatewayEVMZEVMTestLogs // Event containing the contract specifics and raw log

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
func (it *GatewayEVMZEVMTestLogsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayEVMZEVMTestLogs)
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
		it.Event = new(GatewayEVMZEVMTestLogs)
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
func (it *GatewayEVMZEVMTestLogsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayEVMZEVMTestLogsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayEVMZEVMTestLogs represents a Logs event raised by the GatewayEVMZEVMTest contract.
type GatewayEVMZEVMTestLogs struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogs is a free log retrieval operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) FilterLogs(opts *bind.FilterOpts) (*GatewayEVMZEVMTestLogsIterator, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.FilterLogs(opts, "logs")
	if err != nil {
		return nil, err
	}
	return &GatewayEVMZEVMTestLogsIterator{contract: _GatewayEVMZEVMTest.contract, event: "logs", logs: logs, sub: sub}, nil
}

// WatchLogs is a free log subscription operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) WatchLogs(opts *bind.WatchOpts, sink chan<- *GatewayEVMZEVMTestLogs) (event.Subscription, error) {

	logs, sub, err := _GatewayEVMZEVMTest.contract.WatchLogs(opts, "logs")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayEVMZEVMTestLogs)
				if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "logs", log); err != nil {
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

// ParseLogs is a log parse operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_GatewayEVMZEVMTest *GatewayEVMZEVMTestFilterer) ParseLogs(log types.Log) (*GatewayEVMZEVMTestLogs, error) {
	event := new(GatewayEVMZEVMTestLogs)
	if err := _GatewayEVMZEVMTest.contract.UnpackLog(event, "logs", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
