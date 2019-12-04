package myUtil

func SplitJSON(strJSON string) (ss []string) {

	flag := false
	start, count := 0, 0

	for i, c := range strJSON {
		if c == '{' {
			if flag == false {
				flag = true
				start = i
				count = 1
			} else {
				count++
			}
		} else if c == '}' {
			count--

			if count == 0 {
				flag = false
				ss = append(ss, strJSON[start:i+1])
			}
		}
	}
	return ss
}
