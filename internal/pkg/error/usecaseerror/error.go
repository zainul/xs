package usecaseerror

var onwErrors map[string]string

const (
	NotFoundCode                  = "00000"
	UserNotFound                  = "User not found"
	InternalServerError           = "Internal server error"
	ValidationFailed              = "Please fill the correct information body"
	FailedToGenerateAccountNumber = "Failed to generate account number"
)

func init() {
	errs := make(map[string]string)
	errs[UserNotFound] = "UCU00001"
	errs[InternalServerError] = "UCU00002"
	errs[ValidationFailed] = "UCU00003"
	errs[FailedToGenerateAccountNumber] = "UCU00004"

	onwErrors = errs
}

func GetCode(err string) string {
	if val, ok := onwErrors[err]; ok {
		return val
	}

	return NotFoundCode
}
