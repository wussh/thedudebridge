package config

import (
	"github.com/tkanos/gonfig"
	"log"
)

type Configuration struct{
	ChatID string
	BotToken string
	InfluxHost string
	InfluxPort string
	InfluxDB string
	InfluxMeasurement string
	ListenPort string
}

var configuration Configuration

func init(){
	err := gonfig.GetConf("config.json", &configuration)
	if err != nil {
		log.Panic(err)
	}
}

func GetConf() Configuration{
	return configuration
}