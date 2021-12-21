package task01

import (
	"fmt"
	"strings"
	"testing"
)

func TestSayHello(t *testing.T) {
	testName := "John"
	expected := fmt.Sprintf("Hello %s", testName)
	msg := SayHello(testName)
	if !strings.EqualFold(msg, expected) {
		fmt.Println([]rune(msg))
		t.Errorf("Unexpected result:\n\tExpected: %q\n\tGot: %q", expected, msg)
	}
}
