<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# client

```go
import "github.com/SuperBuker/terraform-provider-dns-he-net/client"
```

## Index

- [Variables](<#variables>)


## Variables

```go
var CookieStore = struct {
    Dummy     auth.AuthStore
    Simple    auth.AuthStore
    Encrypted auth.AuthStore
}{
    auth.Dummy,
    auth.Simple,
    auth.Encrypted,
}
```

```go
var NewAuth = auth.NewAuth
```

```go
var NewClient = client.NewClient
```

```go
var With = struct {
    Debug     func() client.Option
    Proxy     func(string) client.Option
    UserAgent func(string) client.Option
    Options   func(...client.Option) client.Options
}{
    client.WithDebug,
    client.WithProxy,
    client.WithUserAgent,
    func(options ...client.Option) client.Options { return options },
}
```



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
