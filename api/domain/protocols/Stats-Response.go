package protocols

type StatsResponse struct {
	CountSimianDNA int `json:"count_simian_dna"`
	CountHumanDNA int `json:"count_human_dna"`
	Ratio float64 `json:"ratio"`
}
