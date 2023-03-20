package db

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"

	"cosmos-gon-testnet/log"
	"cosmos-gon-testnet/types"
)

type DBCli interface {
	BatchSaveIssueNFT(issueNFTs []*types.IssueNFT) error
	BatchSaveMintNFT(mintNFTs []*types.MintNFT) error
	BatchSaveTransferNFT(TransferNFTs []*types.TransferNFT) error
	GetIssueNFTFromDb() ([]*types.IssueNFT, error)
	GetMintNFTFromDb() ([]*types.MintNFT, error)
	GetNFTTransfer(baseClassId, nftId string) ([]string, error)
	GetAllNFT() ([]*types.AllNFT, error)
}

type DbCli struct {
	Conn *sqlx.DB
}

func InitDB() (*DbCli, error) {
	port, _ := strconv.Atoi(viper.GetString("postgres.port"))
	pdqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		viper.GetString("postgres.host"), port, viper.GetString("postgres.user"),
		viper.GetString("postgres.password"), viper.GetString("postgres.name"))
	db, err := sqlx.Connect("postgres", pdqlInfo)
	if err != nil {
		logger.Errorf("Connected database failed. err:%v\n", err)
		return nil, err
	}

	dbConnectionTimeout := time.NewTimer(15 * time.Second)
	go func() {
		<-dbConnectionTimeout.C
		logger.Fatalf("timeout while connecting to the database")
	}()
	err = db.Ping()
	if err != nil {
		logger.Errorf("ping db fail, err:%v", err)
	}

	dbConnectionTimeout.Stop()

	db.SetConnMaxIdleTime(time.Second * 30)
	db.SetConnMaxLifetime(time.Second * 60)
	db.SetMaxOpenConns(200)
	db.SetMaxIdleConns(200)

	logger.Info("Successfully connected database!")
	return &DbCli{
		Conn: db,
	}, nil
}

func (dc *DbCli) BatchSaveIssueNFT(issueNFTs []*types.IssueNFT) error {
	batchSize := 500
	for b := 0; b < len(issueNFTs); b += batchSize {
		logger.Infof("Start saving %d batch of IssueNFT\n", b+1)
		start := b
		end := b + batchSize
		if len(issueNFTs) < end {
			end = len(issueNFTs)
		}
		numArgs := 6
		valueStrings := make([]string, 0, batchSize)
		valueArgs := make([]interface{}, 0, batchSize*numArgs)

		for i, v := range issueNFTs[start:end] {
			valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d)",
				i*numArgs+1, i*numArgs+2, i*numArgs+3, i*numArgs+4, i*numArgs+5, i*numArgs+6))
			valueArgs = append(valueArgs, v.BlockHeight)
			valueArgs = append(valueArgs, v.Timestamp)
			valueArgs = append(valueArgs, v.TxHash)
			valueArgs = append(valueArgs, v.ClassID)
			valueArgs = append(valueArgs, v.Name)
			valueArgs = append(valueArgs, v.Sender)
		}

		sql := fmt.Sprintf(`
			INSERT INTO issue_NFT (
				block_height,
				timestamp,
				tx_hash,
				class_id,
				name,
				sender
			)
			VALUES %v
			ON  CONFLICT (tx_hash) DO UPDATE SET
				block_height = EXCLUDED.block_height,
				timestamp = EXCLUDED.timestamp,
				class_id = EXCLUDED.class_id,
				name = EXCLUDED.name,
				sender = EXCLUDED.sender;
		`, strings.Join(valueStrings, ","))
		_, err := dc.Conn.Exec(sql, valueArgs...)
		if err != nil {
			logger.Errorf("saving IssueNFT batch %v fail. err:%v \n", b+1, err)
			return err
		}

		logger.Infof("saving IssueNFT %v completed\n", b+1)
	}
	return nil
}

func (dc *DbCli) BatchSaveMintNFT(mintNFTs []*types.MintNFT) error {
	batchSize := 500
	for b := 0; b < len(mintNFTs); b += batchSize {
		logger.Infof("Start saving %d batch of MintNFT\n", b+1)
		start := b
		end := b + batchSize
		if len(mintNFTs) < end {
			end = len(mintNFTs)
		}
		numArgs := 8
		valueStrings := make([]string, 0, batchSize)
		valueArgs := make([]interface{}, 0, batchSize*numArgs)

		for i, v := range mintNFTs[start:end] {
			valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d)",
				i*numArgs+1, i*numArgs+2, i*numArgs+3, i*numArgs+4, i*numArgs+5, i*numArgs+6, i*numArgs+7, i*numArgs+8))
			valueArgs = append(valueArgs, v.BlockHeight)
			valueArgs = append(valueArgs, v.Timestamp)
			valueArgs = append(valueArgs, v.TxHash)
			valueArgs = append(valueArgs, v.ClassID)
			valueArgs = append(valueArgs, v.NFTID)
			valueArgs = append(valueArgs, v.Recipient)
			valueArgs = append(valueArgs, v.Owner)
			valueArgs = append(valueArgs, v.TokenUri)
		}

		sql := fmt.Sprintf(`
			INSERT INTO mint_NFT (
				block_height,
				timestamp,
				tx_hash,
				class_id,
				nft_id,
				recipient,
				owner,
				token_uri
			)
			VALUES %v
			ON  CONFLICT (tx_hash) DO UPDATE SET
				block_height = EXCLUDED.block_height,
				timestamp = EXCLUDED.timestamp,
				class_id = EXCLUDED.class_id,
				nft_id = EXCLUDED.nft_id,
				recipient = EXCLUDED.recipient,
				owner = EXCLUDED.owner,
				token_uri = EXCLUDED.token_uri;
		`, strings.Join(valueStrings, ","))
		_, err := dc.Conn.Exec(sql, valueArgs...)
		if err != nil {
			logger.Errorf("saving MintNFT batch %v fail. err:%v \n", b+1, err)
			return err
		}

		logger.Infof("saving MintNFT %v completed\n", b+1)
	}
	return nil
}

