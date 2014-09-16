findsum
=======

Find slice with specified sum

```
$ findsum git:(master) ✗ go test -v --bench=. findsum --benchtime=5s
=== RUN TestHappyPath
--- PASS: TestHappyPath (0.00 seconds)
=== RUN TestHappyPathBruteForce
--- PASS: TestHappyPathBruteForce (0.00 seconds)
=== RUN TestNotFound
--- PASS: TestNotFound (0.00 seconds)
=== RUN TestNotFoundBruteForce
--- PASS: TestNotFoundBruteForce (0.00 seconds)
=== RUN: ExampleFindSum
--- PASS: ExampleFindSum (24.256us)
=== RUN: ExampleFindSumErrNotFound
--- PASS: ExampleFindSumErrNotFound (13.35us)
PASS
BenchmarkFindSum      100000        110501 ns/op
BenchmarkFindSumBruteForce        10    1085130667 ns/op
ok      findsum 24.182s
```
