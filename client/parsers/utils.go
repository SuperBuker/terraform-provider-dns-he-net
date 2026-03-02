package parsers

import (
	"fmt"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/utils"
)

// errParsingNode returns a tailored ErrParsing error.
func errParsingNode(path, field string, err error) error {
	return &ErrParsing{
		fmt.Sprintf("%s // %s", path, field),
		err,
	}
}

// NetworkPrefixToArpaZone converts a delegated prefix to its corresponding ARPA zone.
func NetworkPrefixToArpaZone(prefix models.NetworkPrefix) (models.Zone, error) {
	arpaName, err := utils.IPv6SubnetToArpaZone(prefix.Value)
	if err != nil {
		return models.Zone{}, err
	}
	return models.Zone{
		ID:   prefix.ID,
		Name: arpaName,
	}, nil
}
