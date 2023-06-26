package internal

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/auth"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/logging"
	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/datasources"
	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/resources"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces
var (
	_ provider.Provider = &dnsProvider{}
)

// New is a helper function to simplify provider server and testing implementation.
func New(build BuildFlags, run RunFlags) func() provider.Provider {
	return func() provider.Provider {
		return &dnsProvider{
			Build: build,
			Run:   run,
		}
	}
}

// dnsProvider is the provider implementation.
type dnsProvider struct {
	Build BuildFlags
	Run   RunFlags
}

// dnsProviderModel maps provider schema data to a Go type.
type dnsProviderModel struct {
	Username  types.String `tfsdk:"username"`
	Password  types.String `tfsdk:"password"`
	OTPSecret types.String `tfsdk:"otp_secret"`
	StoreType types.String `tfsdk:"store_type"`
}

// Metadata returns the provider type name.
func (p *dnsProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "dns-he-net"
}

// Schema defines the provider-level schema for configuration data.
func (p *dnsProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description:         "dns.he.net provider configuration",
		MarkdownDescription: "dns.he.net provider configuration",
		Attributes: map[string]schema.Attribute{
			"username": schema.StringAttribute{
				Description:         "dns.he.net account username",
				MarkdownDescription: "dns.he.net account username",
				Required:            true,
			},
			"password": schema.StringAttribute{
				Description:         "dns.he.net account password",
				MarkdownDescription: "dns.he.net account password",
				Required:            true,
				Sensitive:           true,
			},
			"otp_secret": schema.StringAttribute{
				Description:         "dns.he.net OTP secret (optional)",
				MarkdownDescription: "dns.he.net OTP secret (optional)",
				Optional:            true,
				Sensitive:           true,
			},
			"store_type": schema.StringAttribute{
				Description:         "dns.he.net cache store type (optional)",
				MarkdownDescription: "dns.he.net cache store type (optional)",
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("dummy", "simple", "encrypted"),
				},
			},
		},
	}
}

