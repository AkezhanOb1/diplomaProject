package company

import (
	pb "github.com/AkezhanOb1/diplomaProject/api/proto/business/companies"
	config "github.com/AkezhanOb1/diplomaProject/configs"
	"log"
	"github.com/jackc/pgx/v4"
	"context"

)


//GetBusinessCompanyServicesRepository is a repository that responsible to all the requests to DB
//about business categories
func GetBusinessCompanyServicesRepository(ctx context.Context, companyID int64) (*pb.GetBusinessCompanyServicesResponse, error) {
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
	sqlQuery := `SELECT id, name, duration, price FROM business_company_service 
				    WHERE company_id=$1;`

	rows, err := conn.Query(ctx, sqlQuery, companyID)
	if err != nil {
		return nil, err
	}

	var services pb.GetBusinessCompanyServicesResponse

	for rows.Next() {
		var service pb.BusinessCompanyService

		err = rows.Scan(
			&service.CompanyServiceID,
			&service.CompanyServiceName,
			&service.CompanyServiceDuration,
			&service.CompanyServicePrice,
		)

		if err != nil {
			return nil, err
		}

		services.BusinessCompanyService = append(services.BusinessCompanyService, &service)
	}

	return &services, nil
}
