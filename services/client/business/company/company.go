package company

import (
	"context"
	config "github.com/AkezhanOb1/diplomaProject/configs"
	gq "github.com/AkezhanOb1/diplomaProject/api/graphQL/graph/model"
	pb "github.com/AkezhanOb1/diplomaProject/api/proto/business/companies"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"log"
)




//GetBusinessCompany is a client for graphQL on gRPC services
func GetBusinessCompany(ctx context.Context, companyID int64) (*pb.GetBusinessCompanyResponse, error) {
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

	c := pb.NewBusinessCompaniesServiceClient(cc)

	e := pb.GetBusinessCompanyRequest{
		BusinessCompanyID:   companyID,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}

	company, err := c.GetBusinessCompany(ctx, &e)
	if err != nil {
		return nil, err
	}

	return company, nil
}


//GetBusinessCompanies is a client for graphQL on gRPC services
func GetBusinessCompanies(ctx context.Context) (*pb.GetBusinessCompaniesResponse, error) {
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

	c := pb.NewBusinessCompaniesServiceClient(cc)
	e := empty.Empty{
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}

	companies, err := c.GetBusinessCompanies(ctx, &e)
	if err != nil {
		return nil, err
	}

	return companies, nil
}

//GetBusinessCompanyServices is a client for graphQL on gRPC services
func GetBusinessCompanyServices(ctx context.Context, companyID int64) (*pb.GetBusinessCompanyServicesResponse, error) {
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

	c := pb.NewBusinessCompaniesServiceClient(cc)
	e := pb.GetBusinessCompanyServicesRequest{
		BusinessCompanyID:   companyID,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}


	services, err := c.GetBusinessCompanyServices(ctx, &e)
	if err != nil {
		return nil, err
	}

	return services, nil
}



//CreateBusinessCompany is a client function for registration a new business company
func CreateBusinessCompany(ctx context.Context, req gq.CreateBusinessCompanyRequest) (*pb.CreateBusinessCompanyResponse, error) {
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

	c := pb.NewBusinessCompaniesServiceClient(cc)

	r := pb.CreateBusinessCompanyRequest{
		BusinessCompanyName:       req.BusinessCompanyName,
		BusinessCompanyCategoryID: req.BusinessCompanyCategoryID,
	}

	company, err := c.CreateBusinessCompany(ctx, &r)
	if err != nil {
		return nil, err
	}

	return company, nil
}
