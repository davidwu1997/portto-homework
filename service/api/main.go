package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"portto-homework/pkg/api"
	"portto-homework/pkg/db"
	"portto-homework/pkg/eth"
	"portto-homework/pkg/log"
)

var (
	rpcEndpoint = "https://data-seed-prebsc-2-s3.binance.org:8545"
)

func main() {
	pg, err := db.NewPostgres()
	if err != nil {
		panic(err)
	}
	ethclient := eth.New(rpcEndpoint, pg)
	router := gin.Default()
	router.Use(api.InjectContext)
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, "healthy")
	})

	root := router.Group("/")
	if err := api.MountEthRoutes(root, ethclient); err != nil {
		log.Global().Error("api.MountEthRoutes failed")
		panic(err)
	}

	if err := router.Run(); err != nil {
		log.Global().Error("router.Run failed")
	}
}
