package gormate

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type TimeFormat time.Time

func (tf TimeFormat) MarshalJSON() ([]byte, error) {
	str := time.Time(tf).Format(time.RFC3339)
	return []byte("\"" + str + "\""), nil
}

func (tf TimeFormat) Time() time.Time {
	return time.Time(tf)
}

func (tf *TimeFormat) FromTime(t time.Time) {
	*tf = TimeFormat(t)
}

// Value to driver.Value.
func (tf TimeFormat) Value() (driver.Value, error) {
	return time.Time(tf), nil
}

// Scan value of time.Time
func (tf *TimeFormat) Scan(v interface{}) error {
	var t time.Time
	switch s := v.(type) {
	case time.Time:
		t = s
	case []byte:
		var err error
		t, err = time.Parse("2006-01-02 15:04:05", string(s))
		if err != nil {
			return err
		}
	case string:
		var err error
		t, err = time.Parse("2006-01-02 15:04:05", s)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("date: Unsupport scanning type %T", v)
	}
	tf.FromTime(t)
	return nil
}
