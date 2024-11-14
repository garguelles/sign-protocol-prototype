// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// Attestation is an auto generated low-level Go binding around an user-defined struct.
type Attestation struct {
	SchemaId            uint64
	LinkedAttestationId uint64
	AttestTimestamp     uint64
	RevokeTimestamp     uint64
	Attester            common.Address
	ValidUntil          uint64
	DataLocation        uint8
	Revoked             bool
	Recipients          [][]byte
	Data                []byte
}

// OffchainAttestation is an auto generated low-level Go binding around an user-defined struct.
type OffchainAttestation struct {
	Attester  common.Address
	Timestamp uint64
}

// Schema is an auto generated low-level Go binding around an user-defined struct.
type Schema struct {
	Registrant   common.Address
	Revocable    bool
	DataLocation uint8
	MaxValidFor  uint64
	Hook         common.Address
	Timestamp    uint64
	Data         string
}

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AttestationAlreadyRevoked\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AttestationInvalidDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AttestationIrrevocable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AttestationNonexistent\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AttestationWrongAttester\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDelegateSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LegacySPRequired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OffchainAttestationAlreadyRevoked\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OffchainAttestationExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OffchainAttestationNonexistent\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Paused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SchemaNonexistent\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SchemaWrongRegistrant\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"attestationId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"indexingKey\",\"type\":\"string\"}],\"name\":\"AttestationMade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"attestationId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"AttestationRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"attestationId\",\"type\":\"string\"}],\"name\":\"OffchainAttestationMade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"attestationId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"OffchainAttestationRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"schemaId\",\"type\":\"uint64\"}],\"name\":\"SchemaRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"schemaId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"linkedAttestationId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"attestTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revokeTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"attester\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"validUntil\",\"type\":\"uint64\"},{\"internalType\":\"enumDataLocation\",\"name\":\"dataLocation\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"revoked\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"recipients\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structAttestation\",\"name\":\"attestation\",\"type\":\"tuple\"},{\"internalType\":\"contractIERC20\",\"name\":\"resolverFeesERC20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"resolverFeesERC20Amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"indexingKey\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"delegateSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"attest\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"schemaId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"linkedAttestationId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"attestTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revokeTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"attester\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"validUntil\",\"type\":\"uint64\"},{\"internalType\":\"enumDataLocation\",\"name\":\"dataLocation\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"revoked\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"recipients\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structAttestation\",\"name\":\"attestation\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"indexingKey\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"delegateSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"attest\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"schemaId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"linkedAttestationId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"attestTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revokeTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"attester\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"validUntil\",\"type\":\"uint64\"},{\"internalType\":\"enumDataLocation\",\"name\":\"dataLocation\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"revoked\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"recipients\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structAttestation\",\"name\":\"attestation\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"resolverFeesETH\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"indexingKey\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"delegateSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"attest\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"schemaId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"linkedAttestationId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"attestTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revokeTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"attester\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"validUntil\",\"type\":\"uint64\"},{\"internalType\":\"enumDataLocation\",\"name\":\"dataLocation\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"revoked\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"recipients\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structAttestation[]\",\"name\":\"attestations\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"resolverFeesETH\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"indexingKeys\",\"type\":\"string[]\"},{\"internalType\":\"bytes\",\"name\":\"delegateSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"attestBatch\",\"outputs\":[{\"internalType\":\"uint64[]\",\"name\":\"attestationIds\",\"type\":\"uint64[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"schemaId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"linkedAttestationId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"attestTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revokeTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"attester\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"validUntil\",\"type\":\"uint64\"},{\"internalType\":\"enumDataLocation\",\"name\":\"dataLocation\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"revoked\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"recipients\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structAttestation[]\",\"name\":\"attestations\",\"type\":\"tuple[]\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"resolverFeesERC20Tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"resolverFeesERC20Amount\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"indexingKeys\",\"type\":\"string[]\"},{\"internalType\":\"bytes\",\"name\":\"delegateSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"attestBatch\",\"outputs\":[{\"internalType\":\"uint64[]\",\"name\":\"attestationIds\",\"type\":\"uint64[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"schemaId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"linkedAttestationId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"attestTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revokeTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"attester\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"validUntil\",\"type\":\"uint64\"},{\"internalType\":\"enumDataLocation\",\"name\":\"dataLocation\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"revoked\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"recipients\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structAttestation[]\",\"name\":\"attestations\",\"type\":\"tuple[]\"},{\"internalType\":\"string[]\",\"name\":\"indexingKeys\",\"type\":\"string[]\"},{\"internalType\":\"bytes\",\"name\":\"delegateSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"attestBatch\",\"outputs\":[{\"internalType\":\"uint64[]\",\"name\":\"attestationIds\",\"type\":\"uint64[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"offchainAttestationId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"delegateAttester\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"delegateSignature\",\"type\":\"bytes\"}],\"name\":\"attestOffchain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"attestationIds\",\"type\":\"string[]\"},{\"internalType\":\"address\",\"name\":\"delegateAttester\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"delegateSignature\",\"type\":\"bytes\"}],\"name\":\"attestOffchainBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"attestationCounter\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"attestationId\",\"type\":\"uint64\"}],\"name\":\"getAttestation\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"schemaId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"linkedAttestationId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"attestTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revokeTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"attester\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"validUntil\",\"type\":\"uint64\"},{\"internalType\":\"enumDataLocation\",\"name\":\"dataLocation\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"revoked\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"recipients\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structAttestation\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"schemaId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"linkedAttestationId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"attestTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revokeTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"attester\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"validUntil\",\"type\":\"uint64\"},{\"internalType\":\"enumDataLocation\",\"name\":\"dataLocation\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"revoked\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"recipients\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structAttestation[]\",\"name\":\"attestations\",\"type\":\"tuple[]\"}],\"name\":\"getDelegatedAttestBatchHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"schemaId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"linkedAttestationId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"attestTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revokeTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"attester\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"validUntil\",\"type\":\"uint64\"},{\"internalType\":\"enumDataLocation\",\"name\":\"dataLocation\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"revoked\",\"type\":\"bool\"},{\"internalType\":\"bytes[]\",\"name\":\"recipients\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structAttestation\",\"name\":\"attestation\",\"type\":\"tuple\"}],\"name\":\"getDelegatedAttestHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"offchainAttestationIds\",\"type\":\"string[]\"}],\"name\":\"getDelegatedOffchainAttestBatchHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"offchainAttestationId\",\"type\":\"string\"}],\"name\":\"getDelegatedOffchainAttestHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"offchainAttestationIds\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"reasons\",\"type\":\"string[]\"}],\"name\":\"getDelegatedOffchainRevokeBatchHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"offchainAttestationId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"getDelegatedOffchainRevokeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"registrant\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"revocable\",\"type\":\"bool\"},{\"internalType\":\"enumDataLocation\",\"name\":\"dataLocation\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"maxValidFor\",\"type\":\"uint64\"},{\"internalType\":\"contractISPHook\",\"name\":\"hook\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"internalType\":\"structSchema[]\",\"name\":\"schemas\",\"type\":\"tuple[]\"}],\"name\":\"getDelegatedRegisterBatchHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"registrant\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"revocable\",\"type\":\"bool\"},{\"internalType\":\"enumDataLocation\",\"name\":\"dataLocation\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"maxValidFor\",\"type\":\"uint64\"},{\"internalType\":\"contractISPHook\",\"name\":\"hook\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"internalType\":\"structSchema\",\"name\":\"schema\",\"type\":\"tuple\"}],\"name\":\"getDelegatedRegisterHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64[]\",\"name\":\"attestationIds\",\"type\":\"uint64[]\"},{\"internalType\":\"string[]\",\"name\":\"reasons\",\"type\":\"string[]\"}],\"name\":\"getDelegatedRevokeBatchHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"attestationId\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"getDelegatedRevokeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"offchainAttestationId\",\"type\":\"string\"}],\"name\":\"getOffchainAttestation\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"attester\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structOffchainAttestation\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"schemaId\",\"type\":\"uint64\"}],\"name\":\"getSchema\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"registrant\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"revocable\",\"type\":\"bool\"},{\"internalType\":\"enumDataLocation\",\"name\":\"dataLocation\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"maxValidFor\",\"type\":\"uint64\"},{\"internalType\":\"contractISPHook\",\"name\":\"hook\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"internalType\":\"structSchema\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"schemaCounter_\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"attestationCounter_\",\"type\":\"uint64\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"registrant\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"revocable\",\"type\":\"bool\"},{\"internalType\":\"enumDataLocation\",\"name\":\"dataLocation\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"maxValidFor\",\"type\":\"uint64\"},{\"internalType\":\"contractISPHook\",\"name\":\"hook\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"internalType\":\"structSchema\",\"name\":\"schema\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"delegateSignature\",\"type\":\"bytes\"}],\"name\":\"register\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"schemaId\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"registrant\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"revocable\",\"type\":\"bool\"},{\"internalType\":\"enumDataLocation\",\"name\":\"dataLocation\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"maxValidFor\",\"type\":\"uint64\"},{\"internalType\":\"contractISPHook\",\"name\":\"hook\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"internalType\":\"structSchema[]\",\"name\":\"schemas\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"delegateSignature\",\"type\":\"bytes\"}],\"name\":\"registerBatch\",\"outputs\":[{\"internalType\":\"uint64[]\",\"name\":\"schemaIds\",\"type\":\"uint64[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"attestationId\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"delegateSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"revoke\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"attestationId\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"internalType\":\"contractIERC20\",\"name\":\"resolverFeesERC20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"resolverFeesERC20Amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"delegateSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"revoke\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"attestationId\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"resolverFeesETH\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"delegateSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"revoke\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64[]\",\"name\":\"attestationIds\",\"type\":\"uint64[]\"},{\"internalType\":\"string[]\",\"name\":\"reasons\",\"type\":\"string[]\"},{\"internalType\":\"bytes\",\"name\":\"delegateSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"revokeBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64[]\",\"name\":\"attestationIds\",\"type\":\"uint64[]\"},{\"internalType\":\"string[]\",\"name\":\"reasons\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"resolverFeesETH\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"delegateSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"revokeBatch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64[]\",\"name\":\"attestationIds\",\"type\":\"uint64[]\"},{\"internalType\":\"string[]\",\"name\":\"reasons\",\"type\":\"string[]\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"resolverFeesERC20Tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"resolverFeesERC20Amount\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"delegateSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"revokeBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"offchainAttestationId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"delegateSignature\",\"type\":\"bytes\"}],\"name\":\"revokeOffchain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"offchainAttestationIds\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"reasons\",\"type\":\"string[]\"},{\"internalType\":\"bytes\",\"name\":\"delegateSignature\",\"type\":\"bytes\"}],\"name\":\"revokeOffchainBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"schemaCounter\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"hook\",\"type\":\"address\"}],\"name\":\"setGlobalHook\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"}],\"name\":\"setPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Contract *ContractCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Contract *ContractSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Contract.Contract.UPGRADEINTERFACEVERSION(&_Contract.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Contract *ContractCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Contract.Contract.UPGRADEINTERFACEVERSION(&_Contract.CallOpts)
}

