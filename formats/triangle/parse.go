package canonical

import (
	"github.com/arina999999997/daily_theater_converter/common"
	"github.com/arina999999997/daily_theater_converter/formats/canonical"

	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

var location *time.Location

func init() {
	location, _ = time.LoadLocation("Asia/Tokyo")
}
func parse(context common.Context, rawData []byte) ([]common.DailyTheater, error) {
	result := []common.DailyTheater{}
	reader := csv.NewReader(bytes.NewBuffer(rawData))
	for {
		line, err := reader.Read()
		if err != nil {
			if err != io.EOF {
				return nil, err
			}
			break
		}
		dailyTheater, err := parseLine(context, line)
		if err != nil {
			return nil, err
		}
		result = append(result, dailyTheater)
	}
	return result, nil
}

// the line format is as follow:
// language,id,published_at,daily_theater_menu_text,daily_theater_content
// - published_at is in the format yyyy-mm-dd HH:MM:SS, the zone is Etc/UTC
// - daily_theater_menu_text is always the title for daily theater
// - daily_theater_content has both the title and text
func parseLine(context common.Context, field []string) (common.DailyTheater, error) {
	dailyTheater := common.DailyTheater{}
	if len(field) != 5 {
		return dailyTheater, errors.New(fmt.Sprintln(`Wrong number of field: `, field))
	}
	dailyTheater.Language = field[0]
	if dailyTheater.Language == "jp" {
		dailyTheater.Language = "ja"
	}
	dailyTheaterId, err := strconv.Atoi(field[1])
	if err != nil {
		return dailyTheater, err
	}
	dailyTheater.DailyTheaterId = int32(dailyTheaterId)
	timePoint, err := time.Parse(`2006-01-02 15:04:00`, field[2])
	if err != nil {
		return dailyTheater, err
	}
	timePoint = timePoint.In(location)
	dailyTheater.Year = int32(timePoint.Year())
	dailyTheater.Month = int32(timePoint.Month())
	dailyTheater.Day = int32(timePoint.Day())
	contents := strings.Split(field[4], "<:dt_title_end/>")
	if len(contents) != 2 {
		return dailyTheater, errors.New(fmt.Sprintf(`Unexpected daily theater content: %s`, contents))
	}
	start := strings.Index(contents[0], "『")
	end := strings.Index(contents[0], "』")
	if (start == -1) || (end == -1) || (start > end) {
		return dailyTheater, errors.New(fmt.Sprintf(`Unexpected daily theater title format: %s`, contents[0]))
	}
	dailyTheater.Title = contents[0][start+3 : end] // 『 has width 3
	lines, err := canonical.ParseDetailText(context, contents[1])
	if err != nil {
		return dailyTheater, err
	}
	dailyTheater.Lines = lines
	return dailyTheater, nil
}
