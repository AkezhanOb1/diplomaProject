package company

import (
	pb "github.com/AkezhanOb1/diplomaProject/api/proto/business/companies"
	config "github.com/AkezhanOb1/diplomaProject/configs"
	"log"
	"github.com/jackc/pgx/v4"
	"context"

)


//GetBusinessCompanyOperationHoursRepository is a repository that responsible to all the requests to DB
//about business categories
func GetBusinessCompanyOperationHoursRepository(ctx context.Context, companyID int64) (*pb.GetBusinessCompanyOperationHoursResponse, error) {
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

	sqlQuery := `SELECT id, business_company_id, day_of_week, open_time, close_time
					FROM business_company_operation_hours
				 WHERE business_company_id=$1;`

	rows, err := conn.Query(ctx, sqlQuery, companyID)
	if err != nil {
		return nil, err
	}

	var operationHours pb.GetBusinessCompanyOperationHoursResponse

	for rows.Next() {
		var operationHour pb.BusinessCompanyOperationHour

		err = rows.Scan(
			&operationHour.CompanyOperationHourID,
			&operationHour.BusinessCompanyID,
			&operationHour.DayOfWeek,
			&operationHour.OpenTime,
			&operationHour.CloseTime,
		)

		if err != nil {
			return nil, err
		}

		operationHours.BusinessCompanyOperationHour = append(operationHours.BusinessCompanyOperationHour, &operationHour)
	}

	return &operationHours, nil
}
