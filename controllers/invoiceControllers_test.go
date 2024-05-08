package controllers

import (
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
func TestGetAllInvoicesFromPeriod(t *testing.T) {
	t.Parallel()
	// Create a new Gin router
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Create a new mock database connection
	db, mock, _ := sqlmock.New()

	// Create a new GORM DB instance with the mock database connection
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
        DriverName:           "postgres",
        Conn:                 db,
        PreferSimpleProtocol: true,
	}), &gorm.Config{})

    if err != nil {
        t.Fatalf("gorm postgres fatal: %v", err)
	}

	// Create a new instance of the InvoiceController with the mock DB
	controller := NewInvoiceController(gormDB)
	// set router
	router.GET("/invoices", controller.GetAllInvoicesFromPeriod)

	// Define the test case
	t.Run("GetAllInvoicesFromPeriod", func(t *testing.T) {
		// Set up the mock database response
		mockTime,_ := time.Parse(time.RFC3339, "2024-05-07T14:57:53.8867612-02:00")
		rows := sqlmock.NewRows([]string{"id", "amount", "tax", "fee","due_date" }).
			AddRow(1, 100.0, 10.0, 5.0, mockTime).
			AddRow(2, 200.0, 20.0, 10.0, mockTime)
			
		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM \"invoices\" WHERE \"invoices\".\"deleted_at\" IS NULL")).WillReturnRows(rows)

		// Create a new HTTP request
		req, _ := http.NewRequest("GET", "/invoices", nil)

		// Create a new HTTP response recorder
		rec := httptest.NewRecorder()

	

		// Perform the request
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		expectedBody := `[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"IssuedDate":"0001-01-01T00:00:00Z","PaymentAmount":0,"Fee":5,"FeeRate":0,"Tax":10,"TaxRate":0,"TotalAmount":0,"DueDate":"2024-05-07T14:57:53.8867612-02:00","Status":""},{"ID":2,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"IssuedDate":"0001-01-01T00:00:00Z","PaymentAmount":0,"Fee":10,"FeeRate":0,"Tax":20,"TaxRate":0,"TotalAmount":0,"DueDate":"2024-05-07T14:57:53.8867612-02:00","Status":""}]`
		if rec.Body.String() != expectedBody {
			t.Errorf("Expected response body %s, but got %s", expectedBody, rec.Body.String())
		}
		assert.Equal(t,expectedBody,rec.Body.String())
	
	})

	t.Run("GetAllInvoicesFromPeriod_WithFromDateHeader", func(t *testing.T) {
		// Set up the mock database response
		mockTime ,_ := time.Parse(time.RFC3339, "2024-05-07T14:57:53.8867612-03:00")
		mockTime2, _ := time.Parse(time.RFC3339, "2023-05-07T14:57:53.8867612-02:00")
		rows := sqlmock.NewRows([]string{"id", "amount", "tax", "fee","due_date" }).
			AddRow(1, 100.0, 10.0, 5.0, mockTime).
			AddRow(2, 200.0, 20.0, 10.0,mockTime2)

		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM \"invoices\" WHERE due_date \u003e= $1 AND \"invoices\".\"deleted_at\" IS NULL")).WillReturnRows(rows)

		// Create a new HTTP request
		req, _ := http.NewRequest("GET", "/invoices", nil)
		req.Header.Set("from_date", "2022-01-30T04:20:28-02:00")


		// // Create a new HTTP response recorder
		rec := httptest.NewRecorder()

		
			// Perform the request
			router.ServeHTTP(rec, req)


		assert.Equal(t, http.StatusOK, rec.Code)

		expectedBody := `[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"IssuedDate":"0001-01-01T00:00:00Z","PaymentAmount":0,"Fee":5,"FeeRate":0,"Tax":10,"TaxRate":0,"TotalAmount":0,"DueDate":"2024-05-07T14:57:53.8867612-03:00","Status":""},{"ID":2,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"IssuedDate":"0001-01-01T00:00:00Z","PaymentAmount":0,"Fee":10,"FeeRate":0,"Tax":20,"TaxRate":0,"TotalAmount":0,"DueDate":"2023-05-07T14:57:53.8867612-02:00","Status":""}]`
		if rec.Body.String() != expectedBody {
			t.Errorf("Expected response body %s, but got %s", expectedBody, rec.Body.String())
		}
		assert.Equal(t,expectedBody,rec.Body.String())

		
	})
	t.Run("GetAllInvoicesFromPeriod_WithToDateHeader", func(t *testing.T) {
		// Set up the mock database response
		mockTime,_ := time.Parse(time.RFC3339, "2022-05-07T14:57:53.8867612-02:00")
		mockTime2,_ := time.Parse(time.RFC3339, "2023-05-07T14:57:53.8867612-03:00")
		rows := sqlmock.NewRows([]string{"id", "amount", "tax", "fee","due_date" }).
			AddRow(1, 100.0, 10.0, 5.0, mockTime).
			AddRow(2, 200.0, 20.0, 10.0,mockTime2)
			
		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM \"invoices\" WHERE due_date \u003c= $1 AND \"invoices\".\"deleted_at\" IS NULL")).WillReturnRows(rows)

		// Create a new HTTP request
		req, _ := http.NewRequest("GET", "/invoices", nil)

		req.Header.Set("to_date", "2024-05-06T14:57:53.8867612-03:00")

		// // Create a new HTTP response recorder
		rec := httptest.NewRecorder()

		
			// Perform the request
			router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
	
		expectedBody := `[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"IssuedDate":"0001-01-01T00:00:00Z","PaymentAmount":0,"Fee":5,"FeeRate":0,"Tax":10,"TaxRate":0,"TotalAmount":0,"DueDate":"2022-05-07T14:57:53.8867612-02:00","Status":""},{"ID":2,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"IssuedDate":"0001-01-01T00:00:00Z","PaymentAmount":0,"Fee":10,"FeeRate":0,"Tax":20,"TaxRate":0,"TotalAmount":0,"DueDate":"2023-05-07T14:57:53.8867612-03:00","Status":""}]`
		if rec.Body.String() != expectedBody {
			t.Errorf("Expected response body %s, but got %s", expectedBody, rec.Body.String())
		}

		assert.Equal(t,expectedBody,rec.Body.String())


		
	})
	t.Run("GetAllInvoicesFromPeriod_WitBothhHeaders", func(t *testing.T) {
		// Set up the mock database response
		mockTime,_ := time.Parse(time.RFC3339, "2024-05-07T14:57:53.8867612-03:00")
		mockTime2,_ := time.Parse(time.RFC3339, "2022-05-07T14:57:53.8867612-02:00")
		rows := sqlmock.NewRows([]string{"id", "amount", "tax", "fee","due_date" }).
			AddRow(1, 100.0, 10.0, 5.0, mockTime).
			AddRow(2, 200.0, 20.0, 10.0,mockTime2)
			
		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM \"invoices\" WHERE (due_date BETWEEN $1 AND $2) AND \"invoices\".\"deleted_at\" IS NULL")).WillReturnRows(rows)

		// Create a new HTTP request
		req, _ := http.NewRequest("GET", "/invoices", nil)
		req.Header.Set("from_date", "2021-01-30T04:20:28-02:00")
		req.Header.Set("to_date", "2023-11-30T04:20:28-02:00")

		// // Create a new HTTP response recorder
		rec := httptest.NewRecorder()

		
			// Perform the request
			router.ServeHTTP(rec, req)

		// Check the response status code
		if rec.Code != http.StatusOK {
			t.Errorf("Expected status code %d, but got %d", http.StatusOK, rec.Code)
		}

		assert.Equal(t, http.StatusOK, rec.Code)

		expectedBody := `[{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"IssuedDate":"0001-01-01T00:00:00Z","PaymentAmount":0,"Fee":5,"FeeRate":0,"Tax":10,"TaxRate":0,"TotalAmount":0,"DueDate":"2024-05-07T14:57:53.8867612-03:00","Status":""},{"ID":2,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"IssuedDate":"0001-01-01T00:00:00Z","PaymentAmount":0,"Fee":10,"FeeRate":0,"Tax":20,"TaxRate":0,"TotalAmount":0,"DueDate":"2022-05-07T14:57:53.8867612-02:00","Status":""}]`
		if rec.Body.String() != expectedBody {
			t.Errorf("Expected response body %s, but got %s", expectedBody, rec.Body.String())
		}
		
		assert.Equal(t,expectedBody,rec.Body.String())
	})
	t.Run("GetAllInvoicesFromPeriod_BadRequest", func(t *testing.T) {


		// Create a new HTTP request
		req, _ := http.NewRequest("GET", "/invoices", nil)
		req.Header.Set("to_date", "2021-01-30T04:20:28-02:00")
		req.Header.Set("from_date", "2023-11-30T04:20:28-02:00")

		// // Create a new HTTP response recorder
		rec := httptest.NewRecorder()

		
		// Perform the request
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)

	})
}		


