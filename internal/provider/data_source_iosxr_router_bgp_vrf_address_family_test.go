// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIosxrRouterBGPVRFAddressFamily(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrRouterBGPVRFAddressFamilyPrerequisitesConfig + testAccDataSourceIosxrRouterBGPVRFAddressFamilyConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf_address_family.test", "maximum_paths_ebgp_multipath", "10"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf_address_family.test", "maximum_paths_ibgp_multipath", "10"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf_address_family.test", "label_mode_per_ce", "false"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf_address_family.test", "label_mode_per_vrf", "false"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf_address_family.test", "redistribute_connected", "true"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf_address_family.test", "redistribute_connected_metric", "10"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf_address_family.test", "redistribute_static", "true"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf_address_family.test", "redistribute_static_metric", "10"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf_address_family.test", "aggregate_addresses.0.address", "10.0.0.0"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf_address_family.test", "aggregate_addresses.0.masklength", "8"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf_address_family.test", "aggregate_addresses.0.as_set", "false"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf_address_family.test", "aggregate_addresses.0.as_confed_set", "false"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf_address_family.test", "aggregate_addresses.0.summary_only", "false"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf_address_family.test", "networks.0.address", "10.1.0.0"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf_address_family.test", "networks.0.masklength", "16"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf_address_family.test", "redistribute_ospf.0.router_tag", "OSPF1"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf_address_family.test", "redistribute_ospf.0.match_internal", "true"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf_address_family.test", "redistribute_ospf.0.match_internal_external", "true"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf_address_family.test", "redistribute_ospf.0.match_internal_nssa_external", "false"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf_address_family.test", "redistribute_ospf.0.match_external", "false"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf_address_family.test", "redistribute_ospf.0.match_external_nssa_external", "false"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf_address_family.test", "redistribute_ospf.0.match_nssa_external", "false"),
					resource.TestCheckResourceAttr("data.iosxr_router_bgp_vrf_address_family.test", "redistribute_ospf.0.metric", "100"),
				),
			},
		},
	})
}

const testAccDataSourceIosxrRouterBGPVRFAddressFamilyPrerequisitesConfig = `
resource "iosxr_gnmi" "PreReq0" {
	path = "Cisco-IOS-XR-um-vrf-cfg:/vrfs/vrf[vrf-name=VRF1]"
	attributes = {
	}
}

resource "iosxr_gnmi" "PreReq1" {
	path = "Cisco-IOS-XR-um-vrf-cfg:/vrfs/vrf[vrf-name=VRF1]/Cisco-IOS-XR-um-router-bgp-cfg:rd/Cisco-IOS-XR-um-router-bgp-cfg:two-byte-as"
	attributes = {
		as-number = "1"
		index = "1"
	}
	depends_on = [iosxr_gnmi.PreReq0, ]
}

resource "iosxr_gnmi" "PreReq2" {
	path = "Cisco-IOS-XR-um-router-bgp-cfg:/router/bgp/as[as-number=65001]"
	attributes = {
	}
}

resource "iosxr_gnmi" "PreReq3" {
	path = "Cisco-IOS-XR-um-router-bgp-cfg:/router/bgp/as[as-number=65001]/address-families/address-family[af-name=vpnv4-unicast]"
	attributes = {
	}
	depends_on = [iosxr_gnmi.PreReq1, iosxr_gnmi.PreReq2, ]
}

`

const testAccDataSourceIosxrRouterBGPVRFAddressFamilyConfig = `

resource "iosxr_router_bgp_vrf_address_family" "test" {
	as_number = "65001"
	vrf_name = "VRF1"
	af_name = "ipv4-unicast"
	maximum_paths_ebgp_multipath = 10
	maximum_paths_ibgp_multipath = 10
	label_mode_per_ce = false
	label_mode_per_vrf = false
	redistribute_connected = true
	redistribute_connected_metric = 10
	redistribute_static = true
	redistribute_static_metric = 10
	aggregate_addresses = [{
		address = "10.0.0.0"
		masklength = 8
		as_set = false
		as_confed_set = false
		summary_only = false
	}]
	networks = [{
		address = "10.1.0.0"
		masklength = 16
	}]
	redistribute_ospf = [{
		router_tag = "OSPF1"
		match_internal = true
		match_internal_external = true
		match_internal_nssa_external = false
		match_external = false
		match_external_nssa_external = false
		match_nssa_external = false
		metric = 100
	}]
	depends_on = [iosxr_gnmi.PreReq0, iosxr_gnmi.PreReq1, iosxr_gnmi.PreReq2, iosxr_gnmi.PreReq3, ]
}

data "iosxr_router_bgp_vrf_address_family" "test" {
	as_number = "65001"
	vrf_name = "VRF1"
	af_name = "ipv4-unicast"
	depends_on = [iosxr_router_bgp_vrf_address_family.test]
}
`
