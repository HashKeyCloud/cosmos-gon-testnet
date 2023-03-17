package iris

/*import (
	"cosmos-gon-testnet/db"
	"cosmos-gon-testnet/log"
	"cosmos-gon-testnet/types"
	"github.com/spf13/viper"
)

type DbCli struct {
	*db.DbCli
}

func InitDbCli() (*DbCli, error) {
	dbConf := &types.DatabaseConfig{
		Username: viper.GetString("postgres.User"),
		Password: viper.GetString("postgres.Password"),
		Name:     viper.GetString("postgres.Name"),
		Host:     viper.GetString("postgres.Host"),
		Port:     viper.GetString("postgres.Port"),
	}
	dbCli, err := db.InitDB(dbConf)
	if err != nil {
		logger.Error("connect cosmos database server error: ", err)
		return nil, err
	}

	return &DbCli{
		DbCli: dbCli,
	}, nil
}

var logger = log.DBLogger.WithField("module", "db")*/
