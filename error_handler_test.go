package errorhandler

import (
	"errors"
	"testing"
)

func TestErrorHandler_Error(t *testing.T) {
	type fields struct {
		Detail    string
		ErrorInfo ErrorInfo
		HttpInfo  HttpInfo
		TraceInfo *TraceInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "empty detail",
			fields: fields{
				Detail: "",
			},
			want: "",
		}, {
			name: "nonempty detail",
			fields: fields{
				Detail: "chicken fillet",
			},
			want: "chicken fillet",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ErrorHandler{
				Detail:    tt.fields.Detail,
				ErrorInfo: tt.fields.ErrorInfo,
				HttpInfo:  tt.fields.HttpInfo,
				TraceInfo: tt.fields.TraceInfo,
			}
			if got := err.Error(); got != tt.want {
				t.Errorf("ErrorHandler.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddTrace(t *testing.T) {
	type args struct {
		err error
	}
	type want struct {
		wantErr             bool
		wantEmptyFuncName   bool
		wantEmptyLineNumber bool
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "empty err",
			args: args{
				err: nil,
			},
			want: want{
				wantEmptyFuncName:   true,
				wantEmptyLineNumber: true,
			},
		},
		{
			name: "standard error",
			args: args{
				err: errors.New("this is error"),
			},
			want: want{
				wantErr: true,
			},
		},
		{
			name: "ErrorHandler error",
			args: args{
				err: ErrorHandler{},
			},
			want: want{
				wantErr: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := AddTrace(tt.args.err)
			if (err != nil) != tt.want.wantErr {
				t.Errorf("AddTrace() error = %v, wantErr %v", err, tt.want.wantErr)
			}

			if !tt.want.wantErr {
				return
			}

			errH, ok := err.(ErrorHandler)
			if !ok {
				t.Error("Error not an ErrorHandler")
			}

			if (errH.TraceInfo.FuncName == "") != tt.want.wantEmptyFuncName {
				t.Error("Empty func name")
			}

			if (errH.TraceInfo.LineNumber == 0) != tt.want.wantEmptyLineNumber {
				t.Error("Empty func name")
			}
		})
	}
}
