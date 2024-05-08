# Candidate: Daniel Noriaki Kurosawa
## Commands

to start docker:

` docker-compose up -d `

to stop docker:

`docker-compose down`


to run API:

`go run .`


to run migrations (if needed): 

` go run migrate/migrate.go`

to run tests:

`go test ./controllers`

##  Endpoints:

GET http://localhost:8000/api/invoices 

Shows Invoices (optionally headers can be passed in order to specify a range)

optional headers:


 `from_date` e.g: `2020-11-30T04:20:28-02:00`


 `to_date` e.g.: `2020-11-30T04:20:28-02:00`

POST http://localhost:8000/api/invoices

Creates invoices and calculates fees and taxes

example payload:

`{
    "IssuedDate":    "2023-11-30T14:20:28.000+07:00",
    "PaymentAmount": 10000.0,
    "DueDate":       "2023-11-30T14:20:28.000+07:00",
    "Status":        "Pending"
}`

## File structure:

main.go:
initializes DB and sets routers

invoiceRouters.go:
Sets the Paths and calls controllers in order to execute business logic

models.go:
Contains the structs used to interface with the db and used when making 

paymentConstants.go: 
removes important values from code and centralizes them should they need alterations

migrate.go: 
Used only when DB migrations are needed

database.go: 
connects to the database

envLoader.go:
Fetches sensitive configs from app.env


### Others

DB Used:

Postgres

Reasoning: allows for one-to-many relationships needed in the following relations  
 customer to invoices, company to users, company to customers, company to invoices

### Libs used: 
gorm "gorm.io/gorm" (handling database)

gin-gonic "github.com/gin-gonic/gin" (REST API)

"github.com/DATA-DOG/go-sqlmock" (mocking db for tests)


notes :

 removed app.env on purpose from gitignore in case it's insides need to get evaluated