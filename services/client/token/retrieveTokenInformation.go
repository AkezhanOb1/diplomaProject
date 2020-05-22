
package token

import (
	"context"
	pb "github.com/AkezhanOb1/diplomaProject/api/proto/auth"
	c "github.com/AkezhanOb1/diplomaProject/configs"
	"google.golang.org/grpc"
)

//RetrieveTokenInformation is
func RetrieveTokenInformation(accessToken string) (*pb.RetrieveTokenInformationResponse, error) {
	opts := grpc.WithInsecure()
	cc, err := grpc.Dial(c.TokenServer, opts)
	if err != nil {
		return nil, err
	}
	defer cc.Close()

	client := pb.NewCompanyServicesClient(cc)
	request := &pb.RetrieveTokenInformationRequest{
		AccessToken: accessToken,
	}

	resp, err := client.RetrieveTokenInformation(context.Background(), request)
	if err != nil {
		return nil, err
	}

	return resp, err
}

