package algorithms

// Merge two arrays
func Merge (left, right []int) ([]int){
	list := make([]int, len(left) + len(right))
	i := 0
	j := 0
	k := 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			list[k] = left[i]
			k++
			i++
		} else {
			list[k] = right[j]
			k++
			j++
		}
	}

	for i < len(left) {
		list[k] = left[i]
		k++
		i++
	}

	for j < len(right) {
		list[k] = right[j]
		k++
		j++
	}
	return list
}

// Merge two arrays in parallel
func MergeP (left, right []int) ([]int){
	//TODO
	return Merge(left, right)
}

// Merge sort with no extra allocation
func MergeSort(data []int) ([]int){
	if len(data) > 1 {
		middle := len(data)/2
		left := MergeSort(data[:middle])
		right := MergeSort(data[middle:])
		return Merge(left, right)
	}
	
	return data
}

// Parallel Merge sort with no extra allocation
func PMergeSort(data []int) ([]int){
	//TODO
	return MergeSort(data)
}



// Merge sort with no extra allocation
func MergeSortFast(data []int){
	if len(data) > 1 {
		middle := len(data)/2
		MergeSortFast(data[:middle])
		MergeSortFast(data[middle:])
		QuickSort(data)
		//MergeFast(data, middle)
	}
}
