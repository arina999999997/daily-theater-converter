package canonical

import (
	"github.com/arina999999997/daily_theater_converter/common"

	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// parse the lines from the text
func ParseDetailText(context common.Context, text string) ([]common.Line, error) {
	lineTexts := strings.Split(text, `<:dt_line_end/>`)
	lines := []common.Line{}
	for i, line := range lineTexts {
		if i == len(lineTexts)-1 {
			if line != "" {
				return nil, errors.New("Unexpected end of detail text")
			}
			break
		}
		if !strings.HasPrefix(line, "<:th_ch") {
			return nil, errors.New(fmt.Sprintf("Unexpected start of line. Line: %s", line))
		}
		tagEnd := strings.Index(line, "/>")
		if tagEnd == -1 {
			return nil, errors.New(fmt.Sprintf("Unexpected start of line. Line: %s", line))
		}
		characterId := line[7:tagEnd]
		text := line[tagEnd+2:]
		lines = append(lines, common.Line{
			Character: getUniqueName(characterId),
			Text:      text,
		})
	}
	return lines, nil
}

func parse(context common.Context, rawData []byte) ([]common.DailyTheater, error) {
	detail := DailyTheaterDetail{}
	err := json.Unmarshal(rawData, &detail)
	if err != nil {
		return nil, err
	}
	lines, err := ParseDetailText(context, detail.DetailText.DotUnderText)
	if err != nil {
		return nil, err
	}
	dailyTheater := common.DailyTheater{
		Language:       context.Language,
		DailyTheaterId: detail.DailyTheaterId,
		Title:          detail.Title.DotUnderText,
		Lines:          lines,
		Year:           detail.Year,
		Month:          detail.Month,
		Day:            detail.Day,
	}
	return []common.DailyTheater{dailyTheater}, nil
}
