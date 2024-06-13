package custom_makers

import (
	"github.com/arina999999997/daily_theater_converter/common"
)

// setting up https://zura3395.github.io/daily-theatre/
var zuraMakerConfig *CustomMakerConfig

func zuraMakerParse(context common.Context, rawData []byte) ([]common.DailyTheater, error) {
	context.Credit = "Translated by SIFAStheatre and Idol Story"
	a, b := parse(zuraMakerConfig, context, rawData)
	return a, b
}

func zuraMakerSerialize(context common.Context, dailyTheater common.DailyTheater) ([]byte, error) {
	a, b := serialize(zuraMakerConfig, context, dailyTheater)
	return a, b
}

func init() {
	config := getDefaultConfig()
	config.AddDefaultCharacter("YuTakasaki", "30")
	config.AddDefaultCharacter("KanonShibuya", "31")
	config.AddDefaultCharacter("KekeTang", "32")
	config.AddDefaultCharacter("ChisatoArashi", "33")
	config.AddDefaultCharacter("SumireHeanna", "34")
	config.AddDefaultCharacter("RenHazuki", "35")
	config.AddDefaultCharacter("KinakoSakurakoji", "36")
	config.AddDefaultCharacter("MeiYoneme", "37")
	config.AddDefaultCharacter("ShikiWakana", "38")
	config.AddDefaultCharacter("NatsumiOnitsuka", "39")
	config.AddDefaultCharacter("WienMargarete", "40")
	config.AddDefaultCharacter("TomariOnitsuka", "41")
	zuraMakerConfig = config

	common.AddFormat("zura_maker", zuraMakerParse, zuraMakerSerialize)
}
