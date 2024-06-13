package canonical

import (
	"github.com/arina999999997/daily_theater_converter/common"

	"encoding/json"
	"fmt"
)

func SerializeDetailText(dailyTheater common.DailyTheater) string {
	detailText := ""
	for _, line := range dailyTheater.Lines {
		detailText += fmt.Sprintf(`<:th_ch%s/>%s<:dt_line_end/>`, getGameId(line.Character), line.Text)
	}
	return detailText
}

func serialize(_ common.Context, dailyTheater common.DailyTheater) ([]byte, error) {
	detail := DailyTheaterDetail{
		DailyTheaterId: dailyTheater.DailyTheaterId,
		Title: LocalizedText{
			DotUnderText: dailyTheater.Title,
		},
		DetailText: LocalizedText{
			DotUnderText: SerializeDetailText(dailyTheater),
		},
		Year:  dailyTheater.Year,
		Month: dailyTheater.Month,
		Day:   dailyTheater.Day,
	}
	bytes, err := json.Marshal(detail)
	return bytes, err
}
