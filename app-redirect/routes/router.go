package routes

import (
	"net/http"

	"github.com/Outtech105k/ShortUrlServer/app-redirect/utils"
	"github.com/gin-gonic/gin"
)

func SetupRouter(appCtx *utils.AppContext) *gin.Engine {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello, Redirect.")
	})

	return r
}
