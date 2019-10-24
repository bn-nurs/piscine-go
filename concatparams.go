package piscine

func ConcatParams(args []string) string {
	var s = ""
	var lastindex = 0
	for index := range args {
		lastindex = index
	}
	for index, v := range args {
		s += v
		if index != lastindex {
			s += "\n"
		}
	}
	return s
}
