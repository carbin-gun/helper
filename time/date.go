package time

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type Date time.Time
type DateTime time.Time
type MilliSecondTime time.Time

const (
	datePattern     = "2006-01-02"
	dateTimePattern = "2006-01-02 15:04"
)

//MarshalJson json serialize
func (obj Date) MarshalJson() ([]byte, error) {
	myDate := time.Time(obj)
	src := myDate.Format(datePattern)
	return []byte(src), nil
}

//Value database store
func (obj Date) Value() (driver.Value, error) {
	myDate := time.Time(obj)
	src := myDate.Format(datePattern)
	return []byte(src), nil
}

func (obj DateTime) MarshalJson() ([]byte, error) {
	myDate := time.Time(obj)
	src := myDate.Format(dateTimePattern)
	return []byte(src), nil
}

//Value database store
func (obj DateTime) Value() (driver.Value, error) {
	myDate := time.Time(obj)
	src := myDate.Format(dateTimePattern)
	return []byte(src), nil
}

func (obj MilliSecondTime) MilliSeconds() int64 {
	myDate := time.Time(obj)
	return myDate.Unix() * 1000
}
func (obj MilliSecondTime) MarshalJson() ([]byte, error) {
	milliSeconds := obj.MilliSeconds()
	return json.Marshal(milliSeconds)
}

//Value database store
func (obj MilliSecondTime) Value() (driver.Value, error) {
	milliSeconds := obj.MilliSeconds()
	return json.Marshal(milliSeconds)
}
