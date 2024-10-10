package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"net/http"
	"news_web/global"
)

func LikeArticle(ctx *gin.Context) {
	articleId := ctx.Param("id")

	likeKey := "article:" + articleId

	if err := global.RedisDb.Incr(context.Background(), likeKey).Err(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "Successfully liked the article"})
}

func GetArticleLikes(ctx *gin.Context) {
	articleId := ctx.Param("id")

	likeKey := "article:" + articleId

	likes, err := global.RedisDb.Get(context.Background(), likeKey).Result()

	if err == redis.Nil {
		likes = "0"
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"likes": likes})
}
