package main

import "fmt"

/* Make the `IPAddr` type implement `fmt.Stringer` to print the address as a 
   dotted quad.
   
   For instance, `IPAddr{1,2,3,4}` should print as "1.2.3.4"  */

type IPAddr [4]byte

// Add a "String() string" method to IPAddr.
func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
} // there has to be a more idiomatic way to implement this

func main() {
	hosts := map[string]IPAddr{
		"loopback": {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
