package errors

import "fmt"

type CustomErr struct {
	HTTPCode int
	Code     int
	MsgCode  *int
	Verbs    map[string]string
	Data     interface{}
	Err      error
	Errs     map[string][]interface{}
}

func NewErr(httpCode, code int) *CustomErr {
	return &CustomErr{
		HTTPCode: httpCode,
		Code:     code,
	}
}
func (e *CustomErr) SetError(err error) *CustomErr {
	e.Err = err
	return e
}

func (e *CustomErr) Error() string {
	if e.Err == nil {
		return GetMessage(e.Code)
	}
	return fmt.Sprintf("%+v", e.Err)
}
