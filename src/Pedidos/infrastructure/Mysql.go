package infrastructure

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"Pedidos-Api/src/Pedidos/domain"
	"Pedidos-Api/src/Pedidos/domain/entities"
	"Pedidos-Api/src/core"
)

type MysqlPedidoRepository struct {
	conn *sql.DB
}

func NewMysqlPedidoRepository() domain.IPedido {
	conn := core.GetDB()
	return &MysqlPedidoRepository{conn: conn}
}

func (mysql *MysqlPedidoRepository) Save(pedido entities.Pedido) error {
	result, err := mysql.conn.Exec(
		"INSERT INTO pedidos (event_id, user_name, email, quantity, status, created_at) VALUES (?, ?, ?, ?, ?, ?)",
		pedido.EventID,
		pedido.UserName,
		pedido.Email,
		pedido.Quantity,
		pedido.Status,
		time.Now(),
	)
	if err != nil {
		log.Println("Error al guardar el pedido:", err)
		return err
	}

	idInserted, err := result.LastInsertId()
	if err != nil {
		log.Println("Error al obtener el ID insertado:", err)
		return err
	}

	pedido.ID = int(idInserted)
	return nil
}

func (mysql *MysqlPedidoRepository) Update(id int, pedido entities.Pedido) error {
	result, err := mysql.conn.Exec(
		"UPDATE pedidos SET event_id = ?, user_name = ?, email = ?, quantity = ?, status = ? WHERE id = ?",
		pedido.EventID,
		pedido.UserName,
		pedido.Email,
		pedido.Quantity,
		pedido.Status,
		id,
	)
	if err != nil {
		log.Println("Error al actualizar el pedido:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error al obtener filas afectadas:", err)
		return err
	}

	if rowsAffected == 0 {
		log.Println("No se encontró el pedido con ID:", id)
		return fmt.Errorf("pedido con ID %d no encontrado", id)
	}

	return nil
}

func (mysql *MysqlPedidoRepository) Delete(id int) error {
	_, err := mysql.conn.Exec("DELETE FROM pedidos WHERE id = ?", id)
	if err != nil {
		log.Println("Error al eliminar el pedido:", err)
		return err
	}
	return nil
}

func (mysql *MysqlPedidoRepository) FindByID(id int) (entities.Pedido, error) {
	var pedido entities.Pedido
	row := mysql.conn.QueryRow("SELECT id, event_id, user_name, email, quantity, status, created_at FROM pedidos WHERE id = ?", id)

	err := row.Scan(
		&pedido.ID,
		&pedido.EventID,
		&pedido.UserName,
		&pedido.Email,
		&pedido.Quantity,
		&pedido.Status,
		&pedido.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Pedido no encontrado:", err)
			return entities.Pedido{}, fmt.Errorf("pedido con ID %d no encontrado", id)
		}
		log.Println("Error al buscar el pedido por ID:", err)
		return entities.Pedido{}, err
	}

	return pedido, nil
}

func (mysql *MysqlPedidoRepository) GetAll() ([]entities.Pedido, error) {
	var pedidos []entities.Pedido

	rows, err := mysql.conn.Query("SELECT id, event_id, user_name, email, quantity, status, created_at FROM pedidos")
	if err != nil {
		log.Println("Error al obtener todos los pedidos:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var pedido entities.Pedido
		err := rows.Scan(
			&pedido.ID,
			&pedido.EventID,
			&pedido.UserName,
			&pedido.Email,
			&pedido.Quantity,
			&pedido.Status,
			&pedido.CreatedAt,
		)
		if err != nil {
			log.Println("Error al escanear pedido:", err)
			return nil, err
		}
		pedidos = append(pedidos, pedido)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error después de iterar filas:", err)
		return nil, err
	}

	return pedidos, nil
}

func (mysql *MysqlPedidoRepository) GetByEventID(eventID int) ([]entities.Pedido, error) {
	var pedidos []entities.Pedido

	rows, err := mysql.conn.Query("SELECT id, event_id, user_name, email, quantity, status, created_at FROM pedidos WHERE event_id = ?", eventID)
	if err != nil {
		log.Println("Error al obtener pedidos por event_id:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var pedido entities.Pedido
		err := rows.Scan(
			&pedido.ID,
			&pedido.EventID,
			&pedido.UserName,
			&pedido.Email,
			&pedido.Quantity,
			&pedido.Status,
			&pedido.CreatedAt,
		)
		if err != nil {
			log.Println("Error al filtrar los pedidos:", err)
			return nil, err
		}
		pedidos = append(pedidos, pedido)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error al filtrar los pedidos:", err)
		return nil, err
	}

	return pedidos, nil
}
