package main

import (
    "database/sql"
    "fmt"
    "github.com/zionist/sm/db"
    "runtime"
)

func main() {
	fmt.Printf("Hello, world.\n")
	runtime.GOMAXPROCS(runtime.NumCPU())

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable",
        db.DB_USER, db.DB_PASSWORD, db.DB_NAME, db.DB_HOST)
	db1, err := sql.Open("postgres", dbinfo)
	db1.SetMaxIdleConns(100)
	db.CheckErr(err)
	defer db1.Close()

	done := make(chan int)
	fmt.Println("#done 1")
	for i:=0; i < db.RUNS; i++ {
		go db.Test(done, db1)
	}
	fmt.Println("#done 2")
	for i:=0; i< db.RUNS; i++ {
		<-done
	}
	close(done)
	fmt.Println("#done 3")
}
