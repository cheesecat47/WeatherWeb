package app

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"database/sql" 
	_ "github.com/go-sql-driver/mysql"
	
	"routes/v1"
)

func initDB() {
	// http://golang.site/go/article/107-MySql-사용---쿼리
	const MYSQL_ROOT_PW string = os.Getenv("MYSQL_ROOT_PASSWORD")
	const MYSQL_DB string = os.Getenv("MYSQL_DATABASE")

}

func Start(port string) {
	// TODO: refactor
	// gin.SetMode(gin.ReleaseMode)
	gopath := os.Getenv("GOPATH")
	// https://gin-gonic.com/ko-kr/docs/examples/write-log/
	// https://github.com/gin-gonic/gin#how-to-write-log-file
	// https://jeonghwan-kim.github.io/dev/2019/01/14/go-time.html
	f, _ := os.Create("logs/" + time.Now().Format(time.RFC3339) + ".log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	router := gin.Default()

	// https://riptutorial.com/ko/go/example/29299/gin을-사용한-restfull-프로젝트-api#undefined
	router.Use('/v1', )
	

	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	// https://github.com/gin-gonic/gin#manually
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	// https://github.com/gin-gonic/gin#graceful-shutdown-or-restart
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