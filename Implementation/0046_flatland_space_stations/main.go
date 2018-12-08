package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type City struct {
	numberStation int32
	hasStation    bool
}

type Cities []City

func populateCity(n int32, c []int32) (cities Cities) {
	for i := int32(0); i < n; i++ {
		city := City{
			i,
			false,
		}
		cities = append(cities, city)
	}

	for _, v1 := range c {
		for i := int32(0); i < n; i++ {
			// For update use pointer
			attr := &cities[i]
			if v1 == cities[i].numberStation {
				attr.hasStation = true
				break
			}
		}
	}

	return
}

// Complete the flatlandSpaceStations function below.
func flatlandSpaceStations(n int32, c []int32) (maxDistance int32) {
	var (
		tmpDistance    int32
		tmpMaxDistance int32
		initStation    = false
		cities         = populateCity(n, c)
	)

	//fmt.Println(cities)

	for _, v := range cities {
		if !v.hasStation {
			tmpDistance++
		} else {
			if initStation {
				tmpMaxDistance = (tmpDistance + 1) / 2
				// Set the max distance
				if tmpMaxDistance >= maxDistance {
					maxDistance = tmpMaxDistance
				}
			} else {
				tmpMaxDistance = tmpDistance
			}
			//fmt.Println(tmpMaxDistance)
			tmpDistance = 0
			initStation = true
		}
	}

	// When does not exist station on the last city
	if tmpDistance != 0 {
		tmpMaxDistance = tmpDistance
		if tmpMaxDistance >= maxDistance {
			maxDistance = tmpMaxDistance
		}
	}
	return
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nm := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nm[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	mTemp, err := strconv.ParseInt(nm[1], 10, 64)
	checkError(err)
	m := int32(mTemp)

	cTemp := strings.Split(readLine(reader), " ")

	var c []int32

	for i := 0; i < int(m); i++ {
		cItemTemp, err := strconv.ParseInt(cTemp[i], 10, 64)
		checkError(err)
		cItem := int32(cItemTemp)
		c = append(c, cItem)
	}

	result := flatlandSpaceStations(n, c)

	fmt.Fprintf(writer, "%d\n", result)

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
