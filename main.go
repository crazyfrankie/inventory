package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"inventory/config"
	"inventory/route"
)

func main() {
	// 加载配置文件
	dbConfig, err := config.LoadConfig("config/config.ini")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 初始化数据库
	_, err = config.InitDB(dbConfig)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// 自动迁移数据库结构
	config.InitModel()

	// 设置gin框架
	router := route.SetupRouter()

	// 使用优雅启停方式运行服务
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// 启动服务器
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// 等待中断信号以优雅关闭服务器
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 关闭服务器，等待所有请求处理完毕后再退出
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited gracefully")
}
