package main

import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import (
    "fmt"
    "runtime"
    "time"
)

func main(){
    runtime.GOMAXPROCS(10)
    db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/bm")
    if err != nil {
        panic(err)
    }
    count := 0
    countChan := make(chan int, 1000)

    stmt, err := db.Prepare("select col, value from bm where id = ?")
    if err != nil {
        panic(err)
    }

    defer stmt.Close()

    for i := 0; i < 1000; i++ {
        go func() {
            var col string
            var value string
            count++
            rows, err := db.Query("select col, value from bm where id = ?", i)
            if err != nil {
                panic(err)
            }
            defer rows.Close()
            for rows.Next() {
                rows.Scan(&col, &value)
                fmt.Println(col, value)
            }

            countChan <- -1
        }()
    }

    for {
        select {
        case inc := <-countChan:
            count += inc
        case <-time.After(30 * time.Second):
            panic("oh noes")
        }
        if count == 0 {
            break
        }
    }
}
