package customerrors

import (
	"fmt"
	"net/http"
)

type conflict struct {
	code       code
	httpStatus int
}

func (e *conflict) Error() string {
	return e.code.string()
}

func NewConflict() *conflict {
	return &conflict{code: ConflictCode, httpStatus: http.StatusConflict}
}

func (e *conflict) GetErrorCode() string {
	return string(e.code)
}

func (e *conflict) GetHttpStatus() int {
	return e.httpStatus
}

func (e *conflict) Equals(target error) bool {
	return e.Error() == target.Error()
}

type invalidParameter struct {
	code       code
	httpStatus int
	message    string
}

func (e *invalidParameter) Error() string {
	return fmt.Sprintf("code: %s, message: %s", e.code, e.message)
}

func (e *invalidParameter) Equals(target error) bool {
	return e.Error() == target.Error()
}

func NewInvalidParameter(message string) *invalidParameter {
	return &invalidParameter{code: InvalidParameterCode, httpStatus: http.StatusBadRequest, message: message}
}

type notFound struct {
	code       code
	httpStatus int
}

func (e *notFound) Error() string {
	return e.code.string()
}

func NewNotFound() *notFound {
	return &notFound{code: NotFoundCode, httpStatus: http.StatusNotFound}
}

type internalServerError struct {
	code       code
	httpStatus int
}

func (e *internalServerError) Error() string {
	return e.code.string()
}

func NewInternalServerError() *internalServerError {
	return &internalServerError{code: InternalServerErrorCode, httpStatus: http.StatusInternalServerError}
}

type unauthorized struct {
	code       code
	httpStatus int
	message    string
}

func (e *unauthorized) Error() string {
	return fmt.Sprintf("code: %s, message: %s", e.code, e.message)
}

func (e *unauthorized) Equals(target error) bool {
	return e.Error() == target.Error()
}

func NewUnauthorized(message string) *unauthorized {
	return &unauthorized{code: UnauthorizedCode, httpStatus: http.StatusUnauthorized, message: message}
}
