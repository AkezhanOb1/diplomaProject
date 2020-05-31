package subCategories

import (
	"context"
	pb "github.com/AkezhanOb1/diplomaProject/api/proto/business/subCategories"
	config "github.com/AkezhanOb1/diplomaProject/configs"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"os"
)


//GetBusinessSubCategory is a client for graphQL on gRPC services
func GetBusinessSubCategory(ctx context.Context, id int64) (*pb.BusinessSubCategoryResponse, error) {

	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	cert := wd + "/static/qaqtus_me.crt"
	credential, err := credentials.NewClientTLSFromFile(cert, "")
	if err != nil {
		return nil, err
	}
	cc, err := grpc.Dial(config.RpcServerAddress, grpc.WithTransportCredentials(credential))
	if err != nil {
		return nil, err
	}
	defer cc.Close()


	c := pb.NewBusinessSubCategoryServiceClient(cc)

	e := pb.BusinessSubCategoryRequest{
		BusinessSubCategoryID:   id,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}

	subCategory, err := c.GetBusinessSubCategory(ctx, &e)
	if err != nil {
		return nil, err
	}

	return subCategory, nil
}


//GetBusinessSubCategories is a client for graphQL on gRPC services
func GetBusinessSubCategories(ctx context.Context) (*pb.BusinessSubCategoriesResponse, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	cert := wd + "/static/qaqtus_me.crt"
	credential, err := credentials.NewClientTLSFromFile(cert, "")
	if err != nil {
		return nil, err
	}
	cc, err := grpc.Dial(config.RpcServerAddress, grpc.WithTransportCredentials(credential))
	if err != nil {
		return nil, err
	}

	defer cc.Close()


	c := pb.NewBusinessSubCategoryServiceClient(cc)
	e := empty.Empty{}

	subCategories, err := c.GetBusinessSubCategories(ctx, &e)
	if err != nil {
		return nil, err
	}

	return subCategories, nil
}


//GetBusinessSubCategoriesUnderCategory is a client for graphQL on gRPC services
func GetBusinessSubCategoriesUnderCategory(ctx context.Context, categoryID int64) (*pb.BusinessSubCategoriesUnderCategoryResponse, error) {

	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	cert := wd + "/static/qaqtus_me.crt"
	credential, err := credentials.NewClientTLSFromFile(cert, "")
	if err != nil {
		return nil, err
	}
	cc, err := grpc.Dial(config.RpcServerAddress, grpc.WithTransportCredentials(credential))
	if err != nil {
		return nil, err
	}

	defer cc.Close()


	c := pb.NewBusinessSubCategoryServiceClient(cc)

	e := pb.BusinessSubCategoriesUnderCategoryRequest{
		BusinessCategoryID:   categoryID,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}

	subCategories, err := c.GetBusinessSubCategoriesUnderCategory(ctx, &e)
	if err != nil {
		return nil, err
	}

	return subCategories, nil
}