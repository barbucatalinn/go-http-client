//go:build !integration
// +build !integration

package client

import (
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_backoffStatusCheck(t *testing.T) {
	t.Parallel()

	type args struct {
		resp *http.Response
	}
	tests := []struct {
		name string
		args args
		want *time.Duration
	}{
		{
			name: "nil response",
			args: args{
				resp: nil,
			},
			want: nil,
		},
		{
			name: "status code 429 & no header",
			args: args{
				resp: &http.Response{
					StatusCode: http.StatusTooManyRequests,
				},
			},
			want: nil,
		},
		{
			name: "status code 503 & no header",
			args: args{
				resp: &http.Response{
					StatusCode: http.StatusServiceUnavailable,
				},
			},
			want: nil,
		},
		{
			name: "status code 429 with header in seconds",
			args: args{
				resp: &http.Response{
					StatusCode: http.StatusTooManyRequests,
					Header: map[string][]string{
						retryAfterHeaderKey: []string{"120"},
					},
				},
			},
			want: func() *time.Duration {
				m := time.Duration(120) * time.Second
				return &m
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := backoffStatusCheck(tt.args.resp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("backoffStatusCheck() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultBackoffStrategy(t *testing.T) {
	t.Parallel()

	type args struct {
		in0 int
	}
	tests := []struct {
		name string
		args args
		want time.Duration
	}{
		{
			name: "default",
			args: args{},
			want: time.Duration(1) * time.Second,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DefaultBackoffStrategy(tt.args.in0); got != tt.want {
				t.Errorf("DefaultBackoffStrategy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExponentialBackoffStrategy(t *testing.T) {
	t.Parallel()

	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want time.Duration
	}{
		{
			name: "1",
			args: args{
				i: 1,
			},
			want: time.Duration(2) * time.Second,
		},
		{
			name: "2",
			args: args{
				i: 2,
			},
			want: time.Duration(4) * time.Second,
		},
		{
			name: "3",
			args: args{
				i: 3,
			},
			want: time.Duration(8) * time.Second,
		},
		{
			name: "4",
			args: args{
				i: 4,
			},
			want: time.Duration(16) * time.Second,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExponentialBackoffStrategy(tt.args.i); got != tt.want {
				t.Errorf("ExponentialBackoffStrategy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinearBackoffStrategy(t *testing.T) {
	t.Parallel()

	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want time.Duration
	}{
		{
			name: "1",
			args: args{
				i: 1,
			},
			want: time.Duration(1) * time.Second,
		},
		{
			name: "2",
			args: args{
				i: 2,
			},
			want: time.Duration(2) * time.Second,
		},
		{
			name: "3",
			args: args{
				i: 3,
			},
			want: time.Duration(3) * time.Second,
		},
		{
			name: "4",
			args: args{
				i: 4,
			},
			want: time.Duration(4) * time.Second,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LinearBackoffStrategy(tt.args.i); got != tt.want {
				t.Errorf("LinearBackoffStrategy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExponentialJitterBackoffStrategy(t *testing.T) {
	t.Parallel()

	result := ExponentialJitterBackoffStrategy(1)
	assert.True(t, result.Seconds() >= float64(1))
	assert.True(t, result.Seconds() <= float64(3))
}

func TestLinearJitterBackoffStrategy(t *testing.T) {
	t.Parallel()

	result := LinearJitterBackoffStrategy(1)
	assert.True(t, result.Seconds() >= 0.667)
	assert.True(t, result.Seconds() <= 1.333)
}

func Test_jitter(t *testing.T) {
	t.Parallel()

	result := jitter(1)
	assert.True(t, result.Seconds() >= 0.667)
	assert.True(t, result.Seconds() <= 1.333)
}
