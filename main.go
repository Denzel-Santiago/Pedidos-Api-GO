package main

import (
	"Pedidos-Api/src/pedidos/infrastructure"
	"Pedidos-Api/src/pedidos/infrastructure/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	deps, err := infrastructure.SetupDependencies()
	if err != nil {
		log.Fatalf("Failed to setup dependencies: %v", err)
	}

	r := gin.Default()
	routes.SetupPedidosRoutes(r, deps.PedidosHandler)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
