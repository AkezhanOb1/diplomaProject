package company

import (
	"context"
	pb "github.com/AkezhanOb1/diplomaProject/api/proto/business/companies"
	db "github.com/AkezhanOb1/diplomaProject/repositories/business/company"
	"github.com/golang/protobuf/ptypes/empty"
)

type Server struct {}


//GetBusinessCompany is
func (*Server) GetBusinessCompany(ctx context.Context, request *pb.GetBusinessCompanyRequest) (*pb.GetBusinessCompanyResponse, error) {
	businessCompany, err := db.GetBusinessCompanyRepository(ctx, request.BusinessCompanyID)
	if err != nil {
		return nil, err
	}

	return businessCompany, nil
}

//GetBusinessCompanies is
func (*Server) GetBusinessCompanies(ctx context.Context, emp *empty.Empty) (*pb.GetBusinessCompaniesResponse, error) {
	businessCompanies, err := db.GetBusinessCompaniesRepository(ctx)
	if err != nil {
		return nil, err
	}

	return businessCompanies, nil
}

//GetBusinessCompanyServices is
func (*Server) GetBusinessCompanyServices(ctx context.Context, request *pb.GetBusinessCompanyServicesRequest) (*pb.GetBusinessCompanyServicesResponse, error) {
	businessCompanyID := request.GetBusinessCompanyID()
	businessCompany, err := db.GetBusinessCompanyServicesRepository(ctx, businessCompanyID)
	if err != nil {
		return nil, err
	}

	return businessCompany, nil
}


//CreateBusinessCompany is
func (*Server) CreateBusinessCompany(ctx context.Context, request *pb.CreateBusinessCompanyRequest) (*pb.CreateBusinessCompanyResponse, error) {
	businessCompany, err := db.CreateCompanyRepository(ctx, request)
	if err != nil {
		return nil, err
	}

	return businessCompany, nil
}

