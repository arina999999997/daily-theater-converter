package elichika

import (
	"github.com/arina999999997/daily_theater_converter/common"
	"github.com/arina999999997/daily_theater_converter/formats/canonical"

	"encoding/json"
)

func parse(context common.Context, rawData []byte) ([]common.DailyTheater, error) {
	elichikaDailyTheater := DailyTheater{}
	err := json.Unmarshal(rawData, &elichikaDailyTheater)
	if err != nil {
		return nil, err
	}
	lines, err := canonical.ParseDetailText(context, elichikaDailyTheater.DetailText)
	if err != nil {
		return nil, err
	}
	dailyTheater := common.DailyTheater{
		Language:       elichikaDailyTheater.Language,
		DailyTheaterId: elichikaDailyTheater.DailyTheaterId,
		Title:          elichikaDailyTheater.Title,
		Lines:          lines,
		Year:           elichikaDailyTheater.Year,
		Month:          elichikaDailyTheater.Month,
		Day:            elichikaDailyTheater.Day,
	}
	return []common.DailyTheater{dailyTheater}, nil
}
