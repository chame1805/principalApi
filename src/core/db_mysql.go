package core

import (
	"fmt"
	"log"
	"os"
	"database/sql"

	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
)

type Conn_MySQL struct {
	DB  *sql.DB
	Err string
}

func GetDBPool() *Conn_MySQL {
	// Intentamos cargar las variables de entorno desde el archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error al cargar el archivo .env: %v", err)
	}

	// Obtenemos las variables de entorno para la conexión a MySQL
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbSchema := os.Getenv("DB_NAME")

	// Imprimimos las variables para asegurarnos de que se han cargado correctamente
	fmt.Println("DB User:", dbUser)
	fmt.Println("DB Host:", dbHost)
	fmt.Println("DB Schema:", dbSchema)

	// Generamos el DSN (Data Source Name) para la conexión
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbUser, dbPass, dbHost, dbSchema)

	// Establecemos la conexión con la base de datos
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error al abrir la base de datos: %v", err)
	}

	// Establecemos las conexiones máximas
	db.SetMaxOpenConns(10)

	// Verificamos la conexión a la base de datos
	if err := db.Ping(); err != nil {
		db.Close()
		log.Fatalf("Error al verificar la conexión a la base de datos: %v", err)
	}

	// Si todo está bien, devolvemos la conexión a la base de datos
	return &Conn_MySQL{DB: db}
}

func (conn *Conn_MySQL) ExecutePreparedQuery(query string, values ...interface{}) (sql.Result, error) {
	stmt, err := conn.DB.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("error al preparar la consulta: %w", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(values...)
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar la consulta preparada: %w", err)
	}

	return result, nil
}

func (conn *Conn_MySQL) FetchRows(query string, values ...interface{}) (*sql.Rows) {
	rows, err := conn.DB.Query(query, values...)
	if err != nil {
		fmt.Printf("error al ejecutar la consulta SELECT: %v", err)
	}
	return rows
}