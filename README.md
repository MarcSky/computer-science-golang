# computer-science-golang
My personal repo with different computer-science implementation of algorithms and data structures

Comparing search algorithm
```
Search in array with 1.000.000 items

BenchmarkSimpleSearch O(log n)
BenchmarkSimpleSearch-4   	    2186	    524613 ns/op

BenchmarkBinarySearch O(n)
BenchmarkBinarySearch-4   	28309273	        40.6 ns/op
```

Comparing sort algorithms
```
BenchmarkMergeSort 1m items O(log n)
BenchmarkMergeSort-4   	      12	  96197076 ns/op

BenchmarkSelectionSort 100k items O(n^2)
BenchmarkSelectionSort-4   	 1	8143228671 ns/op

```