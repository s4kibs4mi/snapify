package worker

import (
	"github.com/RichardKnop/machinery/v1"
	"github.com/s4kibs4mi/snapify/config"
	"github.com/s4kibs4mi/snapify/log"
	"os"
)

var worker *machinery.Worker
var err error

func NewWorker() {
	cnf := config.RabbitMQ().Worker
	worker = MachineryServer().NewWorker(cnf.Name, cnf.Count)
	err = worker.Launch()
	if err != nil {
		log.Log().Errorln("Couldn't launch worker", err)
		os.Exit(-1)
	}
}
