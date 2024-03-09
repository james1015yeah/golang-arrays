package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Stats struct {
	mean   float64
	median float64
	mode   int
}

func GetMode(array []int) int {
	mode := 0
	occurrences := 0
	currentValue := array[0]
	currentOccurrences := 0
	n := len(array)
	for i := 0; i < n; i++ {
		if array[i] != currentValue {
			if currentOccurrences > occurrences {
				mode = currentValue
				occurrences = currentOccurrences
			}
			currentOccurrences = 1
			currentValue = array[i]
		} else {
			currentOccurrences++
		}
	}
	return mode
}

func GetStats(array []int) *Stats {
	sum := 0.0
	n := len(array)
	sort.Ints(array)
	for i := 0; i < n; i++ {
		sum += float64(array[i])
	}
	mean := sum / float64(n)
	median := float64(array[n/2]+array[n/2-1]) / 2.0
	if n%2 == 1 {
		median = float64(array[n/2])
	}

	mode := GetMode(array)

	stats := new(Stats)
	stats.mean = mean
	stats.median = median
	stats.mode = mode

	return stats
}

func main() {
	var size int
	fmt.Printf("Enter size of your array: ")
	fmt.Scanln(&size)

	var arr = make([]int, size)
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < size; i++ {
		fmt.Printf("Enter a number %d: ", i)
		scanner.Scan()
		input := scanner.Text()
		num, err := strconv.Atoi(strings.TrimSpace(input))
		if err != nil {
			fmt.Println("Invalid input!")
			i--
			continue
		}
		arr[i] = num
	}
	fmt.Println("Your array is:", arr)

	//highest to lowest
	fmt.Println(" ")
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})
	fmt.Println("Sort 1:", arr)

	//lowest to highest
	sort.Ints(arr)

	fmt.Println("Sort 2:", arr)

	fmt.Println(" ")

	//even numbers
	fmt.Print("Even Numbers: ")

	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println(" ")
	//odd numbers
	fmt.Print("Odd Numbers: ")
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 != 0 {
			fmt.Print(arr[i], " ")
		}
	}

	answer := GetStats(arr)
	fmt.Println(" ")
	fmt.Println(" ")

	//mean, median, mode
	fmt.Printf("Mean: \t%.1f\n", answer.mean)
	fmt.Printf("Median: %.1f\n", answer.median)
	fmt.Printf("Mode: \t%d\n", answer.mode)
}
