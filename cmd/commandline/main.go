package main

import (
	"bufio"
	"excercise4"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//fmt.Println("c or f")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	choice := scanner.Text()
	if choice == "c" {
		//fmt.Println("temp?")
		scanner.Scan()
		celsiusStr := scanner.Text()
		celsius, err := strconv.ParseFloat(celsiusStr, 64)
		if err != nil {
			panic(err)
		}

		f := excercise4.Converter{}.ConvertToF(celsius)
		fmt.Printf("%.2f", f)

	} else if choice == "f" {
		panic("")
	}
}
