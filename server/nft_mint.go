package server

import (
	"cosmos-gon-testnet/types"
	"errors"
)

func GetMintNFTs() ([]*types.MintNFT, error) {
	mintNFT, err := MonitorCli.DbCli.GetMintNFTFromDb()
	if err != nil {
		logger.Error("get mint NFT fail, err:", err)
		return nil, err
	}
	return mintNFT, nil
}

func GetMintNFT(classIds []string) ([]*types.MintNFT, error) {
	if len(classIds) <= 0 {
		return nil, errors.New("Need to specify Class ID")
	}
	nfts := make([]*types.MintNFT, 0)
	mintNFTs, err := MonitorCli.DbCli.GetMintNFTFromDb()
	if err != nil {
		logger.Error("get mint NFT fail, err:", err)
		return nil, err
	}
	nftClassId := make(map[string]struct{}, 0)
	for _, classId := range classIds {
		nftClassId[classId] = struct{}{}
	}

	for _, nft := range mintNFTs {
		if _, ok := nftClassId[nft.ClassID]; ok {
			nfts = append(nfts, nft)
		}
	}

	return nfts, nil
}
