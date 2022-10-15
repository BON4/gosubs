# Subscription benchmarks

## Subscibtion List with filters (Create, Select)
```
sqlc:
BenchmarkSubList-4   	     405	   3395468 ns/op
BenchmarkSubList-4   	     432	   2706136 ns/op
BenchmarkSubList-4   	     339	   3445595 ns/op

boil:
BenchmarkSubList-4           324       6197685 ns/op
BenchmarkSubList-4   	     494	   2549310 ns/op
BenchmarkSubList-4   	     438	   3022263 ns/op
```

## Subscription Update (Create, Update)
```
sqlc:
BenchmarkUpdate-4   	     154	   6837069 ns/op
BenchmarkUpdate-4   	     217	   6442926 ns/op
BenchmarkUpdate-4   	     249	   6531757 ns/op

boil:
BenchmarkCreate-4   	     296	   4447802 ns/op
BenchmarkCreate-4   	     266	   4640069 ns/op
BenchmarkCreate-4   	     146	   6955864 ns/op
```
