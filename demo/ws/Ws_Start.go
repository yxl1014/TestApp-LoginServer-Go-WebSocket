package ws

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

func WebSocket() {
	go WebsocketManager.Start()
	go WebsocketManager.SendService()
	go WebsocketManager.SendService()
	go WebsocketManager.SendGroupService()
	go WebsocketManager.SendGroupService()
	go WebsocketManager.SendAllService()
	go WebsocketManager.SendAllService()
	go TestSendGroup()
	go TestSendAll()

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	wsGroup := router.Group("/ws")
	{
		wsGroup.GET("/:channel", WebsocketManager.WsClient)
	}

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server Start Error: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown Error:", err)
	}
	log.Println("Server Shutdown")
}
