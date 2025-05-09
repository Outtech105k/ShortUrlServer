package routes

import (
	"net/http"

	"github.com/Outtech105k/ShortUrlServer/app-ctl/controllers"
	"github.com/Outtech105k/ShortUrlServer/app-ctl/utils"
	"github.com/gin-gonic/gin"
)

func SetupRouter(appCtx *utils.AppContext) *gin.Engine {
	r := gin.Default()
	api := r.Group("/api/shorturl")
	{
		api.GET("/", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "Hello. This is short URL service!")
		})

		api.POST("/set", controllers.SetUrlHandler(appCtx))
	}

	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*")

	gui := r.Group("/service/shorturl")
	{
		gui.GET("/", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "index.html", nil)
		})
	}

	return r
}
