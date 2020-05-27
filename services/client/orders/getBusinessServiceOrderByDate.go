package orders

import (
	"context"
	c "github.com/AkezhanOb1/diplomaProject/configs"
	pb "github.com/AkezhanOb1/diplomaProject/api/proto/order"
	"google.golang.org/grpc"
)

//GetBusinessServiceOrderByDate is
func GetBusinessServiceOrderByDate(ctx context.Context, businessServiceID int64, date string) (*pb.GetBusinessServiceOrderByDateResponse, error) {
	opts := grpc.WithInsecure()
	cc, err := grpc.Dial(c.BusinessServiceOrderServer, opts)
	if err != nil {
		return nil, err
	}
	defer cc.Close()



	c := pb.NewBusinessServiceOrdersClient(cc)
	e := pb.GetBusinessServiceOrderByDateRequest{
		BusinessServiceID:businessServiceID,
		Date:	date,
	}

	orders, err := c.GetBusinessServiceOrderByDate(context.Background(), &e)
	if err != nil {
		return nil, err
	}

	return orders, nil
}
