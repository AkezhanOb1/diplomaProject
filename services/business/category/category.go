package category


import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"

	pb "github.com/AkezhanOb1/diplomaProject/api/proto/business/categories"
	db "github.com/AkezhanOb1/diplomaProject/repositories/business/category"
)

type Server struct {}

func (*Server) GetBusinessCategories(ctx context.Context, emp *empty.Empty) (*pb.BusinessCategoriesResponse, error) {
	businessCategories, err := db.GetCategoriesRepository()
	if err != nil {
		panic(err)
	}

	return businessCategories, nil
}

//GetBusinessCategory is
func (*Server) GetBusinessCategory(ctx context.Context, request *pb.BusinessCategoryRequest) (*pb.BusinessCategoryResponse, error) {
	businessCategories, err := db.GetCategoryRepository(request.GetBusinessCategoryID())
	if err != nil {
		panic(err)
	}

	return businessCategories, nil
}



