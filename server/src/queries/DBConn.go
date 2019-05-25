package queries

import(
  _ "github.com/go-sql-driver/mysql"
  "database/sql"
  "fmt"
)

func DBConn() (db *sql.DB) {
  dbDriver := "mysql"
  dbUser := "root"
  dbPass := "root"
  dbName := "songs-share"

  connStr := fmt.Sprintf("%s:%s@tcp(songs-share-db:3306)/%s", dbUser, dbPass, dbName)

  fmt.Println(fmt.Sprintf("Using %s as connection string for MySQL Database", connStr))

  db, err := sql.Open(dbDriver, connStr)

  if err != nil {
    panic(err.Error())
  }

  return db
}
