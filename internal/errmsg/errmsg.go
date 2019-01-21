package errmsg

//go:generate errgen -type=errs -prefix=GOWRTR-
type errs struct {
	StructNameIsNilErr                                error `errmsg:"struct name must not be empty, but it gets empty (caused at %s)" vars:"caller string"`
	StructFieldNameIsEmptyErr                         error `errmsg:"field name must not be empty, but it gets empty (caused at %s)" vars:"caller string"`
	StructFieldTypeIsEmptyErr                         error `errmsg:"field type must not be empty, but it gets empty (caused at %s)" vars:"caller string"`
	FuncParameterNameIsEmptyErr                       error `errmsg:"func parameter name must not be empty, but it gets empty (caused at %s)" vars:"caller string"`
	LastFuncParameterTypeIsEmptyErr                   error `errmsg:"the last func parameter type must not be empty, but it gets empty (caused at %s)" vars:"caller string"`
	FuncNameIsEmptyError                              error `errmsg:"name of func must not be empty, but it gets empty (caused at %s)" vars:"caller string"`
	InterfaceNameIsEmptyError                         error `errmsg:"name of interface must not be empty, but it gets empty (caused at %s)" vars:"caller string"`
	FuncReceiverNameIsEmptyError                      error `errmsg:"name of func receiver must not be empty, but it gets empty (caused at %s)" vars:"caller string"`
	FuncReceiverTypeIsEmptyError                      error `errmsg:"type of func receiver must not be empty, but it gets empty (caused at %s)" vars:"caller string"`
	FuncSignatureIsNilError                           error `errmsg:"func signature must not be nil, bit it gets nil (caused at %s)" vars:"caller string"`
	AnonymousFuncSignatureIsNilError                  error `errmsg:"anonymous func signature must not be nil, bit it gets nil (caused at %s)" vars:"caller string"`
	FuncInvocationParameterIsEmptyError               error `errmsg:"a parameter of function invocation must not be nil, but it gets nil (caused at %s)" vars:"caller string"`
	CodeFormatterError                                error `errmsg:"code formatter raises error: command=\"%s\", err=\"%s\", msg=\"%s\"" vars:"cmd string, msg string, err error"`
	CaseConditionIsEmptyError                         error `errmsg:"condition of case must not be empty, but it gets empty (caused at %s)" vars:"caller string"`
	IfConditionIsEmptyError                           error `errmsg:"condition of if must not be empty, but it gets empty (caused at %s)" vars:"caller string"`
	UnnamedReturnTypeAppearsAfterNamedReturnTypeError error `errmsg:"unnamed return type appears after named return type (caused at %s)" vars:"caller string"`
	ValueOfCompositeLiteralIsEmptyError               error `errmsg:"a value of composite literal must not be empty, but it gets empty (caused at %s)" vars:"caller string"`
}
