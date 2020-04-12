package companyService

import (
	"context"
	pb "github.com/AkezhanOb1/diplomaProject/api/proto/business/companyServices"
	config "github.com/AkezhanOb1/diplomaProject/configs"
	"log"

	"github.com/jackc/pgx/v4"
)


//UpdateCompanyServiceRepository is a repository that responsible to inserting data into the
//company service table in database
func UpdateCompanyServiceRepository(ctx context.Context, request *pb.UpdateCompanyServiceRequest) (*pb.UpdateCompanyServiceResponse, error) {
	conn, err := pgx.Connect(ctx, config.PostgresConnection)
	if err != nil {
		return nil, err
	}

	defer func() {
		err =  conn.Close(context.Background())
		if err != nil {
			log.Println(err)
		}
	}()

	sqlQuery := `UPDATE business_company_service
					SET name=$1, duration=$2, price=$3, service_id=$4, company_id=$5
				 WHERE id=$6;`


	companyID := request.GetBusinessCompanyID()
	serviceID := request.GetBusinessServiceID()
	companyServiceID := request.GetCompanyServiceID()
	companyServiceName := request.GetCompanyServiceName()
	duration := request.GetCompanyServiceDuration()
	price := request.GetCompanyServicePrice()

	_, err = conn.Exec(
		context.Background(),
		sqlQuery,
		companyServiceName,
		duration,
		price,
		serviceID,
		companyID,
		companyServiceID,
	)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.UpdateCompanyServiceResponse{
		CompanyService: &pb.CompanyService{
			CompanyServiceID:               companyServiceID,
			CompanyServiceName:             companyServiceName,
			CompanyServiceDuration:         duration,
			CompanyServicePrice: 			price,
			BusinessCompanyID:       		companyID,
			BusinessServiceID:              serviceID,
		},
	}, nil
}

