# Api

### Authenticate
Sign up and Log in

Method | URL | Request Details
------------ | ------------- | ------------
POST | `/login` | Login

#### Login

Request Payload
```json
{
	"userName": "string",
	"password": "string"
}
```

### Users
Manage platform users

Method | URL | Request Details
------------ | ------------- | ------------
GET | `/v1/users` | [Info]()
GET | `/v1/users/:id` | [Info]()
POST | `/v1/users` | [Info]()
PUT | `/v1/users/:id` | [Info]()
DELETE | `/v1/users/:id` | [Info]()

#### Create an User

Request Payload
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
