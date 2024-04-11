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
*
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
*
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
*
Compare returns an integer comparing two strings lexicographically. The result will be 0 if a == b, -1 if a < b, and +1 if a > b.
golang의 문자열은 c언어와 다르게 \0으로 끝나지 않는다.
*/
func Compare(a, b string) int {
	var (
		runes_a           []rune = []rune(a)
		runes_b           []rune = []rune(b)
		len_a, len_b, ret int
	)

	len_a = len(runes_a)
	len_b = len(runes_b)
	ret = len_a - len_b

	//check different length
	if ret > 0 {
		return 1
	}
	if ret < 0 {
		return -1
	}
	//compare char
	for i, chr := range runes_a {
		if chr > runes_b[i] {
			return 1
		}
		if chr < runes_b[i] {
			return -1
		}
	}
	return 0
}
