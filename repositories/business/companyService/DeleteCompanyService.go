package companyService

import (
	"context"
	pb "github.com/AkezhanOb1/diplomaProject/api/proto/business/companyServices"
	config "github.com/AkezhanOb1/diplomaProject/configs"
	"log"

	"github.com/jackc/pgx/v4"
)


//DeleteCompanyServiceRepository is a repository that responsible to inserting data into the
//company service table in database
func DeleteCompanyServiceRepository(ctx context.Context, companyServiceID int64) (*pb.DeleteCompanyServiceResponse, error) {
	conn, err := pgx.Connect(ctx, config.PostgresConnection)
	if err != nil {
		return nil, err
	}

	defer func() {
		err =  conn.Close(context.Background())
		if err != nil {
			log.Println(err)
		}
	}()

	sqlQuery := `DELETE FROM business_company_service WHERE id=$1 RETURNING *;`

	row := conn.QueryRow(ctx, sqlQuery, companyServiceID)

	var service pb.CompanyService

	err = row.Scan(
		&service.CompanyServiceID,
		&service.BusinessCompanyID,
		&service.BusinessServiceID,
		&service.CompanyServiceName,
		&service.CompanyServiceDuration,
		&service.CompanyServicePrice,
	)

	return &pb.DeleteCompanyServiceResponse{
		CompanyService: &service,
	}, nil
}

