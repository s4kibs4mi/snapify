package main

import (
	"github.com/s4kibs4mi/snapify/cmd"
	"github.com/s4kibs4mi/snapify/log"
)

func main() {
	log.SetupLog()
	if err := cmd.Execute(); err != nil {
		log.Log().Errorln("Failed to execute command : ", err)
	}
}
