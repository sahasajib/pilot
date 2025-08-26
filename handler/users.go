package handler

import (
	"fmt"
	"net/http"

	_"github.com/sahasajib/pilot/database"
)

func Users(w http.ResponseWriter, r *http.Request){
	// var users database.Users


	fmt.Fprintln(w, "user not created", http.StatusOK)


}