package service_test

import (
	"testing"

	"github.com/JuanigTorres/xmen-finder/service"
)

var MUTANT_DNA_LIST = [][]string {
	{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"},
	{"ATGCGA", "CCGTGC", "TTGTTT", "AGTTGG", "CCCTTA", "TCACTG"},
	{"ATGCGA", "CCGTGC", "TAGTCT", "AGTCTG", "CCCTTA", "TCTCTG"},
	{"ATGCGA", "CCGTGC", "TTTTCT", "AGTCTG", "CCCTTA", "TCTCTG"},
}

var HUMAN_DNA_LIST = [][]string {
	{"ATGCGA", "CCGTGC", "TTGTTT", "AGTAGG", "CCCTTA", "TCACTG"},
	{"ATGCGA", "CCGTGC", "TAGTTT", "AGTTGG", "CCCTTA", "TCACTG"},
	{"ATGCGA", "CCGTGC", "TTCTCT", "AGGCTG", "CCAATA", "TCTCTG"},
	{"ATGCGA", "GCGTGC", "TACTTT", "AGGCTG", "CCCACA", "TCTCTG"},
}

var INVALID_DNA_LIST = [][]string {
	{},
	{"ATGCGA", "CCGTGC", "TAGTTT", "AGTTGG", "CCCTTA"},
	{"ATGCGA", "CCGZGC", "TAGTTT", "AGTTGG", "CCCTTA", "TCACTG"},
	{"ATGCGA", "CCGTGC", "TAGTTT", "AGT11G", "CCCTTA", "TCACTG"},
}

func TestMutantDNAs(t *testing.T) {
	for _, dna := range MUTANT_DNA_LIST {
		truthy := service.IsMutant(dna)
		if !truthy {
			t.FailNow()
		}
	}
}

func TestHumanDNAs(t *testing.T) {
	for _, dna := range HUMAN_DNA_LIST {
		truthy := service.IsMutant(dna)
		if truthy {
			t.FailNow()
		}
	}
}

func TestInvalidDNAs(t *testing.T) {
	for _, dna := range INVALID_DNA_LIST {
		truthy := service.IsMutant(dna)
		if truthy {
			t.FailNow()
		}
	}
}