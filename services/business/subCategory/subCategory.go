
package subCategory

import (
	"context"
	pb "github.com/AkezhanOb1/diplomaProject/api/proto/business/subCategories"
	db "github.com/AkezhanOb1/diplomaProject/repositories/business/subCategory"
	"github.com/golang/protobuf/ptypes/empty"
)


type Server struct {}

//GetBusinessSubCategory is
func (*Server) GetBusinessSubCategory(ctx context.Context, request *pb.BusinessSubCategoryRequest) (*pb.BusinessSubCategoryResponse, error) {
	subCategoryID := request.GetBusinessSubCategoryID()

	SubCategory, err := db.GetSubCategoryRepository(ctx, subCategoryID)
	if err != nil {
		return nil, err
	}

	return SubCategory, nil
}

//GetBusinessSubCategories is
func (*Server) GetBusinessSubCategories(ctx context.Context, emp *empty.Empty) (*pb.BusinessSubCategoriesResponse, error) {

	subCategories, err := db.GetSubCategoriesRepository(ctx)
	if err != nil {
		return nil, err
	}

	return subCategories, nil
}

//GetBusinessSubCategoriesUnderCategory is
func (*Server) GetBusinessSubCategoriesUnderCategory(ctx context.Context, request *pb.BusinessSubCategoriesUnderCategoryRequest) (*pb.BusinessSubCategoriesUnderCategoryResponse, error) {
	categoryID := request.GetBusinessCategoryID()

	subCategories, err := db.GetSubCategoriesUnderCategoryRepository(ctx, categoryID)
	if err != nil {
		return nil, err
	}

	return subCategories, nil
}



