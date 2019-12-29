package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "sort"
)

type ShouldMoveCallback func(int32) bool

type ValueToHitMap struct {
	m map[int32]int32
}

type Memory struct {
	vtm *ValueToHitMap
}

type Cache struct {
	vtm *ValueToHitMap
}

func (vtm *ValueToHitMap) Init() { vtm.m = map[int32]int32{} }

func (source *ValueToHitMap) MoveValueTo(destination *ValueToHitMap, value int32) {

	fmt.Printf("Moving %v from source=%v destination=%v\n", value, source, destination)

	hitTimes, ok := source.m[value]

	if ok {
		delete(source.m, value)
		destination.m[value] = hitTimes
	} else {
		fmt.Printf("Missing value %v from map %v\n", value, source)
	}
}

func (source *valueToHitMap) UpdateValues(valueCountMap map[int32]int32, destination *ValueToHitMap, shouldMoveCallback ShouldMoveCallback) {
    for key, value := range source.m {
        countedValue, ok := valueCountMap[key]

        if ok {
            source.m[key] = value + countedValue
        } else {
            if value > 0 {
                source.m[key] = value - 1
            }
        }

        if (shouldMoveCallback(source.m[key])) {
            source.MoveValueTo(destination, key)
            delete(valueCountMap, key)
        }
    }

    fmt.Printf("After UpdateValues map should be: %v\n", *source.m)
}

func (memory *Memory) AddFully(m map[int32]int32) {
    for k, v := range m {
        if v > 2 {
            v = 2
        }

        memory.vtm.m[k] = v
    }
}

func (m *Memory) GetReadThreshold(value int32) bool { return value > 5 }

func (c *Cache) GetTimerTickTreshold(value int32) bool { return value <= 3 }

/*
 * Complete the 'cacheContents' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts 2D_INTEGER_ARRAY callLogs as parameter.
 */
func cacheContents(callLogs [][]int32) []int32 {
    sort.Slice(callLogs, func(i, j int) bool {
		return callLogs[i][0] < callLogs[j][0]	
	})

    memory := &Memory{ vtm: &ValueToHitMap{ } }
    cache  := &Cache{ vtm: &ValueToHitMap{ } }

    memory.vtm.Init()
    cache.vtm.Init()

	for index := 0; index < len(callLogs); {
		currentTimestamp := callLogs[index][0]
        valueCountMap    := map[int32]int32 { }

		for (index < len(callLogs)) && (currentTimestamp == callLogs[index][0]) {
			readValue := callLogs[index][1]
            readValueCount, ok := valueCountMap[readValue]

            if (!ok) {
                readValueCount = 0
            }

            valueCountMap[readValue] = readValueCount + 1
			index++
		}

        fmt.Printf("Updating with values: %v\n", valueCountMap)
        memory.vtm.UpdateValues(valueCountMap, cache.vtm, memory.GetReadThreshold)

        fmt.Printf("Updating with values  #2: %v\n", valueCountMap)
        cache.vtm.UpdateValues(valueCountMap, memory.vtm, cache.GetTimerTickTreshold)

        if len(valueCountMap) > 0 {
            memory.AddFully(valueCountMap)
        }

        fmt.Printf("M=%v\n", memory.vtm)
        fmt.Printf("C=%v\n\n\n", cache.vtm)
	}

    fmt.Printf("M=%v\n", memory)
    fmt.Printf("C=%v\n", cache)

	return []int32{1,2,3,5}
}

func main() {
    logs := [][]int32{
        {1, 1},
        {2, 1},
        {3, 1},
        {4, 2},
        {5, 2},
        {6, 2},
    }

    cacheContents(logs)

    return

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

    callLogsRows, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)

    callLogsColumns, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)

    var callLogs [][]int32
    for i := 0; i < int(callLogsRows); i++ {
        callLogsRowTemp := strings.Split(strings.TrimRight(readLine(reader)," \t\r\n"), " ")

        var callLogsRow []int32
        for _, callLogsRowItem := range callLogsRowTemp {
            callLogsItemTemp, err := strconv.ParseInt(callLogsRowItem, 10, 64)
            checkError(err)
            callLogsItem := int32(callLogsItemTemp)
            callLogsRow = append(callLogsRow, callLogsItem)
        }

        if len(callLogsRow) != int(callLogsColumns) {
            panic("Bad input")
        }

        callLogs = append(callLogs, callLogsRow)
    }

    result := cacheContents(callLogs)

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
