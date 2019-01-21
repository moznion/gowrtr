package errmsg

//go:generate errgen -type=errs -prefix=GOWRTR-
type errs struct {
	StructNameIsNilErr                                error `errmsg:"struct name must not be empty, but it gets empty"`
	StructFieldNameIsEmptyErr                         error `errmsg:"field name must not be empty, but it gets empty"`
	StructFieldTypeIsEmptyErr                         error `errmsg:"field type must not be empty, but it gets empty"`
	FuncParameterNameIsEmptyErr                       error `errmsg:"func parameter name must not be empty, but it gets empty (caused at %s)" vars:"caller string"`
	LastFuncParameterTypeIsEmptyErr                   error `errmsg:"the last func parameter type must not be empty, but it gets empty (caused at %s)" vars:"caller string"`
	FuncNameIsEmptyError                              error `errmsg:"name of func must not be empty, but it gets empty"`
	InterfaceNameIsEmptyError                         error `errmsg:"name of interface must not be empty, but it gets empty"`
	FuncReceiverNameIsEmptyError                      error `errmsg:"name of func receiver must not be empty, but it gets empty"`
	FuncReceiverTypeIsEmptyError                      error `errmsg:"type of func receiver must not be empty, but it gets empty"`
	FuncSignatureIsNilError                           error `errmsg:"func signature must not be nil, bit it gets nil (caused at %s)" vars:"caller string"`
	AnonymousFuncSignatureIsNilError                  error `errmsg:"anonymous func signature must not be nil, bit it gets nil (caused at %s)" vars:"caller string"`
	FuncInvocationParameterIsEmptyError               error `errmsg:"a parameter of function invocation must not be nil, but it gets nil"`
	CodeFormatterError                                error `errmsg:"code formatter raises error: command=\"%s\", err=\"%s\", msg=\"%s\"" vars:"cmd string, msg string, err error"`
	CaseConditionIsEmptyError                         error `errmsg:"condition of case must not be empty, but it gets empty (caused at %s)" vars:"caller string"`
	IfConditionIsEmptyError                           error `errmsg:"condition of if must not be empty, but it gets empty"`
	UnnamedReturnTypeAppearsAfterNamedReturnTypeError error `errmsg:"unnamed return type appears after named return type"`
	ValueOfCompositeLiteralIsEmptyError               error `errmsg:"a value of composite literal must not be empty, but it gets empty"`
}
