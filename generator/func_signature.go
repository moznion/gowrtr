package generator

import (
	"regexp"
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
	name                       string
	typ                        string
	genericsTypeParameterNames []string
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

	if len(frt.genericsTypeParameterNames) > 0 {
		stmt += "[" + strings.Join(frt.genericsTypeParameterNames, ", ") + "]"
	}

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
	typeParameters     TypeParameters
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
	return NewFuncReturnTypeWithTypeParam(typ, []string{}, name...)
}

// NewFuncReturnTypeWithTypeParam ret	urns a new `FuncReturnType` with the generics type parameter name, e.g. `T`.
func NewFuncReturnTypeWithTypeParam(typ string, genericsTypeParameterNames []string, name ...string) *FuncReturnType {
	n := ""
	if len(name) > 0 {
		n = name[0]
	}

	return &FuncReturnType{
		name:                       n,
		typ:                        typ,
		genericsTypeParameterNames: genericsTypeParameterNames,
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
		typeParameters:     f.typeParameters,
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
		typeParameters:     f.typeParameters,
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
		typeParameters:     f.typeParameters,
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
		typeParameters:     f.typeParameters,
	}
}

// TypeParameters sets the TypeParameters onto the caller FuncSignature.
func (f *FuncSignature) TypeParameters(typeParameters TypeParameters) *FuncSignature {
	return &FuncSignature{
		funcName:           f.funcName,
		funcParameters:     f.funcParameters,
		returnTypes:        f.returnTypes,
		paramCallers:       f.paramCallers,
		funcNameCaller:     f.funcNameCaller,
		returnTypesCallers: f.returnTypesCallers,
		typeParameters:     typeParameters,
	}
}

// Generate generates a signature of the func as golang code.
func (f *FuncSignature) Generate(indentLevel int) (string, error) {
	if f.funcName == "" {
		return "", errmsg.FuncNameIsEmptyError(f.funcNameCaller)
	}

	stmt := f.funcName

	if f.typeParameters != nil && len(f.typeParameters) > 0 {
		typeParametersStmt, err := f.typeParameters.Generate(indentLevel)
		if err != nil {
			return "", err
		}
		stmt += typeParametersStmt
	}

	var typeBoundaries []int
	typeExisted := true
	typeMissingCaller := ""
	for i, param := range f.funcParameters {
		if param.name == "" {
			return "", errmsg.FuncParameterNameIsEmptyErr(f.paramCallers[i])
		}

		typeExisted = param.typ != ""
		if typeExisted {
			typeBoundaries = append(typeBoundaries, i)
		}
	}

	if !typeExisted {
		return "", errmsg.LastFuncParameterTypeIsEmptyErr(typeMissingCaller)
	}

	stmt += "("

	groups := make([]string, len(typeBoundaries))
	prevBoundary := 0
	for groupIndex, boundary := range typeBoundaries {
		group := f.funcParameters[prevBoundary : boundary+1]

		chunks := make([]string, len(group))
		for i, param := range group {
			chunk := param.name
			if param.typ != "" {
				chunk += " " + param.typ
			}
			chunks[i] = chunk
		}

		groups[groupIndex] = strings.Join(chunks, ", ")

		prevBoundary = boundary + 1
	}

	if len(groups) > 1 {
		indent := BuildIndent(indentLevel)
		nextIndent := BuildIndent(indentLevel + 1)

		stmt += "\n" + indent + nextIndent + strings.Join(groups, ",\n"+nextIndent) + ",\n" + indent
	} else if len(groups) == 1 {
		stmt += groups[0]
	}

	stmt += ")"

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

			genericsTrimmedRetType := regexp.MustCompile(`\[[^]]*]`).ReplaceAllString(retType, "")
			isNamedRetType := strings.Contains(genericsTrimmedRetType, " ")
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
