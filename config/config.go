package config

import (
	"fmt"
	"log"
)

var (
	HttpPort string
)

func init() {
	LoadServer()
}

func LoadServer() {
	HttpPort = ":3000"
	if HttpPort == "" {
		log.Fatal("$PORT must be set")
	} else {
		fmt.Println("Port:", HttpPort)
	}

}
