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

// StructNameIsNilErrWrap wraps the error.
func StructNameIsNilErrWrap(
	caller string,
	err error,
) error {
	return errors.Wrap(err, "[GOWRTR-1] struct name must not be empty, but it gets empty (caused at %s)")
}

// StructFieldNameIsEmptyErr returns the error.
func StructFieldNameIsEmptyErr(caller string) error {
	return fmt.Errorf(`[GOWRTR-2] field name must not be empty, but it gets empty (caused at %s)`, caller)
}

// StructFieldNameIsEmptyErrWrap wraps the error.
func StructFieldNameIsEmptyErrWrap(
	caller string,
	err error,
) error {
	return errors.Wrap(err, "[GOWRTR-2] field name must not be empty, but it gets empty (caused at %s)")
}

// StructFieldTypeIsEmptyErr returns the error.
func StructFieldTypeIsEmptyErr(caller string) error {
	return fmt.Errorf(`[GOWRTR-3] field type must not be empty, but it gets empty (caused at %s)`, caller)
}

// StructFieldTypeIsEmptyErrWrap wraps the error.
func StructFieldTypeIsEmptyErrWrap(
	caller string,
	err error,
) error {
	return errors.Wrap(err, "[GOWRTR-3] field type must not be empty, but it gets empty (caused at %s)")
}

// FuncParameterNameIsEmptyErr returns the error.
func FuncParameterNameIsEmptyErr(caller string) error {
	return fmt.Errorf(`[GOWRTR-4] func parameter name must not be empty, but it gets empty (caused at %s)`, caller)
}

// FuncParameterNameIsEmptyErrWrap wraps the error.
func FuncParameterNameIsEmptyErrWrap(
	caller string,
	err error,
) error {
	return errors.Wrap(err, "[GOWRTR-4] func parameter name must not be empty, but it gets empty (caused at %s)")
}

// LastFuncParameterTypeIsEmptyErr returns the error.
func LastFuncParameterTypeIsEmptyErr(caller string) error {
	return fmt.Errorf(`[GOWRTR-5] the last func parameter type must not be empty, but it gets empty (caused at %s)`, caller)
}

// LastFuncParameterTypeIsEmptyErrWrap wraps the error.
func LastFuncParameterTypeIsEmptyErrWrap(
	caller string,
	err error,
) error {
	return errors.Wrap(err, "[GOWRTR-5] the last func parameter type must not be empty, but it gets empty (caused at %s)")
}

// FuncNameIsEmptyError returns the error.
func FuncNameIsEmptyError(caller string) error {
	return fmt.Errorf(`[GOWRTR-6] name of func must not be empty, but it gets empty (caused at %s)`, caller)
}

// FuncNameIsEmptyErrorWrap wraps the error.
func FuncNameIsEmptyErrorWrap(
	caller string,
	err error,
) error {
	return errors.Wrap(err, "[GOWRTR-6] name of func must not be empty, but it gets empty (caused at %s)")
}

// InterfaceNameIsEmptyError returns the error.
func InterfaceNameIsEmptyError(caller string) error {
	return fmt.Errorf(`[GOWRTR-7] name of interface must not be empty, but it gets empty (caused at %s)`, caller)
}

// InterfaceNameIsEmptyErrorWrap wraps the error.
func InterfaceNameIsEmptyErrorWrap(
	caller string,
	err error,
) error {
	return errors.Wrap(err, "[GOWRTR-7] name of interface must not be empty, but it gets empty (caused at %s)")
}

// FuncReceiverNameIsEmptyError returns the error.
func FuncReceiverNameIsEmptyError(caller string) error {
	return fmt.Errorf(`[GOWRTR-8] name of func receiver must not be empty, but it gets empty (caused at %s)`, caller)
}

// FuncReceiverNameIsEmptyErrorWrap wraps the error.
func FuncReceiverNameIsEmptyErrorWrap(
	caller string,
	err error,
) error {
	return errors.Wrap(err, "[GOWRTR-8] name of func receiver must not be empty, but it gets empty (caused at %s)")
}

// FuncReceiverTypeIsEmptyError returns the error.
func FuncReceiverTypeIsEmptyError(caller string) error {
	return fmt.Errorf(`[GOWRTR-9] type of func receiver must not be empty, but it gets empty (caused at %s)`, caller)
}

// FuncReceiverTypeIsEmptyErrorWrap wraps the error.
func FuncReceiverTypeIsEmptyErrorWrap(
	caller string,
	err error,
) error {
	return errors.Wrap(err, "[GOWRTR-9] type of func receiver must not be empty, but it gets empty (caused at %s)")
}

