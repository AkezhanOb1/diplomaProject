package orders

import (
	"context"
	graph "github.com/AkezhanOb1/diplomaProject/api/graphQL/graph/model"
	pb "github.com/AkezhanOb1/diplomaProject/api/proto/order"
	config "github.com/AkezhanOb1/diplomaProject/configs"
	"google.golang.org/grpc"
)

//GetBusinessServiceOrdersByEmail is
func GetBusinessServiceOrdersByEmail(ctx context.Context, request *graph.GetBusinessServiceOrdersByEmailRequest) (*pb.GetBusinessServiceOrdersByEmailResponse, error) {
	opts := grpc.WithInsecure()
	cc, err := grpc.Dial(config.BusinessServiceOrderServer, opts)
	if err != nil {
		return nil, err
	}
	defer cc.Close()

	var count int64
	if request.Pagination.Count != nil {
		count = *request.Pagination.Count
	}
	c := pb.NewBusinessServiceOrdersClient(cc)
	e := pb.GetBusinessServiceOrdersByEmailRequest{
		Email:      request.Email,
		Pagination: &pb.Pagination{
			Limit:  request.Pagination.Limit,
			Offset: request.Pagination.Offset,
			Count:  count,
		},
	}

	orders, err := c.GetBusinessServiceOrdersByEmail(context.Background(), &e)
	if err != nil {
		return nil, err
	}

	return orders, nil
}
