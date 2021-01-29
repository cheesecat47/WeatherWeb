package server

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cheesecat47/webpractice/api/model"
	"github.com/gin-gonic/gin"
)

// Start 서버 시작
func Start(port string) {
	// TODO: refactor
	// gin.SetMode(gin.ReleaseMode)
	gopath := os.Getenv("GOPATH")
	log.Printf("gopath: %s", gopath)
	f, _ := os.Create("logs/" + time.Now().Format(time.RFC3339) + ".log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	router := gin.Default()

	go model.InitDB()

	// router.Use('/v1', )
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, this is go api server.",
		})
	})

	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
