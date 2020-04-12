package companyService

import (
	"context"
	pb "github.com/AkezhanOb1/diplomaProject/api/proto/business/companyServices"
	db "github.com/AkezhanOb1/diplomaProject/repositories/business/companyService"
	"github.com/golang/protobuf/ptypes/empty"
)


type Server struct {}

//GetCompanyService is
func (*Server) GetCompanyService(ctx context.Context, request *pb.GetCompanyServiceRequest) (*pb.GetCompanyServiceResponse, error) {
	serviceID := request.GetCompanyServiceID()

	service, err := db.GetCompanyServiceRepository(ctx, serviceID)
	if err != nil {
		return nil, err
	}

	return service, nil
}

//GetCompanyServices is
func (*Server) GetCompanyServices(ctx context.Context, emp *empty.Empty) (*pb.GetCompanyServicesResponse, error) {

	services, err := db.GetCompanyServicesRepository(ctx)
	if err != nil {
		return nil, err
	}

	return services, nil
}

//GetCompanyServicesUnderSubCategory is
func (*Server) GetCompanyServicesUnderSubCategory(ctx context.Context, request *pb.GetCompanyServicesUnderSubCategoryRequest) (*pb.GetCompanyServicesUnderSubCategoryResponse, error) {
	subCategoryID := request.GetSubCategoryID()

	services, err := db.GetCompanyServicesUnderSubCategoryRepository(ctx, subCategoryID)
	if err != nil {
		return nil, err
	}

	return services, nil
}


//CreateCompanyService is
func (*Server) CreateCompanyService(ctx context.Context, request *pb.CreateCompanyServiceRequest) (*pb.CreateCompanyServiceResponse, error) {
	newService, err := db.CreateCompanyServiceRepository(ctx, request)
	if err != nil {
		return nil, err
	}

	return newService, nil
}

//UpdateCompanyService is
func (*Server) UpdateCompanyService(ctx context.Context, request *pb.UpdateCompanyServiceRequest) (*pb.UpdateCompanyServiceResponse, error) {
	updatedService, err := db.UpdateCompanyServiceRepository(ctx, request)
	if err != nil {
		return nil, err
	}

	return updatedService, nil
}

//DeleteCompanyService is
func (*Server) DeleteCompanyService(ctx context.Context, request *pb.DeleteCompanyServiceRequest) (*pb.DeleteCompanyServiceResponse, error) {
	var companyServiceID = request.GetCompanyServiceID()

	service, err := db.DeleteCompanyServiceRepository(ctx,companyServiceID)
	if err != nil {
		return nil, err
	}

	return service, nil
}



