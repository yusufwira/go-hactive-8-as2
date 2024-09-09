package main

import (
	"assigment2/connection"
	connectDb "assigment2/connection"
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type Server struct {
	HttpServer *http.Server
}

func run(ctx context.Context) error {
	gormDB, err := connectDb.NewConnection()
	if err != nil {
		os.Exit(1)
	}

	router := InitRouter(gormDB)
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", 8080),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s := &Server{
		HttpServer: server,
	}

	go func() {
		log.Info().Msg(fmt.Sprintf("success to listen and serve on port : %d", 8080))
		if err := s.HttpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("failed to listen and serve: %+v")
		}
	}()
	<-ctx.Done()
	return nil

}

func SetupRouter() *gin.Engine {
	// set the runtime mode
	var mode = gin.ReleaseMode
	gin.SetMode(mode)

	// create a new router instance
	router := gin.Default()

	// set up middlewares
	// router.Use(gin.LoggerWithFormatter(logger.HTTPLogger))
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowHeaders:    []string{"Origin", "Accept", "Content-Type", "Authorization", "User-Agent"},
		ExposeHeaders:   []string{"Content-Length"},
	}))
	router.Use(gin.Recovery())

	return router
}

func InitRouter(gormDB *connection.GormDB) *gin.Engine {
	router := SetupRouter()
	api := router.Group("/api/v1")
	InitRequestOrderRoute(api, gormDB).Routes()
	return router
}
