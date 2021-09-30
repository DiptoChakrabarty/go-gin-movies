# go-gin-movies

- This is a simple API developed using gin framework in go
- Makes use of sqlite3 as database
- Functonalities add,view,delete,update movies in database
- Users can register themselves and then login for which jwt is generated
- If you like the project please do :star: it :innocent:

# How to Run 

```sh

- Clone the Repository

- Install binaries using go mod download

- Run using go run main.go

- Head over to http://localhost:8000/api/v1/health to check if running or no

- To generate swagger docs run the swagger.sh file

- Make sure to get swagger by using the command go get -u github.com/swaggo/swag/cmd/swag

- Replace go path in file with your go path

- Service testing is in the tests folder , run go test to run tests
```

## To Update Swagger
```sh
To update swagger run the swagger bash script and then run the application
```


# Routes

| SL No | Route | Method | Functionality |
| ---- | ------ | ---- | ------- |
| 1 | /api/v1/auth/register | POST | Registers a User |
| 2 | /api/v1/auth/login | POST | Login and get JWT |
| 3 | /api/v1/auth/userdisplay | GET | Displays all users |
| 4 | /api/v1/movies | GET | View all movies |
| 5 | /api/v1/movies/:id | GET | Displays one movie |
| 6 | /api/v1/moives | POST | Add a movie |
| 7 | /api/v1/moives/:id | PUT | Update a movie |
| 8 | /api/v1/movies/:id | DELETE | Delete a movie |
| 9 | /api/v1/health | GET | Check server health |
| 10 | /swagger/index.html | GET | Check Swagger docs |


