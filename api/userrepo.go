package main

import (
	"time"

	"github.com/shien/weightrec-backend/pkg/config"
	"github.com/shien/weightrec-backend/pkg/db"
)

type UserRepo struct {
	DB *db.DB
}

func Init() *UserRepo {

	con := config.Get()
	dbconf := con.GetDBConnStr()

	d, err := db.NewClient(dbconf)
	ur := &UserRepo{
		DB: d,
	}
	if err != nil {
		panic(err)
	}

	return ur
}

func (ur *UserRepo) Close() {
	ur.Close()
}

func (ur *UserRepo) AddUser(userName string) error {
	err := ur.DB.InsertUser(userName)
	if err != nil {
		return err
	}
	return nil
}

func (ur *UserRepo) ReadUserWeight(userid string, recordedDate time.Time) (float32, error) {
	weight, err := ur.DB.SelectUserWeight(userid, recordedDate)
	if err != nil {
		return 0, err
	}
	return weight, nil
}
