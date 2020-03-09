package main

import(
	"context"
	"google.golang.org/grpc"
	pb "github.com/AkezhanOb1/diplomaProject/api/proto/business/companies"
	"log"
)

func main(){
	cc, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer func(){
		err := cc.Close()
		if err != nil {
			panic(err)
		}
	}()

	c := pb.NewBusinessCompaniesServiceClient(cc)

	req := pb.CreateBusinessCompanyRequest{
		BusinessCompanyName:       "Cactus",
		BusinessCompanyCategoryID: 1,
	}

	company, err := c.CreateBusinessCompany(context.Background(), &req)
	if err != nil {
		panic(err)
	}

	log.Println(company)
}
