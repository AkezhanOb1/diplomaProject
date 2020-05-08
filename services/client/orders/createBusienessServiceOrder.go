package orders

import (
	"context"
	graph "github.com/AkezhanOb1/diplomaProject/api/graphQL/graph/model"
	pb "github.com/AkezhanOb1/orders/api"
	c "github.com/AkezhanOb1/diplomaProject/configs"

	"google.golang.org/grpc"
)

//CreateBusinessServiceOrder is
func CreateBusinessServiceOrder(ctx context.Context, req graph.CreateBusinessServiceOrderRequest) (*pb.CreateBusinessServiceOrderResponse, error) {
	opts := grpc.WithInsecure()
	cc, err := grpc.Dial(c.BusinessServiceOrderServer, opts)
	if err != nil {
		return nil, err
	}
	defer cc.Close()

	client := pb.NewBusinessServiceOrdersClient(cc)
	request := &pb.CreateBusinessServiceOrderRequest{
		ClientID:                req.ClientID,
		BusinessServiceID:       req.BusinessServiceID,
		OrderDate:               req.OrderDate,
		PrePaid:                 req.PrePaid,
		ClientFirstName:         req.ClientFirstName,
		ClientPhoneNumber:       req.ClientPhoneNumber,
		ClientPhoneNumberPrefix: req.ClientPhoneNumberPrefix,
		ClientCommentary:        req.ClientCommentary,
	}

	newOrder, err := client.CreateBusinessServiceOrder(context.Background(), request)
	if err != nil {
		return nil, err
	}

	return newOrder, nil
}
