package main

import(
  _ "github.com/go-sql-driver/mysql"
  "database/sql"
  "fmt"
  "log"
)

func main() {
  cnn, err := sql.Open("mysql", "root:root@tcp(songs-share-db:3306)/songs-share")
  if err != nil {
    log.Fatal(err)
  }
  
  var uuid string
  var firstName string
  var lastName string
  var username string

  if err := cnn.QueryRow("SELECT * FROM users LIMIT 1").Scan(&uuid, &firstName, &lastName, &username); err != nil {
    log.Fatal(err)
  }

  fmt.Println(uuid, firstName, lastName, username)
}
