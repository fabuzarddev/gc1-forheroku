package service

import (
	"errors"
	"gc1/model"
	"gc1/repository"
	"strings"
)

type EmployeeService interface {
	GetAllEmployees() ([]model.ShortEmployee, error)
	GetEmployeeById(id int) (model.Employee, error)
	CreateEmployee(Employee model.Employee) (model.ShortEmployee, error)
	UpdateEmployee(id int, Employee model.Employee) (model.Employee, error)
	DeleteEmployee(id int) error
}

type employeeService struct {
	repo repository.EmployeeRepository
}

func NewEmployeeService(r repository.EmployeeRepository) EmployeeService {
	return &employeeService{repo: r}
}

// Get all employees

func (s *employeeService) GetAllEmployees() ([]model.ShortEmployee, error) {
	return s.repo.GetAllEmployees()
}

// Get employee by ids

func (s *employeeService) GetEmployeeById(id int) (model.Employee, error) {
	return s.repo.GetEmployeeById(id)
}

// Create employee

func (s *employeeService) CreateEmployee(employee model.Employee) (model.ShortEmployee, error) {
	if employee.Name == "" || employee.Email == "" || employee.Phone == "" {
		return model.ShortEmployee{}, errors.New("nama, email, dan phone wajib di isi")
	}

	createdEmployee, err := s.repo.CreateEmployee(employee)
	if err != nil {
		// Cek duplicate
		if strings.Contains(err.Error(), "Duplicate entry") {
			return model.ShortEmployee{}, errors.New("email sudah digunakan")
		}
		return model.ShortEmployee{}, err
	}

	return createdEmployee, nil
}

// Update employee

func (s *employeeService) UpdateEmployee(id int, Employee model.Employee) (model.Employee, error) {
	if Employee.Name == "" || Employee.Email == "" || Employee.Phone == "" {
		return model.Employee{}, errors.New("nama, email, dan phone wajib di isi")
	}
	return s.repo.UpdateEmployee(id, Employee)

}

// Delete employee

func (s *employeeService) DeleteEmployee(id int) error {
	return s.repo.DeleteEmployee(id)
}
