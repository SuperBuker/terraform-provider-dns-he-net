package parsers

import (
	"errors"
	"fmt"
	"net"
	"strings"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
)

// errParsingNode returns a tailored ErrParsing error.
func errParsingNode(path, field string, err error) error {
	return &ErrParsing{
		fmt.Sprintf("%s // %s", path, field),
		err,
	}
}

// subnetToArpaZone converts a subnet string to its corresponding ARPA zone name.
func subnetToArpaZone(subnet string) (string, error) {
	// Example 2001:470:1f13:1::/64 to 1.0.0.0.3.1.f.1.0.7.4.0.1.0.0.2.ip6.arpa
	// TODO: This function can be improved with regexp

	var nameParts []string
	// Remove the /XX part
	subnet = strings.Split(subnet, "/")[0]
	// Expand the IPv6 address
	ip := net.ParseIP(subnet)
	if ip == nil {
		return "", errors.New("invalid IP")
	}
	ip = ip.To16()
	for i := len(ip) - 1; i >= 0; i-- {
		b := ip[i]
		nameParts = append(nameParts, fmt.Sprintf("%x", b&0x0F))
		nameParts = append(nameParts, fmt.Sprintf("%x", (b>>4)&0x0F))

		// Check if nameParts slice is [0, 0, 0, 0] to avoid adding unnecessary zeros
		if len(nameParts) == 4 {
			zeros := true
			for j := range 4 {
				if nameParts[j] != "0" {
					zeros = false
					break
				}
			}

			if zeros {
				nameParts = nameParts[:0] // reset slice
			}
		}
	}

	if len(nameParts) == 0 {
		nameParts = append(nameParts, "0")
	}

	return strings.Join(nameParts, ".") + ".ip6.arpa", nil
}

// NetworkPrefixToArpaZone converts a delegated prefix to its corresponding ARPA zone.
func NetworkPrefixToArpaZone(prefix models.NetworkPrefix) (models.Zone, error) {
	arpaName, err := subnetToArpaZone(prefix.Value)
	if err != nil {
		return models.Zone{}, err
	}
	return models.Zone{
		ID:   prefix.ID,
		Name: arpaName,
	}, nil
}
