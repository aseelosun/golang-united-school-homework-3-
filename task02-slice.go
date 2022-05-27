package homework

func reverse(input []int64) (result []int64) {
	var res []int64
	for i := len(input) - 1; i >= 0; i-- {
		res = append(res, input[i])
	}
	return res
}
