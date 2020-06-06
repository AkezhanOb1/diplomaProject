package company

import (
	"context"
	pb "github.com/AkezhanOb1/diplomaProject/api/proto/business/companies"
	config "github.com/AkezhanOb1/diplomaProject/configs"
	"github.com/jackc/pgx/v4"
)


//GetBusinessCompanyImagesRepository is a repository that responsible to all the requests to DB
//about business categories
func GetBusinessCompanyImagesRepository(ctx context.Context, businessCompanyID int64) ([]*pb.BusinessCompanyImage, error) {
	conn, err := pgx.Connect(ctx, config.PostgresConnection)
	if err != nil {
		return nil, err
	}

	defer conn.Close(ctx)

	sqlQuery := `SELECT id, path from business_company_image
		WHERE business_company_id=$1 ORDER BY id;`

	rows, err := conn.Query(ctx, sqlQuery, businessCompanyID)
	if err != nil {
		return nil, err
	}

	var images []*pb.BusinessCompanyImage

	for rows.Next() {
		var image pb.BusinessCompanyImage

		err = rows.Scan(
			&image.ImageID,
			&image.ImagePath,
		)

		if err != nil {
			return nil, err
		}

		images = append(images, &image)
	}

	return images, nil
}
