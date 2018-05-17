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

func roll(n, d, m int) int {
	sum := 0
	for i := 0; i < n; i++ {
		sum += rand.Intn(d) + 1
	}
	return sum + m
}

func init() {
	flag.Usage = func() {
		fmt.Printf("Usage: %s file\n", os.Args[0])
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
			m, _ := strconv.Atoi(matches[3])
			fmt.Println(roll(n, d, m))
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
