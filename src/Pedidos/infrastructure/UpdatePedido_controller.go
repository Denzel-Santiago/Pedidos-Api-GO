package infrastructure

import (
	"net/http"
	"strconv"

	"Pedidos-Api/src/Pedidos/application"
	"Pedidos-Api/src/Pedidos/domain/entities"
	"github.com/gin-gonic/gin"
)

type UpdatePedidoController struct {
	updatePedidoUseCase *application.UpdatePedidoUseCase
}

func NewUpdatePedidoController(updatePedidoUseCase *application.UpdatePedidoUseCase) *UpdatePedidoController {
	return &UpdatePedidoController{
		updatePedidoUseCase: updatePedidoUseCase,
	}
}

func (ctrl *UpdatePedidoController) Run(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var pedido entities.Pedido
	if err := c.ShouldBindJSON(&pedido); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = ctrl.updatePedidoUseCase.Run(id, pedido)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar el pedido"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pedido actualizado exitosamente"})
}