// Configure prepares a dns.he.net "API" client for data sources and resources.
func (p *dnsProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tflog.Info(ctx, "Configuring dns.he.net client")

	// Retrieve provider data from configuration
	var config dnsProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// If practitioner provided a configuration value for any of the
	// attributes, it must be a known value.

	if config.Username.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("username"),
			"Unknown dns.he.net Username",
			"The provider cannot create the dns.he.net API client as there is an unknown configuration value for the dns.he.net API username. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the DHN_USERNAME environment variable.",
		)
	}

	if config.Password.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("password"),
			"Unknown dns.he.net Password",
			"The provider cannot create the dns.he.net API client as there is an unknown configuration value for the dns.he.net API password. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the DHN_PASSWORD environment variable.",
		)
	}

	if config.OTPSecret.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("otp_secret"),
			"Unknown dns.he.net OTP Secret",
			"The provider cannot create the dns.he.net API client as there is an unknown configuration value for the dns.he.net API otp_secret. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the DHN_OTP_SECRET environment variable.",
		)
	}

	if config.StoreType.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("store_type"),
			"Unknown dns.he.net OTP Secret",
			"The provider cannot create the dns.he.net API client as there is an unknown configuration value for the dns.he.net API store_type. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the DHN_STORE_TYPE environment variable.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	// Default values to environment variables, but override
	// with Terraform configuration value if set.

	username := os.Getenv("DHN_USERNAME")
	password := os.Getenv("DHN_PASSWORD")
	otpSecret := os.Getenv("DHN_OTP_SECRET")
	storeType := os.Getenv("DHN_STORE_TYPE")

	if !config.Username.IsNull() {
		username = config.Username.ValueString()
	}

	if !config.Password.IsNull() {
		password = config.Password.ValueString()
	}

	if !config.OTPSecret.IsNull() {
		otpSecret = config.OTPSecret.ValueString()
	}

	if !config.StoreType.IsNull() {
		storeType = config.StoreType.ValueString()
	}

	// If any of the expected configurations are missing, return
	// errors with provider-specific guidance.

	if username == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("username"),
			"Missing dns.he.net Username",
			"The provider cannot create the HashiCups API client as there is a missing or empty value for the dns.he.net API username. "+
				"Set the username value in the configuration or use the DHN_USERNAME environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if password == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("password"),
			"Missing dns.he.net Password",
			"The provider cannot create the dns.he.net API client as there is a missing or empty value for the dns.he.net API password. "+
				"Set the password value in the configuration or use the DHN_PASSWORD environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	var storeMode auth.AuthStore
	if storeType == "" {
		resp.Diagnostics.AddWarning(
			"Missing dns.he.net Store Type",
			`Applying default value: "dummy".`,
		)
	} else {
		switch x := strings.ToLower(storeType); x {
		case "dummy":
			storeMode = auth.Dummy
		case "simple":
			storeMode = auth.Simple
		case "encrypted":
			storeMode = auth.Encrypted
		default:
			resp.Diagnostics.AddAttributeError(
				path.Root("store_type"),
				"Invalid dns.he.net Store Type",
				fmt.Sprintf("Attribute value must be one of: [\"dummy\", \"simple\"m, \"encrypted\"], got: \"%s\"", x),
			)
		}
	}

	if resp.Diagnostics.HasError() {
		return
	}

	ctx = tflog.SetField(ctx, "dhn_username", username)
	ctx = tflog.SetField(ctx, "dhn_password", password)
	ctx = tflog.SetField(ctx, "dhn_otp_secret", otpSecret)
	ctx = tflog.SetField(ctx, "dhn_store_type", storeType)
	ctx = tflog.MaskFieldValuesWithFieldKeys(ctx, "dhn_password")
	ctx = tflog.MaskFieldValuesWithFieldKeys(ctx, "dhn_otp_secret")

	tflog.Debug(ctx, "Creating terrser:one client")

	// Create a new tesser:one client using the configuration values
	auth, err := client.NewAuth(username, password, otpSecret, storeMode)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create dns.he.net API Authenticator",
			"An unexpected error occurred when creating the dns.he.net API client. "+
				"If the error is not clear, please contact the provider developers.\n\n"+
				"dns.he.net Authenticator Error: "+err.Error(),
		)
		return
	}

	// Configure User-Agent
	ua := UserAgentString(ctx, p.Build.Version, req.TerraformVersion)
	options := client.With.Options(client.With.UserAgent(ua))

	// Configure Debug flag
	if p.Run.Debug {
		options = append(options, client.With.Debug())
	}

	// Create a new dns.he.net client using the configuration values
	c, err := client.NewClient(ctx, auth, logging.NewTlog(), options...)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create dns.he.net API Client",
			"An unexpected error occurred when creating the dns.he.net API client. "+
				"If the error is not clear, please contact the provider developers.\n\n"+
				"dns.he.net Client Error: "+err.Error(),
		)
		return
	}

	// Make the HashiCups client available during DataSource and Resource
	// type Configure methods.
	resp.DataSourceData = c
	resp.ResourceData = c

	tflog.Info(ctx, "Configured dns.he.net client", map[string]any{"success": true})
}

// DataSources defines the data sources implemented in the provider.
func (p *dnsProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		datasources.NewAccount,
		datasources.NewZoneIndex,
		datasources.NewZone,
		datasources.NewRecordIndex,
		datasources.NewA,
		datasources.NewAAAA,
		datasources.NewAFSDB,
		datasources.NewALIAS,
		datasources.NewCAA,
		datasources.NewCNAME,
		datasources.NewHINFO,
		datasources.NewLOC,
		datasources.NewMX,
		datasources.NewNAPTR,
		datasources.NewNS,
		datasources.NewPTR,
		datasources.NewRP,
		datasources.NewSOA,
		datasources.NewSPF,
		datasources.NewSRV,
		datasources.NewSSHFP,
		datasources.NewTXT,
	}
}

// Resources defines the resources implemented in the provider.
func (p *dnsProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		resources.NewA,
		resources.NewAAAA,
		resources.NewAFSDB,
		resources.NewALIAS,
		resources.NewCAA,
		resources.NewCNAME,
		resources.NewDDNSKey,
		resources.NewHINFO,
		resources.NewLOC,
		resources.NewMX,
		resources.NewNAPTR,
		resources.NewNS,
		resources.NewPTR,
		resources.NewRP,
		resources.NewSPF,
		resources.NewSRV,
		resources.NewSSHFP,
		resources.NewTXT,
	}
}
