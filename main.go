package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var num int // -n

func main() {
	flag.IntVar(&num, "n", 10, `number The first number lines of each input file shall be copied 
    to standard output. The application shall ensure that the number option-argument is a positive decimal integer.

    When a file contains less than number lines, it shall be copied to standard output in its entirety. This shall not be an error.

    If no options are specified, head shall act as if -n 10 had been specified.`)
	flag.Parse()

	fi, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if fi.Mode()&os.ModeNamedPipe == 0 {
		fmt.Fprint(os.Stderr, "No hay datos para recibir")
		os.Exit(1)
	}

	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	arr := strings.Split(string(bytes), "\n")
	print(arr)
	os.Exit(0)
}

func print(data []string) {
	if num > len(data) {
		fmt.Println(strings.Join(data, "\n"))
		return
	}

	fmt.Println(strings.Join(data[:num], "\n"))
}
