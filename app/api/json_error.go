package api

import "fmt"

type JsonError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (e JsonError) Error() string {
	return fmt.Sprintf("Error (%d): %s", e.Status, e.Message)
}

func (e JsonError) String() string {
	return fmt.Sprintf("Error (%d): %s", e.Status, e.Message)
}
