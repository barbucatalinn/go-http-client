# client
--
    import "."


## Usage

#### func  DefaultBackoffStrategy

```go
func DefaultBackoffStrategy(_ int) time.Duration
```
DefaultBackoffStrategy always returns 1 second

#### func  DefaultRetryPolicy

```go
func DefaultRetryPolicy(ctx context.Context, resp *http.Response, err error) (bool, error)
```
DefaultRetryPolicy provides a default callback for Client.Retry, which will
retry on connection errors and server errors

#### func  ExponentialBackoffStrategy

```go
func ExponentialBackoffStrategy(i int) time.Duration
```
ExponentialBackoffStrategy returns ever-increasing backoffs by a power of 2

#### func  ExponentialJitterBackoffStrategy

```go
func ExponentialJitterBackoffStrategy(i int) time.Duration
```
ExponentialJitterBackoffStrategy returns ever-increasing backoffs by a power of
2 with +/- 0-33% to prevent synchronized requests

#### func  LinearBackoffStrategy

```go
func LinearBackoffStrategy(i int) time.Duration
```
LinearBackoffStrategy returns increasing durations, each a second longer than
the last

#### func  LinearJitterBackoffStrategy

```go
func LinearJitterBackoffStrategy(i int) time.Duration
```
LinearJitterBackoffStrategy returns increasing durations, each a second longer
than the last with +/- 0-33% to prevent synchronized requests.

#### type BackoffStrategy

```go
type BackoffStrategy func(attemptNum int) time.Duration
```

BackoffStrategy specifies a strategy for how long to wait between retries

#### type BaseClient

```go
type BaseClient struct {
}
```

BaseClient wraps the http.Client and exposes all the functionality of the
http.Client but with additional functionality

#### func  NewClient

```go
func NewClient(l *logrus.Logger) *BaseClient
```
NewClient creates a new BaseClient with default values

#### func (*BaseClient) Delete

```go
func (c *BaseClient) Delete(ctx context.Context, url string) (*Response, error)
```
Delete provides the functionality to send "DELETE" requests

#### func (*BaseClient) Do

```go
func (c *BaseClient) Do(req *Request) (*Response, error)
```
Do wraps calling an HTTP method with retries

#### func (*BaseClient) Get

```go
func (c *BaseClient) Get(ctx context.Context, url string) (*Response, error)
```
Get provides the functionality to send "GET" requests

#### func (*BaseClient) Head

```go
func (c *BaseClient) Head(ctx context.Context, url string) (*Response, error)
```
Head provides the functionality to send "HEAD" requests

#### func (*BaseClient) NewRequest

```go
func (c *BaseClient) NewRequest(ctx context.Context, method, url string, rawBody interface{}) (*Request, error)
```
NewRequest creates a new wrapped request with the provided context

#### func (*BaseClient) Options

```go
func (c *BaseClient) Options(ctx context.Context, url string) (*Response, error)
```
Options provides the functionality to send "OPTIONS" requests

#### func (*BaseClient) Patch

```go
func (c *BaseClient) Patch(ctx context.Context, url, contentType string, body interface{}) (*Response, error)
```
Patch provides the functionality to send "PATCH" requests

#### func (*BaseClient) Post

```go
func (c *BaseClient) Post(ctx context.Context, url, contentType string, body interface{}) (*Response, error)
```
Post provides the functionality to send "POST" requests

#### func (*BaseClient) Put

```go
func (c *BaseClient) Put(ctx context.Context, url, contentType string, body interface{}) (*Response, error)
```
Put provides the functionality to send "PUT" requests

#### func (*BaseClient) WithBackoffStrategy

```go
func (c *BaseClient) WithBackoffStrategy(backoffStrategy BackoffStrategy) *BaseClient
```
WithBackoffStrategy sets the backoff value and returns the BaseClient

#### func (*BaseClient) WithBasicAuth

```go
func (c *BaseClient) WithBasicAuth(username, password string) *BaseClient
```
WithBasicAuth sets the auth object values (basic auth) and returns the
BaseClient

#### func (*BaseClient) WithBearerAuth

```go
func (c *BaseClient) WithBearerAuth(token string) *BaseClient
```
WithBearerAuth sets the auth object values (bearer auth) and returns the
BaseClient

#### func (*BaseClient) WithCustomAuth

```go
func (c *BaseClient) WithCustomAuth(scheme, token string) *BaseClient
```
WithCustomAuth sets the auth object values (custom auth) and returns the
BaseClient

#### func (*BaseClient) WithRetryMax

```go
func (c *BaseClient) WithRetryMax(retryMax int) *BaseClient
```
WithRetryMax sets the RetryMax value and returns the BaseClient

#### func (*BaseClient) WithRetryPolicy

```go
func (c *BaseClient) WithRetryPolicy(retryPolicy RetryPolicy) *BaseClient
```
WithRetryPolicy sets the retry value and returns the BaseClient

#### type DataDump

```go
type DataDump struct {
	RequestDump  []byte
	ResponseDump []byte
}
```

DataDump is a struct containing the request and the response

#### type Request

```go
type Request struct {
	*http.Request
}
```

Request wraps the metadata needed to create HTTP requests

#### func (*Request) SetHeader

```go
func (r *Request) SetHeader(key, value string) *Request
```
SetHeader method is to set a single header key/value pair

#### func (*Request) SetHeaders

```go
func (r *Request) SetHeaders(headers map[string]string) *Request
```
SetHeaders method sets multiple headers key/value pairs

#### type Response

```go
type Response struct {
	RawResponse *http.Response

	DataDump *DataDump
}
```

Response is a wrapper for the response

#### func (*Response) GetBody

```go
func (r *Response) GetBody() ([]byte, error)
```
GetBody returns the body as []byte array

#### func (*Response) GetHeaders

```go
func (r *Response) GetHeaders() http.Header
```
GetHeaders returns the header map of the response

#### func (*Response) GetStatus

```go
func (r *Response) GetStatus() string
```
GetStatus returns the status string of the response

#### func (*Response) GetStatusCode

```go
func (r *Response) GetStatusCode() int
```
GetStatusCode returns the status code of the response

#### func (*Response) GetStringBody

```go
func (r *Response) GetStringBody() (string, error)
```
GetStringBody returns the body as string

#### func (*Response) UnmarshalJSONResponse

```go
func (r *Response) UnmarshalJSONResponse(target interface{}) error
```
UnmarshalJSONResponse unmarshalls the response body into the provided target
object

#### type RetryPolicy

```go
type RetryPolicy func(ctx context.Context, resp *http.Response, err error) (bool, error)
```

RetryPolicy specifies a policy for handling retries
