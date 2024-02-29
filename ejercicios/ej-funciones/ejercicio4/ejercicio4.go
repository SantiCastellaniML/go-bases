package main

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func main() {
	nums := []float32{1, 2, 3, 4, 5}

	fn, err := operation(minimum)

	if err != "" {
		println(err)
		return
	}

	result := fn(nums)

	println(result)
}

func operation(operation string) (func([]float32) float32, string) {
	switch operation {
	case minimum:
		return min, ""
	case average:
		return avg, ""
	case maximum:
		return max, ""
	default:
		return nil, "invalid operation"
	}
}

func min(numbers []float32) float32 {
	var min float32 = numbers[0]
	for _, number := range numbers {
		if number < min {
			min = number
		}
	}
	return min
}

func avg(numbers []float32) float32 {
	var sum float32
	for _, number := range numbers {
		sum += number
	}
	return sum / float32(len(numbers))
}

func max(numbers []float32) float32 {
	var max float32 = numbers[0]
	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	return max
}
