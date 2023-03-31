package client

import (
	"context"
	"fmt"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/client/filters"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/client/params"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
)

func (c *Client) GetDomains(ctx context.Context) ([]models.Domain, error) {
	// TODO: validate authentication

	resp, err := c.client.R().
		SetContext(ctx).
		SetResult([]models.Domain{}).
		Get(endpoint)

	if err != nil {
		return nil, err
	}

	domains, _ := resp.Result().([]models.Domain)

	return domains, nil
}

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

	domains, _ := resp.Result().([]models.Domain) // TODO: validate

	_domain, _ := filters.LatestDomain(domains)

	return _domain, nil
}

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
