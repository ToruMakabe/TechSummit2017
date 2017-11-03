package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	log.Println("Hostname: ", hostname)

	log.Println("Network Interfaces: ")
	ifaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			panic(err)
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			log.Println("IP Address: ", ip)
		}
	}

	for {
		resp, err := http.Get("http://localhost:8080")
		if err != nil {
			panic(err)
		}
		log.Println("Get Server Headers: ", resp.Header)
		log.Println("Get Server Status: ", resp.Status)

		time.Sleep(10 * time.Second)
	}
}
