package slide200

import "fmt"

type Stringer interface {
	String() string
}

func ToStringer(v interface{}) (Stringer, error) {
	_, ok := v.(Stringer)
	if ok {

	} else {
		e := &MyError{"Not a stringer"}
		return e, e
	}
}

type MyError struct {
	msg string
}

func (e *MyError) String() string {
	return fmt.Sprintf("msg: %s", e.msg)
}

func (e *MyError) Error() string {
	return e.msg
}
