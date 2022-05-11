//go:build ignore

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var models = []string{
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/1771/Cisco-IOS-XE-native.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/1771/cisco-semver.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/1771/ietf-inet-types.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/1771/Cisco-IOS-XE-types.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/1771/Cisco-IOS-XE-features.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/1771/Cisco-IOS-XE-interface-common.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/1771/Cisco-IOS-XE-parser.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/1771/Cisco-IOS-XE-license.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/1771/Cisco-IOS-XE-line.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/1771/Cisco-IOS-XE-logging.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/1771/Cisco-IOS-XE-ip.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/1771/Cisco-IOS-XE-ipv6.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/1771/Cisco-IOS-XE-interfaces.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/1771/Cisco-IOS-XE-isis.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/1771/Cisco-IOS-XE-snmp.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/1771/Cisco-IOS-XE-segment-routing.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/1771/Cisco-IOS-XE-ospf-obsolete.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/1771/Cisco-IOS-XE-vlan.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/1771/Cisco-IOS-XE-vtp.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/1771/Cisco-IOS-XE-device-tracking.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/1771/Cisco-IOS-XE-nd.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/1771/Cisco-IOS-XE-multicast.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/1771/Cisco-IOS-XE-l2vpn.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/1771/Cisco-IOS-XE-ethernet.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/1771/Cisco-IOS-XE-mpls.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/1771/Cisco-IOS-XE-ethernet-oam.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/1771/Cisco-IOS-XE-ethernet-cfm-efp.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/1771/Cisco-IOS-XE-pppoe.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/1771/Cisco-IOS-XE-bgp.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/1771/Cisco-IOS-XE-ospfv3.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/1771/Cisco-IOS-XE-ospf.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/1771/Cisco-IOS-XE-switch.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/1771/ietf-yang-types.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/1771/Cisco-IOS-XE-igmp.yang",
	"https://raw.githubusercontent.com/YangModels/yang/main/vendor/cisco/xe/1771/Cisco-IOS-XE-arp.yang",
}

const (
	modelsPath = "./gen/models/"
)

func main() {
	for _, model := range models {
		f := strings.Split(model, "/")
		path := filepath.Join(modelsPath, f[len(f)-1])
		if _, err := os.Stat(path); err != nil {
			err := downloadModel(path, model)
			if err != nil {
				panic(err)
			}
			fmt.Println("Downloaded model: " + path)
		}
	}
}

func downloadModel(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
