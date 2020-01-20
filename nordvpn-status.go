package main

import (
	"bufio"
	"fmt"
	"github.com/schollz/closestmatch"
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
	"os"
	"os/exec"
	"strings"
)

var (
	app        = kingpin.New("nordvpn-status", "Utility to retrieve NordVPN status elements. Designed for use in Polybar/Lemonbar.")
	status     = app.Flag("status", "Connection status").Bool()
	server     = app.Flag("server", "Current server").Bool()
	country    = app.Flag("country", "Country").Bool()
	city       = app.Flag("city", "City").Bool()
	ip         = app.Flag("ip", "Your new IP").Bool()
	technology = app.Flag("technology", "Current technology").Bool()
	protocol   = app.Flag("protocol", "Current protocol").Bool()
	transfer   = app.Flag("transfer", "Transfer").Bool()
	uptime     = app.Flag("uptime", "Uptime").Bool()
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
func matchKey(arg string) string {
	keys := []string{
		"Status",
		"Current server",
		"Country",
		"City",
		"Your new IP",
		"Current technology",
		"Current protocol",
		"Transfer",
		"Uptime"}

	if Contains(keys, arg) == true {
		return arg
	} else {
		bagSizes := []int{30}
		cm := closestmatch.New(keys, bagSizes)
		return cm.Closest(arg)
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
	if len(os.Args) == 1 {
		// return value of Status if no args are passed in
		fmt.Printf("%s", nordvpn["Status"])
	} else if len(os.Args) == 2 {
		key := matchKey(os.Args[1])
		fmt.Printf("%s", nordvpn[key])
	} else {
		fmt.Fprintf(os.Stderr, "nordvpn-status can only accept one argument")
		os.Exit(1)
	}
}
