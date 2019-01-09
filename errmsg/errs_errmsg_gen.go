// This package was auto generated.
// DO NOT EDIT BY YOUR HAND!

package errmsg

import "errors"

func StructNameIsNilErr() error {
	return errors.New("[GOWRTR-1] struct name must not be empty, but it gets empty")
}

func StructFieldNameIsEmptyErr() error {
	return errors.New("[GOWRTR-2] field name must not be empty, but it gets empty")
}

func StructFieldTypeIsEmptyErr() error {
	return errors.New("[GOWRTR-3] field type must not be empty, but it gets empty")
}

func ErrsList() []string {
	return []string{
		`[GOWRTR-1] struct name must not be empty, but it gets empty`,
		`[GOWRTR-2] field name must not be empty, but it gets empty`,
		`[GOWRTR-3] field type must not be empty, but it gets empty`,
	}
}
