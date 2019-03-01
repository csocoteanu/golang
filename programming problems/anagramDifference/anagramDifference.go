package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

func ToRuneCountMap(string string) map[rune]int32 {
    m := make(map[rune]int32)

    for _, character := range string {
        counter, ok := m[character]

        if ok {
            m[character] = counter + 1
        } else {
            m[character] = 1
        }
    }

    return m
}


func ComputeDifferences(m map[rune]int32, string string) int32 {
    result := int32(len(string))

    for _, character := range string {
        counter, ok := m[character]

        if ok {
            result--

            if counter == 1 {
                    delete(m, character)
                } else {
                    m[character] = counter - 1
                }
        }
    }

    return result
}

/*
 * Complete the 'getMinimumDifference' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. STRING_ARRAY a
 *  2. STRING_ARRAY b
 */
func getMinimumDifference(a []string, b []string) []int32 {
    if len(a) != len(b) {
        panic("Array lengths differ")
    }

    result := make([]int32, len(a))

    for i := 0; i < len(a); i++ {
        if len(a[i]) != len(b[i]) {
            result[i] = -1
        } else {
            result[i] = ComputeDifferences(ToRuneCountMap(a[i]), b[i])
        }
    }

    return result
}

func main() {
    os.Setenv("INPUT_PATH", "./in.txt")
    os.Setenv("OUTPUT_PATH", "./out.txt")

    file, err := os.Open(os.Getenv("INPUT_PATH"))
    checkError(err)

    reader := bufio.NewReaderSize(file, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer file.Close()
    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    aCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)

    var a []string

    for i := 0; i < int(aCount); i++ {
        aItem := readLine(reader)
        a = append(a, aItem)
    }

    bCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)

    var b []string

    for i := 0; i < int(bCount); i++ {
        bItem := readLine(reader)
        b = append(b, bItem)
    }

    result := getMinimumDifference(a, b)

    for i, resultItem := range result {
        fmt.Fprintf(writer, "%d", resultItem)

        if i != len(result) - 1 {
            fmt.Fprintf(writer, "\n")
        }
    }

    fmt.Fprintf(writer, "\n")

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