// AttestationCounter is a free data retrieval call binding the contract method 0x2c996e03.
//
// Solidity: function attestationCounter() view returns(uint64)
func (_Contract *ContractCaller) AttestationCounter(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "attestationCounter")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// AttestationCounter is a free data retrieval call binding the contract method 0x2c996e03.
//
// Solidity: function attestationCounter() view returns(uint64)
func (_Contract *ContractSession) AttestationCounter() (uint64, error) {
	return _Contract.Contract.AttestationCounter(&_Contract.CallOpts)
}

// AttestationCounter is a free data retrieval call binding the contract method 0x2c996e03.
//
// Solidity: function attestationCounter() view returns(uint64)
func (_Contract *ContractCallerSession) AttestationCounter() (uint64, error) {
	return _Contract.Contract.AttestationCounter(&_Contract.CallOpts)
}

// GetAttestation is a free data retrieval call binding the contract method 0x62252880.
//
// Solidity: function getAttestation(uint64 attestationId) view returns((uint64,uint64,uint64,uint64,address,uint64,uint8,bool,bytes[],bytes))
func (_Contract *ContractCaller) GetAttestation(opts *bind.CallOpts, attestationId uint64) (Attestation, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getAttestation", attestationId)

	if err != nil {
		return *new(Attestation), err
	}

	out0 := *abi.ConvertType(out[0], new(Attestation)).(*Attestation)

	return out0, err

}

// GetAttestation is a free data retrieval call binding the contract method 0x62252880.
//
// Solidity: function getAttestation(uint64 attestationId) view returns((uint64,uint64,uint64,uint64,address,uint64,uint8,bool,bytes[],bytes))
func (_Contract *ContractSession) GetAttestation(attestationId uint64) (Attestation, error) {
	return _Contract.Contract.GetAttestation(&_Contract.CallOpts, attestationId)
}

// GetAttestation is a free data retrieval call binding the contract method 0x62252880.
//
// Solidity: function getAttestation(uint64 attestationId) view returns((uint64,uint64,uint64,uint64,address,uint64,uint8,bool,bytes[],bytes))
func (_Contract *ContractCallerSession) GetAttestation(attestationId uint64) (Attestation, error) {
	return _Contract.Contract.GetAttestation(&_Contract.CallOpts, attestationId)
}

