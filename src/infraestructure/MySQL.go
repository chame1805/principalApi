package infraestructure

import (
	"fmt"
	"log"
	"principalApi/src/core"
	
)

type MySQLReserva struct {
	conn *core.Conn_MySQL
}

// Constructor
func NewMySQLReserva() *MySQLReserva {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &MySQLReserva{conn: conn}
}

// Implementación de Save para reservas
func (mysql *MySQLReserva) Save(name string, fecha string, hora string, numeroPersonas int, service string) error {
	query := "INSERT INTO reservas (name, fecha, hora, numeroPersonas, service) VALUES (?, ?, ?, ?, ?)"
	result, err := mysql.conn.DB.Exec(query, name, fecha, hora, numeroPersonas, service)
	if err != nil {
		log.Printf("Error al guardar la reserva: %v", err)
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		log.Println("[MySQL] - No se insertó ninguna reserva")
		return fmt.Errorf("no se insertó ninguna reserva")
	}

	log.Println("[MySQL] - Reserva guardada exitosamente")
	return nil
}

// ✅ Implementación correcta de GetAllReservas
func (mysql *MySQLReserva) GetAllReservas() ([]map[string]interface{}, error) {
	query := "SELECT id, name, fecha, hora, numeroPersonas, service FROM reservas"
	rows, err := mysql.conn.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reservas []map[string]interface{}
	for rows.Next() {
		var id int
		var name, fecha, hora, service string
		var numeroPersonas int
		if err := rows.Scan(&id, &name, &fecha, &hora, &numeroPersonas, &service); err != nil {
			return nil, err
		}
		reserva := map[string]interface{}{
			"id":             id,
			"name":           name,
			"fecha":          fecha,
			"hora":           hora,
			"numeroPersonas": numeroPersonas,
			"service":        service,
		}
		reservas = append(reservas, reserva)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return reservas, nil
}

// Implementación de GetReserva por ID
func (mysql *MySQLReserva) GetReserva(id int) (map[string]interface{}, error) {
	query := "SELECT id, name, fecha, hora, numeroPersonas, service FROM reservas WHERE id = ?"
	rows, err := mysql.conn.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reserva map[string]interface{}
	if rows.Next() {
		var id int
		var name, fecha, hora, service string
		var numeroPersonas int
		if err := rows.Scan(&id, &name, &fecha, &hora, &numeroPersonas, &service); err != nil {
			return nil, err
		}
		reserva = map[string]interface{}{
			"id":             id,
			"name":           name,
			"fecha":          fecha,
			"hora":           hora,
			"numeroPersonas": numeroPersonas,
			"service":        service,
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return reserva, nil
}

// Implementación de Update para reservas
func (mysql *MySQLReserva) Update(id int, name string, fecha string, hora string, numeroPersonas int, service string) error {
	query := "UPDATE reservas SET name = ?, fecha = ?, hora = ?, numeroPersonas = ?, service = ? WHERE id = ?"
	_, err := mysql.conn.DB.Exec(query, name, fecha, hora, numeroPersonas, service, id)
	if err != nil {
		log.Printf("Error al actualizar la reserva: %v", err)
		return err
	}

	log.Println("[MySQL] - Reserva actualizada exitosamente")
	return nil
}

// Implementación de Delete para reservas
func (mysql *MySQLReserva) Delete(id int) error {
	query := "DELETE FROM reservas WHERE id = ?"
	_, err := mysql.conn.DB.Exec(query, id)
	if err != nil {
		log.Printf("Error al eliminar la reserva: %v", err)
		return err
	}

	log.Println("[MySQL] - Reserva eliminada exitosamente")
	return nil
}
