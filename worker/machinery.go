package worker

import (
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	cfg "github.com/s4kibs4mi/snapify/config"
	"github.com/s4kibs4mi/snapify/tasks"
)

var machineryServer *machinery.Server
var msErr error

func NewMachineryServer() error {
	machineryServer, msErr = machinery.NewServer(&config.Config{
		Broker:        cfg.RabbitMQ().Broker,
		DefaultQueue:  cfg.RabbitMQ().DefaultQueue,
		ResultBackend: cfg.RabbitMQ().ResultBackend,
		AMQP: &config.AMQPConfig{
			ExchangeType:  cfg.RabbitMQ().AMQP.ExchangeType,
			Exchange:      cfg.RabbitMQ().AMQP.Exchange,
			BindingKey:    cfg.RabbitMQ().AMQP.BindingKey,
			PrefetchCount: cfg.RabbitMQ().AMQP.PrefetchCount,
		},
		ResultsExpireIn: 3600,
	})
	if msErr != nil {
		return msErr
	}
	return nil
}

func MachineryServer() *machinery.Server {
	return machineryServer
}

func RegisterTasks() error {
	if err := machineryServer.RegisterTask(tasks.TakeScreenShotTaskName, tasks.TakeScreenShot); err != nil {
		return err
	}
	return nil
}
