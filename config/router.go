package config

import (
	"net/http"

	gr "github.com/Tiratom/gin-study/grpc"
	"github.com/Tiratom/gin-study/presentation"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	engine := gin.Default()

	// ダミーメソッド（叩けるか試すために作った）
	engine.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "hello"})
	})

	// ダミーメソッド（叩けるか試すために作った）
	engine.GET("/cat", func(ctx *gin.Context) {
		cs := &presentation.CatServiceServer{}
		msg := &gr.GetMyCatMessage{TargetCat: "tama"}
		rsp, err := cs.GetMyCat(ctx, msg)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "ERROR"})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": rsp})
		}
	})

	return engine
}
