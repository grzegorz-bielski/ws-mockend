package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const configPath = "./config.json"

type UnknownJSON = map[string]*json.RawMessage

type Config interface {
	stringify() string
}

type WSConfig struct {
	Route    string      `json:"route"`
	Interval int         `json:"interval"`
	Msg      UnknownJSON `json:"data"`
}

func (wsd *WSConfig) stringify() string {
	bytes, err := json.Marshal(wsd.Msg)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return string(bytes)
}

func readConfig(path string) []byte {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return raw
}

func getConfig() []WSConfig {
	var wsConfig []WSConfig

	err := json.Unmarshal(readConfig(configPath), &wsConfig)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return wsConfig
}
