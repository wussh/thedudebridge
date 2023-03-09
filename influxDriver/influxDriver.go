package influxDriver

import (
	"log"
	"time"
	"thedudebridge/config"
	_ "github.com/influxdata/influxdb1-client"
	client "github.com/influxdata/influxdb1-client/v2"
)

func SendInflux(
	service string,
	name string,
	address string,
	status int,
	desc string){
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://"+config.GetConf().InfluxHost+":"+config.GetConf().InfluxPort,
	})

	if err != nil {
		log.Println("Error creating InfluxDB Client: ", err.Error())
	}

	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database: config.GetConf().InfluxDB,
		Precision: "s",
	})

	tags := map[string]string{
		"address" : address,
	}

	fields := map[string]interface{} {
		"name" : name,
		"service" : service,
		"status" : status,
		"desc" : desc,
	}

	point, err := client.NewPoint(
		config.GetConf().InfluxMeasurement,
		tags,
		fields,
		time.Now(),
	)

	bp.AddPoint(point)

	c.Write(bp)

}