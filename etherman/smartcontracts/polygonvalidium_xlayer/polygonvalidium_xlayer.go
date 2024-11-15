// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package polygonvalidium_xlayer

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

// PolygonRollupBaseEtrogBatchData is an auto generated low-level Go binding around an user-defined struct.
type PolygonRollupBaseEtrogBatchData struct {
	Transactions         []byte
	ForcedGlobalExitRoot [32]byte
	ForcedTimestamp      uint64
	ForcedBlockHashL1    [32]byte
}

// PolygonValidiumEtrogValidiumBatchData is an auto generated low-level Go binding around an user-defined struct.
type PolygonValidiumEtrogValidiumBatchData struct {
	TransactionsHash     [32]byte
	ForcedGlobalExitRoot [32]byte
	ForcedTimestamp      uint64
	ForcedBlockHashL1    [32]byte
}

// PolygonvalidiumXlayerMetaData contains all meta data concerning the PolygonvalidiumXlayer contract.
var PolygonvalidiumXlayerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRootV2\",\"name\":\"_globalExitRootManager\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"_pol\",\"type\":\"address\"},{\"internalType\":\"contractIPolygonZkEVMBridgeV2\",\"name\":\"_bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"contractPolygonRollupManager\",\"name\":\"_rollupManager\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BatchAlreadyVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchNotSequencedOrNotSequenceEnd\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceedMaxVerifyBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalAccInputHashDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchBelowLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalPendingStateNumInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesAlreadyActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesDecentralized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesNotAllowedOnEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForcedDataDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GasTokenNetworkMustBeZeroOnEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalExitRootNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpiredAfterEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HugeTokenMetadataNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchAboveLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitSequencedBatchDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializeTransaction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeBatchTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeForceBatchTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierBatchFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1InfoRootIndexInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxTimestampSequenceInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewPendingStateTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewStateRootNotInsidePrime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewTrustedAggregatorTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughMaticAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughPOLAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldStateRootDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRollupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedAggregator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedSequencer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateNotConsolidable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequenceWithDataAvailabilityNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequenceZeroBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampBelowForcedTimestamp\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StoredRootMustBeDifferentThanNewRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SwitchToSameValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransactionsLengthAboveMax\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutNotExpired\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AcceptAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"forceBatchNum\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"lastGlobalExitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"}],\"name\":\"ForceBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"lastGlobalExitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"}],\"name\":\"InitialSequenceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"targetBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"accInputHashToRollback\",\"type\":\"bytes32\"}],\"name\":\"RollbackBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"l1InfoRoot\",\"type\":\"bytes32\"}],\"name\":\"SequenceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"}],\"name\":\"SequenceForceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newDataAvailabilityProtocol\",\"type\":\"address\"}],\"name\":\"SetDataAvailabilityProtocol\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newForceBatchAddress\",\"type\":\"address\"}],\"name\":\"SetForceBatchAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newforceBatchTimeout\",\"type\":\"uint64\"}],\"name\":\"SetForceBatchTimeout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"SetTrustedSequencer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"SetTrustedSequencerURL\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"SwitchSequenceWithDataAvailability\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"TransferAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"VerifyBatches\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"GLOBAL_EXIT_ROOT_MANAGER_L2\",\"outputs\":[{\"internalType\":\"contractIBasePolygonZkEVMGlobalExitRoot\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_BRIDGE_LIST_LEN_LEN\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_BRIDGE_PARAMS\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_CONSTANT_BYTES\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_DATA_LEN_EMPTY_METADATA\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_EFFECTIVE_PERCENTAGE\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SIGNATURE_INITIALIZE_TX_R\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SIGNATURE_INITIALIZE_TX_S\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SIGNATURE_INITIALIZE_TX_V\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TIMESTAMP_RANGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMBridgeV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculatePolPerForceBatch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dataAvailabilityProtocol\",\"outputs\":[{\"internalType\":\"contractIDataAvailabilityProtocol\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"polAmount\",\"type\":\"uint256\"}],\"name\":\"forceBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceBatchAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceBatchTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"forcedBatches\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenNetwork\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"networkID\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_gasTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_gasTokenNetwork\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_gasTokenMetadata\",\"type\":\"bytes\"}],\"name\":\"generateInitializeTransaction\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootManager\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRootV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"networkID\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_gasTokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"sequencerURL\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_networkName\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isSequenceWithDataAvailabilityAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastAccInputHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatchSequenced\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"onVerifyBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pol\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"targetBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"accInputHashToRollback\",\"type\":\"bytes32\"}],\"name\":\"rollbackBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupManager\",\"outputs\":[{\"internalType\":\"contractPolygonRollupManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"forcedGlobalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"forcedTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"forcedBlockHashL1\",\"type\":\"bytes32\"}],\"internalType\":\"structPolygonRollupBaseEtrog.BatchData[]\",\"name\":\"batches\",\"type\":\"tuple[]\"},{\"internalType\":\"uint32\",\"name\":\"indexL1InfoRoot\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"maxSequenceTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"expectedFinalAccInputHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"l2Coinbase\",\"type\":\"address\"}],\"name\":\"sequenceBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"transactionsHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"forcedGlobalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"forcedTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"forcedBlockHashL1\",\"type\":\"bytes32\"}],\"internalType\":\"structPolygonValidiumEtrog.ValidiumBatchData[]\",\"name\":\"batches\",\"type\":\"tuple[]\"},{\"internalType\":\"uint32\",\"name\":\"indexL1InfoRoot\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"maxSequenceTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"expectedFinalAccInputHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"l2Coinbase\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"dataAvailabilityMessage\",\"type\":\"bytes\"}],\"name\":\"sequenceBatchesValidium\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"forcedGlobalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"forcedTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"forcedBlockHashL1\",\"type\":\"bytes32\"}],\"internalType\":\"structPolygonRollupBaseEtrog.BatchData[]\",\"name\":\"batches\",\"type\":\"tuple[]\"}],\"name\":\"sequenceForceBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIDataAvailabilityProtocol\",\"name\":\"newDataAvailabilityProtocol\",\"type\":\"address\"}],\"name\":\"setDataAvailabilityProtocol\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newForceBatchAddress\",\"type\":\"address\"}],\"name\":\"setForceBatchAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newforceBatchTimeout\",\"type\":\"uint64\"}],\"name\":\"setForceBatchTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"setTrustedSequencer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"setTrustedSequencerURL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"newIsSequenceWithDataAvailabilityAllowed\",\"type\":\"bool\"}],\"name\":\"switchSequenceWithDataAvailability\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"transferAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencerURL\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x61010060405234801562000011575f80fd5b506040516200524e3803806200524e83398101604081905262000034916200006f565b6001600160a01b0393841660a052918316608052821660c0521660e052620000d4565b6001600160a01b03811681146200006c575f80fd5b50565b5f805f806080858703121562000083575f80fd5b8451620000908162000057565b6020860151909450620000a38162000057565b6040860151909350620000b68162000057565b6060860151909250620000c98162000057565b939692955090935050565b60805160a05160c05160e051615066620001e85f395f818161055901528181611027015281816110ee01528181611110015281816112ac015281816114db01528181611647015281816117c60152818161194601528181611cc80152818161222c0152818161267b0152818161272801528181612e0e01528181612ed301528181613c2301528181613c9c01528181613cbe0152613d6601525f818161070d01528181610b4e015281816128f2015281816129fa015281816133fd015281816134d2015261371401525f81816107dc01528181610bfc01528181611b3f0152818161301b01526137c201525f818161080e015281816108eb0152818161107a015281816111bc0152612ff001526150665ff3fe608060405234801561000f575f80fd5b5060043610610319575f3560e01c80636ff512cc116101a8578063b910e0f9116100f3578063d7bc90ff1161009e578063e7a7ed0211610079578063e7a7ed0214610850578063eaeb077b14610864578063f35dda4714610877578063f851a4401461087f575f80fd5b8063d7bc90ff146107fe578063e46761c414610809578063e57a0b4c14610830575f80fd5b8063c89e42df116100ce578063c89e42df146107a4578063cfa8ed47146107b7578063d02103ca146107d7575f80fd5b8063b910e0f914610761578063c754c7ed14610774578063c7fffd4b1461079c575f80fd5b80639e00187711610153578063a652f26c1161012e578063a652f26c1461072f578063ada8f91914610742578063b0afe15414610755575f80fd5b80639e001877146106da5780639f26f840146106f5578063a3c573eb14610708575f80fd5b80637cd76b8b116101835780637cd76b8b146106ac5780638c3d7301146106bf57806391cafe32146106c7575f80fd5b80636ff512cc1461064a578063712570221461065d5780637a5460c514610670575f80fd5b806340b5de6c1161026857806352bdeb6d11610213578063676870d2116101ee578063676870d21461061a5780636b8616ce146106225780636e05d2cd14610641575f80fd5b806352bdeb6d146105c3578063542028d5146105ff578063669adece14610607575f80fd5b806349b7b8021161024357806349b7b802146105545780634c21fef31461057b5780634e487706146105b0575f80fd5b806340b5de6c146104bb57806342308fab14610513578063456052671461051b575f80fd5b806326782247116102c857806332c2d153116102a357806332c2d1531461044b5780633c351e101461045e5780633cbc795b1461047e575f80fd5b806326782247146103d35780632acdc2b6146104185780632c111c061461042b575f80fd5b8063107bf28c116102f8578063107bf28c1461039c57806311e892d4146103a4578063165e8a8d146103be575f80fd5b8062d0295d1461031d578063035089631461033857806305835f3714610353575b5f80fd5b6103256108a4565b6040519081526020015b60405180910390f35b610340602081565b60405161ffff909116815260200161032f565b61038f6040518060400160405280600881526020017f80808401c9c3809400000000000000000000000000000000000000000000000081525081565b60405161032f9190614154565b61038f6109aa565b6103ac60f981565b60405160ff909116815260200161032f565b6103d16103cc36600461421d565b610a36565b005b6001546103f39073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161032f565b6103d16104263660046142fc565b6113bb565b6008546103f39073ffffffffffffffffffffffffffffffffffffffff1681565b6103d1610459366004614317565b6114d9565b6009546103f39073ffffffffffffffffffffffffffffffffffffffff1681565b6009546104a69074010000000000000000000000000000000000000000900463ffffffff1681565b60405163ffffffff909116815260200161032f565b6104e27fff0000000000000000000000000000000000000000000000000000000000000081565b6040517fff00000000000000000000000000000000000000000000000000000000000000909116815260200161032f565b610325602481565b60075461053b9068010000000000000000900467ffffffffffffffff1681565b60405167ffffffffffffffff909116815260200161032f565b6103f37f000000000000000000000000000000000000000000000000000000000000000081565b603c546105a09074010000000000000000000000000000000000000000900460ff1681565b604051901515815260200161032f565b6103d16105be366004614356565b6115a8565b61038f6040518060400160405280600281526020017f80b800000000000000000000000000000000000000000000000000000000000081525081565b61038f6117b7565b6103d1610615366004614371565b6117c4565b610340601f81565b610325610630366004614356565b60066020525f908152604090205481565b61032560055481565b6103d161065836600461439b565b61187b565b6103d161066b3660046144f2565b611944565b61038f6040518060400160405280600281526020017f80b900000000000000000000000000000000000000000000000000000000000081525081565b6103d16106ba36600461439b565b611efd565b6103d1611fc6565b6103d16106d536600461439b565b612098565b6103f373a40d5f56745a118d0906a34e69aec8c0db1cb8fa81565b6103d16107033660046145da565b6121b0565b6103f37f000000000000000000000000000000000000000000000000000000000000000081565b61038f61073d366004614619565b6127f4565b6103d161075036600461439b565b612bd2565b6103256405ca1ab1e081565b6103d161076f36600461468a565b612c9b565b60075461053b90700100000000000000000000000000000000900467ffffffffffffffff1681565b6103ac60e481565b6103d16107b236600461470a565b612d05565b6002546103f39073ffffffffffffffffffffffffffffffffffffffff1681565b6103f37f000000000000000000000000000000000000000000000000000000000000000081565b610325635ca1ab1e81565b6103f37f000000000000000000000000000000000000000000000000000000000000000081565b603c546103f39073ffffffffffffffffffffffffffffffffffffffff1681565b60075461053b9067ffffffffffffffff1681565b6103d161087236600461473c565b612d97565b6103ac601b81565b5f546103f39062010000900473ffffffffffffffffffffffffffffffffffffffff1681565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201525f90819073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016906370a0823190602401602060405180830381865afa158015610930573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109549190614784565b6007549091505f9061097e9067ffffffffffffffff680100000000000000008204811691166147c8565b67ffffffffffffffff169050805f03610999575f9250505090565b6109a381836147f0565b9250505090565b600480546109b790614828565b80601f01602080910402602001604051908101604052809291908181526020018280546109e390614828565b8015610a2e5780601f10610a0557610100808354040283529160200191610a2e565b820191905f5260205f20905b815481529060010190602001808311610a1157829003601f168201915b505050505081565b60025473ffffffffffffffffffffffffffffffffffffffff163314610a87576040517f11e7be1500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b865f819003610ac2576040517fcb591a5f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103e8811115610afe576040517fb59f753a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610b09602442614879565b8667ffffffffffffffff161115610b4c576040517f0a00feb300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166379e2cf976040518163ffffffff1660e01b81526004015f604051808303815f87803b158015610bb1575f80fd5b505af1158015610bc3573d5f803e3d5ffd5b50506040517fef4eeb3500000000000000000000000000000000000000000000000000000000815263ffffffff8a1660048201525f92507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16915063ef4eeb3590602401602060405180830381865afa158015610c57573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610c7b9190614784565b905080610cb4576040517f0cc1522c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6007546005546801000000000000000090910467ffffffffffffffff1690815f5b85811015610f9a575f8e8e83818110610cf057610cf0614892565b905060800201803603810190610d0691906148bf565b604081015190915067ffffffffffffffff1615610ef35784610d278161490b565b9550505f815f0151826020015183604001518460600151604051602001610d8c9493929190938452602084019290925260c01b7fffffffffffffffff000000000000000000000000000000000000000000000000166040830152604882015260680190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152815160209283012067ffffffffffffffff89165f90815260069093529120549091508114610e14576040517fce3d755e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84825f0151836020015184604001518e8660600151604051602001610ead969594939291909586526020860194909452604085019290925260c01b7fffffffffffffffff000000000000000000000000000000000000000000000000166060808501919091521b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166068830152607c820152609c0190565b60405160208183030381529060405280519060200120945060065f8767ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f905550610f87565b8051604080516020810187905290810191909152606080820188905260c08e901b7fffffffffffffffff0000000000000000000000000000000000000000000000001660808301528b901b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000001660888201525f609c82015260bc016040516020818303038152906040528051906020012093505b5080610f9281614931565b915050610cd5565b5060075467ffffffffffffffff9081169084161115610fe5576040517fc630a00d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60058290558467ffffffffffffffff848116908316146110e2575f61100a83866147c8565b905061102067ffffffffffffffff821683614968565b91506110a17f00000000000000000000000000000000000000000000000000000000000000008267ffffffffffffffff166110596108a4565b611063919061497b565b73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016919061325d565b50600780547fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff166801000000000000000067ffffffffffffffff8716021790555b801561126b576111e4337f0000000000000000000000000000000000000000000000000000000000000000837f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663477fa2706040518163ffffffff1660e01b8152600401602060405180830381865afa158015611177573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061119b9190614784565b6111a5919061497b565b73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016929190613336565b603c546040517f3b51be4b00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911690633b51be4b9061123e908d908c908c906004016149d9565b5f6040518083038186803b158015611254575f80fd5b505afa158015611266573d5f803e3d5ffd5b505050505b6040517f9a908e7300000000000000000000000000000000000000000000000000000000815267ffffffffffffffff87166004820152602481018490525f907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690639a908e73906044016020604051808303815f875af1158015611307573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061132b91906149fb565b90508a8414611366576040517fda5bceb900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8067ffffffffffffffff167f3e54d0825ed78523037d00a81759237eb436ce774bd546993ee67a1b67b6e766876040516113a291815260200190565b60405180910390a2505050505050505050505050505050565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff163314611411576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603c5474010000000000000000000000000000000000000000900460ff1615158115150361146b576040517f5f0e7abe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603c80547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1674010000000000000000000000000000000000000000831515021790556040517ff32a0473f809a720a4f8af1e50d353f1caf7452030626fdaac4273f5e6587f41905f90a150565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163314611548576040517fb9b3a2c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff168367ffffffffffffffff167f9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f59668460405161159b91815260200190565b60405180910390a3505050565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff1633146115fe576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b62093a8067ffffffffffffffff82161115611645576040517ff5e37f2f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166315064c966040518163ffffffff1660e01b8152600401602060405180830381865afa1580156116ae573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906116d29190614a16565b6117335760075467ffffffffffffffff700100000000000000000000000000000000909104811690821610611733576040517ff5e37f2f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600780547fffffffffffffffff0000000000000000ffffffffffffffffffffffffffffffff1670010000000000000000000000000000000067ffffffffffffffff8416908102919091179091556040519081527fa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b906020015b60405180910390a150565b600380546109b790614828565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163314611833576040517fb9b3a2c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600581905560405181815267ffffffffffffffff8316907f1125aaf62d132d8e2d02005114f8fc360ff204c3105e4f1a700a1340dc55d5b19060200160405180910390a25050565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff1633146118d1576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527ff54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0906020016117ac565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1633146119b3576040517fb9b3a2c800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f54610100900460ff16158080156119d157505f54600160ff909116105b806119ea5750303b1580156119ea57505f5460ff166001145b611a7b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084015b60405180910390fd5b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015611ad7575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b5f611ae18561339a565b6009549091505f90611b2b90889073ffffffffffffffffffffffffffffffffffffffff81169074010000000000000000000000000000000000000000900463ffffffff16856127f4565b90505f818051906020012090505f4290505f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16633ed691ef6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611ba6573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611bca9190614784565b90505f808483858f611bdd600143614968565b60408051602081019790975286019490945260608086019390935260c09190911b7fffffffffffffffff000000000000000000000000000000000000000000000000166080850152901b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016608883015240609c82015260bc01604080518083037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe001815290829052805160209091012060058190557f9a908e73000000000000000000000000000000000000000000000000000000008252600160048301526024820181905291507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690639a908e73906044016020604051808303815f875af1158015611d23573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611d4791906149fb565b508c5f60026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508b60025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508860039081611dd79190614a76565b506004611de48982614a76565b508c60085f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555062069780600760106101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055507f060116213bcbf54ca19fd649dc84b59ab2bbd200ab199770e4d923e222a28e7f85838e604051611e8493929190614b8e565b60405180910390a15050505050508015611ef4575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050505050565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff163314611f53576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603c80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fd331bd4c4cd1afecb94a225184bded161ff3213624ba4fb58c4f30c5a861144a906020016117ac565b60015473ffffffffffffffffffffffffffffffffffffffff163314612017576040517fd1ec4b2300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001545f80547fffffffffffffffffffff0000000000000000000000000000000000000000ffff1673ffffffffffffffffffffffffffffffffffffffff9092166201000081029290921790556040519081527f056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e9060200160405180910390a1565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff1633146120ee576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60085473ffffffffffffffffffffffffffffffffffffffff1661213d576040517fc89374d800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600880547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527f5fbd7dd171301c4a1611a84aac4ba86d119478560557755f7927595b082634fb906020016117ac565b60085473ffffffffffffffffffffffffffffffffffffffff1680158015906121ee575073ffffffffffffffffffffffffffffffffffffffff81163314155b15612225576040517f24eff8c300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b4262093a807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166330c27dde6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612293573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906122b791906149fb565b6122c19190614bcc565b67ffffffffffffffff161115612303576040517f3d49ed4c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b815f81900361233e576040517fcb591a5f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103e881111561237a576040517fb59f753a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60075467ffffffffffffffff808216916123a291849168010000000000000000900416614879565b11156123da576040517fc630a00d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6007546005546801000000000000000090910467ffffffffffffffff16905f5b83811015612675575f87878381811061241557612415614892565b90506020028101906124279190614bed565b61243090614c29565b90508361243c8161490b565b825180516020918201208185015160408087015160608801519151959a509295505f946124a8948794929101938452602084019290925260c01b7fffffffffffffffff000000000000000000000000000000000000000000000000166040830152604882015260680190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152815160209283012067ffffffffffffffff89165f90815260069093529120549091508114612530576040517fce3d755e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff86165f90815260066020526040812055612554600188614968565b84036125c35742600760109054906101000a900467ffffffffffffffff1684604001516125819190614bcc565b67ffffffffffffffff1611156125c3576040517fc44a082100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60208381015160408086015160608088015183519586018b90529285018790528481019390935260c01b7fffffffffffffffff0000000000000000000000000000000000000000000000001660808401523390911b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166088830152609c82015260bc01604051602081830303815290604052805190602001209450505050808061266d90614931565b9150506123fa565b506126a37f0000000000000000000000000000000000000000000000000000000000000000846110596108a4565b60058190556007805467ffffffffffffffff841668010000000000000000027fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff9091161790556040517f9a908e730000000000000000000000000000000000000000000000000000000081525f9073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690639a908e7390612774908790869060040167ffffffffffffffff929092168252602082015260400190565b6020604051808303815f875af1158015612790573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906127b491906149fb565b60405190915067ffffffffffffffff8216907f648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a4905f90a250505050505050565b60605f85858573a40d5f56745a118d0906a34e69aec8c0db1cb8fa5f8760405160240161282696959493929190614c94565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167ff811bff70000000000000000000000000000000000000000000000000000000017905283519091506060905f036129765760f9601f83516128ba9190614cf6565b6040518060400160405280600881526020017f80808401c9c380940000000000000000000000000000000000000000000000008152507f00000000000000000000000000000000000000000000000000000000000000006040518060400160405280600281526020017f80b800000000000000000000000000000000000000000000000000000000000081525060e4876040516020016129609796959493929190614d11565b6040516020818303038152906040529050612a7a565b815161ffff10156129b3576040517f248b8f8200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b815160f96129c2602083614cf6565b6040518060400160405280600881526020017f80808401c9c380940000000000000000000000000000000000000000000000008152507f00000000000000000000000000000000000000000000000000000000000000006040518060400160405280600281526020017f80b90000000000000000000000000000000000000000000000000000000000008152508588604051602001612a679796959493929190614df3565b6040516020818303038152906040529150505b8051602080830191909120604080515f80825293810180835292909252601b908201526405ca1ab1e06060820152635ca1ab1e608082015260019060a0016020604051602081039080840390855afa158015612ad8573d5f803e3d5ffd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff8116612b50576040517fcd16196600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040515f90612b959084906405ca1ab1e090635ca1ab1e90601b907fff0000000000000000000000000000000000000000000000000000000000000090602001614ed5565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190529450505050505b949350505050565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff163314612c28576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6906020016117ac565b603c5474010000000000000000000000000000000000000000900460ff16612cef576040517f821935b400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612cfd8686868686866135fc565b505050505050565b5f5462010000900473ffffffffffffffffffffffffffffffffffffffff163314612d5b576040517f4755657900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6003612d678282614a76565b507f6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20816040516117ac9190614154565b60085473ffffffffffffffffffffffffffffffffffffffff168015801590612dd5575073ffffffffffffffffffffffffffffffffffffffff81163314155b15612e0c576040517f24eff8c300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166315064c966040518163ffffffff1660e01b8152600401602060405180830381865afa158015612e75573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612e999190614a16565b15612ed0576040517f39258d1800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663604691696040518163ffffffff1660e01b8152600401602060405180830381865afa158015612f3a573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612f5e9190614784565b905082811115612f9a576040517f2354600f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611388841115612fd6576040517fa29a6c7c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61301873ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016333084613336565b5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16633ed691ef6040518163ffffffff1660e01b8152600401602060405180830381865afa158015613082573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906130a69190614784565b6007805491925067ffffffffffffffff909116905f6130c48361490b565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055505085856040516130fb929190614f30565b6040519081900390208142613111600143614968565b60408051602081019590955284019290925260c01b7fffffffffffffffff000000000000000000000000000000000000000000000000166060830152406068820152608801604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152815160209283012060075467ffffffffffffffff165f9081526006909352912055323303613207576007546040805183815233602082015260608183018190525f90820152905167ffffffffffffffff909216917ff94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc9319181900360800190a2612cfd565b60075460405167ffffffffffffffff909116907ff94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc9319061324d90849033908b908b90614f3f565b60405180910390a2505050505050565b60405173ffffffffffffffffffffffffffffffffffffffff83166024820152604481018290526133319084907fa9059cbb00000000000000000000000000000000000000000000000000000000906064015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152613e73565b505050565b60405173ffffffffffffffffffffffffffffffffffffffff808516602483015283166044820152606481018290526133949085907f23b872dd00000000000000000000000000000000000000000000000000000000906084016132af565b50505050565b606073ffffffffffffffffffffffffffffffffffffffff8216156135f7576040517fc00f14ab00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063c00f14ab906024015f60405180830381865afa158015613441573d5f803e3d5ffd5b505050506040513d5f823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682016040526134869190810190614f7e565b6040517f318aee3d00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84811660048301529192505f9182917f00000000000000000000000000000000000000000000000000000000000000009091169063318aee3d906024016040805180830381865afa158015613518573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061353c9190614fe7565b915091508163ffffffff165f146135b3576009805463ffffffff841674010000000000000000000000000000000000000000027fffffffffffffffff00000000000000000000000000000000000000000000000090911673ffffffffffffffffffffffffffffffffffffffff8416171790556135f4565b600980547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff86161790555b50505b919050565b60025473ffffffffffffffffffffffffffffffffffffffff16331461364d576040517f11e7be1500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b845f819003613688576040517fcb591a5f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103e88111156136c4576040517fb59f753a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6136cf602442614879565b8467ffffffffffffffff161115613712576040517f0a00feb300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166379e2cf976040518163ffffffff1660e01b81526004015f604051808303815f87803b158015613777575f80fd5b505af1158015613789573d5f803e3d5ffd5b50506040517fef4eeb3500000000000000000000000000000000000000000000000000000000815263ffffffff881660048201525f92507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16915063ef4eeb3590602401602060405180830381865afa15801561381d573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906138419190614784565b90508061387a576040517f0cc1522c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6007546005546801000000000000000090910467ffffffffffffffff1690815f5b85811015613b96575f8c8c838181106138b6576138b6614892565b90506020028101906138c89190614bed565b6138d190614c29565b8051805160209091012060408201519192509067ffffffffffffffff1615613ab157856138fd8161490b565b9650505f8183602001518460400151856060015160405160200161395f9493929190938452602084019290925260c01b7fffffffffffffffff000000000000000000000000000000000000000000000000166040830152604882015260680190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152815160209283012067ffffffffffffffff8a165f908152600690935291205490915081146139e7576040517fce3d755e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60208381015160408086015160608088015183519586018c90529285018790528481019390935260c01b7fffffffffffffffff000000000000000000000000000000000000000000000000166080840152908c901b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166088830152609c82015260bc0160405160208183030381529060405280519060200120955060065f8867ffffffffffffffff1667ffffffffffffffff1681526020019081526020015f205f905550613b81565b8151516201d4c01015613af0576040517fa29a6c7c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805160208101879052908101829052606080820189905260c08d901b7fffffffffffffffff0000000000000000000000000000000000000000000000001660808301528a901b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000001660888201525f609c82015260bc016040516020818303038152906040528051906020012094505b50508080613b8e90614931565b91505061389b565b5060075467ffffffffffffffff9081169084161115613be1576040517fc630a00d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60058290558467ffffffffffffffff84811690831614613c96575f613c0683866147c8565b9050613c1c67ffffffffffffffff821683614968565b9150613c557f00000000000000000000000000000000000000000000000000000000000000008267ffffffffffffffff166110596108a4565b50600780547fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff166801000000000000000067ffffffffffffffff8716021790555b613d25337f0000000000000000000000000000000000000000000000000000000000000000837f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663477fa2706040518163ffffffff1660e01b8152600401602060405180830381865afa158015611177573d5f803e3d5ffd5b6040517f9a908e7300000000000000000000000000000000000000000000000000000000815267ffffffffffffffff87166004820152602481018490525f907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690639a908e73906044016020604051808303815f875af1158015613dc1573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613de591906149fb565b9050888414613e20576040517fda5bceb900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8067ffffffffffffffff167f3e54d0825ed78523037d00a81759237eb436ce774bd546993ee67a1b67b6e76687604051613e5c91815260200190565b60405180910390a250505050505050505050505050565b5f613ed4826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff16613f7e9092919063ffffffff16565b8051909150156133315780806020019051810190613ef29190614a16565b613331576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401611a72565b6060612bca84845f85855f808673ffffffffffffffffffffffffffffffffffffffff168587604051613fb0919061501f565b5f6040518083038185875af1925050503d805f8114613fea576040519150601f19603f3d011682016040523d82523d5f602084013e613fef565b606091505b50915091506140008783838761400b565b979650505050505050565b606083156140a05782515f036140995773ffffffffffffffffffffffffffffffffffffffff85163b614099576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401611a72565b5081612bca565b612bca83838151156140b55781518083602001fd5b806040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a729190614154565b5f5b838110156141035781810151838201526020016140eb565b50505f910152565b5f81518084526141228160208601602086016140e9565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081525f614166602083018461410b565b9392505050565b63ffffffff8116811461417e575f80fd5b50565b80356135f78161416d565b67ffffffffffffffff8116811461417e575f80fd5b80356135f78161418c565b73ffffffffffffffffffffffffffffffffffffffff8116811461417e575f80fd5b80356135f7816141ac565b5f8083601f8401126141e8575f80fd5b50813567ffffffffffffffff8111156141ff575f80fd5b602083019150836020828501011115614216575f80fd5b9250929050565b5f805f805f805f8060c0898b031215614234575f80fd5b883567ffffffffffffffff8082111561424b575f80fd5b818b0191508b601f83011261425e575f80fd5b81358181111561426c575f80fd5b8c60208260071b8501011115614280575f80fd5b602083019a508099505061429660208c01614181565b97506142a460408c016141a1565b965060608b013595506142b960808c016141cd565b945060a08b01359150808211156142ce575f80fd5b506142db8b828c016141d8565b999c989b5096995094979396929594505050565b801515811461417e575f80fd5b5f6020828403121561430c575f80fd5b8135614166816142ef565b5f805f60608486031215614329575f80fd5b83356143348161418c565b925060208401359150604084013561434b816141ac565b809150509250925092565b5f60208284031215614366575f80fd5b81356141668161418c565b5f8060408385031215614382575f80fd5b823561438d8161418c565b946020939093013593505050565b5f602082840312156143ab575f80fd5b8135614166816141ac565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6040516080810167ffffffffffffffff81118282101715614406576144066143b6565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715614453576144536143b6565b604052919050565b5f67ffffffffffffffff821115614474576144746143b6565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b5f82601f8301126144af575f80fd5b81356144c26144bd8261445b565b61440c565b8181528460208386010111156144d6575f80fd5b816020850160208301375f918101602001919091529392505050565b5f805f805f8060c08789031215614507575f80fd5b8635614512816141ac565b95506020870135614522816141ac565b945060408701356145328161416d565b93506060870135614542816141ac565b9250608087013567ffffffffffffffff8082111561455e575f80fd5b61456a8a838b016144a0565b935060a089013591508082111561457f575f80fd5b5061458c89828a016144a0565b9150509295509295509295565b5f8083601f8401126145a9575f80fd5b50813567ffffffffffffffff8111156145c0575f80fd5b6020830191508360208260051b8501011115614216575f80fd5b5f80602083850312156145eb575f80fd5b823567ffffffffffffffff811115614601575f80fd5b61460d85828601614599565b90969095509350505050565b5f805f806080858703121561462c575f80fd5b84356146378161416d565b93506020850135614647816141ac565b925060408501356146578161416d565b9150606085013567ffffffffffffffff811115614672575f80fd5b61467e878288016144a0565b91505092959194509250565b5f805f805f8060a0878903121561469f575f80fd5b863567ffffffffffffffff8111156146b5575f80fd5b6146c189828a01614599565b90975095505060208701356146d58161416d565b935060408701356146e58161418c565b92506060870135915060808701356146fc816141ac565b809150509295509295509295565b5f6020828403121561471a575f80fd5b813567ffffffffffffffff811115614730575f80fd5b612bca848285016144a0565b5f805f6040848603121561474e575f80fd5b833567ffffffffffffffff811115614764575f80fd5b614770868287016141d8565b909790965060209590950135949350505050565b5f60208284031215614794575f80fd5b5051919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b67ffffffffffffffff8281168282160390808211156147e9576147e961479b565b5092915050565b5f82614823577f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b500490565b600181811c9082168061483c57607f821691505b602082108103614873577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b8082018082111561488c5761488c61479b565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f608082840312156148cf575f80fd5b6148d76143e3565b823581526020830135602082015260408301356148f38161418c565b60408201526060928301359281019290925250919050565b5f67ffffffffffffffff8083168181036149275761492761479b565b6001019392505050565b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036149615761496161479b565b5060010190565b8181038181111561488c5761488c61479b565b808202811582820484141761488c5761488c61479b565b81835281816020850137505f602082840101525f60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b838152604060208201525f6149f2604083018486614992565b95945050505050565b5f60208284031215614a0b575f80fd5b81516141668161418c565b5f60208284031215614a26575f80fd5b8151614166816142ef565b601f821115613331575f81815260208120601f850160051c81016020861015614a575750805b601f850160051c820191505b81811015612cfd57828155600101614a63565b815167ffffffffffffffff811115614a9057614a906143b6565b614aa481614a9e8454614828565b84614a31565b602080601f831160018114614af6575f8415614ac05750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b178555612cfd565b5f858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b82811015614b4257888601518255948401946001909101908401614b23565b5085821015614b7e57878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b01905550565b606081525f614ba0606083018661410b565b905083602083015273ffffffffffffffffffffffffffffffffffffffff83166040830152949350505050565b67ffffffffffffffff8181168382160190808211156147e9576147e961479b565b5f82357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81833603018112614c1f575f80fd5b9190910192915050565b5f60808236031215614c39575f80fd5b614c416143e3565b823567ffffffffffffffff811115614c57575f80fd5b614c63368286016144a0565b825250602083013560208201526040830135614c7e8161418c565b6040820152606092830135928101929092525090565b5f63ffffffff808916835273ffffffffffffffffffffffffffffffffffffffff8089166020850152818816604085015280871660608501528086166080850152505060c060a0830152614cea60c083018461410b565b98975050505050505050565b61ffff8181168382160190808211156147e9576147e961479b565b5f7fff00000000000000000000000000000000000000000000000000000000000000808a60f81b1683527fffff0000000000000000000000000000000000000000000000000000000000008960f01b1660018401528751614d79816003860160208c016140e9565b80840190507fffffffffffffffffffffffffffffffffffffffff0000000000000000000000008860601b1660038201528651614dbc816017840160208b016140e9565b808201915050818660f81b16601782015284519150614de28260188301602088016140e9565b016018019998505050505050505050565b7fff000000000000000000000000000000000000000000000000000000000000008860f81b1681525f7fffff000000000000000000000000000000000000000000000000000000000000808960f01b1660018401528751614e5b816003860160208c016140e9565b80840190507fffffffffffffffffffffffffffffffffffffffff0000000000000000000000008860601b1660038201528651614e9e816017840160208b016140e9565b808201915050818660f01b16601782015284519150614ec48260198301602088016140e9565b016019019998505050505050505050565b5f8651614ee6818460208b016140e9565b9190910194855250602084019290925260f81b7fff000000000000000000000000000000000000000000000000000000000000009081166040840152166041820152604201919050565b818382375f9101908152919050565b84815273ffffffffffffffffffffffffffffffffffffffff84166020820152606060408201525f614f74606083018486614992565b9695505050505050565b5f60208284031215614f8e575f80fd5b815167ffffffffffffffff811115614fa4575f80fd5b8201601f81018413614fb4575f80fd5b8051614fc26144bd8261445b565b818152856020838501011115614fd6575f80fd5b6149f28260208301602086016140e9565b5f8060408385031215614ff8575f80fd5b82516150038161416d565b6020840151909250615014816141ac565b809150509250929050565b5f8251614c1f8184602087016140e956fea264697066735822122062954baf90cdfaf5e95bb7f8074dd8342e9465aa1c72cb3ea96bf0b7d02b417f64736f6c63430008140033",
}

