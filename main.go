package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"tyranny.rocket.equation/fuel"
)

func main() {
	var totalFuel float64
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter the mass of each module (enter empty line to end the read process):")
	for {
		scanner.Scan()
		entry := scanner.Text()
		if len(entry) != 0 {
			if f, err := strconv.ParseFloat(entry, 64); err == nil {
				totalFuel += fuel.Required(f)
			} else {
				fmt.Println("wrong mass entered")
				break
			}
		} else {
			break
		}
	}
	fmt.Printf("the total fuel required is %v\n", totalFuel)
}
