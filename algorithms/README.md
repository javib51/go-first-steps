#Benchmark
```
$ go test -bench=.
BenchmarkBuiltin1k-8             	   10000	    272381 ns/op	       0 B/op	       0 allocs/op
BenchmarkMergeSort1k-8           	   10000	    237085 ns/op	  194758 B/op	    1000 allocs/op
BenchmarkMergeSortNoAlloc1k-8    	    1000	   2963341 ns/op	       0 B/op	       0 allocs/op
BenchmarkQuickSort1k-8           	   20000	     99908 ns/op	       0 B/op	       0 allocs/op
BenchmarkBubbleSort1k-8          	     100	 159580817 ns/op	       0 B/op	       0 allocs/op
BenchmarkSelectionSort1k-8       	     100	  97742567 ns/op	       0 B/op	       0 allocs/op
BenchmarkBuiltin10k-8            	     500	   2612607 ns/op	       0 B/op	       0 allocs/op
BenchmarkMergeSort10k-8          	     500	   2894314 ns/op	 1867577 B/op	    9999 allocs/op
BenchmarkMergeSortNoAlloc10k-8   	     100	  23209897 ns/op	       0 B/op	       0 allocs/op
BenchmarkQuickSort10k-8          	    2000	   1003968 ns/op	       0 B/op	       0 allocs/op
BenchmarkBubbleSort10k-8         	      10	1599452079 ns/op	       0 B/op	       0 allocs/op
BenchmarkSelectionSort10k-8      	      20	1954045911 ns/op	       0 B/op	       0 allocs/op
BenchmarkBuiltin1M-8             	       5	 260554146 ns/op	       6 B/op	       0 allocs/op
BenchmarkMergeSort1M-8           	       5	 265356415 ns/op	186757734 B/op	  999999 allocs/op
BenchmarkMergeSortNoAlloc1M-8    	       1	2193789300 ns/op	       0 B/op	       0 allocs/op
BenchmarkQuickSort1M-8           	      20	 104114982 ns/op	       0 B/op	       0 allocs/op
BenchmarkBuiltin20M-8            	       1	5691058304 ns/op	      32 B/op	       1 allocs/op
BenchmarkMergeSort20M-8          	       1	4612089187 ns/op	4055167056 B/op	20000000 allocs/op
BenchmarkMergeSortNoAlloc20M-8   	       1	201918663092 ns/op	       0 B/op	       0 allocs/op
BenchmarkQuickSort20M-8          	       1	1993231695 ns/op	       0 B/op	       0 allocs/op
BenchmarkBuiltin100M-8           	       1	31204786322 ns/op	      32 B/op	       1 allocs/op
BenchmarkQuickSort100M-8         	       1	10841803837 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/javib51/go-first-steps/algorithms	408.057s

```
