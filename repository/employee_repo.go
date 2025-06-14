package repository

import (
	"database/sql"
	"gc1/model"
)

type EmployeeRepository interface {
	GetAllEmployees() ([]model.Employee, error)
	GetEmployeeById(id int) (model.Employee, error)
	CreateEmployee(Employee model.Employee) (model.Employee, error)
	UpdateEmployee(id int, Employee model.Employee) (model.Employee, error)
	DeleteEmployee(id int) error
}

type employeeRepository struct {
	db *sql.DB
}

func NewEmployeeRepository(db *sql.DB) EmployeeRepository {
	return &employeeRepository{db}
}

// Get all employees

func (r *employeeRepository) GetAllEmployees() ([]model.Employee, error) {
	rows, err := r.db.Query("select id,name,email,phone from employees")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var employee []model.Employee
	for rows.Next() {
		var p model.Employee
		if err := rows.Scan(&p.ID, &p.Name, &p.Email, &p.Phone); err != nil {
			return nil, err
		}
		employee = append(employee, p)
	}
	return employee, nil
}

// Get employee by id

func (r *employeeRepository) GetEmployeeById(id int) (model.Employee, error) {
	var p model.Employee
	err := r.db.QueryRow("select id,name,email,phone from employees where id = ?", id).Scan(&p.ID, &p.Name, &p.Email, &p.Phone)
	return p, err

}

// Create Employee
func (r *employeeRepository) CreateEmployee(employee model.Employee) (model.Employee, error) {
	result, err := r.db.Exec("INSERT INTO employees (name,email,phone) VALUES (?, ?,?)", employee.Name, employee.Email, employee.Phone)
	if err != nil {
		return model.Employee{}, err
	}
	id, _ := result.LastInsertId()
	employee.ID = int(id)
	return employee, nil
}

// Update employee
func (r *employeeRepository) UpdateEmployee(id int, employee model.Employee) (model.Employee, error) {
	_, err := r.db.Exec("UPDATE employees SET name=?, email=?, phone=? WHERE id=?", employee.Name, employee.Email, employee.Phone, id)
	if err != nil {
		return model.Employee{}, err
	}
	employee.ID = id
	return employee, nil
}

// Delete employee

func (r *employeeRepository) DeleteEmployee(id int) error {
	_, err := r.db.Exec("DELETE FROM employees WHERE id=?", id)
	return err
}
