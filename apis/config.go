package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type ConfigData struct {
	DBName     string
	DBIp       string
	DBPort     string
	Instance   string
	Domain     string
	ServerPort string
	dbTimeOut  uint32
}

func readConfig(fileName string) ConfigData {
	jsonFile, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	defer jsonFile.Close()
	var configdata ConfigData
	json.Unmarshal(byteValue, &configdata)
	return configdata
}
