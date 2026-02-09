package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/webomindapps-dev/coolaid-backend/config"
	cronjob "github.com/webomindapps-dev/coolaid-backend/cron_job"
	"github.com/webomindapps-dev/coolaid-backend/db"
	"github.com/webomindapps-dev/coolaid-backend/internal/api/graphql"
	ticketservice "github.com/webomindapps-dev/coolaid-backend/internal/domain/ticket"
	service "github.com/webomindapps-dev/coolaid-backend/internal/service/container"
	"github.com/webomindapps-dev/coolaid-backend/middleware"
	"github.com/webomindapps-dev/coolaid-backend/oplog"
	"github.com/webomindapps-dev/coolaid-backend/typesense"
)

var router *gin.Engine

func Init() {
	config.LoadConfigs()
}

// Bootstrap loads the required config
func Bootstrap() {
	Init()

	registerMiddleware()
}

func StartApplication() {

	gin.SetMode(gin.ReleaseMode)
	router = gin.Default()
	router.HandleMethodNotAllowed = true

	Bootstrap()

	server := &http.Server{
		Addr:    ":" + strconv.Itoa(config.App.Port),
		Handler: router,
	}

	go func() {
		appCtx := context.Background()

		err := db.Connect(appCtx)
		if err != nil {
			log.Fatalf("Error connecting db %s", err.Error())
		}

		typesense.Connect(appCtx)
		graphqlServer := connectGraphql()
		connectCronJobs()

		mapUrls(graphqlServer)

		fmt.Printf("server started and listening at port: %d \n", config.App.Port)

		//Service requests
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			oplog.Error(context.TODO(), fmt.Errorf("failed to start server: %w", err))
		}

	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	oplog.Info(context.TODO(), "shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		oplog.Error(context.TODO(), fmt.Errorf("failed to stop server : %w", err))
	} else {
		oplog.Info(context.TODO(), "server gracefully stopped")
	}
}

// Register Middleware
func registerMiddleware() {

	cfg := middleware.CORSConfig{
		AllowedOrigins: []string{
			"http://localhost:3050",
			"https://coolaid-frontend-virid.vercel.app/",
		},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	}

	router.Use(middleware.CORSMiddleware(cfg))
	router.Use(middleware.UserAgentMiddleware())
}

// Connect Graphql
func connectGraphql() *graphql.Server {
	services := service.NewContainer(db.DB, typesense.TS)

	return graphql.NewServer(services)
}

// Cron Job Handler
func connectCronJobs() {
	ticketSvc := ticketservice.NewService(db.DB)

	// 3. Register cron jobs
	scheduler := cronjob.RegisterAll(ticketSvc)

	// 4. Start cron scheduler
	scheduler.Start()
}
