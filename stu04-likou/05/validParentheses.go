package main

func isValid(s string) bool {
	b := false

	// 计数器
	num := 0

	for i, _ := range s {
		if string(s[i]) == "(" {
			num++
		}

	}

	return b

}

func main() {
	isValid("123")
}
