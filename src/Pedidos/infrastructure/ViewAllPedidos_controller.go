package infrastructure

import (
	"net/http"

	"Pedidos-Api/src/Pedidos/application"
	"github.com/gin-gonic/gin"
)

type ViewAllPedidosController struct {
	viewAllPedidosUseCase *application.ViewAllPedidosUseCase
}

func NewViewAllPedidosController(viewAllPedidosUseCase *application.ViewAllPedidosUseCase) *ViewAllPedidosController {
	return &ViewAllPedidosController{
		viewAllPedidosUseCase: viewAllPedidosUseCase,
	}
}

func (ctrl *ViewAllPedidosController) Run(c *gin.Context) {
	pedidos, err := ctrl.viewAllPedidosUseCase.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los pedidos"})
		return
	}

	c.JSON(http.StatusOK, pedidos)
}