// FuncSignatureIsNilError returns the error.
func FuncSignatureIsNilError(caller string) error {
	return fmt.Errorf(`[GOWRTR-10] func signature must not be nil, bit it gets nil (caused at %s)`, caller)
}

// FuncSignatureIsNilErrorWrap wraps the error.
func FuncSignatureIsNilErrorWrap(
	caller string,
	err error,
) error {
	return errors.Wrap(err, "[GOWRTR-10] func signature must not be nil, bit it gets nil (caused at %s)")
}

// AnonymousFuncSignatureIsNilError returns the error.
func AnonymousFuncSignatureIsNilError(caller string) error {
	return fmt.Errorf(`[GOWRTR-11] anonymous func signature must not be nil, bit it gets nil (caused at %s)`, caller)
}

// AnonymousFuncSignatureIsNilErrorWrap wraps the error.
func AnonymousFuncSignatureIsNilErrorWrap(
	caller string,
	err error,
) error {
	return errors.Wrap(err, "[GOWRTR-11] anonymous func signature must not be nil, bit it gets nil (caused at %s)")
}

// FuncInvocationParameterIsEmptyError returns the error.
func FuncInvocationParameterIsEmptyError(caller string) error {
	return fmt.Errorf(`[GOWRTR-12] a parameter of function invocation must not be nil, but it gets nil (caused at %s)`, caller)
}

// FuncInvocationParameterIsEmptyErrorWrap wraps the error.
func FuncInvocationParameterIsEmptyErrorWrap(
	caller string,
	err error,
) error {
	return errors.Wrap(err, "[GOWRTR-12] a parameter of function invocation must not be nil, but it gets nil (caused at %s)")
}

// CodeFormatterError returns the error.
func CodeFormatterError(
	cmd string,
	msg string,
	fmterr error,
) error {
	return fmt.Errorf(`[GOWRTR-13] code formatter raises error: command='%s', err='%s', msg='%s'`, cmd, msg, fmterr)
}

// CodeFormatterErrorWrap wraps the error.
func CodeFormatterErrorWrap(
	cmd string,
	msg string,
	fmterr error,
	err error,
) error {
	return errors.Wrap(err, "[GOWRTR-13] code formatter raises error: command='%s', err='%s', msg='%s'")
}

// CaseConditionIsEmptyError returns the error.
func CaseConditionIsEmptyError(caller string) error {
	return fmt.Errorf(`[GOWRTR-14] condition of case must not be empty, but it gets empty (caused at %s)`, caller)
}

// CaseConditionIsEmptyErrorWrap wraps the error.
func CaseConditionIsEmptyErrorWrap(
	caller string,
	err error,
) error {
	return errors.Wrap(err, "[GOWRTR-14] condition of case must not be empty, but it gets empty (caused at %s)")
}

// IfConditionIsEmptyError returns the error.
func IfConditionIsEmptyError(caller string) error {
	return fmt.Errorf(`[GOWRTR-15] condition of if must not be empty, but it gets empty (caused at %s)`, caller)
}

// IfConditionIsEmptyErrorWrap wraps the error.
func IfConditionIsEmptyErrorWrap(
	caller string,
	err error,
) error {
	return errors.Wrap(err, "[GOWRTR-15] condition of if must not be empty, but it gets empty (caused at %s)")
}

// UnnamedReturnTypeAppearsAfterNamedReturnTypeError returns the error.
func UnnamedReturnTypeAppearsAfterNamedReturnTypeError(caller string) error {
	return fmt.Errorf(`[GOWRTR-16] unnamed return type appears after named return type (caused at %s)`, caller)
}

// UnnamedReturnTypeAppearsAfterNamedReturnTypeErrorWrap wraps the error.
func UnnamedReturnTypeAppearsAfterNamedReturnTypeErrorWrap(
	caller string,
	err error,
) error {
	return errors.Wrap(err, "[GOWRTR-16] unnamed return type appears after named return type (caused at %s)")
}

// ValueOfCompositeLiteralIsEmptyError returns the error.
func ValueOfCompositeLiteralIsEmptyError(caller string) error {
	return fmt.Errorf(`[GOWRTR-17] a value of composite literal must not be empty, but it gets empty (caused at %s)`, caller)
}

// ValueOfCompositeLiteralIsEmptyErrorWrap wraps the error.
func ValueOfCompositeLiteralIsEmptyErrorWrap(
	caller string,
	err error,
) error {
	return errors.Wrap(err, "[GOWRTR-17] a value of composite literal must not be empty, but it gets empty (caused at %s)")
}

