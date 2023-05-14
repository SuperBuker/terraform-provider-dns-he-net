package client

import (
	"context"
	"fmt"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/client/filters"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/client/params"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/utils"
)

// getRecordsParams returns the generic query parameters for the record
// operations.
func getRecordsParams(zoneID uint) map[string]string {
	return map[string]string{
		"hosted_dns_zoneid":   fmt.Sprint(zoneID),
		"menu":                "edit_zone",
		"hosted_dns_editzone": "",
	}
}

// GetRecords retrieves all records from the API and returns them in a slice.
func (c *Client) GetRecords(ctx context.Context, zoneID uint) ([]models.Record, error) {
	resp, err := c.client.R().
		SetQueryParams(getRecordsParams(zoneID)).
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
	id, idIsSet := record.GetID()
	form := record.Serialise()

	if idIsSet {
		params.RecordUpdate(form)
	} else {
		params.RecordCreate(form)
	}

	resp, err := c.client.R().
		SetFormData(form).
		SetQueryParams(getRecordsParams(record.GetZoneID())).
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
		_record, ok := filters.LatestRecord(records)

		if !ok {
			return nil, &ErrItemNotFound{Resource: "record"} // TODO: to improve
		}

		return _record.ToX() // returns models.RecordX, err
	} else if _record, ok := filters.RecordById(records, id); ok {
		return _record.ToX() // returns models.RecordX, err
	} else {
		return nil, &ErrItemNotFound{Resource: "record"} // TODO: to improve
	}
}

// DeleteRecord deletes a record, returns an error.
func (c *Client) DeleteRecord(ctx context.Context, record models.RecordX) error {
	form := record.Refs()
	params.RecordDelete(form)

	_, err := c.client.R().
		SetFormData(form).
		SetQueryParams(getRecordsParams(record.GetZoneID())).
		SetContext(ctx).
		Post(endpoint)

	return err
}
