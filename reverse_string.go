package reverse_string

func ReverseString(input string) (output string) {
	// solution goes here
	r := []rune(input)

	for indx1, indx2 := 0, len(r)-1; indx1 < indx2; {
		r[indx1], r[indx2] = r[indx2], r[indx1]
		indx1++
		indx2--
	}
	output = string(r)
	return
}
