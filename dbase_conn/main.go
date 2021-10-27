package main

import(
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main(){
	fmt.Println("Start Connection Mysql")

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/tegra_dev")

	error_handler(err)

	defer db.Close()
	
}

func error_handler(err error){
	if err != nil{
		panic(err.Error())
	}
}
