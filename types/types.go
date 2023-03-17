package types

import "time"

type DatabaseConfig struct {
	Username string
	Password string
	Name     string
	Host     string
	Port     string
}

type IssueNFT struct {
	BlockHeight int64     `db:"block_height"`
	Timestamp   time.Time `db:"timestamp"`
	TxHash      string    `db:"tx_hash"`
	ClassID     string    `db:"class_id"`
	Name        string    `db:"name"`
	Sender      string    `db:"sender"`
}

type MintNFT struct {
	BlockHeight int64     `db:"block_height"`
	Timestamp   time.Time `db:"timestamp"`
	TxHash      string    `db:"tx_hash"`
	ClassID     string    `db:"class_id"`
	NFTID       string    `db:"nft_id"`
	Recipient   string    `db:"recipient"`
	Owner       string    `db:"owner"`
	TokenUri    string    `db:"token_uri"`
}
type AllNFT struct {
	ClassID string `db:"class_id" uri:"class_id" binding:"required"`
	NFTID   string `db:"nft_id" uri:"class_id" binding:"required"`
}

type TransferNFT struct {
	BlockHeight int64     `db:"block_height"`
	Timestamp   time.Time `db:"timestamp"`
	TxHash      string    `db:"tx_hash"`
	ClassID     string    `db:"class_id"`
	BaseClassID string    `db:"base_class_id"`
	NFTID       []string  `db:"nft_id"`
	Sender      string    `db:"sender"`
	Receiver    string    `db:"receiver"`
}

type PacketData struct {
	ClassData string   `json:"classData"`
	ClassID   string   `json:"classId"`
	ClassURI  string   `json:"classUri"`
	Receiver  string   `json:"receiver"`
	Sender    string   `json:"sender"`
	TokenData []string `json:"tokenData"`
	TokenIds  []string `json:"tokenIds"`
	TokenUris []string `json:"tokenUris"`
}

type NFTTransactionPath struct {
	BaseClassID     string
	NFTID           string
	TransferNFTPath []string
}

type NFTWhale struct {
	BaseClassID string `db:"base_class_id"`
}
