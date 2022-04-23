package powerset_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/chanced/powerset"
)

type test[T any] struct {
	set    []T
	expect [][]T
}

func Test(t *testing.T) {
	intTests := []test[int]{
		{[]int{1, 2, 3}, [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}}},
		{[]int{1, 2, 3, 4}, [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}, {4}, {1, 4}, {2, 4}, {1, 2, 4}, {3, 4}, {1, 3, 4}, {2, 3, 4}, {1, 2, 3, 4}}},
	}

	for _, test := range intTests {
		res := powerset.Compute(test.set)
		if len(res) != len(test.expect) {
			t.Errorf("Expected %v, got %v", test.expect, res)
		}
		for i, r := range res {
			if !reflect.DeepEqual(r, test.expect[i]) {
				t.Errorf("Expected result[%d] to be %v, got %v", i, test.expect[i], r)
			}
		}

	}

	stringTests := []test[string]{
		{[]string{"a", "b", "c"}, [][]string{{}, {"a"}, {"b"}, {"a", "b"}, {"c"}, {"a", "c"}, {"b", "c"}, {"a", "b", "c"}}},
		{[]string{"a", "b", "c", "d"}, [][]string{{}, {"a"}, {"b"}, {"a", "b"}, {"c"}, {"a", "c"}, {"b", "c"}, {"a", "b", "c"}, {"d"}, {"a", "d"}, {"b", "d"}, {"a", "b", "d"}, {"c", "d"}, {"a", "c", "d"}, {"b", "c", "d"}, {"a", "b", "c", "d"}}},
	}
	for _, test := range stringTests {
		res := powerset.Compute(test.set)
		if len(res) != len(test.expect) {
			t.Errorf("Expected %v, got %v", test.expect, res)
		}
		for i, r := range res {
			if !reflect.DeepEqual(r, test.expect[i]) {
				t.Errorf("Expected result[%d] to be %v, got %v", i, test.expect[i], r)
			}
		}
	}
}

func ExampleCompute() {
	fmt.Println(powerset.Compute([]string{"a", "b", "c"}))

	fmt.Println(powerset.Compute([]int{1, 2, 3}))
	// output: [[] [a] [b] [a b] [c] [a c] [b c] [a b c]]
	// [[] [1] [2] [1 2] [3] [1 3] [2 3] [1 2 3]]
}
