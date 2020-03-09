package owner

import (
	"context"
	pb "github.com/AkezhanOb1/diplomaProject/api/proto/business/owners"
	config "github.com/AkezhanOb1/diplomaProject/configs"
	"log"

	"github.com/jackc/pgx/v4"
)

//TODO: MAKE WITH TX
//CreateOwnerRepository is a repository that
//responsible to inserting data into the owner
//table in database
func BindCompanyToBusinessOwner(ctx context.Context, request *pb.BindCompanyToBusinessOwnerRequest) (*pb.BindCompanyToBusinessOwnerResponse, error) {
	conn, err := pgx.Connect(ctx, config.PostgresConnection)
	if err != nil {
		return nil, err
	}

	defer func() {
		err :=  conn.Close(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}()

	sqlQuery := `INSERT INTO owners_businesses(
						business_owner_id, business_company_id)
                 VALUES ($1, $2) RETURNING true;`

	row := conn.QueryRow(ctx, sqlQuery,request.GetBusinessOwnerID(), request.GetBusinessCompanyID())

	var success bool
	err = row.Scan(&success)
	if err != nil {
		return nil, err
	}

	return &pb.BindCompanyToBusinessOwnerResponse{
		BindSuccess: success,
	}, nil
}