// ErrsType represents the error type.
type ErrsType int

const (

	// StructNameIsNilErrType represents the error type for StructNameIsNilErr.
	StructNameIsNilErrType ErrsType = iota
	// StructFieldNameIsEmptyErrType represents the error type for StructFieldNameIsEmptyErr.
	StructFieldNameIsEmptyErrType
	// StructFieldTypeIsEmptyErrType represents the error type for StructFieldTypeIsEmptyErr.
	StructFieldTypeIsEmptyErrType
	// FuncParameterNameIsEmptyErrType represents the error type for FuncParameterNameIsEmptyErr.
	FuncParameterNameIsEmptyErrType
	// LastFuncParameterTypeIsEmptyErrType represents the error type for LastFuncParameterTypeIsEmptyErr.
	LastFuncParameterTypeIsEmptyErrType
	// FuncNameIsEmptyErrorType represents the error type for FuncNameIsEmptyError.
	FuncNameIsEmptyErrorType
	// InterfaceNameIsEmptyErrorType represents the error type for InterfaceNameIsEmptyError.
	InterfaceNameIsEmptyErrorType
	// FuncReceiverNameIsEmptyErrorType represents the error type for FuncReceiverNameIsEmptyError.
	FuncReceiverNameIsEmptyErrorType
	// FuncReceiverTypeIsEmptyErrorType represents the error type for FuncReceiverTypeIsEmptyError.
	FuncReceiverTypeIsEmptyErrorType
	// FuncSignatureIsNilErrorType represents the error type for FuncSignatureIsNilError.
	FuncSignatureIsNilErrorType
	// AnonymousFuncSignatureIsNilErrorType represents the error type for AnonymousFuncSignatureIsNilError.
	AnonymousFuncSignatureIsNilErrorType
	// FuncInvocationParameterIsEmptyErrorType represents the error type for FuncInvocationParameterIsEmptyError.
	FuncInvocationParameterIsEmptyErrorType
	// CodeFormatterErrorType represents the error type for CodeFormatterError.
	CodeFormatterErrorType
	// CaseConditionIsEmptyErrorType represents the error type for CaseConditionIsEmptyError.
	CaseConditionIsEmptyErrorType
	// IfConditionIsEmptyErrorType represents the error type for IfConditionIsEmptyError.
	IfConditionIsEmptyErrorType
	// UnnamedReturnTypeAppearsAfterNamedReturnTypeErrorType represents the error type for UnnamedReturnTypeAppearsAfterNamedReturnTypeError.
	UnnamedReturnTypeAppearsAfterNamedReturnTypeErrorType
	// ValueOfCompositeLiteralIsEmptyErrorType represents the error type for ValueOfCompositeLiteralIsEmptyError.
	ValueOfCompositeLiteralIsEmptyErrorType
	// ErrsUnknownType represents unknown type for Errs
	ErrsUnknownType
)

// ListErrs returns the list of errors.
func ListErrs() []string {
	return []string{"[GOWRTR-1] struct name must not be empty, but it gets empty (caused at %s)", "[GOWRTR-2] field name must not be empty, but it gets empty (caused at %s)", "[GOWRTR-3] field type must not be empty, but it gets empty (caused at %s)", "[GOWRTR-4] func parameter name must not be empty, but it gets empty (caused at %s)", "[GOWRTR-5] the last func parameter type must not be empty, but it gets empty (caused at %s)", "[GOWRTR-6] name of func must not be empty, but it gets empty (caused at %s)", "[GOWRTR-7] name of interface must not be empty, but it gets empty (caused at %s)", "[GOWRTR-8] name of func receiver must not be empty, but it gets empty (caused at %s)", "[GOWRTR-9] type of func receiver must not be empty, but it gets empty (caused at %s)", "[GOWRTR-10] func signature must not be nil, bit it gets nil (caused at %s)", "[GOWRTR-11] anonymous func signature must not be nil, bit it gets nil (caused at %s)", "[GOWRTR-12] a parameter of function invocation must not be nil, but it gets nil (caused at %s)", "[GOWRTR-13] code formatter raises error: command='%s', err='%s', msg='%s'", "[GOWRTR-14] condition of case must not be empty, but it gets empty (caused at %s)", "[GOWRTR-15] condition of if must not be empty, but it gets empty (caused at %s)", "[GOWRTR-16] unnamed return type appears after named return type (caused at %s)", "[GOWRTR-17] a value of composite literal must not be empty, but it gets empty (caused at %s)"}
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
