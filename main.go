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

func open() (input *os.File) {
	if flag.NArg() > 0 {
		file, err := os.Open(flag.Arg(0))

		if err != nil {
			log.Fatal(err)
		}

		input = file
	} else {
		input = os.Stdin
	}

	return
}

func roll(n, d, a int) (draws []int, result int) {
	for i := 0; i < n; i++ {
		r := rand.Intn(d)
		draws = append(draws, r)
	}

	sum := 0

	for _, draw := range draws {
		sum += draw
	}

	result = sum + a

	return
}

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

	scanner := bufio.NewScanner(open())

	for scanner.Scan() {
		text := scanner.Text()

		if r.MatchString(text) {
			matches := r.FindStringSubmatch(text)
			n, _ := strconv.Atoi(matches[1])
			d, _ := strconv.Atoi(matches[2])
			a, _ := strconv.Atoi(matches[3])

			draws, result := roll(n, d, a)

			add := strings.Join(strings.Split(strings.Trim(fmt.Sprint(draws), "[]"), " "), "+")
			fmt.Printf("%s:(%s)%s=%d\n", matches[0], add, matches[3], result)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
