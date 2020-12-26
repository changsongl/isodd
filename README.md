# isodd

### Introduction:
isodd is for checking whether a integer, float or string is odd.
It is similar with the is-odd package which is written in javascript.
The package is 100% test coverage.

### How to use it:

````shell script
# download package with go mod
go get github.com/changsongl/isodd

# import package in your program
import "github.com/changsongl/isodd"
````

````go
// some examples

var strNum1 string = "1"
isOdd, err := isodd.String(strNum1) // true, nil
var strNum2 string = "0"
isOdd, err := isodd.String(strNum2) // false, nil

var num int64 = -1
isOdd := isodd.Int64(num) // true
var num int64 = 0
isOdd := isodd.Int64(num) // false

// float type: it convert float to int type and 
// then check whether it is odd
var num float32 = -1.32
isOdd := isodd.Float32(num) // true 
var num float32 = 0
isOdd := isodd.Int64(num) // false

var num interface = "0"
isOdd := isodd.Interface(num) // false
````

### Code Coverage:
````shell script
go test -c -covermode=count -coverpkg ./...

./isodd.test 

PASS
coverage: 100.0% of statements in ./...
````
