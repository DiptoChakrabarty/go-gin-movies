# go-gin-movies

- This is a simple API developed using gin framework in go
- Makes use of sqlite3 as database
- Functonalities add,view,delete,update movies in database
- Users can login and jwt is generated

# How to Run 

```sh

- Clone the Repository

- Install binaries using go mod download

- Run using go run main.go

- Head over to http://localhost:8000/api/v1/health to check if running or no
```


# Routes

| SL No | Route | Method | Functionality |
| ---- | ------ | ---- | ------- |
| 1 | /api/v1/auth/login | POST | Login and get JWT |
| 2 | /api/v1/movies | GET | View all movies |
| 3 | /api/v1/moives | POST | Add a movie |
| 4 | /api/v1/moives/:id | PUT | Update a movie |
| 5 | /api/v1/movies/:id | DELETE | Delete a movie |
| 6 | /api/v1/health | GET | Check server health |
| 7 | /swagger/index.html | GET | Check Swagger docs |


