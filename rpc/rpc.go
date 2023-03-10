package rpc

import (
	"context"
	"cosmos-gon-testnet/log"
	"cosmos-gon-testnet/types"
	abciv1beta1 "cosmossdk.io/api/cosmos/base/abci/v1beta1"
	base "cosmossdk.io/api/cosmos/base/tendermint/v1beta1"
	nft "cosmossdk.io/api/cosmos/nft/v1beta1"
	tx "cosmossdk.io/api/cosmos/tx/v1beta1"
	"encoding/json"
	"fmt"
	"strings"

	/*distribution "cosmossdk.io/api/cosmos/distribution/v1beta1"
	gov "cosmossdk.io/api/cosmos/gov/v1beta1"
	staking "cosmossdk.io/api/cosmos/staking/v1beta1"*/
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client interface {
}

type ChainCli struct {
	/*StakingQueryCli staking.QueryClient
	GovQueryCli     gov.QueryClient
	DistributionCli distribution.QueryClient*/
	BaseQuaryCli base.ServiceClient
	TxCli        tx.ServiceClient
	NFTCli       nft.QueryClient
}

func InitChainRpcCli(endpoint string) (*grpc.ClientConn, error) {
	dialOpts := []grpc.DialOption{
		// grpc.WithInsecure(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		// Maximum receive value 128 MB
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(128 * 1024 * 1024)),
	}
	// Create a connection to the gRPC server.
	grpcConn, err := grpc.Dial(endpoint, dialOpts...)
	if err != nil {
		logger.Error("Failed to create cosmos gRPC client, err:", err)
		return nil, err
	}
	return grpcConn, nil
}

/*func (cc *ChainCli) GetBlockHeight() int64 {
	queryValRankingRequest := &base.GetLatestValidatorSetRequest{}
	valSetResponse, err := cc.BaseQuaryCli.GetLatestValidatorSet(context.Background(), queryValRankingRequest)
	if err != nil {
		logger.Error("Query block height failed, err:", err)
	}
	return valSetResponse.BlockHeight
}*/

func (cc *ChainCli) GetIssueNFT() ([]*types.IssueNFT, error) {
	issueNFTs := make([]*types.IssueNFT, 0)
	txs, err := cc.getTxsEventRequest([]string{"message.action='/irismod.nft.MsgIssueDenom'"})
	if err != nil {
		logger.Error("get Issue NFT Txs Event fail,err:", err)
		return nil, err
	}
	for _, tx := range txs {
		var denomId, denomName, creator string
		for _, txLog := range tx.GetLogs() {
			for _, txEvent := range txLog.GetEvents() {
				for _, attributes := range txEvent.GetAttributes() {
					switch attributes.GetKey() {
					case "denom_id":
						denomId = attributes.GetValue()
					case "denom_name":
						denomName = attributes.GetValue()
					case "creator":
						creator = attributes.GetValue()
					}
				}
			}
		}
		issueNFTs = append(issueNFTs, &types.IssueNFT{
			BlockHeight: tx.GetHeight(),
			TxHash:      tx.GetTxhash(),
			ClassID:     denomId,
			Name:        denomName,
			Sender:      creator,
		})
	}
	return issueNFTs, nil
}

func (cc *ChainCli) GetMintNFT(classId ...string) ([]*types.MintNFT, error) {
	mintNFTs := make([]*types.MintNFT, 0)
	txs := make([]*abciv1beta1.TxResponse, 0)
	var err error
	switch len(classId) {
	case 0:
		txs, err = cc.getTxsEventRequest([]string{"message.action='/irismod.nft.MsgMintNFT'"})
		if err != nil {
			logger.Error("get Issue NFT Txs Event fail,err:", err)
			return nil, err
		}
	case 1:
		classIdEvent := "mint_nft.denom_id='" + classId[0] + "'"
		txs, err = cc.getTxsEventRequest([]string{"message.action='/irismod.nft.MsgMintNFT'", classIdEvent})
		if err != nil {
			logger.Error("get Issue NFT Txs Event fail,err:", err)
			return nil, err
		}
	default:
		logger.Error("Please enter a unique class ID or do not enter any value. If you enter a unique class ID, you will get mint transactions of the NFT, if you do not enter any value, you will get all mint events.")
	}

	for _, tx := range txs {
		var denomId, nftName, recipient, owner, tokenUri string
		for _, txLog := range tx.GetLogs() {
			for _, txEvent := range txLog.GetEvents() {

				for _, attributes := range txEvent.GetAttributes() {
					switch attributes.GetKey() {
					case "class_id":
						denomId = attributes.GetValue()
					case "id":
						nftName = attributes.GetValue()
					case "recipient":
						recipient = attributes.GetValue()
					case "owner":
						owner = attributes.GetValue()
					case "token_uri":
						tokenUri = attributes.GetValue()
					}
				}
			}
		}

		mintNFTs = append(mintNFTs, &types.MintNFT{
			BlockHeight: tx.GetHeight(),
			TxHash:      tx.GetTxhash(),
			ClassID:     denomId,
			NFTName:     nftName,
			Recipient:   recipient,
			Owner:       owner,
			TokenUri:    tokenUri,
		})
	}
	return mintNFTs, nil
}

func (cc *ChainCli) GetNFTTransferByContract(baseclassIds ...string) ([]*types.TransferNFT, error) {
	return cc.getNFTTransferCrossChain([]string{"message.action='/cosmwasm.wasm.v1.MsgExecuteContract'", "send_packet.packet_dst_port='nft-transfer'"}, baseclassIds...)
}

