package orders

import (
	"context"
	pb "github.com/AkezhanOb1/orders/api"
	c "github.com/AkezhanOb1/diplomaProject/configs"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

//GetBusinessServiceOrders is
func GetBusinessServiceOrders(ctx context.Context) (*pb.GetBusinessServiceOrdersResponse, error) {
	opts := grpc.WithInsecure()
	cc, err := grpc.Dial(c.BusinessServiceOrderServer, opts)
	if err != nil {
		return nil, err
	}
	defer cc.Close()



	c := pb.NewBusinessServiceOrdersClient(cc)
	e := empty.Empty{}

	orders, err := c.GetBusinessServiceOrders(context.Background(), &e)
	if err != nil {
		return nil, err
	}

	return orders, nil
}
