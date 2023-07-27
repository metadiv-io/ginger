package types

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func NewError(code string, msg string) *Error {
	return &Error{
		Code:    code,
		Message: msg,
	}
}
