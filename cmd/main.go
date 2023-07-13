package main

import (
	"github.com/firdasafridi/stat-hybrid/internal/config"
	"github.com/firdasafridi/stat-hybrid/lib/common/log"
)

const (
	appName = "pii-lot"
)

func main() {
	log.Infoln("Starting new service...")

	cfg, err := config.New("piilot")
	if err != nil {
		log.Fatalln("Can't get config file", err)
	}

	mHandler := app(cfg)

	httpServer := newRoutes(mHandler)

	log.Errorln(startServer(cfg, httpServer))
}
