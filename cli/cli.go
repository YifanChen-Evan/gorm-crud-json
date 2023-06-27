package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/YifanChen-Evan/gorm-read-json/database"
)

// define command line interface struct to store command line arguments
type CLI struct {
	filePath            *string
	employeeName        *string
	employeeAge         *int
	employeeGender      *string
	employeeNationality *string
	employeeEmail       *string
	employeeDepartment  *string
	employeeReason      *string
	employeeStartDate   *string
	employeeDaysOff     *int
	command             *string

	repo database.EmployeeRepository
}

// define a function to run command line interface
func (c *CLI) Run() {
	c.parseFlags()
	c.initDatabase()

	// call function based on user input
	fmt.Println("---- &&&& ----" + *c.command)
	switch *c.command {
	case "load":
		c.loadDataCommand()
	case "add":
		c.addEmployeeCommand()
	case "get":
		c.getEmployeeCommand()
	case "update":
		c.updateEmployeeCommand()
	case "delete":
		c.deleteEmployeeCommand()
	default:
		fmt.Println("Invalid command. Please choose: load, add, get, update, delete")
	}
}

// define a function to parse the command line arguments
func (c *CLI) parseFlags() { // method signature: "parseFlags" is the method belonged to "CLI" struct
	// ---------- flag.String(argName, defaultValue, argDescription) *string ----------
	c.filePath = flag.String("file", "", "Please enter the path of the local JSON file you want to read:")
	c.employeeName = flag.String("name", "", "Please enter the name of employee:")
	c.employeeAge = flag.Int("age", 0, "Please enter the age of employee:")
	c.employeeGender = flag.String("gender", "", "Please enter the gender of employee:")
	c.employeeNationality = flag.String("nationality", "", "Please enter the nationality of employee:")
	c.employeeEmail = flag.String("email", "", "Please enter the email of employee:")
	c.employeeDepartment = flag.String("department", "", "Please enter the department of employee:")
	c.employeeReason = flag.String("reason", "", "Please enter the reason:")
	c.employeeStartDate = flag.String("startDate", "", "Please enter the start date:")
	c.command = flag.String("command", "", "Command ")
	c.employeeDaysOff = flag.Int("daysOff", 0, "Please enter the days off:")
	flag.Parse() // parse command-line flags into corresponding variables. Must be called after all flags are defined and before flags are accessed by the program.
	fmt.Println("---- $$$$$ ----" + *c.employeeName)
}

// define a function to initialize the database connection and create the table
func (c *CLI) initDatabase() {
	// connect to the database
	db, err := database.ConnectDatabase() // multiple assignment
	if err != nil {
		fmt.Println("failed to connect database:", err)
		os.Exit(1)
	}
	fmt.Println("connected to the database")

	// create table automatically by using Grom
	err = database.DB.AutoMigrate(&database.Employee{})
	if err != nil {
		fmt.Println("failed to create table:", err)
		os.Exit(1) // immediately terminate the execution of the current program and return a specific exit status code 1. Normal exit: status code 0.
	}
	fmt.Println("table created successfully")

	// create an instance of EmployeeRepository
	c.repo = database.EmployeeRepositoryInstance(db)
}

// load data from JSON file into the database
func (c *CLI) loadDataCommand() {
	err := c.repo.LoadDataFromJSON(*c.filePath)

	if err != nil {
		fmt.Println("failed to load data:", err)
	} else {
		fmt.Println("load data successfully")
	}
}

// add a new employee to the database
func (c *CLI) addEmployeeCommand() {
	employee := &database.Employee{
		// initialize fields using commande-line flags
		Name:        *c.employeeName,
		Age:         *c.employeeAge,
		Gender:      *c.employeeGender,
		Nationality: *c.employeeNationality,
		Email:       *c.employeeEmail,
		Department:  *c.employeeDepartment,
		Reason:      *c.employeeReason,
		StartDate:   *c.employeeStartDate,
		DaysOff:     *c.employeeDaysOff,
	}

	err := c.repo.AddEmployee(employee)

	if err != nil {
		fmt.Println("failed to add employee:", err)
	} else {
		fmt.Println("add employee successfully")
	}
}

// retrieve employee information from the database
func (c *CLI) getEmployeeCommand() {
	fmt.Println(c.employeeName)
	employee, err := c.repo.GetEmployeeByName(*c.employeeName) // get the specific employee information by the input name

	if err != nil {
		fmt.Println("failed to read employee:", err)
	} else {
		fmt.Println("Employee Details:")
		fmt.Println("Name:", employee.Name)
		fmt.Println("Age:", employee.Age)
		fmt.Println("Gender:", employee.Gender)
		fmt.Println("Nationality:", employee.Nationality)
		fmt.Println("Email:", employee.Email)
		fmt.Println("Department:", employee.Department)
		fmt.Println("Reason:", employee.Reason)
		fmt.Println("StartDate:", employee.StartDate)
		fmt.Println("DaysOff:", employee.DaysOff)
	}
}

// update an employee's information in the database
func (c *CLI) updateEmployeeCommand() {
	employee, err := c.repo.GetEmployeeByName(*c.employeeName) // get the specific employee information by the input name
	if err != nil {
		fmt.Println("failed to update employee:", err)
		return
	}

	// update infomation with input values
	employee.Age = *c.employeeAge
	employee.Gender = *c.employeeGender
	employee.Nationality = *c.employeeNationality
	employee.Email = *c.employeeEmail
	employee.Department = *c.employeeDepartment
	employee.Reason = *c.employeeReason
	employee.StartDate = *c.employeeStartDate
	employee.DaysOff = *c.employeeDaysOff

	err = c.repo.UpdateEmployeeByName(employee)

	if err != nil {
		fmt.Println("failed to update employee:", err)
	} else {
		fmt.Println("update employee successfully")
	}
}

// delete an employee from the database
func (c *CLI) deleteEmployeeCommand() {
	err := c.repo.DeleteEmployeeByName(*c.employeeName)

	if err != nil {
		fmt.Println("failed to delete employee:", err)
	} else {
		fmt.Println("delete employee successfully")
	}
}
