package domain

import (
	"Pedidos-Api/src/pedidos/domain/entities"
	"database/sql"
)

type PedidoRepository interface {
	Save(pedido entities.Pedido) error
	FindByID(id int) (*entities.Pedido, error)
}

type pedidoRepository struct {
	db *sql.DB
}

func NewPedidoRepository(db *sql.DB) PedidoRepository {
	return &pedidoRepository{db: db}
}

func (r *pedidoRepository) Save(pedido entities.Pedido) error {
	_, err := r.db.Exec("INSERT INTO orders (event_id, user_name, email, quantity, status) VALUES (?, ?, ?, ?, ?)",
		pedido.EventID, pedido.UserName, pedido.Email, pedido.Quantity, pedido.Status)
	return err
}

func (r *pedidoRepository) FindByID(id int) (*entities.Pedido, error) {
	var pedido entities.Pedido
	query := "SELECT id, event_id, user_name, email, quantity, status, created_at FROM orders WHERE id = ?"
	err := r.db.QueryRow(query, id).Scan(&pedido.ID, &pedido.EventID, &pedido.UserName, &pedido.Email, &pedido.Quantity, &pedido.Status, &pedido.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &pedido, nil
}
