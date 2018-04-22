package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"regexp"
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
	fmt.Printf("%d\n", rand.Intn(6))

	r, _ := regexp.Compile("^([0-9]+)d([0-9]+)(([+-])([0-9])+)?$")
	fmt.Println(r.FindStringSubmatch("3d3+2"))
}
