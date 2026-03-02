package client

import (
	"context"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/utils"
)

// GetAllZones retrieves all Domain and ARPA zones from the API and returns them in a slice
func (c *Client) GetAllZones(ctx context.Context) ([]models.Zone, error) {
	resp, err := c.client.R().
		SetContext(ctx).
		SetResult([]GenericZone{}).
		Get(endpoint)

	if err != nil {
		return nil, err
	}

	zones, ok := resp.Result().([]models.Zone)
	if !ok {
		return nil, utils.NewErrCasting([]models.Zone{}, resp.Result())
	}

	return zones, nil
}
