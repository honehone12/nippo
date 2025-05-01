package input

import (
	"fmt"
	"testing"
)

func TestInput(t *testing.T) {
	if c := Clean(" (\"'<>&`#--%*;\\ | =) "); c != "" {
		t.Fatalf("result %s is not expected", c)
	}

	if c := Clean("--- ------ -- -"); c != "-   -" {
		t.Fatalf("result %s is not expected", c)
	}

	if c := Clean("------"); c != "" {
		t.Fatalf("result %s is not expected", c)
	}
}

func TestEmail(t *testing.T) {
	if _, err := ValidEmail("somethin99g@go.com"); err != nil {
		t.Fatal("should be valid")
	}

	if _, err := ValidEmail("somethin99g@go.co.xx"); err != nil {
		t.Fatal("should be valid")
	}

	if _, err := ValidEmail("something.some+some@golang.com"); err != nil {
		t.Fatal("should be valid")
	}

	// ignore double dot '..' because we don't use path here
	if _, err := ValidEmail("..-..++@go.com"); err != nil {
		t.Fatal("should be valid")
	}

	if _, err := ValidEmail("something--some@go.com"); err == nil {
		t.Fatal("should be invalid")
	}

	if _, err := ValidEmail("some@thing@go.com"); err == nil {
		t.Fatal("should be invalid")
	}

	if _, err := ValidEmail("something@go"); err == nil {
		t.Fatal("should be invalid")
	}

	if _, err := ValidEmail("something"); err == nil {
		t.Fatal("should be invalid")
	}

	if _, err := ValidEmail(""); err == nil {
		t.Fatal("should be invalid")
	}

	runes := "*%&|#=\"'`</>;()"
	for _, r := range runes {
		if _, err := ValidEmail(fmt.Sprintf("some%csome@go.com", r)); err == nil {
			t.Fatal("should be invalid")
		}
	}
}
