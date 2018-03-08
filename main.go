package main

import (
	"log"
)

func main() {
	config, err := NewConfig("config.yml")
	if err != nil {
		log.Fatalf("config error - %v\n", err)
	}
	go StartNotModifiedProxy(config)
	log.Printf("Last modified proxy server at :%v, :%v(force check)\n", config.NotModifiedPort, config.CheckModifiedPort)
	StartCheckModifiedProxy(config)
}
