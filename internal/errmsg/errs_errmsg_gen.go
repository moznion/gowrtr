// This package was auto generated.
// DO NOT EDIT BY YOUR HAND!

package errmsg

import (
	"errors"
	"fmt"
)

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
func FuncParameterNameIsEmptyErr(caller string) error {
	return fmt.Errorf(`[GOWRTR-4] func parameter name must not be empty, but it gets empty (caused at %s)`, caller)
}

// LastFuncParameterTypeIsEmptyErr returns the error.
func LastFuncParameterTypeIsEmptyErr(caller string) error {
	return fmt.Errorf(`[GOWRTR-5] the last func parameter type must not be empty, but it gets empty (caused at %s)`, caller)
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
func FuncSignatureIsNilError(caller string) error {
	return fmt.Errorf(`[GOWRTR-10] func signature must not be nil, bit it gets nil (caused at %s)`, caller)
}

// AnonymousFuncSignatureIsNilError returns the error.
func AnonymousFuncSignatureIsNilError(caller string) error {
	return fmt.Errorf(`[GOWRTR-11] anonymous func signature must not be nil, bit it gets nil (caused at %s)`, caller)
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
func CaseConditionIsEmptyError(caller string) error {
	return fmt.Errorf(`[GOWRTR-14] condition of case must not be empty, but it gets empty (caused at %s)`, caller)
}

// IfConditionIsEmptyError returns the error.
func IfConditionIsEmptyError() error {
	return errors.New(`[GOWRTR-15] condition of if must not be empty, but it gets empty`)
}

// UnnamedReturnTypeAppearsAfterNamedReturnTypeError returns the error.
func UnnamedReturnTypeAppearsAfterNamedReturnTypeError() error {
	return errors.New(`[GOWRTR-16] unnamed return type appears after named return type`)
}

// ValueOfCompositeLiteralIsEmptyError returns the error.
func ValueOfCompositeLiteralIsEmptyError() error {
	return errors.New(`[GOWRTR-17] a value of composite literal must not be empty, but it gets empty`)
}

// ErrsList returns the list of errors.
func ErrsList() []string {
	return []string{"[GOWRTR-1] struct name must not be empty, but it gets empty", "[GOWRTR-2] field name must not be empty, but it gets empty", "[GOWRTR-3] field type must not be empty, but it gets empty", "[GOWRTR-4] func parameter name must not be empty, but it gets empty (caused at %s)", "[GOWRTR-5] the last func parameter type must not be empty, but it gets empty (caused at %s)", "[GOWRTR-6] name of func must not be empty, but it gets empty", "[GOWRTR-7] name of interface must not be empty, but it gets empty", "[GOWRTR-8] name of func receiver must not be empty, but it gets empty", "[GOWRTR-9] type of func receiver must not be empty, but it gets empty", "[GOWRTR-10] func signature must not be nil, bit it gets nil (caused at %s)", "[GOWRTR-11] anonymous func signature must not be nil, bit it gets nil (caused at %s)", "[GOWRTR-12] a parameter of function invocation must not be nil, but it gets nil", "[GOWRTR-13] code formatter raises error: command=\"%s\", err=\"%s\", msg=\"%s\"", "[GOWRTR-14] condition of case must not be empty, but it gets empty (caused at %s)", "[GOWRTR-15] condition of if must not be empty, but it gets empty", "[GOWRTR-16] unnamed return type appears after named return type", "[GOWRTR-17] a value of composite literal must not be empty, but it gets empty"}
}
