package infrastructure

import (
	"Pedidos-Api/src/pedidos/application"
	"Pedidos-Api/src/pedidos/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PedidosController struct {
	createPedidoUseCase *application.CreatePedidosUseCase
}

func NewPedidosController(createPedidoUseCase *application.CreatePedidosUseCase) *PedidosController {
	return &PedidosController{createPedidoUseCase: createPedidoUseCase}
}

func (c *PedidosController) CreatePedido(ctx *gin.Context) {
	var pedido entities.Pedido
	if err := ctx.ShouldBindJSON(&pedido); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.createPedidoUseCase.Execute(pedido); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Pedido created successfully"})
}
