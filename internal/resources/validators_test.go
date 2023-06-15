package resources

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidators(t *testing.T) {
	assert.True(t, domainRegexp.MatchString("example.com"))

	assert.True(t, ipv4Regexp.MatchString("0.0.0.0"))

	assert.True(t, ipv6Regexp.MatchString("::1"))

	assert.True(t, afsdbRegexp.MatchString("1 example.com"))

	assert.True(t, locRegexp.MatchString("1 2 3.000 N 4 5 6.000 E 7.00m 8.00m 9.00m 10.00m"))

	assert.True(t, spfRegexp.MatchString(`"v=spf1 ~all"`))

	assert.True(t, srvRegexp.MatchString("_sip._tcp.example.com"))

	assert.True(t, sshfpRegexp.MatchString("1 1 123456789abcdef67890123456789abcdef67890"))

	assert.True(t, txtRegexp.MatchString(`"hello world"`))
}

func TestDomainRegexp(t *testing.T) {
	matrix := map[string]bool{
		"example.com":              true,
		"example.com.":             false,
		"example":                  false,
		"example.":                 false,
		"example..com":             false,
		"example.com. ":            false,
		" example.com":             false,
		"example. com":             false,
		"example-1620.example.com": true,
		"-example.com.":            false,
		".com":                     false,
		"123456789012345678901234567890123456789012345678901234567890123.com":   true,
		"1234567890123456789012345678901234567890123456789012345678901234.com":  false,
		"a.123456789012345678901234567890123456789012345678901234567890123.com": true,
		"ex--ample.ex__ample.example.com":                                       true,
	}

	for domain, expected := range matrix {
		assert.Equal(t, expected, domainRegexp.MatchString(domain))
	}
}
