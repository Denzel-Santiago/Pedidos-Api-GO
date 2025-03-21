//pedidos-api-go/src/Pedidos/infrastructure/Dependencies.go
package infrastructure

import (
	"Pedidos-Api/src/Pedidos/application"
)

func InitPedidoDependencies() (
	*CreatePedidoController,
	*ViewPedidoController,
	*UpdatePedidoController,
	*DeletePedidoController,
	*ViewAllPedidosController,
	*GetPedidoController,
) {

	repo := NewMysqlPedidoRepository()

	createUseCase := application.NewCreatePedidoUseCase(repo)
	viewUseCase := application.NewViewPedidoUseCase(repo)
	updateUseCase := application.NewUpdatePedidoUseCase(repo)
	deleteUseCase := application.NewDeletePedidoUseCase(repo)
	viewAllUseCase := application.NewViewAllPedidosUseCase(repo)
	getPedidoUseCase := application.NewGetPedidoUseCase(repo)

	// Crear controladores
	createController := NewCreatePedidoController(createUseCase)
	viewController := NewViewPedidoController(viewUseCase)
	updateController := NewUpdatePedidoController(updateUseCase)
	deleteController := NewDeletePedidoController(deleteUseCase)
	viewAllController := NewViewAllPedidosController(viewAllUseCase)
	getPedidoController := NewGetPedidoController(getPedidoUseCase)

	return createController, viewController, updateController, deleteController, viewAllController, getPedidoController
}
