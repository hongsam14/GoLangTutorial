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
	const str string = "한글"
	var (
		leng int
	)

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
	const str_a = ""
	var (
		str_b     string
		org, mine int
	)
	//init
	str_b = GetRandomString()
	org = strings.Compare(str_a, str_b)
	mine = Compare(str_a, str_b)

	if org != mine {
		t.Fatalf(`Compare("%s", "%s") = %d, want for %d`, str_a, str_b, mine, org)
	}
}

func TestCompareBNil(t *testing.T) {
	const str_b string = ""
	var (
		str_a     string
		org, mine int
	)
	//init
	str_a = GetRandomString()
	org = strings.Compare(str_a, str_b)
	mine = Compare(str_a, str_b)

	if org != mine {
		t.Fatalf(`Compare("%s", "%s") = %d, want for %d`, str_a, str_b, mine, org)
	}
}

func TestCompareBothNil(t *testing.T) {
	const (
		str_a string = ""
		str_b string = ""
	)
	var (
		org, mine int
	)
	//init
	org = strings.Compare(str_a, str_b)
	mine = Compare(str_a, str_b)

	if org != mine {
		t.Fatalf(`Compare("%s", "%s") = %d, want for %d`, str_a, str_b, mine, org)
	}
}

func TestContainsContinueCase(t *testing.T) {
	const (
		str    string = "aab"
		substr string = "aaaaaaaaaaaaaaaabaaaa"
	)
	var (
		org, mine bool
	)
	//init
	org = strings.Contains(str, substr)
	mine = Contains(str, substr)
	//compare
	if org != mine {
		t.Fatalf(`Contains("%s", "%s") = %t, want for %t`, str, substr, mine, org)
	}
}

func TestContainsEmptyCase(t *testing.T) {
	const (
		str    string = ""
		substr string = "aaaaaaaaaa"
	)
	var (
		org, mine bool
	)
	//init
	org = strings.Contains(str, substr)
	mine = Contains(str, substr)
	//compare
	if org != mine {
		t.Fatalf(`Contains("%s", "%s") = %t, want for %t`, str, substr, mine, org)
	}
}

func TestContainsEmptyCase2(t *testing.T) {
	const (
		substr string = "aaaaaaaaa"
		str    string = ""
	)
	var (
		org, mine bool
	)
	//init
	org = strings.Contains(str, substr)
	mine = Contains(str, substr)
	//compare
	if org != mine {
		t.Fatalf(`Contains("%s", "%s") = %t, want for %t`, str, substr, mine, org)
	}
}

func TestContains(t *testing.T) {
	var (
		str       string
		substr    string
		org, mine bool
	)
	//init
	substr = GetRandomString()
	str = GetRandomString() + substr + GetRandomString()
	org = strings.Contains(str, substr)
	mine = Contains(str, substr)
	//compare
	if org != mine {
		t.Fatalf(`Contains("%s", "%s") = %t, want for %t`, str, substr, mine, org)
	}
}

func TestContainsFunc(t *testing.T) {
	var (
		str, fstr string
		org, mine bool
		f         func(rune) bool
	)

	//init
	str = GetRandomString()
	f = func(r rune) bool {
		return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u'
	}
	fstr = `func(r rune) bool { return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' }`

	org = strings.ContainsFunc(str, f)
	mine = ContainsFunc(str, f)
	//compare
	if org != mine {
		t.Fatalf(`Contains("%s", "%s") = %t, want for %t`, str, fstr, mine, org)
	}
}

func TestCount(t *testing.T) {
	var (
		str             string
		substr          string
		org, mine, loop int
	)
	//init
	loop = rand.Intn(10)
	substr = GetRandomString()
	for i := 0; i < loop; i++ {
		str = GetRandomString() + substr + GetRandomString()
	}
	org = strings.Count(str, substr)
	mine = Count(str, substr)
	//compare
	if org != mine {
		t.Fatalf(`Count("%s", "%s") = %d, want for %d`, str, substr, mine, org)
	}
}

func TestCountEmpty(t *testing.T) {
	const substr string = ""
	var (
		str       string
		org, mine int
	)
	//init
	str = GetRandomString()
	org = strings.Count(str, substr)
	mine = Count(str, substr)
	//compare
	if org != mine {
		t.Fatalf(`Count("%s", "%s") = %d, want for %d`, str, substr, mine, org)
	}
}

func TestCut(t *testing.T) {
	var (
		str                          string
		substr                       string
		org_l, org_r, mine_l, mine_r string
		org, mine                    bool
	)
	//init
	substr = GetRandomString()
	str = GetRandomString() + substr + GetRandomString()
	org_l, org_r, org = strings.Cut(str, substr)
	mine_l, mine_r, mine = Cut(str, substr)
	//compare
	if org_l != mine_l || org_r != mine_r || org != mine {
		t.Fatalf(`Cut("%s", "%s") = ("%s", "%s", %t), want for ("%s", "%s", %t)`, str, substr, mine_l, mine_r, mine, org_l, org_r, org)
	}
}

func TestField(t *testing.T) {
	var (
		str       string
		org, mine []string
		loop      int
	)
	//init
	loop = rand.Intn(10)
	for i := 0; i < loop; i++ {
		str = GetRandomString() + "    " + GetRandomString() + "   "
	}
	org = strings.Fields(str)
	mine = Fields(str)
	//compare
	for i, s := range org {
		if s != mine[i] {
			t.Fatalf(`Fields("%s") = %v, want for %v`, str, mine, org)
		}
	}
}
