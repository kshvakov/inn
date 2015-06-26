package inn

var (
	k10   = [...]int{2, 4, 10, 3, 5, 9, 4, 6, 8}
	k11_1 = [...]int{7, 2, 4, 10, 3, 5, 9, 4, 6, 8}
	k11_2 = [...]int{3, 7, 2, 4, 10, 3, 5, 9, 4, 6, 8}
	dMap  = map[rune]int{
		48: 0,
		49: 1,
		50: 2,
		51: 3,
		52: 4,
		53: 5,
		54: 6,
		55: 7,
		56: 8,
		57: 9,
	}
)

func Check(in string) bool {

	c := len(in)

	if c == 0 {

		return true
	}

	var (
		sum   = 0
		value = make([]int, 0, c)
	)

	for _, v := range in {

		if r, found := dMap[v]; found {

			value = append(value, r)

		} else {

			return false
		}
	}

	switch c {

	case 10:

		for k, v := range k10 {

			sum += int(value[k]) * v
		}

		return sum%11 == value[9]

	case 12:

		for k, v := range k11_1 {

			sum += value[k] * v
		}

		if sum%11 != value[10] {

			return false
		}

		sum = 0

		for k, v := range k11_2 {

			sum += value[k] * v
		}

		return sum%11 == value[11]

	}

	return false
}
