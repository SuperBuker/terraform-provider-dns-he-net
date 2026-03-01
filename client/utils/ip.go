package utils

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

// ParseIPNet parses a subnet string and returns a net.IPNet struct.
// Supports both IPv4 and IPv6 subnets.
func ParseIPNet(subnet string) (net.IPNet, error) {
	// Decompose subnet into IP and mask
	parts := strings.Split(subnet, "/")
	if len(parts) != 2 {
		return net.IPNet{}, fmt.Errorf("invalid subnet format: %s", subnet)
	}

	// Validate IP address
	ip := net.ParseIP(parts[0])
	if ip == nil {
		return net.IPNet{}, fmt.Errorf("invalid IP address: %s", parts[0])
	}

	// Parse mask size
	maskSize, err := strconv.Atoi(parts[1])
	if err != nil {
		return net.IPNet{}, fmt.Errorf("invalid mask: %s", parts[1])
	}

	// Create the appropriate mask based on IP version
	var mask net.IPMask
	if ip.To4() != nil {
		mask = net.CIDRMask(maskSize, 32)
	} else {
		mask = net.CIDRMask(maskSize, 128)
	}

	if mask == nil {
		return net.IPNet{}, fmt.Errorf("invalid mask size: %d", maskSize)
	}

	return net.IPNet{
		IP:   ip,
		Mask: mask,
	}, nil
}
