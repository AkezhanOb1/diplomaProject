package owner

import (
	"context"
	pb "github.com/AkezhanOb1/diplomaProject/api/proto/business/owners"
	db "github.com/AkezhanOb1/diplomaProject/repositories/business/owner"
	"github.com/AkezhanOb1/diplomaProject/services/email"
)

type Server struct{}



//GetBusinessOwnerCompanies
func (s *Server) GetBusinessOwnerCompanies(ctx context.Context, request *pb.GetBusinessOwnerCompaniesRequest) (*pb.GetBusinessOwnerCompaniesResponse, error) {
	companies, err := db.GetBusinessOwnerCompaniesRepository(ctx, request.GetEmail())
	if err != nil {
		return nil, err
	}

	return companies, nil
}

//CreateBusinessOwner is
func (s *Server) CreateBusinessOwner(ctx context.Context, request *pb.CreateBusinessOwnerRequest) (*pb.CreateBusinessOwnerResponse, error) {
	hashedPassword, err := hashPassword(request.GetBusinessOwnerPassword())
	if err != nil {
		return nil, err
	}
	request.BusinessOwnerPassword = hashedPassword

	businessOwner, err := db.CreateOwnerRepository(ctx, request)
	if err != nil {
		return nil, err
	}

	err = email.BusinessRegistrationEmail(request.GetBusinessOwnerEmail())
	if err != nil {
		return nil, err
	}

	bindRequest := &pb.BindCompanyToBusinessOwnerRequest{
		BusinessOwnerID:      businessOwner.BusinessOwner.GetBusinessOwnerID(),
		BusinessCompanyID:    request.GetBusinessCompanyID(),
	}

	_, err = s.BindCompanyToBusinessOwner(context.Background(), bindRequest)
	if err != nil {
		return nil, err
	}

	return businessOwner, nil
}

//BindCompanyToBusinessOwner is
func (*Server) BindCompanyToBusinessOwner(ctx context.Context, request *pb.BindCompanyToBusinessOwnerRequest) (*pb.BindCompanyToBusinessOwnerResponse, error) {
	binding, err := db.BindCompanyToBusinessOwner(ctx, request)
	if err != nil {
		return nil, err
	}

	return binding, nil
}

//CheckOwnerPassword is
func (*Server) CheckOwnerPassword(ctx context.Context, request *pb.CheckOwnerPasswordRequest) (*pb.CheckOwnerPasswordResponse, error) {
	password, err := db.GetOwnerPasswordRepository(ctx, request.GetEmail())
	if err != nil {
		return nil, err
	}

	err = comparePassword(*password, request.GetPassword())
	if err != nil {
		return nil, err
	}
	return &pb.CheckOwnerPasswordResponse{
		Valid:true,
	}, nil
}