package helper

func Clamp(value, min, max int) int {
	if value < min {
		return min
	}
	if value > max {
		return max
	}

	return value
}

func SumOfList(list []int) int {
	sum := 0
	for _, n := range list {
		sum += n
	}
	return sum
}
