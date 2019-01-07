package main

import "fmt"

type Vertex struct {
    X int
    Y int
}

var (
v1          = Vertex{1, 2}  // has type Vertex
v2          = Vertex{X: 1}  // Y:0 is implicit
v3          = Vertex{}      // X:0 and Y:0
p           = &Vertex{1, 2} // has type *Vertex
primes      = []int {2, 3, 5, 7, 11, 13}
slicedArray = primes[1 : 4] // array of refs to primes...any changes in slicedArray are reflected in primes
customSlice = []struct {
                    i int
                    b bool
              }{
                    {2, true},
                    {3, false},
                    {5, true},
                    {7, true},
                    {11, false},
                    {13, true},
              }
companiesMap = map[string]Vertex{
                    "Bell Labs": { 40, -74, },
                    "Google"   : { 37, -122, },
               }
)

func printSliceInfo(s []int) {
    fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printPrimeSliceInfo() {
    fmt.Println("======= Printing custom slice info =======>")

    s := primes[ : ]
    printSliceInfo(s)

    // Slice the slice to give it zero length.
    s = s[ : 0]
    printSliceInfo(s)

    // Extend its length.
    s = s[ : 4]
    printSliceInfo(s)

    // Drop its first two values.
    s = s[2 : ]
    printSliceInfo(s)

    printSliceInfo(primes)

    fmt.Println("======= Ending custom slice info =======>")
}

func printPowersOfTwo() {
    var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

    for i, v := range pow {
        fmt.Printf("2**%d = %d\n", i, v)
    }

    fmt.Printf("==> Powers of two: %v\n", pow);
}

func adder() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x
        return sum
    }
}

func fibonacci() func() int {
    fibonacci0 := 0
    fibonacci1 := 1

    return func() int {
        nextFibonacci := fibonacci0 + fibonacci1

        fibonacci0 = fibonacci1
        fibonacci1 = nextFibonacci

        return nextFibonacci
    }
}

func printFibonacciSequence(upToNumber int) {
    fmt.Println("== fibonacci sequence start ==>")

    fibonacciFn := fibonacci()

    for i := 0; i < upToNumber; i++ {
        fmt.Printf("%d => %d\n", (i + 1), fibonacciFn())
    }

    fmt.Println("== fibonacci sequence end ==>")
}

func main() {
    fmt.Println(v1, p, v2, v3, primes)
    fmt.Printf("== sliced array ==> %v\n", slicedArray)
    fmt.Printf("== anounymous struct array ==> %v\n", customSlice)

    printPrimeSliceInfo()

    var pVertex *Vertex = &v3
    fmt.Printf("== This vertex pointer is nil (NULL) ==> %v \n", pVertex)

    printPowersOfTwo()

    fmt.Printf("==== Maps ===> %v\n", companiesMap)

    value, key := companiesMap["mumu"]
    fmt.Printf("==== Inexisting key ===> %v %v\n", value, key)

    pos, neg := adder(), adder()
    for i := 0; i < 10; i++ {
        fmt.Println(
            pos(i),
            neg(-2*i),
        )
    }

    printFibonacciSequence(10)
}
