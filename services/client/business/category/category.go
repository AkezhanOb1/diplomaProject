package main

import(
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	 pb "github.com/AkezhanOb1/diplomaProject/api/proto/business/categories"
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

	c := pb.NewBusinessCategoryServiceClient(cc)
	e := empty.Empty{
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}
	categories, err := c.GetBusinessCategories(context.Background(), &e)
	if err != nil {
		panic(err)
	}

	for _, b := range categories.GetBusinessCategories() {
		log.Println(b.GetBusinessCategoryID(), b.GetBusinessCategoryName())
	}
	log.Println(categories)
}
