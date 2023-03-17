package server

import "cosmos-gon-testnet/types"

func GetTransferNFTs() ([]*types.NFTTransactionPath, error) {
	allNfts, err := MonitorCli.DbCli.GetAllNFT()
	if err != nil {
		logger.Error("Failed to query all NFT class IDs， err:", err)
		return nil, err
	}

	nftsTransactionPath := make([]*types.NFTTransactionPath, 0)

	for _, nft := range allNfts {
		transferNFT, err1 := MonitorCli.DbCli.GetNFTTransfer(nft.ClassID, nft.NFTID)
		if err1 != nil {
			logger.Error("get mint NFT fail, err:", err1)
			return nil, err1
		}

		nftsTransactionPath = append(nftsTransactionPath, &types.NFTTransactionPath{
			BaseClassID:     nft.ClassID,
			NFTID:           nft.NFTID,
			TransferNFTPath: transferNFT,
		})
	}

	return nftsTransactionPath, nil
}

func GetTransferNFT(nfts []*types.AllNFT) ([]*types.NFTTransactionPath, error) {
	nftsTransactionPath := make([]*types.NFTTransactionPath, 0)
	for _, nft := range nfts {
		transferNFT, err := MonitorCli.DbCli.GetNFTTransfer(nft.ClassID, nft.NFTID)
		if err != nil {
			logger.Error("get mint NFT fail, err:", err)
			return nil, err
		}

		nftsTransactionPath = append(nftsTransactionPath, &types.NFTTransactionPath{
			BaseClassID:     nft.ClassID,
			NFTID:           nft.NFTID,
			TransferNFTPath: transferNFT,
		})
	}
	return nftsTransactionPath, nil
}
