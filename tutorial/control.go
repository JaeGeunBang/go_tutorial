package main

func main() {
	// if 문
	k := 1
	if k == 1 {
		println("One")
	} else if k == 2 {
		println("Two")
	}

	max := 5
	if val := k * 2; val < max {
		println(val)
	}

	// Switch 문
	var name string
	var category = 1
	switch category {
	case 1:
		name = "Paper Book"
	case 2:
		name = "eBook"
	default:
		name = "Other"
	}
	println(name)

	// 반복문
	sum := 0
	for i := 1 ; i <= 100; i++ {
		sum += i
	}
	println(sum)

	names := []string{"1", "2", "3"}
	for index, name:= range names {
		println(index, name)
	}

	main2()
}

// break, continue, goto 예제
func main2() {
	var a = 1
	for a < 15 {
		if a == 5 {
			a += a
			continue // for루프 시작으로제
		}
		a++
		if a > 10 {
			break  //루프 빠져나옴
		}
	}
	if a == 11 {
		goto END //goto 사용예
	}
	println(a)

END:
	println("End")
}