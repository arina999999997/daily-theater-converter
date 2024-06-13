package common

type Serializer = func(Context, DailyTheater) ([]byte, error)

var serializers = map[string]Serializer{}
