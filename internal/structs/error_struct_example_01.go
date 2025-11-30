package structs

import "fmt"

type CustomError struct {
	ID      int
	Message string
}

func (err *CustomError) Error() string {
	return fmt.Sprintf("CustomError: ID - %d message - %s", err.ID, err.Message)
}

func FindId(id int) (string, error) {
	if id != 5 {
		return "", &CustomError{
			ID:      id,
			Message: "Not found",
		}
	}

	return "Found", nil
}
