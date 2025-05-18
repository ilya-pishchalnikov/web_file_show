package main

import (
	"encoding/json"
	"filesender/util"
	"fmt"
	"log"
	"os"
	"time"
)

type Config struct {
	FileName    string `json:"fileName"`
	PostMethod  string `json:"postMethod"`
	Periodicity int    `json:"periodicity"`
	UserName    string `json:"userName"`
	Password    string `json:"password"`
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

	interval := time.Duration(config.Periodicity) * time.Second

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// Этот код будет выполняться каждые N секунд
			fmt.Println(time.Now())

			text, err := os.ReadFile(config.FileName)
			if err != nil {
				fmt.Println(err.Error())
			}

			err = util.PostFile(text, config.PostMethod, config.UserName, config.Password)

			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}

}
