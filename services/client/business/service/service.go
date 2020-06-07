package service

import (
	"context"
	pb "github.com/AkezhanOb1/diplomaProject/api/proto/business-service"
	config "github.com/AkezhanOb1/diplomaProject/configs"
	"github.com/golang/protobuf/ptypes/empty"

	"google.golang.org/grpc"
)


//GetBusinessService is a client for graphQL on gRPC services
func GetBusinessService(ctx context.Context, id int64) (*pb.GetBusinessServiceResponse, error) {

	cc, err := grpc.Dial(config.BusinessServiceServer, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer cc.Close()

	c := pb.NewBusinessServicesClient(cc)

	e := pb.GetBusinessServiceRequest{
		BusinessServiceID:   id,
	}

	service, err := c.GetBusinessService(ctx, &e)
	if err != nil {
		return nil, err
	}

	return service, nil
}


//GetBusinessServices is a client for graphQL on gRPC services
func GetBusinessServices(ctx context.Context) (*pb.GetBusinessServicesResponse, error) {
	cc, err := grpc.Dial(config.BusinessServiceServer, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	defer cc.Close()


	c := pb.NewBusinessServicesClient(cc)
	e := empty.Empty{}

	services, err := c.GetBusinessServices(ctx, &e)
	if err != nil {
		return nil, err
	}

	return services, nil
}


//GetServicesUnderSubCategory is a client for graphQL on gRPC services
func GetServicesUnderSubCategory(ctx context.Context, subCategoryID int64) (*pb.GetServicesUnderSubCategoryResponse, error) {
	cc, err := grpc.Dial(config.BusinessServiceServer, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	defer  cc.Close()


	c := pb.NewBusinessServicesClient(cc)
	e := pb.GetServicesUnderSubCategoryRequest{
		SubCategoryID:   subCategoryID,
	}

	services, err := c.GetServicesUnderSubCategory(ctx, &e)
	if err != nil {
		return nil, err
	}

	return services, nil
}


//CreateBusinessService is a client for graphQL on gRPC services
func CreateBusinessService(ctx context.Context, serviceName string, subCategories []int64) (*pb.CreateBusinessServiceResponse, error) {
	cc, err := grpc.Dial(config.RpcServerAddress, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	defer cc.Close()

	c := pb.NewBusinessServicesClient(cc)
	e := pb.CreateBusinessServiceRequest{
		BusinessServiceName: serviceName,
		SubCategories:  subCategories,
	}

	services, err := c.CreateBusinessService(ctx, &e)
	if err != nil {
		return nil, err
	}

	return services, nil
}





