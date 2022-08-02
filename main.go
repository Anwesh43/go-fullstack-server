package main 
import (
	"fullstackserver/server"
	"fullstackserver/db"
	"fmt"
)
func main() {

	database, err := db.Connect("localhost", "devme", "password", "fullstackdb", "5432")
	if err == nil  {
		fmt.Println("connected to db successfully")
	}
	server.Start(database)
	fmt.Println("started server on 3000")
}