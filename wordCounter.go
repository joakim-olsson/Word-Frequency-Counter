/**
 * @Author: Joakim Olsson <lomo133>
 * @Date:   2019-04-02T18:03:55+02:00
 * @Last modified by:   lomo133
 * @Last modified time: 2019-04-04T19:50:45+02:00
 */

package main

import (
	"fmt"
	"io/ioutil"
	"runtime"
	"strings"
	"time"
)

const DataFile = "sample.txt"

/* Return the word frequencies of the text argument.
Split load optimally across processor cores.*/
func WordCount(text string) map[string]int {

	text1 := strings.Replace(text, ".", "", -1)
	text2 := strings.Replace(text1, ",", "", -1)
	text3 := strings.ToLower(text2)
	words := strings.Fields(text3)
	numCPU := runtime.NumCPU()

	var divided [][]string
	chunkSize := (len(words) + numCPU - 1) / numCPU

	// Dividing the array into the same amount of sub-arrays as CPU's
	for i := 0; i < len(words); i += chunkSize {
		end := i + chunkSize
		if end > len(words) {
			end = len(words)
		}
		divided = append(divided, words[i:end])
	}

	// Splits up the work into as many goroutines as CPU's
	count := len(divided)
	results := make(chan map[string]int, count)
	for _, array := range divided {
		go func(array []string) {
			results <- Frequency(array)
		}(array)
	}

	// Concatenates all maps together into one map.
	freq := make(map[string]int)
	for range divided {
		for a, frequency := range <-results {
			freq[a] += frequency
		}
	}
	return freq
}

// Helper function which counts the words for each sub-array.
func Frequency(s []string) map[string]int {
	freq := make(map[string]int)
	for _, word := range s {
		freq[word]++
	}
	return freq
}

/* Benchmark how long it takes to count word frequencies in text numRuns times.

 Return the total time elapsed. */
func benchmark(text string, numRuns int) int64 {
	start := time.Now()
	for i := 0; i < numRuns; i++ {
		WordCount(text)
	}
	runtimeMillis := time.Since(start).Nanoseconds() / 1e6

	return runtimeMillis
}

// Print the results of a benchmark
func printResults(runtimeMillis int64, numRuns int) {
	fmt.Printf("amount of runs: %d\n", numRuns)
	fmt.Printf("total time: %d ms\n", runtimeMillis)
	average := float64(runtimeMillis) / float64(numRuns)
	fmt.Printf("average time/run: %.2f ms\n", average)
}

func main() {
	data, err := ioutil.ReadFile(DataFile)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Printf("%#v", WordCount(string(data)))

	numRuns := 100
	runtimeMillis := benchmark(string(data), numRuns)
	printResults(runtimeMillis, numRuns)
}
