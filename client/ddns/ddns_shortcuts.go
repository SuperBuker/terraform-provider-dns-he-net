package ddns

import (
	"context"

	"github.com/go-resty/resty/v2"
)

func UpdateIP(ctx context.Context, hostname, password, myip string) (bool, error) {
	return New(resty.New()).UpdateIP(ctx, hostname, password, myip)
}

func UpdateTXT(ctx context.Context, hostname, password, txt string) (bool, error) {
	return New(resty.New()).UpdateTXT(ctx, hostname, password, txt)
}

func CheckAuth(ctx context.Context, hostname, password string) (bool, error) {
	return New(resty.New()).CheckAuth(ctx, hostname, password)
}
