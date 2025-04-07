package input

import (
	"testing"
)

func TestInput(t *testing.T) {
	c := Clean("(\"'<>&`#--%*;\\AND OR UNION | =)")

	if c != "" {
		t.Fatalf("result %s is not expected", c)
	}
}
