package company

import (
	pb "github.com/AkezhanOb1/diplomaProject/api/proto/business/companies"
	config "github.com/AkezhanOb1/diplomaProject/configs"
	"log"
	"github.com/jackc/pgx/v4"
	"context"

)


//GetBusinessCompanyOperationHourByDayRepository is a repository that responsible to all the requests to DB
//about business categories
func GetBusinessCompanyOperationHourByDayRepository(ctx context.Context,companyID int64, dayOfWeek int64) (*pb.GetBusinessCompanyOperationHourByDayResponse, error) {
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
				 WHERE business_company_id=$1 AND day_of_week=$2;`

	row := conn.QueryRow(ctx, sqlQuery, companyID, dayOfWeek)

	var operationHour pb.BusinessCompanyOperationHour

	err = row.Scan(
		&operationHour.CompanyOperationHourID,
		&operationHour.BusinessCompanyID,
		&operationHour.DayOfWeek,
		&operationHour.OpenTime,
		&operationHour.CloseTime,
	)


	if err != nil {
		return nil, err
	}

	return &pb.GetBusinessCompanyOperationHourByDayResponse{
		BusinessCompanyOperationHour: &operationHour,
	}, nil
}