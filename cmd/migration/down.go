package migration

import (
	"github.com/s4kibs4mi/snapify/app"
	"github.com/s4kibs4mi/snapify/config"
	"github.com/s4kibs4mi/snapify/log"
	"github.com/s4kibs4mi/snapify/models"
	"github.com/spf13/cobra"
	"os"
)

var DownCmd = &cobra.Command{
	Use:   "down",
	Short: "down drops database tables",
	Run:   down,
	PreRun: func(cmd *cobra.Command, args []string) {
		if err := config.LoadConfig(downCmdConfigPath, downCmdConfigName); err != nil {
			log.Log().Errorln("Failed to read config : ", err)
			os.Exit(-1)
		}

		if err := app.ConnectSQLDB(); err != nil {
			log.Log().Errorln("Failed to connect to database : ", err)
			os.Exit(-1)
		}
	},
}

var downCmdConfigPath string
var downCmdConfigName string

func init() {
	DownCmd.Flags().StringVar(&downCmdConfigPath, "config_path", "", "configuration path")
	DownCmd.Flags().StringVar(&downCmdConfigName, "config_name", "", "configuration file name without extension")
}

func down(cmd *cobra.Command, args []string) {
	tx := app.DB().Begin()

	ss := models.Screenshot{}
	if err := tx.Model(&ss).DropTableIfExists(&ss).Error; err != nil {
		log.Log().Errorln(err)
		tx.Rollback()
		return
	}

	if err := tx.Commit().Error; err != nil {
		log.Log().Errorln(err)
		return
	}

	log.Log().Infoln("Migration down completed.")
}
