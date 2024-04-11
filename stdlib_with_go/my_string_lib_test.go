package stdlibwithgo

/**
https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
*/

import (
	"math/rand"
	"strings"
	"testing"
)

// package-variable
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GetRandomString() string {
	var ran []rune = make([]rune, rand.Intn(10))

	for i := range ran {
		ran[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(ran)
}

func TestLen(t *testing.T) {
	var (
		leng     int
		org_leng int
		str      string
	)

	str = GetRandomString()

	leng = Len(str)
	org_leng = len(str)
	if leng != org_leng {
		t.Fatalf(`Len("%s") = %d, want for %d`, str, leng, org_leng)
	}
}

func TestUTF8Len(t *testing.T) {
	var (
		leng int
		str  string
	)

	str = "한글"
	leng = Len(str)

	if leng != 2 {
		t.Fatalf(`Len("%s") = %d, want for 2`, str, leng)
	}
}

// check address is different
func TestClone(t *testing.T) {
	var (
		org string
		cln string
	)

	org = GetRandomString()
	cln = Clone(org)

	if org != cln {
		t.Fatalf(`Clone("%s") = "%s", address: %x origin: , cpy: %x`, org, cln, &org, &cln)
	}
}

func TestCompare(t *testing.T) {
	var (
		str_a     string
		str_b     string
		org, mine int
	)
	//init
	str_a = GetRandomString()
	str_b = GetRandomString()
	org = strings.Compare(str_a, str_b)
	mine = Compare(str_a, str_b)

	if org != mine {
		t.Fatalf(`Compare("%s", "%s") = %d, want for %d`, str_a, str_b, mine, org)
	}
}

func TestCompareANil(t *testing.T) {
	var (
		str_a     string
		str_b     string
		org, mine int
	)
	//init
	str_a = ""
	str_b = GetRandomString()
	org = strings.Compare(str_a, str_b)
	mine = Compare(str_a, str_b)

	if org != mine {
		t.Fatalf(`Compare("%s", "%s") = %d, want for %d`, str_a, str_b, mine, org)
	}
}

func TestCompareBNil(t *testing.T) {
	var (
		str_a     string
		str_b     string
		org, mine int
	)
	//init
	str_a = GetRandomString()
	str_b = ""
	org = strings.Compare(str_a, str_b)
	mine = Compare(str_a, str_b)

	if org != mine {
		t.Fatalf(`Compare("%s", "%s") = %d, want for %d`, str_a, str_b, mine, org)
	}
}

func TestCompareBothNil(t *testing.T) {
	var (
		str_a     string
		str_b     string
		org, mine int
	)
	//init
	str_a = ""
	str_b = ""
	org = strings.Compare(str_a, str_b)
	mine = Compare(str_a, str_b)

	if org != mine {
		t.Fatalf(`Compare("%s", "%s") = %d, want for %d`, str_a, str_b, mine, org)
	}
}
