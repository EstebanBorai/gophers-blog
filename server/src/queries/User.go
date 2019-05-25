package queries

import (
  "fmt"
  "encoding/json"
)

type User struct {
  Id string
  FirstName string
  LastName string
  UserName string
}

func UserIndex() string {
  conn := DBConn()

  queryResults, err := conn.Query("SELECT * FROM users LIMIT 1")

  if err != nil {
    panic(err.Error())
  }

  var results = []User{}

  for queryResults.Next() {
    var id, firstName, lastName, userName string

    err = queryResults.Scan(&id, &firstName, &lastName, &userName)

    if err != nil {
      panic(err.Error())
    }

    user := User{
      Id: id,
      FirstName: firstName,
      LastName: lastName,
      UserName: userName,
    }

    fmt.Println(fmt.Sprintf("Found: %s // %s // %s // %s @ songs-share.users\n",
      user.Id, user.FirstName, user.LastName, user.UserName))

    results = append(results, user)
  }

  defer conn.Close()
  jsonEncoded, err := json.Marshal(results)

  return string(jsonEncoded)
}

// func InsertUser(user User{}) string {
//   conn := DBConn()

//   queryResults, err := conn.Query(fmt.Sprintf("INSERT INTO users(firstName, lastName, username)
//     VALUES('%s', '%s', '%s')", user.FirstName, user.LastName, user.UserName, ))

//   if err != nil {
//     panic(err.Error())
//   }

//   var result = User{}
// }
