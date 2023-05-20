package util

import "strings"

func RemovePort(ip string) string {
	return strings.Split(ip, ":")[0]
}