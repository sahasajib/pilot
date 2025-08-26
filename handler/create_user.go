package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/sahasajib/pilot/database"
	"github.com/sahasajib/pilot/util"
)


func CreateUser(w http.ResponseWriter, r *http.Request){
	var user database.Users

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil{
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	

	hashPassword, err := util.HashPassword(user.Password)
	if err != nil{
		log.Printf("password not intreget %v", err)
		http.Error(w, "Hash password not created", http.StatusBadRequest)
		return
	}

	db := database.DB
	if db == nil {
        log.Printf("Database connection is nil")
        http.Error(w, "Database not initialized", http.StatusInternalServerError)
        return
    }

	tx, err := db.Begin()
	if err != nil {
		log.Printf("Database connection not set: %v", err)
		http.Error(w, "database connection not cunnected", http.StatusInternalServerError)
		return
	}
	defer tx.Rollback()

	query := `INSERT INTO users (name, phone, email, password, is_delete) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err = tx.QueryRow(query, user.Name, user.Phone, user.Email, string(hashPassword), user.IsDelete).Scan(&user.ID)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tx.Commit(); err != nil{
		log.Printf("Error commiting for create user :%v", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	user.Password = ""
	util.SendData(w, user, http.StatusCreated)

}