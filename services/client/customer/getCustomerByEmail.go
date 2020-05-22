package client

import (
	"context"
	pb "github.com/AkezhanOb1/diplomaProject/api/proto/customer"
	"github.com/AkezhanOb1/diplomaProject/configs"
	"google.golang.org/grpc"
)


//GetCustomerByEmail is a client function for creating a business owner
func GetCustomerByEmail(ctx context.Context, request *pb.GetCustomerByEmailRequest) (*pb.GetCustomerByEmailResponse, error) {
	cc, err := grpc.Dial(config.CustomerServerAddress, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	defer cc.Close()

	c := pb.NewCustomerServiceClient(cc)

	businessOwner, err := c.GetCustomerByEmail(ctx, request)
	if err != nil {
		return nil, err
	}

	return businessOwner, nil
}


