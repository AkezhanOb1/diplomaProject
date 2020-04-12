package main

import (
	"context"
	"log"
	s "github.com/AkezhanOb1/diplomaProject/services/client/business/service"
)

func main() {
	log.Println("Hello world")
	log.Println(s.GetBusinessService(context.Background(), 2))
}