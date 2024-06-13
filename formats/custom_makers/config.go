package custom_makers

// a CustomMakerCharacter is defined as the one of the default characters, or the characters who have an OfficialGroup
// when referenced by the UniqueName, if a character doesn't have a CustomMakerCharacter object,
//
//	then they're treated as GenericCustom
type CustomMakerCharacter struct {
	UniqueName    string `json:"unique_name"`
	Id            string `json:"id"`
	OfficialGroup string `json:"official_group"`
}

type CustomMakerConfig struct {
	UniqueNameToCharacter  map[string]CustomMakerCharacter
	CharacterIdToCharacter map[string]CustomMakerCharacter
	Groups                 map[string]CustomMakerGroup
}

func (c *CustomMakerConfig) Init() {
	if c.UniqueNameToCharacter == nil {
		c.UniqueNameToCharacter = map[string]CustomMakerCharacter{}
		c.CharacterIdToCharacter = map[string]CustomMakerCharacter{}
		c.Groups = map[string]CustomMakerGroup{}
	}
}

func (c *CustomMakerConfig) AddCharacter(uniqueName string, character CustomMakerCharacter) {
	c.Init()
	character.UniqueName = uniqueName
	_, exist := c.UniqueNameToCharacter[uniqueName]
	if exist {
		panic("character already exist")
	}
	_, exist = c.CharacterIdToCharacter[character.Id]
	if exist {
		panic("character id already exist")
	}
	c.UniqueNameToCharacter[uniqueName] = character
	c.CharacterIdToCharacter[character.Id] = character
}

func (c *CustomMakerConfig) AddDefaultCharacter(uniqueName, id string) {
	c.AddCharacter(uniqueName, CustomMakerCharacter{
		Id: id,
	})
}

func (c *CustomMakerConfig) AddOfficialCustomCharacter(uniqueName, id, group string) {
	c.AddCharacter(uniqueName, CustomMakerCharacter{
		Id:            id,
		OfficialGroup: group,
	})
}

func getDefaultConfig() *CustomMakerConfig {
	config := new(CustomMakerConfig)
	config.AddDefaultCharacter("HonokaKosaka", "0")
	config.AddDefaultCharacter("EliAyase", "1")
	config.AddDefaultCharacter("KotoriMinami", "2")
	config.AddDefaultCharacter("UmiSonoda", "3")
	config.AddDefaultCharacter("RinHoshizora", "4")
	config.AddDefaultCharacter("MakiNishikino", "5")
	config.AddDefaultCharacter("NozomiTojo", "6")
	config.AddDefaultCharacter("HanayoKoizumi", "7")
	config.AddDefaultCharacter("NicoYazawa", "8")
	config.AddDefaultCharacter("ChikaTakami", "9")
	config.AddDefaultCharacter("RikoSakurauchi", "10")
	config.AddDefaultCharacter("KananMatsuura", "11")
	config.AddDefaultCharacter("DiaKurosawa", "12")
	config.AddDefaultCharacter("YouWatanabe", "13")
	config.AddDefaultCharacter("YoshikoTsushima", "14")
	config.AddDefaultCharacter("HanamaruKunikida", "15")
	config.AddDefaultCharacter("MariOhara", "16")
	config.AddDefaultCharacter("RubyKurosawa", "17")
	config.AddDefaultCharacter("AyumuUehara", "18")
	config.AddDefaultCharacter("KasumiNakasu", "19")
	config.AddDefaultCharacter("ShizukuOsaka", "20")
	config.AddDefaultCharacter("KarinAsaka", "21")
	config.AddDefaultCharacter("AiMiyashita", "22")
	config.AddDefaultCharacter("KanataKonoe", "23")
	config.AddDefaultCharacter("SetsunaYuki", "24")
	config.AddDefaultCharacter("EmmaVerde", "25")
	config.AddDefaultCharacter("RinaTennoji", "26")
	config.AddDefaultCharacter("ShiorikoMifune", "27")
	config.AddDefaultCharacter("LanzhuZhong", "28")
	config.AddDefaultCharacter("MiaTaylor", "29")
	config.AddDefaultCharacter("???", "???")
	config.AddDefaultCharacter("___", "Everyone")
	return config
}
