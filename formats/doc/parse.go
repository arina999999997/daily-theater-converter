package doc

import (
	"github.com/arina999999997/daily_theater_converter/common"

	"errors"
	"fmt"
	"log"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// Format is as follow:
// - Original Japanese text followed by English text of the same day
// - The year is supplied in the file name (just renamed manually)
// - To parse, we will continue parsing from top to bottom of the file, keeping everything
// - A title have to have format:
//   - <dd>/<mm> - <title>
//   - Reading a title will flush the existing daily theater and create a new one
// - A line have to have format:
//   - <name>:<text> for Japanese
//   - <name> : <text> for English
//   - The actual interface is to split by the first : and strip from both side
//   - if the name is not valid, then it's not a line, an error is displayed just to be safe.
//   - The parsed line is appended to the current daily theater, and the name is used to decide the language
// - If a line is not of the above format then it's ignored.
//
// Because this format is not fully machine friendly, the parsing process will not stop if it encounter "errors":
// - Some errors are not critical and doesn't affect the accuracy of the parsed data.
// - Human verification is required to make sure the parsed data is good
// Generally, this format should only be used manually

var startDate = time.Date(2020, time.February, 2, 0, 0, 0, 0, time.UTC)

func tryParseTitle(context common.Context, year int, text string) *common.DailyTheater {
	index := strings.Index(text, "-")
	if index == -1 {
		return nil
	}
	dateString := strings.TrimSpace(text[:index])
	titleString := strings.TrimSpace(text[index+1:])
	if len(dateString) != 5 {
		log.Printf("File: \"%s\". Unexpected title line: \"%s\"\n", context.FilePath, dateString)
		return nil
	}
	date, err := time.Parse(`02/01/2006`, fmt.Sprintf(`%s/%d`, dateString, year))
	if err != nil {
		log.Println(`File: "`, context.FilePath, `". Not a fatal error: `, err, `Line: "`, text, `"`)
		return nil
	}

	dailyTheater := new(common.DailyTheater)
	dailyTheater.Year = int32(date.Year())
	dailyTheater.Month = int32(date.Month())
	dailyTheater.Day = int32(date.Day())
	diff := date.Sub(startDate)
	dailyTheater.DailyTheaterId = int32(diff.Hours()/24 + 1000000)
	if dailyTheater.DailyTheaterId < 1000001 || dailyTheater.DailyTheaterId > 1001244 {
		log.Println(`File: "`, context.FilePath, `". Not a fatal error: id out of expected range. Line: "`, text, `"`)
	}
	dailyTheater.Title = titleString
	return dailyTheater
}

func tryParseLine(context common.Context, text string) (common.Line, string) {
	index := strings.Index(text, ":")
	dialogStart := index + 1
	if index == -1 {
		index = strings.Index(text, "：")
		dialogStart = index + 3 // ：take up 3 char in utf-8 encoding
		if index == -1 {
			return common.Line{}, ""
		}
	}
	nameString := strings.TrimSpace(text[:index])
	textString := strings.TrimSpace(text[dialogStart:])
	uniqueName, exist := enNameToUniqueName[nameString]
	if exist { // english name
		return common.Line{
			Character: uniqueName,
			Text:      textString,
		}, "en"
	}
	uniqueName, exist = jaNameToUniqueName[nameString]
	if exist {
		return common.Line{
			Character: uniqueName,
			Text:      textString,
		}, "ja"
	}
	log.Printf("File: \"%s\". Unexpected (spoken) line: \"%s\"\n", context.FilePath, text)
	return common.Line{}, ""
}

func parse(context common.Context, rawData []byte) ([]common.DailyTheater, error) {
	if context.FilePath == "" {
		return nil, errors.New(fmt.Sprintf(`Format "%s" only works with files`, formatName))
	}
	fullText := string(rawData)
	textLines := strings.Split(fullText, "\r\n")
	result := []common.DailyTheater{}
	year, err := strconv.Atoi(filepath.Base(context.FilePath)[:4])
	if err != nil {
		panic(fmt.Sprintf(`Error: Can't parse year for file: %s`, context.FilePath))
	}
	var current *common.DailyTheater
	for _, textLine := range textLines {
		if (textLine == "Translation : ") || (textLine == "Translation:") || (textLine == "Translation :") {
			continue
		}
		line, language := tryParseLine(context, textLine)
		if language != "" { // valid line
			if current.Language == "" {
				current.Language = language
			} else if current.Language != language {
				panic(fmt.Sprintf(`Error: Language changed. File: %s, line: "%s"`, context.FilePath, textLine))
			}
			current.Lines = append(current.Lines, line)
			continue
		}
		new := tryParseTitle(context, year, textLine)
		if new != nil {
			if current != nil {
				result = append(result, *current)
			}
			current = new
		}
	}
	if current != nil {
		result = append(result, *current)
	}
	for i := range result {
		if result[i].Language == "en" {
			result[i].Lines = append(result[i].Lines, common.Line{
				Character: "???",
				Text:      "Translated by SIFAStheatre and Idol Story",
			})
		}
	}
	return result, nil
}
