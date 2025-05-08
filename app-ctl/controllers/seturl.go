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

		if err := setUrlHandlerVaridate(&r); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// UseUppercase, UseLowercase, UseNumbers, IDLengthのデフォルト値を設定
		if r.UseUppercase == nil {
			r.UseUppercase = utils.BoolPtr(false)
		}
		if r.UseLowercase == nil {
			r.UseLowercase = utils.BoolPtr(true)
		}
		if r.UseNumbers == nil {
			r.UseNumbers = utils.BoolPtr(true)
		}
		if r.IDLength == nil {
			r.IDLength = utils.Uint32Ptr(6)
		}

		var customId string
		if r.CustomID == nil {
			// カスタムIDが指定されていない場合、4文字カスタムIDの生成（最大10回試行）
			customIdIsExists := false
			for i := 0; i < 10; i++ {
				var err error
				customId, err = utils.MakeRandomStr(
					*r.IDLength,
					*r.UseUppercase,
					*r.UseLowercase,
					*r.UseNumbers,
				)
				if err != nil {
					if err == utils.ErrNoCharacterSet {
						c.JSON(http.StatusBadRequest, gin.H{"error": "no character types available."})
						return
					}
					c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error."})
					log.Printf("MakeRandomStr error: %v", err)
					return
				}

				// 生成されたカスタムIDがRedisに存在するか確認
				customIdIsExists, err = appCtx.Redis.IsExists(customId)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error."})
					log.Printf("Redis exists error: %v", err)
					return
				}

				if !customIdIsExists {
					break
				}
			}
			if customIdIsExists {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error."})
				log.Printf("Custom ID generation failed after 10 attempts.")
				return
			}
		} else {
			customId = *r.CustomID
			// カスタムIDが指定されている場合、Redisに存在するか確認
			customIdIsExists, err := appCtx.Redis.IsExists(customId)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error."})
				log.Printf("Redis exists error: %v", err)
				return
			}

			if customIdIsExists {
				c.JSON(http.StatusConflict, gin.H{"error": "custom_id already used."})
				return
			}
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

func setUrlHandlerVaridate(r *models.SetUrlRequest) error {
	if r.CustomID != nil {
		if r.UseUppercase != nil || r.UseLowercase != nil || r.UseNumbers != nil || r.IDLength != nil {
			return fmt.Errorf("custom_id is specified, but use_uppercase, use_lowercase, use_numbers, id_length are also specified")
		}
	}

	return nil
}
