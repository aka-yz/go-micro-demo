package main

import (
	"github.com/aka-yz/go-micro-core"
	_ "github.com/aka-yz/go-micro-core/providers"
	"go-micro-demo/internal/api"
	cfg "go-micro-demo/internal/config"
	"go-micro-demo/internal/models"
	"go-micro-demo/internal/router"
	"go-micro-demo/internal/service"
)

func main() {
	go_micro_core.Run(
		models.Provider,
		service.Provider,
		cfg.Provider,
		router.Provider,
		api.Provider,
	)
}
