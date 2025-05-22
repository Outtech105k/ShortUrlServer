package routes

import (
	"net/http"

	"github.com/Outtech105k/ShortUrlServer/app-ctl/controllers"
	"github.com/Outtech105k/ShortUrlServer/app-ctl/utils"
	"github.com/gin-gonic/gin"
)

func SetupRouter(appCtx *utils.AppContext) *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})

	r.POST("/set", controllers.SetUrlHandler(appCtx))

	return r
}
