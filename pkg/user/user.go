package user

import (
	"errors"
	"time"
)

// User is user that have weight and bfp
type User struct {
	name                  string
	weightData            map[recordedDate]float32
	bodyFatPercentageData map[recordedDate]float32
}

type recordedDate struct {
	date time.Time
}

// CreateUser create user in this app
func CreateUser(userName string) User {
	return User{
		name:                  userName,
		weightData:            make(map[recordedDate]float32),
		bodyFatPercentageData: make(map[recordedDate]float32),
	}
}

// Name return user name
func (u *User) Name() string {
	return u.name
}

// AddWeight add weight to map
func (u *User) AddWeight(weight float32, recordedWeightDate time.Time) error {
	rdate := recordedDate{
		date: alignDate(recordedWeightDate),
	}
	if _, ok := u.weightData[rdate]; ok {
		return errors.New("Key is exists")
	}

	u.weightData[rdate] = weight
	return nil
}

// GetWeight return user weight at recordedDate
func (u *User) GetWeight(recordedWeightDate time.Time) (float32, error) {
	rdate := recordedDate{
		date: alignDate(recordedWeightDate),
	}

	if val, ok := u.weightData[rdate]; ok {
		return val, nil
	}

	return 0.0, errors.New("Key is not exists")
}

// GetWeightFromRange return user weight for startDate to endDate
func (u *User) GetWeightFromRange(startDate time.Time, endDate time.Time) (map[time.Time]float32, error) {

	if startDate.After(endDate) {
		return nil, errors.New("start date behind end date")
	}

	s := alignDate(startDate)
	e := alignDate(endDate)

	r := make(map[time.Time]float32)

	// include end date to returning map
	for t := s; e.Add(time.Duration(24) * time.Hour).After(t); t = t.Add(time.Duration(24) * time.Hour) {
		rd := recordedDate{
			date: t,
		}
		if val, ok := u.weightData[rd]; ok {
			r[t] = val
		}
	}

	return r, nil
}

func alignDate(t time.Time) time.Time {
	return time.Date(t.Year(),
		t.Month(),
		t.Day(),
		00,
		00,
		00,
		0,
		time.Local)
}
