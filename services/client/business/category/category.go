package category

import (
	"context"
	"log"

	pb "github.com/AkezhanOb1/diplomaProject/api/proto/business/categories"
	config "github.com/AkezhanOb1/diplomaProject/configs"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

//GetBusinessCategories is a client for graphQL on gRPC services
func GetBusinessCategories(ctx context.Context) ([]*pb.BusinessCategory, error) {
	cc, err := grpc.Dial(config.RpcServerAddress, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer cc.Close()

	c := pb.NewBusinessCategoryServiceClient(cc)
	e := empty.Empty{}

	categories, err := c.GetBusinessCategories(ctx, &e)
	if err != nil {
		return nil, err
	}

	return categories.BusinessCategories, nil
}

//GetBusinessCategory is a client for graphQL on gRPC services
func GetBusinessCategory(ctx context.Context, id int64) (*pb.BusinessCategoryResponse, error) {
	cc, err := grpc.Dial(config.RpcServerAddress, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer func() {
		err = cc.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	c := pb.NewBusinessCategoryServiceClient(cc)
	e := pb.BusinessCategoryRequest{
		BusinessCategoryID:   id,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}

	category, err := c.GetBusinessCategory(ctx, &e)
	if err != nil {
		return nil, err
	}

	return category, nil
}
