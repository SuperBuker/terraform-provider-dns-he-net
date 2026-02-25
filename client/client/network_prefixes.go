package client

import (
	"context"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/utils"
)

// GetNetworkPrefixes retrieves all prefixes from the API and returns them in a slice
func (c *Client) GetNetworkPrefixes(ctx context.Context) ([]models.NetworkPrefix, error) {
	resp, err := c.client.R().
		SetContext(ctx).
		SetResult([]models.NetworkPrefix{}).
		Get(endpoint)

	if err != nil {
		return nil, err
	}

	prefixes, ok := resp.Result().([]models.NetworkPrefix)
	if !ok {
		return nil, utils.NewErrCasting([]models.NetworkPrefix{}, resp.Result())
	}

	return prefixes, nil
}
