package stdlibwithgo

import (
	"fmt"
	"strings"
)

/**************************************
My_Strings_lib

strings 라이브러리를 만들어보며 golang에서의 문자열에 대해 공부해보자.
Origin:
https://pkg.go.dev/strings#pkg-overview
***************************************/

/*
len() 함수는 문자열의 byte 크기를 return 한다. 하지만 golang에서 문자열은 UTF-8인코딩을 사용하며, 각 문자의 길이가 다를 수 있다.
[]rune은 문자열을 유니코드 slice로 변환하는데 사용한다. rune은 golang에서 유니코드 포인트를 나타내는 타입으로 int32와 동일하다.
*/
func Len(str string) int {
	var runes []rune

	//when string is empty
	if str == "" {
		return 0
	}
	runes = []rune(str)
	return len(runes)
}

/*
https://stackoverflow.com/questions/47352449/immutable-string-and-pointer-address/47352588#47352588
https://www.freecodecamp.org/news/new-vs-make-functions-in-go/
https://stackoverflow.com/questions/75502153/is-string-a-reference-type-or-a-value-type
golang에서 string type은 immutable 타입이다.
golang에서 string type은 c에서 char *와 다르다.
golang에서 string type은 2개의 데이터를 보관한다.

	type StringHeader struct {
	    Data uintptr
	    Len  int
	}

만약 string에 새로운 값을 할당한다면, string header의 주소는 변경되지 않고 내부의 Data 포인터 값이 변경된다.
*/
func Clone(str string) string {
	//variable init
	var cpy_builder strings.Builder = strings.Builder{}

	cpy_builder.Grow(Len(str))
	//when str is empty
	if str == "" {
		return ""
	}

	for _, chr := range str {
		fmt.Fprintf(&cpy_builder, "%c", chr)
	}
	return cpy_builder.String()
}

/*
Compare returns an integer comparing two strings lexicographically. The result will be 0 if a == b, -1 if a < b, and +1 if a > b.
golang의 문자열은 c언어와 다르게 \0으로 끝나지 않는다.
*/
func Compare(a, b string) int {
	var (
		runes_a []rune = []rune(a)
		runes_b []rune = []rune(b)
	)
	//add 0char postfix for empty string cases
	runes_a = append(runes_a, 0)
	runes_b = append(runes_b, 0)

	//compare char
	for i, chr := range runes_a {
		//check character
		if chr > runes_b[i] {
			return 1
		}
		if chr < runes_b[i] {
			return -1
		}
	}
	return 0
}

/*
two pointer를 비스무리우스하게 응용해서 구현.
*/
func Contains(s, substr string) bool {
	var (
		runes_org []rune = []rune(s)
		runes_sub []rune = []rune(substr)
		org_idx   int    = 0
		sub_idx   int    = 0
		leng      int
	)
	//add 0char postfix for empty string cases
	runes_org = append(runes_org, 0)
	runes_sub = append(runes_sub, 0)
	leng = len(runes_org)
	//compare loop
	for ; org_idx < leng; org_idx++ {
		//complete
		if runes_sub[sub_idx] == 0 {
			return true
		}
		//compare character
		if runes_org[org_idx] == runes_sub[sub_idx] {
			sub_idx++
			continue
		}
		//protect continuous case
		if sub_idx > 0 {
			sub_idx = 0
			org_idx--
		}
	}
	return false
}

/*
ContainsFunc reports whether any Unicode code points r within s satisfy f(r).

rune slice와 string은 다르다. ->> https://stackoverflow.com/questions/49062100/is-there-any-difference-between-range-over-string-and-range-over-rune-slice
range는 string에서 byte index와 rune을 반환한다. 그렇기 때문에 UTF-8문자가 1byte가 아닐 경우가 있기 때문에, string에서 index는 1씩 증가한다는 보장이 없다.
예를 들어------
range string
0: 'こ'
3: 'ん'
6: 'に'
9: 'ち'
12: 'は'
15: '世'
18: '界'

range []rune(s)
0: 'こ'
1: 'ん'
2: 'に'
3: 'ち'
4: 'は'
5: '世'
6: '界
*/
func ContainsFunc(s string, f func(rune) bool) bool {
	//compare loop
	for _, chr := range s {
		if f(chr) {
			return true
		}
	}
	return false
}

