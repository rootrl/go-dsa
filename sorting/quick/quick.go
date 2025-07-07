package quick

func quick(arr []int) {
	quickSort(arr, 0, len(arr) - 1)
}

func quickSort(arr []int, l, h int) {
	if l < h {
		p := partition(arr, l, h)

		quickSort(arr, l, p - 1)
		quickSort(arr, p + 1, h)
	}
}

func partition(arr []int, l, h int) int{
	p := h
	pv := arr[h]

	i := l - 1

	for j := l; j < h; j++ {
		if arr[j] < pv {
			i++ 
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[p] = arr[p], arr[i+1]

	return i+1
}
