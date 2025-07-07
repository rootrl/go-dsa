package selection

func selection(arr []int) {

	l := len(arr)

	for i := 0; i < l-1; i++ {

		minIndex := i
		for j := i + 1; j < l; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			arr[minIndex], arr[i] = arr[i], arr[minIndex]
		}
	}

}
