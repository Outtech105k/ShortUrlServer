package routes

import (
	"net/http"

	"github.com/Outtech105k/ShortUrlServer/app-redirect/controllers"
	"github.com/Outtech105k/ShortUrlServer/app-redirect/utils"
	"github.com/gin-gonic/gin"
)

func SetupRouter(appCtx *utils.AppContext) *gin.Engine {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusFound, "https://outtech105.com/service/shorturl")
	})

	r.LoadHTMLGlob("templates/*")

	r.GET("/:shortUrl", controllers.GetUrlHandler(appCtx))

	return r
}