// PolygonvalidiumXlayerABI is the input ABI used to generate the binding from.
// Deprecated: Use PolygonvalidiumXlayerMetaData.ABI instead.
var PolygonvalidiumXlayerABI = PolygonvalidiumXlayerMetaData.ABI

// PolygonvalidiumXlayerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PolygonvalidiumXlayerMetaData.Bin instead.
var PolygonvalidiumXlayerBin = PolygonvalidiumXlayerMetaData.Bin

// DeployPolygonvalidiumXlayer deploys a new Ethereum contract, binding an instance of PolygonvalidiumXlayer to it.
func DeployPolygonvalidiumXlayer(auth *bind.TransactOpts, backend bind.ContractBackend, _globalExitRootManager common.Address, _pol common.Address, _bridgeAddress common.Address, _rollupManager common.Address) (common.Address, *types.Transaction, *PolygonvalidiumXlayer, error) {
	parsed, err := PolygonvalidiumXlayerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PolygonvalidiumXlayerBin), backend, _globalExitRootManager, _pol, _bridgeAddress, _rollupManager)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PolygonvalidiumXlayer{PolygonvalidiumXlayerCaller: PolygonvalidiumXlayerCaller{contract: contract}, PolygonvalidiumXlayerTransactor: PolygonvalidiumXlayerTransactor{contract: contract}, PolygonvalidiumXlayerFilterer: PolygonvalidiumXlayerFilterer{contract: contract}}, nil
}

