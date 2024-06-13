package custom_makers

import (
	"github.com/arina999999997/daily_theater_converter/common"

	// "errors"
	"encoding/json"
	"log"
	"strings"
	"time"
)

var startDate = time.Date(2020, time.February, 2, 0, 0, 0, 0, time.UTC)

func parse(config *CustomMakerConfig, context common.Context, rawData []byte) ([]common.DailyTheater, error) {
	data := CustomMakerData{}
	err := json.Unmarshal(rawData, &data)
	if err != nil {
		return nil, err
	}

	dailyTheater := context.GetDefaultDailyTheater()
	dailyTheater.Title = data.Title

	filePath := strings.Replace(context.FilePath, "\\", "/", -1)
	if filePath != "" {
		// assume the file path is something like /.../.../.../yyyymmdd.json
		// or something like /.../.../.../yyyymmdd.txt
		// or something like /.../.../.../yyyymmdd
		// then we can get the year and day, and from there we can also derive the id, assuming we only work with official story
		tokens := strings.Split(filePath, "/")
		fileName := tokens[len(tokens)-1]
		tokens = strings.Split(fileName, ".")
		dateString := tokens[0]
		date, err := time.Parse("20060102", dateString)
		if err == nil {
			dailyTheater.Year = int32(date.Year())
			dailyTheater.Month = int32(date.Month())
			dailyTheater.Day = int32(date.Day())
			diff := date.Sub(startDate)
			dailyTheater.DailyTheaterId = int32(diff.Hours()/24 + 1000000)
		} else {
			log.Println("Warning: could not deduce date into for file: ", filePath)
		}
	}

	// Adding credit when parsing from specific source
	if context.Credit != "" {
		data.Content = append(data.Content, CustomMakerLine{
			Character: Character{
				Id: "???",
			},
			Text: context.Credit,
		})
	}
	// parse the line
	for _, line := range data.Content {
		commonLine, err := line.ConvertToCommonLine(config, &data, context)
		if err != nil {
			return nil, err
		}
		dailyTheater.Lines = append(dailyTheater.Lines, commonLine)
	}

	return []common.DailyTheater{dailyTheater}, nil
}
