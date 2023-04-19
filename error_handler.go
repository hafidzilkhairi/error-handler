package errorhandler

import "runtime"

func (err ErrorHandler) Error() string {
	return err.Detail
}

func AddTrace(err error) error {
	if err == nil {
		return err
	}

	errHandler, ok := err.(ErrorHandler)
	if !ok {
		errHandler.Detail = err.Error()
	}

	pc, _, lineNum, ok := runtime.Caller(1)
	if !ok {
		return errHandler
	}

	ti := &TraceInfo{
		FuncName:   runtime.FuncForPC(pc).Name(),
		LineNumber: lineNum,
		Child:      errHandler.TraceInfo,
	}

	errHandler.TraceInfo = ti

	return errHandler
}
