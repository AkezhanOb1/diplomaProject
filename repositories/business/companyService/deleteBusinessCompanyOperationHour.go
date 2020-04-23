package companyService

import (
	"context"
	pb "github.com/AkezhanOb1/diplomaProject/api/proto/business/companyServices"
	config "github.com/AkezhanOb1/diplomaProject/configs"

	"github.com/jackc/pgx/v4"
)

//DeleteBusinessCompanyServiceOperationHourRepository is a repository that responsible for inserting data into the company
//table inside the database
func DeleteBusinessCompanyServiceOperationHourRepository(ctx context.Context, operationHourID int64) (*pb.DeleteBusinessCompanyServiceOperationHourResponse, error) {
	conn, err := pgx.Connect(ctx, config.PostgresConnection)
	if err != nil {
		return nil, err
	}

	defer conn.Close(ctx)



	sqlQuery := `DELETE FROM business_company_service_operation_hours WHERE id=$1 RETURNING *;`

	row := conn.QueryRow(ctx, sqlQuery, operationHourID)

	var operationHour pb.BusinessCompanyServiceOperationHour

	err = row.Scan(
		&operationHour.ServiceOperationHourID,
		&operationHour.BusinessCompanyID,
		&operationHour.BusinessServiceID,
		&operationHour.DayOfWeek,
		&operationHour.OpenTime,
		&operationHour.CloseTime,
	)



	return &pb.DeleteBusinessCompanyServiceOperationHourResponse{
		BusinessCompanyServiceOperationHour: &operationHour,
	}, nil
}
