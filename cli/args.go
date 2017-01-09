package cli

type Args []string

func (a Args) First() string {
	return a.Index(0)
}

func (a Args) Index(i int) string {
	if i < 0 || i >= len(a) {
		return ""
	}
	return a[i]
}

func (a Args) Len() int {
	return len(a)
}