// GetDelegatedAttestBatchHash is a free data retrieval call binding the contract method 0x3f50fb76.
//
// Solidity: function getDelegatedAttestBatchHash((uint64,uint64,uint64,uint64,address,uint64,uint8,bool,bytes[],bytes)[] attestations) pure returns(bytes32)
func (_Contract *ContractCaller) GetDelegatedAttestBatchHash(opts *bind.CallOpts, attestations []Attestation) ([32]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getDelegatedAttestBatchHash", attestations)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDelegatedAttestBatchHash is a free data retrieval call binding the contract method 0x3f50fb76.
//
// Solidity: function getDelegatedAttestBatchHash((uint64,uint64,uint64,uint64,address,uint64,uint8,bool,bytes[],bytes)[] attestations) pure returns(bytes32)
func (_Contract *ContractSession) GetDelegatedAttestBatchHash(attestations []Attestation) ([32]byte, error) {
	return _Contract.Contract.GetDelegatedAttestBatchHash(&_Contract.CallOpts, attestations)
}

// GetDelegatedAttestBatchHash is a free data retrieval call binding the contract method 0x3f50fb76.
//
// Solidity: function getDelegatedAttestBatchHash((uint64,uint64,uint64,uint64,address,uint64,uint8,bool,bytes[],bytes)[] attestations) pure returns(bytes32)
func (_Contract *ContractCallerSession) GetDelegatedAttestBatchHash(attestations []Attestation) ([32]byte, error) {
	return _Contract.Contract.GetDelegatedAttestBatchHash(&_Contract.CallOpts, attestations)
}

// GetDelegatedAttestHash is a free data retrieval call binding the contract method 0x8ef1fbc1.
//
// Solidity: function getDelegatedAttestHash((uint64,uint64,uint64,uint64,address,uint64,uint8,bool,bytes[],bytes) attestation) pure returns(bytes32)
func (_Contract *ContractCaller) GetDelegatedAttestHash(opts *bind.CallOpts, attestation Attestation) ([32]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getDelegatedAttestHash", attestation)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDelegatedAttestHash is a free data retrieval call binding the contract method 0x8ef1fbc1.
//
// Solidity: function getDelegatedAttestHash((uint64,uint64,uint64,uint64,address,uint64,uint8,bool,bytes[],bytes) attestation) pure returns(bytes32)
func (_Contract *ContractSession) GetDelegatedAttestHash(attestation Attestation) ([32]byte, error) {
	return _Contract.Contract.GetDelegatedAttestHash(&_Contract.CallOpts, attestation)
}

// GetDelegatedAttestHash is a free data retrieval call binding the contract method 0x8ef1fbc1.
//
// Solidity: function getDelegatedAttestHash((uint64,uint64,uint64,uint64,address,uint64,uint8,bool,bytes[],bytes) attestation) pure returns(bytes32)
func (_Contract *ContractCallerSession) GetDelegatedAttestHash(attestation Attestation) ([32]byte, error) {
	return _Contract.Contract.GetDelegatedAttestHash(&_Contract.CallOpts, attestation)
}

// GetDelegatedOffchainAttestBatchHash is a free data retrieval call binding the contract method 0x263bc92d.
//
// Solidity: function getDelegatedOffchainAttestBatchHash(string[] offchainAttestationIds) pure returns(bytes32)
func (_Contract *ContractCaller) GetDelegatedOffchainAttestBatchHash(opts *bind.CallOpts, offchainAttestationIds []string) ([32]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getDelegatedOffchainAttestBatchHash", offchainAttestationIds)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDelegatedOffchainAttestBatchHash is a free data retrieval call binding the contract method 0x263bc92d.
//
// Solidity: function getDelegatedOffchainAttestBatchHash(string[] offchainAttestationIds) pure returns(bytes32)
func (_Contract *ContractSession) GetDelegatedOffchainAttestBatchHash(offchainAttestationIds []string) ([32]byte, error) {
	return _Contract.Contract.GetDelegatedOffchainAttestBatchHash(&_Contract.CallOpts, offchainAttestationIds)
}

// GetDelegatedOffchainAttestBatchHash is a free data retrieval call binding the contract method 0x263bc92d.
//
// Solidity: function getDelegatedOffchainAttestBatchHash(string[] offchainAttestationIds) pure returns(bytes32)
func (_Contract *ContractCallerSession) GetDelegatedOffchainAttestBatchHash(offchainAttestationIds []string) ([32]byte, error) {
	return _Contract.Contract.GetDelegatedOffchainAttestBatchHash(&_Contract.CallOpts, offchainAttestationIds)
}

// GetDelegatedOffchainAttestHash is a free data retrieval call binding the contract method 0xbe919fb1.
//
// Solidity: function getDelegatedOffchainAttestHash(string offchainAttestationId) pure returns(bytes32)
func (_Contract *ContractCaller) GetDelegatedOffchainAttestHash(opts *bind.CallOpts, offchainAttestationId string) ([32]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getDelegatedOffchainAttestHash", offchainAttestationId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDelegatedOffchainAttestHash is a free data retrieval call binding the contract method 0xbe919fb1.
//
// Solidity: function getDelegatedOffchainAttestHash(string offchainAttestationId) pure returns(bytes32)
func (_Contract *ContractSession) GetDelegatedOffchainAttestHash(offchainAttestationId string) ([32]byte, error) {
	return _Contract.Contract.GetDelegatedOffchainAttestHash(&_Contract.CallOpts, offchainAttestationId)
}

// GetDelegatedOffchainAttestHash is a free data retrieval call binding the contract method 0xbe919fb1.
//
// Solidity: function getDelegatedOffchainAttestHash(string offchainAttestationId) pure returns(bytes32)
func (_Contract *ContractCallerSession) GetDelegatedOffchainAttestHash(offchainAttestationId string) ([32]byte, error) {
	return _Contract.Contract.GetDelegatedOffchainAttestHash(&_Contract.CallOpts, offchainAttestationId)
}

// GetDelegatedOffchainRevokeBatchHash is a free data retrieval call binding the contract method 0xd3dc7b23.
//
// Solidity: function getDelegatedOffchainRevokeBatchHash(string[] offchainAttestationIds, string[] reasons) pure returns(bytes32)
func (_Contract *ContractCaller) GetDelegatedOffchainRevokeBatchHash(opts *bind.CallOpts, offchainAttestationIds []string, reasons []string) ([32]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getDelegatedOffchainRevokeBatchHash", offchainAttestationIds, reasons)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDelegatedOffchainRevokeBatchHash is a free data retrieval call binding the contract method 0xd3dc7b23.
//
// Solidity: function getDelegatedOffchainRevokeBatchHash(string[] offchainAttestationIds, string[] reasons) pure returns(bytes32)
func (_Contract *ContractSession) GetDelegatedOffchainRevokeBatchHash(offchainAttestationIds []string, reasons []string) ([32]byte, error) {
	return _Contract.Contract.GetDelegatedOffchainRevokeBatchHash(&_Contract.CallOpts, offchainAttestationIds, reasons)
}

// GetDelegatedOffchainRevokeBatchHash is a free data retrieval call binding the contract method 0xd3dc7b23.
//
// Solidity: function getDelegatedOffchainRevokeBatchHash(string[] offchainAttestationIds, string[] reasons) pure returns(bytes32)
func (_Contract *ContractCallerSession) GetDelegatedOffchainRevokeBatchHash(offchainAttestationIds []string, reasons []string) ([32]byte, error) {
	return _Contract.Contract.GetDelegatedOffchainRevokeBatchHash(&_Contract.CallOpts, offchainAttestationIds, reasons)
}

// GetDelegatedOffchainRevokeHash is a free data retrieval call binding the contract method 0x985292eb.
//
// Solidity: function getDelegatedOffchainRevokeHash(string offchainAttestationId, string reason) pure returns(bytes32)
func (_Contract *ContractCaller) GetDelegatedOffchainRevokeHash(opts *bind.CallOpts, offchainAttestationId string, reason string) ([32]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getDelegatedOffchainRevokeHash", offchainAttestationId, reason)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDelegatedOffchainRevokeHash is a free data retrieval call binding the contract method 0x985292eb.
//
// Solidity: function getDelegatedOffchainRevokeHash(string offchainAttestationId, string reason) pure returns(bytes32)
func (_Contract *ContractSession) GetDelegatedOffchainRevokeHash(offchainAttestationId string, reason string) ([32]byte, error) {
	return _Contract.Contract.GetDelegatedOffchainRevokeHash(&_Contract.CallOpts, offchainAttestationId, reason)
}

// GetDelegatedOffchainRevokeHash is a free data retrieval call binding the contract method 0x985292eb.
//
// Solidity: function getDelegatedOffchainRevokeHash(string offchainAttestationId, string reason) pure returns(bytes32)
func (_Contract *ContractCallerSession) GetDelegatedOffchainRevokeHash(offchainAttestationId string, reason string) ([32]byte, error) {
	return _Contract.Contract.GetDelegatedOffchainRevokeHash(&_Contract.CallOpts, offchainAttestationId, reason)
}

// GetDelegatedRegisterBatchHash is a free data retrieval call binding the contract method 0x00f0f855.
//
// Solidity: function getDelegatedRegisterBatchHash((address,bool,uint8,uint64,address,uint64,string)[] schemas) pure returns(bytes32)
func (_Contract *ContractCaller) GetDelegatedRegisterBatchHash(opts *bind.CallOpts, schemas []Schema) ([32]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getDelegatedRegisterBatchHash", schemas)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDelegatedRegisterBatchHash is a free data retrieval call binding the contract method 0x00f0f855.
//
// Solidity: function getDelegatedRegisterBatchHash((address,bool,uint8,uint64,address,uint64,string)[] schemas) pure returns(bytes32)
func (_Contract *ContractSession) GetDelegatedRegisterBatchHash(schemas []Schema) ([32]byte, error) {
	return _Contract.Contract.GetDelegatedRegisterBatchHash(&_Contract.CallOpts, schemas)
}

// GetDelegatedRegisterBatchHash is a free data retrieval call binding the contract method 0x00f0f855.
//
// Solidity: function getDelegatedRegisterBatchHash((address,bool,uint8,uint64,address,uint64,string)[] schemas) pure returns(bytes32)
func (_Contract *ContractCallerSession) GetDelegatedRegisterBatchHash(schemas []Schema) ([32]byte, error) {
	return _Contract.Contract.GetDelegatedRegisterBatchHash(&_Contract.CallOpts, schemas)
}

// GetDelegatedRegisterHash is a free data retrieval call binding the contract method 0x570b7434.
//
// Solidity: function getDelegatedRegisterHash((address,bool,uint8,uint64,address,uint64,string) schema) pure returns(bytes32)
func (_Contract *ContractCaller) GetDelegatedRegisterHash(opts *bind.CallOpts, schema Schema) ([32]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getDelegatedRegisterHash", schema)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDelegatedRegisterHash is a free data retrieval call binding the contract method 0x570b7434.
//
// Solidity: function getDelegatedRegisterHash((address,bool,uint8,uint64,address,uint64,string) schema) pure returns(bytes32)
func (_Contract *ContractSession) GetDelegatedRegisterHash(schema Schema) ([32]byte, error) {
	return _Contract.Contract.GetDelegatedRegisterHash(&_Contract.CallOpts, schema)
}

// GetDelegatedRegisterHash is a free data retrieval call binding the contract method 0x570b7434.
//
// Solidity: function getDelegatedRegisterHash((address,bool,uint8,uint64,address,uint64,string) schema) pure returns(bytes32)
func (_Contract *ContractCallerSession) GetDelegatedRegisterHash(schema Schema) ([32]byte, error) {
	return _Contract.Contract.GetDelegatedRegisterHash(&_Contract.CallOpts, schema)
}

// GetDelegatedRevokeBatchHash is a free data retrieval call binding the contract method 0x40f356fa.
//
// Solidity: function getDelegatedRevokeBatchHash(uint64[] attestationIds, string[] reasons) pure returns(bytes32)
func (_Contract *ContractCaller) GetDelegatedRevokeBatchHash(opts *bind.CallOpts, attestationIds []uint64, reasons []string) ([32]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getDelegatedRevokeBatchHash", attestationIds, reasons)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDelegatedRevokeBatchHash is a free data retrieval call binding the contract method 0x40f356fa.
//
// Solidity: function getDelegatedRevokeBatchHash(uint64[] attestationIds, string[] reasons) pure returns(bytes32)
func (_Contract *ContractSession) GetDelegatedRevokeBatchHash(attestationIds []uint64, reasons []string) ([32]byte, error) {
	return _Contract.Contract.GetDelegatedRevokeBatchHash(&_Contract.CallOpts, attestationIds, reasons)
}

// GetDelegatedRevokeBatchHash is a free data retrieval call binding the contract method 0x40f356fa.
//
// Solidity: function getDelegatedRevokeBatchHash(uint64[] attestationIds, string[] reasons) pure returns(bytes32)
func (_Contract *ContractCallerSession) GetDelegatedRevokeBatchHash(attestationIds []uint64, reasons []string) ([32]byte, error) {
	return _Contract.Contract.GetDelegatedRevokeBatchHash(&_Contract.CallOpts, attestationIds, reasons)
}

// GetDelegatedRevokeHash is a free data retrieval call binding the contract method 0xdc50eada.
//
// Solidity: function getDelegatedRevokeHash(uint64 attestationId, string reason) pure returns(bytes32)
func (_Contract *ContractCaller) GetDelegatedRevokeHash(opts *bind.CallOpts, attestationId uint64, reason string) ([32]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getDelegatedRevokeHash", attestationId, reason)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDelegatedRevokeHash is a free data retrieval call binding the contract method 0xdc50eada.
//
// Solidity: function getDelegatedRevokeHash(uint64 attestationId, string reason) pure returns(bytes32)
func (_Contract *ContractSession) GetDelegatedRevokeHash(attestationId uint64, reason string) ([32]byte, error) {
	return _Contract.Contract.GetDelegatedRevokeHash(&_Contract.CallOpts, attestationId, reason)
}

// GetDelegatedRevokeHash is a free data retrieval call binding the contract method 0xdc50eada.
//
// Solidity: function getDelegatedRevokeHash(uint64 attestationId, string reason) pure returns(bytes32)
func (_Contract *ContractCallerSession) GetDelegatedRevokeHash(attestationId uint64, reason string) ([32]byte, error) {
	return _Contract.Contract.GetDelegatedRevokeHash(&_Contract.CallOpts, attestationId, reason)
}

// GetOffchainAttestation is a free data retrieval call binding the contract method 0x6a67651c.
//
// Solidity: function getOffchainAttestation(string offchainAttestationId) view returns((address,uint64))
func (_Contract *ContractCaller) GetOffchainAttestation(opts *bind.CallOpts, offchainAttestationId string) (OffchainAttestation, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getOffchainAttestation", offchainAttestationId)

	if err != nil {
		return *new(OffchainAttestation), err
	}

	out0 := *abi.ConvertType(out[0], new(OffchainAttestation)).(*OffchainAttestation)

	return out0, err

}

// GetOffchainAttestation is a free data retrieval call binding the contract method 0x6a67651c.
//
// Solidity: function getOffchainAttestation(string offchainAttestationId) view returns((address,uint64))
func (_Contract *ContractSession) GetOffchainAttestation(offchainAttestationId string) (OffchainAttestation, error) {
	return _Contract.Contract.GetOffchainAttestation(&_Contract.CallOpts, offchainAttestationId)
}

// GetOffchainAttestation is a free data retrieval call binding the contract method 0x6a67651c.
//
// Solidity: function getOffchainAttestation(string offchainAttestationId) view returns((address,uint64))
func (_Contract *ContractCallerSession) GetOffchainAttestation(offchainAttestationId string) (OffchainAttestation, error) {
	return _Contract.Contract.GetOffchainAttestation(&_Contract.CallOpts, offchainAttestationId)
}

// GetSchema is a free data retrieval call binding the contract method 0x8d1be875.
//
// Solidity: function getSchema(uint64 schemaId) view returns((address,bool,uint8,uint64,address,uint64,string))
func (_Contract *ContractCaller) GetSchema(opts *bind.CallOpts, schemaId uint64) (Schema, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getSchema", schemaId)

	if err != nil {
		return *new(Schema), err
	}

	out0 := *abi.ConvertType(out[0], new(Schema)).(*Schema)

	return out0, err

}

// GetSchema is a free data retrieval call binding the contract method 0x8d1be875.
//
// Solidity: function getSchema(uint64 schemaId) view returns((address,bool,uint8,uint64,address,uint64,string))
func (_Contract *ContractSession) GetSchema(schemaId uint64) (Schema, error) {
	return _Contract.Contract.GetSchema(&_Contract.CallOpts, schemaId)
}

// GetSchema is a free data retrieval call binding the contract method 0x8d1be875.
//
// Solidity: function getSchema(uint64 schemaId) view returns((address,bool,uint8,uint64,address,uint64,string))
func (_Contract *ContractCallerSession) GetSchema(schemaId uint64) (Schema, error) {
	return _Contract.Contract.GetSchema(&_Contract.CallOpts, schemaId)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCallerSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Contract *ContractCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Contract *ContractSession) ProxiableUUID() ([32]byte, error) {
	return _Contract.Contract.ProxiableUUID(&_Contract.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Contract *ContractCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Contract.Contract.ProxiableUUID(&_Contract.CallOpts)
}

// SchemaCounter is a free data retrieval call binding the contract method 0xcb6a3237.
//
// Solidity: function schemaCounter() view returns(uint64)
func (_Contract *ContractCaller) SchemaCounter(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "schemaCounter")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// SchemaCounter is a free data retrieval call binding the contract method 0xcb6a3237.
//
// Solidity: function schemaCounter() view returns(uint64)
func (_Contract *ContractSession) SchemaCounter() (uint64, error) {
	return _Contract.Contract.SchemaCounter(&_Contract.CallOpts)
}

// SchemaCounter is a free data retrieval call binding the contract method 0xcb6a3237.
//
// Solidity: function schemaCounter() view returns(uint64)
func (_Contract *ContractCallerSession) SchemaCounter() (uint64, error) {
	return _Contract.Contract.SchemaCounter(&_Contract.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Contract *ContractCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Contract *ContractSession) Version() (string, error) {
	return _Contract.Contract.Version(&_Contract.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Contract *ContractCallerSession) Version() (string, error) {
	return _Contract.Contract.Version(&_Contract.CallOpts)
}

// Attest is a paid mutator transaction binding the contract method 0x812c255f.
//
// Solidity: function attest((uint64,uint64,uint64,uint64,address,uint64,uint8,bool,bytes[],bytes) attestation, address resolverFeesERC20Token, uint256 resolverFeesERC20Amount, string indexingKey, bytes delegateSignature, bytes extraData) returns(uint64)
func (_Contract *ContractTransactor) Attest(opts *bind.TransactOpts, attestation Attestation, resolverFeesERC20Token common.Address, resolverFeesERC20Amount *big.Int, indexingKey string, delegateSignature []byte, extraData []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "attest", attestation, resolverFeesERC20Token, resolverFeesERC20Amount, indexingKey, delegateSignature, extraData)
}

// Attest is a paid mutator transaction binding the contract method 0x812c255f.
//
// Solidity: function attest((uint64,uint64,uint64,uint64,address,uint64,uint8,bool,bytes[],bytes) attestation, address resolverFeesERC20Token, uint256 resolverFeesERC20Amount, string indexingKey, bytes delegateSignature, bytes extraData) returns(uint64)
func (_Contract *ContractSession) Attest(attestation Attestation, resolverFeesERC20Token common.Address, resolverFeesERC20Amount *big.Int, indexingKey string, delegateSignature []byte, extraData []byte) (*types.Transaction, error) {
	return _Contract.Contract.Attest(&_Contract.TransactOpts, attestation, resolverFeesERC20Token, resolverFeesERC20Amount, indexingKey, delegateSignature, extraData)
}

// Attest is a paid mutator transaction binding the contract method 0x812c255f.
//
// Solidity: function attest((uint64,uint64,uint64,uint64,address,uint64,uint8,bool,bytes[],bytes) attestation, address resolverFeesERC20Token, uint256 resolverFeesERC20Amount, string indexingKey, bytes delegateSignature, bytes extraData) returns(uint64)
func (_Contract *ContractTransactorSession) Attest(attestation Attestation, resolverFeesERC20Token common.Address, resolverFeesERC20Amount *big.Int, indexingKey string, delegateSignature []byte, extraData []byte) (*types.Transaction, error) {
	return _Contract.Contract.Attest(&_Contract.TransactOpts, attestation, resolverFeesERC20Token, resolverFeesERC20Amount, indexingKey, delegateSignature, extraData)
}

// Attest0 is a paid mutator transaction binding the contract method 0xb82916cb.
//
// Solidity: function attest((uint64,uint64,uint64,uint64,address,uint64,uint8,bool,bytes[],bytes) attestation, string indexingKey, bytes delegateSignature, bytes extraData) returns(uint64)
func (_Contract *ContractTransactor) Attest0(opts *bind.TransactOpts, attestation Attestation, indexingKey string, delegateSignature []byte, extraData []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "attest0", attestation, indexingKey, delegateSignature, extraData)
}

// Attest0 is a paid mutator transaction binding the contract method 0xb82916cb.
//
// Solidity: function attest((uint64,uint64,uint64,uint64,address,uint64,uint8,bool,bytes[],bytes) attestation, string indexingKey, bytes delegateSignature, bytes extraData) returns(uint64)
func (_Contract *ContractSession) Attest0(attestation Attestation, indexingKey string, delegateSignature []byte, extraData []byte) (*types.Transaction, error) {
	return _Contract.Contract.Attest0(&_Contract.TransactOpts, attestation, indexingKey, delegateSignature, extraData)
}

// Attest0 is a paid mutator transaction binding the contract method 0xb82916cb.
//
// Solidity: function attest((uint64,uint64,uint64,uint64,address,uint64,uint8,bool,bytes[],bytes) attestation, string indexingKey, bytes delegateSignature, bytes extraData) returns(uint64)
func (_Contract *ContractTransactorSession) Attest0(attestation Attestation, indexingKey string, delegateSignature []byte, extraData []byte) (*types.Transaction, error) {
	return _Contract.Contract.Attest0(&_Contract.TransactOpts, attestation, indexingKey, delegateSignature, extraData)
}

// Attest1 is a paid mutator transaction binding the contract method 0xe09f1b7a.
//
// Solidity: function attest((uint64,uint64,uint64,uint64,address,uint64,uint8,bool,bytes[],bytes) attestation, uint256 resolverFeesETH, string indexingKey, bytes delegateSignature, bytes extraData) payable returns(uint64)
func (_Contract *ContractTransactor) Attest1(opts *bind.TransactOpts, attestation Attestation, resolverFeesETH *big.Int, indexingKey string, delegateSignature []byte, extraData []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "attest1", attestation, resolverFeesETH, indexingKey, delegateSignature, extraData)
}

// Attest1 is a paid mutator transaction binding the contract method 0xe09f1b7a.
//
// Solidity: function attest((uint64,uint64,uint64,uint64,address,uint64,uint8,bool,bytes[],bytes) attestation, uint256 resolverFeesETH, string indexingKey, bytes delegateSignature, bytes extraData) payable returns(uint64)
func (_Contract *ContractSession) Attest1(attestation Attestation, resolverFeesETH *big.Int, indexingKey string, delegateSignature []byte, extraData []byte) (*types.Transaction, error) {
	return _Contract.Contract.Attest1(&_Contract.TransactOpts, attestation, resolverFeesETH, indexingKey, delegateSignature, extraData)
}

// Attest1 is a paid mutator transaction binding the contract method 0xe09f1b7a.
//
// Solidity: function attest((uint64,uint64,uint64,uint64,address,uint64,uint8,bool,bytes[],bytes) attestation, uint256 resolverFeesETH, string indexingKey, bytes delegateSignature, bytes extraData) payable returns(uint64)
func (_Contract *ContractTransactorSession) Attest1(attestation Attestation, resolverFeesETH *big.Int, indexingKey string, delegateSignature []byte, extraData []byte) (*types.Transaction, error) {
	return _Contract.Contract.Attest1(&_Contract.TransactOpts, attestation, resolverFeesETH, indexingKey, delegateSignature, extraData)
}

// AttestBatch is a paid mutator transaction binding the contract method 0x4e6ac1d5.
//
// Solidity: function attestBatch((uint64,uint64,uint64,uint64,address,uint64,uint8,bool,bytes[],bytes)[] attestations, uint256[] resolverFeesETH, string[] indexingKeys, bytes delegateSignature, bytes extraData) payable returns(uint64[] attestationIds)
func (_Contract *ContractTransactor) AttestBatch(opts *bind.TransactOpts, attestations []Attestation, resolverFeesETH []*big.Int, indexingKeys []string, delegateSignature []byte, extraData []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "attestBatch", attestations, resolverFeesETH, indexingKeys, delegateSignature, extraData)
}

// AttestBatch is a paid mutator transaction binding the contract method 0x4e6ac1d5.
//
// Solidity: function attestBatch((uint64,uint64,uint64,uint64,address,uint64,uint8,bool,bytes[],bytes)[] attestations, uint256[] resolverFeesETH, string[] indexingKeys, bytes delegateSignature, bytes extraData) payable returns(uint64[] attestationIds)
func (_Contract *ContractSession) AttestBatch(attestations []Attestation, resolverFeesETH []*big.Int, indexingKeys []string, delegateSignature []byte, extraData []byte) (*types.Transaction, error) {
	return _Contract.Contract.AttestBatch(&_Contract.TransactOpts, attestations, resolverFeesETH, indexingKeys, delegateSignature, extraData)
}

// AttestBatch is a paid mutator transaction binding the contract method 0x4e6ac1d5.
//
// Solidity: function attestBatch((uint64,uint64,uint64,uint64,address,uint64,uint8,bool,bytes[],bytes)[] attestations, uint256[] resolverFeesETH, string[] indexingKeys, bytes delegateSignature, bytes extraData) payable returns(uint64[] attestationIds)
func (_Contract *ContractTransactorSession) AttestBatch(attestations []Attestation, resolverFeesETH []*big.Int, indexingKeys []string, delegateSignature []byte, extraData []byte) (*types.Transaction, error) {
	return _Contract.Contract.AttestBatch(&_Contract.TransactOpts, attestations, resolverFeesETH, indexingKeys, delegateSignature, extraData)
}

// AttestBatch0 is a paid mutator transaction binding the contract method 0x59d9a12f.
//
// Solidity: function attestBatch((uint64,uint64,uint64,uint64,address,uint64,uint8,bool,bytes[],bytes)[] attestations, address[] resolverFeesERC20Tokens, uint256[] resolverFeesERC20Amount, string[] indexingKeys, bytes delegateSignature, bytes extraData) returns(uint64[] attestationIds)
func (_Contract *ContractTransactor) AttestBatch0(opts *bind.TransactOpts, attestations []Attestation, resolverFeesERC20Tokens []common.Address, resolverFeesERC20Amount []*big.Int, indexingKeys []string, delegateSignature []byte, extraData []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "attestBatch0", attestations, resolverFeesERC20Tokens, resolverFeesERC20Amount, indexingKeys, delegateSignature, extraData)
}

// AttestBatch0 is a paid mutator transaction binding the contract method 0x59d9a12f.
//
// Solidity: function attestBatch((uint64,uint64,uint64,uint64,address,uint64,uint8,bool,bytes[],bytes)[] attestations, address[] resolverFeesERC20Tokens, uint256[] resolverFeesERC20Amount, string[] indexingKeys, bytes delegateSignature, bytes extraData) returns(uint64[] attestationIds)
func (_Contract *ContractSession) AttestBatch0(attestations []Attestation, resolverFeesERC20Tokens []common.Address, resolverFeesERC20Amount []*big.Int, indexingKeys []string, delegateSignature []byte, extraData []byte) (*types.Transaction, error) {
	return _Contract.Contract.AttestBatch0(&_Contract.TransactOpts, attestations, resolverFeesERC20Tokens, resolverFeesERC20Amount, indexingKeys, delegateSignature, extraData)
}

// AttestBatch0 is a paid mutator transaction binding the contract method 0x59d9a12f.
//
// Solidity: function attestBatch((uint64,uint64,uint64,uint64,address,uint64,uint8,bool,bytes[],bytes)[] attestations, address[] resolverFeesERC20Tokens, uint256[] resolverFeesERC20Amount, string[] indexingKeys, bytes delegateSignature, bytes extraData) returns(uint64[] attestationIds)
func (_Contract *ContractTransactorSession) AttestBatch0(attestations []Attestation, resolverFeesERC20Tokens []common.Address, resolverFeesERC20Amount []*big.Int, indexingKeys []string, delegateSignature []byte, extraData []byte) (*types.Transaction, error) {
	return _Contract.Contract.AttestBatch0(&_Contract.TransactOpts, attestations, resolverFeesERC20Tokens, resolverFeesERC20Amount, indexingKeys, delegateSignature, extraData)
}

// AttestBatch1 is a paid mutator transaction binding the contract method 0xd98934ff.
//
// Solidity: function attestBatch((uint64,uint64,uint64,uint64,address,uint64,uint8,bool,bytes[],bytes)[] attestations, string[] indexingKeys, bytes delegateSignature, bytes extraData) returns(uint64[] attestationIds)
func (_Contract *ContractTransactor) AttestBatch1(opts *bind.TransactOpts, attestations []Attestation, indexingKeys []string, delegateSignature []byte, extraData []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "attestBatch1", attestations, indexingKeys, delegateSignature, extraData)
}

// AttestBatch1 is a paid mutator transaction binding the contract method 0xd98934ff.
//
// Solidity: function attestBatch((uint64,uint64,uint64,uint64,address,uint64,uint8,bool,bytes[],bytes)[] attestations, string[] indexingKeys, bytes delegateSignature, bytes extraData) returns(uint64[] attestationIds)
func (_Contract *ContractSession) AttestBatch1(attestations []Attestation, indexingKeys []string, delegateSignature []byte, extraData []byte) (*types.Transaction, error) {
	return _Contract.Contract.AttestBatch1(&_Contract.TransactOpts, attestations, indexingKeys, delegateSignature, extraData)
}

// AttestBatch1 is a paid mutator transaction binding the contract method 0xd98934ff.
//
// Solidity: function attestBatch((uint64,uint64,uint64,uint64,address,uint64,uint8,bool,bytes[],bytes)[] attestations, string[] indexingKeys, bytes delegateSignature, bytes extraData) returns(uint64[] attestationIds)
func (_Contract *ContractTransactorSession) AttestBatch1(attestations []Attestation, indexingKeys []string, delegateSignature []byte, extraData []byte) (*types.Transaction, error) {
	return _Contract.Contract.AttestBatch1(&_Contract.TransactOpts, attestations, indexingKeys, delegateSignature, extraData)
}

// AttestOffchain is a paid mutator transaction binding the contract method 0xa01fb774.
//
// Solidity: function attestOffchain(string offchainAttestationId, address delegateAttester, bytes delegateSignature) returns()
func (_Contract *ContractTransactor) AttestOffchain(opts *bind.TransactOpts, offchainAttestationId string, delegateAttester common.Address, delegateSignature []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "attestOffchain", offchainAttestationId, delegateAttester, delegateSignature)
}

// AttestOffchain is a paid mutator transaction binding the contract method 0xa01fb774.
//
// Solidity: function attestOffchain(string offchainAttestationId, address delegateAttester, bytes delegateSignature) returns()
func (_Contract *ContractSession) AttestOffchain(offchainAttestationId string, delegateAttester common.Address, delegateSignature []byte) (*types.Transaction, error) {
	return _Contract.Contract.AttestOffchain(&_Contract.TransactOpts, offchainAttestationId, delegateAttester, delegateSignature)
}

// AttestOffchain is a paid mutator transaction binding the contract method 0xa01fb774.
//
// Solidity: function attestOffchain(string offchainAttestationId, address delegateAttester, bytes delegateSignature) returns()
func (_Contract *ContractTransactorSession) AttestOffchain(offchainAttestationId string, delegateAttester common.Address, delegateSignature []byte) (*types.Transaction, error) {
	return _Contract.Contract.AttestOffchain(&_Contract.TransactOpts, offchainAttestationId, delegateAttester, delegateSignature)
}

// AttestOffchainBatch is a paid mutator transaction binding the contract method 0x1916749d.
//
// Solidity: function attestOffchainBatch(string[] attestationIds, address delegateAttester, bytes delegateSignature) returns()
func (_Contract *ContractTransactor) AttestOffchainBatch(opts *bind.TransactOpts, attestationIds []string, delegateAttester common.Address, delegateSignature []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "attestOffchainBatch", attestationIds, delegateAttester, delegateSignature)
}

// AttestOffchainBatch is a paid mutator transaction binding the contract method 0x1916749d.
//
// Solidity: function attestOffchainBatch(string[] attestationIds, address delegateAttester, bytes delegateSignature) returns()
func (_Contract *ContractSession) AttestOffchainBatch(attestationIds []string, delegateAttester common.Address, delegateSignature []byte) (*types.Transaction, error) {
	return _Contract.Contract.AttestOffchainBatch(&_Contract.TransactOpts, attestationIds, delegateAttester, delegateSignature)
}

// AttestOffchainBatch is a paid mutator transaction binding the contract method 0x1916749d.
//
// Solidity: function attestOffchainBatch(string[] attestationIds, address delegateAttester, bytes delegateSignature) returns()
func (_Contract *ContractTransactorSession) AttestOffchainBatch(attestationIds []string, delegateAttester common.Address, delegateSignature []byte) (*types.Transaction, error) {
	return _Contract.Contract.AttestOffchainBatch(&_Contract.TransactOpts, attestationIds, delegateAttester, delegateSignature)
}

// Initialize is a paid mutator transaction binding the contract method 0x20f2f345.
//
// Solidity: function initialize(uint64 schemaCounter_, uint64 attestationCounter_) returns()
func (_Contract *ContractTransactor) Initialize(opts *bind.TransactOpts, schemaCounter_ uint64, attestationCounter_ uint64) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "initialize", schemaCounter_, attestationCounter_)
}

// Initialize is a paid mutator transaction binding the contract method 0x20f2f345.
//
// Solidity: function initialize(uint64 schemaCounter_, uint64 attestationCounter_) returns()
func (_Contract *ContractSession) Initialize(schemaCounter_ uint64, attestationCounter_ uint64) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, schemaCounter_, attestationCounter_)
}

// Initialize is a paid mutator transaction binding the contract method 0x20f2f345.
//
// Solidity: function initialize(uint64 schemaCounter_, uint64 attestationCounter_) returns()
func (_Contract *ContractTransactorSession) Initialize(schemaCounter_ uint64, attestationCounter_ uint64) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, schemaCounter_, attestationCounter_)
}

