package main

import (
	"log"
	"net/http"
	"github.com/labstack/echo/v4"
	"thedudebridge/telegramDriver"
	"thedudebridge/influxDriver"
	"thedudebridge/config"
)

func responseOK(c echo.Context) error {
	c.Response().Header().Set("X-Bms-Response","OK")
	return c.String(http.StatusOK, "You're set!")
}

func getTheDudeParams(w http.ResponseWriter, r *http.Request){
	service := r.URL.Query()["service"][0]
	name := r.URL.Query()["name"][0]
	address := r.URL.Query()["address"][0]
	status_s := r.URL.Query()["status"][0]
	desc := r.URL.Query()["desc"][0]
	var status int
	if status_s == "up" {
		status = 1
	}else{
		status = 0
	}
	log.Println(service,name,address,status,desc)
	telegramDriver.SendTelegram(service,name,address,status,desc)
	influxDriver.SendInflux(service,name,address,status,desc)
}

func handleRequest(){
	http.HandleFunc("/", getTheDudeParams)
	log.Fatal(http.ListenAndServe(":"+config.GetConf().ListenPort, nil))
}

func main(){
	e := echo.New()
	e.GET("/", responseOK)
	e.Start(":"+config.GetConf().ListenPort)
}