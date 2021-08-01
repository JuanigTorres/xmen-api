package model

import "github.com/JuanigTorres/xmen-api/math/utils"

type XmenRequest struct {
	DNA []string
}

type StatsRequest struct {
	CountMutantDNA int64   `json:"count_mutant_dna"`
	CountHumanDNA  int64   `json:"count_human_dna"`
	Ratio          float32 `json:"ratio"`
}

func StatsRequestNew(mutants, humans int64) *StatsRequest {
	return &StatsRequest{
		CountMutantDNA: mutants,
		CountHumanDNA:  humans,
		Ratio:          utils.SecureDiv(mutants, humans),
	}
}
