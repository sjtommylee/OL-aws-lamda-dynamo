package util

type StringSet map[string]bool

// create a set of strings from a given string slice
func NewStringFromSlice(slice []string) StringSet {
	str := make(StringSet, len(slice))
	for _, value := range slice {
		str[value] = true
	}
	return str
}
