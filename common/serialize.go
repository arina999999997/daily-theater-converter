package common

import (
	"errors"
	"fmt"
)

func Serialize(context Context, dailyTheater DailyTheater) ([]byte, error) {
	serializer, exist := serializers[context.Format]
	if !exist {
		return nil, errors.New(fmt.Sprintf(`Error: format "%s" doesn't have a serializer`, context.Format))
	}
	bytes, err := serializer(context, dailyTheater)
	return bytes, err
}
