package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"websrvfileshow/util"
	"websrvfileshow/web"
)

type Config struct {
	Port    string `json:"port"`
	Cert    string `json:"cert"`
	CertKey string `json:"certKey"`
}

func main() {

	data, err := os.ReadFile(util.GetExecDir() + "config.json")
	if err != nil {
		fmt.Print(err.Error())
		log.Fatal(err)
	}

	var config Config

	err = json.Unmarshal(data, &config)
	if err != nil {
		fmt.Print(err.Error())
		log.Fatal(err)
	}

	web.StartServer(config.Port, config.Cert, config.CertKey)
}
