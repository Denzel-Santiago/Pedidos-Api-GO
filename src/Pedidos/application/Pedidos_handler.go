package application

import (
	"Pedidos-Api/src/pedidos/domain/entities"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PedidosHandler struct {
	createPedidoUseCase *CreatePedidosUseCase
	getPedidoUseCase    *GetPedidosUseCase
}

func NewPedidosHandler(createPedidoUseCase *CreatePedidosUseCase, getPedidoUseCase *GetPedidosUseCase) *PedidosHandler {
	return &PedidosHandler{
		createPedidoUseCase: createPedidoUseCase,
		getPedidoUseCase:    getPedidoUseCase,
	}
}

func (h *PedidosHandler) CreatePedido(c *gin.Context) {
	var pedido entities.Pedido
	if err := c.ShouldBindJSON(&pedido); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.createPedidoUseCase.Execute(pedido); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Pedido created successfully"})
}

func (h *PedidosHandler) GetPedido(c *gin.Context) {
	pedidoID := c.Param("id")
	var id int
	if _, err := fmt.Sscan(pedidoID, &id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pedido ID"})
		return
	}

	pedido, err := h.getPedidoUseCase.Execute(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, pedido)
}
