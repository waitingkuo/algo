package insertion

func Sort(a []int) []int {
	for j := 1; j < len(a); j++ {
		key := a[j]
		// Insert a[j] into the sorted sequence a[0] ~ a[j-1]
		i := j - 1
		for ; i >= 0 && a[i] > key; i-- {
			a[i+1] = a[i]
		}
		a[i+1] = key
	}

	return a
}
