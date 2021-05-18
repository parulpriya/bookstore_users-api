package errors

type RestErr struct {
	Message string `json:"message"`
	StatusCode int `json:"status_code"`
	Error string `json:"error"`
}

func BadRequestErr(message string) *RestErr {
	restErr := RestErr{
		Message:    message,
		StatusCode: 400,
		Error:      "bad_request",
	}
	return &restErr
}

func NotFoundErr(message string) *RestErr {
	restErr := RestErr{
		Message:    message,
		StatusCode: 404,
		Error:      "not_found",
	}
	return &restErr
}