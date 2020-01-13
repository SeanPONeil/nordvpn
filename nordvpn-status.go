package main

import (
	"bufio"
	"fmt"
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

func main() {

	nordvpn := NordVPNCmd()
	if len(os.Args) == 1 {
		// return value of Status if no args are passed in
		fmt.Printf("%s", nordvpn["Status"])
	} else if len(os.Args) == 2 {
		key := os.Args[1]
		fmt.Printf("%s", nordvpn[key])
	} else {
		fmt.Fprintf(os.Stderr, "nordvpn-status can only accept one argument")
		os.Exit(1)
	}
}
