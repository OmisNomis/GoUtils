 package main

import (
	"log"
	"net"
)

func main() {
	addrs, _ := net.InterfaceAddrs()

	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				log.Println(ipnet.IP.String())
			}
		}
	}

}
