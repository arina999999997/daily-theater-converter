package canonical

import (
	"github.com/arina999999997/daily_theater_converter/common"
)

const formatName = "canonical"

func init() {
	common.AddFormat(formatName, parse, serialize)
}
