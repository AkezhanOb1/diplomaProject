package business

import (
	"context"
	"fmt"
	"os"

	config "github.com/AkezhanOb1/diplomaProject/configs"
	"github.com/jackc/pgx/v4"
)

//RegistrationRepository is a repository that
//responsible to all the requests to DB
//about business categories
func RegistrationRepository() {

	conn, err := pgx.Connect(context.Background(), config.PostgresConnection)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connection to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	sqlQuery := `SELECT business_categories.id, business_categories.name 
						FROM business_categories;`

	rows, err := conn.Query(context.Background(), sqlQuery)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	for rows.Next() {
		var id int64
		var name string

		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%d. %s\n", id, name)
	}

}
