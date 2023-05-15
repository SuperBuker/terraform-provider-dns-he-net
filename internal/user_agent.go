package internal

import (
	"context"
	"os"
	"strings"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/user_agent"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

const uaEnvVar = "TF_APPEND_USER_AGENT"

// UserAgentString returns the User-Agent string to use for HTTP requests.
func UserAgentString(ctx context.Context, tfVersion string) (ua string) {
	ua = user_agent.UserAgentProducts{
		{Name: "HashiCorp Terraform", Version: tfVersion, Comment: "+https://www.terraform.io"},
		{Name: "terraform-provider-dns-he-net", Version: "0.0.1", Comment: "+https://registry.terraform.io/providers/SuperBuker/dns-he-net"}, // TODO: set version
	}.String()

	if add := os.Getenv(uaEnvVar); add != "" {
		add = strings.TrimSpace(add)
		if len(add) > 0 {
			ua += " " + add

			ctx = tflog.SetField(ctx, "user_agent", ua)
			tflog.Debug(ctx, "Using modified User-Agent")
		}
	}

	return
}
