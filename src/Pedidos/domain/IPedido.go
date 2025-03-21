//pedidos-api-go/src/Pedidos/domain/IPedido.go
package domain

import (
	"Pedidos-Api/src/Pedidos/domain/entities"
)

type IPedido interface {
	Save(pedido entities.Pedido) error
	Update(id int, pedido entities.Pedido) error
	Delete(id int) error
	FindByID(id int) (entities.Pedido, error)
	GetAll() ([]entities.Pedido, error)
	GetByEventID(eventID int) ([]entities.Pedido, error)
}
