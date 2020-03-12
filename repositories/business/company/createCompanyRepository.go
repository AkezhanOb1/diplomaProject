package company

import (
	"context"
	"log"

	pb "github.com/AkezhanOb1/diplomaProject/api/proto/business/companies"
	config "github.com/AkezhanOb1/diplomaProject/configs"

	"github.com/jackc/pgx/v4"
)

//CreateCompanyRepository is a repository that
//responsible to inserting data into the company
//table in database
func CreateCompanyRepository(ctx context.Context, request *pb.CreateBusinessCompanyRequest) (*pb.CreateBusinessCompanyResponse, error) {
	conn, err := pgx.Connect(ctx, config.PostgresConnection)
	if err != nil {
		return nil, err
	}

	defer func() {
		err =  conn.Close(context.Background())
		if err != nil {
			log.Fatal(err)
		}
	}()

	sqlQuery := `INSERT INTO public.business_company (
				 	 name, category_id)
				 VALUES ($1, $2) RETURNING id;`


	var companyID int64
	companyName := request.GetBusinessCompanyName()
	companyCategoryID := request.GetBusinessCompanyCategoryID()

	row := conn.QueryRow(context.Background(), sqlQuery, companyName, companyCategoryID)


	err = row.Scan(&companyID)
	if err != nil {
		return nil, err
	}


	return &pb.CreateBusinessCompanyResponse{
		BusinessCompany: &pb.BusinessCompany{
			BusinessCompanyID:         companyID,
			BusinessCompanyName:       companyName,
			BusinessCompanyCategoryID: companyCategoryID,
		},
	}, nil
}
