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

// FuncReturnType represents a return type of the func.
type FuncReturnType struct {
	name string
	typ  string
}

// Generate generates a return type of the func as golang code.
func (frt *FuncReturnType) Generate(indentLevel int) (string, error) {
	name := frt.name
	typ := frt.typ

	stmt := name
	if name != "" && typ != "" {
		stmt += " "
	}
	stmt += typ

	return stmt, nil
}

// FuncSignature represents a code generator for the signature of the func.
type FuncSignature struct {
	funcName           string
	funcParameters     []*FuncParameter
	returnTypes        []*FuncReturnType
	paramCallers       []string
	funcNameCaller     string
	returnTypesCallers []string
}

// NewFuncParameter returns a new `FuncSignature`.
func NewFuncParameter(name string, typ string) *FuncParameter {
	return &FuncParameter{
		name: name,
		typ:  typ,
	}
}

// NewFuncReturnType returns a new `FuncReturnType`.
// `name` is an optional parameter. If this parameter is specified, FuncReturnType generates code as named return type.
func NewFuncReturnType(typ string, name ...string) *FuncReturnType {
	n := ""
	if len(name) > 0 {
		n = name[0]
	}
	return &FuncReturnType{
		name: n,
		typ:  typ,
	}
}

// NewFuncSignature returns a new `FuncSignature`.
func NewFuncSignature(funcName string) *FuncSignature {
	return &FuncSignature{
		funcName:       funcName,
		funcNameCaller: fetchClientCallerLine(),
	}
}

// AddParameters adds parameters of the func to `FuncSignature`. This does *not* set, just add.
// This method returns a *new* `FuncSignature`; it means this method acts as immutable.
func (f *FuncSignature) AddParameters(funcParameters ...*FuncParameter) *FuncSignature {
	return &FuncSignature{
		funcName:           f.funcName,
		funcParameters:     append(f.funcParameters, funcParameters...),
		returnTypes:        f.returnTypes,
		paramCallers:       append(f.paramCallers, fetchClientCallerLineAsSlice(len(funcParameters))...),
		funcNameCaller:     f.funcNameCaller,
		returnTypesCallers: f.returnTypesCallers,
	}
}

// Parameters sets parameters of the func to `FuncSignature`. This does *not* add, just set.
// This method returns a *new* `FuncSignature`; it means this method acts as immutable.
func (f *FuncSignature) Parameters(funcParameters ...*FuncParameter) *FuncSignature {
	return &FuncSignature{
		funcName:           f.funcName,
		funcParameters:     funcParameters,
		returnTypes:        f.returnTypes,
		paramCallers:       fetchClientCallerLineAsSlice(len(funcParameters)),
		funcNameCaller:     f.funcNameCaller,
		returnTypesCallers: f.returnTypesCallers,
	}
}

// AddReturnTypes adds return types of the func to `FuncSignature`. This does *not* set, just add.
//
// This method accepts a return type as `string`. If you want to use the parameter as named one,
// please consider using `AddReturnTypeStatements()` instead of this (or use this method with string parameter like: `err error`).
//
// This method returns a *new* `FuncSignature`; it means this method acts as immutable.
func (f *FuncSignature) AddReturnTypes(returnTypes ...string) *FuncSignature {
	types := make([]*FuncReturnType, len(returnTypes))
	for i, typ := range returnTypes {
		types[i] = NewFuncReturnType(typ)
	}
	return f.AddReturnTypeStatements(types...)
}

// AddReturnTypeStatements sets return types of the func to `FuncSignature`. This does *not* add, just set.
// This method returns a *new* `FuncSignature`; it means this method acts as immutable.
func (f *FuncSignature) AddReturnTypeStatements(returnTypes ...*FuncReturnType) *FuncSignature {
	return &FuncSignature{
		funcName:           f.funcName,
		funcParameters:     f.funcParameters,
		returnTypes:        append(f.returnTypes, returnTypes...),
		paramCallers:       f.paramCallers,
		funcNameCaller:     f.funcNameCaller,
		returnTypesCallers: append(f.returnTypesCallers, fetchClientCallerLineAsSlice(len(returnTypes))...),
	}
}

// ReturnTypes sets return types of the func to `FuncSignature`. This does *not* add, just set.
//
// This method accepts a return type as `string`. If you want to use the parameter as named one,
// please consider using `ReturnTypeStatements()` instead of this (or use this method with string parameter like: `err error`).
//
// This method returns a *new* `FuncSignature`; it means this method acts as immutable.
func (f *FuncSignature) ReturnTypes(returnTypes ...string) *FuncSignature {
	types := make([]*FuncReturnType, len(returnTypes))
	for i, typ := range returnTypes {
		types[i] = NewFuncReturnType(typ)
	}
	return f.ReturnTypeStatements(types...)
}

// ReturnTypeStatements sets return types of the func to `FuncSignature`. This does *not* add, just set.
// This method returns a *new* `FuncSignature`; it means this method acts as immutable.
func (f *FuncSignature) ReturnTypeStatements(returnTypes ...*FuncReturnType) *FuncSignature {
	return &FuncSignature{
		funcName:           f.funcName,
		funcParameters:     f.funcParameters,
		returnTypes:        returnTypes,
		paramCallers:       f.paramCallers,
		funcNameCaller:     f.funcNameCaller,
		returnTypesCallers: fetchClientCallerLineAsSlice(len(returnTypes)),
	}
}

// Generate generates a signature of the func as golang code.
func (f *FuncSignature) Generate(indentLevel int) (string, error) {
	if f.funcName == "" {
		return "", errmsg.FuncNameIsEmptyError(f.funcNameCaller)
	}

	stmt := f.funcName + "("

	typeExisted := true
	typeMissingCaller := ""
	params := make([]string, len(f.funcParameters))
	for i, param := range f.funcParameters {
		if param.name == "" {
			return "", errmsg.FuncParameterNameIsEmptyErr(f.paramCallers[i])
		}

		paramSet := param.name
		typeExisted = param.typ != ""
		if typeExisted {
			paramSet += " " + param.typ
		}
		if !typeExisted {
			typeMissingCaller = f.paramCallers[i]
		}
		params[i] = paramSet
	}

	if !typeExisted {
		return "", errmsg.LastFuncParameterTypeIsEmptyErr(typeMissingCaller)
	}

	stmt += strings.Join(params, ", ") + ")"

	returnTypes := f.returnTypes
	switch len(returnTypes) {
	case 0:
		// NOP
	case 1:
		retType, _ := returnTypes[0].Generate(0)
		openingLit := " "
		closingLit := ""
		if strings.Contains(retType, " ") {
			openingLit = " ("
			closingLit = ")"
		}
		stmt += openingLit + retType + closingLit
	default:
		namedRetTypeAppeared := false
		retTypes := make([]string, len(returnTypes))
		for i, r := range returnTypes {
			retType, _ := r.Generate(0)
			retTypes[i] = retType

			isNamedRetType := strings.Contains(retType, " ")
			if !namedRetTypeAppeared {
				namedRetTypeAppeared = isNamedRetType
			}
			if namedRetTypeAppeared && !isNamedRetType {
				return "", errmsg.UnnamedReturnTypeAppearsAfterNamedReturnTypeError(f.returnTypesCallers[i])
			}
		}
		stmt += " (" + strings.Join(retTypes, ", ") + ")"
	}
	return stmt, nil
}
