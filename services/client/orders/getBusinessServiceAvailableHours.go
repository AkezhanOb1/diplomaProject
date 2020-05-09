package orders

import (
	"context"
	c "github.com/AkezhanOb1/diplomaProject/configs"
	pb "github.com/AkezhanOb1/orders/api"
	"google.golang.org/grpc"
)

//GetBusinessServiceAvailableHours is
func GetBusinessServiceAvailableHours(ctx context.Context, businessServiceID int64, date string) (*pb.GetCompanyAvailableHoursByDateResponse, error) {
	opts := grpc.WithInsecure()
	cc, err := grpc.Dial(c.BusinessServiceOrderServer, opts)
	if err != nil {
		return nil, err
	}
	defer cc.Close()



	c := pb.NewBusinessServiceOrdersClient(cc)
	e := pb.GetCompanyAvailableHoursByDateRequest{
		BusinessServiceID:businessServiceID,
		Date: date,
	}

	availableHours, err := c.GetCompanyAvailableHoursByDate(context.Background(), &e)
	if err != nil {
		return nil, err
	}

	return availableHours, nil
}
