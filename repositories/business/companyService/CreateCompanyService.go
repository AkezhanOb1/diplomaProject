package companyService

import (
	"context"
	pb "github.com/AkezhanOb1/diplomaProject/api/proto/business/companyServices"
	config "github.com/AkezhanOb1/diplomaProject/configs"
	"log"

	"github.com/jackc/pgx/v4"
)


//CreateCompanyServiceRepository is a repository that responsible to inserting data into the
//company service table in database
func CreateCompanyServiceRepository(ctx context.Context, request *pb.CreateCompanyServiceRequest) (*pb.CreateCompanyServiceResponse, error) {
	conn, err := pgx.Connect(ctx, config.PostgresConnection)
	if err != nil {
		return nil, err
	}

	defer conn.Close(context.Background())

	sqlQuery := `INSERT INTO business_company_service(company_id, service_id, name, duration, price)
				 VALUES ($1, $2, $3, $4, $5) RETURNING id;`

	var companyServiceID int64

	companyID := request.GetBusinessCompanyID()
	serviceID := request.GetBusinessServiceID()
	companyServiceName := request.GetCompanyServiceName()
	duration := request.GetCompanyServiceDuration()
	price := request.GetCompanyServicePrice()

	row := conn.QueryRow(
		context.Background(),
		sqlQuery,
		companyID,
		serviceID,
		companyServiceName,
		duration,
		price,
	)

	err = row.Scan(&companyServiceID)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.CreateCompanyServiceResponse{
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

