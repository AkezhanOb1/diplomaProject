package companyService

import (
	"context"
	pb "github.com/AkezhanOb1/diplomaProject/api/proto/business/companyServices"
	db "github.com/AkezhanOb1/diplomaProject/repositories/business/companyService"
	"github.com/golang/protobuf/ptypes/empty"
)


type Server struct {}



//GetBusinessCompanyServiceOperationHourByDay is
func (*Server) GetBusinessCompanyServiceOperationHourByDay(ctx context.Context, request *pb.GetBusinessCompanyServiceOperationHourByDayRequest) (*pb.GetBusinessCompanyServiceOperationHourByDayResponse, error) {
	dayOfWeek := request.GetDayOfWeek()
	serviceID := request.GetServiceID()
	operationHour, err := db.GetBusinessCompanyServiceOperationHourByDayRepository(ctx, serviceID, dayOfWeek)
	if err != nil {
		return nil, err
	}

	return operationHour, nil
}

//GetBusinessCompanyServiceOperationHours is
func (*Server) GetBusinessCompanyServiceOperationHours(ctx context.Context, request *pb.GetBusinessCompanyServiceOperationHoursRequest) (*pb.GetBusinessCompanyServiceOperationHoursResponse, error) {
	serviceID := request.GetServiceID()
	operationHours, err := db.GetBusinessCompanyServiceOperationHoursRepository(ctx, serviceID)
	if err != nil {
		return nil, err
	}

	return operationHours, nil
}

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


//CreateBusinessCompanyServiceOperationHour is
func (*Server) CreateBusinessCompanyServiceOperationHour(ctx context.Context, request *pb.CreateBusinessCompanyServiceOperationHourRequest) (*pb.CreateBusinessCompanyServiceOperationHourResponse, error) {
	businessCompanyService, err := db.CreateBusinessCompanyServiceOperationHourRepository(ctx, request)
	if err != nil {
		return nil, err
	}

	return businessCompanyService, nil
}



//UpdateBusinessCompanyServiceOperationHour is
func (*Server) UpdateBusinessCompanyServiceOperationHour(ctx context.Context, request *pb.UpdateBusinessCompanyServiceOperationHourRequest) (*pb.UpdateBusinessCompanyServiceOperationHourResponse, error) {
	updatedOperationHour, err := db.UpdateBusinessCompanyServiceOperationHourRepository(ctx, request.BusinessCompanyServiceOperationHour)
	if err != nil {
		return nil, err
	}

	return updatedOperationHour, nil
}


//DeleteBusinessCompanyServiceOperationHour is
func (*Server) DeleteBusinessCompanyServiceOperationHour(ctx context.Context, request *pb.DeleteBusinessCompanyServiceOperationHourRequest) (*pb.DeleteBusinessCompanyServiceOperationHourResponse, error) {
	var operationHourID = request.OperationHourID
	deletedOperationHour, err := db.DeleteBusinessCompanyServiceOperationHourRepository(ctx, operationHourID)
	if err != nil {
		return nil, err
	}

	return deletedOperationHour, nil
}

