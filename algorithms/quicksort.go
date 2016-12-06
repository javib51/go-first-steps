package algorithms
import "runtime"
import "sync"

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

func QuickSortParallel(data []int) {
	var numCPUs int = runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	quickSortParallel(data, numCPUs, wg)	
	wg.Wait()
}

func quickSortParallel(data []int, numCPUs int, wg *sync.WaitGroup) {
	if numCPUs > 1 {
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
		
		wg1 := new(sync.WaitGroup)
		if last > 0 {
			wg1.Add(1)
			go quickSortParallel(data[:last+1], numCPUs/2, wg1)
		}
		if first < n {
			wg1.Add(1)
			go quickSortParallel(data[first:], numCPUs/2, wg1)
		}
		wg1.Wait()
	} else {
		QuickSort(data)
	}
	wg.Done()
}

/*'Engineering a sort function' by Jon L. Bentley and M. Douglas Mcilroy*/
func QuickSortFaster(data []int) {
}
