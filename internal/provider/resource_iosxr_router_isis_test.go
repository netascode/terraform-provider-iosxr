// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIosxrRouterISIS(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxrRouterISISConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxr_router_isis.test", "process_id", "P1"),
					resource.TestCheckResourceAttr("iosxr_router_isis.test", "is_type", "level-1"),
					resource.TestCheckResourceAttr("iosxr_router_isis.test", "nets.0.net_id", "49.0001.2222.2222.2222.00"),
					resource.TestCheckResourceAttr("iosxr_router_isis.test", "address_families.0.af_name", "ipv4"),
					resource.TestCheckResourceAttr("iosxr_router_isis.test", "address_families.0.saf_name", "unicast"),
					resource.TestCheckResourceAttr("iosxr_router_isis.test", "address_families.0.mpls_ldp_auto_config", "false"),
					resource.TestCheckResourceAttr("iosxr_router_isis.test", "address_families.0.metric_style_narrow", "false"),
					resource.TestCheckResourceAttr("iosxr_router_isis.test", "address_families.0.metric_style_wide", "true"),
					resource.TestCheckResourceAttr("iosxr_router_isis.test", "address_families.0.metric_style_transition", "false"),
					resource.TestCheckResourceAttr("iosxr_router_isis.test", "address_families.0.router_id_ip_address", "1.2.3.4"),
					resource.TestCheckResourceAttr("iosxr_router_isis.test", "address_families.0.default_information_originate", "true"),
					resource.TestCheckResourceAttr("iosxr_router_isis.test", "interfaces.0.interface_name", "GigabitEthernet0/0/0/1"),
					resource.TestCheckResourceAttr("iosxr_router_isis.test", "interfaces.0.circuit_type", "level-1"),
					resource.TestCheckResourceAttr("iosxr_router_isis.test", "interfaces.0.hello_padding_disable", "true"),
					resource.TestCheckResourceAttr("iosxr_router_isis.test", "interfaces.0.hello_padding_sometimes", "false"),
					resource.TestCheckResourceAttr("iosxr_router_isis.test", "interfaces.0.priority", "10"),
					resource.TestCheckResourceAttr("iosxr_router_isis.test", "interfaces.0.point_to_point", "false"),
					resource.TestCheckResourceAttr("iosxr_router_isis.test", "interfaces.0.passive", "false"),
					resource.TestCheckResourceAttr("iosxr_router_isis.test", "interfaces.0.suppressed", "false"),
					resource.TestCheckResourceAttr("iosxr_router_isis.test", "interfaces.0.shutdown", "false"),
				),
			},
			{
				ResourceName:  "iosxr_router_isis.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XR-um-router-isis-cfg:/router/isis/processes/process[process-id=P1]",
			},
		},
	})
}

func testAccIosxrRouterISISConfig_minimum() string {
	return `
	resource "iosxr_router_isis" "test" {
		process_id = "P1"
	}
	`
}

func testAccIosxrRouterISISConfig_all() string {
	return `
	resource "iosxr_router_isis" "test" {
		process_id = "P1"
		is_type = "level-1"
		nets = [{
			net_id = "49.0001.2222.2222.2222.00"
		}]
		address_families = [{
			af_name = "ipv4"
			saf_name = "unicast"
			mpls_ldp_auto_config = false
			metric_style_narrow = false
			metric_style_wide = true
			metric_style_transition = false
			router_id_ip_address = "1.2.3.4"
			default_information_originate = true
		}]
		interfaces = [{
			interface_name = "GigabitEthernet0/0/0/1"
			circuit_type = "level-1"
			hello_padding_disable = true
			hello_padding_sometimes = false
			priority = 10
			point_to_point = false
			passive = false
			suppressed = false
			shutdown = false
		}]
	}
	`
}
