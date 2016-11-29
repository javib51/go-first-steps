package algorithms
//import "runtime"
//import "sync"

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
}
/*
// Merge two arrays
func Merge(list [] int){
	var rs int = len(list)/2
	var li int = 0
	var ri int = rs

	for true {
		if list[li] <= list[ri] {
			if  ri == rs {
				li++
			} else {
				list[li], list[ri] = list[ri], list[li]
				ri = rs
			}
		} else {
			ri++
		}

		if  ri >= len(list) || ri <= li {
			break
		}
	}
}

// Merge sort with no extra allocation
func MergeSort(data []int) ([]int){
	if len(data) > 1 {
		middle := len(data)/2
		MergeSort(data[:middle])
		MergeSort(data[middle:])
		Merge(data)
	}
	return data
}
*/
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

//Jyrki Katajainen, Tomi Pasanen, Jukka Teuhola. ``Practical in-place mergesort'' algorithm
// Space: O(square(n,2)) 
// Time: O(n log n)
func swapvect(data []int, i, j int) {
	data[i], data[j] = data[j], data[i]
}

// Merge two arrays
func MergeNoAlloc(data []int, i, m, j, n, w int) {
	for i < m && j < n {
		if data[i] < data[j] {
			swapvect(data, w, i)
			i++
		} else {
			swapvect(data, w, j)
			j++
		}
		w++
	}

	for i < m {
		swapvect(data, w, i)
		w++
		i++
	}
	
	for j < n {
		swapvect(data, w, j)
		w++
		j++
	}
}

func mergesortNoAlloc(data []int, l, u ,w int) {
	var m int
	if u - l > 1 {
		m = l + (u - l) / 2
		MergeSortNoAlloc(data[l:m])
		MergeSortNoAlloc(data[m:u])
		MergeNoAlloc(data, l, m, m, u, w)
	} else {
		for l < u {
			swapvect(data, l, w)
			l++
			w++
		}
	}
}

// Merge sort with no extra allocation
func MergeSortNoAlloc(data []int) {
	
	var size int = len(data)
	var m int
	var n int
	var w int
	
	if size > 1 {
		m = size / 2
		w = size - m
		mergesortNoAlloc(data, 0, m, w) // the last half contains sorted elements
		for w > 2 {
			n = w
			w = (n + 1) / 2
			mergesortNoAlloc(data, w, n, 0) 
			MergeNoAlloc(data, 0, n - w, n, size, w)
		}
		
		for n = w; n > 0; n-- { //switch to insertion sort
			for m = n; m < size && data[m] < data[m-1]; m++ {
				swapvect(data, m, m - 1);
			}
		}
    }
}

