package models

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type StdTime time.Time

const timeFormat = "2006-01-02 15:04:05"

func (t StdTime) MarshalJSON() ([]byte, error) {
	//do your serializing here
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format(timeFormat))
	return []byte(stamp), nil
}

func (t *StdTime) UnmarshalJSON(data []byte) (err error) {
	local, err := time.LoadLocation("Asia/Chongqing")
	now, err := time.ParseInLocation(timeFormat, string(data), local)
	*t = StdTime(now)
	return
}

// Value insert timestamp into mysql need this function.
func (ts StdTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	var ti = time.Time(ts)
	if ti.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return ti, nil
}

// Scan valueof time.Time
func (ts *StdTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*ts = StdTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

func (t StdTime) String() string {
	return time.Time(t).Format(timeFormat)
}
