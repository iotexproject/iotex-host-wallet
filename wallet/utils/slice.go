package utils

// Contain ...
func Contain(s []string, e string) bool {
	for _, t := range s {
		if t == e {
			return true
		}
	}
	return false
}
