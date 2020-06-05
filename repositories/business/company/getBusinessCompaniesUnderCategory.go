package company

import (
	"context"
	pb "github.com/AkezhanOb1/diplomaProject/api/proto/business/companies"
	config "github.com/AkezhanOb1/diplomaProject/configs"
	"github.com/jackc/pgx/v4"
)


//GetBusinessCompaniesUnderCategoryRepository is a repository that responsible to all the requests to DB
//about business categories
func GetBusinessCompaniesUnderCategoryRepository(ctx context.Context, categoryID int64) (*pb.GetBusinessCompaniesUnderCategoryResponse, error) {
	conn, err := pgx.Connect(ctx, config.PostgresConnection)
	if err != nil {
		return nil, err
	}

	defer conn.Close(ctx)

	sqlQuery := `SELECT id, name, category_id, address FROM business_company WHERE category_id=$1;`

	rows, err := conn.Query(ctx, sqlQuery, categoryID)
	if err != nil {
		return nil, err
	}

	var companies pb.GetBusinessCompaniesUnderCategoryResponse

	for rows.Next() {
		var company pb.BusinessCompany

		err = rows.Scan(
			&company.BusinessCompanyID,
			&company.BusinessCompanyName,
			&company.BusinessCompanyCategoryID,
			&company.BusinessCompanyAddress,
		)

		if err != nil {
			return nil, err
		}

		companies.BusinessCompanies = append(companies.BusinessCompanies, &company)
	}

	return &companies, nil
}
