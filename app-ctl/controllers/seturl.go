package controllers

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Outtech105k/ShortUrlServer/app-ctl/models"
	"github.com/Outtech105k/ShortUrlServer/app-ctl/utils"
	"github.com/gin-gonic/gin"
)

func SetUrlHandler(appCtx *utils.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var r models.SetUrlRequest

		if err := c.ShouldBindJSON(&r); err != nil {
			if err == io.EOF {
				c.JSON(http.StatusBadRequest, gin.H{"error": "request body JSON is empty."})
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
			return
		}

		// カスタムIDが指定されていない場合、4文字カスタムIDの生成（最大10回試行）
		var customId string
		customeIdIsExists := false
		for i := 0; i < 10; i++ {
			var err error
			customId, err = utils.MakeRandomStr(4)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error."})
				log.Printf("MakeRandomStr error: %v", err)
				return
			}

			// 生成されたカスタムIDがRedisに存在するか確認
			customeIdIsExists, err = appCtx.Redis.IsExists(customId)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error."})
				log.Printf("Redis exists error: %v", err)
				return
			}

			if !customeIdIsExists {
				break
			}
		}
		if customeIdIsExists {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error."})
			log.Printf("Custom ID generation failed after 10 attempts.")
			return
		}

		// RedisにURLを保存
		if err := appCtx.Redis.SetURLRecord(customId, r.BaseURL, nil); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error."})
			log.Printf("Redis set URL record error: %v", err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"base_url":  r.BaseURL,
			"short_url": fmt.Sprintf("https://rk2.uk/%s", customId),
		})
	}
}
