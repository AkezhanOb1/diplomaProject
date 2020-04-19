package company

import (
	"context"
	"log"

	pb "github.com/AkezhanOb1/diplomaProject/api/proto/business/companies"
	config "github.com/AkezhanOb1/diplomaProject/configs"

	"github.com/jackc/pgx/v4"
)

//DeleteBusinessCompanyOperationHourRepository is a repository that responsible for inserting data into the company
//table inside the database
func DeleteBusinessCompanyOperationHourRepository(ctx context.Context, operationHourID int64) (*pb.DeleteBusinessCompanyOperationHourResponse, error) {
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


	sqlQuery := `DELETE FROM business_company_operation_hours WHERE id=$1 RETURNING *;`

	row := conn.QueryRow(ctx, sqlQuery, operationHourID)

	var operationHour pb.BusinessCompanyOperationHour

	err = row.Scan(
		&operationHour.CompanyOperationHourID,
		&operationHour.BusinessCompanyID,
		&operationHour.DayOfWeek,
		&operationHour.OpenTime,
		&operationHour.CloseTime,
	)


	return &pb.DeleteBusinessCompanyOperationHourResponse{
		BusinessCompanyOperationHour: &operationHour,
	}, nil
}
