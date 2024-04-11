package stdlibwithgo

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
func Len(s string) int {
	var runes []rune

	runes = []rune(s)
	return len(runes)
}