// PolygonvalidiumXlayer is an auto generated Go binding around an Ethereum contract.
type PolygonvalidiumXlayer struct {
	PolygonvalidiumXlayerCaller     // Read-only binding to the contract
	PolygonvalidiumXlayerTransactor // Write-only binding to the contract
	PolygonvalidiumXlayerFilterer   // Log filterer for contract events
}

// PolygonvalidiumXlayerCaller is an auto generated read-only Go binding around an Ethereum contract.
type PolygonvalidiumXlayerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonvalidiumXlayerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PolygonvalidiumXlayerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonvalidiumXlayerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PolygonvalidiumXlayerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonvalidiumXlayerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PolygonvalidiumXlayerSession struct {
	Contract     *PolygonvalidiumXlayer // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PolygonvalidiumXlayerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PolygonvalidiumXlayerCallerSession struct {
	Contract *PolygonvalidiumXlayerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// PolygonvalidiumXlayerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PolygonvalidiumXlayerTransactorSession struct {
	Contract     *PolygonvalidiumXlayerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// PolygonvalidiumXlayerRaw is an auto generated low-level Go binding around an Ethereum contract.
type PolygonvalidiumXlayerRaw struct {
	Contract *PolygonvalidiumXlayer // Generic contract binding to access the raw methods on
}

// PolygonvalidiumXlayerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PolygonvalidiumXlayerCallerRaw struct {
	Contract *PolygonvalidiumXlayerCaller // Generic read-only contract binding to access the raw methods on
}

// PolygonvalidiumXlayerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PolygonvalidiumXlayerTransactorRaw struct {
	Contract *PolygonvalidiumXlayerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPolygonvalidiumXlayer creates a new instance of PolygonvalidiumXlayer, bound to a specific deployed contract.
func NewPolygonvalidiumXlayer(address common.Address, backend bind.ContractBackend) (*PolygonvalidiumXlayer, error) {
	contract, err := bindPolygonvalidiumXlayer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumXlayer{PolygonvalidiumXlayerCaller: PolygonvalidiumXlayerCaller{contract: contract}, PolygonvalidiumXlayerTransactor: PolygonvalidiumXlayerTransactor{contract: contract}, PolygonvalidiumXlayerFilterer: PolygonvalidiumXlayerFilterer{contract: contract}}, nil
}

// NewPolygonvalidiumXlayerCaller creates a new read-only instance of PolygonvalidiumXlayer, bound to a specific deployed contract.
func NewPolygonvalidiumXlayerCaller(address common.Address, caller bind.ContractCaller) (*PolygonvalidiumXlayerCaller, error) {
	contract, err := bindPolygonvalidiumXlayer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumXlayerCaller{contract: contract}, nil
}

// NewPolygonvalidiumXlayerTransactor creates a new write-only instance of PolygonvalidiumXlayer, bound to a specific deployed contract.
func NewPolygonvalidiumXlayerTransactor(address common.Address, transactor bind.ContractTransactor) (*PolygonvalidiumXlayerTransactor, error) {
	contract, err := bindPolygonvalidiumXlayer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumXlayerTransactor{contract: contract}, nil
}

// NewPolygonvalidiumXlayerFilterer creates a new log filterer instance of PolygonvalidiumXlayer, bound to a specific deployed contract.
func NewPolygonvalidiumXlayerFilterer(address common.Address, filterer bind.ContractFilterer) (*PolygonvalidiumXlayerFilterer, error) {
	contract, err := bindPolygonvalidiumXlayer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumXlayerFilterer{contract: contract}, nil
}

// bindPolygonvalidiumXlayer binds a generic wrapper to an already deployed contract.
func bindPolygonvalidiumXlayer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PolygonvalidiumXlayerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PolygonvalidiumXlayer.Contract.PolygonvalidiumXlayerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PolygonvalidiumXlayer.Contract.PolygonvalidiumXlayerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PolygonvalidiumXlayer.Contract.PolygonvalidiumXlayerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PolygonvalidiumXlayer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PolygonvalidiumXlayer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PolygonvalidiumXlayer.Contract.contract.Transact(opts, method, params...)
}

// GLOBALEXITROOTMANAGERL2 is a free data retrieval call binding the contract method 0x9e001877.
//
// Solidity: function GLOBAL_EXIT_ROOT_MANAGER_L2() view returns(address)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCaller) GLOBALEXITROOTMANAGERL2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PolygonvalidiumXlayer.contract.Call(opts, &out, "GLOBAL_EXIT_ROOT_MANAGER_L2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GLOBALEXITROOTMANAGERL2 is a free data retrieval call binding the contract method 0x9e001877.
//
// Solidity: function GLOBAL_EXIT_ROOT_MANAGER_L2() view returns(address)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerSession) GLOBALEXITROOTMANAGERL2() (common.Address, error) {
	return _PolygonvalidiumXlayer.Contract.GLOBALEXITROOTMANAGERL2(&_PolygonvalidiumXlayer.CallOpts)
}

// GLOBALEXITROOTMANAGERL2 is a free data retrieval call binding the contract method 0x9e001877.
//
// Solidity: function GLOBAL_EXIT_ROOT_MANAGER_L2() view returns(address)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCallerSession) GLOBALEXITROOTMANAGERL2() (common.Address, error) {
	return _PolygonvalidiumXlayer.Contract.GLOBALEXITROOTMANAGERL2(&_PolygonvalidiumXlayer.CallOpts)
}

// INITIALIZETXBRIDGELISTLENLEN is a free data retrieval call binding the contract method 0x11e892d4.
//
// Solidity: function INITIALIZE_TX_BRIDGE_LIST_LEN_LEN() view returns(uint8)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCaller) INITIALIZETXBRIDGELISTLENLEN(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _PolygonvalidiumXlayer.contract.Call(opts, &out, "INITIALIZE_TX_BRIDGE_LIST_LEN_LEN")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// INITIALIZETXBRIDGELISTLENLEN is a free data retrieval call binding the contract method 0x11e892d4.
//
// Solidity: function INITIALIZE_TX_BRIDGE_LIST_LEN_LEN() view returns(uint8)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerSession) INITIALIZETXBRIDGELISTLENLEN() (uint8, error) {
	return _PolygonvalidiumXlayer.Contract.INITIALIZETXBRIDGELISTLENLEN(&_PolygonvalidiumXlayer.CallOpts)
}

// INITIALIZETXBRIDGELISTLENLEN is a free data retrieval call binding the contract method 0x11e892d4.
//
// Solidity: function INITIALIZE_TX_BRIDGE_LIST_LEN_LEN() view returns(uint8)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCallerSession) INITIALIZETXBRIDGELISTLENLEN() (uint8, error) {
	return _PolygonvalidiumXlayer.Contract.INITIALIZETXBRIDGELISTLENLEN(&_PolygonvalidiumXlayer.CallOpts)
}

// INITIALIZETXBRIDGEPARAMS is a free data retrieval call binding the contract method 0x05835f37.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS() view returns(bytes)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCaller) INITIALIZETXBRIDGEPARAMS(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _PolygonvalidiumXlayer.contract.Call(opts, &out, "INITIALIZE_TX_BRIDGE_PARAMS")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// INITIALIZETXBRIDGEPARAMS is a free data retrieval call binding the contract method 0x05835f37.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS() view returns(bytes)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerSession) INITIALIZETXBRIDGEPARAMS() ([]byte, error) {
	return _PolygonvalidiumXlayer.Contract.INITIALIZETXBRIDGEPARAMS(&_PolygonvalidiumXlayer.CallOpts)
}

// INITIALIZETXBRIDGEPARAMS is a free data retrieval call binding the contract method 0x05835f37.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS() view returns(bytes)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCallerSession) INITIALIZETXBRIDGEPARAMS() ([]byte, error) {
	return _PolygonvalidiumXlayer.Contract.INITIALIZETXBRIDGEPARAMS(&_PolygonvalidiumXlayer.CallOpts)
}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS is a free data retrieval call binding the contract method 0x7a5460c5.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS() view returns(bytes)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCaller) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _PolygonvalidiumXlayer.contract.Call(opts, &out, "INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS is a free data retrieval call binding the contract method 0x7a5460c5.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS() view returns(bytes)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerSession) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS() ([]byte, error) {
	return _PolygonvalidiumXlayer.Contract.INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS(&_PolygonvalidiumXlayer.CallOpts)
}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS is a free data retrieval call binding the contract method 0x7a5460c5.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS() view returns(bytes)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCallerSession) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS() ([]byte, error) {
	return _PolygonvalidiumXlayer.Contract.INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS(&_PolygonvalidiumXlayer.CallOpts)
}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA is a free data retrieval call binding the contract method 0x52bdeb6d.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA() view returns(bytes)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCaller) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _PolygonvalidiumXlayer.contract.Call(opts, &out, "INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA is a free data retrieval call binding the contract method 0x52bdeb6d.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA() view returns(bytes)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerSession) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA() ([]byte, error) {
	return _PolygonvalidiumXlayer.Contract.INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA(&_PolygonvalidiumXlayer.CallOpts)
}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA is a free data retrieval call binding the contract method 0x52bdeb6d.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA() view returns(bytes)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCallerSession) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA() ([]byte, error) {
	return _PolygonvalidiumXlayer.Contract.INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA(&_PolygonvalidiumXlayer.CallOpts)
}

// INITIALIZETXCONSTANTBYTES is a free data retrieval call binding the contract method 0x03508963.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES() view returns(uint16)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCaller) INITIALIZETXCONSTANTBYTES(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _PolygonvalidiumXlayer.contract.Call(opts, &out, "INITIALIZE_TX_CONSTANT_BYTES")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// INITIALIZETXCONSTANTBYTES is a free data retrieval call binding the contract method 0x03508963.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES() view returns(uint16)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerSession) INITIALIZETXCONSTANTBYTES() (uint16, error) {
	return _PolygonvalidiumXlayer.Contract.INITIALIZETXCONSTANTBYTES(&_PolygonvalidiumXlayer.CallOpts)
}

// INITIALIZETXCONSTANTBYTES is a free data retrieval call binding the contract method 0x03508963.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES() view returns(uint16)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCallerSession) INITIALIZETXCONSTANTBYTES() (uint16, error) {
	return _PolygonvalidiumXlayer.Contract.INITIALIZETXCONSTANTBYTES(&_PolygonvalidiumXlayer.CallOpts)
}

// INITIALIZETXCONSTANTBYTESEMPTYMETADATA is a free data retrieval call binding the contract method 0x676870d2.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA() view returns(uint16)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCaller) INITIALIZETXCONSTANTBYTESEMPTYMETADATA(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _PolygonvalidiumXlayer.contract.Call(opts, &out, "INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// INITIALIZETXCONSTANTBYTESEMPTYMETADATA is a free data retrieval call binding the contract method 0x676870d2.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA() view returns(uint16)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerSession) INITIALIZETXCONSTANTBYTESEMPTYMETADATA() (uint16, error) {
	return _PolygonvalidiumXlayer.Contract.INITIALIZETXCONSTANTBYTESEMPTYMETADATA(&_PolygonvalidiumXlayer.CallOpts)
}

// INITIALIZETXCONSTANTBYTESEMPTYMETADATA is a free data retrieval call binding the contract method 0x676870d2.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA() view returns(uint16)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCallerSession) INITIALIZETXCONSTANTBYTESEMPTYMETADATA() (uint16, error) {
	return _PolygonvalidiumXlayer.Contract.INITIALIZETXCONSTANTBYTESEMPTYMETADATA(&_PolygonvalidiumXlayer.CallOpts)
}

