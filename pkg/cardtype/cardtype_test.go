package cardtype

import (
	"testing"
)

func TestInvalidCard(t *testing.T) {

	if _, exist := CardsByCode["023"]; exist {
		t.Fatal("Did not expect a nonexistent card!")
	}

}

func TestGetCards(t *testing.T) {

	if expect, found := "American Express", CardsByCode["003"].Description; found != expect {
		t.Fatalf("Expected '%s' but found '%s'", expect, found)
	}

}

func TestCardDescription(t *testing.T) {

	if expect, found := "034: Dankort", CardsByCode["034"].String(); found != expect {
		t.Fatalf("Expected '%s' but found '%s'", expect, found)
	}

}
