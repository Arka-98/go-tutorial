package structs

import "fmt"

type Address struct {
	Street  string
	City    string
	Country string
}

type Organization struct {
	ID      int
	Name    string
	Ceo     string
	Address Address
}

type Employee struct {
	ID   int
	Name string
	Role string
	Org  Organization
	Address
}

func (address Address) GetFullAddress() string {
	return fmt.Sprintln(address.Street, address.City, address.Country)
}

func (org Organization) GetOrgDetails() string {
	return fmt.Sprintln(org.ID, org.Name, org.Ceo, org.Address.GetFullAddress())
}

func (emp Employee) GetEmployeeDetails() string {
	return fmt.Sprintln(emp.ID, emp.Name, emp.Role, emp.Org.GetOrgDetails(), emp.GetFullAddress())
}
