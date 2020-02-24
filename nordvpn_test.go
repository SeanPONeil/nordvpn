package nordvpn

import (
	"testing"
)

func TestToMap(t *testing.T) {
	statusOutput := `Status: Connected
Current server: us4568.nordvpn.com
Country: United States
City: Ann Arbor
Your new IP: 100.100.100.100
Current technology: NordLynx
Transfer: 2.03 MiB received, 413.41 KiB sent
Uptime: 19 minutes 54 seconds`

	outputMap := toMap(statusOutput)
	if len(outputMap) != 8 {
		t.Errorf("ToMap() output size was %d, expected 8", len(outputMap))
	}
}

func TestSplitLines(t *testing.T) {
	s := "foo\nbar\nbaz"

	expected := []string{"foo", "bar", "baz"}

	actual := splitLines(s)

	if len(actual) != len(expected) {
		t.Errorf("SplitLines(%s) = %v, actual %v", s, expected, actual)
	}

	for i, v := range actual {
		if v != expected[i] {
			t.Errorf("expected[%d] = %s, actual %s", i, expected[i], v)
		}
	}
}
