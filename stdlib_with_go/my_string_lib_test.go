package stdlibwithgo

/**
https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
*/

import (
	"math/rand"
	"testing"
)

// package-variable
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GetRandomString() string {
	var ran []rune

	ran = make([]rune, rand.Intn(10))
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
