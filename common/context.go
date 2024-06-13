package common

// used to pass parameters for formats that don't inherently store the info
// each parse should start with these data and overwrite if it has better data
type Context struct {
	Format   string
	Language string

	// these might not be present, an empty value is used if not present
	FilePath       string
	DailyTheaterId int32
	Year           int32
	Month          int32
	Day            int32
	Credit         string
}

func (c *Context) GetDefaultDailyTheater() DailyTheater {
	return DailyTheater{
		Language:       c.Language,
		DailyTheaterId: c.DailyTheaterId,
		Year:           c.Year,
		Month:          c.Month,
		Day:            c.Day,
	}
}

func NewContext(format, language string) Context {
	return Context{
		Format:   format,
		Language: language,
	}
}
