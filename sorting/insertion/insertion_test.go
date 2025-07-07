package insertion

import (
	"testing"
	"reflect"
)


func TestInsertion(t *testing.T) {
	arr := []int{6,1,3,2,5,4}
	
        t.Log("Before: ", arr)	

	insertion(arr)

	t.Log("After: ", arr)

	expected := []int{1,2,3,4,5,6}

	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("Insertion faild, Got: %v, expected: %v", arr, expected)
	}

}
