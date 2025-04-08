// pedidos-api-go/src/Pedidos/infrastructure/routes/Pedido_routes.go
package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"Pedidos-Api/src/Pedidos/infrastructure"

	"github.com/gin-gonic/gin"
)

type Router struct {
	engine *gin.Engine
}

func NewRouter(engine *gin.Engine) *Router {
	return &Router{
		engine: engine,
	}
}

// Función para recibir pedidos y actualizar boletos
func logPedidoHandler(c *gin.Context) {
	var pedido map[string]interface{}

	if err := c.ShouldBindJSON(&pedido); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Imprimir JSON recibido
	prettyJSON, _ := json.MarshalIndent(pedido, "", "  ")
	fmt.Println("📩 Pedido recibido desde consumer.go:")
	fmt.Println(string(prettyJSON))
	fmt.Println("----------------------------")

	// ID del evento desde el JSON recibido
	eventID, ok := pedido["id"].(float64)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de evento inválido"})
		return
	}

	// función para actualizar boletos
	err := actualizarBoletos(int(eventID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// actualización exitosa
	c.JSON(http.StatusOK, gin.H{
		"message": "Pedido recibido y boletos actualizados correctamente",
		"data":    pedido,
	})
}

// obtener el ID del evento basado en la ubicación
func obtenerEventoPorUbicacion(location string) (int, error) {
	url := fmt.Sprintf("http://localhost:8000/events/location/%s", location)

	resp, err := http.Get(url)
	if err != nil {
		return 0, fmt.Errorf("error en la petición: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return 0, fmt.Errorf("evento no encontrado")
	} else if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("error en API 1, código: %d", resp.StatusCode)
	}

	var event struct {
		ID int `json:"id"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&event); err != nil {
		return 0, fmt.Errorf("error al decodificar respuesta")
	}

	return event.ID, nil
}

// petición PUT a la API 1 y reducir los boletos
func actualizarBoletos(eventID int) error {
	url := fmt.Sprintf("http://localhost:8000/events/%d", eventID)

	// petición para reducir 1 boleto
	payload := map[string]interface{}{
		"available_tickets": -1,
	}
	payloadBytes, _ := json.Marshal(payload)

	// 🔄 Hacer la solicitud HTTP PUT
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return fmt.Errorf("error creando la solicitud: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error en la petición: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return fmt.Errorf("evento no encontrado en API 1")
	} else if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error en API 1, código: %d", resp.StatusCode)
	}

	fmt.Println("✅ Boletos actualizados correctamente")
	return nil
}

func (router *Router) Run() {
	_, _, _, _, _, getPedidoController := infrastructure.InitPedidoDependencies()

	pedidoGroup := router.engine.Group("/pedidos")
	{
		pedidoGroup.POST("/log", logPedidoHandler)
		pedidoGroup.GET("/get/:id", getPedidoController.Run)
		pedidoGroup.OPTIONS("/*any", func(c *gin.Context) {
			c.Status(204)
		})
	}
}
