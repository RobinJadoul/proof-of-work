package main

import (
	"crypto/sha256"
	"fmt"
	"os"
	"strconv"
	"encoding/binary"
	"math/bits"
)

// Solve returns an integer, for which the decimal representation as a string is the solution to the PoW:
//  binstr(sha256(prefix + str(i))).startswith("0"*difficulty)
// @arg prefix The prefix for hashing
// @arg difficulty The number of leading 0 bits wanted
// @ret The solution to the PoW.
func Solve(prefix string, difficulty int) int {
	var i int
	for {
		sum := sha256.Sum256([]byte(fmt.Sprintf("%s%v", prefix, i)))
		if bits.LeadingZeros64(binary.BigEndian.Uint64(sum[:8])) >= difficulty {
			break
		}
		i++
	}
	return i
}

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage: %s <prefix> <difficulty>\n", os.Args[0])
		return
	}
	dif, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("Difficulty is not a valid integer: %s\n", os.Args[2])
		return
	}
	i := Solve(os.Args[1], dif)
	fmt.Println(i)
}
