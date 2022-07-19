package main

import "fmt"

type IPAddr [4]byte

// func (p IPAddr) String() string {
// 	return fmt.Sprintf("%v (%v years)", p.String(), p.String())
// }

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	fmt.Println(hosts)
	
	for name, ip := range hosts {
		for i := range ip {
			fmt.Printf("%v: %v\n", name, ip[i])
		}
	}
}
