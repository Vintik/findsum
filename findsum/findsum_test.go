package findsum

import (
	"fmt"
	"math/rand"
	"testing"
)

func ExampleFindSum() {
	a := []int64{1, 2, 3, 0, 1}

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
	a := []int64{1, 2, 10, 0, 1}

	_, _, err := FindSum(a, 8)
	fmt.Printf("sum:8 in a:%v : %v\n", a, err)
	// Output: sum:8 in a:[1 2 10 0 1] : Sum is not found
}

func TestHappyPath(t *testing.T) {
	testHappyPath(t, FindSum)
}

func TestHappyPathBruteForce(t *testing.T) {
	testHappyPath(t, FindSumBruteForce)
}

func TestNotFound(t *testing.T) {
	testNotFound(t, FindSumBruteForce)
}

func TestNotFoundBruteForce(t *testing.T) {
	testNotFound(t, FindSumBruteForce)
}

var bench = struct {
	sum int64
	a   []int64
}{}

func init() {
	bench.a = make([]int64, 256*1024)
	for i, _ := range bench.a {
		bench.a[i] = int64(i)
	}

	l := rand.Intn(len(bench.a) - 16*1024)
	r := l + 16*1024
	for _, v := range bench.a[l:r] {
		bench.sum = bench.sum + v
	}
}

func BenchmarkFindSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _, err := FindSum(bench.a, bench.sum)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkFindSumBruteForce(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _, err := FindSumBruteForce(bench.a, bench.sum)
		if err != nil {
			b.Fatal(err)
		}
	}
}

type TestCollection []struct {
	name        string
	sum         int64
	left, right int
	err         error
	a           []int64
}

func testHappyPath(t *testing.T, fs func([]int64, int64) (int, int, error)) {
	tests := TestCollection{
		{"happy-0-0", 0, 0, 1, nil, []int64{0}},
		{"happy-0-left", 0, 0, 1, nil, []int64{0, 1, 1, 0}},
		{"happy-0-right", 0, 2, 3, nil, []int64{1, 1, 0}},
		{"happy-0-middle", 0, 2, 3, nil, []int64{1, 1, 0, 2, 2}},
		{"happy-basic", 1, 0, 1, nil, []int64{1}},
		{"happy-all", 6, 0, 6, nil, []int64{1, 1, 1, 1, 1, 1}},
		{"happy-inside", 8, 1, 5, nil, []int64{1, 2, 2, 2, 2, 1}},
		{"happy-left", 8, 0, 4, nil, []int64{2, 2, 2, 2, 1}},
		{"happy-right", 8, 1, 5, nil, []int64{1, 2, 2, 2, 2}},
		{"happy-right-only", 10, 4, 5, nil, []int64{1, 2, 2, 2, 10}},
		{"happy-zeros-in-between", 20, 1, 6, nil, []int64{1, 10, 0, 0, 0, 10}},
	}

	for _, test := range tests {
		l, r, err := fs(test.a, test.sum)

		if l != test.left || r != test.right || err != test.err {
			fmt.Printf("%v: sum:%v, a:%v, expected l:%v, r:%v, err:'%v', found l:%v, r:%v, err:'%v'\n",
				test.name, test.sum, test.a, test.left, test.right, test.err, l, r, err)
			t.Fail()

		}
	}

}

func testNotFound(t *testing.T, fs func([]int64, int64) (int, int, error)) {
	tests := TestCollection{
		{"notFound", 25, 0, 0, errNotFound, []int64{1, 1, 2, 5, 10, 20}},
		{"notFound-empty", 0, 0, 0, errNotFound, []int64{}},
		{"notFound-zero", 0, 0, 0, errNotFound, []int64{1}},
	}

	for _, test := range tests {
		l, r, err := fs(test.a, test.sum)

		if l != test.left || r != test.right || err != test.err {
			fmt.Printf("%v: sum:%v, a:%v, expected l:%v, r:%v, err:'%v', found l:%v, r:%v, err:'%v'\n",
				test.name, test.sum, test.a, test.left, test.right, test.err, l, r, err)
			t.Fail()

		}
	}

}
