package doc

// This is the name mapping for the format
var enNameToUniqueName = map[string]string{}
var jaNameToUniqueName = map[string]string{}
var enNameToJaName = map[string]string{}
var jaNameToEnName = map[string]string{}

func init() {
	jaNameToEnName[`穂乃果`] = `Honoka`
	jaNameToEnName[`絵里`] = `Eli`
	jaNameToEnName[`ことり`] = `Kotori`
	jaNameToEnName[`海未`] = `Umi`
	jaNameToEnName[`凛`] = `Rin`
	jaNameToEnName[`真姫`] = `Maki`
	jaNameToEnName[`希`] = `Nozomi`
	jaNameToEnName[`花陽`] = `Hanayo`
	jaNameToEnName[`にこ`] = `Nico`
	jaNameToEnName[`千歌`] = `Chika`
	jaNameToEnName[`梨子`] = `Riko`
	jaNameToEnName[`果南`] = `Kanan`
	jaNameToEnName[`ダイヤ`] = `Dia`
	jaNameToEnName[`曜`] = `You`
	jaNameToEnName[`善子`] = `Yoshiko`
	jaNameToEnName[`花丸`] = `Hanamaru`
	jaNameToEnName[`鞠莉`] = `Mari`
	jaNameToEnName[`ルビィ`] = `Ruby`
	jaNameToEnName[`歩夢`] = `Ayumu`
	jaNameToEnName[`かすみ`] = `Kasumi`
	jaNameToEnName[`しずく`] = `Shizuku`
	jaNameToEnName[`果林`] = `Karin`
	jaNameToEnName[`愛`] = `Ai`
	jaNameToEnName[`彼方`] = `Kanata`
	jaNameToEnName[`せつ菜`] = `Setsuna`
	jaNameToEnName[`エマ`] = `Emma`
	jaNameToEnName[`璃奈`] = `Rina`
	jaNameToEnName[`栞子`] = `Shioriko`
	jaNameToEnName[`ランジュ`] = `Lanzhu`
	jaNameToEnName[`ミア`] = `Mia`
	// special
	jaNameToEnName[`？？？`] = `???`
	jaNameToEnName[`全員`] = `Everyone`
	jaNameToEnName[`（SE）`] = `(SE)` // sound effect
	jaNameToEnName[`はんぺん`] = `Hanpen`

	for jaName, enName := range jaNameToEnName {
		enNameToJaName[enName] = jaName
	}

	enNameToUniqueName[`Honoka`] = `HonokaKosaka`
	enNameToUniqueName[`Eli`] = `EliAyase`
	enNameToUniqueName[`Kotori`] = `KotoriMinami`
	enNameToUniqueName[`Umi`] = `UmiSonoda`
	enNameToUniqueName[`Rin`] = `RinHoshizora`
	enNameToUniqueName[`Maki`] = `MakiNishikino`
	enNameToUniqueName[`Nozomi`] = `NozomiTojo`
	enNameToUniqueName[`Hanayo`] = `HanayoKoizumi`
	enNameToUniqueName[`Nico`] = `NicoYazawa`
	enNameToUniqueName[`Chika`] = `ChikaTakami`
	enNameToUniqueName[`Riko`] = `RikoSakurauchi`
	enNameToUniqueName[`Kanan`] = `KananMatsuura`
	enNameToUniqueName[`Dia`] = `DiaKurosawa`
	enNameToUniqueName[`You`] = `YouWatanabe`
	enNameToUniqueName[`Yoshiko`] = `YoshikoTsushima`
	enNameToUniqueName[`Hanamaru`] = `HanamaruKunikida`
	enNameToUniqueName[`Mari`] = `MariOhara`
	enNameToUniqueName[`Ruby`] = `RubyKurosawa`
	enNameToUniqueName[`Ayumu`] = `AyumuUehara`
	enNameToUniqueName[`Kasumi`] = `KasumiNakasu`
	enNameToUniqueName[`Shizuku`] = `ShizukuOsaka`
	enNameToUniqueName[`Karin`] = `KarinAsaka`
	enNameToUniqueName[`Ai`] = `AiMiyashita`
	enNameToUniqueName[`Kanata`] = `KanataKonoe`
	enNameToUniqueName[`Setsuna`] = `SetsunaYuki`
	enNameToUniqueName[`Emma`] = `EmmaVerde`
	enNameToUniqueName[`Rina`] = `RinaTennoji`
	enNameToUniqueName[`Shioriko`] = `ShiorikoMifune`
	enNameToUniqueName[`Lanzhu`] = `LanzhuZhong`
	enNameToUniqueName[`Mia`] = `MiaTaylor`
	enNameToUniqueName[`???`] = `???`
	enNameToUniqueName[`(SE)`] = `SoundEffect`
	enNameToUniqueName[`Everyone`] = `Everyone`
	enNameToUniqueName[`Hanpen`] = `Hanpen`

	for jaName, enName := range jaNameToEnName {
		jaNameToUniqueName[jaName] = enNameToUniqueName[enName]
	}
}
