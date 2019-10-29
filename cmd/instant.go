package cmd

import (
	"github.com/s4kibs4mi/snapify/config"
	"github.com/s4kibs4mi/snapify/core"
	"github.com/s4kibs4mi/snapify/log"
	"github.com/spf13/cobra"
	"os"
)

var instantCmd = &cobra.Command{
	Use:   "instant",
	Short: "instant takes screen shot of given url via cli",
	Run:   onInstance,
	PreRun: func(cmd *cobra.Command, args []string) {
		if err := config.LoadConfig(configPath, configName); err != nil {
			log.Log().Errorln("Failed to read config : ", err)
			os.Exit(-1)
		}
	},
}

var targetUrl string
var targetDirectory string

func init() {
	instantCmd.Flags().StringVar(&targetUrl, "url", "", "target web page url (ex: https://www.example.com)")
	instantCmd.Flags().StringVar(&targetDirectory, "out", "", "directory to save screenshots (ex: /root/ss)")
	instantCmd.Flags().StringVar(&configPath, "config_path", "", "configuration path")
	instantCmd.Flags().StringVar(&configName, "config_name", "", "configuration file name without extension")
}

func onInstance(cmd *cobra.Command, args []string) {
	if err := core.TakeScreenShotAndSave(targetUrl, targetDirectory); err != nil {
		log.Log().Errorln("Failed to take screen shot : ", err)
		return
	}
	log.Log().Infof("Screen shot taken of %s.", targetUrl)
	log.Log().Infof("Saved to directory : %s", targetDirectory)
}
