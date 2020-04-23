package companyService

import (
	pb "github.com/AkezhanOb1/diplomaProject/api/proto/business/companyServices"
	config "github.com/AkezhanOb1/diplomaProject/configs"
	gm "github.com/AkezhanOb1/diplomaProject/api/graphQL/graph/model"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"log"
	"context"
)


//GetCompanyService is a client for graphQL on gRPC services
func GetCompanyService(ctx context.Context, id int64) (*pb.GetCompanyServiceResponse, error) {

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

	c := pb.NewCompanyServicesClient(cc)

	e := pb.GetCompanyServiceRequest{
		CompanyServiceID:id,
	}

	service, err := c.GetCompanyService(ctx, &e)
	if err != nil {
		return nil, err
	}

	return service, nil
}


//GetCompanyServices is a client for graphQL on gRPC services
func GetCompanyServices(ctx context.Context) (*pb.GetCompanyServicesResponse, error) {
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

	c := pb.NewCompanyServicesClient(cc)
	e := empty.Empty{}

	services, err := c.GetCompanyServices(ctx, &e)
	if err != nil {
		return nil, err
	}

	return services, nil
}


//GetCompanyServicesUnderSubCategory is a client for graphQL on gRPC services
func GetCompanyServicesUnderSubCategory(ctx context.Context, subCategoryID int64) (*pb.GetCompanyServicesUnderSubCategoryResponse, error) {
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

	c := pb.NewCompanyServicesClient(cc)
	e := pb.GetCompanyServicesUnderSubCategoryRequest{
		SubCategoryID:   subCategoryID,
	}

	services, err := c.GetCompanyServicesUnderSubCategory(ctx, &e)
	if err != nil {
		return nil, err
	}

	return services, nil
}



