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
		CompanyServiceID:   id,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
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
	e := empty.Empty{
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}

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
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}

	services, err := c.GetCompanyServicesUnderSubCategory(ctx, &e)
	if err != nil {
		return nil, err
	}

	return services, nil
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
		XXX_NoUnkeyedLiteral:   struct{}{},
		XXX_unrecognized:       nil,
		XXX_sizecache:          0,
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
		XXX_NoUnkeyedLiteral:   struct{}{},
		XXX_unrecognized:       nil,
		XXX_sizecache:          0,
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
		XXX_NoUnkeyedLiteral:   struct{}{},
		XXX_unrecognized:       nil,
		XXX_sizecache:          0,
	}

	deletedService, err := c.DeleteCompanyService(ctx, &e)
	if err != nil {
		return nil, err
	}

	return deletedService, nil
}