// INITIALIZETXDATALENEMPTYMETADATA is a free data retrieval call binding the contract method 0xc7fffd4b.
//
// Solidity: function INITIALIZE_TX_DATA_LEN_EMPTY_METADATA() view returns(uint8)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCaller) INITIALIZETXDATALENEMPTYMETADATA(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _PolygonvalidiumXlayer.contract.Call(opts, &out, "INITIALIZE_TX_DATA_LEN_EMPTY_METADATA")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// INITIALIZETXDATALENEMPTYMETADATA is a free data retrieval call binding the contract method 0xc7fffd4b.
//
// Solidity: function INITIALIZE_TX_DATA_LEN_EMPTY_METADATA() view returns(uint8)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerSession) INITIALIZETXDATALENEMPTYMETADATA() (uint8, error) {
	return _PolygonvalidiumXlayer.Contract.INITIALIZETXDATALENEMPTYMETADATA(&_PolygonvalidiumXlayer.CallOpts)
}

// INITIALIZETXDATALENEMPTYMETADATA is a free data retrieval call binding the contract method 0xc7fffd4b.
//
// Solidity: function INITIALIZE_TX_DATA_LEN_EMPTY_METADATA() view returns(uint8)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCallerSession) INITIALIZETXDATALENEMPTYMETADATA() (uint8, error) {
	return _PolygonvalidiumXlayer.Contract.INITIALIZETXDATALENEMPTYMETADATA(&_PolygonvalidiumXlayer.CallOpts)
}

// INITIALIZETXEFFECTIVEPERCENTAGE is a free data retrieval call binding the contract method 0x40b5de6c.
//
// Solidity: function INITIALIZE_TX_EFFECTIVE_PERCENTAGE() view returns(bytes1)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCaller) INITIALIZETXEFFECTIVEPERCENTAGE(opts *bind.CallOpts) ([1]byte, error) {
	var out []interface{}
	err := _PolygonvalidiumXlayer.contract.Call(opts, &out, "INITIALIZE_TX_EFFECTIVE_PERCENTAGE")

	if err != nil {
		return *new([1]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)

	return out0, err

}

// INITIALIZETXEFFECTIVEPERCENTAGE is a free data retrieval call binding the contract method 0x40b5de6c.
//
// Solidity: function INITIALIZE_TX_EFFECTIVE_PERCENTAGE() view returns(bytes1)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerSession) INITIALIZETXEFFECTIVEPERCENTAGE() ([1]byte, error) {
	return _PolygonvalidiumXlayer.Contract.INITIALIZETXEFFECTIVEPERCENTAGE(&_PolygonvalidiumXlayer.CallOpts)
}

// INITIALIZETXEFFECTIVEPERCENTAGE is a free data retrieval call binding the contract method 0x40b5de6c.
//
// Solidity: function INITIALIZE_TX_EFFECTIVE_PERCENTAGE() view returns(bytes1)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCallerSession) INITIALIZETXEFFECTIVEPERCENTAGE() ([1]byte, error) {
	return _PolygonvalidiumXlayer.Contract.INITIALIZETXEFFECTIVEPERCENTAGE(&_PolygonvalidiumXlayer.CallOpts)
}

// SIGNATUREINITIALIZETXR is a free data retrieval call binding the contract method 0xb0afe154.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_R() view returns(bytes32)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCaller) SIGNATUREINITIALIZETXR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PolygonvalidiumXlayer.contract.Call(opts, &out, "SIGNATURE_INITIALIZE_TX_R")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SIGNATUREINITIALIZETXR is a free data retrieval call binding the contract method 0xb0afe154.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_R() view returns(bytes32)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerSession) SIGNATUREINITIALIZETXR() ([32]byte, error) {
	return _PolygonvalidiumXlayer.Contract.SIGNATUREINITIALIZETXR(&_PolygonvalidiumXlayer.CallOpts)
}

// SIGNATUREINITIALIZETXR is a free data retrieval call binding the contract method 0xb0afe154.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_R() view returns(bytes32)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCallerSession) SIGNATUREINITIALIZETXR() ([32]byte, error) {
	return _PolygonvalidiumXlayer.Contract.SIGNATUREINITIALIZETXR(&_PolygonvalidiumXlayer.CallOpts)
}

// SIGNATUREINITIALIZETXS is a free data retrieval call binding the contract method 0xd7bc90ff.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_S() view returns(bytes32)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCaller) SIGNATUREINITIALIZETXS(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PolygonvalidiumXlayer.contract.Call(opts, &out, "SIGNATURE_INITIALIZE_TX_S")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SIGNATUREINITIALIZETXS is a free data retrieval call binding the contract method 0xd7bc90ff.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_S() view returns(bytes32)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerSession) SIGNATUREINITIALIZETXS() ([32]byte, error) {
	return _PolygonvalidiumXlayer.Contract.SIGNATUREINITIALIZETXS(&_PolygonvalidiumXlayer.CallOpts)
}

// SIGNATUREINITIALIZETXS is a free data retrieval call binding the contract method 0xd7bc90ff.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_S() view returns(bytes32)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCallerSession) SIGNATUREINITIALIZETXS() ([32]byte, error) {
	return _PolygonvalidiumXlayer.Contract.SIGNATUREINITIALIZETXS(&_PolygonvalidiumXlayer.CallOpts)
}

// SIGNATUREINITIALIZETXV is a free data retrieval call binding the contract method 0xf35dda47.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_V() view returns(uint8)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCaller) SIGNATUREINITIALIZETXV(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _PolygonvalidiumXlayer.contract.Call(opts, &out, "SIGNATURE_INITIALIZE_TX_V")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// SIGNATUREINITIALIZETXV is a free data retrieval call binding the contract method 0xf35dda47.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_V() view returns(uint8)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerSession) SIGNATUREINITIALIZETXV() (uint8, error) {
	return _PolygonvalidiumXlayer.Contract.SIGNATUREINITIALIZETXV(&_PolygonvalidiumXlayer.CallOpts)
}

// SIGNATUREINITIALIZETXV is a free data retrieval call binding the contract method 0xf35dda47.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_V() view returns(uint8)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCallerSession) SIGNATUREINITIALIZETXV() (uint8, error) {
	return _PolygonvalidiumXlayer.Contract.SIGNATUREINITIALIZETXV(&_PolygonvalidiumXlayer.CallOpts)
}

// TIMESTAMPRANGE is a free data retrieval call binding the contract method 0x42308fab.
//
// Solidity: function TIMESTAMP_RANGE() view returns(uint256)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCaller) TIMESTAMPRANGE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PolygonvalidiumXlayer.contract.Call(opts, &out, "TIMESTAMP_RANGE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TIMESTAMPRANGE is a free data retrieval call binding the contract method 0x42308fab.
//
// Solidity: function TIMESTAMP_RANGE() view returns(uint256)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerSession) TIMESTAMPRANGE() (*big.Int, error) {
	return _PolygonvalidiumXlayer.Contract.TIMESTAMPRANGE(&_PolygonvalidiumXlayer.CallOpts)
}

// TIMESTAMPRANGE is a free data retrieval call binding the contract method 0x42308fab.
//
// Solidity: function TIMESTAMP_RANGE() view returns(uint256)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCallerSession) TIMESTAMPRANGE() (*big.Int, error) {
	return _PolygonvalidiumXlayer.Contract.TIMESTAMPRANGE(&_PolygonvalidiumXlayer.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PolygonvalidiumXlayer.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerSession) Admin() (common.Address, error) {
	return _PolygonvalidiumXlayer.Contract.Admin(&_PolygonvalidiumXlayer.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCallerSession) Admin() (common.Address, error) {
	return _PolygonvalidiumXlayer.Contract.Admin(&_PolygonvalidiumXlayer.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PolygonvalidiumXlayer.contract.Call(opts, &out, "bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerSession) BridgeAddress() (common.Address, error) {
	return _PolygonvalidiumXlayer.Contract.BridgeAddress(&_PolygonvalidiumXlayer.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCallerSession) BridgeAddress() (common.Address, error) {
	return _PolygonvalidiumXlayer.Contract.BridgeAddress(&_PolygonvalidiumXlayer.CallOpts)
}

// CalculatePolPerForceBatch is a free data retrieval call binding the contract method 0x00d0295d.
//
// Solidity: function calculatePolPerForceBatch() view returns(uint256)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCaller) CalculatePolPerForceBatch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PolygonvalidiumXlayer.contract.Call(opts, &out, "calculatePolPerForceBatch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculatePolPerForceBatch is a free data retrieval call binding the contract method 0x00d0295d.
//
// Solidity: function calculatePolPerForceBatch() view returns(uint256)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerSession) CalculatePolPerForceBatch() (*big.Int, error) {
	return _PolygonvalidiumXlayer.Contract.CalculatePolPerForceBatch(&_PolygonvalidiumXlayer.CallOpts)
}

// CalculatePolPerForceBatch is a free data retrieval call binding the contract method 0x00d0295d.
//
// Solidity: function calculatePolPerForceBatch() view returns(uint256)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCallerSession) CalculatePolPerForceBatch() (*big.Int, error) {
	return _PolygonvalidiumXlayer.Contract.CalculatePolPerForceBatch(&_PolygonvalidiumXlayer.CallOpts)
}

// DataAvailabilityProtocol is a free data retrieval call binding the contract method 0xe57a0b4c.
//
// Solidity: function dataAvailabilityProtocol() view returns(address)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCaller) DataAvailabilityProtocol(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PolygonvalidiumXlayer.contract.Call(opts, &out, "dataAvailabilityProtocol")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DataAvailabilityProtocol is a free data retrieval call binding the contract method 0xe57a0b4c.
//
// Solidity: function dataAvailabilityProtocol() view returns(address)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerSession) DataAvailabilityProtocol() (common.Address, error) {
	return _PolygonvalidiumXlayer.Contract.DataAvailabilityProtocol(&_PolygonvalidiumXlayer.CallOpts)
}

// DataAvailabilityProtocol is a free data retrieval call binding the contract method 0xe57a0b4c.
//
// Solidity: function dataAvailabilityProtocol() view returns(address)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCallerSession) DataAvailabilityProtocol() (common.Address, error) {
	return _PolygonvalidiumXlayer.Contract.DataAvailabilityProtocol(&_PolygonvalidiumXlayer.CallOpts)
}

// ForceBatchAddress is a free data retrieval call binding the contract method 0x2c111c06.
//
// Solidity: function forceBatchAddress() view returns(address)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCaller) ForceBatchAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PolygonvalidiumXlayer.contract.Call(opts, &out, "forceBatchAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ForceBatchAddress is a free data retrieval call binding the contract method 0x2c111c06.
//
// Solidity: function forceBatchAddress() view returns(address)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerSession) ForceBatchAddress() (common.Address, error) {
	return _PolygonvalidiumXlayer.Contract.ForceBatchAddress(&_PolygonvalidiumXlayer.CallOpts)
}

// ForceBatchAddress is a free data retrieval call binding the contract method 0x2c111c06.
//
// Solidity: function forceBatchAddress() view returns(address)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCallerSession) ForceBatchAddress() (common.Address, error) {
	return _PolygonvalidiumXlayer.Contract.ForceBatchAddress(&_PolygonvalidiumXlayer.CallOpts)
}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCaller) ForceBatchTimeout(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _PolygonvalidiumXlayer.contract.Call(opts, &out, "forceBatchTimeout")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerSession) ForceBatchTimeout() (uint64, error) {
	return _PolygonvalidiumXlayer.Contract.ForceBatchTimeout(&_PolygonvalidiumXlayer.CallOpts)
}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCallerSession) ForceBatchTimeout() (uint64, error) {
	return _PolygonvalidiumXlayer.Contract.ForceBatchTimeout(&_PolygonvalidiumXlayer.CallOpts)
}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCaller) ForcedBatches(opts *bind.CallOpts, arg0 uint64) ([32]byte, error) {
	var out []interface{}
	err := _PolygonvalidiumXlayer.contract.Call(opts, &out, "forcedBatches", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerSession) ForcedBatches(arg0 uint64) ([32]byte, error) {
	return _PolygonvalidiumXlayer.Contract.ForcedBatches(&_PolygonvalidiumXlayer.CallOpts, arg0)
}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCallerSession) ForcedBatches(arg0 uint64) ([32]byte, error) {
	return _PolygonvalidiumXlayer.Contract.ForcedBatches(&_PolygonvalidiumXlayer.CallOpts, arg0)
}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCaller) GasTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PolygonvalidiumXlayer.contract.Call(opts, &out, "gasTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerSession) GasTokenAddress() (common.Address, error) {
	return _PolygonvalidiumXlayer.Contract.GasTokenAddress(&_PolygonvalidiumXlayer.CallOpts)
}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCallerSession) GasTokenAddress() (common.Address, error) {
	return _PolygonvalidiumXlayer.Contract.GasTokenAddress(&_PolygonvalidiumXlayer.CallOpts)
}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCaller) GasTokenNetwork(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _PolygonvalidiumXlayer.contract.Call(opts, &out, "gasTokenNetwork")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerSession) GasTokenNetwork() (uint32, error) {
	return _PolygonvalidiumXlayer.Contract.GasTokenNetwork(&_PolygonvalidiumXlayer.CallOpts)
}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCallerSession) GasTokenNetwork() (uint32, error) {
	return _PolygonvalidiumXlayer.Contract.GasTokenNetwork(&_PolygonvalidiumXlayer.CallOpts)
}

// GenerateInitializeTransaction is a free data retrieval call binding the contract method 0xa652f26c.
//
// Solidity: function generateInitializeTransaction(uint32 networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, bytes _gasTokenMetadata) view returns(bytes)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCaller) GenerateInitializeTransaction(opts *bind.CallOpts, networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _gasTokenMetadata []byte) ([]byte, error) {
	var out []interface{}
	err := _PolygonvalidiumXlayer.contract.Call(opts, &out, "generateInitializeTransaction", networkID, _gasTokenAddress, _gasTokenNetwork, _gasTokenMetadata)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GenerateInitializeTransaction is a free data retrieval call binding the contract method 0xa652f26c.
//
// Solidity: function generateInitializeTransaction(uint32 networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, bytes _gasTokenMetadata) view returns(bytes)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerSession) GenerateInitializeTransaction(networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _gasTokenMetadata []byte) ([]byte, error) {
	return _PolygonvalidiumXlayer.Contract.GenerateInitializeTransaction(&_PolygonvalidiumXlayer.CallOpts, networkID, _gasTokenAddress, _gasTokenNetwork, _gasTokenMetadata)
}

// GenerateInitializeTransaction is a free data retrieval call binding the contract method 0xa652f26c.
//
// Solidity: function generateInitializeTransaction(uint32 networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, bytes _gasTokenMetadata) view returns(bytes)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCallerSession) GenerateInitializeTransaction(networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _gasTokenMetadata []byte) ([]byte, error) {
	return _PolygonvalidiumXlayer.Contract.GenerateInitializeTransaction(&_PolygonvalidiumXlayer.CallOpts, networkID, _gasTokenAddress, _gasTokenNetwork, _gasTokenMetadata)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCaller) GlobalExitRootManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PolygonvalidiumXlayer.contract.Call(opts, &out, "globalExitRootManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerSession) GlobalExitRootManager() (common.Address, error) {
	return _PolygonvalidiumXlayer.Contract.GlobalExitRootManager(&_PolygonvalidiumXlayer.CallOpts)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCallerSession) GlobalExitRootManager() (common.Address, error) {
	return _PolygonvalidiumXlayer.Contract.GlobalExitRootManager(&_PolygonvalidiumXlayer.CallOpts)
}

// IsSequenceWithDataAvailabilityAllowed is a free data retrieval call binding the contract method 0x4c21fef3.
//
// Solidity: function isSequenceWithDataAvailabilityAllowed() view returns(bool)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCaller) IsSequenceWithDataAvailabilityAllowed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PolygonvalidiumXlayer.contract.Call(opts, &out, "isSequenceWithDataAvailabilityAllowed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSequenceWithDataAvailabilityAllowed is a free data retrieval call binding the contract method 0x4c21fef3.
//
// Solidity: function isSequenceWithDataAvailabilityAllowed() view returns(bool)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerSession) IsSequenceWithDataAvailabilityAllowed() (bool, error) {
	return _PolygonvalidiumXlayer.Contract.IsSequenceWithDataAvailabilityAllowed(&_PolygonvalidiumXlayer.CallOpts)
}

// IsSequenceWithDataAvailabilityAllowed is a free data retrieval call binding the contract method 0x4c21fef3.
//
// Solidity: function isSequenceWithDataAvailabilityAllowed() view returns(bool)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCallerSession) IsSequenceWithDataAvailabilityAllowed() (bool, error) {
	return _PolygonvalidiumXlayer.Contract.IsSequenceWithDataAvailabilityAllowed(&_PolygonvalidiumXlayer.CallOpts)
}

