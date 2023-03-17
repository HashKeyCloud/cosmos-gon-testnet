package iris

import (
	"cosmos-gon-testnet/rpc"
	base "cosmossdk.io/api/cosmos/base/tendermint/v1beta1"
	nft "cosmossdk.io/api/cosmos/nft/v1beta1"
	tx "cosmossdk.io/api/cosmos/tx/v1beta1"
	"fmt"
	"testing"
)

func TestIssueNFT(t *testing.T) {
	grpcConn, err := rpc.InitChainRpcCli("34.80.93.133:9090")
	if err != nil {
		logger.Error("Failed to create evmos gRPC client, err:", err)
	}

	txCli := tx.NewServiceClient(grpcConn)
	nftCli := nft.NewQueryClient(grpcConn)

	cc := &IrisCli{
		ChainCli: &rpc.ChainCli{
			TxCli:  txCli,
			NFTCli: nftCli,
		},
	}
	nfts, err := cc.GetIssueNFT()
	if err != nil {
		fmt.Println("error:", err)
	}
	for _, nft := range nfts {
		fmt.Println("nft:", nft)
	}
}

func TestGetMintNFT(t *testing.T) {
	grpcConn, err := rpc.InitChainRpcCli("34.80.93.133:9090")
	if err != nil {
		logger.Error("Failed to create evmos gRPC client, err:", err)
	}

	txCli := tx.NewServiceClient(grpcConn)
	nftCli := nft.NewQueryClient(grpcConn)
	baseCli := base.NewServiceClient(grpcConn)

	cc := &IrisCli{
		ChainCli: &rpc.ChainCli{
			BaseQuaryCli: baseCli,
			TxCli:        txCli,
			NFTCli:       nftCli,
		},
	}
	nfts, err := cc.GetMintNFT()
	if err != nil {
		fmt.Println("error:", err)
	}
	for _, nft := range nfts {
		fmt.Println("nft:", nft)
	}
}

func TestNFTTransferByIBC(t *testing.T) {
	grpcConn, err := rpc.InitChainRpcCli("34.80.93.133:9090")
	if err != nil {
		logger.Error("Failed to create evmos gRPC client, err:", err)
	}

	txCli := tx.NewServiceClient(grpcConn)
	nftCli := nft.NewQueryClient(grpcConn)

	cc := &IrisCli{
		ChainCli: &rpc.ChainCli{
			TxCli:  txCli,
			NFTCli: nftCli,
		},
	}

	txs, err := cc.GetNFTTransferByIBC()
	if err != nil {
		fmt.Println("get nfr transfer by ibc fail, err:", err)
	}
	for _, tx := range txs {
		fmt.Println("nft transfer by ibc:", tx)
	}
}

func TestNFTTransferOnIris(t *testing.T) {
	grpcConn, err := rpc.InitChainRpcCli("34.80.93.133:9090")
	if err != nil {
		logger.Error("Failed to create evmos gRPC client, err:", err)
	}

	txCli := tx.NewServiceClient(grpcConn)
	nftCli := nft.NewQueryClient(grpcConn)

	cc := &IrisCli{
		ChainCli: &rpc.ChainCli{
			TxCli:  txCli,
			NFTCli: nftCli,
		},
	}

	txs, err := cc.GetNFTTransferOnIris()
	if err != nil {
		fmt.Println("get nfr transfer by ibc fail, err:", err)
	}
	for _, tx := range txs {
		fmt.Println("nft transfer:", tx)
	}
}

func TestNFTTransferOnOmniFlix(t *testing.T) {
	grpcConn, err := rpc.InitChainRpcCli("65.21.93.56:9090")
	if err != nil {
		logger.Error("Failed to create evmos gRPC client, err:", err)
	}

	txCli := tx.NewServiceClient(grpcConn)
	nftCli := nft.NewQueryClient(grpcConn)

	cc := &IrisCli{
		ChainCli: &rpc.ChainCli{
			TxCli:  txCli,
			NFTCli: nftCli,
		},
	}

	txs, err := cc.GetNFTTransferOnOmniflix()
	if err != nil {
		fmt.Println("get nfr transfer by ibc fail, err:", err)
	}
	for _, tx := range txs {
		fmt.Println("nft transfer:", tx)
	}
}

func TestNFTTransferByContract(t *testing.T) {
	grpcConn, err := rpc.InitChainRpcCli("juno-testnet-grpc.polkachu.com:12690")
	if err != nil {
		logger.Error("Failed to create juno gRPC client, err:", err)
	}

	txCli := tx.NewServiceClient(grpcConn)
	nftCli := nft.NewQueryClient(grpcConn)

	cc := &IrisCli{
		ChainCli: &rpc.ChainCli{
			TxCli:  txCli,
			NFTCli: nftCli,
		},
	}

	txs, err := cc.GetNFTTransferByContract()
	if err != nil {
		fmt.Println("get nfr transfer by ibc fail, err:", err)
	}
	for _, tx := range txs {
		fmt.Println("nft transfer by ibc:", tx)
	}
}

/*func TestTxByHash(t *testing.T) {
	grpcConn, err := rpc.InitChainRpcCli("34.80.93.133:9090")
	if err != nil {
		logger.Error("Failed to create evmos gRPC client, err:", err)
	}

	txCli := tx.NewServiceClient(grpcConn)
	nftCli := nft.NewQueryClient(grpcConn)

	cc := &IrisCli{
		ChainCli: &rpc.ChainCli{
			TxCli:  txCli,
			NFTCli: nftCli,
		},
	}
	cc.GetTxByHash("CB3BD702DFF7FA8B819AB6FC6E277F3766908E7CB68C31FFC5542012FDFC6760")
}*/
