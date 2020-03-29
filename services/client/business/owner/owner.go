package owner

import (
	"context"
	"github.com/AkezhanOb1/diplomaProject/api/graphQL/graph/model"
	config "github.com/AkezhanOb1/diplomaProject/configs"
	pb "github.com/AkezhanOb1/diplomaProject/api/proto/business/owners"
	"google.golang.org/grpc"
	"log"
)


//CreateBusinessOwner is a client function for creating a business owner
func CreateBusinessOwner(ctx context.Context, req model.CreateBusinessOwnerRequest) (*pb.CreateBusinessOwnerResponse, error) {
	cc, err := grpc.Dial(config.RpcServerAddress, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	defer func(){
		err = cc.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	c := pb.NewBusinessCompaniesServiceClient(cc)

	r := pb.CreateBusinessOwnerRequest{
		BusinessOwnerName:              req.BusinessOwnerName,
		BusinessOwnerEmail:             req.BusinessOwnerEmail,
		BusinessOwnerPassword:          req.BusinessOwnerPassword,
		BusinessOwnerPhoneNumberPrefix: req.BusinessOwnerPhoneNumberPrefix,
		BusinessOwnerPhoneNumber:       req.BusinessOwnerPhoneNumber,
		BusinessCompanyID:              req.BusinessCompanyID,
	}

	businessOwner, err := c.CreateBusinessOwner(context.Background(), &r)
	if err != nil {
		return nil, err
	}

	return businessOwner, nil
}