func TestCreateInvoice(t *testing.T) {
	t.Parallel()
	// Create a new Gin router
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Create a new mock database connection
	db, mock, _ := sqlmock.New()
	// Create a new GORM DB instance with the mock database connection
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
        DriverName:           "postgres",
        Conn:                 db,
        PreferSimpleProtocol: true,
	}), &gorm.Config{})

    if err != nil {
        t.Fatalf("gorm postgres fatal: %v", err)
	}
	// Create a new instance of the InvoiceController with the mock DB
	controller := NewInvoiceController(gormDB)
	// Set Router
	router.POST("/invoices", controller.CreateInvoice)
	t.Run("CreateInvoice", func(t *testing.T) {
		// Set up the mock database response
		now ,_ := time.Parse(time.RFC3339, "2022-05-07T14:57:53.8867612-02:00")
		rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "fee","fee_rate","tax","tax_rate","total_amount","due_date","status"}).
			AddRow(1, now, now, 400, 0.04,1000,0.1,10000,now,"Pending")

		mock.ExpectBegin()
		mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO \"invoices\" (\"created_at\",\"updated_at\",\"deleted_at\",\"issued_date\",\"payment_amount\",\"fee\",\"fee_rate\",\"tax\",\"tax_rate\",\"total_amount\",\"due_date\",\"status\") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12) RETURNING \"id\"")).
			WillReturnRows(rows)
		mock.ExpectCommit()

		// Create a new HTTP request
		req, _ := http.NewRequest("POST", "/invoices", strings.NewReader(`{"PaymentAmount": 10000.0}`,))
		req.Header.Set("Content-Type", "application/json")

		// Create a new HTTP response recorder
		rec := httptest.NewRecorder()


		// Perform the request
		router.ServeHTTP(rec, req)

		// Check the response status code
		if rec.Code != http.StatusCreated {
			t.Errorf("Expected status code %d, but got %d", http.StatusCreated, rec.Code)
		}
		assert.Equal(t, http.StatusCreated, rec.Code)
		// Check the response body
		expectedBody := `{"ID":1,"CreatedAt":"2022-05-07T14:57:53.8867612-02:00","UpdatedAt":"2022-05-07T14:57:53.8867612-02:00","DeletedAt":null,"IssuedDate":"0001-01-01T00:00:00Z","PaymentAmount":10000,"Fee":400,"FeeRate":0.04,"Tax":1000,"TaxRate":0.1,"TotalAmount":10000,"DueDate":"2022-05-07T14:57:53.8867612-02:00","Status":"Pending"}`
		
		if rec.Body.String() != expectedBody {
			t.Errorf("Expected response body %s, but got %s", expectedBody, rec.Body.String())
		}

		assert.Equal(t,expectedBody,rec.Body.String())
	})

	t.Run("CreateInvoice_InvalidRequestBody", func(t *testing.T) {


		// Create a new HTTP request with an invalid request body
		req, _ := http.NewRequest("POST", "/invoices", strings.NewReader(
			`{
				"InvalidField": "InvalidValue"
			}`,
		))
		req.Header.Set("Content-Type", "application/json")
	
		// Create a new HTTP response recorder
		rec := httptest.NewRecorder()
	
		// Perform the request
		router.ServeHTTP(rec, req)
	
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
}