// Register is a paid mutator transaction binding the contract method 0xe117a861.
//
// Solidity: function register((address,bool,uint8,uint64,address,uint64,string) schema, bytes delegateSignature) returns(uint64 schemaId)
func (_Contract *ContractTransactor) Register(opts *bind.TransactOpts, schema Schema, delegateSignature []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "register", schema, delegateSignature)
}

// Register is a paid mutator transaction binding the contract method 0xe117a861.
//
// Solidity: function register((address,bool,uint8,uint64,address,uint64,string) schema, bytes delegateSignature) returns(uint64 schemaId)
func (_Contract *ContractSession) Register(schema Schema, delegateSignature []byte) (*types.Transaction, error) {
	return _Contract.Contract.Register(&_Contract.TransactOpts, schema, delegateSignature)
}

// Register is a paid mutator transaction binding the contract method 0xe117a861.
//
// Solidity: function register((address,bool,uint8,uint64,address,uint64,string) schema, bytes delegateSignature) returns(uint64 schemaId)
func (_Contract *ContractTransactorSession) Register(schema Schema, delegateSignature []byte) (*types.Transaction, error) {
	return _Contract.Contract.Register(&_Contract.TransactOpts, schema, delegateSignature)
}

// RegisterBatch is a paid mutator transaction binding the contract method 0x5e7689b7.
//
// Solidity: function registerBatch((address,bool,uint8,uint64,address,uint64,string)[] schemas, bytes delegateSignature) returns(uint64[] schemaIds)
func (_Contract *ContractTransactor) RegisterBatch(opts *bind.TransactOpts, schemas []Schema, delegateSignature []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "registerBatch", schemas, delegateSignature)
}

