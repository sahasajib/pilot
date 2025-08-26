package main

import (
	

	"github.com/sahasajib/pilot/cmd"
	"github.com/sahasajib/pilot/database"
)


func main(){
	database.InitDB()
	cmd.Serve()

}