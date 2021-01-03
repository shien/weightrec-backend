package user

import (
	"time"

	"github.com/shien/weightrec-backend/pkg/db"
	"github.com/shien/weightrec-backend/pkg/dbcfg"
)

type UserRepo struct {
	DB *db.DB
}

func Init() *UserRepo {

	conf := dbcfg.GetDBConfig()
	dbconf := conf.GetDBConnStr()

	d, err := db.NewClient(dbconf)
	ur := &UserRepo{
		DB: d,
	}
	if err != nil {
		panic(err)
	}

	return ur
}

// Close close DB
func (ur *UserRepo) Close() {
	ur.DB.Close()
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
