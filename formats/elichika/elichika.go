package elichika

import (
	"github.com/arina999999997/daily_theater_converter/common"
)

const formatName = "elichika"

func init() {
	common.AddFormat(formatName, parse, serialize)
}
