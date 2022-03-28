package customerrors

import "net/http"

var (
	InvalidParameter    = &customError{code: InvalidParameterCode, httpStatus: http.StatusBadRequest}
	NotFound            = &customError{code: NotFoundCode, httpStatus: http.StatusNotFound}
	Conflict            = &customError{code: ConflictCode, httpStatus: http.StatusConflict}
	InternalServerError = &customError{code: InternalServerErrorCode, httpStatus: http.StatusInternalServerError}
)
