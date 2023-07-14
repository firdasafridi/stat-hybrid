package main

import (
	"github.com/firdasafridi/stat-hybrid/internal/config"
	piihandler "github.com/firdasafridi/stat-hybrid/internal/handler/pii"
	piiuc "github.com/firdasafridi/stat-hybrid/internal/usecase/pii"
	"github.com/firdasafridi/stat-hybrid/lib/hybridencryption"
)

func app(cfg *config.Config) moduleHandler {

	hybridEncryptionRepo, err := hybridencryption.NewHybridEncryption(hybridencryption.RSAOption{
		PubKey:     cfg.Server.HybridEncryption.PublicKey,
		PrivateKey: cfg.Server.HybridEncryption.PrivateKey,
	})

	if err != nil {
		panic(err)
	}

	piiUC := piiuc.New(&piiuc.PII{
		HBELIB: hybridEncryptionRepo,
	})

	piiHandler := piihandler.PIIHandler{
		PIIUC: piiUC,
	}

	return moduleHandler{
		PIIHandler: piiHandler,
	}
}
