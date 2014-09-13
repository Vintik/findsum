package findsum

import (
	"fmt"
	"testing"
)

func ExampleFindSum() {
	a := []int{1, 2, 3, 0, 1}

	l, r, _ := FindSum(a, 4)
	fmt.Printf("found sum:%v in a:%v : %v left:%v, right:%v\n",
		4, a, a[l:r], l, r)

	l, r, _ = FindSum(a, 0)
	fmt.Printf("found 0 sum in a:%v : %v left:%v, right:%v\n",
		a, a[l:r], l, r)
	// Output:
	// found sum:4 in a:[1 2 3 0 1] : [3 0 1] left:2, right:5
	// found 0 sum in a:[1 2 3 0 1] : [0] left:3, right:4
}

func ExampleFindSumErrNotFound() {
	a := []int{1, 2, 10, 0, 1}

	_, _, err := FindSum(a, 8)
	fmt.Printf("sum:8 in a:%v : %v\n", a, err)
	// Output: sum:8 in a:[1 2 10 0 1] : Sum is not found
}

type TestCollection []struct {
	name        string
	sum         int
	left, right int
	err         error
	a           []int
}

func TestHappyPath(testing *testing.T) {
	tests := TestCollection{
		{"happy-0-0", 0, 0, 1, nil, []int{0}},
		{"happy-0-left", 0, 0, 1, nil, []int{0, 1, 1, 0}},
		{"happy-0-right", 0, 2, 3, nil, []int{1, 1, 0}},
		{"happy-0-middle", 0, 2, 3, nil, []int{1, 1, 0, 2, 2}},
		{"happy-basic", 1, 0, 1, nil, []int{1}},
		{"happy-all", 6, 0, 6, nil, []int{1, 1, 1, 1, 1, 1}},
		{"happy-inside", 8, 1, 5, nil, []int{1, 2, 2, 2, 2, 1}},
		{"happy-left", 8, 0, 4, nil, []int{2, 2, 2, 2, 1}},
		{"happy-right", 8, 1, 5, nil, []int{1, 2, 2, 2, 2}},
		{"happy-right-only", 10, 4, 5, nil, []int{1, 2, 2, 2, 10}},
		{"happy-zeros-in-between", 20, 1, 6, nil, []int{1, 10, 0, 0, 0, 10}},
	}

	for _, t := range tests {
		l, r, err := FindSum(t.a, t.sum)

		if l != t.left || r != t.right || err != t.err {
			fmt.Printf("%v: sum:%v, a:%v, expected l:%v, r:%v, err:'%v', found l:%v, r:%v, err:'%v'\n",
				t.name, t.sum, t.a, t.left, t.right, t.err, l, r, err)
			testing.Fail()

		}
	}

}

func TestNotFound(testing *testing.T) {
	tests := TestCollection{
		{"notFound", 25, 0, 0, errNotFound, []int{1, 1, 2, 5, 10, 20}},
		{"notFound-empty", 0, 0, 0, errNotFound, []int{}},
		{"notFound-zero", 0, 0, 0, errNotFound, []int{1}},
	}

	for _, t := range tests {
		l, r, err := FindSum(t.a, t.sum)

		if l != t.left || r != t.right || err != t.err {
			fmt.Printf("%v: sum:%v, a:%v, expected l:%v, r:%v, err:'%v', found l:%v, r:%v, err:'%v'\n",
				t.name, t.sum, t.a, t.left, t.right, t.err, l, r, err)
			testing.Fail()

		}
	}

}
