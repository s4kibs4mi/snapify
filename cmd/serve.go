package cmd

import (
	"context"
	"fmt"
	"github.com/s4kibs4mi/snapify/api"
	"github.com/s4kibs4mi/snapify/app"
	"github.com/s4kibs4mi/snapify/config"
	"github.com/s4kibs4mi/snapify/log"
	"github.com/s4kibs4mi/snapify/worker"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var serveCmd = &cobra.Command{
	Use: "serve",
	Run: serve,
	PreRun: func(cmd *cobra.Command, args []string) {
		if err := config.LoadConfig(configPath); err != nil {
			log.Log().Errorln("Failed to read config : ", err)
			os.Exit(-1)
		}

		if err := app.ConnectSQLDB(); err != nil {
			log.Log().Errorln("Failed to connect to database : ", err)
			os.Exit(-1)
		}
		if err := app.ConnectMinio(); err != nil {
			log.Log().Errorln("Failed to connect to minio : ", err)
			os.Exit(-1)
		}
		if err := worker.NewMachineryServer(); err != nil {
			log.Log().Errorln("Failed to connect to rabbitmq : ", err)
			os.Exit(-1)
		}
		go worker.NewWorker()
		if err := worker.RegisterTasks(); err != nil {
			log.Log().Errorln("Failed to register tasks : ", err)
			os.Exit(-1)
		}
	},
}

var configPath string

func init() {
	serveCmd.Flags().StringVar(&configPath, "config_path", "", "configuration path")
}

func serve(cmd *cobra.Command, args []string) {
	addr := fmt.Sprintf("%s:%d", config.App().Base, config.App().Port)

	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)

	hServer := http.Server{
		Addr:    addr,
		Handler: api.Router(),
	}

	go func() {
		log.Log().Infoln("Http server has been started on", addr)
		if err := hServer.ListenAndServe(); err != nil {
			log.Log().Errorln("Failed to start http server on :", err)
			os.Exit(-1)
		}
	}()

	<-stop

	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	if err := hServer.Shutdown(ctx); err != nil {
		log.Log().Infoln("Http server couldn't shutdown gracefully")
	}
	log.Log().Infoln("Http server has been shutdown gracefully")
}
