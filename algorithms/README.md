#Benchmark
```
go test -bench=.
BenchmarkBuiltin1k-8          	   10000	    265353 ns/op	       0 B/op	       0 allocs/op
BenchmarkMergeSort1k-8        	    5000	    261726 ns/op	  186757 B/op	    1000 allocs/op
BenchmarkQuickSort1k-8        	   20000	     96323 ns/op	       0 B/op	       0 allocs/op
BenchmarkBubbleSort1k-8       	     100	 154205565 ns/op	       0 B/op	       0 allocs/op
BenchmarkSelectionSort1k-8    	     100	  94602331 ns/op	       0 B/op	       0 allocs/op
BenchmarkBuiltin10k-8         	    1000	   2675328 ns/op	       0 B/op	       0 allocs/op
BenchmarkMergeSort10k-8       	     500	   2990902 ns/op	 1867577 B/op	    9999 allocs/op
BenchmarkQuickSort10k-8       	    2000	    964353 ns/op	       0 B/op	       0 allocs/op
BenchmarkBubbleSort10k-8      	      10	1552875770 ns/op	       0 B/op	       0 allocs/op
BenchmarkSelectionSort10k-8   	      20	1891607448 ns/op	       0 B/op	       0 allocs/op
BenchmarkBuiltin1M-8          	       5	 256496034 ns/op	       6 B/op	       0 allocs/op
BenchmarkMergeSort1M-8        	       5	 293447200 ns/op	186757734 B/op	  999999 allocs/op
BenchmarkQuickSort1M-8        	      20	  96176863 ns/op	       0 B/op	       0 allocs/op
BenchmarkBuiltin20M-8         	       1	5614214413 ns/op	      32 B/op	       1 allocs/op
BenchmarkMergeSort20M-8       	       1	4378292069 ns/op	4055167056 B/op	20000000 allocs/op
BenchmarkQuickSort20M-8       	       1	2038609988 ns/op	       0 B/op	       0 allocs/op
BenchmarkBuiltin100M-8        	       1	30421022288 ns/op	      32 B/op	       1 allocs/op
BenchmarkQuickSort100M-8      	       1	10416693446 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/javib51/go-first-steps/algorithms	195.790s

```
