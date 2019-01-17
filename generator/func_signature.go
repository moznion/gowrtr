package generator

import (
	"strings"

	"github.com/moznion/gowrtr/internal/errmsg"
)

// FuncParameter represents a parameter of the func.
type FuncParameter struct {
	name string
	typ  string
}

// FuncSignature represents a code generator for the signature of the func.
type FuncSignature struct {
	funcName       string
	funcParameters []*FuncParameter
	returnTypes    []string
}

// NewFuncParameter returns a new `FuncSignature`.
func NewFuncParameter(name string, typ string) *FuncParameter {
	return &FuncParameter{
		name: name,
		typ:  typ,
	}
}

// NewFuncSignature returns a new `FuncSignature`.
func NewFuncSignature(funcName string) *FuncSignature {
	return &FuncSignature{
		funcName: funcName,
	}
}

// AddParameters adds parameters of the func to `FuncSignature`. This does *not* set, just add.
// This method returns a *new* `FuncSignature`; it means this method acts as immutable.
func (f *FuncSignature) AddParameters(funcParameters ...*FuncParameter) *FuncSignature {
	return &FuncSignature{
		funcName:       f.funcName,
		funcParameters: append(f.funcParameters, funcParameters...),
		returnTypes:    f.returnTypes,
	}
}

// Parameters sets parameters of the func to `FuncSignature`. This does *not* add, just set.
// This method returns a *new* `FuncSignature`; it means this method acts as immutable.
func (f *FuncSignature) Parameters(funcParameters ...*FuncParameter) *FuncSignature {
	return &FuncSignature{
		funcName:       f.funcName,
		funcParameters: funcParameters,
		returnTypes:    f.returnTypes,
	}
}

// AddReturnTypes adds return types of the func to `FuncSignature`. This does *not* set, just add.
// This method returns a *new* `FuncSignature`; it means this method acts as immutable.
func (f *FuncSignature) AddReturnTypes(returnTypes ...string) *FuncSignature {
	return &FuncSignature{
		funcName:       f.funcName,
		funcParameters: f.funcParameters,
		returnTypes:    append(f.returnTypes, returnTypes...),
	}
}

// ReturnTypes sets return types of the func to `FuncSignature`. This does *not* add, just set.
// This method returns a *new* `FuncSignature`; it means this method acts as immutable.
func (f *FuncSignature) ReturnTypes(returnTypes ...string) *FuncSignature {
	return &FuncSignature{
		funcName:       f.funcName,
		funcParameters: f.funcParameters,
		returnTypes:    returnTypes,
	}
}

// Generate generates a signature of the func as golang code.
func (f *FuncSignature) Generate(indentLevel int) (string, error) {
	if f.funcName == "" {
		return "", errmsg.FuncNameIsEmptyError()
	}

	stmt := f.funcName + "("

	typeExisted := true
	params := make([]string, len(f.funcParameters))
	for i, param := range f.funcParameters {
		if param.name == "" {
			return "", errmsg.FuncParameterNameIsEmptyErr()
		}

		paramSet := param.name
		typeExisted = param.typ != ""
		if typeExisted {
			paramSet += " " + param.typ
		}
		params[i] = paramSet
	}

	if !typeExisted {
		return "", errmsg.LastFuncParameterTypeIsEmptyErr()
	}

	stmt += strings.Join(params, ", ") + ")"

	returnTypes := f.returnTypes
	switch len(returnTypes) {
	case 0:
		// NOP
	case 1:
		stmt += " " + returnTypes[0]
	default:
		stmt += " (" + strings.Join(returnTypes, ", ") + ")"
	}
	return stmt, nil
}
