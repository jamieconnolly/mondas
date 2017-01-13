package cli

// Args represents a list of command line arguments.
type Args []string

// First returns the first argument (same as a.Index(0)).
func (a Args) First() string {
	return a.Index(0)
}

// Index returns the i'th argument. It returns an empty
// string if the requested argument does not exist.
func (a Args) Index(i int) string {
	if i < 0 || i >= len(a) {
		return ""
	}
	return a[i]
}

// Len returns the number of arguments.
func (a Args) Len() int {
	return len(a)
}
