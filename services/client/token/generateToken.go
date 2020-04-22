package token

import (
	"context"
	pb "github.com/AkezhanOb1/auth/api"
	c "github.com/AkezhanOb1/diplomaProject/configs"

	"google.golang.org/grpc"
)

//GenerateToken is
func GenerateToken(email string, password string) (*pb.Token, error) {
	opts := grpc.WithInsecure()
	cc, err := grpc.Dial(c.TokenServer, opts)
	if err != nil {
		return nil, err
	}
	defer cc.Close()

	client := pb.NewCompanyServicesClient(cc)
	request := &pb.GenerateTokenRequest{
		Credentials: &pb.ClientCredentials{
			Email:    email,
			Password: password,
		},
	}

	resp, err := client.GenerateToken(context.Background(), request)
	if err != nil {
		return nil, err
	}

	return resp.Token, nil
}
