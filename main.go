package main

import (
	"log"
	c "github.com/AkezhanOb1/diplomaProject/services/client/orders"

)

func main() {
	log.Println("Hello world")
	log.Println(c.GetBusinessServiceOrders())
}