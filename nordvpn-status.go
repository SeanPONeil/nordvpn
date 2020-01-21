package main

import (
	"bufio"
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
	"os/exec"
	"strings"
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

func nordVPNCmd() map[string]string {
	out, err := exec.Command("nordvpn", "status").Output()
	if err != nil {
		log.Fatal(err)
	}

	m := map[string]string{}

	for _, s := range splitLines(string(out)) {
		k, v := parseKeyValue(s)
		m[k] = v
	}

	return m
}

func splitLines(s string) []string {
	var lines []string
	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func parseKeyValue(s string) (string, string) {
	kv := strings.Split(s, ":")
	k := strings.TrimSpace(kv[0])
	v := strings.TrimSpace(kv[1])
	return k, v
}

// Fuzzy matches arg to NordVPN status key,
// and returns a valid key
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

// Returns true if slice contains target string
func Contains(slice []string, target string) bool {
	for _, elem := range slice {
		if elem == target {
			return true
		}
	}
	return false
}

func main() {
	kingpin.Parse()
	nordvpn := nordVPNCmd()
	key := matchKey()
	fmt.Printf("%s", nordvpn[key])
}
