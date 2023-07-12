package main

import (
	"github.com/firdasafridi/stat-hybrid/internal/config"
	piihandler "github.com/firdasafridi/stat-hybrid/internal/handler/pii"
	piiuc "github.com/firdasafridi/stat-hybrid/internal/usecase/pii"
)

func app(cfg *config.Config) moduleHandler {

	piiUC := piiuc.New(&piiuc.PII{
	})

	piiHandler := piihandler.PIIHandler{
		PIIUC: piiUC,
	}

	return moduleHandler{
		PIIHandler: piiHandler,
	}
}
