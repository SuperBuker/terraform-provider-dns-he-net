package utils

import (
	"fmt"
	"net"
	"strings"
)

func IPv4AddrToPTR(ip net.IP) string {
	// Example 192.0.2.1 to 1.2.0.192.in-addr.arpa
	var nameParts []string
	ip = ip.To4()
	for i := len(ip) - 1; i >= 0; i-- {
		nameParts = append(nameParts, fmt.Sprintf("%d", ip[i]))
	}

	return strings.Join(nameParts, ".") + ".in-addr.arpa"
}
