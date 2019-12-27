package utils

import "strings"

func StrToIp(Address string) []string {
	return strings.Split(Address, ",")
}
