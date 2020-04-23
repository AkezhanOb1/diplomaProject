package companyService

import (
	"context"

	pb "github.com/AkezhanOb1/diplomaProject/api/proto/business/companyServices"
	config "github.com/AkezhanOb1/diplomaProject/configs"

	"github.com/jackc/pgx/v4"
)

//CreateBusinessCompanyServiceOperationHourRepository is a repository that responsible for inserting data into the company
//table inside the database
func CreateBusinessCompanyServiceOperationHourRepository(ctx context.Context, request *pb.CreateBusinessCompanyServiceOperationHourRequest) (*pb.CreateBusinessCompanyServiceOperationHourResponse, error) {
	conn, err := pgx.Connect(ctx, config.PostgresConnection)
	if err != nil {
		return nil, err
	}

	defer conn.Close(ctx)


	sqlQuery := `INSERT INTO business_company_service_operation_hours
						(business_company_id,business_service_id, day_of_week, open_time, close_time)
				 VALUES ($1, $2, $3, $4, $5) RETURNING id;`


	var operationHourID int64
	companyID := request.GetBusinessCompanyID()
	serviceID := request.GetBusinessServiceID()
	dayOfWeek := request.GetDayOfWeek()
	openTime := request.GetOpenTime()
	closeTime := request.GetCloseTime()

	row := conn.QueryRow(context.Background(), sqlQuery, companyID, serviceID, dayOfWeek, openTime, closeTime)


	err = row.Scan(&operationHourID)
	if err != nil {
		return nil, err
	}


	return &pb.CreateBusinessCompanyServiceOperationHourResponse{
		BusinessCompanyServiceOperationHour: &pb.BusinessCompanyServiceOperationHour{
			ServiceOperationHourID:         operationHourID,
			BusinessCompanyID:      companyID ,
			BusinessServiceID:  serviceID,
			DayOfWeek: dayOfWeek,
			CloseTime: closeTime,
			OpenTime: openTime,
		},
	}, nil
}
