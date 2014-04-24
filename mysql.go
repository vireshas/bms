package main

import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "fmt"

func main(){
    db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/bm")

    if err != nil{
        fmt.Println(err)
    }

    var col string
    var value string

    stmt, err := db.Prepare("select col, value from bm where id = ?")

    if err != nil{
        fmt.Println(err)
    }

    defer stmt.Close()

    for i := 0; i < 1000; i++ {
        rows, err := stmt.Query(i)

        if err != nil{
            fmt.Println(err)
        }

        defer rows.Close()

        for rows.Next() {
            err := rows.Scan(&col, &value)

            if err != nil{
                fmt.Println(err)
            }

            fmt.Println(col, value)
        }
    }
}
