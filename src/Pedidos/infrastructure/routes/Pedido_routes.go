package routes

import (
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

func (router *Router) Run() {
	// Inicializamos las dependencias de pedidos
	createController, viewController, updateController, deleteController, viewAllController, getPedidoController := infrastructure.InitPedidoDependencies()

	// Grupo de rutas para pedidos
	pedidoGroup := router.engine.Group("/pedidos")
	{
		// ✅ Crear un pedido
		pedidoGroup.POST("/", createController.Run)

		// ✅ Obtener un pedido por ID
		pedidoGroup.GET("/:id", viewController.Run)

		// ✅ Actualizar un pedido por ID
		pedidoGroup.PUT("/:id", updateController.Run)

		// ✅ Eliminar un pedido por ID
		pedidoGroup.DELETE("/:id", deleteController.Run)

		// ✅ Obtener todos los pedidos
		pedidoGroup.GET("/", viewAllController.Run)

		// ✅ Obtener un pedido por ID (alternativa)
		pedidoGroup.GET("/get/:id", getPedidoController.Run)

		// ✅ Añadir manejador OPTIONS para preflight requests
		pedidoGroup.OPTIONS("/*any", func(c *gin.Context) {
			c.Status(204) // Responder con No Content
		})
	}
}
