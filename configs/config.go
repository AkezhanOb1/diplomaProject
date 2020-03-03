package config

import "fmt"

//PostgresAdress is the address of the postgres
var PostgresAdress = "127.0.0.1"

//PostgresPort is the port of the postgres
var PostgresPort = "5432"

//PostgresDataBase is the name of the database
var PostgresDataBase = "diploma"

//PostgresUsername is the name of the user inside DB
var PostgresUsername = "postgres"

//PostgresPassword is the password of the user
var PostgresPassword = "postgres"

//PostgresConnection is the connection string to the database
var PostgresConnection = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	PostgresAdress, PostgresPort, PostgresUsername, PostgresPassword, PostgresDataBase)
