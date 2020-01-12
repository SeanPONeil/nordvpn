package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
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

	cityPtr := flag.Bool("city", false, "Print city of NordVPN connection")

	flag.Parse()

	nordvpn := NordVPNCmd()
	if *cityPtr == true {
		fmt.Printf("%s", nordvpn["City"])
	} else {
		// print connection status by default
		fmt.Printf("%s", nordvpn["Status"])
	}
}
