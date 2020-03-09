package owner

import (
	"context"
	"log"

	pb "github.com/AkezhanOb1/diplomaProject/api/proto/business/owners"
	db "github.com/AkezhanOb1/diplomaProject/repositories/business/owner"
)

type Server struct{}

//CreateBusinessOwner is
func (s *Server) CreateBusinessOwner(ctx context.Context, request *pb.CreateBusinessOwnerRequest) (*pb.CreateBusinessOwnerResponse, error) {
	hashedPassword, err := hashPassword(request.GetBusinessOwnerPassword())
	if err != nil {
		return nil, err
	}
	request.BusinessOwnerPassword = hashedPassword

	businessOwner, err := db.CreateOwnerRepository(ctx, request)
	if err != nil {
		panic(err)
	}

	bindRequest := &pb.BindCompanyToBusinessOwnerRequest{
		BusinessOwnerID:      businessOwner.BusinessOwner.GetBusinessOwnerID(),
		BusinessCompanyID:    request.GetBusinessCompany_ID(),
	}

	_, err = s.BindCompanyToBusinessOwner(context.Background(), bindRequest)
	if err != nil {
		panic(err)
	}

	return businessOwner, nil
}

//BindCompanyToBusinessOwner is
func (*Server) BindCompanyToBusinessOwner(ctx context.Context, request *pb.BindCompanyToBusinessOwnerRequest) (*pb.BindCompanyToBusinessOwnerResponse, error) {
	binding, err := db.BindCompanyToBusinessOwner(ctx, request)
	if err != nil {
		panic(err)
	}

	log.Println(binding)
	return binding, nil
}
