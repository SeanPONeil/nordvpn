package main

import (
	"fmt"
	"github.com/seanponeil/nordvpn"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	status     = kingpin.Flag("status", "Connection status").Bool()
	server     = kingpin.Flag("server", "Current server").Bool()
	country    = kingpin.Flag("country", "Country").Bool()
	city       = kingpin.Flag("city", "City").Bool()
	ip         = kingpin.Flag("ip", "Your new IP").Bool()
	technology = kingpin.Flag("technology", "Current technology").Bool()
	protocol   = kingpin.Flag("protocol", "Current protocol").Bool()
	transfer   = kingpin.Flag("transfer", "Transfer").Bool()
	uptime     = kingpin.Flag("uptime", "Uptime").Bool()
)

// Match flag argument to NordVPN status key
func matchKey() string {
	if *status == true {
		return "Status"
	} else if *server == true {
		return "Current Server"
	} else if *country == true {
		return "Country"
	} else if *city == true {
		return "City"
	} else if *ip == true {
		return "Your new IP"
	} else if *technology == true {
		return "Current technology"
	} else if *protocol == true {
		return "Current protocol"
	} else if *transfer == true {
		return "Transfer"
	} else if *uptime == true {
		return "Uptime"
	} else {
		// return connection status if no flags supplied
		return "Status"
	}
}

func main() {
	kingpin.Parse()
	nordvpn := nordvpn.Status()
	key := matchKey()
	fmt.Printf("%s", nordvpn[key])
}