// LastAccInputHash is a free data retrieval call binding the contract method 0x6e05d2cd.
//
// Solidity: function lastAccInputHash() view returns(bytes32)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCaller) LastAccInputHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PolygonvalidiumXlayer.contract.Call(opts, &out, "lastAccInputHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LastAccInputHash is a free data retrieval call binding the contract method 0x6e05d2cd.
//
// Solidity: function lastAccInputHash() view returns(bytes32)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerSession) LastAccInputHash() ([32]byte, error) {
	return _PolygonvalidiumXlayer.Contract.LastAccInputHash(&_PolygonvalidiumXlayer.CallOpts)
}

// LastAccInputHash is a free data retrieval call binding the contract method 0x6e05d2cd.
//
// Solidity: function lastAccInputHash() view returns(bytes32)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCallerSession) LastAccInputHash() ([32]byte, error) {
	return _PolygonvalidiumXlayer.Contract.LastAccInputHash(&_PolygonvalidiumXlayer.CallOpts)
}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCaller) LastForceBatch(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _PolygonvalidiumXlayer.contract.Call(opts, &out, "lastForceBatch")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerSession) LastForceBatch() (uint64, error) {
	return _PolygonvalidiumXlayer.Contract.LastForceBatch(&_PolygonvalidiumXlayer.CallOpts)
}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCallerSession) LastForceBatch() (uint64, error) {
	return _PolygonvalidiumXlayer.Contract.LastForceBatch(&_PolygonvalidiumXlayer.CallOpts)
}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCaller) LastForceBatchSequenced(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _PolygonvalidiumXlayer.contract.Call(opts, &out, "lastForceBatchSequenced")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerSession) LastForceBatchSequenced() (uint64, error) {
	return _PolygonvalidiumXlayer.Contract.LastForceBatchSequenced(&_PolygonvalidiumXlayer.CallOpts)
}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCallerSession) LastForceBatchSequenced() (uint64, error) {
	return _PolygonvalidiumXlayer.Contract.LastForceBatchSequenced(&_PolygonvalidiumXlayer.CallOpts)
}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCaller) NetworkName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PolygonvalidiumXlayer.contract.Call(opts, &out, "networkName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerSession) NetworkName() (string, error) {
	return _PolygonvalidiumXlayer.Contract.NetworkName(&_PolygonvalidiumXlayer.CallOpts)
}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCallerSession) NetworkName() (string, error) {
	return _PolygonvalidiumXlayer.Contract.NetworkName(&_PolygonvalidiumXlayer.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PolygonvalidiumXlayer.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerSession) PendingAdmin() (common.Address, error) {
	return _PolygonvalidiumXlayer.Contract.PendingAdmin(&_PolygonvalidiumXlayer.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCallerSession) PendingAdmin() (common.Address, error) {
	return _PolygonvalidiumXlayer.Contract.PendingAdmin(&_PolygonvalidiumXlayer.CallOpts)
}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCaller) Pol(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PolygonvalidiumXlayer.contract.Call(opts, &out, "pol")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerSession) Pol() (common.Address, error) {
	return _PolygonvalidiumXlayer.Contract.Pol(&_PolygonvalidiumXlayer.CallOpts)
}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCallerSession) Pol() (common.Address, error) {
	return _PolygonvalidiumXlayer.Contract.Pol(&_PolygonvalidiumXlayer.CallOpts)
}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCaller) RollupManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PolygonvalidiumXlayer.contract.Call(opts, &out, "rollupManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerSession) RollupManager() (common.Address, error) {
	return _PolygonvalidiumXlayer.Contract.RollupManager(&_PolygonvalidiumXlayer.CallOpts)
}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCallerSession) RollupManager() (common.Address, error) {
	return _PolygonvalidiumXlayer.Contract.RollupManager(&_PolygonvalidiumXlayer.CallOpts)
}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCaller) TrustedSequencer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PolygonvalidiumXlayer.contract.Call(opts, &out, "trustedSequencer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerSession) TrustedSequencer() (common.Address, error) {
	return _PolygonvalidiumXlayer.Contract.TrustedSequencer(&_PolygonvalidiumXlayer.CallOpts)
}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCallerSession) TrustedSequencer() (common.Address, error) {
	return _PolygonvalidiumXlayer.Contract.TrustedSequencer(&_PolygonvalidiumXlayer.CallOpts)
}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCaller) TrustedSequencerURL(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PolygonvalidiumXlayer.contract.Call(opts, &out, "trustedSequencerURL")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerSession) TrustedSequencerURL() (string, error) {
	return _PolygonvalidiumXlayer.Contract.TrustedSequencerURL(&_PolygonvalidiumXlayer.CallOpts)
}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerCallerSession) TrustedSequencerURL() (string, error) {
	return _PolygonvalidiumXlayer.Contract.TrustedSequencerURL(&_PolygonvalidiumXlayer.CallOpts)
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerTransactor) AcceptAdminRole(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PolygonvalidiumXlayer.contract.Transact(opts, "acceptAdminRole")
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerSession) AcceptAdminRole() (*types.Transaction, error) {
	return _PolygonvalidiumXlayer.Contract.AcceptAdminRole(&_PolygonvalidiumXlayer.TransactOpts)
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerTransactorSession) AcceptAdminRole() (*types.Transaction, error) {
	return _PolygonvalidiumXlayer.Contract.AcceptAdminRole(&_PolygonvalidiumXlayer.TransactOpts)
}

// ForceBatch is a paid mutator transaction binding the contract method 0xeaeb077b.
//
// Solidity: function forceBatch(bytes transactions, uint256 polAmount) returns()
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerTransactor) ForceBatch(opts *bind.TransactOpts, transactions []byte, polAmount *big.Int) (*types.Transaction, error) {
	return _PolygonvalidiumXlayer.contract.Transact(opts, "forceBatch", transactions, polAmount)
}

// ForceBatch is a paid mutator transaction binding the contract method 0xeaeb077b.
//
// Solidity: function forceBatch(bytes transactions, uint256 polAmount) returns()
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerSession) ForceBatch(transactions []byte, polAmount *big.Int) (*types.Transaction, error) {
	return _PolygonvalidiumXlayer.Contract.ForceBatch(&_PolygonvalidiumXlayer.TransactOpts, transactions, polAmount)
}

// ForceBatch is a paid mutator transaction binding the contract method 0xeaeb077b.
//
// Solidity: function forceBatch(bytes transactions, uint256 polAmount) returns()
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerTransactorSession) ForceBatch(transactions []byte, polAmount *big.Int) (*types.Transaction, error) {
	return _PolygonvalidiumXlayer.Contract.ForceBatch(&_PolygonvalidiumXlayer.TransactOpts, transactions, polAmount)
}

// Initialize is a paid mutator transaction binding the contract method 0x71257022.
//
// Solidity: function initialize(address _admin, address sequencer, uint32 networkID, address _gasTokenAddress, string sequencerURL, string _networkName) returns()
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerTransactor) Initialize(opts *bind.TransactOpts, _admin common.Address, sequencer common.Address, networkID uint32, _gasTokenAddress common.Address, sequencerURL string, _networkName string) (*types.Transaction, error) {
	return _PolygonvalidiumXlayer.contract.Transact(opts, "initialize", _admin, sequencer, networkID, _gasTokenAddress, sequencerURL, _networkName)
}

// Initialize is a paid mutator transaction binding the contract method 0x71257022.
//
// Solidity: function initialize(address _admin, address sequencer, uint32 networkID, address _gasTokenAddress, string sequencerURL, string _networkName) returns()
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerSession) Initialize(_admin common.Address, sequencer common.Address, networkID uint32, _gasTokenAddress common.Address, sequencerURL string, _networkName string) (*types.Transaction, error) {
	return _PolygonvalidiumXlayer.Contract.Initialize(&_PolygonvalidiumXlayer.TransactOpts, _admin, sequencer, networkID, _gasTokenAddress, sequencerURL, _networkName)
}

// Initialize is a paid mutator transaction binding the contract method 0x71257022.
//
// Solidity: function initialize(address _admin, address sequencer, uint32 networkID, address _gasTokenAddress, string sequencerURL, string _networkName) returns()
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerTransactorSession) Initialize(_admin common.Address, sequencer common.Address, networkID uint32, _gasTokenAddress common.Address, sequencerURL string, _networkName string) (*types.Transaction, error) {
	return _PolygonvalidiumXlayer.Contract.Initialize(&_PolygonvalidiumXlayer.TransactOpts, _admin, sequencer, networkID, _gasTokenAddress, sequencerURL, _networkName)
}

// OnVerifyBatches is a paid mutator transaction binding the contract method 0x32c2d153.
//
// Solidity: function onVerifyBatches(uint64 lastVerifiedBatch, bytes32 newStateRoot, address aggregator) returns()
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerTransactor) OnVerifyBatches(opts *bind.TransactOpts, lastVerifiedBatch uint64, newStateRoot [32]byte, aggregator common.Address) (*types.Transaction, error) {
	return _PolygonvalidiumXlayer.contract.Transact(opts, "onVerifyBatches", lastVerifiedBatch, newStateRoot, aggregator)
}

// OnVerifyBatches is a paid mutator transaction binding the contract method 0x32c2d153.
//
// Solidity: function onVerifyBatches(uint64 lastVerifiedBatch, bytes32 newStateRoot, address aggregator) returns()
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerSession) OnVerifyBatches(lastVerifiedBatch uint64, newStateRoot [32]byte, aggregator common.Address) (*types.Transaction, error) {
	return _PolygonvalidiumXlayer.Contract.OnVerifyBatches(&_PolygonvalidiumXlayer.TransactOpts, lastVerifiedBatch, newStateRoot, aggregator)
}

// OnVerifyBatches is a paid mutator transaction binding the contract method 0x32c2d153.
//
// Solidity: function onVerifyBatches(uint64 lastVerifiedBatch, bytes32 newStateRoot, address aggregator) returns()
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerTransactorSession) OnVerifyBatches(lastVerifiedBatch uint64, newStateRoot [32]byte, aggregator common.Address) (*types.Transaction, error) {
	return _PolygonvalidiumXlayer.Contract.OnVerifyBatches(&_PolygonvalidiumXlayer.TransactOpts, lastVerifiedBatch, newStateRoot, aggregator)
}

// RollbackBatches is a paid mutator transaction binding the contract method 0x669adece.
//
// Solidity: function rollbackBatches(uint64 targetBatch, bytes32 accInputHashToRollback) returns()
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerTransactor) RollbackBatches(opts *bind.TransactOpts, targetBatch uint64, accInputHashToRollback [32]byte) (*types.Transaction, error) {
	return _PolygonvalidiumXlayer.contract.Transact(opts, "rollbackBatches", targetBatch, accInputHashToRollback)
}

// RollbackBatches is a paid mutator transaction binding the contract method 0x669adece.
//
// Solidity: function rollbackBatches(uint64 targetBatch, bytes32 accInputHashToRollback) returns()
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerSession) RollbackBatches(targetBatch uint64, accInputHashToRollback [32]byte) (*types.Transaction, error) {
	return _PolygonvalidiumXlayer.Contract.RollbackBatches(&_PolygonvalidiumXlayer.TransactOpts, targetBatch, accInputHashToRollback)
}

// RollbackBatches is a paid mutator transaction binding the contract method 0x669adece.
//
// Solidity: function rollbackBatches(uint64 targetBatch, bytes32 accInputHashToRollback) returns()
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerTransactorSession) RollbackBatches(targetBatch uint64, accInputHashToRollback [32]byte) (*types.Transaction, error) {
	return _PolygonvalidiumXlayer.Contract.RollbackBatches(&_PolygonvalidiumXlayer.TransactOpts, targetBatch, accInputHashToRollback)
}

// SequenceBatches is a paid mutator transaction binding the contract method 0xb910e0f9.
//
// Solidity: function sequenceBatches((bytes,bytes32,uint64,bytes32)[] batches, uint32 indexL1InfoRoot, uint64 maxSequenceTimestamp, bytes32 expectedFinalAccInputHash, address l2Coinbase) returns()
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerTransactor) SequenceBatches(opts *bind.TransactOpts, batches []PolygonRollupBaseEtrogBatchData, indexL1InfoRoot uint32, maxSequenceTimestamp uint64, expectedFinalAccInputHash [32]byte, l2Coinbase common.Address) (*types.Transaction, error) {
	return _PolygonvalidiumXlayer.contract.Transact(opts, "sequenceBatches", batches, indexL1InfoRoot, maxSequenceTimestamp, expectedFinalAccInputHash, l2Coinbase)
}

// SequenceBatches is a paid mutator transaction binding the contract method 0xb910e0f9.
//
// Solidity: function sequenceBatches((bytes,bytes32,uint64,bytes32)[] batches, uint32 indexL1InfoRoot, uint64 maxSequenceTimestamp, bytes32 expectedFinalAccInputHash, address l2Coinbase) returns()
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerSession) SequenceBatches(batches []PolygonRollupBaseEtrogBatchData, indexL1InfoRoot uint32, maxSequenceTimestamp uint64, expectedFinalAccInputHash [32]byte, l2Coinbase common.Address) (*types.Transaction, error) {
	return _PolygonvalidiumXlayer.Contract.SequenceBatches(&_PolygonvalidiumXlayer.TransactOpts, batches, indexL1InfoRoot, maxSequenceTimestamp, expectedFinalAccInputHash, l2Coinbase)
}

// SequenceBatches is a paid mutator transaction binding the contract method 0xb910e0f9.
//
// Solidity: function sequenceBatches((bytes,bytes32,uint64,bytes32)[] batches, uint32 indexL1InfoRoot, uint64 maxSequenceTimestamp, bytes32 expectedFinalAccInputHash, address l2Coinbase) returns()
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerTransactorSession) SequenceBatches(batches []PolygonRollupBaseEtrogBatchData, indexL1InfoRoot uint32, maxSequenceTimestamp uint64, expectedFinalAccInputHash [32]byte, l2Coinbase common.Address) (*types.Transaction, error) {
	return _PolygonvalidiumXlayer.Contract.SequenceBatches(&_PolygonvalidiumXlayer.TransactOpts, batches, indexL1InfoRoot, maxSequenceTimestamp, expectedFinalAccInputHash, l2Coinbase)
}

// SequenceBatchesValidium is a paid mutator transaction binding the contract method 0x165e8a8d.
//
// Solidity: function sequenceBatchesValidium((bytes32,bytes32,uint64,bytes32)[] batches, uint32 indexL1InfoRoot, uint64 maxSequenceTimestamp, bytes32 expectedFinalAccInputHash, address l2Coinbase, bytes dataAvailabilityMessage) returns()
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerTransactor) SequenceBatchesValidium(opts *bind.TransactOpts, batches []PolygonValidiumEtrogValidiumBatchData, indexL1InfoRoot uint32, maxSequenceTimestamp uint64, expectedFinalAccInputHash [32]byte, l2Coinbase common.Address, dataAvailabilityMessage []byte) (*types.Transaction, error) {
	return _PolygonvalidiumXlayer.contract.Transact(opts, "sequenceBatchesValidium", batches, indexL1InfoRoot, maxSequenceTimestamp, expectedFinalAccInputHash, l2Coinbase, dataAvailabilityMessage)
}

// SequenceBatchesValidium is a paid mutator transaction binding the contract method 0x165e8a8d.
//
// Solidity: function sequenceBatchesValidium((bytes32,bytes32,uint64,bytes32)[] batches, uint32 indexL1InfoRoot, uint64 maxSequenceTimestamp, bytes32 expectedFinalAccInputHash, address l2Coinbase, bytes dataAvailabilityMessage) returns()
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerSession) SequenceBatchesValidium(batches []PolygonValidiumEtrogValidiumBatchData, indexL1InfoRoot uint32, maxSequenceTimestamp uint64, expectedFinalAccInputHash [32]byte, l2Coinbase common.Address, dataAvailabilityMessage []byte) (*types.Transaction, error) {
	return _PolygonvalidiumXlayer.Contract.SequenceBatchesValidium(&_PolygonvalidiumXlayer.TransactOpts, batches, indexL1InfoRoot, maxSequenceTimestamp, expectedFinalAccInputHash, l2Coinbase, dataAvailabilityMessage)
}

// SequenceBatchesValidium is a paid mutator transaction binding the contract method 0x165e8a8d.
//
// Solidity: function sequenceBatchesValidium((bytes32,bytes32,uint64,bytes32)[] batches, uint32 indexL1InfoRoot, uint64 maxSequenceTimestamp, bytes32 expectedFinalAccInputHash, address l2Coinbase, bytes dataAvailabilityMessage) returns()
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerTransactorSession) SequenceBatchesValidium(batches []PolygonValidiumEtrogValidiumBatchData, indexL1InfoRoot uint32, maxSequenceTimestamp uint64, expectedFinalAccInputHash [32]byte, l2Coinbase common.Address, dataAvailabilityMessage []byte) (*types.Transaction, error) {
	return _PolygonvalidiumXlayer.Contract.SequenceBatchesValidium(&_PolygonvalidiumXlayer.TransactOpts, batches, indexL1InfoRoot, maxSequenceTimestamp, expectedFinalAccInputHash, l2Coinbase, dataAvailabilityMessage)
}

