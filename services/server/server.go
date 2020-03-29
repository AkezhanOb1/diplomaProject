package main

import (
	"google.golang.org/grpc"
	"log"
	"net"

	category "github.com/AkezhanOb1/diplomaProject/api/proto/business/categories"
	company "github.com/AkezhanOb1/diplomaProject/api/proto/business/companies"
	owner "github.com/AkezhanOb1/diplomaProject/api/proto/business/owners"

	categoryService "github.com/AkezhanOb1/diplomaProject/services/business/category"
	companyService "github.com/AkezhanOb1/diplomaProject/services/business/company"
	ownerService "github.com/AkezhanOb1/diplomaProject/services/business/owner"
)


func main() {
	address := "0.0.0.0:50051"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	log.Printf("Server is listening on %v ...", address)

	s := grpc.NewServer()

	category.RegisterBusinessCategoryServiceServer(s, &categoryService.Server{})
	company.RegisterBusinessCompaniesServiceServer(s, &companyService.Server{})
	owner.RegisterBusinessCompaniesServiceServer(s, &ownerService.Server{})
	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}
}
