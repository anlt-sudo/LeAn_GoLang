package main

func checkRevertString(s1, s2 string) bool {
	if len(s1) == 0 || len(s2) == 0 || len(s1) != len(s2) {
		return false
	}

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[len(s2)-1-i] {
			return false
		}
	}
	return true
}

func checkSymmetricalString(s string) bool {
	if len(s) == 0 {
		return false
	}

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	println("Result check revert string: ", checkRevertString("hhhhhh", "hhhhhh"))
	println("Result check Symmetrical string: ", checkSymmetricalString("hhhhhh"))
	println("Result check revert string: ", checkRevertString("hhhhhh", "hhh1hh"))
	print("Result check Symmetrical string: ", checkSymmetricalString("hhhohhh"))
}
