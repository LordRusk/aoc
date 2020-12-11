// the reason this isn't polished, and is quite bad actually
// is because i was approaching the puzzle the wrong way for
// 45 minutes. So, ya. it works, kinda.
package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("stop the program whenever you find the right number")

	var hash [16]byte

	var i int
	for {
		hash = md5.Sum([]byte(data + strconv.Itoa(i)))

		if hash[0] == 0 && hash[1] == 0 {
			fmt.Printf("input %v gave hash %x\n", data+strconv.Itoa(i), hash)
		}

		i++
	}

}

// my personal input
const data = "bgvyzdsv"
