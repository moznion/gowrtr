package errmsg

//go:generate errgen -type=errs -prefix=GOWRTR-
type errs struct {
	StructNameIsNilErr        error `errmsg:"struct name must not be empty, but it gets empty"`
	StructFieldNameIsEmptyErr error `errmsg:"field name must not be empty, but it gets empty"`
}
