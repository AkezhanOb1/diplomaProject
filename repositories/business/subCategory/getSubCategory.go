package subCategory

import (
	"context"
	"log"

	pb "github.com/AkezhanOb1/diplomaProject/api/proto/business/subCategories"
	config "github.com/AkezhanOb1/diplomaProject/configs"

	"github.com/jackc/pgx/v4"
)

//GetSubCategoryRepository is a repository that responsible to all the requests to DB
//about business categories
func  GetSubCategoryRepository(ctx context.Context,subCategoryID int64) (*pb.BusinessSubCategoryResponse, error) {
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
					FROM public.business_sub_category
                 WHERE id = $1;`

	row := conn.QueryRow(ctx, sqlQuery, subCategoryID)

	var subCategory pb.BusinessSubCategory

	err = row.Scan(
		&subCategory.BusinessCategoryID,
		&subCategory.BusinessSubCategoryName,
		&subCategory.BusinessSubCategoryID,
		)

	if err != nil {
		return nil, err
	}

	return &pb.BusinessSubCategoryResponse{
		BusinessSubCategory: &subCategory,
	}, nil
}
