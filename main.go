package main 
import (
	"fullstackserver/server"
	"fmt"
)
func main() {
	server.Start()
	fmt.Println("started server on 3000")
}