package algorithms

import (
	"testing"
	"sort"
	"math/rand"
	"fmt"
)

func mergeSortGeneral(t *testing.T, size int) (int){
	array := make([]int, size)
	a1 := make([]int, size)
	a2 := make([]int, size)
	//rand.Read(a1)
	for i := 0; i < size; i++ {
		array[i] = rand.Intn(size)
	}
	
	copy(a2, array)
	copy(a1, array)
	
	sort.Ints(a1)
	a2 = MergeSort(a2)
	
	for i := 0; i < size; i++ {
		if a1[i] != a2[i]{
			fmt.Println("unsorted:\n", array)
			fmt.Println("built-in:\n", a1)
			fmt.Println("mergesort:\n", a2)
			t.Error("pos:", i, "-built-in result != mergesort result")
			return -1
		}
	}
	return 0
}

func TestMergeSort(t *testing.T){
	var size int = 1000000
	
	for i := 1; i < size; i *= 10 {
		if  mergeSortGeneral(t, i) == -1 {
			break
		}
	}
}


func BenchmarkMergeSortBuiltin1(b * testing.B) {

	var size int = 20000000
	array := make([]int, size)
	
	for i := 0; i < size; i++ {
		array[i] = rand.Intn(size)
	}

	b.ReportAllocs()
	b.ResetTimer()
	
	sort.Ints(array)
}

func BenchmarkMergeSort1(b * testing.B) {

	var size int = 20000000
	array := make([]int, size)
	
	for i := 0; i < size; i++ {
		array[i] = rand.Intn(size)
	}

	b.ReportAllocs()
	b.ResetTimer()
	
	array = MergeSort(array)
}
