package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

//Config d
type Config struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Secret   string `json:"secret"`
	Database struct {
		Dsn    string `json:"dsn"`
		Driver string `json:"driver"`
	} `json:"database"`
}

//NewConfig d
func NewConfig() Config {
	f, err := os.Open("./config/config.json")
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	c := Config{}
	json.Unmarshal(data, &c)
	return c
}
