package slice

type Insertable []int

// InsertValueAtIndex inserts a value at the specified index
func (a Insertable) InsertValueAtIndex(value int, index int) Insertable {
	a = append(a[:index+1], a[index:]...)
	a[index] = value
	return a
}
