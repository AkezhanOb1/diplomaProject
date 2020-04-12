package companyService

import (
	"context"
	"log"

	pb "github.com/AkezhanOb1/diplomaProject/api/proto/business/companyServices"
	config "github.com/AkezhanOb1/diplomaProject/configs"

	"github.com/jackc/pgx/v4"
)

//GetCompanyServiceRepository is a repository that responsible to all the requests to DB
//about business categories
func  GetCompanyServiceRepository(ctx context.Context, serviceID int64) (*pb.GetCompanyServiceResponse, error) {
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

	sqlQuery := `SELECT bc.id, bc.name, bc.duration, bc.price, b.id, b.name, s.id, s.name
					   FROM business_company_service bc
				 INNER JOIN business_company b ON bc.company_id = b.id
				 INNER JOIN business_service s ON bc.service_id = s.id
				 WHERE bc.id=$1;`

	row := conn.QueryRow(ctx, sqlQuery, serviceID)

	var service pb.CompanyService

	err = row.Scan(
		&service.CompanyServiceID,
		&service.CompanyServiceName,
		&service.CompanyServiceDuration,
		&service.CompanyServicePrice,
		&service.BusinessCompanyID,
		&service.BusinessCompanyName,
		&service.BusinessServiceID,
		&service.BusinessServiceName,
		)

	if err != nil {
		return nil, err
	}

	return &pb.GetCompanyServiceResponse{
		CompanyService: &service,
	}, nil
}
