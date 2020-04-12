package service

import (
	pb "github.com/AkezhanOb1/diplomaProject/api/proto/business/services"
	config "github.com/AkezhanOb1/diplomaProject/configs"
	"context"
	"log"
	"github.com/jackc/pgx/v4"

)

//GetBusinessServiceRepository is a repository that responsible to all the requests to DB
//about business categories
func  GetBusinessServiceRepository(ctx context.Context, serviceID int64) (*pb.GetBusinessServiceResponse, error) {
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
	             WHERE b.id = $1 GROUP BY b.id;`

	row := conn.QueryRow(ctx, sqlQuery, serviceID)

	var service pb.BusinessService

	err = row.Scan(
		&service.BusinessServiceID,
		&service.BusinessServiceName,
		&service.SubCategories,
	)

	if err != nil {
		return nil, err
	}

	return &pb.GetBusinessServiceResponse{
		BusinessService: &service,
	}, nil
}
