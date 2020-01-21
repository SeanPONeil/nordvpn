package nordvpn

import (
	"testing"
)

func TestNordVPNCmd(t *testing.T) {
	out := NordVPNCmd()

	if len(out) != 9 {
		t.Errorf("nordVPNCmd() output size was %d, expected 9", len(out))
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
