package handler

import (
	"cosmos-gon-testnet/log"
	"cosmos-gon-testnet/server"
	"cosmos-gon-testnet/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Get NFT Issue events
// @Description get NFT issue events
// @Tags Server
// @Accept json
// @Produce json
// @Success 200 {object} Response
// @Fail	400 {object}  string "fail"
// @Router /server/issueNFT [get]
func GetIssueNFT(c *gin.Context) {
	issueNFT, err := server.GetIssueNFT()
	if err != nil {
		logger.Error("get nft issue events fail, err:", err)
		sendResponse(c, err, nil)
	}
	sendResponse(c, nil, issueNFT)
}

// @Summary Get All NFT Mint events
// @Description get all NFT mint events
// @Tags Server
// @Accept json
// @Produce json
// @Success 200 {object} Response
// @Fail	400 {object}  string "fail"
// @Router /server/mintNFTs [get]
func GetMintNFTs(c *gin.Context) {
	mintNFT, err := server.GetMintNFTs()
	if err != nil {
		logger.Error("get all nft mint events fail, err:", err)
		sendResponse(c, err, nil)
	}
	sendResponse(c, nil, mintNFT)
}

// @Summary Get NFT Mint events
// @Description get NFT mint events
// @Tags Server
// @Param classId query []string true "classId"
// @Accept json
// @Produce json
// @Success 200 {object} Response
// @Fail	400 {object}  string "fail"
// @Router /server/mintNFT [get]
func GetMintNFT(c *gin.Context) {
	mintNFT, err := server.GetMintNFT(c.QueryArray("classId"))
	if err != nil {
		logger.Error("get all nft mint events fail, err:", err)
		sendResponse(c, err, nil)
	}
	sendResponse(c, nil, mintNFT)
}

// @Summary Get All NFT Transfer Path
// @Description get all NFT transfer path
// @Tags Server
// @Accept json
// @Produce json
// @Success 200 {object} Response
// @Fail	400 {object}  string "fail"
// @Router /server/transferNFTs [get]
func GetTransferNFTs(c *gin.Context) {
	tranferNFTs, err := server.GetTransferNFTs()
	if err != nil {
		logger.Error("get all nft transfer events fail, err:", err)
		sendResponse(c, err, nil)
	}
	sendResponse(c, nil, tranferNFTs)
}

// @Summary Get NFT Mint events
// @Description get NFT mint events
// @Tags Server
// @Param nfts body []types.AllNFT true  "nfts"
// @Accept json
// @Produce json
// @Success 200 {object} Response
// @Fail	400 {object}  string "fail"
// @Router /server/transferNFT [post]
func GetTransferNFT(c *gin.Context) {
	var nfts []*types.AllNFT

	if err := c.ShouldBind(&nfts); err != nil {
		sendResponse(c, err, nil)
		return
	}

	transferNft, err := server.GetTransferNFT(nfts)
	if err != nil {
		logger.Error("get nft transfer path fail, err:", err)
		sendResponse(c, err, nil)
	}
	sendResponse(c, nil, transferNft)

}

type Response struct {
	Result    bool        `json:"result"`
	Code      string      `json:"code"`
	Message   string      `json:"message"`
	MessageZh string      `json:"message_zh"`
	Body      interface{} `json:"body"`
}

func sendResponse(c *gin.Context, err error, data interface{}) {
	code, message, messageZH := decodeErr(err)
	c.JSON(http.StatusOK, Response{
		Result:    err == nil,
		Code:      code,
		Message:   message,
		MessageZh: messageZH,
		Body:      data,
	})
}

func decodeErr(err error) (string, string, string) {
	if err == nil {
		return "0", "SUCCESS", "success"
	} else {
		return "-1", "FAIL" + err.Error(), "fail"
	}
}

var logger = log.HandlerLogger.WithField("module", "handler")
