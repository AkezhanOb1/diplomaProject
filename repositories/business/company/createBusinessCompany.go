package company

import (
	"context"
	pb "github.com/AkezhanOb1/diplomaProject/api/proto/business/companies"
	config "github.com/AkezhanOb1/diplomaProject/configs"

	"github.com/jackc/pgx/v4"
)

//CreateCompanyRepository is a repository that responsible for inserting data into the company
//table inside the database
func CreateCompanyRepository(ctx context.Context, request *pb.CreateBusinessCompanyRequest) (*pb.CreateBusinessCompanyResponse, error) {
	conn, err := pgx.Connect(ctx, config.PostgresConnection)
	if err != nil {
		return nil, err
	}

	defer conn.Close(context.Background())

	sqlQuery := `INSERT INTO public.business_company (
				 	 name, category_id, address)
				 VALUES ($1, $2, $3) RETURNING id;`


	var companyID int64
	companyName := request.GetBusinessCompanyName()
	companyCategoryID := request.GetBusinessCompanyCategoryID()
	companyAddress := request.GetBusinessCompanyAddress()

	row := conn.QueryRow(context.Background(), sqlQuery, companyName, companyCategoryID, companyAddress)


	err = row.Scan(&companyID)
	if err != nil {
		return nil, err
	}


	return &pb.CreateBusinessCompanyResponse{
		BusinessCompany: &pb.BusinessCompany{
			BusinessCompanyID:         companyID,
			BusinessCompanyName:       companyName,
			BusinessCompanyCategoryID: companyCategoryID,
			BusinessCompanyAddress: companyAddress,
		},
	}, nil
}
