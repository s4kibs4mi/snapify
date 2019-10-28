package cmd

import (
	"github.com/s4kibs4mi/snapify/core"
	"github.com/s4kibs4mi/snapify/log"
	"github.com/spf13/cobra"
)

var instantCmd = &cobra.Command{
	Use:   "instant",
	Short: "instant takes screen shot of given url via cli",
	Run:   onInstance,
}

var targetUrl string
var targetDirectory string

func init() {
	instantCmd.Flags().StringVar(&targetUrl, "url", "", "target web page url (ex: https://www.example.com)")
	instantCmd.Flags().StringVar(&targetDirectory, "out", "", "directory to save screenshots (ex: /root/ss)")
}

func onInstance(cmd *cobra.Command, args []string) {
	if err := core.TakeScreenShotAndSave(targetUrl, targetDirectory); err != nil {
		log.Log().Errorln("Failed to take screen shot : ", err)
		return
	}
	log.Log().Infof("Screen shot taken of %s.", targetUrl)
	log.Log().Infof("Saved to directory : %s", targetDirectory)
}
