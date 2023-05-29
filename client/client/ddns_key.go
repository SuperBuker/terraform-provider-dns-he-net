package client

import (
	"context"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/client/params"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/utils"
)

// SetDDNSKey creates or updates a ddns domain key, then returns the result message or an error.
func (c *Client) SetDDNSKey(ctx context.Context, dk models.DDNSKey) (string, error) {
	form := dk.Serialise()

	params.DDNSKeySet(form)

	resp, err := c.client.R().
		SetFormData(form).
		SetQueryParams(getRecordsParams(dk.GetZoneID())). // Yes, this is correct.
		SetContext(ctx).
		SetResult(models.StatusMessage{}).
		Post(endpoint)

	if err != nil {
		return "", err
	}

	statusMsg, ok := resp.Result().(models.StatusMessage)
	if !ok {
		return "", utils.NewErrCasting(models.StatusMessage{}, resp.Result())
	}

	return statusMsg.Data, nil
}
