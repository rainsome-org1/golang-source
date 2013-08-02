package bubblesort

func BubbleSort(values []int) {
	var flags bool = true

	for i:= 0; i < len(values); i++ {
		flags = true

		for j:= 0; j < len(values); j++ {
			if values[j] > values[j + 1] {
				values[j], values[j+1] = values[j +1], values[j]
				flags = false
			}
			if flags == true {
				break
			}
		}
	}
}
