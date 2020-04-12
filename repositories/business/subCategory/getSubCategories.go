package subCategory

import (
	"context"
	"log"

	pb "github.com/AkezhanOb1/diplomaProject/api/proto/business/subCategories"
	config "github.com/AkezhanOb1/diplomaProject/configs"

	"github.com/jackc/pgx/v4"
)

//GetSubCategoriesRepository is a repository that responsible to all the requests to DB
//about business categories
func GetSubCategoriesRepository(ctx context.Context) (*pb.BusinessSubCategoriesResponse, error) {
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

	sqlQuery := `SELECT id, name, category_id
						FROM business_sub_category;`

	rows, err := conn.Query(ctx, sqlQuery)
	if err != nil {
		return nil, err
	}

	var subCategories pb.BusinessSubCategoriesResponse

	for rows.Next() {
		var subCategory pb.BusinessSubCategory

		err = rows.Scan(
			&subCategory.BusinessSubCategoryID,
			&subCategory.BusinessSubCategoryName,
			&subCategory.BusinessCategoryID,
			)

		if err != nil {
			return nil, err
		}

		subCategories.BusinessSubCategories = append(subCategories.BusinessSubCategories, &subCategory)
	}

	return &subCategories, nil
}
