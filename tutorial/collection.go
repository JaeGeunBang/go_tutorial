package main

import "fmt"

func main() {
	// Array
	var a [3]int
	a[0] = 1
	a[1] = 2
	a[2] = 3
	println(a[1])

	var a1 = [3]int{1, 2, 3}
	var a3 = [...]int{1, 2, 3}
	println(a1[1])
	println(a3[1])

	// Multiple Array
	var multiArray [3][4]int
	multiArray[0][0] = 10

	var multiArray2 = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	println(multiArray2[1][2])

	// Slice (java에 ArrayList와 같은 역할. 동적으로 알아서 사이즈가 늘어남)
	var aa []int
	aa = []int{1, 2, 3}
	aa[1] = 10
	fmt.Println(aa)

	s := make([]int, 5, 10)
	println(len(s), cap(s)) // 길이, 캐퍼시티

	// sub-slice
	s2 := []int{0,1,2,3,4,5}
	s2 = append(s2, 6)
	s2 = s2[2:]
	fmt.Println(s2)

	// len=0, cap=3 인 슬라이스
	sliceA := make([]int, 0, 3)
	for i := 1; i <= 15; i++ {
		sliceA = append(sliceA, i)
		fmt.Println(len(sliceA), cap(sliceA))
	}

	fmt.Println(sliceA)

	// Map
	var idMap map[int]string
	idMap = make(map[int]string)
	tickers := map[string]string {
		"GooG": "Google Inc",
		"MSFT": "Microsoft",
		"FB": "Facebook",
	}

	idMap[901] = "Apple"
	tickers["Apple"] = "Apple Inc"
	str := idMap[901]
	println(str)

	noData := tickers["Test"]
	println(noData)

	// Map에 요소 삭제
	delete(idMap, 901)

	// Map에 값이 있는지 체크
	val, exists := tickers["MSFT2"]
	if !exists {
		println(val)
		println("No MSFT tickers")
	}

	// Map 반복문
	for key, val := range tickers {
		println(key, val)
	}
}