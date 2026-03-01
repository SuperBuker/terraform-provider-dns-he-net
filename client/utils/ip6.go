package utils

import (
	"errors"
	"fmt"
	"net"
	"strings"
)

func IPv6AddrToPTR(ip net.IP, zeroed bool) string {
	// Example 2001:470:1f13:1:: to 1.0.0.0.3.1.f.1.0.7.4.0.1.0.0.2.ip6.arpa
	var nameParts []string
	ip = ip.To16()
	for i := len(ip) - 1; i >= 0; i-- {
		b := ip[i]
		nameParts = append(nameParts, fmt.Sprintf("%x", b&0x0F))
		nameParts = append(nameParts, fmt.Sprintf("%x", (b>>4)&0x0F))

		// Check if nameParts slice is [0, 0, 0, 0] to avoid adding unnecessary zeros
		if !zeroed && len(nameParts) == 4 {
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

	if !zeroed && len(nameParts) == 0 {
		nameParts = append(nameParts, "0")
	}

	return strings.Join(nameParts, ".") + ".ip6.arpa"
}

// IPv6SubnetToArpaZone converts a subnet string to its corresponding ARPA zone name.
func IPv6SubnetToArpaZone(subnet string) (string, error) {
	ipNet, err := ParseIPNet(subnet)
	if err != nil {
		return "", fmt.Errorf("invalid subnet: %w", err)
	}

	if ipNet.IP.To4() != nil {
		return "", errors.New("IPv4 subnets are not supported for ARPA zone conversion")
	}

	return IPv6AddrToPTR(ipNet.IP, false), nil
}
