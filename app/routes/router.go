package routes

import (
	"net/http"

	"github.com/Outtech105k/ShortUrlServer/app/controllers"
	"github.com/Outtech105k/ShortUrlServer/app/utils"
	"github.com/gin-gonic/gin"
)

func SetupRouter(appCtx *utils.AppContext) *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})
	r.GET("/:shortUrl", controllers.GetUrlHandler(appCtx))
	r.POST("/set", controllers.SetUrlHandler(appCtx))

	return r
}
