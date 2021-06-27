package mock

var DNAValid = []string{
"CTGAGA",
"CTGAGC",
"TATTGT",
"AGAGGG",
"CCCCTG",
"TCACTG",
}

var DNAMatrixValid = [][]string{
	{"C", "T", "G", "A", "G", "A"},
	{"C", "T", "G", "A" ,"G","C"},
	{"T", "A", "T", "T" ,"G", "T"},
	{"A", "G", "A", "G", "G", "G"},
	{"C", "C", "C", "C", "T", "G"},
	{"T", "C", "A", "C", "T", "G"},
}

var DNAHuman = []string{
"ATGCGA",
"CAGTGC",
"TTATTT",
"AGACGG",
"GCGTCA",
"TCACTG",
}

var DNA3 = []string{
"CTGAGA",
"CTAACC",
"TATTGT",
"AGAGGG",
"CCACTA",
"TCACTG",
}

var DNA4 = []string{
"CTAGG",
"CCGAA",
"TGACG",
"GACCA",
"GACCA",
}

var DNAInvalidMatrix = []string{
"CTAG",
"CCGA",
"TGAC",
"GACC",
"GACC",
}

var DNAInvalidChar = []string{
	"CTGAGX",
	"CTGAGC",
	"TATTGT",
	"AGAGGG",
	"CCCCTG",
	"TCACTG",
}

var DNA6 = []string{
"AACTCA",
"CGGCGC",
"ACTTCG",
"CTCAAT",
"TGACCA",
"CCAACC",
}

var DNA7 = []string{
"CCGCCTGC",
"ACGTGCTG",
"AGATGGGG",
"CAATCGCC",
"ACTGACCT",
"CTTGCCAT",
"AATAACCT",
"TGTCGCGA",
}
