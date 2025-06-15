package main

import (
	"gc1/config"
	"gc1/handler"
	"gc1/repository"
	"gc1/service"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func main() {
	config.LoadEnv()
	db := config.DBInit()
	defer db.Close()

	employeeRepo := repository.NewEmployeeRepository(db)
	employeeService := service.NewEmployeeService(employeeRepo)
	employeeHandler := handler.NewEmployeeHandler(employeeService)

	router := httprouter.New()

	router.GET("/employees", employeeHandler.GetAllEmployees)
	router.GET("/employees/:id", employeeHandler.GetEmployeeById)
	router.POST("/employees/addemployee", employeeHandler.CreateEmployee)
	router.PUT("/employees/:id", employeeHandler.UpdateEmployee)
	router.DELETE("/employees/:id", employeeHandler.DeleteEmployee)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback for local dev
	}

	log.Println("Server running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, router))

}
