package nordvpn

import (
	"bufio"
	"log"
	"os/exec"
	"strings"
)

// Returns the output of
//	nordvpn status
//as a map of key value pairs.
func Status() map[string]string {
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

// Split a string by the delimiter '\n'
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
