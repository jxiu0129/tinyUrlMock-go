package errors

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
