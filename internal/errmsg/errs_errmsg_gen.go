// This package was auto generated.
// DO NOT EDIT BY YOUR HAND!

package errmsg

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

// StructNameIsNilErr returns the error.
func StructNameIsNilErr(caller string) error {
	return fmt.Errorf(`[GOWRTR-1] struct name must not be empty, but it gets empty (caused at %s)`, caller)
}
func StructNameIsNilErrWrap(caller string, err error) error {
	return errors.Wrap(fmt.Errorf(`[GOWRTR-1] struct name must not be empty, but it gets empty (caused at %s)`, caller), err.Error())
}

// StructFieldNameIsEmptyErr returns the error.
func StructFieldNameIsEmptyErr(caller string) error {
	return fmt.Errorf(`[GOWRTR-2] field name must not be empty, but it gets empty (caused at %s)`, caller)
}
func StructFieldNameIsEmptyErrWrap(caller string, err error) error {
	return errors.Wrap(fmt.Errorf(`[GOWRTR-2] field name must not be empty, but it gets empty (caused at %s)`, caller), err.Error())
}

// StructFieldTypeIsEmptyErr returns the error.
func StructFieldTypeIsEmptyErr(caller string) error {
	return fmt.Errorf(`[GOWRTR-3] field type must not be empty, but it gets empty (caused at %s)`, caller)
}
func StructFieldTypeIsEmptyErrWrap(caller string, err error) error {
	return errors.Wrap(fmt.Errorf(`[GOWRTR-3] field type must not be empty, but it gets empty (caused at %s)`, caller), err.Error())
}

// FuncParameterNameIsEmptyErr returns the error.
func FuncParameterNameIsEmptyErr(caller string) error {
	return fmt.Errorf(`[GOWRTR-4] func parameter name must not be empty, but it gets empty (caused at %s)`, caller)
}
func FuncParameterNameIsEmptyErrWrap(caller string, err error) error {
	return errors.Wrap(fmt.Errorf(`[GOWRTR-4] func parameter name must not be empty, but it gets empty (caused at %s)`, caller), err.Error())
}

// LastFuncParameterTypeIsEmptyErr returns the error.
func LastFuncParameterTypeIsEmptyErr(caller string) error {
	return fmt.Errorf(`[GOWRTR-5] the last func parameter type must not be empty, but it gets empty (caused at %s)`, caller)
}
func LastFuncParameterTypeIsEmptyErrWrap(caller string, err error) error {
	return errors.Wrap(fmt.Errorf(`[GOWRTR-5] the last func parameter type must not be empty, but it gets empty (caused at %s)`, caller), err.Error())
}

// FuncNameIsEmptyError returns the error.
func FuncNameIsEmptyError(caller string) error {
	return fmt.Errorf(`[GOWRTR-6] name of func must not be empty, but it gets empty (caused at %s)`, caller)
}
func FuncNameIsEmptyErrorWrap(caller string, err error) error {
	return errors.Wrap(fmt.Errorf(`[GOWRTR-6] name of func must not be empty, but it gets empty (caused at %s)`, caller), err.Error())
}

// InterfaceNameIsEmptyError returns the error.
func InterfaceNameIsEmptyError(caller string) error {
	return fmt.Errorf(`[GOWRTR-7] name of interface must not be empty, but it gets empty (caused at %s)`, caller)
}
func InterfaceNameIsEmptyErrorWrap(caller string, err error) error {
	return errors.Wrap(fmt.Errorf(`[GOWRTR-7] name of interface must not be empty, but it gets empty (caused at %s)`, caller), err.Error())
}

// FuncReceiverNameIsEmptyError returns the error.
func FuncReceiverNameIsEmptyError(caller string) error {
	return fmt.Errorf(`[GOWRTR-8] name of func receiver must not be empty, but it gets empty (caused at %s)`, caller)
}
func FuncReceiverNameIsEmptyErrorWrap(caller string, err error) error {
	return errors.Wrap(fmt.Errorf(`[GOWRTR-8] name of func receiver must not be empty, but it gets empty (caused at %s)`, caller), err.Error())
}

// FuncReceiverTypeIsEmptyError returns the error.
func FuncReceiverTypeIsEmptyError(caller string) error {
	return fmt.Errorf(`[GOWRTR-9] type of func receiver must not be empty, but it gets empty (caused at %s)`, caller)
}
func FuncReceiverTypeIsEmptyErrorWrap(caller string, err error) error {
	return errors.Wrap(fmt.Errorf(`[GOWRTR-9] type of func receiver must not be empty, but it gets empty (caused at %s)`, caller), err.Error())
}

