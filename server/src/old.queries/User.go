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

  if IsUserNameTaken(user.UserName) == true {
    return fmt.Sprintf("%s is already taken!", user.UserName)
  }

  conn, err := sql.Open("mysql", connStr)

  if err != nil {
    panic(err.Error())
  }

  var id string = uuid.New().String()
  var pwdHash string = HashAndSalt(user.Password)

  var query string = fmt.Sprintf(`INSERT INTO users(id, firstName, lastName, username, password) VALUES('%s', '%s', '%s', '%s', '%s')`, 
    id, user.FirstName, user.LastName, user.UserName, pwdHash)

  fmt.Println(query)

  user.Id = id

  queryInsert, err := conn.Prepare(query)

  if err != nil {
    panic(err.Error())
  }
  defer queryInsert.Close()

  queryInsert.Exec()

  jsonEncoded, err := json.Marshal(user)
  return string(jsonEncoded)
}

func IsUserNameTaken(username string) bool {
  conn, err := sql.Open("mysql", connStr)
  var exists int

  if err != nil {
    panic(err.Error())
  }

  var query string = fmt.Sprintf("SELECT COUNT(*) FROM users WHERE username='%s' LIMIT 1",
    username)

  queryResults, err := conn.Query(query)

  for queryResults.Next() {
    err = queryResults.Scan(&exists)
  }

  if exists != 0 {
    return true
  }

  return false
}

func Login(requestBody string) string {
  // Enhance this
  var userCreds UserCredentials
  var passwordHash string

  json.Unmarshal([]byte(requestBody), &userCreds)

  conn, err := sql.Open("mysql", connStr)

  if err != nil {
    panic(err.Error())
  }

  var query string = fmt.Sprintf("SELECT password FROM users WHERE username='%s' LIMIT 1", 
    userCreds.UserName)

  queryResults, err := conn.Query(query)

  for queryResults.Next() {
    err = queryResults.Scan(&passwordHash)

    if err != nil {
      panic(err.Error())
    }

    fmt.Println(passwordHash)
    fmt.Println(HashAndSalt(userCreds.Password))
    fmt.Println(userCreds.Password)
  }

  defer conn.Close()

  if CheckPassword(passwordHash, userCreds.Password) == true {
    jsonEncoded, err := json.Marshal(true)

    if err != nil {
      panic(err.Error())
    }
  
    return string(jsonEncoded)
  } else {
    return string("Bad Request")
  }
}
