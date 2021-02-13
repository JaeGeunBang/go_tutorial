package main

import "fmt"

func main() {
	println("Hello World")

	// 변수
	var i, j, k int = 1,2,3

	// 상수
	const c int = 10
	const (
		Visa = "Visa"
		Maser = "MasterCard"
	)

	// 데이터 타입 (:= --> 알아서 데이터 타입도 정해진다)
	rawLiteral := "아리랑"
	fmt.Println(rawLiteral)
	fmt.Println(i + j + k)

	// 데이터 타입 변환
	str := "ABC"
	bytes := []byte(str)
	str2 := string(bytes)
	println(bytes, str2)

	// 연산자
	d := (i + j) / 5;
	d++;

	// 포인터 연산자
	// &: 변수의 주소 값, *: 변수가 가리키는 주소에 실제 내용
	q := 10
	p := &q
	println(&p)
}
