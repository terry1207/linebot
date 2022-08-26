package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

var (
	HttpPort string
)

var Config struct {
	Port uint `default:"5000" env:"PORT"`
	DB   struct {
		Name     string `env:"DBName" default:"example"`
		Adapter  string `env:"DBAdapter" default:"mysql"`
		Host     string `env:"DBHost" default:"localhost"`
		Port     string `env:"DBPort" default:"3306"`
		User     string `env:"DBUser"`
		Password string `env:"DBPassword"`
	}
}

func init() {
	y, err := ioutil.ReadFile("./internal/config/database.yml")
	if err != nil {
		panic(err)
	}
	yaml.Unmarshal(y, &Config)
	LoadServer()
}

func LoadServer() {
	HttpPort = os.Getenv("PORT")
	//HttpPort = "5000"
	if HttpPort == "" {
		log.Fatal("$PORT must be set")
	} else {
		fmt.Println("Port:", HttpPort)
	}

}
