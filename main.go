package main

import (
	"github.com/aka-yz/go-micro-core"
	"go-micro-demo/pkg/models"
	"go-micro-demo/pkg/service"
)

func main() {
	go_micro_core.Run(
		models.Provider,
		service.Provider,
	)
}
