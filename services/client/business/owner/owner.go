package owner

import (
	"context"
	"github.com/AkezhanOb1/diplomaProject/api/graphQL/graph/model"
	pb "github.com/AkezhanOb1/diplomaProject/api/proto/business-owner"
	config "github.com/AkezhanOb1/diplomaProject/configs"
	"google.golang.org/grpc"
)


//GetBusinessOwnerCompanies is a
func GetBusinessOwnerCompanies(ctx context.Context, email string) (*pb.GetBusinessOwnerCompaniesResponse, error) {
	cc, err := grpc.Dial(config.BusinessOwnerServer, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	defer cc.Close()

	c := pb.NewBusinessOwnerServiceClient(cc)

	r := pb.GetBusinessOwnerCompaniesRequest{
		Email:email,
	}

	businessOwner, err := c.GetBusinessOwnerCompanies(ctx, &r)
	if err != nil {
		return nil, err
	}

	return businessOwner, nil
}

//CreateBusinessOwner is a client function for creating a business owner
func CreateBusinessOwner(ctx context.Context, req model.CreateBusinessOwnerRequest) (*pb.CreateBusinessOwnerResponse, error) {
	cc, err := grpc.Dial(config.BusinessOwnerServer, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	defer cc.Close()

	c := pb.NewBusinessOwnerServiceClient(cc)

	r := pb.CreateBusinessOwnerRequest{
		BusinessOwnerName:              req.BusinessOwnerName,
		BusinessOwnerEmail:             req.BusinessOwnerEmail,
		BusinessOwnerPassword:          req.BusinessOwnerPassword,
		BusinessOwnerPhoneNumberPrefix: req.BusinessOwnerPhoneNumberPrefix,
		BusinessOwnerPhoneNumber:       req.BusinessOwnerPhoneNumber,
		BusinessCompanyID:              req.BusinessCompanyID,
	}

	businessOwner, err := c.CreateBusinessOwner(ctx, &r)
	if err != nil {
		return nil, err
	}

	return businessOwner, nil
}

//CheckOwnerPassword is
func CheckOwnerPassword(ctx context.Context, email string, password string) (*pb.CheckOwnerPasswordResponse, error) {
	cc, err := grpc.Dial(config.BusinessOwnerServer, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	defer cc.Close()

	c := pb.NewBusinessOwnerServiceClient(cc)

	r := pb.CheckOwnerPasswordRequest{
		Email: email,
		Password: password,
	}

	valid, err := c.CheckOwnerPassword(context.Background(), &r)
	if err != nil {
		return nil, err
	}

	return valid, nil
}


func BusinessOwnerPasswordForgot(ctx context.Context, email string) (*pb.BusinessOwnerPasswordForgotResponse, error) {
	cc, err := grpc.Dial(config.BusinessOwnerServer, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	defer cc.Close()

	c := pb.NewBusinessOwnerServiceClient(cc)

	r := pb.BusinessOwnerPasswordForgotRequest{
		BusinessOwnerEmail:    email,
	}

	res, err := c.BusinessOwnerPasswordForgot(context.Background(), &r)
	if err != nil {
		return nil, err
	}

	return res, nil
}


func BusinessOwnerPasswordReset(ctx context.Context, email string, password string) (*pb.ResetBusinessOwnerPasswordResponse, error) {
	cc, err := grpc.Dial(config.BusinessOwnerServer, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	defer cc.Close()

	c := pb.NewBusinessOwnerServiceClient(cc)

	r := pb.ResetBusinessOwnerPasswordRequest{
		BusinessOwnerEmail:    email,
		BusinessOwnerPassword: password,
	}

	owner, err := c.ResetBusinessOwnerPassword(context.Background(), &r)
	if err != nil {
		return nil, err
	}

	return owner, nil
}