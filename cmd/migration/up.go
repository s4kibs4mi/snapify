package migration

import (
	"github.com/s4kibs4mi/snapify/app"
	"github.com/s4kibs4mi/snapify/config"
	"github.com/s4kibs4mi/snapify/log"
	"github.com/s4kibs4mi/snapify/models"
	"github.com/spf13/cobra"
	"os"
)

var UpCmd = &cobra.Command{
	Use:   "up",
	Short: "up creates database tables",
	Run:   up,
	PreRun: func(cmd *cobra.Command, args []string) {
		if err := config.LoadConfig(upCmdConfigPath, upCmdConfigName); err != nil {
			log.Log().Errorln("Failed to read config : ", err)
			os.Exit(-1)
		}

		if err := app.ConnectSQLDB(); err != nil {
			log.Log().Errorln("Failed to connect to database : ", err)
			os.Exit(-1)
		}
	},
}

var upCmdConfigPath string
var upCmdConfigName string

func init() {
	UpCmd.Flags().StringVar(&upCmdConfigPath, "config_path", "", "configuration path")
	UpCmd.Flags().StringVar(&upCmdConfigName, "config_name", "", "configuration file name without extension")
}

func up(cmd *cobra.Command, args []string) {
	tx := app.DB().Begin()

	ss := models.Screenshot{}
	if err := tx.Model(&ss).AutoMigrate(&ss).Error; err != nil {
		log.Log().Errorln(err)
		tx.Rollback()
		return
	}

	if err := tx.Commit().Error; err != nil {
		log.Log().Errorln(err)
		return
	}

	log.Log().Infoln("Migration up completed.")
}
