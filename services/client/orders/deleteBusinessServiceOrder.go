package orders

import (
	"context"
	config "github.com/AkezhanOb1/diplomaProject/configs"
	pb "github.com/AkezhanOb1/diplomaProject/api/proto/order"
	graph "github.com/AkezhanOb1/diplomaProject/api/graphQL/graph/model"

	"google.golang.org/grpc"
)

//DeleteBusinessServiceOrder is
func DeleteBusinessServiceOrder(ctx context.Context, request *graph.DeleteBusinessServiceOrderRequest) (*pb.DeleteBusinessServiceOrderResponse, error) {
	opts := grpc.WithInsecure()
	cc, err := grpc.Dial(config.BusinessServiceOrderServer, opts)
	if err != nil {
		return nil, err
	}
	defer cc.Close()


	c := pb.NewBusinessServiceOrdersClient(cc)
	e := pb.DeleteBusinessServiceOrderRequest{
		OrderID:                 request.OrderID,
	}

	order, err := c.DeleteBusinessServiceOrder(context.Background(), &e)
	if err != nil {
		return nil, err
	}

	return order, nil
}
