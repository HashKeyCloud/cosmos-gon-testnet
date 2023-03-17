package server

import "cosmos-gon-testnet/types"

func GetIssueNFT() ([]*types.IssueNFT, error) {
	issueNFT, err := MonitorCli.DbCli.GetIssueNFTFromDb()
	if err != nil {
		logger.Error("get issue NFT fail, err:", err)
		return nil, err
	}
	return issueNFT, nil
}
