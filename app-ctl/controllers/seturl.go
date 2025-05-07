package controllers

import (
	"net/http"

	"github.com/Outtech105k/ShortUrlServer/app-ctl/models"
	"github.com/gin-gonic/gin"
)

func SetUrl(c *gin.Context) {
	var r models.SetUrlRequest

	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"base": r.BaseURL,
		"id":   r.CustomID,
		"url":  "https://rk2.uk/aaaaa",
	})
}
