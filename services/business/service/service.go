package service

import (
	"context"
	pb "github.com/AkezhanOb1/diplomaProject/api/proto/business/services"
	db "github.com/AkezhanOb1/diplomaProject/repositories/business/service"
	"github.com/golang/protobuf/ptypes/empty"
)


type Server struct {}

//GetBusinessService is
func (*Server) GetBusinessService(ctx context.Context, request *pb.GetBusinessServiceRequest) (*pb.GetBusinessServiceResponse, error) {
	serviceID := request.GetBusinessServiceID()

	service, err := db.GetBusinessServiceRepository(ctx, serviceID)
	if err != nil {
		return nil, err
	}

	return service, nil
}

//GetBusinessServices is
func (*Server) GetBusinessServices(ctx context.Context, emp *empty.Empty) (*pb.GetBusinessServicesResponse, error) {

	services, err := db.GetBusinessServicesRepository(ctx)
	if err != nil {
		return nil, err
	}

	return services, nil
}

//GetServicesUnderSubCategory is
func (*Server) GetServicesUnderSubCategory(ctx context.Context, request *pb.GetServicesUnderSubCategoryRequest) (*pb.GetServicesUnderSubCategoryResponse, error) {
	categoryID := request.GetSubCategoryID()

	services, err := db.GetServicesUnderSubCategoryRepository(ctx, categoryID)
	if err != nil {
		return nil, err
	}

	return services, nil
}


//CreateBusinessService is
func (*Server) CreateBusinessService(ctx context.Context, request *pb.CreateBusinessServiceRequest) (*pb.CreateBusinessServiceResponse, error) {

	var serviceName = request.GetBusinessServiceName()
	var subCategories = request.GetSubCategories()

	service, err := db.CreateBusinessService(ctx, serviceName, subCategories)
	if err != nil {
		return nil, err
	}

	return service, nil
}



