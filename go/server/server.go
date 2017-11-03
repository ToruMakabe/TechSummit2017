package main

import (
	"bytes"
	"encoding/json"
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
	fmt.Fprintln(w, "[Hostname]")
	fmt.Fprintln(w, hostname)

	fmt.Fprintln(w)
	fmt.Fprintln(w, "[Network Interfaces]")
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

	fmt.Fprintln(w)
	fmt.Fprintln(w, "[Env. Variables]")
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
		var buf bytes.Buffer
		err := json.Indent(&buf, []byte(bodyBytes), "", "  ")
		if err != nil {
			panic(err)
		}
		indentJson := buf.String()
		fmt.Fprintln(w, "[Instance Metadata]")
		fmt.Fprintln(w, indentJson)
	}

}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
