package category

import (
	"context"
	"log"

	pb "github.com/AkezhanOb1/diplomaProject/api/proto/business/categories"
	config "github.com/AkezhanOb1/diplomaProject/configs"

	"github.com/jackc/pgx/v4"
)

//GetCategoryRepository is a repository that
//responsible to all the requests to DB
//about business categories
func  GetCategoryRepository(categoryID int64) (*pb.BusinessCategoryResponse, error) {
	conn, err := pgx.Connect(context.Background(), config.PostgresConnection)
	if err != nil {
		return nil, err
	}

	defer func() {
		err :=  conn.Close(context.Background())
		if err != nil {
			log.Fatal(err)
		}
	}()

	sqlQuery := `SELECT business_category.id, business_category.name 
						FROM business_category WHERE business_category.id = $1;`

	row := conn.QueryRow(context.Background(), sqlQuery, categoryID)

	var category pb.BusinessCategory
	err = row.Scan(&category.BusinessCategoryID, &category.BusinessCategoryName)
	if err != nil {
		return nil, err
	}

	return &pb.BusinessCategoryResponse{
		BusinessCategory: &category,
	}, nil
}
