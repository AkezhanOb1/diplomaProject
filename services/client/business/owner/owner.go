package main

import(
	"context"
	"google.golang.org/grpc"
	pb "github.com/AkezhanOb1/diplomaProject/api/proto/business/owners"
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

	req := pb.CreateBusinessOwnerRequest{
		BusinessOwnerName:              "Akezhan",
		BusinessOwnerEmail:             "esbolatovakEzhan@gmail.com",
		BusinessOwnerPassword:          "Yfehsp2203",
		BusinessOwnerPhoneNumberPrefix: "+7",
		BusinessOwnerPhoneNumber:       "7772059339",
		BusinessCompany_ID:             5,
	}

	company, err := c.CreateBusinessOwner(context.Background(), &req)
	if err != nil {
		panic(err)
	}

	log.Println(company)
}
