package buffer

import "fmt"

type CapErr struct {
	msg string
}

func (e *CapErr) Error() string { return e.msg }

func newCapError(funcName string) *CapErr {
	msg := fmt.Sprintf("%v will exceed the cap", funcName)
	return &CapErr{msg}
}

type BuffEmpErr struct {
	msg string
}

func (e *BuffEmpErr) Error() string { return e.msg }

func newBuffEmptyError(funcName string) *BuffEmpErr {
	msg := fmt.Sprintf("%v buffer is empty", funcName)
	return &BuffEmpErr{msg}
}

type BuffBoundsErr struct {
	msg string
}

func (e *BuffBoundsErr) Error() string { return e.msg }

func newBuffBoundsErr(funcName string) *BuffBoundsErr {
	msg := fmt.Sprintf("%v out of bounds", funcName)
	return &BuffBoundsErr{msg}
}
