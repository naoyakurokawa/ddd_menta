package customerrors

type code string

const (
	InvalidParameterCode    code = "InvalidParameter"
	ConflictCode            code = "Conflict"
	InternalServerErrorCode code = "InternalServerError"
	NotFoundCode            code = "NotFound"
)

func (c code) string() string {
	return string(c)
}
