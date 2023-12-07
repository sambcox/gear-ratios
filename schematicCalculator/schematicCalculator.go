package schematiccalculator

import (
	//"fmt"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func CalculateSchematicSum(lines []string) int {

	var sum int
	for index, line := range lines {
		if index == 0 {
			sum += calculateLine("", line, lines[index+1])
		} else if index == len(lines)-1 {
			sum += calculateLine(lines[index-1], line, "")
		} else {
			sum += calculateLine(lines[index-1], line, lines[index+1])
		}
	}
	return sum
}

func calculateLine(priorLine string, line string, subsequentLine string) int {
	parts := strings.Split(line, "")
	priorLineParts := strings.Split(priorLine, "")
	subsequentLineParts := strings.Split(subsequentLine, "")
	lineSum := 0
	partNumber := ""

	for index, part := range parts {
		if unicode.IsDigit(rune(part[0])) && index == len(parts)-1 {
			partNumber += part
			startIndex := max(0, index-(len(partNumber)+1))
			endIndex := min(len(parts)-1, index+1)
			priorAdjacentParts := []string{}
			if priorLine != "" {
				priorAdjacentParts = priorLineParts[startIndex:endIndex]
			}
			subsequentAdjacentParts := []string{}
			if subsequentLine != "" {
				subsequentAdjacentParts = subsequentLineParts[startIndex:endIndex]
			}
			if isNumericPartValid(parts[startIndex], part, priorAdjacentParts, subsequentAdjacentParts) {
				partNumberInt, err := strconv.Atoi(partNumber)
				if err == nil {
					fmt.Println("Adding", partNumberInt)
					lineSum += partNumberInt
				}
			}
			partNumber = ""
		} else if unicode.IsDigit(rune(part[0])) {
			partNumber += part
		} else if len(partNumber) > 0 {
			startIndex := max(0, index-(len(partNumber)+1))
			endIndex := min(len(parts)-1, index+1)
			priorAdjacentParts := []string{}
			if priorLine != "" {
				priorAdjacentParts = priorLineParts[startIndex:endIndex]
			}
			subsequentAdjacentParts := []string{}
			if subsequentLine != "" {
				subsequentAdjacentParts = subsequentLineParts[startIndex:endIndex]
			}
			if isNumericPartValid(parts[startIndex], part, priorAdjacentParts, subsequentAdjacentParts) {
				partNumberInt, err := strconv.Atoi(partNumber)
				if err == nil {
					fmt.Println("Adding", partNumberInt)
					lineSum += partNumberInt
				}
			}
			partNumber = ""
		}
	}
	return lineSum
}

func isNumericPartValid(beginPart string, endPart string, priorAdjacentParts, subsequentAdjacentParts []string) bool {
	return endPart != "." || beginPart != "." || hasAdjacent(priorAdjacentParts) || hasAdjacent(subsequentAdjacentParts)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func hasAdjacent(parts []string) bool {
	for _, part := range parts {
		if part != "." && !unicode.IsDigit(rune(part[0])) {
			return true
		}
	}
	return false
}
