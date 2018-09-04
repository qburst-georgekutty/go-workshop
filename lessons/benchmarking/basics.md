# Benchmarking

&nbsp;

➩  Benchmarking is also a very powerful tool provided by golang, Aspect of your code can be setup to be benchmarked for performance reviews.

* Example for benchmarking

    ```golang
    func Fib(n int) int {
            if n < 2 {
                    return n
            }
            return Fib(n-1) + Fib(n-2)
    }
    ```

    ```golang
    // from fib_test.go
    func BenchmarkFib10(b *testing.B) {
            // run the Fib function b.N times
            for n := 0; n < b.N; n++ {
                    Fib(10)
            }
    }
    ```

  * Running benchmark test

    > go test -run none -bench .

    ```golang
        PASS
        BenchmarkFib10   5000000               509 ns/op
        ok      github.com/qburst-georgekutty/fib       3.084s
    ```