package main

import(
  _ "github.com/go-sql-driver/mysql"
  "database/sql"
  "fmt"
  "log"
)

func main() {
  fmt.Println("Main Process")

  cnn, err := sql.Open("mysql", "root:root@tcp(songs-share-db:3306)/songs-share")
  if err != nil {
    log.Fatal(err)
  }
  
  id := 1
  var name string

  if err := cnn.QueryRow("SELECT username FROM users WHERE id = ? LIMIT 1", id).Scan(&name); err != nil {
    log.Fatal(err)
  }

  fmt.Println(id, name)
}
