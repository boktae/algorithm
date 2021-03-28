package h_index

import (
	"fmt"
	"sort"
)

//문제 https://programmers.co.kr/learn/courses/30/lessons/42747?language=go
/**

문제 설명
H-Index는 과학자의 생산성과 영향력을 나타내는 지표입니다. 어느 과학자의 H-Index를 나타내는 값인 h를 구하려고 합니다. 위키백과1에 따르면, H-Index는 다음과 같이 구합니다.

어떤 과학자가 발표한 논문 n편 중, h번 이상 인용된 논문이 h편 이상이고 나머지 논문이 h번 이하 인용되었다면 h의 최댓값이 이 과학자의 H-Index입니다.

어떤 과학자가 발표한 논문의 인용 횟수를 담은 배열 citations가 매개변수로 주어질 때, 이 과학자의 H-Index를 return 하도록 solution 함수를 작성해주세요.

*/

//내가 푼 코드
func Solution(citations []int) int {
	sort.Ints(citations)

	count := 0
	citationsLen := len(citations)
	for index, citation := range citations {
		subLen := citationsLen - index - 1
		fmt.Printf("%d > %d", citation, subLen)
		if citation > subLen {
			count++
			fmt.Printf("true\n")
		} else {
			fmt.Printf("false \n")
		}
	}
	fmt.Printf("\n")

	return count
}

//남이 작성한 코드
// func Solution(citations []int) int {
//     sort.Sort(sort.Reverse(sort.IntSlice(citations)))
//     for h := 0; h < len(citations); h++ {
//         if citations[h] < h+1 {
//             return h
//         }
//     }

//     return len(citations)
// }