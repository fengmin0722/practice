package main

import (
	"fmt"
	"sort"
)

type Item struct {
	Start, Stop int
}

func canAttendMeeting(items []*Item) bool {
	sort.Slice(items, func(i, j int) bool {
		return items[i].Start < items[j].Start
	})

	for i := 0; i < len(items) - 1; i++ {
		if items[i].Start > items[i+1].Stop {
			return false
		}
	}
	return true
}

func main() {
	items := []*Item {
		&Item{0, 30},
		&Item{15, 20},
		&Item{5, 10},
	}
	fmt.Println(canAttendMeeting(items))
}
