package main

import (
    "fmt"
    "time"
    "math"
)

func assigMulti() (bool, int, string) {
    return true, 3, "Ana are mere"
}

func swap(x, y string) (string, string) {
    return y, x
}

func multiply (x int, y int) int {
    if x == 0 {
        return 0
    }

    if x == 1 {
        return y
    }

    return y + multiply(x - 1, y)
}

func add(x int, y int) int {
    return x + y
}

func printGoDataTypes() {
    var (
        boolean bool        = false
        maxInt  uint64      = (1 << 64) - 1
        randomString string = "This is some random string"
    )

    fmt.Println("========== Printing Go data types ==========")
    fmt.Printf("TYPE: %T Value: %v\n", boolean, boolean)
    fmt.Printf("TYPE: %T Value: %#X\n", maxInt, maxInt)
    fmt.Printf("TYPE: %T Value: '%v'\n", randomString, randomString)
    fmt.Println("============================================")
}

func printMoreDataTypes() {
    i := 42
    f := float64(i)
    u := uint(f)

    fmt.Println(i, f, u)
}

func pow(x, n, lim float64) float64 {
    if v := math.Pow(x, n) ; v < lim {
        return v
    } else {
        fmt.Printf("Invalid value computed: %g > %g\n", v, lim)
    }

    return lim
}

func printDayOfWeek(weekDay int) {
    switch weekDay {
        case 1, 2, 3, 4, 5: {
            fmt.Println("week day!")
        }

        case 6, 7: {
            fmt.Println("weekend!")
        }

        default : {
            fmt.Println("Invalid day!")
        }
    }
}

func printCurrentTime() {
    fmt.Println("The time is", time.Now())
}

func sumUpTo(N int) {

    sum := 1
    for sum < N {
        sum += sum
    }

    fmt.Println(sum)
}

func printFooBar(upToNumber int) {
    var printNumberToConsole bool

    for i := 1; i <= upToNumber; i++ {
        printNumberToConsole = true

        if (i % 3 == 0) {
            printNumberToConsole = false
            fmt.Printf("Foo")
        }
        if (i % 5 == 0) {
            printNumberToConsole = false
            fmt.Printf("Bar")
        }
        if (printNumberToConsole) {
            fmt.Printf("%d", i)
        }

        fmt.Printf("\n")
    }
}

func main() {
    fmt.Println("Welcome to the playground!")

    string1, string2 := swap("hello", "world")
    fmt.Println(string1, string2)

    printCurrentTime()

    var a1 int = 4
    var a2 int = 8
    fmt.Printf("Adding: %d + %d = %d\n", a1, a2, add(a1, a2))

    var m1 int = 4
    var m2 int = 8
    fmt.Println("Multiplying: %d + %d = %d\n", m1, m2, multiply(m1, m2))

    var boolean, integer, string_val = assigMulti()
    fmt.Println(boolean, integer, string_val)

    printGoDataTypes()

    printMoreDataTypes()

    sumUpTo(100)

    printDayOfWeek(34)

    printFooBar(45)
}
