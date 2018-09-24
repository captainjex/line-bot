package main

import (
	"log"
	"net/http"
	"github.com/ciazhar/line-bot/controller"
)

func main() {
	http.HandleFunc("/", controller.Base)
	http.HandleFunc("/callback", controller.Callback)
	
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}