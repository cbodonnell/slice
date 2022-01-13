package slice

type Insertable []int

// InsertValueAtIndex inserts a value at the specified index
// and returns the new slice
func (a Insertable) InsertValueAtIndex(value int, index int) Insertable {
	a = append(a[:index+1], a[index:]...)
	a[index] = value
	return a
}
