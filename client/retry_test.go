//go:build !integration
// +build !integration

package client

import (
	"context"
	"crypto/x509"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"testing"
)

func Test_baseRetryPolicy(t *testing.T) {
	type args struct {
		resp *http.Response
		err  error
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "redirects error",
			args: args{
				resp: &http.Response{},
				err: &url.Error{
					Op:  "",
					URL: "",
					Err: fmt.Errorf("stopped after 5 redirects"),
				},
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "scheme error",
			args: args{
				resp: &http.Response{},
				err: &url.Error{
					Op:  "",
					URL: "",
					Err: fmt.Errorf("unsupported protocol scheme"),
				},
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "unknown authority error",
			args: args{
				resp: &http.Response{},
				err: &url.Error{
					Op:  "",
					URL: "",
					Err: x509.UnknownAuthorityError{},
				},
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "OpError",
			args: args{
				resp: &http.Response{},
				err: &url.Error{
					Op:  "",
					URL: "",
					Err: &net.OpError{},
				},
			},
			want:    true,
			wantErr: true,
		},
		{
			name: "generic error",
			args: args{
				resp: &http.Response{},
				err:  fmt.Errorf("some error"),
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "status code 429",
			args: args{
				resp: &http.Response{StatusCode: http.StatusTooManyRequests},
				err:  nil,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "status code 429",
			args: args{
				resp: &http.Response{StatusCode: http.StatusTooManyRequests},
				err:  nil,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "status code 0",
			args: args{
				resp: &http.Response{StatusCode: 0},
				err:  nil,
			},
			want:    true,
			wantErr: true,
		},
		{
			name: "status code 501",
			args: args{
				resp: &http.Response{StatusCode: 501},
				err:  nil,
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "status code 502",
			args: args{
				resp: &http.Response{StatusCode: 502},
				err:  nil,
			},
			want:    true,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := baseRetryPolicy(tt.args.resp, tt.args.err)
			if (err != nil) != tt.wantErr {
				t.Errorf("baseRetryPolicy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("baseRetryPolicy() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultRetryPolicy(t *testing.T) {
	cc, cancel := context.WithCancel(context.Background())
	cancel()

	type args struct {
		ctx  context.Context
		resp *http.Response
		err  error
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "context canceled",
			args: args{
				ctx:  cc,
				resp: &http.Response{},
				err:  nil,
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "redirects error",
			args: args{
				ctx:  context.Background(),
				resp: &http.Response{},
				err: &url.Error{
					Op:  "",
					URL: "",
					Err: fmt.Errorf("stopped after 5 redirects"),
				},
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "scheme error",
			args: args{
				ctx:  context.Background(),
				resp: &http.Response{},
				err: &url.Error{
					Op:  "",
					URL: "",
					Err: fmt.Errorf("unsupported protocol scheme"),
				},
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "unknown authority error",
			args: args{
				ctx:  context.Background(),
				resp: &http.Response{},
				err: &url.Error{
					Op:  "",
					URL: "",
					Err: x509.UnknownAuthorityError{},
				},
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "OpError",
			args: args{
				ctx:  context.Background(),
				resp: &http.Response{},
				err: &url.Error{
					Op:  "",
					URL: "",
					Err: &net.OpError{},
				},
			},
			want:    true,
			wantErr: true,
		},
		{
			name: "generic error",
			args: args{
				ctx:  context.Background(),
				resp: &http.Response{},
				err:  fmt.Errorf("some error"),
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "status code 429",
			args: args{
				ctx:  context.Background(),
				resp: &http.Response{StatusCode: http.StatusTooManyRequests},
				err:  nil,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "status code 429",
			args: args{
				ctx:  context.Background(),
				resp: &http.Response{StatusCode: http.StatusTooManyRequests},
				err:  nil,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "status code 0",
			args: args{
				ctx:  context.Background(),
				resp: &http.Response{StatusCode: 0},
				err:  nil,
			},
			want:    true,
			wantErr: true,
		},
		{
			name: "status code 501",
			args: args{
				ctx:  context.Background(),
				resp: &http.Response{StatusCode: 501},
				err:  nil,
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "status code 502",
			args: args{
				ctx:  context.Background(),
				resp: &http.Response{StatusCode: 502},
				err:  nil,
			},
			want:    true,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DefaultRetryPolicy(tt.args.ctx, tt.args.resp, tt.args.err)
			if (err != nil) != tt.wantErr {
				t.Errorf("DefaultRetryPolicy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DefaultRetryPolicy() got = %v, want %v", got, tt.want)
			}
		})
	}
}
