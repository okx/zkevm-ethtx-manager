package ethtxmanager

import (
	"context"
	"math/big"
	"testing"
	"time"

	zktypes "github.com/0xPolygon/zkevm-ethtx-manager/config/types"
	"github.com/0xPolygon/zkevm-ethtx-manager/hex"
	zkmanTypes "github.com/0xPolygon/zkevm-ethtx-manager/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

const (
	domain                 = "http://asset-onchain.base-defi.svc.test.local:7001"
	seqAddr                = "1a13bddcc02d363366e04d4aa588d3c125b0ff6f"
	aggAddr                = "66e39a1e507af777e8c385e2d91559e20e306303"
	contractAddr           = "8947dc90862f386968966b22dfe5edf96435bc2f"
	contractAddrAgg        = "1d5298ee11f7cd56fb842b7894346bfb2e47a95f"
	l1ChainID       uint64 = 11155111
	AccessKey              = ""
	SecretKey              = ""
	// domain       = "http://127.0.0.1:7001"
	// seqAddr      = "f39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
	// aggAddr      = "70997970C51812dc3A010C7d01b50e0d17dc79C8"
	// contractAddr = "812cB73e48841a6736bB94c65c56341817cE6304"
)

func TestUnpackMarshalSeqFork12(t *testing.T) {
	client := &Client{
		etherman: mockEtherman{},
		cfg: Config{
			CustodialAssets: CustodialAssetsConfig{
				Enable:            false,
				URL:               domain,
				Symbol:            2882,
				SequencerAddr:     common.HexToAddress(seqAddr),
				AggregatorAddr:    common.HexToAddress(aggAddr),
				WaitResultTimeout: zktypes.NewDuration(4 * time.Minute),
				OperateTypeSeq:    3,
				OperateTypeAgg:    4,
				ProjectSymbol:     3011,
				OperateSymbol:     2,
				SysFrom:           3,
				UserID:            0,
				OperateAmount:     0,
				RequestSignURI:    "/priapi/v1/assetonchain/ecology/ecologyOperate",
				QuerySignURI:      "/priapi/v1/assetonchain/ecology/querySignDataByOrderNo",
				AccessKey:         AccessKey,
				SecretKey:         SecretKey,
			},
		},
	}

	// Fork13 contract example TX data
	// Source: https://sepolia.etherscan.io/tx/0x2e361e9e6515f4ceef4d235a88878bc87a4be0c261995f2ef160264e56691641
	data, _ := hex.DecodeHex("0x165e8a8d00000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000672cc2e7ef4be24211f1efe1cd9faa74033914b15647469e9f6e84ce75384869102ae3f80000000000000000000000004820e45eb6cf99074389544f83fe1f30084bdab700000000000000000000000000000000000000000000000000000000000006e0000000000000000000000000000000000000000000000000000000000000000c14df0077a11ccf22bd14d1ec3e209ba5d8e065ce73556fdc8fb9c19b83af4ec3000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000cb974e114004736273079d8f2fb4c703de79e14eea80f77e14abc5a1cd3d169b00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000014df0077a11ccf22bd14d1ec3e209ba5d8e065ce73556fdc8fb9c19b83af4ec300000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000014df0077a11ccf22bd14d1ec3e209ba5d8e065ce73556fdc8fb9c19b83af4ec3000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000dd6adb9b5339c8211dc51e7a58a554ed96cf79e3ee7fc2584989fc59f9498a8500000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000014df0077a11ccf22bd14d1ec3e209ba5d8e065ce73556fdc8fb9c19b83af4ec3000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000dd6adb9b5339c8211dc51e7a58a554ed96cf79e3ee7fc2584989fc59f9498a8500000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000014df0077a11ccf22bd14d1ec3e209ba5d8e065ce73556fdc8fb9c19b83af4ec3000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000dd6adb9b5339c8211dc51e7a58a554ed96cf79e3ee7fc2584989fc59f9498a85000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000dd6adb9b5339c8211dc51e7a58a554ed96cf79e3ee7fc2584989fc59f9498a850000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000006a1686e10575fe5046ef79d653dde24a3ade0d4a82aad51956ad6268e3f33cca00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000014df0077a11ccf22bd14d1ec3e209ba5d8e065ce73556fdc8fb9c19b83af4ec300000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000556810c9dd47ba22910ff6045631ea61f3c723104859aa392eaa2c3dbadb2bf8c411ff640c48b18ec9b6efe6bfde82414731db333dfb1b09185c10c3281c6dbee61cf39fd6e51aad88f6f4ce6ab8827279cfffb922660000000000000000000000")
	to := common.HexToAddress("0x4820E45eB6Cf99074389544F83FE1F30084bdAb7")
	mTx := zkmanTypes.MonitoredTx{
		From:     common.HexToAddress(seqAddr),
		To:       &to,
		Nonce:    uint64(0),
		Value:    big.NewInt(10),
		Data:     data,
		Gas:      uint64(50000),
		GasPrice: big.NewInt(10),
	}

	args, err := client.unpackSequenceBatchesTx(mTx.Tx())
	if err != nil {
		t.Fatal(err)
	}
	if len(args.Batches) != 12 {
		t.Fatalf("Expected 12 batches, Got = %d\n", 12)
	}
	if args.L2Coinbase != to {
		t.Fatalf("Expected %s, got = %s\n", to, args.L2Coinbase)
	}

	_, err = args.marshal(common.HexToAddress(contractAddr), mTx)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUnpackMarshalAggFork12(t *testing.T) {
	client := &Client{
		etherman: mockEtherman{},
		cfg: Config{
			CustodialAssets: CustodialAssetsConfig{
				Enable:            false,
				URL:               domain,
				Symbol:            2882,
				SequencerAddr:     common.HexToAddress(seqAddr),
				AggregatorAddr:    common.HexToAddress(aggAddr),
				WaitResultTimeout: zktypes.NewDuration(4 * time.Minute),
				OperateTypeSeq:    3,
				OperateTypeAgg:    4,
				ProjectSymbol:     3011,
				OperateSymbol:     2,
				SysFrom:           3,
				UserID:            0,
				OperateAmount:     0,
				RequestSignURI:    "/priapi/v1/assetonchain/ecology/ecologyOperate",
				QuerySignURI:      "/priapi/v1/assetonchain/ecology/querySignDataByOrderNo",
				AccessKey:         AccessKey,
				SecretKey:         SecretKey,
			},
		},
	}

	// Same sample data used in Fork8 (since contract method did not change)
	data, _ := hex.DecodeHex("0x1489ed100000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000500ae00000000000000000000000000000000000000000000000000000000000500b33d2b755195521747e8b32d1f28b033f0d93082872cd253ee42e3bcfaf2cd8a3129e24287e603ff699b4322778e511f2bf9f4d305f199403c24db686e44c5b1cf000000000000000000000000a57d7641fba20d916a7cb61de5af71e9743e4b571666fc7a9b894634feba222cebcf0678a56b2b6348e31af2eddea7b98fc9b36a017778ab8f75848be64ffeaf6974eb09369e01a96c6c443d611289e3b18ebcae02fa881020960c0cd3fc9819024a60e1869baead40db9422af24947c1ec825d31f3643021dd597b4b6799b1669e50d5ccc29dcab571308b4afd9f69856b6f1cf1560e3fdedf10b3b9b0fa9ebb9a09af91ebd61c3722274eef718948620bcc84a22ba517411823319b4fad2bac5ca0d30f72261ca3a291fe11ee3f2d65f21947c2cda3b57c03ea3d576cef736b45185e6a4f7e2fae3d2e1442245ac6cf906032d032cb2ce822d9e16521b862e46fa28bc62ce12f0804fa057e8ecef7bff12d72423095607eb51333715cdc3cd39d316e02ef9d92ce24ec75518353670b87b48a7178bb86cb37a5c2b08b8705977b3d046a8f7777dbcb988f6a8508d9ebe3c3ce42f33985c054be78209809df36aa39ebecc34c44bbdf116ba4226613b5f2562b41139c17ccf95c2791327defc83fc7e611bc8e08aa72a9bca9138652f8ca9742f25eb11f74ffd3737e147acb9dd4cbffc298860cfab785c4e9b5cceb92aafe344262278e991e0d715bd9a66aafe8153b723b9ca7ac050205970cfc85a8dca1cad2a22a2b4642d0217277893d12d1ede382ca42eed04b686a9229eb3fed45d838b27474a3242ee89ded32881b4abc11fd072d7b718af15171364c7f5b285f5da8d303245456ec4e9beb1c2930cb7db35b4ab991a7ea48a2373a4c2ef201ba4151c15e12c35f10e806ab5fd807fb6cf20255ca28e54d5952dba961b510af29296c7240e73818dfef35e5919998ce4885fc9671c0379f63479f2d61197725916b1a814b67a71ce84f0148fd0205ff1bc3e20df8a0e4a1180cecc57cdaa4f690c07bb2010b9e0c3d25daa958dbdbaad240e5900a97e8d0e75197f983b7870188775471344ab8ee51610daf6f634de0407c9d24ca762a38a6f2ba44d4114749c616d972bfbfd7cb939ababbf9e6d1d569b9a3ff396ab3475b3bfe3332b3ac218c2f6e516780ca06e7ec5a98c542686c92c2c2ff09cff58101df7824b01aa6b0b85e206")
	to := common.HexToAddress(contractAddrAgg)
	mTx := zkmanTypes.MonitoredTx{
		From:     common.HexToAddress(aggAddr),
		To:       &to,
		Nonce:    uint64(0),
		Value:    big.NewInt(10),
		Data:     data,
		Gas:      uint64(50000),
		GasPrice: big.NewInt(10),
	}

	args, err := client.unpackVerifyBatchesTrustedAggregatorTx(mTx.Tx())
	if err != nil {
		t.Fatal(err)
	}
	if args.RollupId != 1 {
		t.Fatalf("Got = %d\n", args.RollupId)
	}
	if args.Beneficiary != common.HexToAddress("0xa57d7641FBA20D916A7cb61DE5AF71e9743e4B57") {
		t.Fatalf("Got = %s\n", args.Beneficiary)
	}

	_, err = args.marshal(common.HexToAddress(contractAddrAgg), mTx)
	if err != nil {
		t.Fatal(err)
	}
}

type mockEtherman struct{}

func (m mockEtherman) PendingNonce(ctx context.Context, account common.Address) (uint64, error) {
	//TODO implement me
	panic("implement me")
}

func (m mockEtherman) EstimateGasBlobTx(ctx context.Context, from common.Address, to *common.Address, gasFeeCap *big.Int, gasTipCap *big.Int, value *big.Int, data []byte) (uint64, error) {
	//TODO implement me
	panic("implement me")
}

func (m mockEtherman) GetLatestBlockNumber(ctx context.Context) (uint64, error) {
	//TODO implement me
	panic("implement me")
}

func (m mockEtherman) GetHeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	//TODO implement me
	panic("implement me")
}

func (m mockEtherman) GetSuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	//TODO implement me
	panic("implement me")
}

