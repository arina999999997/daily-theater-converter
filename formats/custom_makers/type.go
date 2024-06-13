package custom_makers

import (
	"github.com/arina999999997/daily_theater_converter/common"

	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"errors"
)

type GenericCustom struct {
	DisplayName     string `json:"s"`
	BackgroundColor string `json:"c"`
}

type OfficialCustom struct {
	DisplayName struct {
		Jp string `json:"jp"`
		En string `json:"en"`
		Zh string `json:"zh"`
		Kr string `json:"kr"`
		Th string `json:"th"`
	} `json:"s"`
	Image string `json:"i"`
}

type CustomMakerGroup struct {
	GroupName string
	Members   []OfficialCustom
}

type Extensions struct {
	CustomCharacter []GenericCustom
	CustomGroup     []CustomMakerGroup
}

func (e Extensions) MarshalJSON() ([]byte, error) {
	result := []byte{}
	result = append(result, []byte(`{"custom":`)...)
	bytes, err := json.Marshal(e.CustomCharacter)
	if err != nil {
		return nil, err
	}
	result = append(result, bytes...)
	for _, group := range e.CustomGroup {
		result = append(result, []byte(`,"`)...)
		result = append(result, []byte(group.GroupName)...)
		result = append(result, []byte(`":`)...)
		bytes, err = json.Marshal(group.Members)
		if err != nil {
			return nil, err
		}
		result = append(result, bytes...)
	}
	result = append(result, '}')
	return result, nil
}

func (e *Extensions) UnmarshalJSON(data []byte) error {
	objectMap := map[string]json.RawMessage{}
	err := json.Unmarshal(data, &objectMap)
	if err != nil {
		return err
	}
	for key, message := range objectMap {
		if key == "custom" {
			err = json.Unmarshal(message, &e.CustomCharacter)
			if err != nil {
				return err
			}
		} else {
			group := CustomMakerGroup{
				GroupName: key,
			}
			err = json.Unmarshal(message, &group.Members)
			if err != nil {
				return err
			}
			e.CustomGroup = append(e.CustomGroup, group)
		}
	}
	return nil
}

type Character struct {
	Id      string
	IdNum   int
	IdGroup string
}

func (c *Character) UnmarshalJSON(data []byte) error {
	if data[0] == '"' { // is a string, marshal as is
		c.Id = string(data[1 : len(data)-1])
	} else { // is a number, but we can save it as is
		c.Id = string(data)
	}
	index := strings.LastIndex(c.Id, "_")
	if index != -1 {
		c.IdGroup = c.Id[:index]
	}
	index, err := strconv.Atoi(c.Id[index+1:])
	if err != nil {
		return err
	}
	c.IdNum = index
	return nil
}

func (c Character) MarshalJSON() ([]byte, error) {
	_, err := strconv.Atoi(c.Id)
	if err != nil { // not a number wrap it in ""
		return []byte(fmt.Sprintf(`"%s"`, c.Id)), nil
	} else { // is a number
		return []byte(c.Id), nil
	}
}

type CustomMakerLine struct {
	Character Character `json:"n"`
	Text      string    `json:"d"`
}

type CustomMakerData struct {
	Header     string            `json:"header"` // can be missing, but will always be marshaled by this library
	Title      string            `json:"title"`  // as above
	Extensions Extensions        `json:"extensions"`
	Content    []CustomMakerLine `json:"content"`
}

func (c *Character) GetUniqueName(config *CustomMakerConfig, data *CustomMakerData, context common.Context) (string, error) {
	character, exist := config.CharacterIdToCharacter[c.Id]
	if exist {
		return character.UniqueName, nil
	}
	// this character is not recognised, we use the raw name as is, which is the name in custom
	if c.IdGroup != "custom" {
		return "", errors.New(fmt.Sprintf(`Character with id "%s" isn't a known character, and isn't defined in the custom character lists. Either the data is wrong, or this library is outdated and need updating.`, c.Id))
	}
	return data.Extensions.CustomCharacter[c.IdNum].DisplayName, nil
}

func (c *CustomMakerLine) ConvertToCommonLine(config *CustomMakerConfig, data *CustomMakerData, context common.Context) (common.Line, error) {
	uniqueName, err := c.Character.GetUniqueName(config, data, context)
	return common.Line{
		Character: uniqueName,
		Text:      c.Text,
	}, err
}
