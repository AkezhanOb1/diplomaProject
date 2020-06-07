package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pbCategory "github.com/AkezhanOb1/diplomaProject/api/proto/business/categories"
	pbCompanyService "github.com/AkezhanOb1/diplomaProject/api/proto/business/companyServices"
	pbSubCategory "github.com/AkezhanOb1/diplomaProject/api/proto/business/subCategories"

	"github.com/AkezhanOb1/diplomaProject/services/business/category"
	"github.com/AkezhanOb1/diplomaProject/services/business/companyService"
	"github.com/AkezhanOb1/diplomaProject/services/business/subCategory"
)

func main() {
	address := "0.0.0.0:50051"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	log.Printf("Server is listening on %v ...", address)

	s := grpc.NewServer()

	pbSubCategory.RegisterBusinessSubCategoryServiceServer(s, &subCategory.Server{})
	pbCategory.RegisterBusinessCategoryServiceServer(s, &category.Server{})
	pbCompanyService.RegisterCompanyServicesServer(s, &companyService.Server{})

	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}
}
