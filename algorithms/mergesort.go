package algorithms
//import "runtime"
//import "sync"
/*
// Merge two arrays
func Merge(left, right []int) ([]int){
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

// Merge sort with no extra allocation
func MergeSort(data []int) ([]int){
	if len(data) > 1 {
		middle := len(data)/2
		left := MergeSort(data[:middle])
		right := MergeSort(data[middle:])
		return Merge(left, right)
	}
	
	return data
}*/

// Merge two arrays
func Merge(list, left, right []int){
	//list := make([]int, len(left) + len(right))
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
}

// Merge sort with no extra allocation
func MergeSort(data []int) ([]int){
	if len(data) > 1 {
		middle := len(data)/2
		MergeSort(data[:middle])
		MergeSort(data[middle:])
		Merge(data, data[:middle], data[middle:])
	}
	return data
}

/*
// Parallel Merge sort with no extra allocation
func PMergeSort(data []int){
	
	var threads int = runtime.NumCPU()
	var n int = len(data)
	var block int = n/threads

	//runtime.GOMAXPROCS(threads)

	var wg sync.WaitGroup
	
	for i:=0; i < threads; i++ {
		wg.Add(1)
		go func( pos, b int, data []int){
			start := pos*b
			end := pos*b + b
			QuickSort(data[start:end])
			wg.Done()
		}(i, block, data)
	}
	wg.Wait()

	for i:=threads; i > 1; i*=0.5 {
		block = n/i
		
	}
	
	
}*/



// Merge sort with no extra allocation
func MergeSortNoAlloc(data []int){
	if len(data) > 1 {
		middle := len(data)/2
		MergeSortNoAlloc(data[:middle])
		MergeSortNoAlloc(data[middle:])
		QuickSort(data)
		//MergeNoAlloc(data, middle)
	}
}
