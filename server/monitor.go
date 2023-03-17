package server

import (
	"cosmos-gon-testnet/db"
	"cosmos-gon-testnet/log"
	"cosmos-gon-testnet/rpc"
	"cosmos-gon-testnet/rpc/iris"
	"cosmos-gon-testnet/rpc/juno"
	"cosmos-gon-testnet/rpc/omniflix"
	"cosmos-gon-testnet/rpc/stargaze"
	"cosmos-gon-testnet/rpc/uptick"
	"cosmos-gon-testnet/types"
	"cosmos-gon-testnet/utils"
	"github.com/spf13/viper"
	"time"
)

var (
	MonitorCli *Monitor
)

type Monitor struct {
	RpcCli map[string]rpc.Client
	/*IrisRpcCli     rpc.Client
	JunoRpcCli     rpc.Client
	OmniflixRpcCli rpc.Client
	StargazeRpcCli rpc.Client
	UptickRpcCli   rpc.Client*/
	DbCli db.DBCli
}

func InitMonitor() {
	rpcCli := make(map[string]rpc.Client)
	irisRpcCli, err := iris.InitIrisRpcCli()
	if err != nil {
		logger.Error("connect iris rpc client error:", err)
		return
	}
	rpcCli["iris"] = irisRpcCli

	junoRpcCli, err := juno.InitJunoRpcCli()
	if err != nil {
		logger.Error("connect juno rpc client error:", err)
		return
	}
	rpcCli["juno"] = junoRpcCli

	omniflixRpcCli, err := omniflix.InitOmniflixRpcCli()
	if err != nil {
		logger.Error("connect omniflix rpc client error:", err)
		return
	}
	rpcCli["omniflix"] = omniflixRpcCli

	stargazeRpcCli, err := stargaze.InitStargazeRpcCli()
	if err != nil {
		logger.Error("connect stargaze rpc client error:", err)
		return
	}
	rpcCli["stargaze"] = stargazeRpcCli

	uptickRpcCli, err := uptick.InitUptickRpcCli()
	if err != nil {
		logger.Error("connect uptick rpc client error:", err)
		return
	}
	rpcCli["uptick"] = uptickRpcCli

	dbCli, err := db.InitDB()
	if err != nil {
		logger.Error("connect db client error:", err)
		return
	}

	MonitorCli = &Monitor{
		RpcCli: rpcCli,
		DbCli:  dbCli,
	}

	logger.Info("init monitor successful")
}

func (m *Monitor) StartMonitor() {
	epochTicker := time.NewTicker(time.Duration(viper.GetInt("monitor.timeInterval")) * time.Second)
	for range epochTicker.C {
		for chainName, rpcCli := range m.RpcCli {
			logger.Infof("Start getting events on %v\n", chainName)
			issueNFTs := make([]*types.IssueNFT, 0)
			mintNFTs := make([]*types.MintNFT, 0)
			transferNFTs := make([]*types.TransferNFT, 0)

			issueNFT, err := rpcCli.GetIssueNFT()
			if err != nil {
				res := utils.Retry(func() bool {
					issueNFT, err = rpcCli.GetIssueNFT()
					if err != nil {
						return false
					} else {
						return true
					}
				}, []int{1, 3})
				if !res {
					logger.Errorf("Failed to get issue NFT event on %v, err:%v\n", chainName, err)
				}
			}
			issueNFTs = append(issueNFTs, issueNFT...)

			mintNFT, err := rpcCli.GetMintNFT()
			if err != nil {
				res := utils.Retry(func() bool {
					mintNFT, err = rpcCli.GetMintNFT()
					if err != nil {
						return false
					} else {
						return true
					}
				}, []int{1, 3})
				if !res {
					logger.Errorf("Failed to get mint NFT event on %v, err:%v\n", chainName, err)
				}
			}

			for _, mintNft := range mintNFT {
				transferNFTs = append(transferNFTs, &types.TransferNFT{
					BlockHeight: mintNft.BlockHeight,
					Timestamp:   mintNft.Timestamp,
					TxHash:      mintNft.TxHash,
					ClassID:     mintNft.ClassID,
					BaseClassID: mintNft.ClassID,
					NFTID:       []string{mintNft.NFTID},
					Sender:      "0x",
					Receiver:    mintNft.Owner,
				})
			}
			mintNFTs = append(mintNFTs, mintNFT...)

			if chainName == "juno" || chainName == "stargaze" {
				transferNFT, err := rpcCli.GetNFTTransferByContract()
				if err != nil {
					res := utils.Retry(func() bool {
						transferNFT, err = rpcCli.GetNFTTransferByContract()
						if err != nil {
							return false
						} else {
							return true
						}
					}, []int{1, 3})
					if !res {
						logger.Errorf("Failed to get transfer NFT event on %v, err:%v\n", chainName, err)
					}
				}
				transferNFTs = append(transferNFTs, transferNFT...)
			} else {
				transferNFT, err := rpcCli.GetNFTTransferByIBC()
				if err != nil {
					res := utils.Retry(func() bool {
						transferNFT, err = rpcCli.GetNFTTransferByIBC()
						if err != nil {
							return false
						} else {
							return true
						}
					}, []int{1, 3})
					if !res {
						logger.Errorf("Failed to get transfer NFT event on %v, err:%v\n", chainName, err)
					}
				}
				transferNFTs = append(transferNFTs, transferNFT...)
			}
			err1 := m.DbCli.BatchSaveIssueNFT(issueNFTs)
			if err1 != nil {
				logger.Error("Failed to save issue NFT event, err:", err1)
			} else {
				logger.Info("Save issue NFT event successfully")
			}

			err2 := m.DbCli.BatchSaveMintNFT(mintNFTs)
			if err2 != nil {
				logger.Error("Failed to save mint NFT event, err:", err2)
			} else {
				logger.Info("Save mint NFT event successfully")
			}

			err3 := m.DbCli.BatchSaveTransferNFT(transferNFTs)
			if err3 != nil {
				logger.Error("Failed to save transfer NFT event, err:", err3)
			} else {
				logger.Info("Save transfer NFT event successfully")
			}
		}

	}
}

var logger = log.ServerLogger.WithField("module", "service")
