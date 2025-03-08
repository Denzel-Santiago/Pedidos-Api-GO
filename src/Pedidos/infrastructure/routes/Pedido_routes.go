package routes

import (
	"Pedidos-Api/src/pedidos/application"
	"github.com/gin-gonic/gin"
)

func SetupPedidosRoutes(r *gin.Engine, pedidosHandler *application.PedidosHandler) {
	r.POST("/pedidos", pedidosHandler.CreatePedido)
	r.GET("/pedidos/:id", pedidosHandler.GetPedido)
}
