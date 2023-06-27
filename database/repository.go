package database

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"gorm.io/gorm"
)

type EmployeeRepository interface {
	LoadDataFromJSON(filePath string) error
	AddEmployee(employee *Employee) error
	GetEmployeeByName(name string) (*Employee, error)
	UpdateEmployeeByName(employee *Employee) error
	DeleteEmployeeByName(name string) error
}

// define a struct to implement "EmployeeRepository" interface (reason: seperate the method implementation of the interface from the data storage logic)
type employeeRepository struct {
	DB *gorm.DB
}

// "EmployeeRepositoryInstance" will create a instance of "EmployeeRepository" interface
func EmployeeRepositoryInstance(db *gorm.DB) EmployeeRepository {
	return &employeeRepository{
		DB: db,
	}
}

// create a function load data from local JSON file and store data in database
func (repo *employeeRepository) LoadDataFromJSON(filePath string) error {
	// -------------read local JSON file-------------
	jsonData, err := ioutil.ReadFile(filePath) // read local JSON file
	if err != nil {                            // fail
		return fmt.Errorf("failed to read JSON file: %w", err)
	}

	// -------------parse local JSON file-------------
	var employees []Employee                   // define an empty slice to store parsed JSON data
	err = json.Unmarshal(jsonData, &employees) // parse "jsonData" and store to "employees" slice
	if err != nil {                            // fail
		return fmt.Errorf("failed to parse JSON: %w", err)
	}

	// -------------delete existing data from the table-------------
	err = repo.DB.Exec("DELETE FROM employees").Error
	if err != nil {
		return fmt.Errorf("failed to delete existing data: %w", err)
	}

	// -------------store initial data in database-------------
	// using "Create()" method of Gorm to populate into database
	err = repo.DB.Create(&employees).Error // pass a slice to insert
	if err != nil {                        // fail
		return fmt.Errorf("failed to insert data in database: %w", err)
	}

	return nil // success
}

// ------------------------------CRUD------------------------------
// CREATE ----- add new employee
func (repo *employeeRepository) AddEmployee(employee *Employee) error {
	return repo.DB.Create(employee).Error
}

// READ ----- get specific employee by name in current database
func (repo *employeeRepository) GetEmployeeByName(name string) (*Employee, error) {
	var employee Employee
	err := repo.DB.Where("name = ?", name).First(&employee).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("we cannot get the employee, because he/she does not exsit") // error strings should not be capitalized and end with punctuation.
		}
		return nil, fmt.Errorf("failed to get the emplyee: %w", err)
	}
	return &employee, nil
}

// UPDATE ----- update specific employee by name in current database
func (repo *employeeRepository) UpdateEmployeeByName(employee *Employee) error {
	err := repo.DB.Model(&Employee{}).Where("name = ?", employee.Name).Updates(employee).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("we cannot update the employee, because he/she does not exsit")
		}
		return fmt.Errorf("failed to update the emplyee: %w", err)
	}
	return nil
}

// DELETE ----- delete specific employee by name in current database
func (repo *employeeRepository) DeleteEmployeeByName(name string) error {
	err := repo.DB.Delete(&Employee{}, "name = ?", name).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("we cannot delete the employee, because he/she does not exsit")
		}
		return fmt.Errorf("failed to delete the emplyee: %w", err)
	}
	return nil
}
