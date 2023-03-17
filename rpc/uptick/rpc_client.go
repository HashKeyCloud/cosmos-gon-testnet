package uptick

import (
	"fmt"

	nft "cosmossdk.io/api/cosmos/nft/v1beta1"
	tx "cosmossdk.io/api/cosmos/tx/v1beta1"
	"github.com/spf13/viper"

	"cosmos-gon-testnet/log"
	"cosmos-gon-testnet/rpc"
)

type UptickCli struct {
	*rpc.ChainCli
}

// InitRpcCli Create a new RPC server
func InitUptickRpcCli() (*UptickCli, error) {
	endpoint := fmt.Sprintf("%s:%s", viper.GetString("gRpc.uptickIp"), viper.GetString("gRpc.uptickPort"))
	grpcConn, err := rpc.InitChainRpcCli(endpoint)
	if err != nil {
		logger.Error("Failed to create uptick gRPC client, err:", err)
		return nil, err
	}

	txCli := tx.NewServiceClient(grpcConn)
	nftCli := nft.NewQueryClient(grpcConn)

	return &UptickCli{
		ChainCli: &rpc.ChainCli{
			TxCli:  txCli,
			NFTCli: nftCli,
		},
	}, err
}

var logger = log.RPCLogger.WithField("module", "rpc")
