package daily_theater_converter

import (
	"github.com/arina999999997/daily_theater_converter/common"
	_ "github.com/arina999999997/daily_theater_converter/formats"

	"errors"
	"fmt"
)

// Direct converting
// - Only work for files that contain a single daily theatre, even if the actual format support multiple daily theater / files

func Convert(parseContext, serializeContext common.Context, data []byte) ([]byte, error) {
	input, err := common.Parse(parseContext, data)
	if err != nil {
		return nil, err
	}
	if len(input) != 1 {
		return nil, errors.New(fmt.Sprintf(`Unexpected amount of daily theater in input: %d`, len(input)))
	}
	output, err := common.Serialize(serializeContext, input[0])
	if err != nil {
		return nil, err
	}
	return output, nil
}
