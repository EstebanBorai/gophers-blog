# Users API
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
