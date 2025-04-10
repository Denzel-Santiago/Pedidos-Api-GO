// pedidos-api-go/src/Pedidos/infrastructure/routes/Pedido_routes.go
// pedidos-api-go/src/Pedidos/infrastructure/routes/Pedido_routes.go
package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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

func logPedidoHandler(c *gin.Context) {
	var pedido map[string]interface{}

	if err := c.ShouldBindJSON(&pedido); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// 🖨️ Imprimir JSON recibido
	prettyJSON, _ := json.MarshalIndent(pedido, "", "  ")
	fmt.Println("📩 Pedido recibido desde consumer.go:")
	fmt.Println(string(prettyJSON))
	fmt.Println("----------------------------")

	// 📌 Llamar a la función para actualizar boletos
	err := actualizarBoletos(pedido)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// ✅ Confirmar actualización exitosa
	c.JSON(http.StatusOK, gin.H{
		"message": "Pedido recibido y evento actualizado correctamente",
		"data":    pedido,
	})
}

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

func actualizarBoletos(pedido map[string]interface{}) error {
	// Extraer el ID del evento
	eventID, ok := pedido["id"].(float64)
	if !ok {
		return fmt.Errorf("ID de evento inválido o no proporcionado")
	}

	// Crear un nuevo mapa solo con los datos necesarios
	payload := map[string]interface{}{
		"id":        eventID,
		"operation": "decrement", // Nueva bandera para indicar la operación
	}

	// Opcional: incluir otros campos si son necesarios
	if name, ok := pedido["name"].(string); ok {
		payload["name"] = name
	}
	if location, ok := pedido["location"].(string); ok {
		payload["location"] = location
	}

	url := fmt.Sprintf("http://localhost:8000/events/%d", int(eventID))

	// Convertir el mapa a JSON
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error serializando los datos: %v", err)
	}

	// Resto del código se mantiene igual...
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
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("error en API 1, código: %d, respuesta: %s", resp.StatusCode, string(body))
	}

	fmt.Println("✅ Solicitud de decremento enviada correctamente")
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
