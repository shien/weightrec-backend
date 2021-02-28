package user

import (
	"time"

	"github.com/shien/weightrec-backend/pkg/db"
	"github.com/shien/weightrec-backend/pkg/dbcfg"
)

// UserRepo has user data
type UserRepo struct {
	DB *db.DB
}

// Init initialize user repository
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

func (ur *UserRepo) AddUser(mailAddress string) error {
	err := ur.DB.InsertUser(mailAddress)
	if err != nil {
		return err
	}
	return nil
}

// IsUserExists exists user had mail address
func (ur *UserRepo) IsUserExists(mailAddress string) (bool, error) {
	result, err := ur.DB.SelectUser(mailAddress)
	if err != nil {
		return false, err
	}
	if result != mailAddress {
		return false, nil
	}
	return true, nil
}

func (ur *UserRepo) ReadUserWeight(userid string, recordedDate time.Time) (float32, error) {
	weight, err := ur.DB.SelectUserWeight(userid, recordedDate)
	if err != nil {
		return 0, err
	}
	return weight, nil
}
