package category

import (
	"context"
	"log"

	pb "github.com/AkezhanOb1/diplomaProject/api/proto/business/categories"
	config "github.com/AkezhanOb1/diplomaProject/configs"

	"github.com/jackc/pgx/v4"
)

//GetCategoriesRepository is a repository that
//responsible to all the requests to DB
//about business categories
func GetCategoriesRepository() (*pb.BusinessCategoriesResponse, error) {
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
						FROM business_category;`

	rows, err := conn.Query(context.Background(), sqlQuery)
	if err != nil {
		return nil, err
	}

	var businessCategories pb.BusinessCategoriesResponse

	for rows.Next() {
		var businessCategory pb.BusinessCategory

		err := rows.Scan(&businessCategory.BusinessCategoryID, &businessCategory.BusinessCategoryName)
		if err != nil {
			return nil, err
		}
		businessCategories.BusinessCategories = append(businessCategories.BusinessCategories, &businessCategory)
	}

	return &businessCategories, nil
}
