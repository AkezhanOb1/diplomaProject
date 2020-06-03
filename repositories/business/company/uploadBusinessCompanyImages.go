package company

import (
    "context"

    config "github.com/AkezhanOb1/diplomaProject/configs"
    "github.com/jackc/pgx/v4"
)

//UploadBusinessCompanyImageRepository is a repository
func UploadBusinessCompanyImageRepository(ctx context.Context, path string, businessCompanyID int64) (*int64, error) {
    conn, err := pgx.Connect(ctx, config.PostgresConnection)
    if err != nil {
        return nil, err
    }

    defer conn.Close(ctx)


    sqlQuery := `INSERT INTO business_company_image(path, business_company_id)
				 VALUES ($1, $2) RETURNING id;`


    var imageID int64

    row := conn.QueryRow(ctx, sqlQuery, path, businessCompanyID)
    err = row.Scan(&imageID)
    if err != nil {
        return nil, err
    }


    return &imageID, nil
}
