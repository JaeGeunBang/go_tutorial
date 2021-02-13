package main

func main() {
	msg := "Hello"
	say(msg)
	println(msg)

	say2(&msg)
	println(msg)

	say3("this", "is", "a", "book")
	say3("Hi")

	count, total := sum(1,7,3,5,9)
	println(count, total)

	// 익명함수
	sum := func(n ...int) int {
		s := 0
		for _, i := range n {
			s += i
		}
		return s
	}
	println(sum)

	// 일급함수
	//// 함수를 파라미터로 전달할 수 있다. (익명함수를 변수 할당후 변수를 전달)
	add := func(i int, j int) int {
		return i + j
	}
	r1 := calc(add, 10, 20)
	println(r1)

	// 클로저
	next := nextValue()
	println(next())
	println(next())
	println(next())

}

// call by value
func say(msg string) {
	println(msg)
}

// call by reference
func say2(msg *string) {
	println(*msg)
	*msg = "Changed"
}

// 가변 인자 함수
func say3(msg ...string) {
	for _, s := range msg {
		println(s)
	}
}

// 반환값이 있는 함
func sum(nums ...int) (int, int) {
	count:= 0
	s := 0
	for _, n := range nums {
		s += n
		count++
	}
	return count, s
}

// 일급함수
func calc(f func(int, int) int, a int, b int) int {
	result := f(a, b)
	return result
}

// 클로저
//// 함수 바깥에 있는 변수를 참조하는 함수 값을 일컫는다.
//// 이때 함수는 바깥의 변수를 마치 함수안으로 끌어들인 듯이 그 변수를 읽거나 쓴다.
//// ex. 내부에 있는 func()이 바깥에 있는 변수 i를 마치 끌어들인 듯이 사용한다.
func nextValue() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}