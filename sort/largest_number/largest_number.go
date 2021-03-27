package lagest_number

import (
	"sort"
	"strconv"
	"strings"
)

//문제 https://programmers.co.kr/learn/courses/30/lessons/42748
/**
	가장 큰 수

문제 설명:
0 또는 양의 정수가 주어졌을 때, 정수를 이어 붙여 만들 수 있는 가장 큰 수를 알아내 주세요.

예를 들어, 주어진 정수가 [6, 10, 2]라면 [6102, 6210, 1062, 1026, 2610, 2106]를 만들 수 있고, 이중 가장 큰 수는 6210입니다.

0 또는 양의 정수가 담긴 배열 numbers가 매개변수로 주어질 때, 순서를 재배치하여 만들 수 있는 가장 큰 수를 문자열로 바꾸어 return 하도록 solution 함수를 작성해주세요.

*/


type customSort []string

func (s customSort) Len() int {
    return len(s)
}

func (s customSort) Swap(i, j int){
    s[i], s[j] = s[j], s[i]
}

func (s customSort) Less(i, j int) bool{

    if result := strings.Compare(s[i]+s[j], s[j]+s[i]); result == 1{
        return true
    } else{
        return false
    }
}

func Solution(numbers []int) string {
	strs := make([]string, len(numbers))

	for index, number := range numbers {
		strs[index] = strconv.Itoa(number)
	}

	sort.Slice(strs, func(i, j int) bool {
		return strs[i]+strs[j] > strs[j]+strs[i]
	})

	result := strings.Join(strs, "")
	n, _ := strconv.Atoi(result)
	if n == 0 {
		return "0"
	}

	return result
}
