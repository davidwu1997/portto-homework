package api

import (
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"

	"portto-homework/pkg/context"
	"portto-homework/pkg/eth"
)

func MountEthRoutes(group *gin.RouterGroup, ethClient eth.Eth) error {
	handler := &ethHandler{
		ethClient: ethClient,
	}
	blocks := group.Group("/blocks")
	blocks.GET("/", handler.getBlocks)
	blocks.GET("/:hash", handler.getBlock)

	transaction := group.Group("/transaction")
	transaction.GET("/:txHash", handler.getTransaction)

	return nil
}

type ethHandler struct {
	ethClient eth.Eth
}

func (h *ethHandler) getBlocks(c *gin.Context) {
	ctx := c.MustGet("ctx").(context.Context)

	limitStr := c.Query("limit")
	n, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, errMsg(err))
		return
	}

	blocks, err := h.ethClient.GetBlocks(ctx, uint64(n))
	if err != nil {
		c.JSON(http.StatusBadRequest, errMsg(err))
		return
	}

	// remove tx fields to fit resp spec
	for _, b := range blocks {
		b.Transactions = nil
	}

	c.JSON(http.StatusOK, gin.H{"blocks": blocks})
}

func (h *ethHandler) getBlock(c *gin.Context) {
	ctx := c.MustGet("ctx").(context.Context)

	hashHexStr := c.Param("hash")
	hash := common.HexToHash(hashHexStr)

	block, err := h.ethClient.GetBlock(ctx, hash)
	if err == eth.ErrNotFound {
		c.JSON(http.StatusNotFound, errMsg(err))
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, errMsg(err))
		return
	}

	c.JSON(http.StatusOK, block)
}

func (h *ethHandler) getTransaction(c *gin.Context) {
	ctx := c.MustGet("ctx").(context.Context)

	txHashHexStr := c.Param("txHash")
	txHash := common.HexToHash(txHashHexStr)

	tx, err := h.ethClient.GetTransaction(ctx, txHash)
	if err == eth.ErrNotFound {
		c.JSON(http.StatusNotFound, errMsg(err))
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, errMsg(err))
		return
	}

	c.JSON(http.StatusOK, tx)
}
