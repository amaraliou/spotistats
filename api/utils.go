package api

import "strconv"

func MillisecsToSongTime(duration_ms int) string {
	minutes := (duration_ms / 1000) / 60
	seconds := duration_ms/1000 - minutes*60
	duration := strconv.Itoa(minutes) + ":" + strconv.Itoa(seconds)
	return duration
}
