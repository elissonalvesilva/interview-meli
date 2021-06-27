package repository

type AnalyzeGRPC interface {
	AnalyzeDNA([]string) (string, error)
}
