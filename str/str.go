package str

func In(src string, strs ...string) bool {
	for _, v := range strs {
		if src == v {
			return true
		}
	}
	return false
}
