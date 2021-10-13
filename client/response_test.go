//go:build !integration
// +build !integration

package client

import (
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func TestResponse_GetStatus(t *testing.T) {
	t.Parallel()

	type fields struct {
		RawResponse *http.Response
		DataDump    *DataDump
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "OK status",
			fields: fields{
				RawResponse: &http.Response{
					Status: "OK",
				},
			},
			want: "OK",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Response{
				RawResponse: tt.fields.RawResponse,
				DataDump:    tt.fields.DataDump,
			}
			if got := r.GetStatus(); got != tt.want {
				t.Errorf("GetStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResponse_GetStatusCode(t *testing.T) {
	t.Parallel()

	type fields struct {
		RawResponse *http.Response
		DataDump    *DataDump
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "200",
			fields: fields{
				RawResponse: &http.Response{
					StatusCode: 200,
				},
			},
			want: 200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Response{
				RawResponse: tt.fields.RawResponse,
				DataDump:    tt.fields.DataDump,
			}
			if got := r.GetStatusCode(); got != tt.want {
				t.Errorf("GetStatusCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResponse_GetHeaders(t *testing.T) {
	t.Parallel()

	type fields struct {
		RawResponse *http.Response
		DataDump    *DataDump
	}
	tests := []struct {
		name   string
		fields fields
		want   http.Header
	}{
		{
			name: "return headers",
			fields: fields{
				RawResponse: &http.Response{
					Header: map[string][]string{
						"key1": []string{"value1"},
						"key2": []string{"value2"},
					},
				},
			},
			want: map[string][]string{
				"key1": []string{"value1"},
				"key2": []string{"value2"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Response{
				RawResponse: tt.fields.RawResponse,
				DataDump:    tt.fields.DataDump,
			}
			if got := r.GetHeaders(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetHeaders() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResponse_GetBody(t *testing.T) {
	t.Parallel()

	type fields struct {
		RawResponse *http.Response
		DataDump    *DataDump
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "return empty",
			fields: fields{
				RawResponse: nil,
			},
			want:    []byte{},
			wantErr: false,
		},
		{
			name: "return byte",
			fields: fields{
				RawResponse: &http.Response{
					Body: ioutil.NopCloser(strings.NewReader("test")),
				},
			},
			want:    []byte(`test`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Response{
				RawResponse: tt.fields.RawResponse,
				DataDump:    tt.fields.DataDump,
			}
			got, err := r.GetBody()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBody() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBody() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResponse_GetStringBody(t *testing.T) {
	type fields struct {
		RawResponse *http.Response
		DataDump    *DataDump
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "return string body",
			fields: fields{
				RawResponse: &http.Response{
					Body: ioutil.NopCloser(strings.NewReader("test")),
				},
			},
			want:    "test",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Response{
				RawResponse: tt.fields.RawResponse,
				DataDump:    tt.fields.DataDump,
			}
			got, err := r.GetStringBody()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStringBody() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetStringBody() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResponse_UnmarshalJSON(t *testing.T) {
	type product struct {
		Code string `json:"code"`
		Name string `json:"name"`
	}
	var p product

	type fields struct {
		RawResponse *http.Response
		DataDump    *DataDump
	}
	type args struct {
		target interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "return error",
			fields: fields{
				RawResponse: &http.Response{
					Body: ioutil.NopCloser(strings.NewReader("test")),
				},
			},
			args: args{
				target: p,
			},
			wantErr: true,
		},
		{
			name: "unmarshal",
			fields: fields{
				RawResponse: &http.Response{
					Body: ioutil.NopCloser(strings.NewReader(`{"code":"test", "name":"pkg1"}`)),
				},
			},
			args: args{
				target: p,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Response{
				RawResponse: tt.fields.RawResponse,
				DataDump:    tt.fields.DataDump,
			}
			if err := r.UnmarshalJSON(tt.args.target); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJson() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
