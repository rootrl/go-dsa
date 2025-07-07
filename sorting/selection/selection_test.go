package selection

import (
	"reflect"
	"testing"
)

func TestSelection(t *testing.T) {
	arr := []int{6, 1, 3, 2, 5, 4}

	t.Log("Before: ", arr)

	selection(arr)

	t.Log("After: ", arr)

	expected := []int{1, 2, 3, 4, 5, 6}

	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("Selection faild, Got: %v, expected: %v", arr, expected)
	}

}
