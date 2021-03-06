package day9

import (
	"bufio"
	"log"
	"os"
	"path"
	"strconv"
)

func readInput(filename string) []int {
	file, err := os.Open(path.Join("testdata", filename))

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	res := []int{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		res = append(res, val)
	}

	return res
}

func part1(size int, in []int) int {
	// map to handle repeated values in window
	window := make(map[int]int)

	// initialize preamble
	for i := 0; i < size; i++ {
		window[in[i]]++
	}

	for i := size; i < len(in); i++ {
		prev := in[i-size]
		curr := in[i]
		// check if elements in window sum to curr
		flag := false

		for k := range window {
			if window[curr-k] != 0 {
				flag = true // need at least one pair to continue
				break
			}
		}
		if !flag {
			return curr
		}

		window[curr]++

		if window[prev] <= 1 {
			delete(window, prev)
		} else {
			window[prev]--
		}
	}

	return 0
}

func minMax(in []int) (int, int) {
	var min, max int

	for i, v := range in {
		switch {
		case v < min || i == 0:
			min = v
		case v > max:
			max = v
		}
	}
	return min, max
}

func part2(target int, in []int) int {
	l := 0
	r := 0
	res := 0

	for l < len(in) && r < len(in) {
		switch {
		case res > target:
			res -= in[l]
			l++
		case res < target:
			res += in[r]
			r++
		case res == target:
			min, max := minMax(in[l:r])
			return min + max
		}
	}

	return 0
}

func run(size int, filename string) (int, int) {
	in := readInput(filename)
	p1 := part1(size, in)
	p2 := part2(p1, in)
	return p1, p2
}
