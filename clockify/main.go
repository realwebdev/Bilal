package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "clockify"
)

func main() {

	psqlconnstring := fmt.Sprintf("host= %s port = %d  user = %s password = %s dbname = %s sslmode= disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconnstring)

	CheckError(err)

	insertstatement := `insert into "User"("user_id" , "email" ,  "user_name" , "password") values('1' , 'haseeb@gmail.com' , 'haseeb' , 'sasketchwan')`

	_, e := db.Exec(insertstatement)

	CheckError(e)

}

func CheckError(err error) {

	if err != nil {
		panic(err)
	}
}
