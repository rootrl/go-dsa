package bubble

func bubble(arr []int) {
	l := len(arr)

	for i := 0; i < l; i++ {
		swaped := false
		for j := 0; j < l - i - 1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swaped = true
			}
		}

		if !swaped {
			break
		}
	}
}
