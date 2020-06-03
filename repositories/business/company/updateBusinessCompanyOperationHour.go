package company

import (
	"context"
	"log"

	pb "github.com/AkezhanOb1/diplomaProject/api/proto/business/companies"
	config "github.com/AkezhanOb1/diplomaProject/configs"

	"github.com/jackc/pgx/v4"
)

//UpdateBusinessCompanyOperationHourRepository is a repository that responsible for inserting data into the company
//table inside the database
func UpdateBusinessCompanyOperationHourRepository(ctx context.Context, request *pb.BusinessCompanyOperationHour) (*pb.UpdateBusinessCompanyOperationHourResponse, error) {
	conn, err := pgx.Connect(ctx, config.PostgresConnection)
	if err != nil {
		return nil, err
	}

	defer conn.Close(ctx)


	sqlQuery := `UPDATE business_company_operation_hours
				 SET business_company_id=$1,
	 				 day_of_week=$2,
	                 open_time=$3,
	                 close_time=$4
				 WHERE id=$5`



	companyID := request.GetBusinessCompanyID()
	dayOfWeek := request.GetDayOfWeek()
	openTime := request.GetOpenTime()
	closeTime := request.GetCloseTime()
	operationHourID := request.CompanyOperationHourID

	_, err = conn.Exec(context.Background(), sqlQuery, companyID, dayOfWeek, openTime, closeTime, operationHourID)
	if err != nil {
		log.Println(err)
		return nil, err
	}


	return &pb.UpdateBusinessCompanyOperationHourResponse{
		BusinessCompanyOperationHour: &pb.BusinessCompanyOperationHour{
			CompanyOperationHourID:         operationHourID,
			BusinessCompanyID:      companyID ,
			DayOfWeek: dayOfWeek,
			CloseTime: closeTime,
			OpenTime: openTime,
		},
	}, nil
}
