package algorithms


func swap(a, b *int) {
    *a, *b = *b, *a
}

func QuickSort(data []int) {

	var n int = len(data)
	var middle int = (n-1)/2

	var pivot int = data[middle]
	
	var first int = 0
	var last int = n-1
	
	//partition
	for first <= last {
		for data[first] < pivot { first++;}
		for data[last] > pivot { last--;}
		if first <= last {
			swap(&data[first],&data[last]);
			first++
			last--
		}
	}
	
	if last > 0 {
		QuickSort(data[:last+1])
	}
	if first < n {
		QuickSort(data[first:])
	}	
}