func (m mockEtherman) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	//TODO implement me
	panic("implement me")
}

func (m mockEtherman) GetTx(ctx context.Context, txHash common.Hash) (*types.Transaction, bool, error) {
	//TODO implement me
	panic("implement me")
}

func (m mockEtherman) GetTxReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	//TODO implement me
	panic("implement me")
}

func (m mockEtherman) WaitTxToBeMined(ctx context.Context, tx *types.Transaction, timeout time.Duration) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (m mockEtherman) SendTx(ctx context.Context, tx *types.Transaction) error {
	//TODO implement me
	panic("implement me")
}

func (m mockEtherman) CurrentNonce(ctx context.Context, account common.Address) (uint64, error) {
	//TODO implement me
	panic("implement me")
}

func (m mockEtherman) SuggestedGasPrice(ctx context.Context) (*big.Int, error) {
	//TODO implement me
	panic("implement me")
}

func (m mockEtherman) EstimateGas(ctx context.Context, from common.Address, to *common.Address, value *big.Int, data []byte) (uint64, error) {
	//TODO implement me
	panic("implement me")
}

func (m mockEtherman) CheckTxWasMined(ctx context.Context, txHash common.Hash) (bool, *types.Receipt, error) {
	//TODO implement me
	panic("implement me")
}

func (m mockEtherman) SignTx(ctx context.Context, sender common.Address, tx *types.Transaction) (*types.Transaction, error) {
	//TODO implement me
	panic("implement me")
}

func (m mockEtherman) GetRevertMessage(ctx context.Context, tx *types.Transaction) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (m mockEtherman) GetZkEVMAddressAndL1ChainID() (common.Address, common.Address, uint64, error) {
	return common.HexToAddress(contractAddr), common.HexToAddress(contractAddrAgg), l1ChainID, nil
}
