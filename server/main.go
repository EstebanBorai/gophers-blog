package main

import(
  _ "github.com/go-sql-driver/mysql"
  queries "queries"
  "fmt"
)

func main() {
  queries.UserIndex()
}
