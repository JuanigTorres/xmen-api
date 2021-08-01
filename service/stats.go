package service

import (
	"github.com/JuanigTorres/xmen-api/database"
	"github.com/JuanigTorres/xmen-api/model"
)

func GetStats() *model.StatsRequest {
	mutants := database.NumberOfDNAs(true)
	humans := database.NumberOfDNAs(false)
	return model.StatsRequestNew(mutants, humans)
}
