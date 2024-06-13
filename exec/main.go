package main

// This is the demo executable file
import (
	_ "github.com/arina999999997/daily_theater_converter"
	"github.com/arina999999997/daily_theater_converter/common"

	"fmt"
	"io/fs"
	"log"
	"os"
	pathLib "path"
	"path/filepath"
)

func main() {
	if len(os.Args) != 7 {
		log.Fatalf("Usage: %s <language> <format_to_parse> <path/to/dir/to/parse> <format_to_serialize> <path/to/dir/to/format_to_serialize> <file_type>", os.Args[0])
	}
	language := os.Args[1]
	parseFormat := os.Args[2]
	parseDir := os.Args[3]
	parseContext := common.NewContext(parseFormat, language)
	serializeFormat := os.Args[4]
	serializeDir := os.Args[5]
	fileType := os.Args[6]
	serializeContext := common.NewContext(serializeFormat, language)
	all := map[string]map[int32]common.DailyTheater{}
	err := filepath.Walk(parseDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		extension := pathLib.Ext(path)
		if extension != ".json" && extension != ".txt" && extension != ".csv" && extension != "" {
			return nil
		}
		// file
		dailyTheaters, err := common.ParseFile(parseContext, path)
		if err != nil {
			log.Println("Error parsing file: ", path, ": ", err, ". Skipping!")
		}
		for _, dailyTheater := range dailyTheaters {
			_, exist := all[dailyTheater.Language]
			if !exist {
				all[dailyTheater.Language] = map[int32]common.DailyTheater{}
			}
			_, exist = all[dailyTheater.Language][dailyTheater.DailyTheaterId]
			if exist {
				log.Printf(`Warning: multiple daily theater with id %d and language %s`, dailyTheater.DailyTheaterId, dailyTheater.Language)
			}
			all[dailyTheater.Language][dailyTheater.DailyTheaterId] = dailyTheater
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	os.MkdirAll(serializeDir, 0644)
	for language, idMap := range all {
		for id, dailyTheater := range idMap {
			bytes, err := common.Serialize(serializeContext, dailyTheater)
			if err != nil {
				panic(err)
			}
			fileName := fmt.Sprintf("%s-%d.%s", language, id, fileType)
			err = os.WriteFile(fmt.Sprintf(`%s/%s`, serializeDir, fileName), bytes, 0644)
			if err != nil {
				panic(err)
			}
		}
	}
}
