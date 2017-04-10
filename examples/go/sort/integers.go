package sort

import (
	"math/rand"
	"time"
	"sort"
	"fmt"
)

func main() {
	ints := createUnsortedSlice(0, 100, 100)
	sort.Sort(ints)
	fmt.Println(ints)
}

type Integers []int

func (s Integers) Len() int {
	return len(s)
}
func (s Integers) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s Integers) Less(i, j int) bool {
	return s[i] < s[j]
}

// createUnsortedSlice will generate a random unsorted list of integers where
// min is the minimum value an int might be
// max is the maximum value an int might be
// len is the length of the list to return
func createUnsortedSlice(min, max, len int) Integers {
	rand.Seed(time.Now().UTC().UnixNano())
	var ints []int
	for i := 0; i < len; i++ {
		var x int
		x = min + rand.Intn(max-min)
		ints = append(ints, x)
	}
	return Integers(ints)
}
