package main

import (
	"reflect"
	"testing"
)

func TestSplitLines(t *testing.T) {
	expected := `foo
	bar
	baz`

	t.Logf("expected %s", expected)

	actual := SplitLines(expected)

	t.Logf("actual %s", actual)
	if reflect.DeepEqual(expected, actual) {
		t.Errorf("TestSplitLines failed, expected %s, actual %s", expected, actual)
	}
}
