package corpus

// WordFrequency stores a word and its frequency
type WordFrequency struct {
	Word  string
	Count int
}

// WordFrequencies stores a slice of WordFrequencies
type WordFrequencies []WordFrequency

func (slice WordFrequencies) Len() int {
	return len(slice)
}

func (slice WordFrequencies) Less(x, y int) bool {
	return slice[x].Count < slice[y].Count
}

func (slice WordFrequencies) Swap(x, y int) {
	slice[x], slice[y] = slice[y], slice[x]
}
