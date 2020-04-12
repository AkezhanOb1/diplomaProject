package service

import (
	pb "github.com/AkezhanOb1/diplomaProject/api/proto/business/services"
	config "github.com/AkezhanOb1/diplomaProject/configs"
	"log"
	"github.com/jackc/pgx/v4"
	"context"

)


//GetServicesUnderSubCategoryRepository is a repository that responsible to all the requests to DB
//about business categories
func GetServicesUnderSubCategoryRepository(ctx context.Context, categoryID int64) (*pb.GetServicesUnderSubCategoryResponse, error) {
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

	sqlQuery := `SELECT b.id, b.name, array_agg(bsc.id) FROM business_service b
				 INNER JOIN business_sub_categories_services
					ON b.id = business_sub_categories_services.business_services_id
				 INNER JOIN business_sub_category bsc
					ON business_sub_categories_services.sub_categories_id = bsc.id
	             WHERE bsc.id=$1 GROUP BY b.id;`

	rows, err := conn.Query(ctx, sqlQuery, categoryID)
	if err != nil {
		return nil, err
	}

	var services pb.GetServicesUnderSubCategoryResponse

	for rows.Next() {
		var service pb.BusinessService

		err = rows.Scan(
			&service.BusinessServiceID,
			&service.BusinessServiceName,
			&service.SubCategories,
		)

		if err != nil {
			return nil, err
		}

		services.BusinessServices = append(services.BusinessServices, &service)
	}

	return &services, nil
}
