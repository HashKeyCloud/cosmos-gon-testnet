package omniflix

import (
	"fmt"

	nft "cosmossdk.io/api/cosmos/nft/v1beta1"
	tx "cosmossdk.io/api/cosmos/tx/v1beta1"
	"github.com/spf13/viper"

	"cosmos-gon-testnet/log"
	"cosmos-gon-testnet/rpc"
)

type OmniflixCli struct {
	*rpc.ChainCli
}

// InitRpcCli Create a new RPC server
func InitOmniflixRpcCli() (*OmniflixCli, error) {
	endpoint := fmt.Sprintf("%s:%s", viper.GetString("gRpc.omniflixIp"), viper.GetString("gRpc.omniflixPort"))
	grpcConn, err := rpc.InitChainRpcCli(endpoint)
	if err != nil {
		logger.Error("Failed to create omniflix gRPC client, err:", err)
		return nil, err
	}

	txCli := tx.NewServiceClient(grpcConn)
	nftCli := nft.NewQueryClient(grpcConn)

	return &OmniflixCli{
		ChainCli: &rpc.ChainCli{
			TxCli:  txCli,
			NFTCli: nftCli,
		},
	}, err
}

var logger = log.RPCLogger.WithField("module", "rpc")
