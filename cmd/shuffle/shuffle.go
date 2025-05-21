package main

import (
	"fmt"
	"math/rand"
)

const (
	CHARSET = "bcdefghjkmnpqrstwxyzBCDEFGHJKMNPQRSTWXYZ:_+!|~"
	CODE    = "ae3ay4ap2agqac2gkhgmghg5hmh2glb2gfa2ga"

	FACTOR = 1
	CYCLES = 500
)

func main() {

	fmt.Println("")

	fmt.Println("###############")
	fmt.Println("## SHUFFLING ##")
	fmt.Println("###############")

	fmt.Println("")

	for cycle := 1; cycle <= CYCLES; cycle++ {

		chars := []byte(CHARSET)

		seed := int64(cycle * FACTOR)
		r := rand.New(rand.NewSource(seed))

		r.Shuffle(len(chars), func(i, j int) {
			chars[i], chars[j] = chars[j], chars[i]
		})

		result := rewrite(chars)
		fmt.Printf("%3d [%05d] %s\n", cycle, seed, result)

		if cycle%10 == 0 {
			fmt.Println("")
		}

	}

	fmt.Println("")

}

func rewrite(mapping []byte) []byte {

	seq := []byte(CODE)

	for i, c := range seq {

		if c < 'a' || c > 'z' {
			continue
		}

		p := int(c - 'a')
		seq[i] = mapping[p]

	}

	return seq

}
