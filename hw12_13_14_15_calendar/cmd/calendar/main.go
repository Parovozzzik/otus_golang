package main

import (
	"context"
	"database/sql"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	app "github.com/Parovozzzik/otus_golang/hw12_13_14_15_calendar/internal/app"
	config "github.com/Parovozzzik/otus_golang/hw12_13_14_15_calendar/internal/config"
	logger "github.com/Parovozzzik/otus_golang/hw12_13_14_15_calendar/internal/logger"
	internalhttp "github.com/Parovozzzik/otus_golang/hw12_13_14_15_calendar/server/http"
	client "github.com/Parovozzzik/otus_golang/hw12_13_14_15_calendar/storage/sql"
)

var configFile string

var database *sql.DB

func init() {
	flag.StringVar(&configFile, "config")
}

func main() {
	flag.Parse()

	logger.Init()
	logger := logger.GetLogger()
	logger.Info("logger initialized")

	cnf := config.GetConfig()
	//logg := logger.New(config.Logger.Level)

	db, err := client.NewClient(cnf.MySql.Host, cnf.MySql.Port, cnf.MySql.Username, cnf.MySql.Password, cnf.MySql.Database)
	if err != nil {
		logger.Info(err)
	}

	storage = db
	logger.Info("database initialized")
	defer database.Close()

	calendar := app.New(logger, storage)

	server := internalhttp.NewServer(logger, calendar)

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		if err := server.Stop(ctx); err != nil {
			logger.Error("failed to stop http server: " + err.Error())
		}
	}()

	logger.Info("calendar is running...")

	if err := server.Start(ctx); err != nil {
		logger.Error("failed to start http server: " + err.Error())
		cancel()
		os.Exit(1) //nolint:gocritic
	}
}
