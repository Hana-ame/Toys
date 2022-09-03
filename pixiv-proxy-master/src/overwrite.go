package main

import (
	"strings"
)

func overWrite(s, host string) string {
	// fmt.Println(pairsByHost)
	// fmt.Println(s, host)
	for k, v := range pairsByHost[host] {
		s = strings.Replace(s, k, v, -1)
		// fmt.Println(1)
	}
	// fmt.Println(s)
	return s
}
