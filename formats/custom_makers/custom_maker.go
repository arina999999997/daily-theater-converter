package custom_makers

import (
	"github.com/arina999999997/daily_theater_converter/common"

	"encoding/json"
)

// setting up https://twy.name/LLAS/mainichi/
var customMakerConfig *CustomMakerConfig

func customMakerParse(context common.Context, rawData []byte) ([]common.DailyTheater, error) {
	a, b := parse(customMakerConfig, context, rawData)
	return a, b
}

func customMakerSerialize(context common.Context, dailyTheater common.DailyTheater) ([]byte, error) {
	a, b := serialize(customMakerConfig, context, dailyTheater)
	return a, b
}

func init() {
	config := getDefaultConfig()
	config.AddOfficialCustomCharacter("KanonShibuya", "liella_0", "liella")
	config.AddOfficialCustomCharacter("KekeTang", "liella_1", "liella")
	config.AddOfficialCustomCharacter("ChisatoArashi", "liella_2", "liella")
	config.AddOfficialCustomCharacter("SumireHeanna", "liella_3", "liella")
	config.AddOfficialCustomCharacter("RenHazuki", "liella_4", "liella")
	config.AddOfficialCustomCharacter("KinakoSakurakoji", "liella_5", "liella")
	config.AddOfficialCustomCharacter("MeiYoneme", "liella_6", "liella")
	config.AddOfficialCustomCharacter("ShikiWakana", "liella_7", "liella")
	config.AddOfficialCustomCharacter("NatsumiOnitsuka", "liella_8", "liella")
	config.AddOfficialCustomCharacter("WienMargarete", "liella_9", "liella")
	config.AddOfficialCustomCharacter("TomariOnitsuka", "liella_10", "liella")

	config.AddOfficialCustomCharacter("KahoHinoshita", "hasunosora_0", "hasunosora")
	config.AddOfficialCustomCharacter("SayakaMurano", "hasunosora_1", "hasunosora")
	config.AddOfficialCustomCharacter("KozueOtomune", "hasunosora_2", "hasunosora")
	config.AddOfficialCustomCharacter("TsuzuriYugiri", "hasunosora_3", "hasunosora")
	config.AddOfficialCustomCharacter("RurinoOsawa", "hasunosora_4", "hasunosora")
	config.AddOfficialCustomCharacter("MegumiFujishima", "hasunosora_5", "hasunosora")
	// these members are not present at the time of implementing this
	// send and issue or a pull request if you wish to update it
	// config.AddOfficialCustomCharacter("GinkoMomose", "hasunosora_6", "hasunosora")
	// config.AddOfficialCustomCharacter("KosuzuKachimachi", "hasunosora_7", "hasunosora")
	// config.AddOfficialCustomCharacter("HimeAnyoji", "hasunosora_8", "hasunosora")
	customMakerConfig = config

	hasunosoraJson := `[{"s":{"jp":"花帆","en":"Kaho","zh":"花帆","kr":"카호","th":"คาโฮะ"},"i":"kaho"},{"s":{"jp":"さやか","en":"Sayaka","zh":"沙耶香","kr":"사야카","th":"ซายากะ"},"i":"sayaka"},{"s":{"jp":"梢","en":"Kozue","zh":"梢","kr":"코즈에","th":"โคซุโอะ"},"i":"kozue"},{"s":{"jp":"綴理","en":"Tsuzuri","zh_t":"綴理","zh_s":"缀理","kr":"츠즈리","th":"ซึสุริ"},"i":"tsuzuri"},{"s":{"jp":"瑠璃乃","en":"Rurino","zh":"瑠璃乃","kr":"루리노","th":"รุริโนะ"},"i":"rurino"},{"s":{"jp":"慈","en":"Megumi","zh":"慈","kr":"메구미","th":"เมกุมิ"},"i":"megumi"}]`
	hasunosora := CustomMakerGroup{
		GroupName: "hasunosora",
	}
	err := json.Unmarshal([]byte(hasunosoraJson), &hasunosora.Members)
	if err != nil {
		panic(err)
	}
	config.Groups["hasunosora"] = hasunosora

	liellaJson := `[{"s":{"jp":"かのん","en":"Kanon","zh":"香音","kr":"카논","th":"คานง"},"i":"kanon"},{"s":{"jp":"可可","en":"Keke","zh":"可可","kr":"쿠쿠","th":"เขอเข่อ"},"i":"keke"},{"s":{"jp":"千砂都","en":"Chisato","zh":"千砂都","kr":"치사토","th":"จิซาโตะ"},"i":"chisato"},{"s":{"jp":"すみれ","en":"Sumire","zh":"堇","kr":"스미레","th":"สุมิเระ"},"i":"sumire"},{"s":{"jp":"恋","en":"Ren","zh_t":"戀","zh_s":"恋","kr":"렌","th":"เร็ง"},"i":"ren"},{"s":{"jp":"きな子","en":"Kinako","zh":"希奈子","kr":"키나코","th":"คินาโกะ"},"i":"kinako"},{"s":{"jp":"メイ","en":"Mei","zh":"芽衣","kr":"메이","th":"เมย์"},"i":"mei"},{"s":{"jp":"四季","en":"Shiki","zh":"四季","kr":"시키","th":"ชิกิ"},"i":"shiki"},{"s":{"jp":"夏美","en":"Natsumi","zh":"夏美","kr":"나츠미","th":"นัตสึมิ"},"i":"natsumi"},{"s":{"jp":"マルガレーテ","en":"Margarete","zh_t":"瑪格麗特","zh_s":"玛格丽特","kr":"마르가레테","th":"มากาเร็ต"},"i":"margarete"},{"s":{"jp":"冬毬","en":"Tomari","zh":"冬毬","kr":"토마리","th":"โทมาริ"},"i":"tomari"}]`
	liella := CustomMakerGroup{
		GroupName: "liella",
	}
	err = json.Unmarshal([]byte(liellaJson), &liella.Members)
	if err != nil {
		panic(err)
	}
	config.Groups["liella"] = liella

	common.AddFormat("custom_maker", customMakerParse, customMakerSerialize)
}
