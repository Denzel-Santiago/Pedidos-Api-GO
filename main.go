//pedidos-api-go/main.go
package main

import (
	"fmt"

	pedidosRut "Pedidos-Api/src/Pedidos/infrastructure/routes" // Importamos las rutas de pedidos
	"Pedidos-Api/src/core"
	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializar la base de datos
	core.InitDB()

	// Crear un enrutador de Gin
	r := gin.Default()

	// ✅ Registrar Middleware CORS
	r.Use(core.CORSMiddleware())

	// ✅ Configurar las rutas de pedidos
	pedidosRouter := pedidosRut.NewRouter(r)
	pedidosRouter.Run()

	fmt.Println("¡API en Funcionamiento :D!")

	// ✅ Iniciar el servidor
	err := r.Run(":8001")
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
