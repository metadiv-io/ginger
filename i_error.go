package ginger

func NewError(code, locale string) IError {
	return &Error{
		Code:    code,
		Message: GetError(code, locale),
	}
}

type IError interface {
	GetCode() string
	GetMessage() string
	SetCode(code string)
	SetMessage(msg string)
}