// RegisterBatch is a paid mutator transaction binding the contract method 0x5e7689b7.
//
// Solidity: function registerBatch((address,bool,uint8,uint64,address,uint64,string)[] schemas, bytes delegateSignature) returns(uint64[] schemaIds)
func (_Contract *ContractSession) RegisterBatch(schemas []Schema, delegateSignature []byte) (*types.Transaction, error) {
	return _Contract.Contract.RegisterBatch(&_Contract.TransactOpts, schemas, delegateSignature)
}

// RegisterBatch is a paid mutator transaction binding the contract method 0x5e7689b7.
//
// Solidity: function registerBatch((address,bool,uint8,uint64,address,uint64,string)[] schemas, bytes delegateSignature) returns(uint64[] schemaIds)
func (_Contract *ContractTransactorSession) RegisterBatch(schemas []Schema, delegateSignature []byte) (*types.Transaction, error) {
	return _Contract.Contract.RegisterBatch(&_Contract.TransactOpts, schemas, delegateSignature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

// Revoke is a paid mutator transaction binding the contract method 0x24855e60.
//
// Solidity: function revoke(uint64 attestationId, string reason, bytes delegateSignature, bytes extraData) returns()
func (_Contract *ContractTransactor) Revoke(opts *bind.TransactOpts, attestationId uint64, reason string, delegateSignature []byte, extraData []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "revoke", attestationId, reason, delegateSignature, extraData)
}

// Revoke is a paid mutator transaction binding the contract method 0x24855e60.
//
// Solidity: function revoke(uint64 attestationId, string reason, bytes delegateSignature, bytes extraData) returns()
func (_Contract *ContractSession) Revoke(attestationId uint64, reason string, delegateSignature []byte, extraData []byte) (*types.Transaction, error) {
	return _Contract.Contract.Revoke(&_Contract.TransactOpts, attestationId, reason, delegateSignature, extraData)
}

// Revoke is a paid mutator transaction binding the contract method 0x24855e60.
//
// Solidity: function revoke(uint64 attestationId, string reason, bytes delegateSignature, bytes extraData) returns()
func (_Contract *ContractTransactorSession) Revoke(attestationId uint64, reason string, delegateSignature []byte, extraData []byte) (*types.Transaction, error) {
	return _Contract.Contract.Revoke(&_Contract.TransactOpts, attestationId, reason, delegateSignature, extraData)
}

// Revoke0 is a paid mutator transaction binding the contract method 0x9b7f8318.
//
// Solidity: function revoke(uint64 attestationId, string reason, address resolverFeesERC20Token, uint256 resolverFeesERC20Amount, bytes delegateSignature, bytes extraData) returns()
func (_Contract *ContractTransactor) Revoke0(opts *bind.TransactOpts, attestationId uint64, reason string, resolverFeesERC20Token common.Address, resolverFeesERC20Amount *big.Int, delegateSignature []byte, extraData []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "revoke0", attestationId, reason, resolverFeesERC20Token, resolverFeesERC20Amount, delegateSignature, extraData)
}

