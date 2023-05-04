package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/twpayne/go-polyline"
)

func main() {
	coords := make([][]float64, 0)
	file, err := os.Open("./example.log")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		if index := strings.Index(line, "Location arrived"); index != -1 && index == 59 {
			strs := strings.Split(line[index+len("Location arrived"):], ",")
			if len(strs) == 2 {
				lat, err1 := strconv.ParseFloat(strings.TrimSpace(strs[0]), 64)
				lon, err2 := strconv.ParseFloat(strings.TrimSpace(strs[1]), 64)
				if err1 == nil && err2 == nil {
					coords = append(coords, []float64{lat, lon})
					i++
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	if len(coords) == 0 {
		fmt.Println("No coordinates found in file.")
		return
	}

	encoded := polyline.EncodeCoords(coords)
	fmt.Println(encoded)
}
