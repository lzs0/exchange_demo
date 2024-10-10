package controllers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"news_web/global"
	"news_web/models"
	"time"
)

func CreateExchangeRate(ctx *gin.Context) {
	var exchangeRate models.ExchangeRate
	if err := ctx.ShouldBindJSON(&exchangeRate); err != nil {
		fmt.Println("err:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "not found the msg"})
		return
	}

	exchangeRate.Date = time.Now()

	err := global.Db.AutoMigrate(&exchangeRate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	err = global.Db.Create(&exchangeRate).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusCreated, exchangeRate)

}

func GetExchangeRate(ctx *gin.Context) {
	var exchangeRates []models.ExchangeRate

	if err := global.Db.Find(&exchangeRates).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}
		return
	}
	ctx.JSON(http.StatusOK, exchangeRates)
}
