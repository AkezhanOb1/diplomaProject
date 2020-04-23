package companyService

import (
	"context"
	pb "github.com/AkezhanOb1/diplomaProject/api/proto/business/companyServices"
	config "github.com/AkezhanOb1/diplomaProject/configs"
	"github.com/jackc/pgx/v4"
)

//UpdateBusinessCompanyServiceOperationHour is a repository that responsible for inserting data into the company
//table inside the database
func UpdateBusinessCompanyServiceOperationHourRepository(ctx context.Context, request *pb.BusinessCompanyServiceOperationHour) (*pb.UpdateBusinessCompanyServiceOperationHourResponse, error) {
	conn, err := pgx.Connect(ctx, config.PostgresConnection)
	if err != nil {
		return nil, err
	}

	defer conn.Close(ctx)


	sqlQuery := `UPDATE business_company_service_operation_hours
				 SET business_company_id=$1,
					 business_service_id=$2,
	 				 day_of_week=$3,
	                 open_time=$4,
	                 close_time=$5
				 WHERE id=$6`



	companyID := request.GetBusinessCompanyID()
	serviceID := request.GetBusinessServiceID()
	dayOfWeek := request.GetDayOfWeek()
	openTime := request.GetOpenTime()
	closeTime := request.GetCloseTime()
	operationHourID := request.GetServiceOperationHourID()

	_, err = conn.Exec(context.Background(), sqlQuery, companyID, serviceID, dayOfWeek, openTime, closeTime, operationHourID)
	if err != nil {
		return nil, err
	}


	return &pb.UpdateBusinessCompanyServiceOperationHourResponse{
		BusinessCompanyServiceOperationHour: &pb.BusinessCompanyServiceOperationHour{
			ServiceOperationHourID:         operationHourID,
			BusinessCompanyID:      companyID ,
			BusinessServiceID: serviceID,
			DayOfWeek: dayOfWeek,
			CloseTime: closeTime,
			OpenTime: openTime,
		},
	}, nil
}
