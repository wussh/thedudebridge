package	telegramDriver

import (
	"net/http"
	"log"
	"io/ioutil"
	"strconv"
	"thedudebridge/config"
)

func SendTelegram(
	service string,
	name string,
	address string,
	status int,
	desc string){
	telegram_request := "https://api.telegram.org/bot"+config.GetConf().BotToken+"/sendMessage?chat_id=-"+config.GetConf().ChatID+"&text=the dude notification Service "+service+" on "+name+" "+address+" is now "+strconv.Itoa(status)+" "+desc
	request, err := http.Get(telegram_request)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(request.Body)
	log.Printf(string(body))
}