package cmd

import (
	"github.com/s4kibs4mi/snapify/core"
	"github.com/s4kibs4mi/snapify/log"
	"github.com/spf13/cobra"
)

var instanceCmd = &cobra.Command{
	Use:   "instance",
	Short: "instance takes screen shot of given url via cli",
	Run:   onInstance,
}

var targetUrl string
var targetDirectory string

func init() {
	instanceCmd.Flags().StringVar(&targetUrl, "url", "", "target web page url (ex: https://www.example.com)")
	instanceCmd.Flags().StringVar(&targetDirectory, "out", "", "directory to save screenshots (ex: /root/ss)")
}

func onInstance(cmd *cobra.Command, args []string) {
	if err := core.TakeScreenShot(targetUrl, targetDirectory); err != nil {
		log.Log().Errorln("Failed to take screen shot : ", err)
		return
	}
	log.Log().Infof("Screen shot taken of %s.", targetUrl)
	log.Log().Infof("Saved to directory : %s", targetDirectory)
}
