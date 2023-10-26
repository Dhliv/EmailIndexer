package faststrings_test

import (
	"testing"

	"github.com/dhliv/EmailIndexing/src/fast_strings"
)

func TestFastStringsConcatWithTrimForFastString(t *testing.T) {
	fs := fast_strings.NewFastString()
	a := "       Este es un buen string       "
	b := "   Este es otro\nbuen string    "
	ab := "Este es un buen string   Este es otro\nbuen string"
	fs.Concat(&a)
	fs.Concat(&b)

	c := "Este es un string de prueba     "
	d := " igual que este      otro "
	cd := "Este es un string de prueba igual que este      otro"
	fs2 := fast_strings.NewFastString()
	fs2.Concat(&c)
	fs2.Concat(&d)

	fs.ConcatFastString(fs2)
	abcd := ab + cd
	res := fs.GetString()

	areEqual := len(*res) == len(abcd)
	if areEqual {
		for i := 0; i < len(abcd); i++ {
			if abcd[i] != (*res)[i] {
				areEqual = false;
				break
			}
		}
	}

	if !areEqual {
		t.Errorf("%v != %v\n", *res, abcd)
	}
}