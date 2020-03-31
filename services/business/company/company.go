package company

import (
	"context"
	pb "github.com/AkezhanOb1/diplomaProject/api/proto/business/companies"
	db "github.com/AkezhanOb1/diplomaProject/repositories/business/company"

)

type Server struct {}

//CreateBusinessCompany is
func (*Server) CreateBusinessCompany(ctx context.Context, request *pb.CreateBusinessCompanyRequest) (*pb.CreateBusinessCompanyResponse, error) {
	businessCompany, err := db.CreateCompanyRepository(ctx, request)
	if err != nil {
		return nil, err
	}

	return businessCompany, nil
}

