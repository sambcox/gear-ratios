package main

import (
	"bufio"
	"fmt"
	"github.com/sambcox/gear-ratios/schematicCalculator"
	"os"
)

func main() {
	file, err := os.Open("schematicinput.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	schematicSum := schematiccalculator.CalculateSchematicSum(lines)

	fmt.Println("Sum of part numbers:", schematicSum)
}
