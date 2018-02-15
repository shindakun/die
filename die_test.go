package die_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/shindakun/die"
)

func TestRoll(t *testing.T) {
	rand.Seed(1) // should result in 2 on roll.
	r, err := die.Roll("1d20")
	fmt.Println(r, err)
	if err != nil {
		t.Errorf("die.Roll failed, got: %v, want: %v", err, nil)
	}
	if r != 2 {
		t.Errorf("die.Roll failed, got: %v, want: %v", r, 2)
	}
}
