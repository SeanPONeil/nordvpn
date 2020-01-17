package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/schollz/closestmatch"
	"log"
	"os"
	"os/exec"
	"strings"
)

func NordVPNCmd() map[string]string {
	out, err := exec.Command("nordvpn", "status").Output()
	if err != nil {
		log.Fatal(err)
	}

	m := map[string]string{}

	for _, s := range splitLines(string(out)) {
		k, v := ParseKeyValue(s)
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

func ParseKeyValue(s string) (string, string) {
	kv := strings.Split(s, ":")
	k := strings.TrimSpace(kv[0])
	v := strings.TrimSpace(kv[1])
	return k, v
}

// Fuzzy matches arg to NordVPN status key,
// and returns a valid key
func MatchKey(arg string) string {
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
		bagSizes := []int{10}
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
	statusPtr := flag.Bool("status", false, "Connection status")
	serverPtr := flag.Bool("server", false, "Current server")
	countryPtr := flag.Bool("country", false, "Country")
	cityPtr := flag.Bool("city", false, "City")
	ipPtr := flag.Bool("ip", false, "Your new IP")
	techPtr := flag.Bool("technology", false, "Current technology")
	transferPtr := flag.Bool("transfer", false, "Transfer")
	uptimePtr := flag.Bool("uptimePtr", false, "Uptime")
	flag.Parse()

	nordvpn := NordVPNCmd()
	if len(os.Args) == 1 {
		// return value of Status if no args are passed in
		fmt.Printf("%s", nordvpn["Status"])
	} else if len(os.Args) == 2 {
		key := MatchKey(os.Args[1])
		fmt.Printf("%s", nordvpn[key])
	} else {
		fmt.Fprintf(os.Stderr, "nordvpn-status can only accept one argument")
		os.Exit(1)
	}
}
