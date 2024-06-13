package elichika

import (
	"github.com/arina999999997/daily_theater_converter/common"
	"github.com/arina999999997/daily_theater_converter/formats/canonical"

	"encoding/json"
)

func serialize(_ common.Context, dailyTheater common.DailyTheater) ([]byte, error) {
	detailText := canonical.SerializeDetailText(dailyTheater)
	elichikaDailyTheater := DailyTheater{
		Language:       dailyTheater.Language,
		DailyTheaterId: dailyTheater.DailyTheaterId,
		Title:          dailyTheater.Title,
		DetailText:     detailText,
		Year:           dailyTheater.Year,
		Month:          dailyTheater.Month,
		Day:            dailyTheater.Day,
	}
	bytes, err := json.Marshal(elichikaDailyTheater)
	return bytes, err
}