//GetBusinessCompanyOperationHourByDay is a client for graphQL on gRPC services
func GetBusinessCompanyServiceOperationHourByDay(ctx context.Context, serviceID int64, dayOfWeek int64) (*pb.GetBusinessCompanyServiceOperationHourByDayResponse, error) {
	cc, err := grpc.Dial(config.RpcServerAddress, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	defer cc.Close()


	c := pb.NewCompanyServicesClient(cc)
	r := pb.GetBusinessCompanyServiceOperationHourByDayRequest{
		ServiceID:            serviceID,
		DayOfWeek:            dayOfWeek,
	}

	operationHour, err := c.GetBusinessCompanyServiceOperationHourByDay(ctx, &r)
	if err != nil {
		return nil, err
	}


	return operationHour, nil
}

//GetBusinessCompanyOperationHours is a client for graphQL on gRPC services
func GetBusinessCompanyServiceOperationHours(ctx context.Context, serviceID int64) (*pb.GetBusinessCompanyServiceOperationHoursResponse, error) {
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

	c := pb.NewCompanyServicesClient(cc)
	r := pb.GetBusinessCompanyServiceOperationHoursRequest{
		ServiceID:   serviceID,
	}


	operationHours, err := c.GetBusinessCompanyServiceOperationHours(ctx, &r)
	if err != nil {
		return nil, err
	}

	return operationHours, nil
}



//CreateBusinessService is a client for graphQL on gRPC services
func CreateBusinessService(ctx context.Context, request gm.CreateCompanyServiceRequest) (*pb.CreateCompanyServiceResponse, error) {
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

	c := pb.NewCompanyServicesClient(cc)
	e := pb.CreateCompanyServiceRequest{
		CompanyServiceName:     request.CompanyServiceName,
		CompanyServiceDuration: request.CompanyServiceDuration,
		CompanyServicePrice:    request.CompanyServicePrice,
		BusinessServiceID:      request.BusinessServiceID,
		BusinessCompanyID:      request.BusinessCompanyID,
	}

	services, err := c.CreateCompanyService(ctx, &e)
	if err != nil {
		return nil, err
	}

	return services, nil
}



//UpdateCompanyService is a client for graphQL on gRPC services
func UpdateCompanyService(ctx context.Context, request gm.UpdateCompanyServiceRequest) (*pb.UpdateCompanyServiceResponse, error) {
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

	c := pb.NewCompanyServicesClient(cc)
	e := pb.UpdateCompanyServiceRequest{
		CompanyServiceID: 	    request.CompanyServiceID,
		CompanyServiceName:     request.CompanyServiceName,
		CompanyServiceDuration: request.CompanyServiceDuration,
		CompanyServicePrice:    request.CompanyServicePrice,
		BusinessServiceID:      request.BusinessServiceID,
		BusinessCompanyID:      request.BusinessCompanyID,
	}

	updatedService, err := c.UpdateCompanyService(ctx, &e)
	if err != nil {
		return nil, err
	}

	return updatedService, nil
}



//DeleteCompanyService is a client for graphQL on gRPC services
func DeleteCompanyService(ctx context.Context, request gm.DeleteCompanyServiceRequest) (*pb.DeleteCompanyServiceResponse, error) {
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

	c := pb.NewCompanyServicesClient(cc)
	e := pb.DeleteCompanyServiceRequest{
		CompanyServiceID: 	   request.CompanyServiceID,
	}

	deletedService, err := c.DeleteCompanyService(ctx, &e)
	if err != nil {
		return nil, err
	}

	return deletedService, nil
}



//CreateBusinessCompanyServiceOperationHours is a client function for registration a new business company
func CreateBusinessCompanyServiceOperationHour(ctx context.Context, req gm.CreateBusinessCompanyServiceOperationHoursRequest) (*pb.CreateBusinessCompanyServiceOperationHourResponse, error) {
	cc, err := grpc.Dial(config.RpcServerAddress, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	defer cc.Close()

	c := pb.NewCompanyServicesClient(cc)

	r := pb.CreateBusinessCompanyServiceOperationHourRequest{
		BusinessCompanyID:    req.BusinessCompanyID,
		BusinessServiceID:    req.BusinessServiceID,
		DayOfWeek:            req.DayOfWeek,
		OpenTime:             req.OpenTime,
		CloseTime:            req.CloseTime,
	}

	newOperationHour, err := c.CreateBusinessCompanyServiceOperationHour(ctx, &r)
	if err != nil {
		return nil, err
	}

	return newOperationHour, nil
}


//UpdateBusinessCompanyServiceOperationHour is a client function for registration a new business company
func UpdateBusinessCompanyServiceOperationHour(ctx context.Context, req gm.UpdateBusinessCompanyServiceOperationHoursRequest) (*pb.UpdateBusinessCompanyServiceOperationHourResponse, error) {
	cc, err := grpc.Dial(config.RpcServerAddress, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	defer cc.Close()


	c := pb.NewCompanyServicesClient(cc)

	r := pb.UpdateBusinessCompanyServiceOperationHourRequest{
		BusinessCompanyServiceOperationHour: &pb.BusinessCompanyServiceOperationHour{
			ServiceOperationHourID: req.OperationHourID,
			BusinessCompanyID:      req.BusinessCompanyID,
			BusinessServiceID:      req.BusinessServiceID,
			DayOfWeek:              req.DayOfWeek,
			OpenTime:               req.OpenTime,
			CloseTime:              req.CloseTime,
		},
	}

	updatedOperationHour, err := c.UpdateBusinessCompanyServiceOperationHour(ctx, &r)
	if err != nil {
		return nil, err
	}

	return updatedOperationHour, nil
}


//DeleteBusinessCompanyServiceOperationHour is a client function for registration a new business company
func DeleteBusinessCompanyServiceOperationHour(ctx context.Context, req gm.DeleteBusinessCompanyServiceOperationHoursRequest) (*pb.DeleteBusinessCompanyServiceOperationHourResponse, error) {
	cc, err := grpc.Dial(config.RpcServerAddress, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	defer cc.Close()


	c := pb.NewCompanyServicesClient(cc)

	r := pb.DeleteBusinessCompanyServiceOperationHourRequest{
		OperationHourID:      req.OperationHourID,
	}

	deletedOperationHour, err := c.DeleteBusinessCompanyServiceOperationHour(ctx, &r)
	if err != nil {
		return nil, err
	}

	return deletedOperationHour, nil
}









