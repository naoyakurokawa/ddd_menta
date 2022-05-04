package customerrors

import (
	"net/http"
)

type conflict struct {
	err        error
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
	err        error
	code       code
	httpStatus int
}

func (e *invalidParameter) Error() string {
	return e.code.string()
}

func NewInvalidParameter() *invalidParameter {
	return &invalidParameter{code: InvalidParameterCode, httpStatus: http.StatusBadRequest}
}

type notFound struct {
	err        error
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
	err        error
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
	err        error
	code       code
	httpStatus int
}

func (u *unauthorized) Error() string {
	return u.code.string()
}

func NewUnauthorized() *unauthorized {
	return &unauthorized{code: UnauthorizedCode, httpStatus: http.StatusUnauthorized}
}
