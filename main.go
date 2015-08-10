package main

import (
	"database/sql"
	"fmt"
	"github.com/zionist/sm/contants"
	"runtime"
)

func main() {
	fmt.Printf("Hello, world.\n")
	runtime.GOMAXPROCS(runtime.NumCPU())

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable",
		constants.DB_USER, constants.DB_PASSWORD, constants.DB_NAME, constants.DB_HOST)
	db1, err := sql.Open("postgres", dbinfo)
	db1.SetMaxIdleConns(100)
	constants.CheckErr(err)
	defer db1.Close()

	done := make(chan int)
	fmt.Println("#done 1")
	for i := 0; i < constants.RUNS; i++ {
		go constants.Test(done, db1)
	}
	fmt.Println("#done 2")
	for i := 0; i < constants.RUNS; i++ {
		<-done
	}
	close(done)
	fmt.Println("#done 3")
}
