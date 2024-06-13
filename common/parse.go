package common

import (
	"errors"
	"fmt"
	"os"
)

func Parse(context Context, rawData []byte) ([]DailyTheater, error) {
	parser, exist := parsers[context.Format]
	if !exist {
		return nil, errors.New(fmt.Sprintf(`Error: format "%s" doesn't have a parser`, context.Format))
	}
	dailyTheaters, err := parser(context, rawData)
	return dailyTheaters, err
}

func ParseFile(context Context, path string) ([]DailyTheater, error) {
	rawData, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	context.FilePath = path
	dailyTheaters, err := Parse(context, rawData)
	return dailyTheaters, err
}
