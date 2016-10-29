package api

type JsonError struct {
	ErrorMessage string `json:"errorMessage"`
}

func (e JsonError) Error() string {
	return e.ErrorMessage
}

func (e JsonError) String() string {
	return e.ErrorMessage
}
