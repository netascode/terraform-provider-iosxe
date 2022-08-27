//go:build testAll

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIosxeService(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxeServiceConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxe_service.test", "pad", "true"),
					resource.TestCheckResourceAttr("iosxe_service.test", "password_encryption", "true"),
					resource.TestCheckResourceAttr("iosxe_service.test", "password_recovery", "true"),
					resource.TestCheckResourceAttr("iosxe_service.test", "timestamps", "true"),
					resource.TestCheckResourceAttr("iosxe_service.test", "timestamps_debug", "true"),
					resource.TestCheckResourceAttr("iosxe_service.test", "timestamps_debug_datetime", "true"),
					resource.TestCheckResourceAttr("iosxe_service.test", "timestamps_debug_datetime_msec", "true"),
					resource.TestCheckResourceAttr("iosxe_service.test", "timestamps_debug_datetime_localtime", "true"),
					resource.TestCheckResourceAttr("iosxe_service.test", "timestamps_debug_datetime_show_timezone", "true"),
					resource.TestCheckResourceAttr("iosxe_service.test", "timestamps_debug_datetime_year", "true"),
					resource.TestCheckResourceAttr("iosxe_service.test", "timestamps_debug_uptime", "true"),
					resource.TestCheckResourceAttr("iosxe_service.test", "timestamps_log", "true"),
					resource.TestCheckResourceAttr("iosxe_service.test", "timestamps_log_datetime", "true"),
					resource.TestCheckResourceAttr("iosxe_service.test", "timestamps_log_datetime_msec", "true"),
					resource.TestCheckResourceAttr("iosxe_service.test", "timestamps_log_datetime_localtime", "true"),
					resource.TestCheckResourceAttr("iosxe_service.test", "timestamps_log_datetime_show_timezone", "true"),
					resource.TestCheckResourceAttr("iosxe_service.test", "timestamps_log_datetime_year", "true"),
					resource.TestCheckResourceAttr("iosxe_service.test", "timestamps_log_uptime", "true"),
					resource.TestCheckResourceAttr("iosxe_service.test", "dhcp", "true"),
					resource.TestCheckResourceAttr("iosxe_service.test", "tcp_keepalives_in", "true"),
					resource.TestCheckResourceAttr("iosxe_service.test", "tcp_keepalives_out", "true"),
				),
			},
			{
				ResourceName:  "iosxe_service.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XE-native:native/service",
			},
		},
	})
}

func testAccIosxeServiceConfig_minimum() string {
	return `
	resource "iosxe_service" "test" {
	}
	`
}

func testAccIosxeServiceConfig_all() string {
	return `
	resource "iosxe_service" "test" {
		pad = true
		password_encryption = true
		password_recovery = true
		timestamps = true
		timestamps_debug = true
		timestamps_debug_datetime = true
		timestamps_debug_datetime_msec = true
		timestamps_debug_datetime_localtime = true
		timestamps_debug_datetime_show_timezone = true
		timestamps_debug_datetime_year = true
		timestamps_debug_uptime = true
		timestamps_log = true
		timestamps_log_datetime = true
		timestamps_log_datetime_msec = true
		timestamps_log_datetime_localtime = true
		timestamps_log_datetime_show_timezone = true
		timestamps_log_datetime_year = true
		timestamps_log_uptime = true
		dhcp = true
		tcp_keepalives_in = true
		tcp_keepalives_out = true
	}
	`
}
