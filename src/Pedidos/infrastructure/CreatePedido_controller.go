package infrastructure

import (
	"net/http"

	"Pedidos-Api/src/Pedidos/application"
	"Pedidos-Api/src/Pedidos/domain/entities"
	"github.com/gin-gonic/gin"
)

type CreatePedidoController struct {
	createPedidoUseCase *application.CreatePedidoUseCase
}

func NewCreatePedidoController(createPedidoUseCase *application.CreatePedidoUseCase) *CreatePedidoController {
	return &CreatePedidoController{
		createPedidoUseCase: createPedidoUseCase,
	}
}

func (ctrl *CreatePedidoController) Run(c *gin.Context) {
	var pedido entities.Pedido

	if errJSON := c.ShouldBindJSON(&pedido); errJSON != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos del pedido inválidos",
			"error":   errJSON.Error(),
		})
		return
	}

	pedidoCreado, errAdd := ctrl.createPedidoUseCase.Run(&pedido)

	if errAdd != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error al agregar el pedido",
			"error":   errAdd.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "El pedido ha sido agregado",
		"pedido":  pedidoCreado,
	})
}
