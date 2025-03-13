package infrastructure

import (
	"net/http"
	"strconv"

	"Pedidos-Api/src/Pedidos/application"
	"github.com/gin-gonic/gin"
)

type GetPedidoController struct {
	getPedidoUseCase *application.GetPedidoUseCase
}

func NewGetPedidoController(getPedidoUseCase *application.GetPedidoUseCase) *GetPedidoController {
	return &GetPedidoController{
		getPedidoUseCase: getPedidoUseCase,
	}
}

func (ctrl *GetPedidoController) Run(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	pedido, err := ctrl.getPedidoUseCase.Run(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pedido no encontrado"})
		return
	}

	c.JSON(http.StatusOK, pedido)
}
