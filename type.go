package errorhandler

type (
	ErrorHandler struct {
		Detail    string
		ErrorInfo ErrorInfo
		HttpInfo  HttpInfo
		TraceInfo *TraceInfo
	}

	ErrorInfo struct {
		ErrorCode    string
		ErrorMessage string
	}

	HttpInfo struct {
		HttpCode int
	}

	TraceInfo struct {
		FuncName   string
		LineNumber int
		Child      *TraceInfo
	}
)
