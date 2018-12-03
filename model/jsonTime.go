package model

import (
	"time"
	"database/sql/driver"
	"fmt"
)

type JsonDateTime struct {
	T     time.Time
	Valid bool
}

func DateTimeNow() JsonDateTime {
	return JsonDateTime{T: time.Now(), Valid: true}
}

func (j *JsonDateTime) Scan(value interface{}) error {
	j.T, j.Valid = value.(time.Time)
	return nil
}

func (j JsonDateTime) Value() (driver.Value, error) {
	if !j.Valid {
		return nil, nil
	}
	return time.Time(j.T), nil
}

func (j JsonDateTime) MarshalJSON() ([]byte, error) {
	if j.Valid {
		var stamp = fmt.Sprintf("\"%s\"", time.Time(j.T).Format("2006-01-02 15:04:05"))
		return []byte(stamp), nil
	} else {
		return []byte(`null`), nil
	}
}

func (j *JsonDateTime) UnmarshalJSON(data []byte) (error) {
	local, err := time.ParseInLocation(`"2006-01-02 15:04:05"`, string(data), time.Local)
	j.T = local
	j.Valid = err == nil
	return err
}

type JsonDate time.Time

func (j *JsonDate) Scan(value interface{}) error {
	*j = JsonDate(value.(time.Time))
	return nil
}

func (j JsonDate) Value() (driver.Value, error) {
	return time.Time(j), nil
}

func (j JsonDate) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(j).Format("2006-01-02"))
	return []byte(stamp), nil
}

func (j *JsonDate) UnmarshalJSON(data []byte) (error) {
	local, err := time.ParseInLocation(`"2006-01-02"`, string(data), time.Local)
	*j = JsonDate(local)
	return err
}

type JsonTime time.Time

func (j *JsonTime) Scan(value interface{}) error {
	*j = JsonTime(value.(time.Time))
	return nil
}

func (j JsonTime) Value() (driver.Value, error) {
	return time.Time(j), nil
}

func (j JsonTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(j).Format("15:04:05"))
	return []byte(stamp), nil
}

func (j *JsonTime) UnmarshalJSON(data []byte) error {
	local, err := time.ParseInLocation(`"15:04:05"`, string(data), time.Local)
	*j = JsonTime(local)
	return err
}

type JsonTimeHS time.Time

func (j *JsonTimeHS) Scan(value interface{}) error {
	*j = JsonTimeHS(value.(time.Time))
	return nil
}

func (j JsonTimeHS) Value() (driver.Value, error) {
	return time.Time(j), nil
}

func (j JsonTimeHS) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(j).Format("15:04"))
	return []byte(stamp), nil
}

func (j *JsonTimeHS) UnmarshalJSON(data []byte) error {
	local, err := time.ParseInLocation(`"15:04"`, string(data), time.Local)
	*j = JsonTimeHS(local)
	return err
}
