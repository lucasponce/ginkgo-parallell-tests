# ginkgo-parallell-tests

Exercice for refactoring a long sequential test into indepent tests that ginkgo can run in parallel.

```
# Run the sequential approach (parallel flag won't work as the tests are coded in an ordered way
ginkgo -p test test_sequential/
```

```
[...]
Ran 24 of 24 Specs in 250.200 seconds
SUCCESS! -- 24 Passed | 0 Failed | 0 Pending | 0 Skipped
PASS

Ginkgo ran 1 suite in 4m10.810731534s
Test Suite Passed
```

```
# Run the parallel approach (different tests could be run in parallel)
ginkgo -p test test_parallel/
```
```
[...]
Ran 24 of 24 Specs in 30.008 seconds
SUCCESS! -- 24 Passed | 0 Failed | 0 Pending | 0 Skipped


Ginkgo ran 1 suite in 30.353073841s
Test Suite Passed
```
