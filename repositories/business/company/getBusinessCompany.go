package company

import (
	pb "github.com/AkezhanOb1/diplomaProject/api/proto/business/companies"
	config "github.com/AkezhanOb1/diplomaProject/configs"
	"context"
	"log"
	"github.com/jackc/pgx/v4"

)

//GetBusinessCompanyRepository is a repository that responsible to all the requests to DB
//about business categories
func  GetBusinessCompanyRepository(ctx context.Context, companyID int64) (*pb.GetBusinessCompanyResponse, error) {
	conn, err := pgx.Connect(ctx, config.PostgresConnection)
	if err != nil {
		return nil, err
	}

	defer func() {
		err =  conn.Close(ctx)
		if err != nil {
			log.Println(err)
		}
	}()

	sqlQuery := `SELECT id, name, category_id FROM business_company WHERE id=$1;`

	row := conn.QueryRow(ctx, sqlQuery, companyID)

	var company pb.BusinessCompany

	err = row.Scan(
		&company.BusinessCompanyID,
		&company.BusinessCompanyName,
		&company.BusinessCompanyCategoryID,
	)

	if err != nil {
		return nil, err
	}

	return &pb.GetBusinessCompanyResponse{
		BusinessCompany: &company,
	}, nil
}
