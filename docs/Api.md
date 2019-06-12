# Api

### Authenticate
Sign up and Log in

Method | URL | Authentication
------------ | ------------- | ------------
POST | `/login` | None
POST | `/signup` | None

#### Login
Validates username and password and returns a new JWT

Request Body
```json
{
  "userName": "string",
  "password": "string"
}
```

Response Body
```json
{
  "code": "number",
  "expire": "Date",
  "token": "string"
}
```

#### Sign Up
Registers a new user

Request Body
```json
{
  "userName": "string",
  "password": "string",
  "firstName": "string",
  "lastName": "string",
  "email": "string",
  "birthday": "Date"
}
```

Response Body
```json
{
  "id": "string",
  "userName": "string",
  "firstName": "string",
  "lastName": "string",
  "email": "string",
  "birthday": "Date",
  "dateJoined": "Date"
}
```

### Users
Manage platform users

Method | URL | Authentication
------------ | ------------- | ------------
GET | `/v1/users` | Bearer Token (JWT)
GET | `/v1/users/:id` | Bearer Token (JWT)
PUT | `/v1/users/:id` | Bearer Token (JWT)
DELETE | `/v1/users/:id` | Bearer Token (JWT)

#### Create an User

Request Body
```json
{
  "userName": "string",
  "password": "string",
  "firstName": "string",
  "lastName": "string",
  "email": "string",
  "birthday": "Date"
}
```
