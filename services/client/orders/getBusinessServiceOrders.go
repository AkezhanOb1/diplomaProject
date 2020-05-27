package orders

import (
	"context"
	c "github.com/AkezhanOb1/diplomaProject/configs"
	pb "github.com/AkezhanOb1/diplomaProject/api/proto/order"
	"google.golang.org/grpc"
)

//GetBusinessServiceOrders is
func GetBusinessServiceOrders(ctx context.Context, businessServiceID int64) (*pb.GetBusinessServiceOrdersResponse, error) {
	opts := grpc.WithInsecure()
	cc, err := grpc.Dial(c.BusinessServiceOrderServer, opts)
	if err != nil {
		return nil, err
	}
	defer cc.Close()



	c := pb.NewBusinessServiceOrdersClient(cc)
	e := pb.GetBusinessServiceOrdersRequest{
		BusinessServiceID:businessServiceID,
	}

	orders, err := c.GetBusinessServiceOrders(context.Background(), &e)
	if err != nil {
		return nil, err
	}

	return orders, nil
}
