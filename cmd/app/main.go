package main

import (
	"github.com/shashaneRanasinghe/simpleMongoAPI/internal/config"
	"github.com/shashaneRanasinghe/simpleMongoAPI/pkg/server"
	"github.com/tryfix/log"
)

func main() {

	config.LoadConfigs()
	closeChannel := server.Serve()
	<-closeChannel

	log.Info("Server Stopped")
}
