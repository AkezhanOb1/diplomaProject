package companyService

import (
	"context"
	pb "github.com/AkezhanOb1/diplomaProject/api/proto/business/companyServices"
	config "github.com/AkezhanOb1/diplomaProject/configs"

	"github.com/jackc/pgx/v4"
)

//GetCompanyServicesUnderCategoryRepository is a repository that responsible to all the requests to DB
//about business categories
func GetCompanyServicesUnderCategoryRepository(ctx context.Context, categoryID int64) (*pb.GetCompanyServicesUnderCategoryResponse, error) {
	conn, err := pgx.Connect(ctx, config.PostgresConnection)
	if err != nil {
		return nil, err
	}

	defer conn.Close(ctx)


	sqlQuery := `SELECT bc.id, bc.name, bc.duration, bc.price, b.id, b.name, s.id, s.name
					FROM business_company_service bc
				INNER JOIN business_company b
				 ON bc.company_id = b.id
				INNER JOIN business_service s
				 ON bc.service_id = s.id
				INNER JOIN business_sub_categories_services bscs
				 ON s.id = bscs.business_services_id
				INNER JOIN business_sub_category bsc
				 on bscs.sub_categories_id = bsc.id
				WHERE bsc.category_id=$1;`

	rows, err := conn.Query(ctx, sqlQuery, categoryID)
	if err != nil {
		return nil, err
	}

	var services pb.GetCompanyServicesUnderCategoryResponse

	for rows.Next() {
		var service pb.CompanyService

		err = rows.Scan(
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

		services.CompanyServices = append(services.CompanyServices, &service)
	}

	return &services, nil
}
