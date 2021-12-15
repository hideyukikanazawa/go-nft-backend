package main

import (
	"log"
	"github.com/hideyukikanazawa/go-nft-backend/util"

)


func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	log.Info(config.ALCHEMY_API_KEY)
}
