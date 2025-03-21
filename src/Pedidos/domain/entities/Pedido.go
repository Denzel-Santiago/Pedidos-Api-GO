//pedidos-api-go/src/Pedidos/domain/entities/Pedido.go
package entities

import (
	"time"
)

type Pedido struct {
	ID        int       `json:"id"`
	EventID   int       `json:"event_id"`
	UserName  string    `json:"user_name"`
	Email     string    `json:"email"`
	Quantity  int       `json:"quantity"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
