package main

import (
	"context"
	"log"
	"net/http"
	"news_web/config"
	"news_web/router"
	"os"
	"os/signal"
	"time"
)

func main() {
	config.InitConfig()
	r := router.SetRouter()

	srv := &http.Server{
		Addr:    config.AppConfig.App.Port,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen:%v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("shutdown server....")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
