package infrastructure

import (
	"net/http"
	"strconv"

	"Pedidos-Api/src/Pedidos/application"
	"github.com/gin-gonic/gin"
)

type ViewPedidoController struct {
	viewPedidoUseCase *application.ViewPedidoUseCase
}

func NewViewPedidoController(viewPedidoUseCase *application.ViewPedidoUseCase) *ViewPedidoController {
	return &ViewPedidoController{
		viewPedidoUseCase: viewPedidoUseCase,
	}
}

func (ctrl *ViewPedidoController) Run(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	pedido, err := ctrl.viewPedidoUseCase.Run(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pedido no encontrado"})
		return
	}

	c.JSON(http.StatusOK, pedido)
}
