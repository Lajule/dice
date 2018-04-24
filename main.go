package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func init() {
	flag.Usage = func() {
		fmt.Printf("Usage: %s [flags] file\n", os.Args[0])
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()

	rand.Seed(time.Now().UTC().UnixNano())

	r, _ := regexp.Compile("^([0-9]+)d([0-9]+)([+-][0-9]+)?$")

	scanner := bufio.NewScanner(strings.NewReader("3d32\n3d32+3\n3d\n5d6+3\n10d3-3"))

	for scanner.Scan() {
		text := scanner.Text()

		fmt.Printf("%s\n", text)

		if r.MatchString(text) {
			groups := r.FindStringSubmatch(text)

			n, _ := strconv.Atoi(groups[1])
			d, _ := strconv.Atoi(groups[2])

			var draws []int

			for i := 0; i < n; i++ {
				r := rand.Intn(d)
				draws = append(draws, r)
			}

			fmt.Printf("%v\n", draws)

			sum := 0

			for _, draw := range draws {
				sum += draw
			}

			fmt.Printf("%d\n", sum)

			a, _ := strconv.Atoi(groups[3])
			fmt.Printf("%d\n", sum+a)

		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
