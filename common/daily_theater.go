package common

type DailyTheater struct {
	Language       string // some format store their own language, some format have to get them from context
	DailyTheaterId int32
	Title          string
	Lines          []Line
	Year           int32
	Month          int32
	Day            int32
}
