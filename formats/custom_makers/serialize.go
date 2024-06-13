package custom_makers

import (
	"github.com/arina999999997/daily_theater_converter/common"

	"encoding/json"
	"fmt"
)

func serialize(config *CustomMakerConfig, context common.Context, dailyTheater common.DailyTheater) ([]byte, error) {
	// serializing has to skip some info:
	// - mainly, the fully customized characters will be missing background colour, and the display name is changed a bit
	data := CustomMakerData{
		Header: "Daily theater",
		Title:  dailyTheater.Title,
	}
	if context.Language != "en" {
		data.Header = "毎日劇場"
	}
	hasGroupExtension := map[string]bool{}
	uniqueNameToCustomCharacterPosition := map[string]int{}

	for _, line := range dailyTheater.Lines {
		customMakerLine := CustomMakerLine{
			Text: line.Text,
		}
		// map to original id
		customMakerCharacter, exist := config.UniqueNameToCharacter[line.Character]
		if exist {
			// default character or supported extension character
			if customMakerCharacter.OfficialGroup != "" {
				hasGroupExtension[customMakerCharacter.OfficialGroup] = true
			}
			customMakerLine.Character.Id = customMakerCharacter.Id
		} else {
			pos, exist := uniqueNameToCustomCharacterPosition[line.Character]
			if !exist {
				uniqueNameToCustomCharacterPosition[line.Character] = len(data.Extensions.CustomCharacter)
				data.Extensions.CustomCharacter = append(data.Extensions.CustomCharacter, GenericCustom{
					DisplayName:     line.Character,
					BackgroundColor: "#D1D1D1",
				})
				pos = uniqueNameToCustomCharacterPosition[line.Character]
			}
			// custom character
			customMakerLine.Character.Id = fmt.Sprintf(`custom_%d`, pos)
		}
		data.Content = append(data.Content, customMakerLine)
	}
	// finally enable the group, if there's any
	for group := range hasGroupExtension {
		data.Extensions.CustomGroup = append(data.Extensions.CustomGroup, config.Groups[group])
	}
	bytes, err := json.Marshal(data)
	return bytes, err
}
