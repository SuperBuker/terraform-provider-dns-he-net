package client

import (
	"context"
	"fmt"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/client/filters"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/client/params"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/utils"
)

// GetDomains retrieves all domains from the API and returns them in a slice
func (c *Client) GetDomains(ctx context.Context) ([]models.Domain, error) {
	resp, err := c.client.R().
		SetContext(ctx).
		SetResult([]models.Domain{}).
		Get(endpoint)

	if err != nil {
		return nil, err
	}

	domains, ok := resp.Result().([]models.Domain)
	if !ok {
		return nil, utils.NewErrCasting([]models.Domain{}, resp.Result())
	}

	return domains, nil
}

// CreateDomain creates a new domain, then returns it, or an error.
func (c *Client) CreateDomain(ctx context.Context, domain string) (models.Domain, error) {
	form := map[string]string{
		"add_domain": domain,
	}
	params.DomainCreate(form)

	resp, err := c.client.R().
		SetFormData(form).
		SetContext(ctx).
		SetResult([]models.Domain{}).
		Post(endpoint)

	if err != nil {
		return models.Domain{}, err
	}

	domains, ok := resp.Result().([]models.Domain)
	if !ok {
		return models.Domain{}, utils.NewErrCasting([]models.Domain{}, resp.Result())
	}

	_domain, ok := filters.LatestDomain(domains)
	if !ok {
		return models.Domain{}, &ErrItemNotFound{Resource: "domain"} // TODO: to improve
	}

	return _domain, nil
}

// DeleteDomain deletes a domain, returns an error.
func (c *Client) DeleteDomain(ctx context.Context, domain models.Domain) error {
	form := map[string]string{
		"account":   c.account,
		"delete_id": fmt.Sprint(domain.Id),
	}
	params.DomainDelete(form)

	_, err := c.client.R().
		SetFormData(form).
		SetContext(ctx).
		Post(endpoint)

	return err
}
