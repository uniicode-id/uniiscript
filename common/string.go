package common

// Hash with DJB2 algorithm
func Hash(str string) int {
	hash := 5381
	for _, c := range str {
		hash = ((hash << 5) + hash) + int(c)
	}
	return hash
}
