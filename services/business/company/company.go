package company

import (
	"context"
	pb "github.com/AkezhanOb1/diplomaProject/api/proto/business/companies"
	db "github.com/AkezhanOb1/diplomaProject/repositories/business/company"
	"github.com/golang/protobuf/ptypes/empty"
)

type Server struct {}


//GetBusinessCompany is
func (*Server) GetBusinessCompany(ctx context.Context, request *pb.GetBusinessCompanyRequest) (*pb.GetBusinessCompanyResponse, error) {
	businessCompany, err := db.GetBusinessCompanyRepository(ctx, request.BusinessCompanyID)
	if err != nil {
		return nil, err
	}

	return businessCompany, nil
}

//GetBusinessCompanies is
func (*Server) GetBusinessCompanies(ctx context.Context, emp *empty.Empty) (*pb.GetBusinessCompaniesResponse, error) {
	businessCompanies, err := db.GetBusinessCompaniesRepository(ctx)
	if err != nil {
		return nil, err
	}

	return businessCompanies, nil
}

//GetBusinessCompanyServices is
func (*Server) GetBusinessCompanyServices(ctx context.Context, request *pb.GetBusinessCompanyServicesRequest) (*pb.GetBusinessCompanyServicesResponse, error) {
	businessCompanyID := request.GetBusinessCompanyID()
	businessCompany, err := db.GetBusinessCompanyServicesRepository(ctx, businessCompanyID)
	if err != nil {
		return nil, err
	}

	return businessCompany, nil
}

//GetBusinessCompanyOperationHourByDay is
func (*Server) GetBusinessCompanyOperationHourByDay(ctx context.Context, request *pb.GetBusinessCompanyOperationHourByDayRequest) (*pb.GetBusinessCompanyOperationHourByDayResponse, error) {
	dayOfWeek := request.GetDayOfWeek()
	companyID := request.GetCompanyID()
	operationHour, err := db.GetBusinessCompanyOperationHourByDayRepository(ctx, companyID, dayOfWeek)
	if err != nil {
		return nil, err
	}

	return operationHour, nil
}

//GetBusinessCompanyOperationHours is
func (*Server) GetBusinessCompanyOperationHours(ctx context.Context, request *pb.GetBusinessCompanyOperationHoursRequest) (*pb.GetBusinessCompanyOperationHoursResponse, error) {
	companyID := request.GetCompanyID()
	operationHours, err := db.GetBusinessCompanyOperationHoursRepository(ctx, companyID)
	if err != nil {
		return nil, err
	}

	return operationHours, nil
}

//GetBusinessCompaniesUnderCategory is
func (*Server) GetBusinessCompaniesUnderCategory(ctx context.Context, request *pb.GetBusinessCompaniesUnderCategoryRequest) (*pb.GetBusinessCompaniesUnderCategoryResponse, error) {
	companies, err := db.GetBusinessCompaniesUnderCategoryRepository(ctx, request.GetCategoryID())
	if err != nil {
		return nil, err
	}

	return companies, nil
}

//CreateBusinessCompany is
func (*Server) CreateBusinessCompany(ctx context.Context, request *pb.CreateBusinessCompanyRequest) (*pb.CreateBusinessCompanyResponse, error) {
	businessCompany, err := db.CreateCompanyRepository(ctx, request)
	if err != nil {
		return nil, err
	}

	return businessCompany, nil
}


//CreateBusinessCompanyOperationHour is
func (*Server) CreateBusinessCompanyOperationHour(ctx context.Context, request *pb.CreateBusinessCompanyOperationHourRequest) (*pb.CreateBusinessCompanyOperationHourResponse, error) {
	businessCompany, err := db.CreateBusinessCompanyOperationHourRepository(ctx, request)
	if err != nil {
		return nil, err
	}

	return businessCompany, nil
}



//CreateBusinessCompanyOperationHour is
func (*Server) UpdateBusinessCompanyOperationHour(ctx context.Context, request *pb.UpdateBusinessCompanyOperationHourRequest) (*pb.UpdateBusinessCompanyOperationHourResponse, error) {
	updatedOperationHour, err := db.UpdateBusinessCompanyOperationHourRepository(ctx, request.BusinessCompanyOperationHour)
	if err != nil {
		return nil, err
	}

	return updatedOperationHour, nil
}



//CreateBusinessCompanyOperationHour is
func (*Server) DeleteBusinessCompanyOperationHour(ctx context.Context, request *pb.DeleteBusinessCompanyOperationHourRequest) (*pb.DeleteBusinessCompanyOperationHourResponse, error) {
	var operationHourID = request.OperationHourID
	deletedOperationHour, err := db.DeleteBusinessCompanyOperationHourRepository(ctx, operationHourID)
	if err != nil {
		return nil, err
	}

	return deletedOperationHour, nil
}

