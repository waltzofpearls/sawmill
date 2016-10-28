package api

type JsonError struct {
	ErrorMessage string `json:"errorMessage"`
}

func (je *JsonError) Error() string {
	return je.ErrorMessage
}

func (je *JsonError) String() string {
	return je.ErrorMessage
}
