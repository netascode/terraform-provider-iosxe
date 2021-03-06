// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIosxeUsername(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxeUsernameConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxe_username.test", "name", "user1"),
					resource.TestCheckResourceAttr("iosxe_username.test", "privilege", "15"),
					resource.TestCheckResourceAttr("iosxe_username.test", "description", "User1 description"),
					resource.TestCheckResourceAttr("iosxe_username.test", "password_encryption", "0"),
					resource.TestCheckResourceAttr("iosxe_username.test", "password", "MyPassword"),
				),
			},
			{
				ResourceName:  "iosxe_username.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XE-native:native/username=user1",
			},
		},
	})
}

func testAccIosxeUsernameConfig_minimum() string {
	return `
	resource "iosxe_username" "test" {
		name = "user1"
	}
	`
}

func testAccIosxeUsernameConfig_all() string {
	return `
	resource "iosxe_username" "test" {
		name = "user1"
		privilege = 15
		description = "User1 description"
		password_encryption = "0"
		password = "MyPassword"
	}
	`
}
