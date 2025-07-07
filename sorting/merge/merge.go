package merge

func merge(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	m := len(arr) / 2
	l := merge(arr[:m])
	r := merge(arr[m:])
	return doMerge(l, r)
}

func doMerge(l []int, r []int) []int {

	arr := make([]int, len(l)+len(r))

	li := 0
	ri := 0

	for i := 0; i < len(arr); i++ {

		if li < len(l) && ri < len(r) {
			if l[li] < r[ri] {
				arr[i] = l[li]
				li++
			} else {
				arr[i] = r[ri]
				ri++
			}
		} else {
			if li < len(l) {
				arr[i] = l[li]
				li++
			}

			if ri < len(r) {
				arr[i] = r[ri]
				ri++
			}
		}

	}

	return arr
}
