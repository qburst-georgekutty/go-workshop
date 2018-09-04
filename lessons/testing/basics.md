# Testing

&nbsp;

➩ **Testing** is built right into the go tools and the standard library. Testing needs to be a vital part of the development process because it can save you a tremendous amount of time throughout the life cycle of the project.

&nbsp;

## Testing Techniques

* Basic Unit [Test](https://github.com/qburst-georgekutty/go-workshop/blob/master/codes//testing/example1/example1.go).

* Table [Test](https://github.com/qburst-georgekutty/go-workshop/blob/master/codes//testing/example2/example2.go).

* Mocking HTTP-based requests using [httptests](https://github.com/qburst-georgekutty/go-workshop/blob/master/codes//testing/example3/example3.go)

* Example [Test](https://github.com/qburst-georgekutty/go-workshop/blob/master/codes//testing/example4/example4.go)

* Sub Tests

  * Sub test

    ```golang
        for i, tt := range tests {
            // sub test
            st := func(t *testing.T) {
                // do testing
            }
            t.Run(tt.name, st)
        }
    ```

  * Parallel Test

    ```golang
        for i, tt := range tests {
            // sub test
            st := func(t *testing.T) {

                // all subtest run parallel
                t.Parallel()

                // do testing
            }
            t.Run(tt.name, st)
        }
    ```

## Coverage

* Making sure your tests cover as much of your code as possible is critical.

    ```golang
        go test -coverprofile cover.out
        go tool cover -html=cover.out
    ```