package api

import "github.com/gin-gonic/gin"

func errMsg(err error) interface{} {
	return gin.H{
		"err": err.Error(),
	}
}
