package client

import (
	"context"
	pb "github.com/AkezhanOb1/diplomaProject/api/proto/auth"
	"github.com/AkezhanOb1/diplomaProject/configs"
	"google.golang.org/grpc"
)



//RetrieveTokenInformation is a client function for creating a business owner
func RetrieveTokenInformation(ctx context.Context, request *pb.RetrieveCustomerTokenInformationRequest) (*pb.RetrieveCustomerTokenInformationResponse, error) {
	cc, err := grpc.Dial(config.TokenServer, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	defer cc.Close()

	c := pb.NewCompanyServicesClient(cc)

	businessOwner, err := c.RetrieveCustomerTokenInformation(ctx, request)
	if err != nil {
		return nil, err
	}

	return businessOwner, nil
}

