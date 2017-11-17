# hashmap - implement HashMap data structure, with collision resolution by separate chaining

Benchmarks of HashMap
---
```
$ go test -benchmem -benchtime=200ms -bench=BenchmarkHashmap_
goos: linux
goarch: amd64
pkg: hashmap
BenchmarkHashmap_Set_16-4         	  300000	      1106 ns/op	      39 B/op	       3 allocs/op
BenchmarkHashmap_Set_64-4         	 1000000	       447 ns/op	      39 B/op	       3 allocs/op
BenchmarkHashmap_Set_128-4        	 1000000	       344 ns/op	      39 B/op	       3 allocs/op
BenchmarkHashmap_Set_1024-4       	 1000000	       228 ns/op	      39 B/op	       3 allocs/op
BenchmarkHashmap_Set_Native-4     	 2000000	       193 ns/op	      31 B/op	       2 allocs/op

BenchmarkHashmap_Get_16-4         	  500000	       559 ns/op	      31 B/op	       2 allocs/op
BenchmarkHashmap_Get_64-4         	 1000000	       287 ns/op	      31 B/op	       2 allocs/op
BenchmarkHashmap_Get_128-4        	 1000000	       232 ns/op	      31 B/op	       2 allocs/op
BenchmarkHashmap_Get_1024-4       	 2000000	       186 ns/op	      31 B/op	       2 allocs/op
BenchmarkHashmap_Get_Native-4     	 2000000	       143 ns/op	      19 B/op	       1 allocs/op

BenchmarkHashmap_Unset_16-4       	  300000	       899 ns/op	      32 B/op	       2 allocs/op
BenchmarkHashmap_Unset_64-4       	  500000	       492 ns/op	      32 B/op	       2 allocs/op
BenchmarkHashmap_Unset_128-4      	 1000000	       410 ns/op	      32 B/op	       2 allocs/op
BenchmarkHashmap_Unset_1024-4     	 1000000	       351 ns/op	      32 B/op	       2 allocs/op
BenchmarkHashmap_Unset_Native-4   	 1000000	       436 ns/op	      31 B/op	       1 allocs/op
PASS
ok  	hashmap	25.039s

```

Benchmarks of Hash function
---
Current hash function based on hash/crc32 function. And convert any Key type into []byte by reflect.

```
$ go test -benchmem -benchtime=1000ms -bench=BenchmarkHash_
goos: linux
goarch: amd64
pkg: hashmap
BenchmarkHash_16-4     	 3000000	       426 ns/op	     208 B/op	       4 allocs/op
--- BENCH: BenchmarkHash_16-4
	hash_benchmark_test.go:39: Len uniqs values: 1 (1)
	hash_benchmark_test.go:39: Len uniqs values: 16 (100)
	hash_benchmark_test.go:39: Len uniqs values: 16 (10000)
	hash_benchmark_test.go:39: Len uniqs values: 16 (1000000)
	hash_benchmark_test.go:39: Len uniqs values: 16 (3000000)
BenchmarkHash_64-4     	 3000000	       426 ns/op	     208 B/op	       4 allocs/op
--- BENCH: BenchmarkHash_64-4
	hash_benchmark_test.go:39: Len uniqs values: 1 (1)
	hash_benchmark_test.go:39: Len uniqs values: 42 (100)
	hash_benchmark_test.go:39: Len uniqs values: 64 (10000)
	hash_benchmark_test.go:39: Len uniqs values: 64 (1000000)
	hash_benchmark_test.go:39: Len uniqs values: 64 (3000000)
BenchmarkHash_128-4    	 3000000	       434 ns/op	     208 B/op	       4 allocs/op
--- BENCH: BenchmarkHash_128-4
	hash_benchmark_test.go:39: Len uniqs values: 1 (1)
	hash_benchmark_test.go:39: Len uniqs values: 42 (100)
	hash_benchmark_test.go:39: Len uniqs values: 128 (10000)
	hash_benchmark_test.go:39: Len uniqs values: 128 (1000000)
	hash_benchmark_test.go:39: Len uniqs values: 128 (3000000)
BenchmarkHash_1024-4   	 3000000	       433 ns/op	     208 B/op	       4 allocs/op
--- BENCH: BenchmarkHash_1024-4
	hash_benchmark_test.go:39: Len uniqs values: 1 (1)
	hash_benchmark_test.go:39: Len uniqs values: 100 (100)
	hash_benchmark_test.go:39: Len uniqs values: 1024 (10000)
	hash_benchmark_test.go:39: Len uniqs values: 1024 (1000000)
	hash_benchmark_test.go:39: Len uniqs values: 1024 (3000000)
PASS
ok  	hashmap	6.898s

```