// FuncSignatureIsNilError returns the error.
func FuncSignatureIsNilError(caller string) error {
	return fmt.Errorf(`[GOWRTR-10] func signature must not be nil, bit it gets nil (caused at %s)`, caller)
}
func FuncSignatureIsNilErrorWrap(caller string, err error) error {
	return errors.Wrap(fmt.Errorf(`[GOWRTR-10] func signature must not be nil, bit it gets nil (caused at %s)`, caller), err.Error())
}

// AnonymousFuncSignatureIsNilError returns the error.
func AnonymousFuncSignatureIsNilError(caller string) error {
	return fmt.Errorf(`[GOWRTR-11] anonymous func signature must not be nil, bit it gets nil (caused at %s)`, caller)
}
func AnonymousFuncSignatureIsNilErrorWrap(caller string, err error) error {
	return errors.Wrap(fmt.Errorf(`[GOWRTR-11] anonymous func signature must not be nil, bit it gets nil (caused at %s)`, caller), err.Error())
}

// FuncInvocationParameterIsEmptyError returns the error.
func FuncInvocationParameterIsEmptyError(caller string) error {
	return fmt.Errorf(`[GOWRTR-12] a parameter of function invocation must not be nil, but it gets nil (caused at %s)`, caller)
}
func FuncInvocationParameterIsEmptyErrorWrap(caller string, err error) error {
	return errors.Wrap(fmt.Errorf(`[GOWRTR-12] a parameter of function invocation must not be nil, but it gets nil (caused at %s)`, caller), err.Error())
}

// CodeFormatterError returns the error.
func CodeFormatterError(cmd string, msg string, fmterr error) error {
	return fmt.Errorf(`[GOWRTR-13] code formatter raises error: command="%s", err="%s", msg="%s"`, cmd, msg, fmterr)
}
func CodeFormatterErrorWrap(cmd string, msg string, fmterr error, err error) error {
	return errors.Wrap(fmt.Errorf(`[GOWRTR-13] code formatter raises error: command="%s", err="%s", msg="%s"`, cmd, msg, fmterr), err.Error())
}

// CaseConditionIsEmptyError returns the error.
func CaseConditionIsEmptyError(caller string) error {
	return fmt.Errorf(`[GOWRTR-14] condition of case must not be empty, but it gets empty (caused at %s)`, caller)
}
func CaseConditionIsEmptyErrorWrap(caller string, err error) error {
	return errors.Wrap(fmt.Errorf(`[GOWRTR-14] condition of case must not be empty, but it gets empty (caused at %s)`, caller), err.Error())
}

// IfConditionIsEmptyError returns the error.
func IfConditionIsEmptyError(caller string) error {
	return fmt.Errorf(`[GOWRTR-15] condition of if must not be empty, but it gets empty (caused at %s)`, caller)
}
func IfConditionIsEmptyErrorWrap(caller string, err error) error {
	return errors.Wrap(fmt.Errorf(`[GOWRTR-15] condition of if must not be empty, but it gets empty (caused at %s)`, caller), err.Error())
}

// UnnamedReturnTypeAppearsAfterNamedReturnTypeError returns the error.
func UnnamedReturnTypeAppearsAfterNamedReturnTypeError(caller string) error {
	return fmt.Errorf(`[GOWRTR-16] unnamed return type appears after named return type (caused at %s)`, caller)
}
func UnnamedReturnTypeAppearsAfterNamedReturnTypeErrorWrap(caller string, err error) error {
	return errors.Wrap(fmt.Errorf(`[GOWRTR-16] unnamed return type appears after named return type (caused at %s)`, caller), err.Error())
}

// ValueOfCompositeLiteralIsEmptyError returns the error.
func ValueOfCompositeLiteralIsEmptyError(caller string) error {
	return fmt.Errorf(`[GOWRTR-17] a value of composite literal must not be empty, but it gets empty (caused at %s)`, caller)
}
func ValueOfCompositeLiteralIsEmptyErrorWrap(caller string, err error) error {
	return errors.Wrap(fmt.Errorf(`[GOWRTR-17] a value of composite literal must not be empty, but it gets empty (caused at %s)`, caller), err.Error())
}

type ErrsType int

