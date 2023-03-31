package client

import (
	"context"
	"fmt"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/client/filters"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/client/params"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
)

func getRecordsParams(domainId uint) map[string]string {
	return map[string]string{
		"hosted_dns_zoneid":   fmt.Sprint(domainId),
		"menu":                "edit_zone",
		"hosted_dns_editzone": "",
	}
}

func (c *Client) GetRecords(ctx context.Context, domainId uint) ([]models.Record, error) {
	resp, err := c.client.R().
		SetQueryParams(getRecordsParams(domainId)).
		SetContext(ctx).
		SetResult([]models.Record{}).
		Get(endpoint)

	if err != nil {
		return nil, err
	}

	records, _ := resp.Result().([]models.Record) // TODO: validate

	return records, nil
}

func (c *Client) SetRecord(ctx context.Context, record models.RecordX) (models.RecordX, error) {
	id, idIsSet := record.GetId()
	form := record.Serialise()

	if idIsSet {
		params.RecordUpdate(form)
	} else {
		params.RecordCreate(form)
	}

	resp, err := c.client.R().
		SetFormData(form).
		SetQueryParams(getRecordsParams(record.GetParentId())).
		SetContext(ctx).
		SetResult([]models.Record{}).
		Post(endpoint)

	if err != nil {
		return nil, err
	}

	records, _ := resp.Result().([]models.Record) // Doubt

	if !idIsSet {
		_record, _ := filters.LatestRecord(records)
		return _record.ToX() // err
	} else if _record, ok := filters.RecordById(records, id); ok {
		return _record.ToX() // err
	} else {
		return nil, nil // missing err not found
	}
}

func (c *Client) DeleteRecord(ctx context.Context, record models.RecordX) error {
	form := record.Refs()
	params.RecordDelete(form)

	_, err := c.client.R().
		SetFormData(form).
		SetQueryParams(getRecordsParams(record.GetParentId())).
		SetContext(ctx).
		Post(endpoint)

	return err
}
