/*
* @Author: sottxiong
* @Date:   2019-08-02 22:26:05
* @Last Modified by:   sottxiong
* @Last Modified time: 2019-08-02 22:36:06
 */
package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type conf struct {
	server   string
	port     int
	usename  string
	password string
	db       string
}

func GetConf(string file) *conf {
	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("Open configuration error: %s\n", err)
		return
	}
	defer f.Close()
	configuration := &conf{}
	decoder := json.NewDecoder(f)
	err = decoder.Decode(configuration)
	if err != nil {
		panic(err)
	}
	return configuration
}
