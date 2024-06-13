package common

// one file can contain multiple daily theater in some format, it's up to the format itself to split them
type Parser = func(Context, []byte) ([]DailyTheater, error)

var parsers = map[string]Parser{}
