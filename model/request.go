package model

import "github.com/JuanigTorres/xmen-api/math/utils"

type XmenRequest struct {
	DNA []string
}

type StatsRequest struct {
	CountMutantDNA int64
	CountHumanDNA  int64
	Ratio          float32
}

func StatsRequestNew(mutants, humans int64) *StatsRequest {
	return &StatsRequest{
		CountMutantDNA: mutants,
		CountHumanDNA:  humans,
		Ratio:          utils.SecureDiv(mutants, humans),
	}
}
