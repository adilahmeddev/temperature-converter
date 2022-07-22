package main

import (
	"bufio"
	"excercise4"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	Run(os.Stdin, os.Stdout)
}

func Run(in io.Reader, out io.Writer) {
	fmt.Fprintf(out, "c or f\n")

	scanner := bufio.NewScanner(in)
	scanner.Scan()
	choice := scanner.Text()
	if choice == "c" {
		//fmt.Fprint(out, "temp?")
		scanner.Scan()
		celsiusStr := scanner.Text()
		celsius, err := strconv.ParseFloat(celsiusStr, 64)
		if err != nil {
			panic(err)
		}

		f, err := excercise4.Converter{}.ConvertToF(celsius)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(out, "%.2f", f)

	} else if choice == "f" {
		panic("")
	}
}
