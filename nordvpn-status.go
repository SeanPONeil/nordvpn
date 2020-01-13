package main

import (
	"bufio"
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

	scanner := bufio.NewScanner(strings.NewReader(string(out)))
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ":")
		key, value := line[0], line[1]
		m[key] = value
	}
	return m
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
