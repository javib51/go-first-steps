#Benchmark
PASS
BenchmarkBuiltin1k-8    	    5000	    315269 ns/op	       0 B/op	       0 allocs/op
BenchmarkMergeSort1k-8  	    5000	    258573 ns/op	  186758 B/op	    1000 allocs/op
BenchmarkQuickSort1k-8  	   10000	    113475 ns/op	       0 B/op	       0 allocs/op
BenchmarkBuiltin10k-8   	     500	   3286986 ns/op	       0 B/op	       0 allocs/op
BenchmarkMergeSort10k-8 	     500	   3221263 ns/op	 1867581 B/op	   10000 allocs/op
BenchmarkQuickSort10k-8 	    2000	   1242389 ns/op	       0 B/op	       0 allocs/op
BenchmarkBuiltin1M-8    	       5	 315598961 ns/op	       6 B/op	       0 allocs/op
BenchmarkMergeSort1M-8  	       5	 262908216 ns/op	186758118 B/op	 1000037 allocs/op
BenchmarkQuickSort1M-8  	      10	 117575999 ns/op	       0 B/op	       0 allocs/op
BenchmarkBuiltin20M-8   	       1	6574788408 ns/op	      32 B/op	       1 allocs/op
BenchmarkMergeSort20M-8 	       1	5059337270 ns/op	4055169200 B/op	20000207 allocs/op
BenchmarkQuickSort20M-8 	       1	2277166759 ns/op	       0 B/op	       0 allocs/op
BenchmarkBuiltin100M-8  	       1	33674428774 ns/op	      32 B/op	       1 allocs/op
BenchmarkQuickSort100M-8	       1	12385765292 ns/op	       0 B/op	       0 allocs/op

ok  	github.com/javib51/go-first-steps/algorithms	94.053s
