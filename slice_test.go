package slice

import (
	"fmt"
	"testing"
)

// TestInsertValueAtIndex verifies that an integer can be
// inserted into an Insertable []int and that the proper
// value is returned
func TestInsertValueAtIndex(t *testing.T) {
	a := Insertable{1, 2, 4, 5}
	length := len(a)
	value := 3
	index := 2
	a = a.InsertValueAtIndex(value, index)

	if a[index] != value || len(a) != length+1 {
		t.Errorf("Failed to insert %d at index %d", value, index)
	}
}

// ExampleInsertable_InsertValueAtIndex provides an example of the
// InsertValueAtIndex method of the Insertable type
func ExampleInsertable_InsertValueAtIndex() {
	a := Insertable{1, 2, 4, 5}
	value := 3
	index := 2
	a = a.InsertValueAtIndex(value, index)
	fmt.Println(a)
	// Output: [1 2 3 4 5]
}
