// Extreme brute force method using GO's concurrency.
// This was made with repl.it on a chromebook because
// my computer got taken away for not having a -B average.
package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

var NUMWORKERS = 5

type workerPool struct {
	fChan   chan string
	workers []int
}

func main() {
	pool := workerPool{
		fChan:   make(chan string),
		workers: make([]int, NUMWORKERS),
	}

	for i := 0; i < NUMWORKERS; i++ {
		pool.workers[i] = i
	}

	for i := 0; i < len(pool.workers); i++ {
		go func(i int, c chan string) {
			var hash [16]byte

			i++
			for {
				hash = md5.Sum([]byte(data + strconv.Itoa(i)))

				if hash[0] == 0 && hash[1] == 0 && hash[2] == 0 {
					c <- data + strconv.Itoa(i)
				}

				i += NUMWORKERS
			}
		}(i, pool.fChan)
	}
	fmt.Printf("%v workers are working on finding the puzzle input...\n", NUMWORKERS)

	puzzleKey := <-pool.fChan

	hash := md5.Sum([]byte(puzzleKey))

	fmt.Printf("input %v gives hash %x\n", puzzleKey, hash)
}

// my personal input
const data = "bgvyzdsv"
