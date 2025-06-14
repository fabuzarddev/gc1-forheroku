package handler

import (
	"gc1/model"
	"gc1/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EmployeeHandler struct {
	Service service.EmployeeService
}

func NewEmployeeHandler(s service.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{Service: s}
}

// Get all product

func (h *EmployeeHandler) GetAllEmployees(c *gin.Context) {
	product, err := h.Service.GetAllEmployees()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}

// Get by id
func (h *EmployeeHandler) GetEmployeeById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid id"})
		return
	}

	employee, err := h.Service.GetEmployeeById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, employee)
}

// Create employee

func (h *EmployeeHandler) CreateEmployee(c *gin.Context) {
	var employee model.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	createdEmployee, err := h.Service.CreateEmployee(employee)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
	}
	c.JSON(http.StatusCreated, createdEmployee)

}

// Update employee
func (h *EmployeeHandler) UpdateEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid id"})
		return
	}
	var employee model.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	updatedEmployee, err := h.Service.UpdateEmployee(id, employee)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid id"})
		return
	}
	c.JSON(http.StatusOK, updatedEmployee)

}

// Delete employee
func (h *EmployeeHandler) DeleteEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid id"})
		return
	}

	if err := h.Service.DeleteEmployee(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Employee deleted successfully"})
}