func (dc *DbCli) BatchSaveTransferNFT(TransferNFTs []*types.TransferNFT) error {
	batchSize := 500
	for b := 0; b < len(TransferNFTs); b += batchSize {
		logger.Infof("Start saving %d batch of TransferNFT\n", b+1)
		start := b
		end := b + batchSize
		if len(TransferNFTs) < end {
			end = len(TransferNFTs)
		}
		numArgs := 8
		valueStrings := make([]string, 0, batchSize)
		valueArgs := make([]interface{}, 0, batchSize*numArgs)

		for i, v := range TransferNFTs[start:end] {
			valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d)",
				i*numArgs+1, i*numArgs+2, i*numArgs+3, i*numArgs+4, i*numArgs+5, i*numArgs+6, i*numArgs+7, i*numArgs+8))
			valueArgs = append(valueArgs, v.BlockHeight)
			valueArgs = append(valueArgs, v.Timestamp)
			valueArgs = append(valueArgs, v.TxHash)
			valueArgs = append(valueArgs, v.ClassID)
			valueArgs = append(valueArgs, v.BaseClassID)
			valueArgs = append(valueArgs, v.NFTID[0])
			valueArgs = append(valueArgs, v.Sender)
			valueArgs = append(valueArgs, v.Receiver)
		}

		sql := fmt.Sprintf(`
			INSERT INTO transfer_NFT (
				block_height,
				timestamp,
				tx_hash,
				class_id,
				base_class_id,
				nft_id,
				sender,
				receiver
			)
			VALUES %v
			ON  CONFLICT (tx_hash, base_class_id, nft_id) DO UPDATE SET
				block_height = EXCLUDED.block_height,
				timestamp = EXCLUDED.timestamp,
				class_id = EXCLUDED.class_id,
				sender = EXCLUDED.sender,
				receiver = EXCLUDED.receiver;
		`, strings.Join(valueStrings, ","))
		_, err := dc.Conn.Exec(sql, valueArgs...)
		if err != nil {
			logger.Errorf("saving TransferNFT batch %v fail. err:%v \n", b+1, err)
			return err
		}

		logger.Infof("saving TransferNFT %v completed\n", b+1)
	}
	return nil
}

func (dc *DbCli) GetIssueNFTFromDb() ([]*types.IssueNFT, error) {
	issueNFTs := make([]*types.IssueNFT, 0)
	sqld := `select * from issue_NFT;`
	err := dc.Conn.Select(&issueNFTs, sqld)
	if err != nil {
		logger.Errorf("Failed to query Issue NFT from db, err:%v\n", err)
		return nil, err
	}
	logger.Infof("query issue NFT from db successful")
	return issueNFTs, nil
}

func (dc *DbCli) GetMintNFTFromDb() ([]*types.MintNFT, error) {
	mintsNFTs := make([]*types.MintNFT, 0)
	sqld := `select * from mint_NFT;`
	err := dc.Conn.Select(&mintsNFTs, sqld)
	if err != nil {
		logger.Errorf("Failed to query Mint NFT from db, err:%v\n", err)
		return nil, err
	}
	logger.Info("query mint NFT from db successful")
	return mintsNFTs, nil
}

func (dc *DbCli) GetNFTTransfer(baseClassId, nftId string) ([]string, error) {
	transfer := make([]string, 0)
	sqld := `select receiver from transfer_NFT WHERE base_class_id = $1 AND nft_id = $2  ORDER BY transfer_NFT.timestamp;`
	err := dc.Conn.Select(&transfer, sqld, baseClassId, nftId)
	if err != nil {
		logger.Errorf("Failed to query NFT transfer info from db, err:%v\n", err)
		return nil, err
	}
	logger.Info("query NFT transfer info from db successful")
	return transfer, nil
}

func (dc *DbCli) GetAllNFT() ([]*types.AllNFT, error) {
	addNFTs := make([]*types.AllNFT, 0)
	sqld := `select class_id, nft_id from mint_NFT;`
	err := dc.Conn.Select(&addNFTs, sqld)
	if err != nil {
		logger.Errorf("Failed to query All NFT from db, err:%v\n", err)
		return nil, err
	}
	logger.Info("query All NFT from db successful")
	return addNFTs, nil
}

/*func (dc *DbCli) GetNFTWhale(baseClassId string) ()*/

var logger = log.DBLogger.WithField("module", "db")
