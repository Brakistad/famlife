package gofamlife_test

import (
	"testing"

	"github.com/Brakistad/famlife/gofamlife"
)


func TestFamlifer(t *testing.T) {
	if gofamlife.Famlifer() != "Hello Famlifer!" {
		t.Fatal("Thats not the right message! :(")
	}
}

