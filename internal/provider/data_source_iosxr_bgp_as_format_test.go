// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIosxrBGPASFormat(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrBGPASFormatConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.iosxr_bgp_as_format.test", "asdot", "false"),
					resource.TestCheckResourceAttr("data.iosxr_bgp_as_format.test", "asplain", "true"),
				),
			},
		},
	})
}

const testAccDataSourceIosxrBGPASFormatConfig = `

resource "iosxr_bgp_as_format" "test" {
	asdot = false
	asplain = true
}

data "iosxr_bgp_as_format" "test" {
	depends_on = [iosxr_bgp_as_format.test]
}
`
