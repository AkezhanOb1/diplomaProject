package company

import (
	"context"
	gq "github.com/AkezhanOb1/diplomaProject/api/graphQL/graph/model"
	pb "github.com/AkezhanOb1/diplomaProject/api/proto/business/companies"
	"google.golang.org/grpc"
	"log"
)

//CreateBusinessCompany is a client function for registration a new
//business company
func CreateBusinessCompany(ctx context.Context, req gq.CreateBusinessCompanyRequest) (*pb.CreateBusinessCompanyResponse, error) {
	cc, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())
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
