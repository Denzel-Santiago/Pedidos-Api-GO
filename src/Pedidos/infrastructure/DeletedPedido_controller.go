package infrastructure

import (
	"net/http"
	"strconv"

	"Pedidos-Api/src/Pedidos/application"
	"github.com/gin-gonic/gin"
)

type DeletePedidoController struct {
	deletePedidoUseCase *application.DeletePedidoUseCase
}

func NewDeletePedidoController(deletePedidoUseCase *application.DeletePedidoUseCase) *DeletePedidoController {
	return &DeletePedidoController{
		deletePedidoUseCase: deletePedidoUseCase,
	}
}

func (ctrl *DeletePedidoController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	errDelete := ctrl.deletePedidoUseCase.Run(id)
	if errDelete != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo eliminar el pedido",
			"error":   errDelete.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Pedido eliminado exitosamente",
	})
}
