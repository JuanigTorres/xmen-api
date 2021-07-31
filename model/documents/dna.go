package documents

import (
	"go.mongodb.org/mongo-driver/bson"
	"strings"
)

type DNADocument struct {
	DNA  string // A guion separated list
	Mutant bool
}

func NewDNADocument(dna []string, isMutant bool) *DNADocument {
	return &DNADocument{
		DNA: strings.Join(dna, "-"),
		Mutant: isMutant,
	}
}

func (document *DNADocument) AsBson() *bson.D {
	return &bson.D {
		{ "dna", document.DNA },
		{ "mutant", document.Mutant },
	}
}