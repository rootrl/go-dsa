package merge

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	arr := []int{6, 1, 3, 2, 5, 4}

	t.Log("Before: ", arr)

	arr = merge(arr)

	t.Log("After: ", arr)

	expected := []int{1, 2, 3, 4, 5, 6}

	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("Merge faild, Got: %v, expected: %v", arr, expected)
	}

}
