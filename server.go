package main

import (
	"./config"
	"./network_method"
	"./service/collection"
	"./service/log"
	"encoding/json"
	"os"
)

const configFile = "config/config.json"

func main() {

	//log file
	f, err := os.OpenFile("logfile.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)

	file, _ := os.Open(configFile)
	defer file.Close()
	decoder := json.NewDecoder(file)
	serverConfig := config.ServerConfig{}
	err = decoder.Decode(&serverConfig)
	if err != nil {
		log.Fatalf("Error while decoding certificates : %v", err)
	}

	collection.MongoDBInit()

	if serverConfig.Server.Status == "active" {
		// fire the gRPC server in a goroutine
		go func() {
			err := network_method.StartGRPCServer(serverConfig.Server, serverConfig.TLS)
			if err != nil {
				log.Fatalf("Failed to start gRPC server: %s", err)
			}
		}()
	}

	if serverConfig.OptionalServer.Status == "active" {
		go func() {
			tls := serverConfig.TLS
			err := network_method.StartRESTServer(serverConfig.OptionalServer, serverConfig.Server, tls.Certificate)
			if err != nil {
				log.Fatalf("Failed to start REST server: %s", err)
			}
		}()
	}

	select {}
}
