package client

import (
	"context"
	"fmt"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/client/filters"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/client/params"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/utils"
)

// GetDomainZones retrieves all domain zones from the API and returns them in a slice
func (c *Client) GetDomainZones(ctx context.Context) ([]models.Zone, error) {
	resp, err := c.client.R().
		SetContext(ctx).
		SetResult([]DomainZone{}).
		Get(endpoint)

	if err != nil {
		return nil, err
	}

	domains, ok := resp.Result().([]models.Zone)
	if !ok {
		return nil, utils.NewErrCasting([]models.Zone{}, resp.Result())
	}

	return domains, nil
}

// CreateDomainZone creates a new domain zone, then returns it, or an error.
func (c *Client) CreateDomainZone(ctx context.Context, name string) (models.Zone, error) {
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

	domains, ok := resp.Result().([]models.Zone)
	if !ok {
		return models.Zone{}, utils.NewErrCasting([]models.Zone{}, resp.Result())
	}

	domain, ok := filters.LatestZone(domains)
	if !ok {
		return models.Zone{}, &ErrItemNotFound{Resource: "domain"} // TODO: to improve
	}

	return domain, nil
}

// DeleteDomainZone deletes a domain zone, returns an error.
func (c *Client) DeleteDomainZone(ctx context.Context, domain models.Zone) error {
	form := map[string]string{
		"account":   c.account,
		"delete_id": fmt.Sprint(domain.ID),
	}
	params.ZoneDelete(form)

	_, err := c.client.R().
		SetFormData(form).
		SetContext(ctx).
		Post(endpoint)

	return err
}
