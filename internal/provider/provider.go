// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/terraform-provider-iosxr/internal/provider/client"
)

// provider satisfies the tfsdk.Provider interface and usually is included
// with all Resource and DataSource implementations.
type provider struct {
	client *client.Client

	// configured is set to true at the end of the Configure method.
	// This can be used in Resource and DataSource implementations to verify
	// that the provider was previously configured.
	configured bool

	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// providerData can be used to store data from the Terraform configuration.
type providerData struct {
	Username types.String         `tfsdk:"username"`
	Password types.String         `tfsdk:"password"`
	Host     types.String         `tfsdk:"host"`
	Devices  []providerDataDevice `tfsdk:"devices"`
}

type providerDataDevice struct {
	Name types.String `tfsdk:"name"`
	Host types.String `tfsdk:"host"`
}

func (p *provider) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"username": {
				MarkdownDescription: "Username for the IOS-XR device. This can also be set as the IOSXR_USERNAME environment variable.",
				Type:                types.StringType,
				Optional:            true,
			},
			"password": {
				MarkdownDescription: "Password for the IOS-XR device. This can also be set as the IOSXR_PASSWORD environment variable.",
				Type:                types.StringType,
				Optional:            true,
				Sensitive:           true,
			},
			"host": {
				MarkdownDescription: "IP or name of the Cisco IOS-XR device. Optionally a port can be added with `:12345`. The default port is `57400`. This can also be set as the IOSXR_HOST environment variable. If no `host` is provided, the `host` of the first device from the `devices` list is being used.",
				Type:                types.StringType,
				Optional:            true,
			},
			"devices": {
				MarkdownDescription: "This can be used to manage a list of devices from a single provider. All devices must use the same credentials. Each resource and data source has an optional attribute named `device`, which can then select a device by its name from this list.",
				Optional:            true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"name": {
						MarkdownDescription: "Device name.",
						Type:                types.StringType,
						Required:            true,
					},
					"host": {
						MarkdownDescription: "IP of the Cisco IOS-XR device.",
						Type:                types.StringType,
						Required:            true,
					},
				}, tfsdk.ListNestedAttributesOptions{}),
			},
		},
	}, nil
}

func (p *provider) Configure(ctx context.Context, req tfsdk.ConfigureProviderRequest, resp *tfsdk.ConfigureProviderResponse) {
	// Retrieve provider data from configuration
	var config providerData
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// User must provide a username to the provider
	var username string
	if config.Username.Unknown {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as username",
		)
		return
	}

	if config.Username.Null {
		username = os.Getenv("IOSXR_USERNAME")
	} else {
		username = config.Username.Value
	}

	if username == "" {
		// Error vs warning - empty value must stop execution
		resp.Diagnostics.AddError(
			"Unable to find username",
			"Username cannot be an empty string",
		)
		return
	}

	// User must provide a password to the provider
	var password string
	if config.Password.Unknown {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as password",
		)
		return
	}

	if config.Password.Null {
		password = os.Getenv("IOSXR_PASSWORD")
	} else {
		password = config.Password.Value
	}

	if password == "" {
		// Error vs warning - empty value must stop execution
		resp.Diagnostics.AddError(
			"Unable to find password",
			"Password cannot be an empty string",
		)
		return
	}

	// User must provide a username to the provider
	var host string
	if config.Host.Unknown {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as host",
		)
		return
	}

	if config.Host.Null {
		host = os.Getenv("IOSXR_HOST")
		if host == "" && len(config.Devices) > 0 {
			host = config.Devices[0].Host.Value
		}
	} else {
		host = config.Host.Value
	}

	if host == "" {
		// Error vs warning - empty value must stop execution
		resp.Diagnostics.AddError(
			"Unable to find host",
			"Host cannot be an empty string",
		)
		return
	}

	client := client.NewClient()

	diags = client.AddTarget(ctx, "", host, username, password)
	resp.Diagnostics.Append(diags...)

	for _, device := range config.Devices {
		diags = client.AddTarget(ctx, device.Name.Value, device.Host.Value, username, password)
		resp.Diagnostics.Append(diags...)
	}

	p.client = &client
	p.configured = true
}

