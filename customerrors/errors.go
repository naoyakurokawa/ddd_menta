package customerrors

type customError struct {
	code       code
	httpStatus int
}

func (e *customError) Error() string {
	return e.code.string()
}