// SequenceForceBatches is a paid mutator transaction binding the contract method 0x9f26f840.
//
// Solidity: function sequenceForceBatches((bytes,bytes32,uint64,bytes32)[] batches) returns()
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerTransactor) SequenceForceBatches(opts *bind.TransactOpts, batches []PolygonRollupBaseEtrogBatchData) (*types.Transaction, error) {
	return _PolygonvalidiumXlayer.contract.Transact(opts, "sequenceForceBatches", batches)
}

// SequenceForceBatches is a paid mutator transaction binding the contract method 0x9f26f840.
//
// Solidity: function sequenceForceBatches((bytes,bytes32,uint64,bytes32)[] batches) returns()
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerSession) SequenceForceBatches(batches []PolygonRollupBaseEtrogBatchData) (*types.Transaction, error) {
	return _PolygonvalidiumXlayer.Contract.SequenceForceBatches(&_PolygonvalidiumXlayer.TransactOpts, batches)
}

// SequenceForceBatches is a paid mutator transaction binding the contract method 0x9f26f840.
//
// Solidity: function sequenceForceBatches((bytes,bytes32,uint64,bytes32)[] batches) returns()
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerTransactorSession) SequenceForceBatches(batches []PolygonRollupBaseEtrogBatchData) (*types.Transaction, error) {
	return _PolygonvalidiumXlayer.Contract.SequenceForceBatches(&_PolygonvalidiumXlayer.TransactOpts, batches)
}

// SetDataAvailabilityProtocol is a paid mutator transaction binding the contract method 0x7cd76b8b.
//
// Solidity: function setDataAvailabilityProtocol(address newDataAvailabilityProtocol) returns()
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerTransactor) SetDataAvailabilityProtocol(opts *bind.TransactOpts, newDataAvailabilityProtocol common.Address) (*types.Transaction, error) {
	return _PolygonvalidiumXlayer.contract.Transact(opts, "setDataAvailabilityProtocol", newDataAvailabilityProtocol)
}

// SetDataAvailabilityProtocol is a paid mutator transaction binding the contract method 0x7cd76b8b.
//
// Solidity: function setDataAvailabilityProtocol(address newDataAvailabilityProtocol) returns()
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerSession) SetDataAvailabilityProtocol(newDataAvailabilityProtocol common.Address) (*types.Transaction, error) {
	return _PolygonvalidiumXlayer.Contract.SetDataAvailabilityProtocol(&_PolygonvalidiumXlayer.TransactOpts, newDataAvailabilityProtocol)
}

// SetDataAvailabilityProtocol is a paid mutator transaction binding the contract method 0x7cd76b8b.
//
// Solidity: function setDataAvailabilityProtocol(address newDataAvailabilityProtocol) returns()
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerTransactorSession) SetDataAvailabilityProtocol(newDataAvailabilityProtocol common.Address) (*types.Transaction, error) {
	return _PolygonvalidiumXlayer.Contract.SetDataAvailabilityProtocol(&_PolygonvalidiumXlayer.TransactOpts, newDataAvailabilityProtocol)
}

// SetForceBatchAddress is a paid mutator transaction binding the contract method 0x91cafe32.
//
// Solidity: function setForceBatchAddress(address newForceBatchAddress) returns()
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerTransactor) SetForceBatchAddress(opts *bind.TransactOpts, newForceBatchAddress common.Address) (*types.Transaction, error) {
	return _PolygonvalidiumXlayer.contract.Transact(opts, "setForceBatchAddress", newForceBatchAddress)
}

// SetForceBatchAddress is a paid mutator transaction binding the contract method 0x91cafe32.
//
// Solidity: function setForceBatchAddress(address newForceBatchAddress) returns()
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerSession) SetForceBatchAddress(newForceBatchAddress common.Address) (*types.Transaction, error) {
	return _PolygonvalidiumXlayer.Contract.SetForceBatchAddress(&_PolygonvalidiumXlayer.TransactOpts, newForceBatchAddress)
}

// SetForceBatchAddress is a paid mutator transaction binding the contract method 0x91cafe32.
//
// Solidity: function setForceBatchAddress(address newForceBatchAddress) returns()
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerTransactorSession) SetForceBatchAddress(newForceBatchAddress common.Address) (*types.Transaction, error) {
	return _PolygonvalidiumXlayer.Contract.SetForceBatchAddress(&_PolygonvalidiumXlayer.TransactOpts, newForceBatchAddress)
}

// SetForceBatchTimeout is a paid mutator transaction binding the contract method 0x4e487706.
//
// Solidity: function setForceBatchTimeout(uint64 newforceBatchTimeout) returns()
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerTransactor) SetForceBatchTimeout(opts *bind.TransactOpts, newforceBatchTimeout uint64) (*types.Transaction, error) {
	return _PolygonvalidiumXlayer.contract.Transact(opts, "setForceBatchTimeout", newforceBatchTimeout)
}

// SetForceBatchTimeout is a paid mutator transaction binding the contract method 0x4e487706.
//
// Solidity: function setForceBatchTimeout(uint64 newforceBatchTimeout) returns()
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerSession) SetForceBatchTimeout(newforceBatchTimeout uint64) (*types.Transaction, error) {
	return _PolygonvalidiumXlayer.Contract.SetForceBatchTimeout(&_PolygonvalidiumXlayer.TransactOpts, newforceBatchTimeout)
}

// SetForceBatchTimeout is a paid mutator transaction binding the contract method 0x4e487706.
//
// Solidity: function setForceBatchTimeout(uint64 newforceBatchTimeout) returns()
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerTransactorSession) SetForceBatchTimeout(newforceBatchTimeout uint64) (*types.Transaction, error) {
	return _PolygonvalidiumXlayer.Contract.SetForceBatchTimeout(&_PolygonvalidiumXlayer.TransactOpts, newforceBatchTimeout)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerTransactor) SetTrustedSequencer(opts *bind.TransactOpts, newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _PolygonvalidiumXlayer.contract.Transact(opts, "setTrustedSequencer", newTrustedSequencer)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerSession) SetTrustedSequencer(newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _PolygonvalidiumXlayer.Contract.SetTrustedSequencer(&_PolygonvalidiumXlayer.TransactOpts, newTrustedSequencer)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerTransactorSession) SetTrustedSequencer(newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _PolygonvalidiumXlayer.Contract.SetTrustedSequencer(&_PolygonvalidiumXlayer.TransactOpts, newTrustedSequencer)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerTransactor) SetTrustedSequencerURL(opts *bind.TransactOpts, newTrustedSequencerURL string) (*types.Transaction, error) {
	return _PolygonvalidiumXlayer.contract.Transact(opts, "setTrustedSequencerURL", newTrustedSequencerURL)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerSession) SetTrustedSequencerURL(newTrustedSequencerURL string) (*types.Transaction, error) {
	return _PolygonvalidiumXlayer.Contract.SetTrustedSequencerURL(&_PolygonvalidiumXlayer.TransactOpts, newTrustedSequencerURL)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerTransactorSession) SetTrustedSequencerURL(newTrustedSequencerURL string) (*types.Transaction, error) {
	return _PolygonvalidiumXlayer.Contract.SetTrustedSequencerURL(&_PolygonvalidiumXlayer.TransactOpts, newTrustedSequencerURL)
}

// SwitchSequenceWithDataAvailability is a paid mutator transaction binding the contract method 0x2acdc2b6.
//
// Solidity: function switchSequenceWithDataAvailability(bool newIsSequenceWithDataAvailabilityAllowed) returns()
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerTransactor) SwitchSequenceWithDataAvailability(opts *bind.TransactOpts, newIsSequenceWithDataAvailabilityAllowed bool) (*types.Transaction, error) {
	return _PolygonvalidiumXlayer.contract.Transact(opts, "switchSequenceWithDataAvailability", newIsSequenceWithDataAvailabilityAllowed)
}

// SwitchSequenceWithDataAvailability is a paid mutator transaction binding the contract method 0x2acdc2b6.
//
// Solidity: function switchSequenceWithDataAvailability(bool newIsSequenceWithDataAvailabilityAllowed) returns()
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerSession) SwitchSequenceWithDataAvailability(newIsSequenceWithDataAvailabilityAllowed bool) (*types.Transaction, error) {
	return _PolygonvalidiumXlayer.Contract.SwitchSequenceWithDataAvailability(&_PolygonvalidiumXlayer.TransactOpts, newIsSequenceWithDataAvailabilityAllowed)
}

// SwitchSequenceWithDataAvailability is a paid mutator transaction binding the contract method 0x2acdc2b6.
//
// Solidity: function switchSequenceWithDataAvailability(bool newIsSequenceWithDataAvailabilityAllowed) returns()
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerTransactorSession) SwitchSequenceWithDataAvailability(newIsSequenceWithDataAvailabilityAllowed bool) (*types.Transaction, error) {
	return _PolygonvalidiumXlayer.Contract.SwitchSequenceWithDataAvailability(&_PolygonvalidiumXlayer.TransactOpts, newIsSequenceWithDataAvailabilityAllowed)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerTransactor) TransferAdminRole(opts *bind.TransactOpts, newPendingAdmin common.Address) (*types.Transaction, error) {
	return _PolygonvalidiumXlayer.contract.Transact(opts, "transferAdminRole", newPendingAdmin)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerSession) TransferAdminRole(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _PolygonvalidiumXlayer.Contract.TransferAdminRole(&_PolygonvalidiumXlayer.TransactOpts, newPendingAdmin)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerTransactorSession) TransferAdminRole(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _PolygonvalidiumXlayer.Contract.TransferAdminRole(&_PolygonvalidiumXlayer.TransactOpts, newPendingAdmin)
}

