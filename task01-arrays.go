package homework

func average(input [15]float32) (result float32) {
	var sum float32
	for i := 0; i < len(input); i++ {
		sum += input[i]
	}
	//Place your code here
	avg := float32(sum) / float32(len(input))
	return avg
}
