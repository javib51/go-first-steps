package algorithms

import (
	"testing"
	"sort"
	"math/rand"
	"fmt"
)

//initTest init two arrays
func initTest(size int) ([]int, []int){
	array := make([]int, size)
	a1 := make([]int, size)
	a2 := make([]int, size)
	//rand.Read(a1)
	for i := 0; i < size; i++ {
		array[i] = rand.Intn(size)
	}
	
	copy(a2, array)
	copy(a1, array)
	
	return a1, a2
}

func initBenchmark(size int) []int{
	
	array := make([]int, size)
	for i := 0; i < size; i++ {
		array[i] = rand.Intn(size)
	}
	
	return array
}

func testArray(t *testing.T, size int, a1, a2 [] int) int {
	for i := 0; i < size; i++ {
		if a1[i] != a2[i]{
			fmt.Println("built-in:\n", a1)
			fmt.Println("mergesort:\n", a2)
			t.Error("pos:", i, "-built-in result != mergesort result")
			return -1
		}
	}
	return 0
}

func mergeSortGeneral(t *testing.T, size int) (int){
	a1, a2 := initTest(size)	
	sort.Ints(a1)
	a2 = MergeSort(a2)

	return testArray(t, size,  a1, a2)
}

func TestMergeSort(t *testing.T){
	var size int = 100000
	
	for i := 1; i <= size; i *= 10 {
		if  mergeSortGeneral(t, i) == -1 {
			break
		}
	}
}

func quickSortGeneral(t *testing.T, size int) (int){
	a1, a2 := initTest(size)	
	sort.Ints(a1)
	QuickSort(a2)

	return testArray(t, size, a1, a2)
}

func TestQuickSort(t *testing.T){
	var size int = 1000000
	
	for i := 1; i <= size; i *= 10 {
		if  mergeSortGeneral(t, i) == -1 {
			break
		}
	}
}

func BenchmarkBuiltin1k(b * testing.B) {

	var size int = 1000 * b.N 

	array := initBenchmark(size)

	b.ReportAllocs()
	b.ResetTimer()
	
	sort.Ints(array)
}

func BenchmarkMergeSort1k(b * testing.B) {

	var size int =  1000 * b.N

	array := initBenchmark(size)

	b.ReportAllocs()
	b.ResetTimer()
	
	array = MergeSort(array)
}



func BenchmarkQuickSort1k(b * testing.B) {

	var size int =  1000 * b.N

	array := initBenchmark(size)

	b.ReportAllocs()
	b.ResetTimer()
	
	QuickSort(array)
}

func BenchmarkBuiltin10k(b * testing.B) {

	var size int = 10000 * b.N 

	array := initBenchmark(size)

	b.ReportAllocs()
	b.ResetTimer()
	
	sort.Ints(array)
}

func BenchmarkMergeSort10k(b * testing.B) {

	var size int =  10000 * b.N

	array := initBenchmark(size)

	b.ReportAllocs()
	b.ResetTimer()
	
	array = MergeSort(array)
}



func BenchmarkQuickSort10k(b * testing.B) {

	var size int =  10000 * b.N

	array := initBenchmark(size)

	b.ReportAllocs()
	b.ResetTimer()
	
	QuickSort(array)
}

func BenchmarkBuiltin1M(b * testing.B) {

	var size int = 1000000 * b.N

	array := initBenchmark(size)

	b.ReportAllocs()
	b.ResetTimer()
	
	sort.Ints(array)
}

func BenchmarkMergeSort1M(b * testing.B) {

	var size int = 1000000 * b.N

	array := initBenchmark(size)

	b.ReportAllocs()
	b.ResetTimer()
	
	array = MergeSort(array)
}



func BenchmarkQuickSort1M(b * testing.B) {

	var size int = 1000000 * b.N

	array := initBenchmark(size)

	b.ReportAllocs()
	b.ResetTimer()
	
	QuickSort(array)
}

func BenchmarkBuiltin20M(b * testing.B) {

	var size int = 20000000 * b.N

	array := initBenchmark(size)

	b.ReportAllocs()
	b.ResetTimer()
	
	sort.Ints(array)
}

func BenchmarkMergeSort20M(b * testing.B) {

	var size int = 20000000 * b.N

	array := initBenchmark(size)

	b.ReportAllocs()
	b.ResetTimer()
	
	array = MergeSort(array)
}



func BenchmarkQuickSort20M(b * testing.B) {

	var size int = 20000000 * b.N

	array := initBenchmark(size)

	b.ReportAllocs()
	b.ResetTimer()
	
	QuickSort(array)
}

func BenchmarkBuiltin100M(b * testing.B) {

	var size int = 100000000 * b.N

	array := initBenchmark(size)

	b.ReportAllocs()
	b.ResetTimer()
	
	sort.Ints(array)
}

func BenchmarkQuickSort100M(b * testing.B) {

	var size int = 100000000 * b.N

	array := initBenchmark(size)

	b.ReportAllocs()
	b.ResetTimer()
	
	QuickSort(array)
}
