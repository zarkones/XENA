package patch

func Patch(value, placeholder string) string {
	if len(value) > len(placeholder) {
		return value[:len(placeholder)]
	}
	patched := value
	for len(patched) != len(placeholder) {
		patched += Replacement
	}
	return patched
}
