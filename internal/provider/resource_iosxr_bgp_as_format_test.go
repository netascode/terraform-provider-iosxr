// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIosxrBGPASFormat(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxrBGPASFormatConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxr_bgp_as_format.test", "asdot", "false"),
					resource.TestCheckResourceAttr("iosxr_bgp_as_format.test", "asplain", "true"),
				),
			},
			{
				ResourceName:  "iosxr_bgp_as_format.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XR-um-router-bgp-cfg:/as-format",
			},
		},
	})
}

func testAccIosxrBGPASFormatConfig_minimum() string {
	return `
	resource "iosxr_bgp_as_format" "test" {
	}
	`
}

func testAccIosxrBGPASFormatConfig_all() string {
	return `
	resource "iosxr_bgp_as_format" "test" {
		asdot = false
		asplain = true
	}
	`
}
