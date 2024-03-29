<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# parsers

```go
import "github.com/SuperBuker/terraform-provider-dns-he-net/client/parsers"
```

Parsers contains functions to parse the response from the server.

## Index

- [Constants](<#constants>)
- [func GetAccount(doc *html.Node) (string, error)](<#func-getaccount>)
- [func GetRecords(doc *html.Node) ([]models.Record, error)](<#func-getrecords>)
- [func GetStatusMessage(doc *html.Node) (models.StatusMessage, error)](<#func-getstatusmessage>)
- [func GetZones(doc *html.Node) ([]models.Zone, error)](<#func-getzones>)
- [func LoginStatus(doc *html.Node) auth.Status](<#func-loginstatus>)
- [func ParseError(doc *html.Node) (issues []string)](<#func-parseerror>)
- [func ParseStatus(doc *html.Node) string](<#func-parsestatus>)
- [func errParsingNode(path, field string, err error) error](<#func-errparsingnode>)
- [func parseRecordNode(node *html.Node) (record models.Record, err error)](<#func-parserecordnode>)
- [func parseZoneNode(node *html.Node) (models.Zone, error)](<#func-parsezonenode>)
- [type ErrNotFound](<#type-errnotfound>)
  - [func (e *ErrNotFound) Error() string](<#func-errnotfound-error>)
  - [func (e *ErrNotFound) Unwrap() []error](<#func-errnotfound-unwrap>)
- [type ErrParsing](<#type-errparsing>)
  - [func (e *ErrParsing) Error() string](<#func-errparsing-error>)
  - [func (e *ErrParsing) Unwrap() []error](<#func-errparsing-unwrap>)
- [type loginStatusQuery](<#type-loginstatusquery>)
  - [func getLoginStatusTuples() []loginStatusQuery](<#func-getloginstatustuples>)


## Constants

```go
const (
    // accountQ is the XPath query for the account field.
    accountQ = "//form[@name='remove_domain']/input[@name='account']"

    // loginOkQ is the XPath query for the logout hyperlink.
    loginOkQ = "//a[@id='_tlogout']"

    // loginOtpQ is the XPath query for the OTP form.
    loginOtpQ = "//input[@id='tfacode']"

    // loginNoAuthQ is the XPath query for the login form.
    loginNoAuthQ = "//form[@name='login']"

    // zonesTableQ is the XPath query for the zones table.
    zonesTableQ = "//table[@id='domains_table']"

    // zoneQ is the XPath query for the zone rows.
    zoneQ = zonesTableQ + "/tbody/tr"

    // zoneIDQ is the XPath query for the zone ID within the zone row.
    zoneIDQ = "//td[@style]/img[@name][@value]"

    // recordsTableQ is the XPath query for the records table.
    recordsTableQ = "//div[@id='dns_main_content']/table[@class='generictable']"

    // recordQ is the XPath query for the record rows.
    recordQ = recordsTableQ + "/tbody/tr[@class]"

    // statusQ is the XPath query for the status message.
    statusQ = "//div[@id='dns_status']"

    // errorQ is the XPath query for the error message.
    errorQ = "//div[@id='dns_err']"
)
```

## func [GetAccount](<https://github.com/SuperBuker/terraform-provider-dns-he-net/tree/master/common/client/parsers/blob/master/client/parsers/account.go#L9>)

```go
func GetAccount(doc *html.Node) (string, error)
```

GetAccount returns the account name from the HTML body.

## func [GetRecords](<https://github.com/SuperBuker/terraform-provider-dns-he-net/tree/master/common/client/parsers/blob/master/client/parsers/records.go#L166>)

```go
func GetRecords(doc *html.Node) ([]models.Record, error)
```

GetRecords returns the records from the HTML body.

## func [GetStatusMessage](<https://github.com/SuperBuker/terraform-provider-dns-he-net/tree/master/common/client/parsers/blob/master/client/parsers/status.go#L23>)

```go
func GetStatusMessage(doc *html.Node) (models.StatusMessage, error)
```

GetStatusMessage returns the status message from the HTML body.

## func [GetZones](<https://github.com/SuperBuker/terraform-provider-dns-he-net/tree/master/common/client/parsers/blob/master/client/parsers/zones.go#L28>)

```go
func GetZones(doc *html.Node) ([]models.Zone, error)
```

GetZones returns the zones from the HTML body.

## func [LoginStatus](<https://github.com/SuperBuker/terraform-provider-dns-he-net/tree/master/common/client/parsers/blob/master/client/parsers/auth.go#L26>)

```go
func LoginStatus(doc *html.Node) auth.Status
```

LoginStatus returns the login status from the HTML body.

## func [ParseError](<https://github.com/SuperBuker/terraform-provider-dns-he-net/tree/master/common/client/parsers/blob/master/client/parsers/error.go#L11>)

```go
func ParseError(doc *html.Node) (issues []string)
```

ParseError returns the error message from the HTML body.

## func [ParseStatus](<https://github.com/SuperBuker/terraform-provider-dns-he-net/tree/master/common/client/parsers/blob/master/client/parsers/status.go#L12>)

```go
func ParseStatus(doc *html.Node) string
```

ParseStatus returns the dns status from the HTML body.

## func [errParsingNode](<https://github.com/SuperBuker/terraform-provider-dns-he-net/tree/master/common/client/parsers/blob/master/client/parsers/utils.go#L6>)

```go
func errParsingNode(path, field string, err error) error
```

errParsingNode returns a tailored ErrParsing error.

## func [parseRecordNode](<https://github.com/SuperBuker/terraform-provider-dns-he-net/tree/master/common/client/parsers/blob/master/client/parsers/records.go#L16>)

```go
func parseRecordNode(node *html.Node) (record models.Record, err error)
```

parseRecordNode parses a record node.

## func [parseZoneNode](<https://github.com/SuperBuker/terraform-provider-dns-he-net/tree/master/common/client/parsers/blob/master/client/parsers/zones.go#L13>)

```go
func parseZoneNode(node *html.Node) (models.Zone, error)
```

parseZoneNode parses a zone node.

## type [ErrNotFound](<https://github.com/SuperBuker/terraform-provider-dns-he-net/tree/master/common/client/parsers/blob/master/client/parsers/errors.go#L9-L11>)

ErrNotFound is returned when the element is not found in the document.

```go
type ErrNotFound struct {
    XPath string
}
```

### func \(\*ErrNotFound\) [Error](<https://github.com/SuperBuker/terraform-provider-dns-he-net/tree/master/common/client/parsers/blob/master/client/parsers/errors.go#L13>)

```go
func (e *ErrNotFound) Error() string
```

### func \(\*ErrNotFound\) [Unwrap](<https://github.com/SuperBuker/terraform-provider-dns-he-net/tree/master/common/client/parsers/blob/master/client/parsers/errors.go#L17>)

```go
func (e *ErrNotFound) Unwrap() []error
```

## type [ErrParsing](<https://github.com/SuperBuker/terraform-provider-dns-he-net/tree/master/common/client/parsers/blob/master/client/parsers/errors.go#L24-L27>)

ErrParsing is returned when an error happens when parsing the document.

```go
type ErrParsing struct {
    XPath string
    Err   error
}
```

### func \(\*ErrParsing\) [Error](<https://github.com/SuperBuker/terraform-provider-dns-he-net/tree/master/common/client/parsers/blob/master/client/parsers/errors.go#L29>)

```go
func (e *ErrParsing) Error() string
```

### func \(\*ErrParsing\) [Unwrap](<https://github.com/SuperBuker/terraform-provider-dns-he-net/tree/master/common/client/parsers/blob/master/client/parsers/errors.go#L33>)

```go
func (e *ErrParsing) Unwrap() []error
```

## type [loginStatusQuery](<https://github.com/SuperBuker/terraform-provider-dns-he-net/tree/master/common/client/parsers/blob/master/client/parsers/auth.go#L11-L14>)

loginStatusQuery is the query to find the login status.

```go
type loginStatusQuery struct {
    status auth.Status
    query  string
}
```

### func [getLoginStatusTuples](<https://github.com/SuperBuker/terraform-provider-dns-he-net/tree/master/common/client/parsers/blob/master/client/parsers/auth.go#L17>)

```go
func getLoginStatusTuples() []loginStatusQuery
```

getLoginStatusTuples returns the login status tuples.



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
