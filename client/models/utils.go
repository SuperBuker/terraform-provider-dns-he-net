package models

import (
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/utils"

	"fmt"
)

var b2s = map[bool]string{false: "0", true: "1"}

func toString(n interface{}) string {
	if utils.IsNil(n) {
		return ""
	}

	switch v := n.(type) {
	case *uint:
		return fmt.Sprint(*v)
	case *uint8:
		return fmt.Sprint(*v)
	case *uint16:
		return fmt.Sprint(*v)
	case *uint32:
		return fmt.Sprint(*v)
	case *uint64:
		return fmt.Sprint(*v)
	default:
		return ""
	}
}
