package db

import (
    "database/sql"
    _ "github.com/lib/pq"
)

const (
    DB_USER     = "vcc"
    DB_PASSWORD = "vcc"
    DB_NAME     = "vcc"
    DB_HOST     = "localhost"
    RUNS        = 100
)


func Test(done chan int, db *sql.DB) {
	for i:=0; i<10; i++ {
	    rows, err := db.Query("SELECT * FROM ac_role")
	    CheckErr(err)

	    for rows.Next() {
		var id int
		var name string
		var static bool
		var can_create_same bool
		var description string
		err = rows.Scan(&id, &name, &static, &can_create_same, &description)
		CheckErr(err)
		//fmt.Println("uid | username")
		//fmt.Printf("%3v | %8v \n", id, name)
	    }
	}
	done <- 1;
}

func CheckErr(err error) {
    if err != nil {
        panic(err)
    }
}

