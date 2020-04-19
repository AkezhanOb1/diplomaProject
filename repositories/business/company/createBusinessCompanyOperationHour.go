package company

import (
	"context"
	"log"

	pb "github.com/AkezhanOb1/diplomaProject/api/proto/business/companies"
	config "github.com/AkezhanOb1/diplomaProject/configs"

	"github.com/jackc/pgx/v4"
)

//CreateBusinessCompanyOperationHourRepository is a repository that responsible for inserting data into the company
//table inside the database
func CreateBusinessCompanyOperationHourRepository(ctx context.Context, request *pb.CreateBusinessCompanyOperationHourRequest) (*pb.CreateBusinessCompanyOperationHourResponse, error) {
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

	sqlQuery := `INSERT INTO business_company_operation_hours(business_company_id, day_of_week, open_time, close_time)
				 VALUES ($1, $2, $3, $4) RETURNING id;`


	var operationHourID int64
	companyID := request.GetBusinessCompanyID()
	dayOfWeek := request.GetDayOfWeek()
	openTime := request.GetOpenTime()
	closeTime := request.GetCloseTime()

	row := conn.QueryRow(context.Background(), sqlQuery, companyID, dayOfWeek, openTime, closeTime)


	err = row.Scan(&operationHourID)
	if err != nil {
		return nil, err
	}


	return &pb.CreateBusinessCompanyOperationHourResponse{
		BusinessCompanyOperationHour: &pb.BusinessCompanyOperationHour{
			CompanyOperationHourID:         operationHourID,
			BusinessCompanyID:      companyID ,
			DayOfWeek: dayOfWeek,
			CloseTime: closeTime,
			OpenTime: openTime,
		},
	}, nil
}
