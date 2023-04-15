<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# status

```go
import "github.com/SuperBuker/terraform-provider-dns-he-net/client/status"
```

## Index

- [func Check(doc *html.Node) error](<#func-check>)
- [func fromAuthStatus(status auth.Status) error](<#func-fromauthstatus>)
- [func fromIssue(issue string) error](<#func-fromissue>)
- [type ErrAuth](<#type-errauth>)
  - [func (e *ErrAuth) Error() string](<#func-errauth-error>)
- [type ErrHeNet](<#type-errhenet>)
  - [func (e *ErrHeNet) Error() string](<#func-errhenet-error>)
- [type ErrNoAuth](<#type-errnoauth>)
  - [func (e *ErrNoAuth) Error() string](<#func-errnoauth-error>)
  - [func (e *ErrNoAuth) Unwrap() []error](<#func-errnoauth-unwrap>)
- [type ErrOTPAuth](<#type-errotpauth>)
  - [func (e *ErrOTPAuth) Error() string](<#func-errotpauth-error>)
  - [func (e *ErrOTPAuth) Unwrap() []error](<#func-errotpauth-unwrap>)
- [type ErrUnknownAuth](<#type-errunknownauth>)
  - [func (e *ErrUnknownAuth) Error() string](<#func-errunknownauth-error>)
  - [func (e *ErrUnknownAuth) Unwrap() []error](<#func-errunknownauth-unwrap>)


## func [Check](<https://github.com/SuperBuker/terraform-provider-dns-he-net/tree/master/common/client/status/blob/master/client/status/status.go#L13>)

```go
func Check(doc *html.Node) error
```

Check checks all possible errors in the response. \- If the user is not fully logged in. \- If there are other contained errors.

## func [fromAuthStatus](<https://github.com/SuperBuker/terraform-provider-dns-he-net/tree/master/common/client/status/blob/master/client/status/parsers.go#L11>)

```go
func fromAuthStatus(status auth.Status) error
```

fromAuthStatus returns an error asssociated to the auth status.

## func [fromIssue](<https://github.com/SuperBuker/terraform-provider-dns-he-net/tree/master/common/client/status/blob/master/client/status/parsers.go#L29>)

```go
func fromIssue(issue string) error
```

fromIssue parses the errors in the response and returns them as &ErrHeNet\{\}.

## type [ErrAuth](<https://github.com/SuperBuker/terraform-provider-dns-he-net/tree/master/common/client/status/blob/master/client/status/errors.go#L4>)

ErrAuth is an error that is returned when authentication fails.

```go
type ErrAuth struct{}
```

### func \(\*ErrAuth\) [Error](<https://github.com/SuperBuker/terraform-provider-dns-he-net/tree/master/common/client/status/blob/master/client/status/errors.go#L6>)

```go
func (e *ErrAuth) Error() string
```

## type [ErrHeNet](<https://github.com/SuperBuker/terraform-provider-dns-he-net/tree/master/common/client/status/blob/master/client/status/errors.go#L60-L62>)

ErrHeNet is an error returned when dns.he.net returns an error. It contains the error message returned in the HTML response.

```go
type ErrHeNet struct {
    error string
}
```

### func \(\*ErrHeNet\) [Error](<https://github.com/SuperBuker/terraform-provider-dns-he-net/tree/master/common/client/status/blob/master/client/status/errors.go#L64>)

```go
func (e *ErrHeNet) Error() string
```

## type [ErrNoAuth](<https://github.com/SuperBuker/terraform-provider-dns-he-net/tree/master/common/client/status/blob/master/client/status/errors.go#L13>)

ErrNoAuth is an error returned when the user is not authenticated. It is used when dns.he.net returns that the client is not authenticated. It includes two wrapped errors, ErrAuth and ErrHeNet.

```go
type ErrNoAuth struct{}
```

### func \(\*ErrNoAuth\) [Error](<https://github.com/SuperBuker/terraform-provider-dns-he-net/tree/master/common/client/status/blob/master/client/status/errors.go#L15>)

```go
func (e *ErrNoAuth) Error() string
```

### func \(\*ErrNoAuth\) [Unwrap](<https://github.com/SuperBuker/terraform-provider-dns-he-net/tree/master/common/client/status/blob/master/client/status/errors.go#L19>)

```go
func (e *ErrNoAuth) Unwrap() []error
```

## type [ErrOTPAuth](<https://github.com/SuperBuker/terraform-provider-dns-he-net/tree/master/common/client/status/blob/master/client/status/errors.go#L29>)

ErrOTPAuth is an error returned when the user is not fully authenticated. It is used when dns.he.net returns that the client lacks OTP authentication. It includes two wrapped errors, ErrAuth and ErrHeNet.

```go
type ErrOTPAuth struct{}
```

### func \(\*ErrOTPAuth\) [Error](<https://github.com/SuperBuker/terraform-provider-dns-he-net/tree/master/common/client/status/blob/master/client/status/errors.go#L31>)

```go
func (e *ErrOTPAuth) Error() string
```

### func \(\*ErrOTPAuth\) [Unwrap](<https://github.com/SuperBuker/terraform-provider-dns-he-net/tree/master/common/client/status/blob/master/client/status/errors.go#L35>)

```go
func (e *ErrOTPAuth) Unwrap() []error
```

## type [ErrUnknownAuth](<https://github.com/SuperBuker/terraform-provider-dns-he-net/tree/master/common/client/status/blob/master/client/status/errors.go#L45>)

ErrUnknownAuth is an error returned when it's not possible to determine the authentication status. It includes two wrapped errors, ErrAuth and ErrHeNet.

```go
type ErrUnknownAuth struct{}
```

### func \(\*ErrUnknownAuth\) [Error](<https://github.com/SuperBuker/terraform-provider-dns-he-net/tree/master/common/client/status/blob/master/client/status/errors.go#L47>)

```go
func (e *ErrUnknownAuth) Error() string
```

### func \(\*ErrUnknownAuth\) [Unwrap](<https://github.com/SuperBuker/terraform-provider-dns-he-net/tree/master/common/client/status/blob/master/client/status/errors.go#L51>)

```go
func (e *ErrUnknownAuth) Unwrap() []error
```



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)