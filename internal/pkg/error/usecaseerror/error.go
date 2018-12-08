package usecaseerror

var errors map[string]string

const (
	NotFoundCode        = "00000"
	UserNotFound        = "User not found"
	InternalServerError = "Internal server error"
)

func init() {
	errs := make(map[string]string)
	errs[UserNotFound] = "UCU00001"

	errors = errs
}

func GetCode(err string) string {
	if val, ok := errors[err]; ok {
		return val
	}

	return NotFoundCode
}
