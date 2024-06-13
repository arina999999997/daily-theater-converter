# Daily theater converter
The converter between daily theater format for sifas. 

This is a go module that can be used to convert between formats. However, the repository also contain some raw data and a standalone executable.
## Technical details
### Common representation
The text is parsed from original format to the common representation (in memory), then can be exported back in any format.

The common representation is as follow:

- Language: The language of the daily theater.
- Id: The daily theater id. This is always present for official daily theater and muat be unique per language.
- Title: The title of the daily theater.
- Year, month, day: The date of the daily theater.
- Lines: The lines spoken in the daily theater in order.

For each line, there is:

- Character - The character who speak the line:

    - The characters are represented by unique names:

        - This concept is shared between format, the format must map to and from an unique name that is shared between all format.
        - To future proof things, we use the official full international name of the character without space, capitalizing the first character of each part of the name.
        - So HonokaKosaka, ChikaTakami, AyumuUehara, KanonShibuya, and so on.
        - This convention should extends to other characters as well.
        - However it's all customizable.

    - There are some special characters that appeared in official daily theater:
        - There is a special "unknown" character, we use `???` for this case.
        - There is a special "later on" character, we use `LaterOn` for this case.
        - There is a special "sound effect" character, we use `SoundEffect` for this case.
        - There is a special "everyone" character, we use `Everyone` for this case.
        - Hanpen the cat show up as a dialog character in one daily theater, we use `Hanpen` for this.

- Text - The text spoken:

    - We assume this is "just plain text"
    - Although similar to the other places in game with this kind of text rendering, color, inline images, or even buttons can probably appear.

### Builtin formats
There are the following builtin formats:

- Canonical network format: 

    - Name in code: `canonical`
    - This is a field in the response of `/dailyTheater/fetchDailyTheater`
    - Contain infos related to the daily theater itself:
        - `daily_theater_id`
        - `title`
        - `detail_text`
        - `year`
        - `month`
        - `day`
    - Is a json object containing the above fields.

- Elichika loading format:

    - Name in code: `elichika`
    - This is the same as the canonical network format, except there's another `language` field and the `title` and `detail_text` field dont't have `dot_under_text` wrapper.
    - Used by [elichika](https://github.com/arina999999997/elichika) to load daily theater into database.

- triangle's logging format:

    - Name in code: `triangle`
    - The format in the csv triangle provided.
    - Contain multiple fields, but the `title` and `detail_text` is mixed together into 1 field.
    - Maybe it was the format before some update(?)

- Custom maker format:

    - Name in code: `custom_maker`
    - This is the export/import format from https://twy.name/LLAS/mainichi/

- Zura's custom maker fork format:

    - Name in code: `zura_maker`
    - This is the export/import format from https://zura3395.github.io/daily-theatre/

- Published doc format:

    - Name in code: `doc`
    - This is the doc format containing both translation and original Japanese published at https://twitter.com/SIFAStheatre/status/1675169461740146688
    - Assume the docs is saved using the download .txt button from Google Drive

### Custom formats
To add a format, implement the respective `parse` and `serialize` function, and then register it. Just check the existing implementations for examples.

Notes:

- `parse` or `serialize` can be missing, in which case can you can only `serialize` or `parse` from the format.
- For the existing formats, only some formats will have `serialize`, otherwise they would only have `parse`.

## Data details
### Original Japanese text
For the original Japanese text:

- We only have network record by triangle from 2020-02-04 to 2021-06-23. This data is considered the most correct and will be used when possible.
- Other than that, we use the transcription provided by SIFAStheatre and Idol Story. This is reproduced by human and can have mistakes:

    - Some phrase might by misinputted.
    - Some Kanji might be replaced with a similar looking Kanji.
    - Some English character might be used in place of the Japanese one.
    - The errors in the name of the character and the title can be programmatically detected, and corrected.
    - However, errors inside the text itself is harder to detect as we don't have access to any good reference point.
    - If you have such data and is willing to share, please reach out.

### Translated English text
For the translated English text:

- We use translation provided by [SIFAStheatre](https://twitter.com/SIFAStheatre) and [Idol Story](https://twitter.com/idoldotst).
- This is the format parsed directly from their published documents, and might potentially contain error compared to what they posted on social media.
- The posts on social media are proofread and are considered more correct.
- Also there can be parsing error.
- So, if you spot a mismatch, please reach out.

### Other translated texts
There is no translated texts to other languages for now. If you know of such resources and want them added, please reach out.

## Credit
Special thanks to the community for perserving and translating the daily theater. In no particular order:

- Special thanks to [SIFAStheatre](https://twitter.com/SIFAStheatre) and [Idol Story](https://twitter.com/idoldotst) for translating and for the original transcript.
- Special thanks to [TWY](https://twy.name/LLAS/mainichi) and [zura3395](https://zura3395.github.io/daily-theatre) for the custom daily theater makers.
- Special thanks to [triangle](https://triangular.dev) for the network log of daily theaters.

## Disclaimer
The raw data and possibly data that appear as literals in the code are presented good faith with implicit or explicit permission from the respective sources. If you happen to be the source of some of the data and want them removed, please reach out to resolve it.
