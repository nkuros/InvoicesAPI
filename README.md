to start docker
` docker-compose up -d `

 to run migrations: 
` go run migrate/migrate.go`

endpoints:

GET http://localhost:8000/api/invoices 
Shows Invoices
optional headers:
 from_date e.g: 2020-11-30T04:20:28-02:00
 to_date e.g.: 2020-11-30T04:20:28-02:00

POST http://localhost:8000/api/invoices
Creates invoices