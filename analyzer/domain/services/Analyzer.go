package services

import (
	"fmt"
	constants "github.com/elissonalvesilva/interview-meli/analyzer/shared"
	"strings"
)

type Analyzer struct {}

const sequencesNeeded = 1
const consecutiveLettersNeeded = 4
var MatrixLength int
var validLetters = map[string]int{
	"A": 1,
	"C": 2,
	"G": 3,
	"T": 4,
}

func NewAnalyzer() *Analyzer {
	return &Analyzer{}
}

func IsValidMatrix(dna []string) bool {
	return len(dna) == len(dna[0])
}

func (a *Analyzer) DNAAnalyzer(dna []string) (string, error) {
	var matrix, err = GetMatrixFromDNAList(dna)

	count := 0

	if err != nil {
		return constants.EmptyDNA, err
	}

	MatrixLength = len(dna)
	dnaMatrixLength := len(dna)

	fmt.Printf("Analyzing DNA: %s \n", dna)
	for row := 0; row < dnaMatrixLength; row++ {
		for col := 0; col < dnaMatrixLength; col++ {
			currentLetter := matrix[row][col]

			if SearchForSimianValidRepetitions(matrix, row, col, row, col, 0, currentLetter) {
				count += 1
				if count >= sequencesNeeded {
					fmt.Printf("Finished Analyze of DNA: %s. This DNA is: %s \n", dna, constants.DNATypeSimian)
					return constants.DNATypeSimian, nil
				}
			}
		}
	}

	fmt.Printf("Finished Analyze of DNA: %s. This DNA is: %s \n", dna, constants.DNATypeHuman)
	return constants.DNATypeHuman, nil
}

func SearchForSimianValidRepetitions(matrix [][]string, lastRow int, lastCol int, currentRow int, currentCol int, countLettersFound int, nextLetter string) bool {
	if countLettersFound == consecutiveLettersNeeded {
		return true
	}

	if !isValidCoordinate(currentRow, currentCol) || !(matrix[currentRow][currentCol] == nextLetter) {
		return false
	}

	var result = false

	if countLettersFound > 0 {
		sumRow := currentRow - lastRow
		sumCol := currentCol - lastCol
		result = SearchForSimianValidRepetitions(matrix, currentRow, currentCol, currentRow + sumRow, currentCol + sumCol, countLettersFound + 1, nextLetter)
	}else {
		result = SearchForSimianValidRepetitions(matrix, lastRow, lastCol, currentRow -1, currentCol + 1, 1, nextLetter) ||
			SearchForSimianValidRepetitions(matrix, lastRow, lastCol, currentRow, currentCol + 1, 1, nextLetter) ||
			SearchForSimianValidRepetitions(matrix, lastRow, lastCol, currentRow + 1, currentCol +1, 1, nextLetter) ||
			SearchForSimianValidRepetitions(matrix, lastRow, lastCol, currentRow + 1, currentCol, 1, nextLetter)
	}

	return result
}

func GetMatrixFromDNAList(dna []string) ([][]string, error){
	if dna == nil {
		return nil, ErrDNAEmpty
	}

	if !IsValidMatrix(dna) {
		return nil, ErrInvalidMatrix
	}

	matrixLength := len(dna)
	var matrixResult [][]string

	for row := 0; row < matrixLength; row++ {

		rowArray := strings.Split(dna[row], "")

		for col := 0; col < len(rowArray); col++ {
			currentCol := rowArray[col]
			if _, exists := validLetters[currentCol]; !exists {
				return nil, ErrInvalidCharMatrix
			}
		}
		matrixResult = append(matrixResult, rowArray)

	}

	return matrixResult, nil
}

func isValidCoordinate(actualRow int, actualCol int) bool {
	return !(actualRow <0 || actualRow >= MatrixLength || actualCol < 0 || actualCol >= MatrixLength)
}
