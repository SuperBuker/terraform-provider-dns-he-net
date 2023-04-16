package client

import (
	"context"
	"fmt"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/client/filters"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/client/params"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/utils"
)

// getRecordsParams returns the genericquery parameters for the record
// operations.
func getRecordsParams(domainId uint) map[string]string {
	return map[string]string{
		"hosted_dns_zoneid":   fmt.Sprint(domainId),
		"menu":                "edit_zone",
		"hosted_dns_editzone": "",
	}
}

// GetRecords retrieves all records from the API and returns them in a slice.
func (c *Client) GetRecords(ctx context.Context, domainId uint) ([]models.Record, error) {
	resp, err := c.client.R().
		SetQueryParams(getRecordsParams(domainId)).
		SetContext(ctx).
		SetResult([]models.Record{}).
		Get(endpoint)

	if err != nil {
		return nil, err
	}

	records, ok := resp.Result().([]models.Record)
	if !ok {
		return nil, utils.NewErrCasting([]models.Record{}, resp.Result())
	}

	return records, nil
}

// SetRecord creates or updates a record, then returns it, or an error.
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

	records, ok := resp.Result().([]models.Record)
	if !ok {
		return nil, utils.NewErrCasting([]models.Record{}, resp.Result())
	}

	if !idIsSet {
		_record, _ := filters.LatestRecord(records)
		return _record.ToX() // err
	} else if _record, ok := filters.RecordById(records, id); ok {
		return _record.ToX() // err
	} else {
		return nil, nil // missing err not found
	}
}

// DeleteRecord deletes a record, returns an error.
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