func (cc *ChainCli) GetNFTTransferByIBC(baseclassIds ...string) ([]*types.TransferNFT, error) {
	return cc.getNFTTransferCrossChain([]string{"message.action='/ibc.applications.nft_transfer.v1.MsgTransfer'"}, baseclassIds...)
}

func (cc *ChainCli) getNFTTransferCrossChain(events []string, baseclassIds ...string) ([]*types.TransferNFT, error) {
	transferNFTsByIBC := make([]*types.TransferNFT, 0)
	txs, err := cc.getTxsEventRequest(events)
	if err != nil {
		logger.Error("get NFT CrossChain Transfer Event fail,err:", err)
		return nil, err
	}
	for _, tx := range txs {
		var classId, baseClassId, sender, receiver string
		nftNames := make([]string, 0)
		for _, txLog := range tx.GetLogs() {
			for _, txEvent := range txLog.GetEvents() {
				for _, attributes := range txEvent.GetAttributes() {
					if attributes.GetKey() == "packet_data" {
						packetData := &types.PacketData{}
						err1 := json.Unmarshal([]byte(attributes.GetValue()), packetData)
						if err1 != nil {
							logger.Error("Failed to get transaction data，err:", err1)
							return nil, err1
						}
						classId = packetData.ClassID
						classIdSplit := strings.Split(classId, "/")
						baseClassId = classIdSplit[len(classIdSplit)-1]
						nftNames = packetData.TokenIds
						sender = packetData.Sender
						receiver = packetData.Receiver
					}
				}
			}
		}

		if len(baseclassIds) > 0 {
			clMap := make(map[string]struct{})
			for _, cl := range baseclassIds {
				clMap[cl] = struct{}{}
			}
			if _, ok := clMap[baseClassId]; ok {
				transferNFTsByIBC = append(transferNFTsByIBC, &types.TransferNFT{
					BlockHeight: tx.GetHeight(),
					TxHash:      tx.GetTxhash(),
					ClassID:     classId,
					BaseClassID: baseClassId,
					NFTName:     nftNames,
					Sender:      sender,
					Receiver:    receiver,
				})
			}
		} else {
			transferNFTsByIBC = append(transferNFTsByIBC, &types.TransferNFT{
				BlockHeight: tx.GetHeight(),
				TxHash:      tx.GetTxhash(),
				ClassID:     classId,
				BaseClassID: baseClassId,
				NFTName:     nftNames,
				Sender:      sender,
				Receiver:    receiver,
			})
		}
	}
	return transferNFTsByIBC, nil
}

func (cc *ChainCli) GetNFTTransfer(baseclassIds ...string) ([]*types.TransferNFT, error) {
	transferNFTs := make([]*types.TransferNFT, 0)
	txs, err := cc.getTxsEventRequest([]string{"message.action='/irismod.nft.MsgTransferNFT'"})
	if err != nil {
		logger.Error("get NFT CrossChain Transfer Event fail,err:", err)
		return nil, err
	}

	for _, tx := range txs {
		var classId, baseClassId, sender, receiver string
		nftName := make([]string, 0)
		for _, txLog := range tx.GetLogs() {
			for _, txEvent := range txLog.GetEvents() {
				for _, attributes := range txEvent.GetAttributes() {
					switch attributes.GetKey() {
					case "denom_id":
						classId = attributes.GetValue()
						classIdSplit := strings.Split(classId, "/")
						baseClassId = classIdSplit[len(classIdSplit)-1]
					case "token_id":
						nftName = []string{attributes.GetValue()}
					case "sender":
						sender = attributes.GetValue()
					case "recipient":
						receiver = attributes.GetValue()
					}
				}
			}
		}
		if len(baseclassIds) > 0 {
			clMap := make(map[string]struct{})
			for _, cl := range baseclassIds {
				clMap[cl] = struct{}{}
			}
			if _, ok := clMap[baseClassId]; ok {
				transferNFTs = append(transferNFTs, &types.TransferNFT{
					BlockHeight: tx.GetHeight(),
					TxHash:      tx.GetTxhash(),
					ClassID:     classId,
					BaseClassID: baseClassId,
					NFTName:     nftName,
					Sender:      sender,
					Receiver:    receiver,
				})
			}
		} else {
			transferNFTs = append(transferNFTs, &types.TransferNFT{
				BlockHeight: tx.GetHeight(),
				TxHash:      tx.GetTxhash(),
				ClassID:     classId,
				BaseClassID: baseClassId,
				NFTName:     nftName,
				Sender:      sender,
				Receiver:    receiver,
			})
		}
	}

	return transferNFTs, nil
}

func (cc *ChainCli) GetTxByHash(hash string) {
	txRequest := &tx.GetTxRequest{
		Hash: hash,
	}
	tx, err := cc.TxCli.GetTx(context.Background(), txRequest)
	if err != nil {
		logger.Error("get tx by hash fail,err:", err)
	}
	fmt.Println("Tx:", tx.GetTxResponse())
}

func (cc *ChainCli) getTxsEventRequest(events []string) ([]*abciv1beta1.TxResponse, error) {
	var page uint64 = 1
	txs := make([]*abciv1beta1.TxResponse, 0)
	for {
		txEventRequest := &tx.GetTxsEventRequest{
			Events: events,
			Page:   page,
			Limit:  100,
		}
		txsResponse, err := cc.TxCli.GetTxsEvent(context.Background(), txEventRequest)
		if err != nil {
			logger.Error("get tx fail, err:", err)
			return nil, err
		}
		txs = append(txs, txsResponse.GetTxResponses()...)
		if txsResponse.Total/100+1 <= page || len(txsResponse.Txs) <= 0 || txsResponse.Txs == nil {
			break
		} else {
			page++
		}
	}
	return txs, nil
}

var logger = log.RPCLogger.WithField("module", "rpc")
