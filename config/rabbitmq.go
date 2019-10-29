package config

import "github.com/spf13/viper"

type AMQP struct {
	Exchange      string
	ExchangeType  string
	BindingKey    string
	PrefetchCount int
}

type Worker struct {
	Name  string
	Count int
}

type RabbitMQCfg struct {
	Broker        string
	DefaultQueue  string
	ResultBackend string
	AMQP          AMQP
	Worker        Worker
}

var rabbitmq RabbitMQCfg

func LoadRabbitMQ() {
	mu.Lock()
	defer mu.Unlock()

	rabbitmq = RabbitMQCfg{
		Broker:        viper.GetString("rabbitmq.broker"),
		DefaultQueue:  viper.GetString("rabbitmq.default_queue"),
		ResultBackend: viper.GetString("rabbitmq.result_backend"),
		Worker: Worker{
			Name:  viper.GetString("rabbitmq.worker.name"),
			Count: viper.GetInt("rabbitmq.worker.count"),
		},
		AMQP: AMQP{
			Exchange:      viper.GetString("rabbitmq.amqp.exchange"),
			ExchangeType:  viper.GetString("rabbitmq.amqp.exchange_type"),
			BindingKey:    viper.GetString("rabbitmq.amqp.binding_key"),
			PrefetchCount: viper.GetInt("rabbitmq.amqp.prefetch_count"),
		},
	}
}

func RabbitMQ() RabbitMQCfg {
	return rabbitmq
}
