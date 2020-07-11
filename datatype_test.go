package gormate_test

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/ipfans/gormate"
	"github.com/stretchr/testify/assert"
)

func TestTimeFormat(t *testing.T) {
	now := time.Now()
	m := map[string]gormate.TimeFormat{
		"time": gormate.TimeFormat(now),
	}
	b, err := json.Marshal(m)
	assert.Nil(t, err)
	assert.Equal(t,
		fmt.Sprintf(`{"time":"%s"}`, now.Format(time.RFC3339)),
		string(b),
	)

	newNow := time.Now()
	assert.NotEqual(t, now, newNow)
	tf := gormate.TimeFormat(now)
	assert.Equal(t, now, tf.Time())
	tf.FromTime(newNow)
	assert.Equal(t, newNow, tf.Time())

	var v driver.Value
	v, err = tf.Value()
	assert.Nil(t, err)
	if !assert.Equal(t, newNow, v) {
		return
	}

	// Postgresql Scanner
	err = tf.Scan(now)
	assert.Nil(t, err)
	assert.Equal(t, now, tf.Time())
	// MySQL Scanner
	err = tf.Scan([]byte(newNow.UTC().Format("2006-01-02 15:04:05")))
	assert.Nil(t, err)
	if !assert.Equal(t,
		newNow.Format("2006-01-02 15:04:05"),
		tf.Time().Local().Format("2006-01-02 15:04:05")) {
		return
	}
	// SQLite Scanner
	err = tf.Scan(now.UTC().Format("2006-01-02 15:04:05"))
	assert.Nil(t, err)
	if !assert.Equal(t,
		now.Format("2006-01-02 15:04:05"),
		tf.Time().Local().Format("2006-01-02 15:04:05")) {
		return
	}
}