func (p *provider) GetResources(ctx context.Context) (map[string]tfsdk.ResourceType, diag.Diagnostics) {
	return map[string]tfsdk.ResourceType{
		"iosxr_gnmi":                                 resourceGnmiType{},
		"iosxr_bgp_as_format":                        resourceBGPASFormatType{},
		"iosxr_hostname":                             resourceHostnameType{},
		"iosxr_interface":                            resourceInterfaceType{},
		"iosxr_l2vpn":                                resourceL2VPNType{},
		"iosxr_l2vpn_xconnect_group_p2p":             resourceL2VPNXconnectGroupP2PType{},
		"iosxr_mpls_ldp":                             resourceMPLSLDPType{},
		"iosxr_oc_system_config":                     resourceOCSystemConfigType{},
		"iosxr_router_bgp":                           resourceRouterBGPType{},
		"iosxr_router_bgp_address_family":            resourceRouterBGPAddressFamilyType{},
		"iosxr_router_bgp_vrf":                       resourceRouterBGPVRFType{},
		"iosxr_router_bgp_vrf_address_family":        resourceRouterBGPVRFAddressFamilyType{},
		"iosxr_router_isis":                          resourceRouterISISType{},
		"iosxr_router_isis_interface_address_family": resourceRouterISISInterfaceAddressFamilyType{},
		"iosxr_router_ospf":                          resourceRouterOSPFType{},
		"iosxr_router_ospf_area_interface":           resourceRouterOSPFAreaInterfaceType{},
		"iosxr_router_ospf_vrf":                      resourceRouterOSPFVRFType{},
		"iosxr_router_ospf_vrf_area_interface":       resourceRouterOSPFVRFAreaInterfaceType{},
		"iosxr_vrf":                                  resourceVRFType{},
	}, nil
}

func (p *provider) GetDataSources(ctx context.Context) (map[string]tfsdk.DataSourceType, diag.Diagnostics) {
	return map[string]tfsdk.DataSourceType{
		"iosxr_gnmi":                                 dataSourceGnmiType{},
		"iosxr_bgp_as_format":                        dataSourceBGPASFormatType{},
		"iosxr_hostname":                             dataSourceHostnameType{},
		"iosxr_interface":                            dataSourceInterfaceType{},
		"iosxr_l2vpn":                                dataSourceL2VPNType{},
		"iosxr_l2vpn_xconnect_group_p2p":             dataSourceL2VPNXconnectGroupP2PType{},
		"iosxr_mpls_ldp":                             dataSourceMPLSLDPType{},
		"iosxr_oc_system_config":                     dataSourceOCSystemConfigType{},
		"iosxr_router_bgp":                           dataSourceRouterBGPType{},
		"iosxr_router_bgp_address_family":            dataSourceRouterBGPAddressFamilyType{},
		"iosxr_router_bgp_vrf":                       dataSourceRouterBGPVRFType{},
		"iosxr_router_bgp_vrf_address_family":        dataSourceRouterBGPVRFAddressFamilyType{},
		"iosxr_router_isis":                          dataSourceRouterISISType{},
		"iosxr_router_isis_interface_address_family": dataSourceRouterISISInterfaceAddressFamilyType{},
		"iosxr_router_ospf":                          dataSourceRouterOSPFType{},
		"iosxr_router_ospf_area_interface":           dataSourceRouterOSPFAreaInterfaceType{},
		"iosxr_router_ospf_vrf":                      dataSourceRouterOSPFVRFType{},
		"iosxr_router_ospf_vrf_area_interface":       dataSourceRouterOSPFVRFAreaInterfaceType{},
		"iosxr_vrf":                                  dataSourceVRFType{},
	}, nil
}

func New(version string) func() tfsdk.Provider {
	return func() tfsdk.Provider {
		return &provider{
			version: version,
		}
	}
}

// convertProviderType is a helper function for NewResource and NewDataSource
// implementations to associate the concrete provider type. Alternatively,
// this helper can be skipped and the provider type can be directly type
// asserted (e.g. provider: in.(*provider)), however using this can prevent
// potential panics.
func convertProviderType(in tfsdk.Provider) (provider, diag.Diagnostics) {
	var diags diag.Diagnostics

	p, ok := in.(*provider)

	if !ok {
		diags.AddError(
			"Unexpected Provider Instance Type",
			fmt.Sprintf("While creating the data source or resource, an unexpected provider type (%T) was received. This is always a bug in the provider code and should be reported to the provider developers.", p),
		)
		return provider{}, diags
	}

	if p == nil {
		diags.AddError(
			"Unexpected Provider Instance Type",
			"While creating the data source or resource, an unexpected empty provider instance was received. This is always a bug in the provider code and should be reported to the provider developers.",
		)
		return provider{}, diags
	}

	return *p, diags
}
