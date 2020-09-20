package main

import (
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
}

func readConfig(fileName string) []byte {
	jsonFile, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	defer jsonFile.Close()

	return byteValue
}
