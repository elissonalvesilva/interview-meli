package main

import (
	"fmt"
	"strings"
)

const sequencesNeeded = 1
const consecLettersNeeded = 4
var matrixLength int

var validLetters = map[string]int{
	"A": 1,
	"C": 2,
	"G": 3,
	"T": 4,
}


func main()  {
	dna1 := []string{
		"CTGAGA",
		"CTGAGC",
		"TATTGT",
		"AGAGGG",
		"CCCCTG",
		"TCACTG",
	}

	dna2 := []string{
		"ATGCGA",
		"CAGTGC",
		"TTATTT",
		"AGACGG",
		"GCGTCA",
		"TCACTG",
	}

	dna3 := []string{
		"CTGAGA",
		"CTAACC",
		"TATTGT",
		"AGAGGG",
		"CCACTA",
		"TCACTG",
	}

	dna4 := []string{
		"CTAGG",
		"CCGAA",
		"TGACG",
		"GACCA",
		"GACCA",
	}

	dna5 := []string{
		"CTAG",
		"CCGA",
		"TGAC",
		"GACC",
		"GACC",
	}

	dna6 := []string{
		"AACTCA",
		"CGGCGC",
		"ACTTCG",
		"CTCAAT",
		"TGACCA",
		"CCAACC",
	}

	dna7 := []string{
		"CCGCCTGC",
		"ACGTGCTG",
		"AGATGGGG",
		"CAATCGCC",
		"ACTGACCT",
		"CTTGCCAT",
		"AATAACCT",
		"TGTCGCGA",
	}

	is := isSimian(dna1)
	fmt.Println(is)
	isS := isSimian(dna2)
	fmt.Println(isS)

	isSS := isSimian(dna3)
	fmt.Println(isSS)

	a := isSimian(dna4)
	fmt.Println(a)

	b := isSimian(dna5)
	fmt.Println(b)

	c := isSimian(dna6)
	fmt.Println(c)

	d := isSimian(dna7)
	fmt.Println(d)
}

func IsValidMatrix(dna []string) bool {

	return len(dna) == len(dna[0])
}

func isSimian(dna []string) bool {
	var matrix = getMatrixFromDNAList(dna)
	count := 0
	if matrix == nil {
		fmt.Println("Empty string list, please insert a valid matrix(at least 4x4)")
		return false
	}

	matrixLength = len(dna)
	dnaMatrixLength := len(dna)

	for row := 0; row < dnaMatrixLength; row++ {
		for col := 0; col < dnaMatrixLength; col++ {
			currentLetter := matrix[row][col]

			if searchForSimianValidRepetitions(matrix, row, col, row, col, 0, currentLetter) {
				count += 1
				if count >= sequencesNeeded {
					return true
				}
			}
		}
	}

	return false
}

func searchForSimianValidRepetitions(matrix [][]string, lastRow int, lastCol int, currentRow int, currentCol int, countLettersFound int, nextLetter string) bool {
	if countLettersFound == consecLettersNeeded {
		return true
	}

	if !isValidCoord(currentRow, currentCol) || !(matrix[currentRow][currentCol] == nextLetter) {
		return false
	}

	var result = false

	if countLettersFound > 0 {
		sumRow := currentRow - lastRow
		sumCol := currentCol - lastCol

		result = searchForSimianValidRepetitions(matrix, currentRow, currentCol, currentRow + sumRow, currentCol + sumCol, countLettersFound + 1, nextLetter)
	}else {
		result = searchForSimianValidRepetitions(matrix, lastRow, lastCol, currentRow -1, currentCol + 1, 1, nextLetter) || // find in top duagibak
			searchForSimianValidRepetitions(matrix, lastRow, lastCol, currentRow, currentCol + 1, 1, nextLetter) ||
			searchForSimianValidRepetitions(matrix, lastRow, lastCol, currentRow + 1, currentCol +1, 1, nextLetter) || // find in bottom diagonal
			searchForSimianValidRepetitions(matrix, lastRow, lastCol, currentRow + 1, currentCol, 1, nextLetter) // bottom

	}

	return result
}

func getMatrixFromDNAList(dna []string) [][]string{
	if dna == nil {
		return nil
	}

	if !IsValidMatrix(dna) {
		fmt.Println("Invalid Matrix")
		return nil
	}

	matrixLength := len(dna)
	var matrixResult [][]string

	for row := 0; row < matrixLength; row++ {

		rowArray := strings.Split(dna[row], "")

		for col := 0; col < len(rowArray); col++ {
			currentCol := rowArray[col]
			if _, exists := validLetters[currentCol]; !exists {
				fmt.Println("Invalid char in the matrix, only permitted: G, T, A, C.")
				return nil
			}
		}
		matrixResult = append(matrixResult, rowArray)

	}

	return matrixResult
}

func isValidCoord(actualRow int, actualCol int) bool {
	return !(actualRow <0 || actualRow >= matrixLength || actualCol < 0 || actualCol >= matrixLength)
}