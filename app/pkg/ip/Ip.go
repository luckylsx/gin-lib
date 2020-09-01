package ip

import (
	"fmt"
	"strconv"
	"strings"
)

// ToInt Return int after IP is converted to int
func ToInt(ip string) (int, bool) {
	ips := strings.Split(ip, ".")
	if len(ips) != 4 {
		return 0, false
	}
	var intIP int
	for i, v := range ips {
		val, err := strconv.Atoi(v)
		if err != nil || val < 0 || val > 255 {
			return 0, false
		}
		intIP = intIP | val<<((3-i)*8)
	}
	return intIP, true
}

// ToString Return string after int is converted to ip
func ToString(ipInt int) (string, bool) {
	ip1 := ipInt >> 24 & 255
	ip2 := ipInt >> 16 & 255
	ip3 := ipInt >> 8 & 255
	ip4 := ipInt & 255
	if ip1 > 255 || ip2 > 255 || ip3 > 255 || ip4 > 255 {
		return "", false
	}
	ipStr := fmt.Sprintf("%d.%d.%d.%d", ip1, ip2, ip3, ip4)
	return ipStr, true
}
