package types

import "fmt"

type Employee struct {
	UserID       int    `json:"userId"`
	Age          int    `json:"age"`
	Name         string `json:"name"`
	DepartmentID int    `json:"departmentId"`
}

func (e Employee) String() string {
	return fmt.Sprintf("UserID: %d; Age: %d; Name: %s; DepartmentID: %d; ", e.UserID, e.Age, e.Name, e.DepartmentID)
}
