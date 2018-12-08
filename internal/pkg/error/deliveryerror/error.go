package deliveryerror

import "github.com/zainul/xs/internal/pkg/error/usecaseerror"

// Error ...
type Error struct {
	ErrorCode string `json:"error_code"`
	ErrorMsg  string `json:"error_message"`
	ErrorCase string `json:"error_case"`
}

// GetError ...
func GetError(errMsg string, caseError error) *Error {
	err := Error{}
	err.ErrorCode = usecaseerror.GetCode(errMsg)
	err.ErrorMsg = errMsg
	err.ErrorCase = caseError.Error()

	return &err
}
