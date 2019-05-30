# Users API

## Authenticate
> Creates a JWT with the current user claims and serve it as Response Cookie
HTTP Method | Request URI
**POST** | `/login`

### Request Body
```json
{
	"userName": "string",
	"password": "string"
}
```

## Create User
> Creates a new user

HTTP Method | Request URI
--- | --- 
**POST** | `/api/v1/users`

### Request Body
```json
{
  "userName": "string",
  "password": "string",
  "firstName": "string",
  "lastName": "string",
  "birthday": "Date",
  "dateJoined": "Date"
}
```

### Response Body
```json 
{
  "id": "string",
  "userName": "string",
  "firstName": "string",
  "lastName": "string",
  "birthday": "Date",
  "dateJoined": "Date"
}
```
