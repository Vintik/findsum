package findsum

import (
	"fmt"
	// "log"
)

var (
	errNotFound = fmt.Errorf("Sum is not found")
)

func FindSum(a []int, sum int) (int, int, error) {
	var s, l, r, offset int

	if sum == 0 {
		offset = 1
		sum = 1
	}

	for {
		switch {
		case s < sum:
			//log.Printf("s < sum: s:%v, sum:%v", s, sum)
			if r == len(a) {
				return 0, 0, errNotFound
			}
			s = s + (a[r] + offset)
			r++

		case s > sum:
			//log.Printf("s > sum: s:%v, sum:%v", s, sum)
			if l == r {
				return 0, 0, errNotFound
			}
			s = s - (a[l] + offset)
			l++

		case s == sum:
			//log.Printf("found: left:%d, right:%d", l, r)
			return l, r, nil

		}
	}
}