/*
Golang에서 접근제한자는 존재하지 않지만, 함수 이름의 첫 글자가 대문자인지 아닌지에 따라 패키지 외부에서 접근 가능한지 아닌지 정할 수 있다.
substr을 발견할 경우 substr의 마지막 char 다음 첫 글자의 idx를 반환, else 문자열 마지막 인덱스 반환
*/
//############ERRROR##############

func get_substr(org, sub []rune, start int) (startpoint, endpoint int, found int) {
	var (
		idx_org         int = start
		idx_sub         int = 0
		idx_start, leng int
	)

	leng = len(org)
	//if substr is empty
	if sub[idx_sub] == 0 {
		return idx_start, idx_org, 1
	}
	for ; idx_org < leng; idx_org++ {
		if org[idx_org] == sub[idx_sub] {
			//set startpoint
			if idx_sub == 0 {
				idx_start = idx_org
			}
			idx_sub++
			//compare complete
			if sub[idx_sub] == 0 {
				return idx_start, idx_org + 1, 1
			}
			continue
		}
		//return to org idx
		if idx_sub > 0 {
			idx_sub = 0
			idx_org--
		}
	}
	return 0, 0, 0
}

/*
Count counts the number of non-overlapping instances of substr in s. If substr is an empty string, Count returns 1 + the number of Unicode code points in s.

Golang에서 접근제한자는 존재하지 않지만, 함수 이름의 첫 글자가 대문자인지 아닌지에 따라 패키지 외부에서 접근 가능한지 아닌지 정할 수 있다.
*/
func Count(s, substr string) int {
	var (
		cnt       int = 0
		idx       int = 0
		equal     int = 0
		leng      int
		runes_org []rune = []rune(s)
		runes_sub []rune = []rune(substr)
	)

	//add 0char postfix for empty string cases
	runes_sub = append(runes_sub, 0)
	leng = len(runes_org)

	if substr == "" {
		return leng + 1
	}

	for idx < leng {
		_, idx, equal = get_substr(runes_org, runes_sub, idx)
		if equal == 0 {
			break
		}
		cnt += equal
	}

	return cnt
}

/*
Cut slices s around the first instance of sep, returning the text before and after sep.
The found result reports whether sep appears in s. If sep does not appear in s, cut returns s, "", false.

golang에서 slice는 인덱싱을 지원한다.
*/

func Cut(s, sep string) (before, after string, found bool) {
	var (
		runes_org             []rune = []rune(s)
		runes_sub             []rune = []rune(sep)
		idx_start, idx_end, f int
	)

	//add 0char postfix for empty string cases
	runes_sub = append(runes_sub, 0)
	//contains
	idx_start, idx_end, f = get_substr(runes_org, runes_sub, 0)
	return string(runes_org[0:idx_start]), string(runes_org[idx_end:]), f == 1
}

/*
Fields splits the string s around each instance of one or more consecutive white space characters, as defined by unicode.
IsSpace, returning a slice of substrings of s or an empty slice if s contains only white space.

make와 new는 다르다. -->https://www.freecodecamp.org/news/new-vs-make-functions-in-go/
*/

func Fields(s string) []string {
	var (
		runes_org []rune = []rune(s)
		ret       []string
		s_ptr     int = -2
		e_ptr     int = -1
	)
	//dynamic allocate
	ret = make([]string, 0, 5)
	//add 0char postfix for empty string cases
	runes_org = append(runes_org, 0)

	for idx, chr := range runes_org {
		if chr != ' ' && s_ptr <= e_ptr {
			s_ptr = idx
			continue
		}
		if (chr == ' ' || chr == 0) && e_ptr < s_ptr {
			e_ptr = idx
			ret = append(ret, string(runes_org[s_ptr:e_ptr]))
			s_ptr = e_ptr
		}
	}
	return ret
}
