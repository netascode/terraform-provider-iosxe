// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceIosxeClock(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxeClockConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.iosxe_clock.test", "calendar_valid", "true"),
					resource.TestCheckResourceAttr("data.iosxe_clock.test", "summer_time_zone", "CET"),
					resource.TestCheckResourceAttr("data.iosxe_clock.test", "summer_time_recurring", "true"),
					resource.TestCheckResourceAttr("data.iosxe_clock.test", "summer_time_recurring_start_week", "1"),
					resource.TestCheckResourceAttr("data.iosxe_clock.test", "summer_time_recurring_start_weekday", "Mon"),
					resource.TestCheckResourceAttr("data.iosxe_clock.test", "summer_time_recurring_start_month", "Jan"),
					resource.TestCheckResourceAttr("data.iosxe_clock.test", "summer_time_recurring_start_time", "00:00"),
					resource.TestCheckResourceAttr("data.iosxe_clock.test", "summer_time_recurring_end_week", "1"),
					resource.TestCheckResourceAttr("data.iosxe_clock.test", "summer_time_recurring_end_weekday", "Mon"),
					resource.TestCheckResourceAttr("data.iosxe_clock.test", "summer_time_recurring_end_month", "Dec"),
					resource.TestCheckResourceAttr("data.iosxe_clock.test", "summer_time_recurring_end_time", "00:00"),
					resource.TestCheckResourceAttr("data.iosxe_clock.test", "summer_time_recurring_offset", "60"),
				),
			},
		},
	})
}

const testAccDataSourceIosxeClockConfig = `

resource "iosxe_clock" "test" {
	delete_mode = "attributes"
	calendar_valid = true
	summer_time_zone = "CET"
	summer_time_recurring = true
	summer_time_recurring_start_week = "1"
	summer_time_recurring_start_weekday = "Mon"
	summer_time_recurring_start_month = "Jan"
	summer_time_recurring_start_time = "00:00"
	summer_time_recurring_end_week = "1"
	summer_time_recurring_end_weekday = "Mon"
	summer_time_recurring_end_month = "Dec"
	summer_time_recurring_end_time = "00:00"
	summer_time_recurring_offset = 60
}

data "iosxe_clock" "test" {
	depends_on = [iosxe_clock.test]
}
`
