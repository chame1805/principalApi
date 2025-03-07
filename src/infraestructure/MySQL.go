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
	// Aquí se usa 'numeroPersonas' como nombre genérico
	query := "INSERT INTO reservas (name, fecha, hora, numeroPersonas, service) VALUES (?, ?, ?, ?, ?)"

	// Agregar logs para depuración
	log.Printf("Ejecutando consulta: %s, con valores: name=%s, fecha=%s, hora=%s, numeroPersonas=%d, service=%s", query, name, fecha, hora, numeroPersonas, service)

	result, err := mysql.conn.ExecutePreparedQuery(query, name, fecha, hora, numeroPersonas, service)
	if err != nil {
		log.Printf("Error al guardar la reserva: %v", err)
		return err
	}

	// Verificar si realmente se insertó un registro
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		log.Println("[MySQL] - No se insertó ninguna reserva")
		return fmt.Errorf("no se insertó ninguna reserva")
	}

	log.Println("[MySQL] - Reserva guardada exitosamente")
	return nil
}

// Implementación de GetAll para reservas
func (mysql *MySQLReserva) GetAll() ([]map[string]interface{}, error) {
	query := "SELECT * FROM reservas"
	rows := mysql.conn.FetchRows(query)
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
			"id":               id,
			"name":             name,
			"fecha":            fecha,
			"hora":             hora,
			"numeroPersonas":   numeroPersonas,  // Cambié el nombre de la clave aquí
			"service":          service,
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
	query := "SELECT * FROM reservas WHERE id = ?"
	rows := mysql.conn.FetchRows(query, id)
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
			"id":               id,
			"name":             name,
			"fecha":            fecha,
			"hora":             hora,
			"numeroPersonas":   numeroPersonas,  // Cambié el nombre de la clave aquí también
			"service":          service,
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
	_, err := mysql.conn.ExecutePreparedQuery(query, name, fecha, hora, numeroPersonas, service, id)
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
	_, err := mysql.conn.ExecutePreparedQuery(query, id)
	if err != nil {
		log.Printf("Error al eliminar la reserva: %v", err)
		return err
	}

	log.Println("[MySQL] - Reserva eliminada exitosamente")
	return nil
}