// PolygonvalidiumXlayerAcceptAdminRoleIterator is returned from FilterAcceptAdminRole and is used to iterate over the raw logs and unpacked data for AcceptAdminRole events raised by the PolygonvalidiumXlayer contract.
type PolygonvalidiumXlayerAcceptAdminRoleIterator struct {
	Event *PolygonvalidiumXlayerAcceptAdminRole // Event containing the contract specifics and raw log

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
func (it *PolygonvalidiumXlayerAcceptAdminRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumXlayerAcceptAdminRole)
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
		it.Event = new(PolygonvalidiumXlayerAcceptAdminRole)
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
func (it *PolygonvalidiumXlayerAcceptAdminRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumXlayerAcceptAdminRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumXlayerAcceptAdminRole represents a AcceptAdminRole event raised by the PolygonvalidiumXlayer contract.
type PolygonvalidiumXlayerAcceptAdminRole struct {
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAcceptAdminRole is a free log retrieval operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerFilterer) FilterAcceptAdminRole(opts *bind.FilterOpts) (*PolygonvalidiumXlayerAcceptAdminRoleIterator, error) {

	logs, sub, err := _PolygonvalidiumXlayer.contract.FilterLogs(opts, "AcceptAdminRole")
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumXlayerAcceptAdminRoleIterator{contract: _PolygonvalidiumXlayer.contract, event: "AcceptAdminRole", logs: logs, sub: sub}, nil
}

// WatchAcceptAdminRole is a free log subscription operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerFilterer) WatchAcceptAdminRole(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumXlayerAcceptAdminRole) (event.Subscription, error) {

	logs, sub, err := _PolygonvalidiumXlayer.contract.WatchLogs(opts, "AcceptAdminRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumXlayerAcceptAdminRole)
				if err := _PolygonvalidiumXlayer.contract.UnpackLog(event, "AcceptAdminRole", log); err != nil {
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

// ParseAcceptAdminRole is a log parse operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerFilterer) ParseAcceptAdminRole(log types.Log) (*PolygonvalidiumXlayerAcceptAdminRole, error) {
	event := new(PolygonvalidiumXlayerAcceptAdminRole)
	if err := _PolygonvalidiumXlayer.contract.UnpackLog(event, "AcceptAdminRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonvalidiumXlayerForceBatchIterator is returned from FilterForceBatch and is used to iterate over the raw logs and unpacked data for ForceBatch events raised by the PolygonvalidiumXlayer contract.
type PolygonvalidiumXlayerForceBatchIterator struct {
	Event *PolygonvalidiumXlayerForceBatch // Event containing the contract specifics and raw log

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
func (it *PolygonvalidiumXlayerForceBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumXlayerForceBatch)
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
		it.Event = new(PolygonvalidiumXlayerForceBatch)
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
func (it *PolygonvalidiumXlayerForceBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumXlayerForceBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumXlayerForceBatch represents a ForceBatch event raised by the PolygonvalidiumXlayer contract.
type PolygonvalidiumXlayerForceBatch struct {
	ForceBatchNum      uint64
	LastGlobalExitRoot [32]byte
	Sequencer          common.Address
	Transactions       []byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterForceBatch is a free log retrieval operation binding the contract event 0xf94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc931.
//
// Solidity: event ForceBatch(uint64 indexed forceBatchNum, bytes32 lastGlobalExitRoot, address sequencer, bytes transactions)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerFilterer) FilterForceBatch(opts *bind.FilterOpts, forceBatchNum []uint64) (*PolygonvalidiumXlayerForceBatchIterator, error) {

	var forceBatchNumRule []interface{}
	for _, forceBatchNumItem := range forceBatchNum {
		forceBatchNumRule = append(forceBatchNumRule, forceBatchNumItem)
	}

	logs, sub, err := _PolygonvalidiumXlayer.contract.FilterLogs(opts, "ForceBatch", forceBatchNumRule)
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumXlayerForceBatchIterator{contract: _PolygonvalidiumXlayer.contract, event: "ForceBatch", logs: logs, sub: sub}, nil
}

// WatchForceBatch is a free log subscription operation binding the contract event 0xf94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc931.
//
// Solidity: event ForceBatch(uint64 indexed forceBatchNum, bytes32 lastGlobalExitRoot, address sequencer, bytes transactions)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerFilterer) WatchForceBatch(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumXlayerForceBatch, forceBatchNum []uint64) (event.Subscription, error) {

	var forceBatchNumRule []interface{}
	for _, forceBatchNumItem := range forceBatchNum {
		forceBatchNumRule = append(forceBatchNumRule, forceBatchNumItem)
	}

	logs, sub, err := _PolygonvalidiumXlayer.contract.WatchLogs(opts, "ForceBatch", forceBatchNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumXlayerForceBatch)
				if err := _PolygonvalidiumXlayer.contract.UnpackLog(event, "ForceBatch", log); err != nil {
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

// ParseForceBatch is a log parse operation binding the contract event 0xf94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc931.
//
// Solidity: event ForceBatch(uint64 indexed forceBatchNum, bytes32 lastGlobalExitRoot, address sequencer, bytes transactions)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerFilterer) ParseForceBatch(log types.Log) (*PolygonvalidiumXlayerForceBatch, error) {
	event := new(PolygonvalidiumXlayerForceBatch)
	if err := _PolygonvalidiumXlayer.contract.UnpackLog(event, "ForceBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonvalidiumXlayerInitialSequenceBatchesIterator is returned from FilterInitialSequenceBatches and is used to iterate over the raw logs and unpacked data for InitialSequenceBatches events raised by the PolygonvalidiumXlayer contract.
type PolygonvalidiumXlayerInitialSequenceBatchesIterator struct {
	Event *PolygonvalidiumXlayerInitialSequenceBatches // Event containing the contract specifics and raw log

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
func (it *PolygonvalidiumXlayerInitialSequenceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumXlayerInitialSequenceBatches)
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
		it.Event = new(PolygonvalidiumXlayerInitialSequenceBatches)
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
func (it *PolygonvalidiumXlayerInitialSequenceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumXlayerInitialSequenceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumXlayerInitialSequenceBatches represents a InitialSequenceBatches event raised by the PolygonvalidiumXlayer contract.
type PolygonvalidiumXlayerInitialSequenceBatches struct {
	Transactions       []byte
	LastGlobalExitRoot [32]byte
	Sequencer          common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterInitialSequenceBatches is a free log retrieval operation binding the contract event 0x060116213bcbf54ca19fd649dc84b59ab2bbd200ab199770e4d923e222a28e7f.
//
// Solidity: event InitialSequenceBatches(bytes transactions, bytes32 lastGlobalExitRoot, address sequencer)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerFilterer) FilterInitialSequenceBatches(opts *bind.FilterOpts) (*PolygonvalidiumXlayerInitialSequenceBatchesIterator, error) {

	logs, sub, err := _PolygonvalidiumXlayer.contract.FilterLogs(opts, "InitialSequenceBatches")
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumXlayerInitialSequenceBatchesIterator{contract: _PolygonvalidiumXlayer.contract, event: "InitialSequenceBatches", logs: logs, sub: sub}, nil
}

// WatchInitialSequenceBatches is a free log subscription operation binding the contract event 0x060116213bcbf54ca19fd649dc84b59ab2bbd200ab199770e4d923e222a28e7f.
//
// Solidity: event InitialSequenceBatches(bytes transactions, bytes32 lastGlobalExitRoot, address sequencer)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerFilterer) WatchInitialSequenceBatches(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumXlayerInitialSequenceBatches) (event.Subscription, error) {

	logs, sub, err := _PolygonvalidiumXlayer.contract.WatchLogs(opts, "InitialSequenceBatches")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumXlayerInitialSequenceBatches)
				if err := _PolygonvalidiumXlayer.contract.UnpackLog(event, "InitialSequenceBatches", log); err != nil {
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

// ParseInitialSequenceBatches is a log parse operation binding the contract event 0x060116213bcbf54ca19fd649dc84b59ab2bbd200ab199770e4d923e222a28e7f.
//
// Solidity: event InitialSequenceBatches(bytes transactions, bytes32 lastGlobalExitRoot, address sequencer)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerFilterer) ParseInitialSequenceBatches(log types.Log) (*PolygonvalidiumXlayerInitialSequenceBatches, error) {
	event := new(PolygonvalidiumXlayerInitialSequenceBatches)
	if err := _PolygonvalidiumXlayer.contract.UnpackLog(event, "InitialSequenceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonvalidiumXlayerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the PolygonvalidiumXlayer contract.
type PolygonvalidiumXlayerInitializedIterator struct {
	Event *PolygonvalidiumXlayerInitialized // Event containing the contract specifics and raw log

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
func (it *PolygonvalidiumXlayerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumXlayerInitialized)
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
		it.Event = new(PolygonvalidiumXlayerInitialized)
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
func (it *PolygonvalidiumXlayerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumXlayerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumXlayerInitialized represents a Initialized event raised by the PolygonvalidiumXlayer contract.
type PolygonvalidiumXlayerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerFilterer) FilterInitialized(opts *bind.FilterOpts) (*PolygonvalidiumXlayerInitializedIterator, error) {

	logs, sub, err := _PolygonvalidiumXlayer.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumXlayerInitializedIterator{contract: _PolygonvalidiumXlayer.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumXlayerInitialized) (event.Subscription, error) {

	logs, sub, err := _PolygonvalidiumXlayer.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumXlayerInitialized)
				if err := _PolygonvalidiumXlayer.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerFilterer) ParseInitialized(log types.Log) (*PolygonvalidiumXlayerInitialized, error) {
	event := new(PolygonvalidiumXlayerInitialized)
	if err := _PolygonvalidiumXlayer.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonvalidiumXlayerRollbackBatchesIterator is returned from FilterRollbackBatches and is used to iterate over the raw logs and unpacked data for RollbackBatches events raised by the Polygonvalidiumetrog contract.
type PolygonvalidiumXlayerRollbackBatchesIterator struct {
	Event *PolygonvalidiumXlayerRollbackBatches // Event containing the contract specifics and raw log

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
func (it *PolygonvalidiumXlayerRollbackBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumXlayerRollbackBatches)
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
		it.Event = new(PolygonvalidiumXlayerRollbackBatches)
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
func (it *PolygonvalidiumXlayerRollbackBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumXlayerRollbackBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumXlayerRollbackBatches represents a RollbackBatches event raised by the Polygonvalidiumetrog contract.
type PolygonvalidiumXlayerRollbackBatches struct {
	TargetBatch            uint64
	AccInputHashToRollback [32]byte
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterRollbackBatches is a free log retrieval operation binding the contract event 0x1125aaf62d132d8e2d02005114f8fc360ff204c3105e4f1a700a1340dc55d5b1.
//
// Solidity: event RollbackBatches(uint64 indexed targetBatch, bytes32 accInputHashToRollback)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerFilterer) FilterRollbackBatches(opts *bind.FilterOpts, targetBatch []uint64) (*PolygonvalidiumXlayerRollbackBatchesIterator, error) {

	var targetBatchRule []interface{}
	for _, targetBatchItem := range targetBatch {
		targetBatchRule = append(targetBatchRule, targetBatchItem)
	}

	logs, sub, err := _PolygonvalidiumXlayer.contract.FilterLogs(opts, "RollbackBatches", targetBatchRule)
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumXlayerRollbackBatchesIterator{contract: _PolygonvalidiumXlayer.contract, event: "RollbackBatches", logs: logs, sub: sub}, nil
}

// WatchRollbackBatches is a free log subscription operation binding the contract event 0x1125aaf62d132d8e2d02005114f8fc360ff204c3105e4f1a700a1340dc55d5b1.
//
// Solidity: event RollbackBatches(uint64 indexed targetBatch, bytes32 accInputHashToRollback)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerFilterer) WatchRollbackBatches(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumXlayerRollbackBatches, targetBatch []uint64) (event.Subscription, error) {

	var targetBatchRule []interface{}
	for _, targetBatchItem := range targetBatch {
		targetBatchRule = append(targetBatchRule, targetBatchItem)
	}

	logs, sub, err := _PolygonvalidiumXlayer.contract.WatchLogs(opts, "RollbackBatches", targetBatchRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumXlayerRollbackBatches)
				if err := _PolygonvalidiumXlayer.contract.UnpackLog(event, "RollbackBatches", log); err != nil {
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

// ParseRollbackBatches is a log parse operation binding the contract event 0x1125aaf62d132d8e2d02005114f8fc360ff204c3105e4f1a700a1340dc55d5b1.
//
// Solidity: event RollbackBatches(uint64 indexed targetBatch, bytes32 accInputHashToRollback)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerFilterer) ParseRollbackBatches(log types.Log) (*PolygonvalidiumXlayerRollbackBatches, error) {
	event := new(PolygonvalidiumXlayerRollbackBatches)
	if err := _PolygonvalidiumXlayer.contract.UnpackLog(event, "RollbackBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonvalidiumXlayerSequenceBatchesIterator is returned from FilterSequenceBatches and is used to iterate over the raw logs and unpacked data for SequenceBatches events raised by the PolygonvalidiumXlayer contract.
type PolygonvalidiumXlayerSequenceBatchesIterator struct {
	Event *PolygonvalidiumXlayerSequenceBatches // Event containing the contract specifics and raw log

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
func (it *PolygonvalidiumXlayerSequenceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumXlayerSequenceBatches)
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
		it.Event = new(PolygonvalidiumXlayerSequenceBatches)
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
func (it *PolygonvalidiumXlayerSequenceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumXlayerSequenceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumXlayerSequenceBatches represents a SequenceBatches event raised by the PolygonvalidiumXlayer contract.
type PolygonvalidiumXlayerSequenceBatches struct {
	NumBatch   uint64
	L1InfoRoot [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSequenceBatches is a free log retrieval operation binding the contract event 0x3e54d0825ed78523037d00a81759237eb436ce774bd546993ee67a1b67b6e766.
//
// Solidity: event SequenceBatches(uint64 indexed numBatch, bytes32 l1InfoRoot)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerFilterer) FilterSequenceBatches(opts *bind.FilterOpts, numBatch []uint64) (*PolygonvalidiumXlayerSequenceBatchesIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _PolygonvalidiumXlayer.contract.FilterLogs(opts, "SequenceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumXlayerSequenceBatchesIterator{contract: _PolygonvalidiumXlayer.contract, event: "SequenceBatches", logs: logs, sub: sub}, nil
}

// WatchSequenceBatches is a free log subscription operation binding the contract event 0x3e54d0825ed78523037d00a81759237eb436ce774bd546993ee67a1b67b6e766.
//
// Solidity: event SequenceBatches(uint64 indexed numBatch, bytes32 l1InfoRoot)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerFilterer) WatchSequenceBatches(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumXlayerSequenceBatches, numBatch []uint64) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _PolygonvalidiumXlayer.contract.WatchLogs(opts, "SequenceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumXlayerSequenceBatches)
				if err := _PolygonvalidiumXlayer.contract.UnpackLog(event, "SequenceBatches", log); err != nil {
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

// ParseSequenceBatches is a log parse operation binding the contract event 0x3e54d0825ed78523037d00a81759237eb436ce774bd546993ee67a1b67b6e766.
//
// Solidity: event SequenceBatches(uint64 indexed numBatch, bytes32 l1InfoRoot)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerFilterer) ParseSequenceBatches(log types.Log) (*PolygonvalidiumXlayerSequenceBatches, error) {
	event := new(PolygonvalidiumXlayerSequenceBatches)
	if err := _PolygonvalidiumXlayer.contract.UnpackLog(event, "SequenceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonvalidiumXlayerSequenceForceBatchesIterator is returned from FilterSequenceForceBatches and is used to iterate over the raw logs and unpacked data for SequenceForceBatches events raised by the PolygonvalidiumXlayer contract.
type PolygonvalidiumXlayerSequenceForceBatchesIterator struct {
	Event *PolygonvalidiumXlayerSequenceForceBatches // Event containing the contract specifics and raw log

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
func (it *PolygonvalidiumXlayerSequenceForceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumXlayerSequenceForceBatches)
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
		it.Event = new(PolygonvalidiumXlayerSequenceForceBatches)
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
func (it *PolygonvalidiumXlayerSequenceForceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumXlayerSequenceForceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumXlayerSequenceForceBatches represents a SequenceForceBatches event raised by the PolygonvalidiumXlayer contract.
type PolygonvalidiumXlayerSequenceForceBatches struct {
	NumBatch uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSequenceForceBatches is a free log retrieval operation binding the contract event 0x648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a4.
//
// Solidity: event SequenceForceBatches(uint64 indexed numBatch)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerFilterer) FilterSequenceForceBatches(opts *bind.FilterOpts, numBatch []uint64) (*PolygonvalidiumXlayerSequenceForceBatchesIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _PolygonvalidiumXlayer.contract.FilterLogs(opts, "SequenceForceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumXlayerSequenceForceBatchesIterator{contract: _PolygonvalidiumXlayer.contract, event: "SequenceForceBatches", logs: logs, sub: sub}, nil
}

// WatchSequenceForceBatches is a free log subscription operation binding the contract event 0x648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a4.
//
// Solidity: event SequenceForceBatches(uint64 indexed numBatch)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerFilterer) WatchSequenceForceBatches(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumXlayerSequenceForceBatches, numBatch []uint64) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _PolygonvalidiumXlayer.contract.WatchLogs(opts, "SequenceForceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumXlayerSequenceForceBatches)
				if err := _PolygonvalidiumXlayer.contract.UnpackLog(event, "SequenceForceBatches", log); err != nil {
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

// ParseSequenceForceBatches is a log parse operation binding the contract event 0x648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a4.
//
// Solidity: event SequenceForceBatches(uint64 indexed numBatch)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerFilterer) ParseSequenceForceBatches(log types.Log) (*PolygonvalidiumXlayerSequenceForceBatches, error) {
	event := new(PolygonvalidiumXlayerSequenceForceBatches)
	if err := _PolygonvalidiumXlayer.contract.UnpackLog(event, "SequenceForceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonvalidiumXlayerSetDataAvailabilityProtocolIterator is returned from FilterSetDataAvailabilityProtocol and is used to iterate over the raw logs and unpacked data for SetDataAvailabilityProtocol events raised by the PolygonvalidiumXlayer contract.
type PolygonvalidiumXlayerSetDataAvailabilityProtocolIterator struct {
	Event *PolygonvalidiumXlayerSetDataAvailabilityProtocol // Event containing the contract specifics and raw log

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
func (it *PolygonvalidiumXlayerSetDataAvailabilityProtocolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumXlayerSetDataAvailabilityProtocol)
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
		it.Event = new(PolygonvalidiumXlayerSetDataAvailabilityProtocol)
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
func (it *PolygonvalidiumXlayerSetDataAvailabilityProtocolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumXlayerSetDataAvailabilityProtocolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumXlayerSetDataAvailabilityProtocol represents a SetDataAvailabilityProtocol event raised by the PolygonvalidiumXlayer contract.
type PolygonvalidiumXlayerSetDataAvailabilityProtocol struct {
	NewDataAvailabilityProtocol common.Address
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterSetDataAvailabilityProtocol is a free log retrieval operation binding the contract event 0xd331bd4c4cd1afecb94a225184bded161ff3213624ba4fb58c4f30c5a861144a.
//
// Solidity: event SetDataAvailabilityProtocol(address newDataAvailabilityProtocol)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerFilterer) FilterSetDataAvailabilityProtocol(opts *bind.FilterOpts) (*PolygonvalidiumXlayerSetDataAvailabilityProtocolIterator, error) {

	logs, sub, err := _PolygonvalidiumXlayer.contract.FilterLogs(opts, "SetDataAvailabilityProtocol")
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumXlayerSetDataAvailabilityProtocolIterator{contract: _PolygonvalidiumXlayer.contract, event: "SetDataAvailabilityProtocol", logs: logs, sub: sub}, nil
}

// WatchSetDataAvailabilityProtocol is a free log subscription operation binding the contract event 0xd331bd4c4cd1afecb94a225184bded161ff3213624ba4fb58c4f30c5a861144a.
//
// Solidity: event SetDataAvailabilityProtocol(address newDataAvailabilityProtocol)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerFilterer) WatchSetDataAvailabilityProtocol(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumXlayerSetDataAvailabilityProtocol) (event.Subscription, error) {

	logs, sub, err := _PolygonvalidiumXlayer.contract.WatchLogs(opts, "SetDataAvailabilityProtocol")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumXlayerSetDataAvailabilityProtocol)
				if err := _PolygonvalidiumXlayer.contract.UnpackLog(event, "SetDataAvailabilityProtocol", log); err != nil {
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

// ParseSetDataAvailabilityProtocol is a log parse operation binding the contract event 0xd331bd4c4cd1afecb94a225184bded161ff3213624ba4fb58c4f30c5a861144a.
//
// Solidity: event SetDataAvailabilityProtocol(address newDataAvailabilityProtocol)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerFilterer) ParseSetDataAvailabilityProtocol(log types.Log) (*PolygonvalidiumXlayerSetDataAvailabilityProtocol, error) {
	event := new(PolygonvalidiumXlayerSetDataAvailabilityProtocol)
	if err := _PolygonvalidiumXlayer.contract.UnpackLog(event, "SetDataAvailabilityProtocol", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonvalidiumXlayerSetForceBatchAddressIterator is returned from FilterSetForceBatchAddress and is used to iterate over the raw logs and unpacked data for SetForceBatchAddress events raised by the PolygonvalidiumXlayer contract.
type PolygonvalidiumXlayerSetForceBatchAddressIterator struct {
	Event *PolygonvalidiumXlayerSetForceBatchAddress // Event containing the contract specifics and raw log

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
func (it *PolygonvalidiumXlayerSetForceBatchAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumXlayerSetForceBatchAddress)
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
		it.Event = new(PolygonvalidiumXlayerSetForceBatchAddress)
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
func (it *PolygonvalidiumXlayerSetForceBatchAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumXlayerSetForceBatchAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumXlayerSetForceBatchAddress represents a SetForceBatchAddress event raised by the PolygonvalidiumXlayer contract.
type PolygonvalidiumXlayerSetForceBatchAddress struct {
	NewForceBatchAddress common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSetForceBatchAddress is a free log retrieval operation binding the contract event 0x5fbd7dd171301c4a1611a84aac4ba86d119478560557755f7927595b082634fb.
//
// Solidity: event SetForceBatchAddress(address newForceBatchAddress)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerFilterer) FilterSetForceBatchAddress(opts *bind.FilterOpts) (*PolygonvalidiumXlayerSetForceBatchAddressIterator, error) {

	logs, sub, err := _PolygonvalidiumXlayer.contract.FilterLogs(opts, "SetForceBatchAddress")
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumXlayerSetForceBatchAddressIterator{contract: _PolygonvalidiumXlayer.contract, event: "SetForceBatchAddress", logs: logs, sub: sub}, nil
}

// WatchSetForceBatchAddress is a free log subscription operation binding the contract event 0x5fbd7dd171301c4a1611a84aac4ba86d119478560557755f7927595b082634fb.
//
// Solidity: event SetForceBatchAddress(address newForceBatchAddress)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerFilterer) WatchSetForceBatchAddress(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumXlayerSetForceBatchAddress) (event.Subscription, error) {

	logs, sub, err := _PolygonvalidiumXlayer.contract.WatchLogs(opts, "SetForceBatchAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumXlayerSetForceBatchAddress)
				if err := _PolygonvalidiumXlayer.contract.UnpackLog(event, "SetForceBatchAddress", log); err != nil {
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

// ParseSetForceBatchAddress is a log parse operation binding the contract event 0x5fbd7dd171301c4a1611a84aac4ba86d119478560557755f7927595b082634fb.
//
// Solidity: event SetForceBatchAddress(address newForceBatchAddress)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerFilterer) ParseSetForceBatchAddress(log types.Log) (*PolygonvalidiumXlayerSetForceBatchAddress, error) {
	event := new(PolygonvalidiumXlayerSetForceBatchAddress)
	if err := _PolygonvalidiumXlayer.contract.UnpackLog(event, "SetForceBatchAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonvalidiumXlayerSetForceBatchTimeoutIterator is returned from FilterSetForceBatchTimeout and is used to iterate over the raw logs and unpacked data for SetForceBatchTimeout events raised by the PolygonvalidiumXlayer contract.
type PolygonvalidiumXlayerSetForceBatchTimeoutIterator struct {
	Event *PolygonvalidiumXlayerSetForceBatchTimeout // Event containing the contract specifics and raw log

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
func (it *PolygonvalidiumXlayerSetForceBatchTimeoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumXlayerSetForceBatchTimeout)
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
		it.Event = new(PolygonvalidiumXlayerSetForceBatchTimeout)
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
func (it *PolygonvalidiumXlayerSetForceBatchTimeoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumXlayerSetForceBatchTimeoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumXlayerSetForceBatchTimeout represents a SetForceBatchTimeout event raised by the PolygonvalidiumXlayer contract.
type PolygonvalidiumXlayerSetForceBatchTimeout struct {
	NewforceBatchTimeout uint64
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSetForceBatchTimeout is a free log retrieval operation binding the contract event 0xa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b.
//
// Solidity: event SetForceBatchTimeout(uint64 newforceBatchTimeout)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerFilterer) FilterSetForceBatchTimeout(opts *bind.FilterOpts) (*PolygonvalidiumXlayerSetForceBatchTimeoutIterator, error) {

	logs, sub, err := _PolygonvalidiumXlayer.contract.FilterLogs(opts, "SetForceBatchTimeout")
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumXlayerSetForceBatchTimeoutIterator{contract: _PolygonvalidiumXlayer.contract, event: "SetForceBatchTimeout", logs: logs, sub: sub}, nil
}

// WatchSetForceBatchTimeout is a free log subscription operation binding the contract event 0xa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b.
//
// Solidity: event SetForceBatchTimeout(uint64 newforceBatchTimeout)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerFilterer) WatchSetForceBatchTimeout(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumXlayerSetForceBatchTimeout) (event.Subscription, error) {

	logs, sub, err := _PolygonvalidiumXlayer.contract.WatchLogs(opts, "SetForceBatchTimeout")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumXlayerSetForceBatchTimeout)
				if err := _PolygonvalidiumXlayer.contract.UnpackLog(event, "SetForceBatchTimeout", log); err != nil {
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

// ParseSetForceBatchTimeout is a log parse operation binding the contract event 0xa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b.
//
// Solidity: event SetForceBatchTimeout(uint64 newforceBatchTimeout)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerFilterer) ParseSetForceBatchTimeout(log types.Log) (*PolygonvalidiumXlayerSetForceBatchTimeout, error) {
	event := new(PolygonvalidiumXlayerSetForceBatchTimeout)
	if err := _PolygonvalidiumXlayer.contract.UnpackLog(event, "SetForceBatchTimeout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonvalidiumXlayerSetTrustedSequencerIterator is returned from FilterSetTrustedSequencer and is used to iterate over the raw logs and unpacked data for SetTrustedSequencer events raised by the PolygonvalidiumXlayer contract.
type PolygonvalidiumXlayerSetTrustedSequencerIterator struct {
	Event *PolygonvalidiumXlayerSetTrustedSequencer // Event containing the contract specifics and raw log

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
func (it *PolygonvalidiumXlayerSetTrustedSequencerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumXlayerSetTrustedSequencer)
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
		it.Event = new(PolygonvalidiumXlayerSetTrustedSequencer)
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
func (it *PolygonvalidiumXlayerSetTrustedSequencerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumXlayerSetTrustedSequencerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumXlayerSetTrustedSequencer represents a SetTrustedSequencer event raised by the PolygonvalidiumXlayer contract.
type PolygonvalidiumXlayerSetTrustedSequencer struct {
	NewTrustedSequencer common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedSequencer is a free log retrieval operation binding the contract event 0xf54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0.
//
// Solidity: event SetTrustedSequencer(address newTrustedSequencer)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerFilterer) FilterSetTrustedSequencer(opts *bind.FilterOpts) (*PolygonvalidiumXlayerSetTrustedSequencerIterator, error) {

	logs, sub, err := _PolygonvalidiumXlayer.contract.FilterLogs(opts, "SetTrustedSequencer")
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumXlayerSetTrustedSequencerIterator{contract: _PolygonvalidiumXlayer.contract, event: "SetTrustedSequencer", logs: logs, sub: sub}, nil
}

// WatchSetTrustedSequencer is a free log subscription operation binding the contract event 0xf54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0.
//
// Solidity: event SetTrustedSequencer(address newTrustedSequencer)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerFilterer) WatchSetTrustedSequencer(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumXlayerSetTrustedSequencer) (event.Subscription, error) {

	logs, sub, err := _PolygonvalidiumXlayer.contract.WatchLogs(opts, "SetTrustedSequencer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumXlayerSetTrustedSequencer)
				if err := _PolygonvalidiumXlayer.contract.UnpackLog(event, "SetTrustedSequencer", log); err != nil {
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

// ParseSetTrustedSequencer is a log parse operation binding the contract event 0xf54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0.
//
// Solidity: event SetTrustedSequencer(address newTrustedSequencer)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerFilterer) ParseSetTrustedSequencer(log types.Log) (*PolygonvalidiumXlayerSetTrustedSequencer, error) {
	event := new(PolygonvalidiumXlayerSetTrustedSequencer)
	if err := _PolygonvalidiumXlayer.contract.UnpackLog(event, "SetTrustedSequencer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonvalidiumXlayerSetTrustedSequencerURLIterator is returned from FilterSetTrustedSequencerURL and is used to iterate over the raw logs and unpacked data for SetTrustedSequencerURL events raised by the PolygonvalidiumXlayer contract.
type PolygonvalidiumXlayerSetTrustedSequencerURLIterator struct {
	Event *PolygonvalidiumXlayerSetTrustedSequencerURL // Event containing the contract specifics and raw log

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
func (it *PolygonvalidiumXlayerSetTrustedSequencerURLIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumXlayerSetTrustedSequencerURL)
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
		it.Event = new(PolygonvalidiumXlayerSetTrustedSequencerURL)
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
func (it *PolygonvalidiumXlayerSetTrustedSequencerURLIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumXlayerSetTrustedSequencerURLIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumXlayerSetTrustedSequencerURL represents a SetTrustedSequencerURL event raised by the PolygonvalidiumXlayer contract.
type PolygonvalidiumXlayerSetTrustedSequencerURL struct {
	NewTrustedSequencerURL string
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedSequencerURL is a free log retrieval operation binding the contract event 0x6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20.
//
// Solidity: event SetTrustedSequencerURL(string newTrustedSequencerURL)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerFilterer) FilterSetTrustedSequencerURL(opts *bind.FilterOpts) (*PolygonvalidiumXlayerSetTrustedSequencerURLIterator, error) {

	logs, sub, err := _PolygonvalidiumXlayer.contract.FilterLogs(opts, "SetTrustedSequencerURL")
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumXlayerSetTrustedSequencerURLIterator{contract: _PolygonvalidiumXlayer.contract, event: "SetTrustedSequencerURL", logs: logs, sub: sub}, nil
}

// WatchSetTrustedSequencerURL is a free log subscription operation binding the contract event 0x6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20.
//
// Solidity: event SetTrustedSequencerURL(string newTrustedSequencerURL)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerFilterer) WatchSetTrustedSequencerURL(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumXlayerSetTrustedSequencerURL) (event.Subscription, error) {

	logs, sub, err := _PolygonvalidiumXlayer.contract.WatchLogs(opts, "SetTrustedSequencerURL")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumXlayerSetTrustedSequencerURL)
				if err := _PolygonvalidiumXlayer.contract.UnpackLog(event, "SetTrustedSequencerURL", log); err != nil {
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

// ParseSetTrustedSequencerURL is a log parse operation binding the contract event 0x6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20.
//
// Solidity: event SetTrustedSequencerURL(string newTrustedSequencerURL)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerFilterer) ParseSetTrustedSequencerURL(log types.Log) (*PolygonvalidiumXlayerSetTrustedSequencerURL, error) {
	event := new(PolygonvalidiumXlayerSetTrustedSequencerURL)
	if err := _PolygonvalidiumXlayer.contract.UnpackLog(event, "SetTrustedSequencerURL", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonvalidiumXlayerSwitchSequenceWithDataAvailabilityIterator is returned from FilterSwitchSequenceWithDataAvailability and is used to iterate over the raw logs and unpacked data for SwitchSequenceWithDataAvailability events raised by the PolygonvalidiumXlayer contract.
type PolygonvalidiumXlayerSwitchSequenceWithDataAvailabilityIterator struct {
	Event *PolygonvalidiumXlayerSwitchSequenceWithDataAvailability // Event containing the contract specifics and raw log

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
func (it *PolygonvalidiumXlayerSwitchSequenceWithDataAvailabilityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumXlayerSwitchSequenceWithDataAvailability)
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
		it.Event = new(PolygonvalidiumXlayerSwitchSequenceWithDataAvailability)
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
func (it *PolygonvalidiumXlayerSwitchSequenceWithDataAvailabilityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumXlayerSwitchSequenceWithDataAvailabilityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumXlayerSwitchSequenceWithDataAvailability represents a SwitchSequenceWithDataAvailability event raised by the PolygonvalidiumXlayer contract.
type PolygonvalidiumXlayerSwitchSequenceWithDataAvailability struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSwitchSequenceWithDataAvailability is a free log retrieval operation binding the contract event 0xf32a0473f809a720a4f8af1e50d353f1caf7452030626fdaac4273f5e6587f41.
//
// Solidity: event SwitchSequenceWithDataAvailability()
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerFilterer) FilterSwitchSequenceWithDataAvailability(opts *bind.FilterOpts) (*PolygonvalidiumXlayerSwitchSequenceWithDataAvailabilityIterator, error) {

	logs, sub, err := _PolygonvalidiumXlayer.contract.FilterLogs(opts, "SwitchSequenceWithDataAvailability")
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumXlayerSwitchSequenceWithDataAvailabilityIterator{contract: _PolygonvalidiumXlayer.contract, event: "SwitchSequenceWithDataAvailability", logs: logs, sub: sub}, nil
}

// WatchSwitchSequenceWithDataAvailability is a free log subscription operation binding the contract event 0xf32a0473f809a720a4f8af1e50d353f1caf7452030626fdaac4273f5e6587f41.
//
// Solidity: event SwitchSequenceWithDataAvailability()
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerFilterer) WatchSwitchSequenceWithDataAvailability(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumXlayerSwitchSequenceWithDataAvailability) (event.Subscription, error) {

	logs, sub, err := _PolygonvalidiumXlayer.contract.WatchLogs(opts, "SwitchSequenceWithDataAvailability")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumXlayerSwitchSequenceWithDataAvailability)
				if err := _PolygonvalidiumXlayer.contract.UnpackLog(event, "SwitchSequenceWithDataAvailability", log); err != nil {
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

// ParseSwitchSequenceWithDataAvailability is a log parse operation binding the contract event 0xf32a0473f809a720a4f8af1e50d353f1caf7452030626fdaac4273f5e6587f41.
//
// Solidity: event SwitchSequenceWithDataAvailability()
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerFilterer) ParseSwitchSequenceWithDataAvailability(log types.Log) (*PolygonvalidiumXlayerSwitchSequenceWithDataAvailability, error) {
	event := new(PolygonvalidiumXlayerSwitchSequenceWithDataAvailability)
	if err := _PolygonvalidiumXlayer.contract.UnpackLog(event, "SwitchSequenceWithDataAvailability", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonvalidiumXlayerTransferAdminRoleIterator is returned from FilterTransferAdminRole and is used to iterate over the raw logs and unpacked data for TransferAdminRole events raised by the PolygonvalidiumXlayer contract.
type PolygonvalidiumXlayerTransferAdminRoleIterator struct {
	Event *PolygonvalidiumXlayerTransferAdminRole // Event containing the contract specifics and raw log

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
func (it *PolygonvalidiumXlayerTransferAdminRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumXlayerTransferAdminRole)
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
		it.Event = new(PolygonvalidiumXlayerTransferAdminRole)
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
func (it *PolygonvalidiumXlayerTransferAdminRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumXlayerTransferAdminRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumXlayerTransferAdminRole represents a TransferAdminRole event raised by the PolygonvalidiumXlayer contract.
type PolygonvalidiumXlayerTransferAdminRole struct {
	NewPendingAdmin common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTransferAdminRole is a free log retrieval operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerFilterer) FilterTransferAdminRole(opts *bind.FilterOpts) (*PolygonvalidiumXlayerTransferAdminRoleIterator, error) {

	logs, sub, err := _PolygonvalidiumXlayer.contract.FilterLogs(opts, "TransferAdminRole")
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumXlayerTransferAdminRoleIterator{contract: _PolygonvalidiumXlayer.contract, event: "TransferAdminRole", logs: logs, sub: sub}, nil
}

// WatchTransferAdminRole is a free log subscription operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerFilterer) WatchTransferAdminRole(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumXlayerTransferAdminRole) (event.Subscription, error) {

	logs, sub, err := _PolygonvalidiumXlayer.contract.WatchLogs(opts, "TransferAdminRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumXlayerTransferAdminRole)
				if err := _PolygonvalidiumXlayer.contract.UnpackLog(event, "TransferAdminRole", log); err != nil {
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

// ParseTransferAdminRole is a log parse operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerFilterer) ParseTransferAdminRole(log types.Log) (*PolygonvalidiumXlayerTransferAdminRole, error) {
	event := new(PolygonvalidiumXlayerTransferAdminRole)
	if err := _PolygonvalidiumXlayer.contract.UnpackLog(event, "TransferAdminRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonvalidiumXlayerVerifyBatchesIterator is returned from FilterVerifyBatches and is used to iterate over the raw logs and unpacked data for VerifyBatches events raised by the PolygonvalidiumXlayer contract.
type PolygonvalidiumXlayerVerifyBatchesIterator struct {
	Event *PolygonvalidiumXlayerVerifyBatches // Event containing the contract specifics and raw log

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
func (it *PolygonvalidiumXlayerVerifyBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumXlayerVerifyBatches)
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
		it.Event = new(PolygonvalidiumXlayerVerifyBatches)
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
func (it *PolygonvalidiumXlayerVerifyBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumXlayerVerifyBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumXlayerVerifyBatches represents a VerifyBatches event raised by the PolygonvalidiumXlayer contract.
type PolygonvalidiumXlayerVerifyBatches struct {
	NumBatch   uint64
	StateRoot  [32]byte
	Aggregator common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVerifyBatches is a free log retrieval operation binding the contract event 0x9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f5966.
//
// Solidity: event VerifyBatches(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerFilterer) FilterVerifyBatches(opts *bind.FilterOpts, numBatch []uint64, aggregator []common.Address) (*PolygonvalidiumXlayerVerifyBatchesIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _PolygonvalidiumXlayer.contract.FilterLogs(opts, "VerifyBatches", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumXlayerVerifyBatchesIterator{contract: _PolygonvalidiumXlayer.contract, event: "VerifyBatches", logs: logs, sub: sub}, nil
}

// WatchVerifyBatches is a free log subscription operation binding the contract event 0x9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f5966.
//
// Solidity: event VerifyBatches(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerFilterer) WatchVerifyBatches(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumXlayerVerifyBatches, numBatch []uint64, aggregator []common.Address) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _PolygonvalidiumXlayer.contract.WatchLogs(opts, "VerifyBatches", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumXlayerVerifyBatches)
				if err := _PolygonvalidiumXlayer.contract.UnpackLog(event, "VerifyBatches", log); err != nil {
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

// ParseVerifyBatches is a log parse operation binding the contract event 0x9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f5966.
//
// Solidity: event VerifyBatches(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_PolygonvalidiumXlayer *PolygonvalidiumXlayerFilterer) ParseVerifyBatches(log types.Log) (*PolygonvalidiumXlayerVerifyBatches, error) {
	event := new(PolygonvalidiumXlayerVerifyBatches)
	if err := _PolygonvalidiumXlayer.contract.UnpackLog(event, "VerifyBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
