package config

import "fmt"

//PostgresAddress is the address of the postgres
const PostgresAddress = "127.0.0.1"

//PostgresPort is the port of the postgres
const PostgresPort = "5432"

//PostgresDataBase is the name of the database
const PostgresDataBase = "diploma"

//PostgresUsername is the name of the user inside DB
const PostgresUsername = "postgres"

//PostgresPassword is the password of the user
const PostgresPassword = "postgres"

//PostgresConnection is the connection string to the database
var PostgresConnection = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	PostgresAddress, PostgresPort, PostgresUsername, PostgresPassword, PostgresDataBase)
