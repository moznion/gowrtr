// This package was auto generated.
// DO NOT EDIT BY YOUR HAND!

package errmsg

import "errors"
import "fmt"

// StructNameIsNilErr returns the error.
func StructNameIsNilErr() error {
	return errors.New(`[GOWRTR-1] struct name must not be empty, but it gets empty`)
}

// StructFieldNameIsEmptyErr returns the error.
func StructFieldNameIsEmptyErr() error {
	return errors.New(`[GOWRTR-2] field name must not be empty, but it gets empty`)
}

// StructFieldTypeIsEmptyErr returns the error.
func StructFieldTypeIsEmptyErr() error {
	return errors.New(`[GOWRTR-3] field type must not be empty, but it gets empty`)
}

// FuncParameterNameIsEmptyErr returns the error.
func FuncParameterNameIsEmptyErr() error {
	return errors.New(`[GOWRTR-4] func parameter name must not be empty, but it gets empty`)
}

// LastFuncParameterTypeIsEmptyErr returns the error.
func LastFuncParameterTypeIsEmptyErr() error {
	return errors.New(`[GOWRTR-5] the last func parameter type must not be empty, but it gets empty`)
}

// FuncNameIsEmptyError returns the error.
func FuncNameIsEmptyError() error {
	return errors.New(`[GOWRTR-6] name of func must not be empty, but it gets empty`)
}

// InterfaceNameIsEmptyError returns the error.
func InterfaceNameIsEmptyError() error {
	return errors.New(`[GOWRTR-7] name of interface must not be empty, but it gets empty`)
}

// FuncReceiverNameIsEmptyError returns the error.
func FuncReceiverNameIsEmptyError() error {
	return errors.New(`[GOWRTR-8] name of func receiver must not be empty, but it gets empty`)
}

// FuncReceiverTypeIsEmptyError returns the error.
func FuncReceiverTypeIsEmptyError() error {
	return errors.New(`[GOWRTR-9] type of func receiver must not be empty, but it gets empty`)
}

// FuncSignatureIsNilError returns the error.
func FuncSignatureIsNilError() error {
	return errors.New(`[GOWRTR-10] func signature must not be nil, bit it gets nil`)
}

// InlineFuncSignatureIsNilError returns the error.
func InlineFuncSignatureIsNilError() error {
	return errors.New(`[GOWRTR-11] inline func signature must not be nil, bit it gets nil`)
}

// FuncInvocationParameterIsEmptyError returns the error.
func FuncInvocationParameterIsEmptyError() error {
	return errors.New(`[GOWRTR-12] a parameter of function invocation must not be nil, but it gets nil`)
}

// CodeFormatterError returns the error.
func CodeFormatterError(cmd string, msg string, err error) error {
	return fmt.Errorf(`[GOWRTR-13] code formatter raises error: command="%s", err="%s", msg="%s"`, cmd, msg, err)
}

// CaseConditionIsEmptyError returns the error.
func CaseConditionIsEmptyError() error {
	return errors.New(`[GOWRTR-14] condition of case must not be empty, but it gets empty`)
}

// ErrsList returns the list of errors.
func ErrsList() []string {
	return []string{
		`[GOWRTR-1] struct name must not be empty, but it gets empty`,
		`[GOWRTR-2] field name must not be empty, but it gets empty`,
		`[GOWRTR-3] field type must not be empty, but it gets empty`,
		`[GOWRTR-4] func parameter name must not be empty, but it gets empty`,
		`[GOWRTR-5] the last func parameter type must not be empty, but it gets empty`,
		`[GOWRTR-6] name of func must not be empty, but it gets empty`,
		`[GOWRTR-7] name of interface must not be empty, but it gets empty`,
		`[GOWRTR-8] name of func receiver must not be empty, but it gets empty`,
		`[GOWRTR-9] type of func receiver must not be empty, but it gets empty`,
		`[GOWRTR-10] func signature must not be nil, bit it gets nil`,
		`[GOWRTR-11] inline func signature must not be nil, bit it gets nil`,
		`[GOWRTR-12] a parameter of function invocation must not be nil, but it gets nil`,
		`[GOWRTR-13] code formatter raises error: command="%s", err="%s", msg="%s"`,
		`[GOWRTR-14] condition of case must not be empty, but it gets empty`,
	}
}
