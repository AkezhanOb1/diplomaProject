package orders

import (
	"context"
	config "github.com/AkezhanOb1/diplomaProject/configs"
	pb "github.com/AkezhanOb1/diplomaProject/api/proto/order"
	graph "github.com/AkezhanOb1/diplomaProject/api/graphQL/graph/model"

	"google.golang.org/grpc"
)

//UpdateBusinessServiceOrder is
func UpdateBusinessServiceOrder(ctx context.Context, request *graph.UpdateBusinessServiceOrderRequest) (*pb.UpdateBusinessServiceOrderResponse, error) {
	opts := grpc.WithInsecure()
	cc, err := grpc.Dial(config.BusinessServiceOrderServer, opts)
	if err != nil {
		return nil, err
	}
	defer cc.Close()


	c := pb.NewBusinessServiceOrdersClient(cc)
	e := pb.UpdateBusinessServiceOrderRequest{
		OrderID:                 request.OrderID,
		BusinessServiceID:       request.BusinessServiceID,
		StartAt:                 request.StartAt,
		PrePaid:                 request.PrePaid,
		ClientFirstName:         request.ClientFirstName,
		ClientPhoneNumber:       request.ClientPhoneNumber,
		ClientPhoneNumberPrefix: request.ClientPhoneNumberPrefix,
		ClientCommentary:        request.ClientCommentary,
	}

	order, err := c.UpdateBusinessServiceOrder(context.Background(), &e)
	if err != nil {
		return nil, err
	}

	return order, nil
}
