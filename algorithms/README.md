#Benchmark
```
$ go test -bench=.
BenchmarkBuiltin1k-8              	   10000	    267515 ns/op	       0 B/op	       0 allocs/op
BenchmarkMergeSort1k-8            	   10000	    245811 ns/op	  194758 B/op	    1000 allocs/op
BenchmarkMergeSortNoAlloc1k-8     	   10000	    280681 ns/op	       0 B/op	       0 allocs/op
BenchmarkQuickSort1k-8            	   20000	     97185 ns/op	       0 B/op	       0 allocs/op
BenchmarkBubbleSort1k-8           	     100	 155477193 ns/op	       0 B/op	       0 allocs/op
BenchmarkSelectionSort1k-8        	     100	  94789903 ns/op	       0 B/op	       0 allocs/op
BenchmarkBuiltin10k-8             	    1000	   2698827 ns/op	       0 B/op	       0 allocs/op
BenchmarkMergeSort10k-8           	     500	   3123647 ns/op	 1867577 B/op	    9999 allocs/op
BenchmarkMergeSortNoAlloc10k-8    	    1000	   2794291 ns/op	       0 B/op	       0 allocs/op
BenchmarkQuickSort10k-8           	    2000	    965543 ns/op	       0 B/op	       0 allocs/op
BenchmarkBubbleSort10k-8          	      10	1541981193 ns/op	       0 B/op	       0 allocs/op
BenchmarkSelectionSort10k-8       	      20	1901538154 ns/op	       0 B/op	       0 allocs/op
BenchmarkBuiltin1M-8              	       5	 259931336 ns/op	       6 B/op	       0 allocs/op
BenchmarkMergeSort1M-8            	       5	 270462166 ns/op	186757734 B/op	  999999 allocs/op
BenchmarkMergeSortNoAlloc1M-8     	       5	 257649926 ns/op	       0 B/op	       0 allocs/op
BenchmarkQuickSort1M-8            	      20	  98625265 ns/op	       0 B/op	       0 allocs/op
BenchmarkBuiltin20M-8             	       1	5603536131 ns/op	      32 B/op	       1 allocs/op
BenchmarkMergeSort20M-8           	       1	4230092072 ns/op	4055166976 B/op	19999999 allocs/op
BenchmarkMergeSortNoAlloc20M-8    	       1	6066574208 ns/op	       0 B/op	       0 allocs/op
BenchmarkQuickSort20M-8           	       1	1917423821 ns/op	       0 B/op	       0 allocs/op
BenchmarkMergeSortNoAlloc100M-8   	       1	34840641722 ns/op	       0 B/op	       0 allocs/op
BenchmarkBuiltin100M-8            	       1	30529100503 ns/op	      32 B/op	       1 allocs/op
BenchmarkQuickSort100M-8          	       1	10545902403 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/javib51/go-first-steps/algorithms	250.125s

```
