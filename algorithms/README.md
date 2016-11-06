#Benchmark
```
$ go test -bench=.
BenchmarkBuiltin1k-8          	   10000	    264130 ns/op	       0 B/op	       0 allocs/op
BenchmarkMergeSort1k-8        	    5000	    298151 ns/op	  186757 B/op	    1000 allocs/op
BenchmarkMergeSortFast1k-8    	    1000	   2190089 ns/op	       0 B/op	       0 allocs/op
BenchmarkQuickSort1k-8        	   20000	     96044 ns/op	       0 B/op	       0 allocs/op
BenchmarkBubbleSort1k-8       	     100	 155349183 ns/op	       0 B/op	       0 allocs/op
BenchmarkSelectionSort1k-8    	     100	  93919342 ns/op	       0 B/op	       0 allocs/op
BenchmarkBuiltin10k-8         	     500	   2544432 ns/op	       0 B/op	       0 allocs/op
BenchmarkMergeSort10k-8       	     500	   2994561 ns/op	 1867577 B/op	    9999 allocs/op
BenchmarkMergeSortFast10k-8   	     100	  23676515 ns/op	       0 B/op	       0 allocs/op
BenchmarkQuickSort10k-8       	    2000	    958968 ns/op	       0 B/op	       0 allocs/op
BenchmarkBubbleSort10k-8      	      10	1532499039 ns/op	       0 B/op	       0 allocs/op
BenchmarkSelectionSort10k-8   	      20	1889269714 ns/op	       0 B/op	       0 allocs/op
BenchmarkBuiltin1M-8          	       5	 253678890 ns/op	       6 B/op	       0 allocs/op
BenchmarkMergeSort1M-8        	       5	 274606595 ns/op	186757750 B/op	 1000000 allocs/op
BenchmarkMergeSortFast1M-8    	       1	2294866421 ns/op	       0 B/op	       0 allocs/op
BenchmarkQuickSort1M-8        	      20	  95706654 ns/op	       0 B/op	       0 allocs/op
BenchmarkBuiltin20M-8         	       1	5514871339 ns/op	      32 B/op	       1 allocs/op
BenchmarkMergeSort20M-8       	       1	4751679368 ns/op	4055166976 B/op	19999999 allocs/op
BenchmarkMergeSortFast20M-8   	       1	180716087848 ns/op	       0 B/op	       0 allocs/op
BenchmarkQuickSort20M-8       	       1	1947765884 ns/op	       0 B/op	       0 allocs/op
BenchmarkBuiltin100M-8        	       1	30690513290 ns/op	      32 B/op	       1 allocs/op
BenchmarkQuickSort100M-8      	       1	10350879858 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/javib51/go-first-steps/algorithms	384.054s

```
