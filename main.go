package main

import (
	"log"
)

func main() {
	config, err := NewConfig("config.yml")
	if err != nil {
		log.Fatalf("config error - %v\n", err)
	}
	log.Printf("Last modified proxy server at :%v\n", config.Port)
	log.Fatalln(StartLastModifiedProxy(config))
}
