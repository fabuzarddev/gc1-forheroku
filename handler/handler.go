package handler

import (
	"encoding/json"
	"gc1/model"
	"gc1/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type EmployeeHandler struct {
	Service service.EmployeeService
}

func NewEmployeeHandler(s service.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{Service: s}
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func (h *EmployeeHandler) GetAllEmployees(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	employees, err := h.Service.GetAllEmployees()
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	if len(employees) == 0 {
		writeJSON(w, http.StatusOK, map[string]string{"message": "Tidak ada data employee"})
		return

	}

	writeJSON(w, http.StatusOK, employees)
}

func (h *EmployeeHandler) GetEmployeeById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
		return
	}

	employee, err := h.Service.GetEmployeeById(id)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, employee)
}

func (h *EmployeeHandler) CreateEmployee(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var employee model.Employee
	if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	created, err := h.Service.CreateEmployee(employee)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	response := model.SuccessResponse{
		Message: "Karyawan berhasil ditambahkan",
		Data:    created,
	}

	writeJSON(w, http.StatusCreated, response)
}

func (h *EmployeeHandler) UpdateEmployee(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
		return
	}

	var employee model.Employee
	if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	updated, err := h.Service.UpdateEmployee(id, employee)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	response := model.SuccessResponse{
		Message: "Karyawan berhasil diupdate",
		Data:    updated,
	}

	writeJSON(w, http.StatusOK, response)
}

func (h *EmployeeHandler) DeleteEmployee(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
		return
	}

	if err := h.Service.DeleteEmployee(id); err != nil {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"message": "Employee deleted successfully"})
}
