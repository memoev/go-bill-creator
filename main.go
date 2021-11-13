package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Easy way to create bills!")
	arrItems := make([]string, 0)
	arrPrices := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Enter your item: ")
		// Scans a line from Stdin(Console)
		scanner.Scan()
		// Holds the string that scanned
		newItem := scanner.Text()
		if len(newItem) != 0 {
			arrItems = append(arrItems, newItem)
			fmt.Println("Enter price: ")
			scanner.Scan()
			newPrice := scanner.Text()
			if len(newPrice) != 0 {
				arrPrices = append(arrPrices, newPrice)
			} else {
				arrPrices = append(arrPrices, "0")
			}
		} else {
			break
		}
	}

	fmt.Println("The total for items:", arrItems)
	fmt.Println("Is:", calculateTotal(arrPrices))
}

func calculateTotal(p []string) float64 {
	var total float64 = 0
	for _, value := range p {
		i, _ := strconv.ParseFloat(value, 64)
		total = total + i
	}
	return total
}