const (
	StructNameIsNilErrType ErrsType = iota
	StructFieldNameIsEmptyErrType
	StructFieldTypeIsEmptyErrType
	FuncParameterNameIsEmptyErrType
	LastFuncParameterTypeIsEmptyErrType
	FuncNameIsEmptyErrorType
	InterfaceNameIsEmptyErrorType
	FuncReceiverNameIsEmptyErrorType
	FuncReceiverTypeIsEmptyErrorType
	FuncSignatureIsNilErrorType
	AnonymousFuncSignatureIsNilErrorType
	FuncInvocationParameterIsEmptyErrorType
	CodeFormatterErrorType
	CaseConditionIsEmptyErrorType
	IfConditionIsEmptyErrorType
	UnnamedReturnTypeAppearsAfterNamedReturnTypeErrorType
	ValueOfCompositeLiteralIsEmptyErrorType
	ErrsUnknownType
)

// ErrsList returns the list of errors.
func ListErrs() []string {
	return []string{"[GOWRTR-1] struct name must not be empty, but it gets empty (caused at %s)", "[GOWRTR-2] field name must not be empty, but it gets empty (caused at %s)", "[GOWRTR-3] field type must not be empty, but it gets empty (caused at %s)", "[GOWRTR-4] func parameter name must not be empty, but it gets empty (caused at %s)", "[GOWRTR-5] the last func parameter type must not be empty, but it gets empty (caused at %s)", "[GOWRTR-6] name of func must not be empty, but it gets empty (caused at %s)", "[GOWRTR-7] name of interface must not be empty, but it gets empty (caused at %s)", "[GOWRTR-8] name of func receiver must not be empty, but it gets empty (caused at %s)", "[GOWRTR-9] type of func receiver must not be empty, but it gets empty (caused at %s)", "[GOWRTR-10] func signature must not be nil, bit it gets nil (caused at %s)", "[GOWRTR-11] anonymous func signature must not be nil, bit it gets nil (caused at %s)", "[GOWRTR-12] a parameter of function invocation must not be nil, but it gets nil (caused at %s)", "[GOWRTR-13] code formatter raises error: command=\"%s\", err=\"%s\", msg=\"%s\"", "[GOWRTR-14] condition of case must not be empty, but it gets empty (caused at %s)", "[GOWRTR-15] condition of if must not be empty, but it gets empty (caused at %s)", "[GOWRTR-16] unnamed return type appears after named return type (caused at %s)", "[GOWRTR-17] a value of composite literal must not be empty, but it gets empty (caused at %s)"}
}

// IdentifyErrs checks the identity of an error
func IdentifyErrs(err error) ErrsType {
	errStr := err.Error()
	switch {
	case strings.HasPrefix(errStr, "[GOWRTR-1]"):
		return StructNameIsNilErrType
	case strings.HasPrefix(errStr, "[GOWRTR-2]"):
		return StructFieldNameIsEmptyErrType
	case strings.HasPrefix(errStr, "[GOWRTR-3]"):
		return StructFieldTypeIsEmptyErrType
	case strings.HasPrefix(errStr, "[GOWRTR-4]"):
		return FuncParameterNameIsEmptyErrType
	case strings.HasPrefix(errStr, "[GOWRTR-5]"):
		return LastFuncParameterTypeIsEmptyErrType
	case strings.HasPrefix(errStr, "[GOWRTR-6]"):
		return FuncNameIsEmptyErrorType
	case strings.HasPrefix(errStr, "[GOWRTR-7]"):
		return InterfaceNameIsEmptyErrorType
	case strings.HasPrefix(errStr, "[GOWRTR-8]"):
		return FuncReceiverNameIsEmptyErrorType
	case strings.HasPrefix(errStr, "[GOWRTR-9]"):
		return FuncReceiverTypeIsEmptyErrorType
	case strings.HasPrefix(errStr, "[GOWRTR-10]"):
		return FuncSignatureIsNilErrorType
	case strings.HasPrefix(errStr, "[GOWRTR-11]"):
		return AnonymousFuncSignatureIsNilErrorType
	case strings.HasPrefix(errStr, "[GOWRTR-12]"):
		return FuncInvocationParameterIsEmptyErrorType
	case strings.HasPrefix(errStr, "[GOWRTR-13]"):
		return CodeFormatterErrorType
	case strings.HasPrefix(errStr, "[GOWRTR-14]"):
		return CaseConditionIsEmptyErrorType
	case strings.HasPrefix(errStr, "[GOWRTR-15]"):
		return IfConditionIsEmptyErrorType
	case strings.HasPrefix(errStr, "[GOWRTR-16]"):
		return UnnamedReturnTypeAppearsAfterNamedReturnTypeErrorType
	case strings.HasPrefix(errStr, "[GOWRTR-17]"):
		return ValueOfCompositeLiteralIsEmptyErrorType
	default:
		return ErrsUnknownType
	}
}
