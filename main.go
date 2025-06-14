package main

import (
	"gc1/config"
	"gc1/handler"
	"gc1/repository"
	"gc1/service"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	db := config.DBInit()
	defer db.Close()

	employeeRepo := repository.NewEmployeeRepository(db)
	employeeService := service.NewEmployeeService(employeeRepo)
	employeeHandler := handler.NewEmployeeHandler(employeeService)

	r := gin.Default()

	EmployeeRoutes := r.Group("/employees")
	{
		EmployeeRoutes.GET("", employeeHandler.GetAllEmployees)
		EmployeeRoutes.GET("/:id", employeeHandler.GetEmployeeById)
		EmployeeRoutes.POST("/addemployee", employeeHandler.CreateEmployee)
		EmployeeRoutes.PUT("/:id", employeeHandler.UpdateEmployee)
		EmployeeRoutes.DELETE("/:id", employeeHandler.DeleteEmployee)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback for local
	}
	r.Run(":" + port)

}
