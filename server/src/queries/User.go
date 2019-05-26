package queries

import (
  "fmt"
  "encoding/json"
  _ "github.com/go-sql-driver/mysql"
  uuid "github.com/google/uuid"
  "database/sql"
)

type User struct {
  Id string `json:"id"`
  FirstName string `json:"firstName"`
  LastName string `json:"lastName"`
  UserName string `json:"userName"`
  Password string `json:"-"`
}

type UserCredentials struct {
  UserName string
  Password string
}

func UserIndex() string {
  conn, err := sql.Open("mysql", connStr)

  queryResults, err := conn.Query("SELECT * FROM users")

  if err != nil {
    panic(err.Error())
  }

  var results = []User{}

  for queryResults.Next() {
    var id, firstName, lastName, userName, password string

    err = queryResults.Scan(&id, &firstName, &lastName, &userName, &password)

    if err != nil {
      panic(err.Error())
    }

    user := User{
      Id: id,
      FirstName: firstName,
      LastName: lastName,
      UserName: userName,
      Password: password,
    }

    fmt.Println(fmt.Sprintf("Found: %s // %s // %s // %s @ songs-share.users\n",
      user.Id, user.FirstName, user.LastName, user.UserName))

    results = append(results, user)
  }

  defer conn.Close()
  jsonEncoded, err := json.Marshal(results)

  return string(jsonEncoded)
}

func InsertUser(requestBody string) string {
  var user User
  json.Unmarshal([]byte(requestBody), &user)

  conn, err := sql.Open("mysql", connStr)

  if err != nil {
    panic(err.Error())
  }

  var id string = uuid.New().String()

  var query string = fmt.Sprintf(`INSERT INTO users(id, firstName, lastName, username) VALUES('%s', '%s', '%s', '%s')`, 
    id, user.FirstName, user.LastName, user.UserName)

  fmt.Println(query)

  queryInsert, err := conn.Prepare(query)

  if err != nil {
    panic(err.Error())
  }
  defer queryInsert.Close()

  queryInsert.Exec()

  jsonEncoded, err := json.Marshal(user)
  return string(jsonEncoded)
}

func Login(requestBody string) string {
  // Enhance this
  var userCreds UserCredentials
  var user User

  json.Unmarshal([]byte(requestBody), &userCreds)

  conn, err := sql.Open("mysql", connStr)

  if err != nil {
    panic(err.Error())
  }

  var query string = fmt.Sprintf("SELECT * FROM users WHERE userName='%s' AND password='%s'", 
    userCreds.UserName, userCreds.Password)

  queryResults, err := conn.Query(query)

  for queryResults.Next() {
    var id, firstName, lastName, userName, password string

    err = queryResults.Scan(&id, &firstName, &lastName, &userName, &password)

    if err != nil {
      panic(err.Error())
    }

    user = User{
      Id: id,
      FirstName: firstName,
      LastName: lastName,
      UserName: userName,
      Password: password,
    }
  }

  defer conn.Close()
  jsonEncoded, err := json.Marshal(user)

  return string(jsonEncoded)
}
