package main

import (
	"fmt"
	"math"
)

/*
关键点：

情况一：猜中了（solution[index] == guess[index]）：猜中次数+1
情况二：没猜中（solution[index] != guess[index]）：统计字母的出现次数
接着对情况二继续处理，伪处理就是要统计solution和guess中出现相同字母次数的总和。
举例来说，solution中出现'R'字母5次，guess中出现'R'字母3次，则认为出现了3次相同的字母。
 */

func masterMind(solution, guess string) []int {
	size := len(solution)

	hit, false_hit := 0, 0
	s_count := make(map[int]int)
	g_count := make(map[int]int)

	for i := 0; i < size; i++ {
		if solution[i] == guess[i] {
			hit++
		} else {
			s_count[int(solution[i])]++
			g_count[int(guess[i])]++
		}
	}

	for _, elem := range []int{'R', 'Y', 'G', 'B'}  {
		fmt.Println(fmt.Sprintf("elem[%c], s count[%d], g count[%d]", elem, s_count[elem], g_count[elem]))
		false_hit += int(math.Min(float64(s_count[elem]),float64(g_count[elem])))
	}

	return []int{hit, false_hit}
}

func main() {
	g_count := make(map[int]int)
	g_count[10]++
	fmt.Println(g_count)

	fmt.Println([]int{'R', 'Y', 'G', 'B'})

	fmt.Println(masterMind("RGRB", "BGGR"))
}
