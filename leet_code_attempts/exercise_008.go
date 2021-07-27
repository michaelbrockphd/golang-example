package leet_code_attempts

const offset = 32

const ascii_range = 25

//const index_9_ascii = 57

func myAtoi(str string) int32 {
	rtn := int32(0)

	allowed := make([]bool, ascii_range)

	i := 24

	for 14 < i {
		allowed[i] = true
		i -= 1
	}

	allowed[13] = true
	allowed[12] = true
	allowed[10] = true
	allowed[0] = true

	for _, c := range str {
		switch o := int32(c - offset); {
		case 14 < o:
			rtn = rtn*10 + (o - 16)
		}
	}

	return rtn
}
