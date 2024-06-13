package canonical

import (
	"log"
)

var uniqueNameToGameId = map[string]string{}
var gameIdToUniqueName = map[string]string{}

// each character is tagged as <th_ch%s/>, %s is filled with ID
// for this to display correctly in the game, the corresponding petag must be present
// for characters that appear in the game itself (Muse, Aqours, Niji), the ID is a number represented stored in the above mapping
// for characters that doesn't appear in the game, we use the name itself as the ID

func addPair(uniqueName, gameId string) {
	oldId, exist := uniqueNameToGameId[uniqueName]
	if exist {
		log.Fatalf(`Error: The unique name "%s" already exist, previous game ID: "%s"`, uniqueName, oldId)
	}
	oldName, exist := gameIdToUniqueName[gameId]
	if exist {
		log.Fatalf(`Error: The game ID "%s" already exist, previous unique name: "%s"`, gameId, oldName)
	}
	uniqueNameToGameId[uniqueName] = gameId
	gameIdToUniqueName[gameId] = uniqueName
}

func getUniqueName(gameId string) string {
	name, exist := gameIdToUniqueName[gameId]
	if exist {
		return name
	}
	return gameId
}

func getGameId(uniqueName string) string {
	id, exist := uniqueNameToGameId[uniqueName]
	if exist {
		return id
	}
	return uniqueName
}

func init() {
	addPair("HonokaKosaka", "0001")
	addPair("EliAyase", "0002")
	addPair("KotoriMinami", "0003")
	addPair("UmiSonoda", "0004")
	addPair("RinHoshizora", "0005")
	addPair("MakiNishikino", "0006")
	addPair("NozomiTojo", "0007")
	addPair("HanayoKoizumi", "0008")
	addPair("NicoYazawa", "0009")
	addPair("ChikaTakami", "0101")
	addPair("RikoSakurauchi", "0102")
	addPair("KananMatsuura", "0103")
	addPair("DiaKurosawa", "0104")
	addPair("YouWatanabe", "0105")
	addPair("YoshikoTsushima", "0106")
	addPair("HanamaruKunikida", "0107")
	addPair("MariOhara", "0108")
	addPair("RubyKurosawa", "0109")
	addPair("AyumuUehara", "0201")
	addPair("KasumiNakasu", "0202")
	addPair("ShizukuOsaka", "0203")
	addPair("KarinAsaka", "0204")
	addPair("AiMiyashita", "0205")
	addPair("KanataKonoe", "0206")
	addPair("SetsunaYuki", "0207")
	addPair("EmmaVerde", "0208")
	addPair("RinaTennoji", "0209")
	addPair("ShiorikoMifune", "0210")
	addPair("LanzhuZhong", "0211")
	addPair("MiaTaylor", "0212")

	// special name
	addPair("???", "0000")
	addPair("LaterOn", "1001")
	addPair("SoundEffect", "1002")
	addPair("Everyone", "1003")
	addPair("Hanpen", "5206")
}
