package company

import (
	"bytes"
	"context"
	config "github.com/AkezhanOb1/diplomaProject/configs"
	gq "github.com/AkezhanOb1/diplomaProject/api/graphQL/graph/model"
	pb "github.com/AkezhanOb1/diplomaProject/api/proto/business-company"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"log"

	"github.com/AkezhanOb1/diplomaProject/api/graphQL/graph/model"
)




//GetBusinessCompany is a client for graphQL on gRPC services
func GetBusinessCompany(ctx context.Context, companyID int64) (*pb.GetBusinessCompanyResponse, error) {
	cc, err := grpc.Dial(config.BusinessCompanyServer, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer cc.Close()

	c := pb.NewBusinessCompaniesServiceClient(cc)

	e := pb.GetBusinessCompanyRequest{
		BusinessCompanyID:   companyID,
	}

	company, err := c.GetBusinessCompany(ctx, &e)
	if err != nil {
		return nil, err
	}

	return company, nil
}


//GetBusinessCompanies is a client for graphQL on gRPC services
func GetBusinessCompanies(ctx context.Context) (*pb.GetBusinessCompaniesResponse, error) {
	cc, err := grpc.Dial(config.BusinessCompanyServer, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	defer  cc.Close()

	c := pb.NewBusinessCompaniesServiceClient(cc)
	e := empty.Empty{}

	companies, err := c.GetBusinessCompanies(ctx, &e)
	if err != nil {
		return nil, err
	}

	return companies, nil
}

//GetBusinessCompanyServices is a client for graphQL on gRPC services
func GetBusinessCompanyServices(ctx context.Context, companyID int64) (*pb.GetBusinessCompanyServicesResponse, error) {
	cc, err := grpc.Dial(config.BusinessCompanyServer, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	defer cc.Close()


	c := pb.NewBusinessCompaniesServiceClient(cc)
	e := pb.GetBusinessCompanyServicesRequest{
		BusinessCompanyID:   companyID,
	}


	services, err := c.GetBusinessCompanyServices(ctx, &e)
	if err != nil {
		return nil, err
	}

	return services, nil
}

//GetBusinessCompanyOperationHourByDay is a client for graphQL on gRPC services
func GetBusinessCompanyOperationHourByDay(ctx context.Context, companyID int64, dayOfWeek int64) (*pb.GetBusinessCompanyOperationHourByDayResponse, error) {
	cc, err := grpc.Dial(config.BusinessCompanyServer, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	defer cc.Close()


	c := pb.NewBusinessCompaniesServiceClient(cc)
	e := pb.GetBusinessCompanyOperationHourByDayRequest{
		CompanyID:            companyID,
		DayOfWeek:            dayOfWeek,
	}

	operationHour, err := c.GetBusinessCompanyOperationHourByDay(ctx, &e)
	if err != nil {
		return nil, err
	}


	return operationHour, nil
}

//GetBusinessCompanyOperationHours is a client for graphQL on gRPC services
func GetBusinessCompanyOperationHours(ctx context.Context, companyID int64) (*pb.GetBusinessCompanyOperationHoursResponse, error) {
	cc, err := grpc.Dial(config.BusinessCompanyServer, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	defer cc.Close()


	c := pb.NewBusinessCompaniesServiceClient(cc)
	e := pb.GetBusinessCompanyOperationHoursRequest{
		CompanyID:   companyID,
	}


	operationHours, err := c.GetBusinessCompanyOperationHours(ctx, &e)
	if err != nil {
		return nil, err
	}

	return operationHours, nil
}

//GetBusinessCompaniesUnderCategory is a client for graphQL on gRPC services
func GetBusinessCompaniesUnderCategory(ctx context.Context, categoryID int64) (*pb.GetBusinessCompaniesUnderCategoryResponse, error) {
	cc, err := grpc.Dial(config.BusinessCompanyServer, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	defer func(){
		err = cc.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	c := pb.NewBusinessCompaniesServiceClient(cc)
	e := pb.GetBusinessCompaniesUnderCategoryRequest{
		CategoryID:categoryID,
	}


	companies, err := c.GetBusinessCompaniesUnderCategory(ctx, &e)
	if err != nil {
		return nil, err
	}

	return companies, nil
}


//CreateBusinessCompany is a client function for registration a new business company
func CreateBusinessCompany(ctx context.Context, req gq.CreateBusinessCompanyRequest) (*pb.CreateBusinessCompanyResponse, error) {
	cc, err := grpc.Dial(config.BusinessCompanyServer, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	defer cc.Close()


	c := pb.NewBusinessCompaniesServiceClient(cc)

	r := pb.CreateBusinessCompanyRequest{
		BusinessCompanyName:       req.BusinessCompanyName,
		BusinessCompanyCategoryID: req.BusinessCompanyCategoryID,
		BusinessCompanyAddress: req.BusinessCompanyAddress,
	}

	company, err := c.CreateBusinessCompany(ctx, &r)
	if err != nil {
		return nil, err
	}

	return company, nil
}


//CreateBusinessCompanyOperationHour is a client function for registration a new business company
func CreateBusinessCompanyOperationHour(ctx context.Context, req gq.CreateBusinessCompanyOperationHoursRequest) (*pb.CreateBusinessCompanyOperationHourResponse, error) {
	cc, err := grpc.Dial(config.BusinessCompanyServer, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	defer cc.Close()


	c := pb.NewBusinessCompaniesServiceClient(cc)

	r := pb.CreateBusinessCompanyOperationHourRequest{
		BusinessCompanyID:    req.BusinessCompanyID,
		DayOfWeek:            req.DayOfWeek,
		OpenTime:             req.OpenTime,
		CloseTime:            req.CloseTime,
	}

	newOperationHour, err := c.CreateBusinessCompanyOperationHour(ctx, &r)
	if err != nil {
		return nil, err
	}

	return newOperationHour, nil
}


//CreateBusinessCompany is a client function for registration a new business company
func UpdateBusinessCompanyOperationHour(ctx context.Context, req gq.UpdateBusinessCompanyOperationHoursRequest) (*pb.UpdateBusinessCompanyOperationHourResponse, error) {
	cc, err := grpc.Dial(config.BusinessCompanyServer, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	defer cc.Close()


	c := pb.NewBusinessCompaniesServiceClient(cc)

	r := pb.UpdateBusinessCompanyOperationHourRequest{
		BusinessCompanyOperationHour: &pb.BusinessCompanyOperationHour{
			CompanyOperationHourID: req.CompanyOperationHourID,
			BusinessCompanyID:      req.BusinessCompanyID,
			DayOfWeek:              req.DayOfWeek,
			OpenTime:               req.OpenTime,
			CloseTime:              req.CloseTime,
		},
	}

	updatedOperationHour, err := c.UpdateBusinessCompanyOperationHour(ctx, &r)
	if err != nil {
		return nil, err
	}

	return updatedOperationHour, nil
}


//DeleteBusinessCompanyOperationHour is a client function for registration a new business company
func DeleteBusinessCompanyOperationHour(ctx context.Context, req gq.DeleteBusinessCompanyOperationHoursRequest) (*pb.DeleteBusinessCompanyOperationHourResponse, error) {
	cc, err := grpc.Dial(config.BusinessCompanyServer, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	defer cc.Close()


	c := pb.NewBusinessCompaniesServiceClient(cc)

	r := pb.DeleteBusinessCompanyOperationHourRequest{
		OperationHourID:      req.CompanyOperationHourID,
	}

	deletedOperationHour, err := c.DeleteBusinessCompanyOperationHour(ctx, &r)
	if err != nil {
		return nil, err
	}

	return deletedOperationHour, nil
}


func BusinessCompanyImageUpload(ctx context.Context, input model.BusinessCompanyImageUploadRequest) (*pb.BusinessCompanyImageUploadResponse, error) {
	cc, err := grpc.Dial(config.BusinessCompanyServer, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	defer cc.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(input.File.File)
	c := pb.NewBusinessCompaniesServiceClient(cc)

	r := pb.BusinessCompanyImageUploadRequest{
		BusinessCompanyID: input.BussinessCompanyID,
		Image:             &pb.File{
			File:        buf.Bytes(),
			FileName:    input.File.Filename,
			Size:        input.File.Size,
			ContentType: input.File.ContentType,
		},
	}

	uploadImage, err := c.BusinessCompanyImageUpload(ctx, &r)
	if err != nil {
		return nil, err
	}

	return uploadImage, nil
}