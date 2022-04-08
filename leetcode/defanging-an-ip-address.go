package main

import (
	"fmt"
	"strings"
)

// source: https://leetcode.com/problems/defanging-an-ip-address/

func defangIPaddr1(address string) string {
	return strings.Replace(address, ".", "[.]", 3)
}

func defangIPaddr(address string) string {
	var res string
	for i := range address {
		if address[i] == '.' {
			res += "[.]"
			continue
		}
		res += string(address[i])
	}
	return res
}

func main() {
	// Example 1
	var address1 = "1.1.1.1"
	fmt.Println("Expected: \"1[.]1[.]1[.]1\"	Output: ", defangIPaddr(address1))

	// Example 2
	var address2 = "255.100.50.0"
	fmt.Println("Expected: \"255[.]100[.]50[.]0\"	Output: ", defangIPaddr(address2))

}
