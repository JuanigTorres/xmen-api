package service

import (
	"github.com/JuanigTorres/xmen-finder/database"
	"github.com/JuanigTorres/xmen-finder/model"
)

func GetStats() *model.StatsRequest {
	mutants := database.NumberOfDNAs(true)
	humans := database.NumberOfDNAs(false)
	return model.StatsRequestNew(mutants, humans)
}
