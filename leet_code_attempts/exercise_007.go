package leet_code_attempts

func reverse(src int32) int32 {
	rtn := int32(0)

	tmp := int32(0)

	rem := int32(0)

	cur := src

	for cur != 0 {
		rem = (cur % 10)

		tmp = ((rtn * 10) + rem)

		if ((tmp - rem) / 10) != rtn {
			cur = 0
			rtn = 0

		} else {
			rtn = tmp

			cur /= 10
		}
	}

	return rtn
}
