package main

import "time"

// DB operate databases IO.
type DB interface {
	ReadUserLog(string) (float32, time.Time, error)
	WriteUserLog(string, float32, time.Time) error
	ReadUserLogFromRange(string, time.Time, time.Time) (map[time.Time]float32, error)
	WriteUserLogRange(string, map[time.Time]float32) (float32, time.Time, error)
}

// func ReadUserLog(id string) (float32, time.Time, error) {

//}
