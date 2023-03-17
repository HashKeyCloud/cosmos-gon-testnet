package main

import (
	"cosmos-gon-testnet/handler"
	"cosmos-gon-testnet/middelwares"
	"cosmos-gon-testnet/server"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"

	"cosmos-gon-testnet/config"
	_ "cosmos-gon-testnet/docs"
	"cosmos-gon-testnet/log"
)

func init() {
	var cfg = pflag.StringP("config", "c", "config/.conf.yaml", "config file path.")
	pflag.Parse()
	err := config.InitConfig(*cfg)
	if err != nil {
		fmt.Println("read config file error:", err)
	}
	log.InitLogger(log.Logger, viper.GetString("log.path"))
	log.InitLogger(log.RPCLogger, viper.GetString("log.path"))
	log.InitLogger(log.MailLogger, viper.GetString("log.path"))
	log.InitLogger(log.HandlerLogger, viper.GetString("log.path"))
	log.InitLogger(log.ServerLogger, viper.GetString("log.path"))
	log.InitLogger(log.DBLogger, viper.GetString("log.path"))

	server.InitMonitor()
}

// @title Cosmos GoN API
// @version 1.0
// @description This is a Cosmos GoN dashboard.
// @termsOfService http://hashquark.io

// @contact.name API Support
// @contact.url http://www.hashquark.io
// @contact.email wangzezheng@hashquark.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8080
// @BasePath

func main() {
	logger := log.Logger.WithField("module", "main")
	logger.Info("Successfully read config file!")
	go server.MonitorCli.StartMonitor()

	r := gin.New()
	r.Use(middelwares.Cors())
	// http://127.0.0.1:8080/swagger/index.html
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.GET("/server/issueNFT", handler.GetIssueNFT)
	r.GET("/server/mintNFTs", handler.GetMintNFTs)
	r.GET("/server/mintNFT", handler.GetMintNFT)
	r.GET("/server/transferNFTs", handler.GetTransferNFTs)
	r.POST("/server/transferNFT", handler.GetTransferNFT)

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
	http.ListenAndServe(":8081", nil).Error()
}
