package handler

import (
	"fmt"
	"net/http"
)

func User(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "hello sajib", http.StatusOK)
}