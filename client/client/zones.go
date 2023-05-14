package client

import (
	"context"
	"fmt"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/client/filters"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/client/params"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/utils"
)

// GetZones retrieves all zones from the API and returns them in a slice
func (c *Client) GetZones(ctx context.Context) ([]models.Zone, error) {
	resp, err := c.client.R().
		SetContext(ctx).
		SetResult([]models.Zone{}).
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

// CreateZone creates a new zone, then returns it, or an error.
func (c *Client) CreateZone(ctx context.Context, name string) (models.Zone, error) {
	form := map[string]string{
		"add_domain": name,
	}
	params.ZoneCreate(form)

	resp, err := c.client.R().
		SetFormData(form).
		SetContext(ctx).
		SetResult([]models.Zone{}).
		Post(endpoint)

	if err != nil {
		return models.Zone{}, err
	}

	zones, ok := resp.Result().([]models.Zone)
	if !ok {
		return models.Zone{}, utils.NewErrCasting([]models.Zone{}, resp.Result())
	}

	zone, ok := filters.LatestZone(zones)
	if !ok {
		return models.Zone{}, &ErrItemNotFound{Resource: "zone"} // TODO: to improve
	}

	return zone, nil
}

// DeleteZone deletes a zone, returns an error.
func (c *Client) DeleteZone(ctx context.Context, zone models.Zone) error {
	form := map[string]string{
		"account":   c.account,
		"delete_id": fmt.Sprint(zone.ID),
	}
	params.ZoneDelete(form)

	_, err := c.client.R().
		SetFormData(form).
		SetContext(ctx).
		Post(endpoint)

	return err
}
