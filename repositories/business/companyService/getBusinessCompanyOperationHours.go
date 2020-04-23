package companyService

import (
	"context"
	pb "github.com/AkezhanOb1/diplomaProject/api/proto/business/companyServices"
	config "github.com/AkezhanOb1/diplomaProject/configs"
	"github.com/jackc/pgx/v4"
)


//GetBusinessCompanyServiceOperationHoursRepository is a repository that responsible to all the requests to DB
//about business categories
func GetBusinessCompanyServiceOperationHoursRepository(ctx context.Context, serviceID int64) (*pb.GetBusinessCompanyServiceOperationHoursResponse, error) {
	conn, err := pgx.Connect(ctx, config.PostgresConnection)
	if err != nil {
		return nil, err
	}

	defer conn.Close(ctx)


	sqlQuery := `SELECT id, business_company_id, business_service_id, day_of_week, open_time, close_time
					FROM business_company_service_operation_hours
				 WHERE business_service_id=$1;`

	rows, err := conn.Query(ctx, sqlQuery, serviceID)
	if err != nil {
		return nil, err
	}

	var operationHours pb.GetBusinessCompanyServiceOperationHoursResponse

	for rows.Next() {
		var operationHour pb.BusinessCompanyServiceOperationHour

		err = rows.Scan(
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

		operationHours.BusinessCompanyServiceOperationHour = append(operationHours.BusinessCompanyServiceOperationHour, &operationHour)
	}

	return &operationHours, nil
}
