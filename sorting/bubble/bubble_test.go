package bubble

import (
	"reflect"
	"testing"
)

func TestBubble(t *testing.T) {
	arr := []int{6, 1, 3, 2, 5, 4}

	t.Log("Before: ", arr)

	bubble(arr)

	t.Log("After: ", arr)

	expected := []int{1, 2, 3, 4, 5, 6}

	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("Bubble faild, Got: %v, expected: %v", arr, expected)
	}

}
