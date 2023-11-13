package ginger

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *Error) GetCode() string {
	return e.Code
}

func (e *Error) GetMessage() string {
	return e.Message
}

func (e *Error) SetCode(code string) {
	e.Code = code
}

func (e *Error) SetMessage(msg string) {
	e.Message = msg
}