// Revoke0 is a paid mutator transaction binding the contract method 0x9b7f8318.
//
// Solidity: function revoke(uint64 attestationId, string reason, address resolverFeesERC20Token, uint256 resolverFeesERC20Amount, bytes delegateSignature, bytes extraData) returns()
func (_Contract *ContractSession) Revoke0(attestationId uint64, reason string, resolverFeesERC20Token common.Address, resolverFeesERC20Amount *big.Int, delegateSignature []byte, extraData []byte) (*types.Transaction, error) {
	return _Contract.Contract.Revoke0(&_Contract.TransactOpts, attestationId, reason, resolverFeesERC20Token, resolverFeesERC20Amount, delegateSignature, extraData)
}

// Revoke0 is a paid mutator transaction binding the contract method 0x9b7f8318.
//
// Solidity: function revoke(uint64 attestationId, string reason, address resolverFeesERC20Token, uint256 resolverFeesERC20Amount, bytes delegateSignature, bytes extraData) returns()
func (_Contract *ContractTransactorSession) Revoke0(attestationId uint64, reason string, resolverFeesERC20Token common.Address, resolverFeesERC20Amount *big.Int, delegateSignature []byte, extraData []byte) (*types.Transaction, error) {
	return _Contract.Contract.Revoke0(&_Contract.TransactOpts, attestationId, reason, resolverFeesERC20Token, resolverFeesERC20Amount, delegateSignature, extraData)
}

// Revoke1 is a paid mutator transaction binding the contract method 0xf83b92b7.
//
// Solidity: function revoke(uint64 attestationId, string reason, uint256 resolverFeesETH, bytes delegateSignature, bytes extraData) payable returns()
func (_Contract *ContractTransactor) Revoke1(opts *bind.TransactOpts, attestationId uint64, reason string, resolverFeesETH *big.Int, delegateSignature []byte, extraData []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "revoke1", attestationId, reason, resolverFeesETH, delegateSignature, extraData)
}

// Revoke1 is a paid mutator transaction binding the contract method 0xf83b92b7.
//
// Solidity: function revoke(uint64 attestationId, string reason, uint256 resolverFeesETH, bytes delegateSignature, bytes extraData) payable returns()
func (_Contract *ContractSession) Revoke1(attestationId uint64, reason string, resolverFeesETH *big.Int, delegateSignature []byte, extraData []byte) (*types.Transaction, error) {
	return _Contract.Contract.Revoke1(&_Contract.TransactOpts, attestationId, reason, resolverFeesETH, delegateSignature, extraData)
}

// Revoke1 is a paid mutator transaction binding the contract method 0xf83b92b7.
//
// Solidity: function revoke(uint64 attestationId, string reason, uint256 resolverFeesETH, bytes delegateSignature, bytes extraData) payable returns()
func (_Contract *ContractTransactorSession) Revoke1(attestationId uint64, reason string, resolverFeesETH *big.Int, delegateSignature []byte, extraData []byte) (*types.Transaction, error) {
	return _Contract.Contract.Revoke1(&_Contract.TransactOpts, attestationId, reason, resolverFeesETH, delegateSignature, extraData)
}

// RevokeBatch is a paid mutator transaction binding the contract method 0x65ae13b3.
//
// Solidity: function revokeBatch(uint64[] attestationIds, string[] reasons, bytes delegateSignature, bytes extraData) returns()
func (_Contract *ContractTransactor) RevokeBatch(opts *bind.TransactOpts, attestationIds []uint64, reasons []string, delegateSignature []byte, extraData []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "revokeBatch", attestationIds, reasons, delegateSignature, extraData)
}

// RevokeBatch is a paid mutator transaction binding the contract method 0x65ae13b3.
//
// Solidity: function revokeBatch(uint64[] attestationIds, string[] reasons, bytes delegateSignature, bytes extraData) returns()
func (_Contract *ContractSession) RevokeBatch(attestationIds []uint64, reasons []string, delegateSignature []byte, extraData []byte) (*types.Transaction, error) {
	return _Contract.Contract.RevokeBatch(&_Contract.TransactOpts, attestationIds, reasons, delegateSignature, extraData)
}

// RevokeBatch is a paid mutator transaction binding the contract method 0x65ae13b3.
//
// Solidity: function revokeBatch(uint64[] attestationIds, string[] reasons, bytes delegateSignature, bytes extraData) returns()
func (_Contract *ContractTransactorSession) RevokeBatch(attestationIds []uint64, reasons []string, delegateSignature []byte, extraData []byte) (*types.Transaction, error) {
	return _Contract.Contract.RevokeBatch(&_Contract.TransactOpts, attestationIds, reasons, delegateSignature, extraData)
}

// RevokeBatch0 is a paid mutator transaction binding the contract method 0xd7681235.
//
// Solidity: function revokeBatch(uint64[] attestationIds, string[] reasons, uint256[] resolverFeesETH, bytes delegateSignature, bytes extraData) payable returns()
func (_Contract *ContractTransactor) RevokeBatch0(opts *bind.TransactOpts, attestationIds []uint64, reasons []string, resolverFeesETH []*big.Int, delegateSignature []byte, extraData []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "revokeBatch0", attestationIds, reasons, resolverFeesETH, delegateSignature, extraData)
}

// RevokeBatch0 is a paid mutator transaction binding the contract method 0xd7681235.
//
// Solidity: function revokeBatch(uint64[] attestationIds, string[] reasons, uint256[] resolverFeesETH, bytes delegateSignature, bytes extraData) payable returns()
func (_Contract *ContractSession) RevokeBatch0(attestationIds []uint64, reasons []string, resolverFeesETH []*big.Int, delegateSignature []byte, extraData []byte) (*types.Transaction, error) {
	return _Contract.Contract.RevokeBatch0(&_Contract.TransactOpts, attestationIds, reasons, resolverFeesETH, delegateSignature, extraData)
}

// RevokeBatch0 is a paid mutator transaction binding the contract method 0xd7681235.
//
// Solidity: function revokeBatch(uint64[] attestationIds, string[] reasons, uint256[] resolverFeesETH, bytes delegateSignature, bytes extraData) payable returns()
func (_Contract *ContractTransactorSession) RevokeBatch0(attestationIds []uint64, reasons []string, resolverFeesETH []*big.Int, delegateSignature []byte, extraData []byte) (*types.Transaction, error) {
	return _Contract.Contract.RevokeBatch0(&_Contract.TransactOpts, attestationIds, reasons, resolverFeesETH, delegateSignature, extraData)
}

