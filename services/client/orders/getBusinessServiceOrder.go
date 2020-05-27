package orders

import (
	"context"
	pb "github.com/AkezhanOb1/diplomaProject/api/proto/order"
	c "github.com/AkezhanOb1/diplomaProject/configs"
	"google.golang.org/grpc"
)

//GetBusinessServiceOrder is
func GetBusinessServiceOrder(ctx context.Context, orderID int64) (*pb.GetBusinessServiceOrderResponse, error) {
	opts := grpc.WithInsecure()
	cc, err := grpc.Dial(c.BusinessServiceOrderServer, opts)
	if err != nil {
		return nil, err
	}
	defer cc.Close()

	client := pb.NewBusinessServiceOrdersClient(cc)
	request := &pb.GetBusinessServiceOrderRequest{
		BusinessServiceOrderID: orderID,
	}

	order, err := client.GetBusinessServiceOrder(context.Background(), request)
	if err != nil {
		return nil, err
	}

	return order, nil
}
