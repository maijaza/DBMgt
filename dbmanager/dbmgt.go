package dbmanager

import (
	"encoding/json"
	"fmt"
	"os"
)

func loadConfig() DBConfig {
	var config DBConfig
	configFile, err := os.Open("./configuration/DB.json")
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}

func GetConfig() DBConfig {
	fmt.Println("Load Config")
	return loadConfig()
}

func Test() {
	fmt.Println("test")
}
