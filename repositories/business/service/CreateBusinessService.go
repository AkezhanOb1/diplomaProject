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
func  CreateBusinessService(ctx context.Context, serviceName string, subCategories []int64) (*pb.CreateBusinessServiceResponse, error) {
	conn, err := pgx.Connect(context.Background(), config.PostgresConnection)
	if err != nil {
		return nil, err
	}

	defer conn.Close(ctx)


	sqlQuery := `INSERT INTO business_service(name) VALUES ($1) RETURNING id;`

	var serviceID int64
	err = conn.QueryRow(context.Background(), sqlQuery, serviceName).Scan(&serviceID)
	if err != nil {
		return nil, err
	}

	for i := range subCategories {
		ctx = context.Background()
		tx, err := conn.Begin(ctx)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		sqlQuery = `INSERT INTO business_sub_categories_services(business_services_id, sub_categories_id)
		VALUES ($1,$2);`

		_, err = tx.Exec(ctx, sqlQuery, serviceID, subCategories[i])
		if err != nil {
			log.Println(err)
			return nil, err
		}

		err = tx.Commit(ctx)
		if err != nil {
			log.Println(err)
			return nil, err
		}

	}

	return &pb.CreateBusinessServiceResponse{
		BusinessService: &pb.BusinessService{
			BusinessServiceID:    serviceID,
			BusinessServiceName:  serviceName,
			SubCategories:        subCategories,
		},
	}, nil
}