// RevokeBatch1 is a paid mutator transaction binding the contract method 0xfd397913.
//
// Solidity: function revokeBatch(uint64[] attestationIds, string[] reasons, address[] resolverFeesERC20Tokens, uint256[] resolverFeesERC20Amount, bytes delegateSignature, bytes extraData) returns()
func (_Contract *ContractTransactor) RevokeBatch1(opts *bind.TransactOpts, attestationIds []uint64, reasons []string, resolverFeesERC20Tokens []common.Address, resolverFeesERC20Amount []*big.Int, delegateSignature []byte, extraData []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "revokeBatch1", attestationIds, reasons, resolverFeesERC20Tokens, resolverFeesERC20Amount, delegateSignature, extraData)
}

// RevokeBatch1 is a paid mutator transaction binding the contract method 0xfd397913.
//
// Solidity: function revokeBatch(uint64[] attestationIds, string[] reasons, address[] resolverFeesERC20Tokens, uint256[] resolverFeesERC20Amount, bytes delegateSignature, bytes extraData) returns()
func (_Contract *ContractSession) RevokeBatch1(attestationIds []uint64, reasons []string, resolverFeesERC20Tokens []common.Address, resolverFeesERC20Amount []*big.Int, delegateSignature []byte, extraData []byte) (*types.Transaction, error) {
	return _Contract.Contract.RevokeBatch1(&_Contract.TransactOpts, attestationIds, reasons, resolverFeesERC20Tokens, resolverFeesERC20Amount, delegateSignature, extraData)
}

// RevokeBatch1 is a paid mutator transaction binding the contract method 0xfd397913.
//
// Solidity: function revokeBatch(uint64[] attestationIds, string[] reasons, address[] resolverFeesERC20Tokens, uint256[] resolverFeesERC20Amount, bytes delegateSignature, bytes extraData) returns()
func (_Contract *ContractTransactorSession) RevokeBatch1(attestationIds []uint64, reasons []string, resolverFeesERC20Tokens []common.Address, resolverFeesERC20Amount []*big.Int, delegateSignature []byte, extraData []byte) (*types.Transaction, error) {
	return _Contract.Contract.RevokeBatch1(&_Contract.TransactOpts, attestationIds, reasons, resolverFeesERC20Tokens, resolverFeesERC20Amount, delegateSignature, extraData)
}

// RevokeOffchain is a paid mutator transaction binding the contract method 0x574a3066.
//
// Solidity: function revokeOffchain(string offchainAttestationId, string reason, bytes delegateSignature) returns()
func (_Contract *ContractTransactor) RevokeOffchain(opts *bind.TransactOpts, offchainAttestationId string, reason string, delegateSignature []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "revokeOffchain", offchainAttestationId, reason, delegateSignature)
}

// RevokeOffchain is a paid mutator transaction binding the contract method 0x574a3066.
//
// Solidity: function revokeOffchain(string offchainAttestationId, string reason, bytes delegateSignature) returns()
func (_Contract *ContractSession) RevokeOffchain(offchainAttestationId string, reason string, delegateSignature []byte) (*types.Transaction, error) {
	return _Contract.Contract.RevokeOffchain(&_Contract.TransactOpts, offchainAttestationId, reason, delegateSignature)
}

// RevokeOffchain is a paid mutator transaction binding the contract method 0x574a3066.
//
// Solidity: function revokeOffchain(string offchainAttestationId, string reason, bytes delegateSignature) returns()
func (_Contract *ContractTransactorSession) RevokeOffchain(offchainAttestationId string, reason string, delegateSignature []byte) (*types.Transaction, error) {
	return _Contract.Contract.RevokeOffchain(&_Contract.TransactOpts, offchainAttestationId, reason, delegateSignature)
}

// RevokeOffchainBatch is a paid mutator transaction binding the contract method 0xba97637d.
//
// Solidity: function revokeOffchainBatch(string[] offchainAttestationIds, string[] reasons, bytes delegateSignature) returns()
func (_Contract *ContractTransactor) RevokeOffchainBatch(opts *bind.TransactOpts, offchainAttestationIds []string, reasons []string, delegateSignature []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "revokeOffchainBatch", offchainAttestationIds, reasons, delegateSignature)
}

// RevokeOffchainBatch is a paid mutator transaction binding the contract method 0xba97637d.
//
// Solidity: function revokeOffchainBatch(string[] offchainAttestationIds, string[] reasons, bytes delegateSignature) returns()
func (_Contract *ContractSession) RevokeOffchainBatch(offchainAttestationIds []string, reasons []string, delegateSignature []byte) (*types.Transaction, error) {
	return _Contract.Contract.RevokeOffchainBatch(&_Contract.TransactOpts, offchainAttestationIds, reasons, delegateSignature)
}

// RevokeOffchainBatch is a paid mutator transaction binding the contract method 0xba97637d.
//
// Solidity: function revokeOffchainBatch(string[] offchainAttestationIds, string[] reasons, bytes delegateSignature) returns()
func (_Contract *ContractTransactorSession) RevokeOffchainBatch(offchainAttestationIds []string, reasons []string, delegateSignature []byte) (*types.Transaction, error) {
	return _Contract.Contract.RevokeOffchainBatch(&_Contract.TransactOpts, offchainAttestationIds, reasons, delegateSignature)
}

// SetGlobalHook is a paid mutator transaction binding the contract method 0x10c037a7.
//
// Solidity: function setGlobalHook(address hook) returns()
func (_Contract *ContractTransactor) SetGlobalHook(opts *bind.TransactOpts, hook common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setGlobalHook", hook)
}

// SetGlobalHook is a paid mutator transaction binding the contract method 0x10c037a7.
//
// Solidity: function setGlobalHook(address hook) returns()
func (_Contract *ContractSession) SetGlobalHook(hook common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetGlobalHook(&_Contract.TransactOpts, hook)
}

// SetGlobalHook is a paid mutator transaction binding the contract method 0x10c037a7.
//
// Solidity: function setGlobalHook(address hook) returns()
func (_Contract *ContractTransactorSession) SetGlobalHook(hook common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetGlobalHook(&_Contract.TransactOpts, hook)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool paused) returns()
func (_Contract *ContractTransactor) SetPause(opts *bind.TransactOpts, paused bool) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setPause", paused)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool paused) returns()
func (_Contract *ContractSession) SetPause(paused bool) (*types.Transaction, error) {
	return _Contract.Contract.SetPause(&_Contract.TransactOpts, paused)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool paused) returns()
