package slices

func Deduplicate[T string](slice []T) []T {
	ids := make(map[T]bool)
	l := []T{}
	for _, item := range slice {
		if _, v := ids[item]; !v {
			ids[item] = true
			l = append(l, item)
		}
	}
	return l
}
