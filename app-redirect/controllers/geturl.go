package controllers

import (
	"log"
	"net/http"

	"github.com/Outtech105k/ShortUrlServer/app-redirect/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func GetUrlHandler(appCtx *utils.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		shortUrl := c.Param("shortUrl")

		// Redisに問い合わせてURLを取得
		baseUrl, err := appCtx.Redis.GetBaseUrl(shortUrl)
		if err != nil {
			// 保存されていない(nil)場合は404を返す
			if err == redis.Nil {
				c.JSON(http.StatusNotFound, gin.H{
					"error": "URL not found",
				})
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to retrieve base URL",
			})
			log.Printf("Failed to retrieve base URL: %v", err)
			return
		}

		c.Redirect(http.StatusFound, baseUrl)
	}
}
