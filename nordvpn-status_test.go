package main

import (
	"testing"
)

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
