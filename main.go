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

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		text := scanner.Text()

		if r.MatchString(text) {
			groups := r.FindStringSubmatch(text)

			n, _ := strconv.Atoi(groups[1])
			d, _ := strconv.Atoi(groups[2])

			var draws []int

			for i := 0; i < n; i++ {
				r := rand.Intn(d)
				draws = append(draws, r)
			}

			sum := 0

			for _, draw := range draws {
				sum += draw
			}

			additional, _ := strconv.Atoi(groups[3])

			fmt.Printf("%s:(%s)%s=%d\n", text, strings.Join(strings.Split(strings.Trim(fmt.Sprint(draws), "[]"), " "), "+"), groups[3], sum + additional)

		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
