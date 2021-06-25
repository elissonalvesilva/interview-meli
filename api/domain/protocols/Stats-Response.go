package protocols

type StatsResponse struct {
	CountSimianDNA int64 `json:"count_simian_dna"`
	CountHumanDNA int64 `json:"count_human_dna"`
	Ratio float64 `json:"ratio"`
}
