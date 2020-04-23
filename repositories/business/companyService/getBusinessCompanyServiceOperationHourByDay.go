package companyService

import (
	"context"
	pb "github.com/AkezhanOb1/diplomaProject/api/proto/business/companyServices"
	config "github.com/AkezhanOb1/diplomaProject/configs"
	"github.com/jackc/pgx/v4"
)


//GetBusinessCompanyServiceOperationHourByDayRepository is a repository that responsible to all the requests to DB
//about business categories
func GetBusinessCompanyServiceOperationHourByDayRepository(ctx context.Context,serviceID int64, dayOfWeek int64) (*pb.GetBusinessCompanyServiceOperationHourByDayResponse, error) {
	conn, err := pgx.Connect(ctx, config.PostgresConnection)
	if err != nil {
		return nil, err
	}

	defer conn.Close(ctx)

	sqlQuery := `SELECT id, business_company_id, business_service_id, day_of_week, open_time, close_time
					FROM business_company_service_operation_hours 
				 WHERE business_service_id=$1 AND day_of_week=$2;`

	row := conn.QueryRow(ctx, sqlQuery, serviceID, dayOfWeek)

	var operationHour pb.BusinessCompanyServiceOperationHour

	err = row.Scan(
		&operationHour.ServiceOperationHourID,
		&operationHour.BusinessCompanyID,
		&operationHour.BusinessServiceID,
		&operationHour.DayOfWeek,
		&operationHour.OpenTime,
		&operationHour.CloseTime,
	)


	if err != nil {
		return nil, err
	}

	return &pb.GetBusinessCompanyServiceOperationHourByDayResponse{
		BusinessCompanyServiceOperationHour: &operationHour,
	}, nil
}