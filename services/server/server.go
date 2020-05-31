package main

import (
	"google.golang.org/grpc/credentials"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	pbCategory "github.com/AkezhanOb1/diplomaProject/api/proto/business/categories"
	pbCompany "github.com/AkezhanOb1/diplomaProject/api/proto/business/companies"
	pbCompanyService "github.com/AkezhanOb1/diplomaProject/api/proto/business/companyServices"
	pbOwner "github.com/AkezhanOb1/diplomaProject/api/proto/business/owners"
	pbBusinessService "github.com/AkezhanOb1/diplomaProject/api/proto/business/services"
	pbSubCategory "github.com/AkezhanOb1/diplomaProject/api/proto/business/subCategories"

	"github.com/AkezhanOb1/diplomaProject/services/business/category"
	"github.com/AkezhanOb1/diplomaProject/services/business/company"
	"github.com/AkezhanOb1/diplomaProject/services/business/companyService"
	"github.com/AkezhanOb1/diplomaProject/services/business/owner"
	"github.com/AkezhanOb1/diplomaProject/services/business/service"
	"github.com/AkezhanOb1/diplomaProject/services/business/subCategory"
)

func main() {
	address := "0.0.0.0:50051"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	// Create the TLS credentials
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	cert := wd + "/static/qaqtus_me.crt"
	key := wd + "/static/qaqtus.me.key"
	creds, err := credentials.NewServerTLSFromFile(cert, key)
	if err != nil {
		panic(err)
	}

	log.Printf("Server is listening on %v ...", address)

	s := grpc.NewServer(grpc.Creds(creds))

	pbSubCategory.RegisterBusinessSubCategoryServiceServer(s, &subCategory.Server{})
	pbCategory.RegisterBusinessCategoryServiceServer(s, &category.Server{})
	pbCompany.RegisterBusinessCompaniesServiceServer(s, &company.Server{})
	pbOwner.RegisterBusinessOwnerServiceServer(s, &owner.Server{})
	pbBusinessService.RegisterBusinessServicesServer(s, &service.Server{})
	pbCompanyService.RegisterCompanyServicesServer(s, &companyService.Server{})

	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}
}
