package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(w, "Hostname: ", hostname)

	fmt.Fprintln(w)
	fmt.Fprintln(w, "Env. Variables: ")
	for _, pair := range os.Environ() {
		fmt.Fprintln(w, pair)
	}

	fmt.Fprintln(w)
	url := "http://169.254.169.254/metadata/instance?api-version=2017-08-01"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Metadata", "true")
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintln(w, "Failed attempt to get instance metadata: ", err)
	} else {
		defer resp.Body.Close()
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		bodyString := string(bodyBytes)
		fmt.Fprintln(w, "Instance Metadata: ")
		fmt.Fprintln(w, bodyString)
	}

	fmt.Fprintln(w)
	fmt.Fprintln(w, "Network Interfaces: ")
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
			fmt.Fprintln(w, "IP Address: ", ip)
		}
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
