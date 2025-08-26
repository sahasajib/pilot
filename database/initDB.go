package database

import (
	"database/sql"
	"log"
	"log/slog"
	_ "github.com/lib/pq"
)

var DB *sql.DB
func InitDB(){
	log.Printf("Starting database initialization")
	var err error
	connString := "host=localhost port=5432 user=postgres password=12345678 dbname=pilot sslmode=disable"
	DB, err = sql.Open("postgres", connString)
	if err != nil{
		log.Fatal("Faild to connect to the database", err)
	}
	err = DB.Ping()
	if err != nil{
		log.Fatal("Faild to Ping the database", err)
	}
	slog.Info("Database connection established successfully")


	createTableUser := `CREATE TABLE IF NOT EXISTS users(
		id SERIAL PRIMARY KEY,
		name VARCHAR(50) NOT NULL,
		phone VARCHAR(50) NOT NULL,
		email VARCHAR(50) NOT NULL,
		password TEXT NOT NULL,
		is_delete BOOLEAN DEFAULT FALSE
	);`
	_, err = DB.Exec(createTableUser)
	if err != nil{
		log.Fatal("user table not created")
	} 
	slog.Info("User table created successfully")
	log.Printf("Database initialization completed")
}