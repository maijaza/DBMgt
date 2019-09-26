package dbmanager

import (
	"encoding/json"
	"fmt"
	"os"
)

func loadConfig(s string) DBConfig {
	var config DBConfig
	configFile, err := os.Open("./configuration/" + s + ".json")
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}

func GetConfig(str string) DBConfig {
	fmt.Println("Load Config")
	var confile string
	if str == "" {
		confile = "DB"
	} else {
		confile = str
	}
	return loadConfig(confile)
}

func Test() {
	fmt.Println("test")
}