func (_Contract *ContractTransactorSession) SetPause(paused bool) (*types.Transaction, error) {
	return _Contract.Contract.SetPause(&_Contract.TransactOpts, paused)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Contract *ContractTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Contract *ContractSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Contract.Contract.UpgradeToAndCall(&_Contract.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Contract *ContractTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Contract.Contract.UpgradeToAndCall(&_Contract.TransactOpts, newImplementation, data)
}

// ContractAttestationMadeIterator is returned from FilterAttestationMade and is used to iterate over the raw logs and unpacked data for AttestationMade events raised by the Contract contract.
type ContractAttestationMadeIterator struct {
	Event *ContractAttestationMade // Event containing the contract specifics and raw log

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
func (it *ContractAttestationMadeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAttestationMade)
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
		it.Event = new(ContractAttestationMade)
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
func (it *ContractAttestationMadeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAttestationMadeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAttestationMade represents a AttestationMade event raised by the Contract contract.
type ContractAttestationMade struct {
	AttestationId uint64
	IndexingKey   string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAttestationMade is a free log retrieval operation binding the contract event 0x38b331f00373c8b0f9af5a4273ca1b91e894fec999bfa4ec787e8c90a5c8b05c.
//
// Solidity: event AttestationMade(uint64 attestationId, string indexingKey)
func (_Contract *ContractFilterer) FilterAttestationMade(opts *bind.FilterOpts) (*ContractAttestationMadeIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "AttestationMade")
	if err != nil {
		return nil, err
	}
	return &ContractAttestationMadeIterator{contract: _Contract.contract, event: "AttestationMade", logs: logs, sub: sub}, nil
}

// WatchAttestationMade is a free log subscription operation binding the contract event 0x38b331f00373c8b0f9af5a4273ca1b91e894fec999bfa4ec787e8c90a5c8b05c.
//
// Solidity: event AttestationMade(uint64 attestationId, string indexingKey)
func (_Contract *ContractFilterer) WatchAttestationMade(opts *bind.WatchOpts, sink chan<- *ContractAttestationMade) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "AttestationMade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAttestationMade)
				if err := _Contract.contract.UnpackLog(event, "AttestationMade", log); err != nil {
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

// ParseAttestationMade is a log parse operation binding the contract event 0x38b331f00373c8b0f9af5a4273ca1b91e894fec999bfa4ec787e8c90a5c8b05c.
//
// Solidity: event AttestationMade(uint64 attestationId, string indexingKey)
func (_Contract *ContractFilterer) ParseAttestationMade(log types.Log) (*ContractAttestationMade, error) {
	event := new(ContractAttestationMade)
	if err := _Contract.contract.UnpackLog(event, "AttestationMade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAttestationRevokedIterator is returned from FilterAttestationRevoked and is used to iterate over the raw logs and unpacked data for AttestationRevoked events raised by the Contract contract.
type ContractAttestationRevokedIterator struct {
	Event *ContractAttestationRevoked // Event containing the contract specifics and raw log

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
func (it *ContractAttestationRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAttestationRevoked)
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
		it.Event = new(ContractAttestationRevoked)
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
func (it *ContractAttestationRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAttestationRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAttestationRevoked represents a AttestationRevoked event raised by the Contract contract.
type ContractAttestationRevoked struct {
	AttestationId uint64
	Reason        string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAttestationRevoked is a free log retrieval operation binding the contract event 0x294c02ce66be799794a712d72e9926c40a3cb2738a51571523ffe23be2c6fb01.
//
// Solidity: event AttestationRevoked(uint64 attestationId, string reason)
func (_Contract *ContractFilterer) FilterAttestationRevoked(opts *bind.FilterOpts) (*ContractAttestationRevokedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "AttestationRevoked")
	if err != nil {
		return nil, err
	}
	return &ContractAttestationRevokedIterator{contract: _Contract.contract, event: "AttestationRevoked", logs: logs, sub: sub}, nil
}

// WatchAttestationRevoked is a free log subscription operation binding the contract event 0x294c02ce66be799794a712d72e9926c40a3cb2738a51571523ffe23be2c6fb01.
//
// Solidity: event AttestationRevoked(uint64 attestationId, string reason)
func (_Contract *ContractFilterer) WatchAttestationRevoked(opts *bind.WatchOpts, sink chan<- *ContractAttestationRevoked) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "AttestationRevoked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAttestationRevoked)
				if err := _Contract.contract.UnpackLog(event, "AttestationRevoked", log); err != nil {
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

// ParseAttestationRevoked is a log parse operation binding the contract event 0x294c02ce66be799794a712d72e9926c40a3cb2738a51571523ffe23be2c6fb01.
//
// Solidity: event AttestationRevoked(uint64 attestationId, string reason)
func (_Contract *ContractFilterer) ParseAttestationRevoked(log types.Log) (*ContractAttestationRevoked, error) {
	event := new(ContractAttestationRevoked)
	if err := _Contract.contract.UnpackLog(event, "AttestationRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Contract contract.
type ContractInitializedIterator struct {
	Event *ContractInitialized // Event containing the contract specifics and raw log

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
func (it *ContractInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractInitialized)
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
		it.Event = new(ContractInitialized)
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
func (it *ContractInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractInitialized represents a Initialized event raised by the Contract contract.
type ContractInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Contract *ContractFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractInitializedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractInitializedIterator{contract: _Contract.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Contract *ContractFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractInitialized) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractInitialized)
				if err := _Contract.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Contract *ContractFilterer) ParseInitialized(log types.Log) (*ContractInitialized, error) {
	event := new(ContractInitialized)
	if err := _Contract.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOffchainAttestationMadeIterator is returned from FilterOffchainAttestationMade and is used to iterate over the raw logs and unpacked data for OffchainAttestationMade events raised by the Contract contract.
type ContractOffchainAttestationMadeIterator struct {
	Event *ContractOffchainAttestationMade // Event containing the contract specifics and raw log

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
func (it *ContractOffchainAttestationMadeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOffchainAttestationMade)
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
		it.Event = new(ContractOffchainAttestationMade)
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
func (it *ContractOffchainAttestationMadeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOffchainAttestationMadeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOffchainAttestationMade represents a OffchainAttestationMade event raised by the Contract contract.
type ContractOffchainAttestationMade struct {
	AttestationId string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOffchainAttestationMade is a free log retrieval operation binding the contract event 0x6f0bf119c2fc55752051c1cfde5d145fdd499c7f748df7920409c72319ff6027.
//
// Solidity: event OffchainAttestationMade(string attestationId)
func (_Contract *ContractFilterer) FilterOffchainAttestationMade(opts *bind.FilterOpts) (*ContractOffchainAttestationMadeIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OffchainAttestationMade")
	if err != nil {
		return nil, err
	}
	return &ContractOffchainAttestationMadeIterator{contract: _Contract.contract, event: "OffchainAttestationMade", logs: logs, sub: sub}, nil
}

// WatchOffchainAttestationMade is a free log subscription operation binding the contract event 0x6f0bf119c2fc55752051c1cfde5d145fdd499c7f748df7920409c72319ff6027.
//
// Solidity: event OffchainAttestationMade(string attestationId)
func (_Contract *ContractFilterer) WatchOffchainAttestationMade(opts *bind.WatchOpts, sink chan<- *ContractOffchainAttestationMade) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OffchainAttestationMade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOffchainAttestationMade)
				if err := _Contract.contract.UnpackLog(event, "OffchainAttestationMade", log); err != nil {
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

// ParseOffchainAttestationMade is a log parse operation binding the contract event 0x6f0bf119c2fc55752051c1cfde5d145fdd499c7f748df7920409c72319ff6027.
//
// Solidity: event OffchainAttestationMade(string attestationId)
func (_Contract *ContractFilterer) ParseOffchainAttestationMade(log types.Log) (*ContractOffchainAttestationMade, error) {
	event := new(ContractOffchainAttestationMade)
	if err := _Contract.contract.UnpackLog(event, "OffchainAttestationMade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOffchainAttestationRevokedIterator is returned from FilterOffchainAttestationRevoked and is used to iterate over the raw logs and unpacked data for OffchainAttestationRevoked events raised by the Contract contract.
type ContractOffchainAttestationRevokedIterator struct {
	Event *ContractOffchainAttestationRevoked // Event containing the contract specifics and raw log

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
func (it *ContractOffchainAttestationRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOffchainAttestationRevoked)
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
		it.Event = new(ContractOffchainAttestationRevoked)
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
func (it *ContractOffchainAttestationRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOffchainAttestationRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOffchainAttestationRevoked represents a OffchainAttestationRevoked event raised by the Contract contract.
type ContractOffchainAttestationRevoked struct {
	AttestationId string
	Reason        string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOffchainAttestationRevoked is a free log retrieval operation binding the contract event 0x8cdef7ac4262df432ef960b4d902483c56118d54647a70746d99abc05cb970ba.
//
// Solidity: event OffchainAttestationRevoked(string attestationId, string reason)
func (_Contract *ContractFilterer) FilterOffchainAttestationRevoked(opts *bind.FilterOpts) (*ContractOffchainAttestationRevokedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OffchainAttestationRevoked")
	if err != nil {
		return nil, err
	}
	return &ContractOffchainAttestationRevokedIterator{contract: _Contract.contract, event: "OffchainAttestationRevoked", logs: logs, sub: sub}, nil
}

// WatchOffchainAttestationRevoked is a free log subscription operation binding the contract event 0x8cdef7ac4262df432ef960b4d902483c56118d54647a70746d99abc05cb970ba.
//
// Solidity: event OffchainAttestationRevoked(string attestationId, string reason)
func (_Contract *ContractFilterer) WatchOffchainAttestationRevoked(opts *bind.WatchOpts, sink chan<- *ContractOffchainAttestationRevoked) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OffchainAttestationRevoked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOffchainAttestationRevoked)
				if err := _Contract.contract.UnpackLog(event, "OffchainAttestationRevoked", log); err != nil {
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

// ParseOffchainAttestationRevoked is a log parse operation binding the contract event 0x8cdef7ac4262df432ef960b4d902483c56118d54647a70746d99abc05cb970ba.
//
// Solidity: event OffchainAttestationRevoked(string attestationId, string reason)
func (_Contract *ContractFilterer) ParseOffchainAttestationRevoked(log types.Log) (*ContractOffchainAttestationRevoked, error) {
	event := new(ContractOffchainAttestationRevoked)
	if err := _Contract.contract.UnpackLog(event, "OffchainAttestationRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Contract contract.
type ContractOwnershipTransferredIterator struct {
	Event *ContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOwnershipTransferred)
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
		it.Event = new(ContractOwnershipTransferred)
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
func (it *ContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOwnershipTransferred represents a OwnershipTransferred event raised by the Contract contract.
type ContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractOwnershipTransferredIterator{contract: _Contract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOwnershipTransferred)
				if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) ParseOwnershipTransferred(log types.Log) (*ContractOwnershipTransferred, error) {
	event := new(ContractOwnershipTransferred)
	if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSchemaRegisteredIterator is returned from FilterSchemaRegistered and is used to iterate over the raw logs and unpacked data for SchemaRegistered events raised by the Contract contract.
type ContractSchemaRegisteredIterator struct {
	Event *ContractSchemaRegistered // Event containing the contract specifics and raw log

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
func (it *ContractSchemaRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSchemaRegistered)
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
		it.Event = new(ContractSchemaRegistered)
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
func (it *ContractSchemaRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSchemaRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSchemaRegistered represents a SchemaRegistered event raised by the Contract contract.
type ContractSchemaRegistered struct {
	SchemaId uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSchemaRegistered is a free log retrieval operation binding the contract event 0xd0f9818d35b9c7d941f89e81a08a7f4761384ae32aeaf4a913b94319a321e7ff.
//
// Solidity: event SchemaRegistered(uint64 schemaId)
func (_Contract *ContractFilterer) FilterSchemaRegistered(opts *bind.FilterOpts) (*ContractSchemaRegisteredIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "SchemaRegistered")
	if err != nil {
		return nil, err
	}
	return &ContractSchemaRegisteredIterator{contract: _Contract.contract, event: "SchemaRegistered", logs: logs, sub: sub}, nil
}

// WatchSchemaRegistered is a free log subscription operation binding the contract event 0xd0f9818d35b9c7d941f89e81a08a7f4761384ae32aeaf4a913b94319a321e7ff.
//
// Solidity: event SchemaRegistered(uint64 schemaId)
func (_Contract *ContractFilterer) WatchSchemaRegistered(opts *bind.WatchOpts, sink chan<- *ContractSchemaRegistered) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "SchemaRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSchemaRegistered)
				if err := _Contract.contract.UnpackLog(event, "SchemaRegistered", log); err != nil {
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

// ParseSchemaRegistered is a log parse operation binding the contract event 0xd0f9818d35b9c7d941f89e81a08a7f4761384ae32aeaf4a913b94319a321e7ff.
//
// Solidity: event SchemaRegistered(uint64 schemaId)
func (_Contract *ContractFilterer) ParseSchemaRegistered(log types.Log) (*ContractSchemaRegistered, error) {
	event := new(ContractSchemaRegistered)
	if err := _Contract.contract.UnpackLog(event, "SchemaRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Contract contract.
type ContractUpgradedIterator struct {
	Event *ContractUpgraded // Event containing the contract specifics and raw log

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
func (it *ContractUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractUpgraded)
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
		it.Event = new(ContractUpgraded)
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
func (it *ContractUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractUpgraded represents a Upgraded event raised by the Contract contract.
type ContractUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Contract *ContractFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ContractUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ContractUpgradedIterator{contract: _Contract.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Contract *ContractFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ContractUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractUpgraded)
				if err := _Contract.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Contract *ContractFilterer) ParseUpgraded(log types.Log) (*ContractUpgraded, error) {
	event := new(ContractUpgraded)
	if err := _Contract.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
