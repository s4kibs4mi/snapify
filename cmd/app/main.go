package main

import (
	"context"
	gSQL "database/sql"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/s4kibs4mi/snapify/api"
	"github.com/s4kibs4mi/snapify/config"
	"github.com/s4kibs4mi/snapify/dao"
	"github.com/s4kibs4mi/snapify/ent"
	"github.com/s4kibs4mi/snapify/log"
	"github.com/s4kibs4mi/snapify/services"
	"os"
	"os/signal"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	logger := log.New()
	logger.Info("Logger initialized")

	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Fatal(err)
	}
	logger.Info("Config loaded")

	dbConUrl := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s", cfg.DBCfg.Username,
		cfg.DBCfg.Password, cfg.DBCfg.Host, cfg.DBCfg.Port, cfg.DBCfg.Name)
	logger.Info("dbConUrl", dbConUrl)

	sqlClient, err := gSQL.Open("pgx", dbConUrl)
	if err != nil {
		logger.Fatal(err)
	}

	driver := sql.OpenDB(dialect.Postgres, sqlClient)
	entDriver := ent.Driver(driver)
	client := ent.NewClient(entDriver)
	if err != nil {
		logger.Fatal("failed opening connection to postgresql", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		logger.Fatal("failed creating schema resources", err)
	}

	screenshotDao := dao.NewScreenshotDao(client, logger)
	tokenDao := dao.NewTokenDao(client, logger)

	storageService, err := services.NewMinioService(cfg, logger)
	if err != nil {
		logger.Fatal("failed to initialize storage service", err)
	}

	if err := storageService.Init(); err != nil {
		logger.Fatal("failed to execute init task", err)
	}

	queueService, err := services.NewQueueService(cfg, logger)
	if err != nil {
		logger.Fatal("failed to initialize queuing service", err)
	}

	handlers := api.NewHandlers(screenshotDao, tokenDao, queueService, storageService, logger)
	apiServer := api.NewServer(cfg)

	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)

	go func() {
		if err := apiServer.Start(handlers); err != nil {
			logger.Error(err)
			stop <- os.Interrupt
		}
	}()
	logger.Info("Api server running...")

	<-stop

	apiServer.Stop()
	logger.Info("Api server stopped")
}