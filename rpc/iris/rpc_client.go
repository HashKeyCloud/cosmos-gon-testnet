package iris

import (
	"fmt"

	/*base "cosmossdk.io/api/cosmos/base/tendermint/v1beta1"
	distribution "cosmossdk.io/api/cosmos/distribution/v1beta1"
	gov "cosmossdk.io/api/cosmos/gov/v1beta1"
	staking "cosmossdk.io/api/cosmos/staking/v1beta1"*/
	nft "cosmossdk.io/api/cosmos/nft/v1beta1"
	tx "cosmossdk.io/api/cosmos/tx/v1beta1"
	"github.com/spf13/viper"

	"cosmos-gon-testnet/log"
	"cosmos-gon-testnet/rpc"
)

type IrisCli struct {
	*rpc.ChainCli
}

// InitRpcCli Create a new RPC service
func InitIrisRpcCli() (*IrisCli, error) {
	endpoint := fmt.Sprintf("%s:%s", viper.GetString("gRpc.irisIp"), viper.GetString("gRpc.irisPort"))
	grpcConn, err := rpc.InitChainRpcCli(endpoint)
	if err != nil {
		logger.Error("Failed to create iris gRPC client, err:", err)
		return nil, err
	}

	/*stakingQueryCli := staking.NewQueryClient(grpcConn)
	govQueryCli := gov.NewQueryClient(grpcConn)
	baseCli := base.NewServiceClient(grpcConn)
	distributionCli := distribution.NewQueryClient(grpcConn)*/
	txCli := tx.NewServiceClient(grpcConn)
	nftCli := nft.NewQueryClient(grpcConn)

	return &IrisCli{
		ChainCli: &rpc.ChainCli{
			TxCli:  txCli,
			NFTCli: nftCli,
			/*StakingQueryCli: stakingQueryCli,
			GovQueryCli:     govQueryCli,
			BaseQuaryCli:    baseCli,
			DistributionCli: distributionCli,*/
		},
	}, err
}

var logger = log.RPCLogger.WithField("module", "rpc")
