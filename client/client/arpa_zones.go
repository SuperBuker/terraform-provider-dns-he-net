package client

import (
	"context"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/utils"
)

// GetArpaZones retrieves all ARPA zones from the API and returns them in a slice
func (c *Client) GetArpaZones(ctx context.Context) ([]models.Zone, error) {
	resp, err := c.client.R().
		SetContext(ctx).
		SetResult([]ArpaZone{}).
		Get(endpoint)

	if err != nil {
		return nil, err
	}

	arpa_zones, ok := resp.Result().([]models.Zone)
	if !ok {
		return nil, utils.NewErrCasting([]models.Zone{}, resp.Result())
	}

	return arpa_zones, nil
}
