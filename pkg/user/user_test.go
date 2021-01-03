package user

import (
	"log"
	"testing"
	"time"
)

func TestCreateUser(t *testing.T) {

	name := "test-user"
	u := CreateUser(name)
	if u.Name() != name {
		t.Errorf("got: %v\nwant: %v", u.Name(), name)
	}
}

func TestGetWeight(t *testing.T) {
	name := "test-user"
	u := CreateUser(name)

	var testWeight float32 = 62.3
	date := time.Date(2020, 11, 22, 23, 59, 59, 0, time.Local)
	err := u.AddWeight(testWeight, date)
	if err != nil {
		t.Errorf("Failed to add weight")
	}

	weight, err := u.GetWeight(date)
	if err != nil {
		t.Errorf("Failed to get weight")
	}

	if weight != testWeight {
		t.Errorf("got: %f\nwant: %f", weight, testWeight)
	}

	testNotExistsDate := time.Date(2020, 11, 29, 23, 59, 59, 0, time.Local)
	_, err2 := u.GetWeight(testNotExistsDate)
	if err2 == nil {
		t.Errorf("Could not get Error")
	}
}

func TestGetWeightFromRange(t *testing.T) {
	name := "test-user"
	u := CreateUser(name)

	startDay := 22
	endDay := 26

	start := time.Date(2020, 11, startDay, 22, 21, 00, 0, time.Local)
	end := time.Date(2020, 11, endDay, 12, 31, 00, 0, time.Local)

	testWeight := make(map[time.Time]float32)
	testDate := make([]time.Time, endDay-startDay+1)

	for i := 0; i < endDay-startDay+1; i++ {
		testDate[i] = time.Date(2020, 11, startDay+i, 00, 00, 00, 0, time.Local)
	}

	testWeight[testDate[0]] = 52.1
	testWeight[testDate[1]] = 52.1
	testWeight[testDate[2]] = 53.2
	testWeight[testDate[3]] = 55.1
	testWeight[testDate[4]] = 54.3

	for j := 0; j < endDay-startDay+1; j++ {
		u.AddWeight(testWeight[testDate[j]], testDate[j])
	}

	result, err := u.GetWeightFromRange(start, end)
	if err != nil {
		t.Errorf("Failed to get weight for week")
	}
	log.Println(result)

	for _, d := range testDate {
		if testWeight[d] != result[d] {
			t.Errorf("got: %f\nwant: %f", result[d], testWeight[d])
		}
	}

	_, err2 := u.GetWeightFromRange(end, start)
	if err2 == nil {
		t.Errorf("Should get error")
	}

}
