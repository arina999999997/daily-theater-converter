package doc

import (
	"github.com/arina999999997/daily_theater_converter/common"
)

const formatName = "doc"

func init() {
	common.AddFormat(formatName, parse, nil)
}
