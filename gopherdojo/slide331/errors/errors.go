package errors

import "fmt"

const (
	OmikujiErrorCode                = "100"
	OmikujiServiceErrorCode         = "101"
	OmikujiServerErrorCode          = "102"
	OmikujiRecoveryErrorCode        = "103"
	OmikujiRecoveryUnknownErrorCode = "104"
)

type OmikujiException struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

func (e *OmikujiException) Error() string {
	return fmt.Sprintf("[OmikujiException] message: [%v], code: [%v].", e.Message, e.Code)
}

func NewOmikujiException(msg, code string) *OmikujiException {
	return &OmikujiException{
		msg,
		code,
	}
}

func ThrowOmikujiException(msg, code string) {
	e := NewOmikujiException(msg, code)
	panic(e)
}
