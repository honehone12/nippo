package input

import (
	"html"
	"testing"
)

func TestInput(t *testing.T) {
	c := Clean("\"'<>&`#--%*;\\AND OR UNION")

	if html.UnescapeString(c) != "\"'<>&" {
		t.Fatal("result is not expected")
	}
}
