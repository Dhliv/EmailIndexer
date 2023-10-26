package faststrings_test

import (
	"testing"

	"github.com/dhliv/EmailIndexing/src/fast_strings"
)

func TestFastStringsConcat(t *testing.T) {
	fs := fast_strings.NewFastString()
	a := "Este es un buen string"
	b := "Este es otro\nbuen string"
	ab := "Este es un buen stringEste es otro\nbuen string"
	fs.Concat(&a)
	fs.Concat(&b)
	res := fs.GetString()

	areEqual := len(*res) == len(ab)
	if areEqual {
		for i := 0; i < len(ab); i++ {
			if ab[i] != (*res)[i] {
				areEqual = false;
				break
			}
		}
	}

	if !areEqual {
		t.Errorf("%v != %v\n", *res, ab)
	}
}

func TestFastStringsConcatWithTrim(t *testing.T) {
	fs := fast_strings.NewFastString()
	a := "       Este es un buen string       "
	b := "   Este es otro\nbuen string    "
	ab := "Este es un buen string   Este es otro\nbuen string"
	fs.Concat(&a)
	fs.Concat(&b)
	res := fs.GetString()

	areEqual := len(*res) == len(ab)
	if areEqual {
		for i := 0; i < len(ab); i++ {
			if ab[i] != (*res)[i] {
				areEqual = false;
				break
			}
		}
	}

	if !areEqual {
		t.Errorf("%v != %v\n", *res, ab)
	}
}


func TestFastStringsCutPrefix(t *testing.T) {
	fs := fast_strings.NewFastString()
	a := "       Este es un buen string       "
	b := "   Este es otro\nbuen string    "
	c := "Este es"
	ab := "un buen string   Este es otro\nbuen string"
	fs.Concat(&a)
	fs.Concat(&b)
	hasPrefix := fs.CutPrefix(&c)
	if !hasPrefix {
		t.Errorf("%v Trimmed contains prefix %v\n", a, c)
	}

	res := fs.GetString()
	areEqual := len(*res) == len(ab)
	if areEqual {
		for i := 0; i < len(ab); i++ {
			if ab[i] != (*res)[i] {
				areEqual = false;
				break
			}
		}
	}

	if !areEqual {
		t.Errorf("%v != %v\n", *res, ab)
	}
}

func TestFastStringsCutNonePrefix(t *testing.T) {
	fs := fast_strings.NewFastString()
	a := "dijkstra"
	c := "aho-corasick"
	ab := "dijkstra"
	fs.Concat(&a)
	hasPrefix := fs.CutPrefix(&c)
	if hasPrefix {
		t.Errorf("%v Trimmed doesn't contains prefix %v\n", a, c)
	}

	res := fs.GetString()
	areEqual := len(*res) == len(ab)
	if areEqual {
		for i := 0; i < len(ab); i++ {
			if ab[i] != (*res)[i] {
				areEqual = false;
				break
			}
		}
	}

	if !areEqual {
		t.Errorf("%v != %v\n", *res, ab)
	}
}