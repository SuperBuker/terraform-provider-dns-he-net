package filters

import (
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
)

// NetworkPrefixById returns a network prefix by its ID.
// If the network prefix is not found, it returns an empty string and false.
func NetworkPrefixById(prefixes []models.NetworkPrefix, id uint) (models.NetworkPrefix, bool) {
	for _, prefix := range prefixes {
		if prefix.ID == id {
			return prefix, true
		}
	}

	return models.NetworkPrefix{}, false
}

// NetworkPrefixByValue returns a network prefix by its value.
// If the network prefix is not found, it returns an empty string and false.
func NetworkPrefixByValue(prefixes []models.NetworkPrefix, value string) (models.NetworkPrefix, bool) {
	for _, prefix := range prefixes {
		if prefix.Value == value {
			return prefix, true
		}
	}

	return models.NetworkPrefix{}, false
}
