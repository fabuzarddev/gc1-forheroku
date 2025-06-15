package repository

import (
	"database/sql"
	"fmt"
	"gc1/model"
	"time"
)

type EmployeeRepository interface {
	GetAllEmployees() ([]model.ShortEmployee, error)
	GetEmployeeById(id int) (model.Employee, error)
	CreateEmployee(Employee model.Employee) (model.ShortEmployee, error)
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

func (r *employeeRepository) GetAllEmployees() ([]model.ShortEmployee, error) {
	rows, err := r.db.Query("select id,name,email from employees")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var employee []model.ShortEmployee
	for rows.Next() {
		var p model.ShortEmployee
		if err := rows.Scan(&p.ID, &p.Name, &p.Email); err != nil {
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
func (r *employeeRepository) CreateEmployee(employee model.Employee) (model.ShortEmployee, error) {
	result, err := r.db.Exec("INSERT INTO employees (name,email,phone) VALUES (?, ?,?)", employee.Name, employee.Email, employee.Phone)
	if err != nil {
		return model.ShortEmployee{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return model.ShortEmployee{}, err
	}

	if rowsAffected == 0 {
		return model.ShortEmployee{}, fmt.Errorf("karyawan tidak dapat di temukan")
	}

	id, _ := result.LastInsertId()
	employee.ID = int(id)

	newEmployee := model.ShortEmployee{
		ID:    int(id),
		Name:  employee.Name,
		Email: employee.Email,
	}

	return newEmployee, nil
}

// Update employee
func (r *employeeRepository) UpdateEmployee(id int, employee model.Employee) (model.Employee, error) {
	result, err := r.db.Exec(
		"UPDATE employees SET name=?, email=?, phone=?, updated_at=NOW() WHERE id=?",
		employee.Name, employee.Email, employee.Phone, id,
	)
	if err != nil {
		return model.Employee{}, err
	}

	// Variables to hold raw date strings
	var createdAtStr, updatedAtStr string
	var updated model.Employee

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return model.Employee{}, err
	}

	if rowsAffected == 0 {
		return model.Employee{}, fmt.Errorf("karyawan tidak dapat di temukan")
	}

	// SELECT the updated row
	err = r.db.QueryRow(
		`SELECT id, name, email, phone, created_at, updated_at FROM employees WHERE id=?`, id,
	).Scan(
		&updated.ID,
		&updated.Name,
		&updated.Email,
		&updated.Phone,
		&createdAtStr,
		&updatedAtStr,
	)
	if err != nil {
		return model.Employee{}, err
	}

	// Manually parse the time
	layout := "2006-01-02 15:04:05" // MySQL DATETIME format
	updated.CreatedAt, _ = time.Parse(layout, createdAtStr)
	updated.UpdatedAt, _ = time.Parse(layout, updatedAtStr)

	return updated, nil
}

// Delete employee

func (r *employeeRepository) DeleteEmployee(id int) error {
	result, err := r.db.Exec("DELETE FROM employees WHERE id=?", id)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("karyawan tidak dapat di temukan")
	}

	return err
}
