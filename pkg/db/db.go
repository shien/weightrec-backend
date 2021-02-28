package db

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

// WeightDB operate databases IO.
type WeightDB interface {
	ReadUserLog(string time.Time) (float32, error)
	WriteUserLog(string, float32, time.Time) error
	ReadUserLogFromRange(string, time.Time, time.Time) (map[time.Time]float32, error)
	WriteUserLogRange(string, map[time.Time]float32) (float32, time.Time, error)
	Close()
}

// DB has pointer of sql.DB
type DB struct {
	Client *sql.DB
}

// NewClient create connection database
func NewClient(dbinfo string) (*DB, error) {
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		return nil, err
	}
	return &DB{
		Client: db,
	}, nil
}

// Close close connection to database
func (d *DB) Close() error {
	return d.Client.Close()
}

// ReadUserLog read log from db
func (d *DB) SelectUserWeight(id string, recordedDate time.Time) (float32, error) {
	var weight float32 = 0.0
	err := d.Client.QueryRow("select * from users.users where id = $2 and recorded_time = $2", id, recordedDate).Scan(&weight)
	if err != nil {
		return 0.0, err
	}
	return weight, nil
}

// SelectUser read user from db
func (d *DB) SelectUser(mailAddress string) (string, error) {
	var user string = ""
	err := d.Client.QueryRow("select mail_address from users.users where mail_address = $1", mailAddress).Scan(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (d *DB) InsertUser(mailAddress string) error {
	_, err := d.Client.ExecContext(context.Background(), "insert into users.users (mail_address) values ($1)", mailAddress)
	if err != nil {
		return err
	}
	return nil
}
