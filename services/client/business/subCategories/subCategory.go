package subCategories

import (
	"context"
	config "github.com/AkezhanOb1/diplomaProject/configs"
	pb "github.com/AkezhanOb1/diplomaProject/api/proto/business/subCategories"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"log"

)


//GetBusinessSubCategory is a client for graphQL on gRPC services
func GetBusinessSubCategory(ctx context.Context, id int64) (*pb.BusinessSubCategoryResponse, error) {

	cc, err := grpc.Dial(config.RpcServerAddress, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer func(){
		err = cc.Close()
		if err != nil {
			log.Println(err)
		}
	}()

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
	cc, err := grpc.Dial(config.RpcServerAddress, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	defer func(){
		err = cc.Close()
		if err != nil {
			log.Println(err)
		}
	}()

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

	cc, err := grpc.Dial(config.RpcServerAddress, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	defer func(){
		err = cc.Close()
		if err != nil {
			log.Println(err)
		}
	}()

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