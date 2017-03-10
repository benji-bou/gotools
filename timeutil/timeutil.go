package timeutil

import (
	"encoding/json"
	"errors"
	"fmt"
	// "log"
	"strconv"
	"time"
)

type UnixTime struct {
	time.Time
}

func (t *UnixTime) UnmarshalJSON(b []byte) error {
	// you can now parse b as thoroughly as you want
	var dateUnix int64
	errUnix := json.Unmarshal(b, &dateUnix)
	if errUnix == nil {
		*t = UnixTime{time.Unix(dateUnix, 0)}
		return nil
	}
	tmpTime := &time.Time{}
	err := json.Unmarshal(b, tmpTime)
	if err == nil {
		*t = UnixTime{*tmpTime}
		return nil
	}
	return errors.New(fmt.Sprint(err, errUnix))
}

func (t *UnixTime) UnmarshalParam(param string) error {
	tmpTime, err := time.Parse(time.RFC3339, param)
	if err == nil {
		*t = UnixTime{tmpTime}
		return nil
	}
	baseInt, errInt := strconv.ParseInt(param, 10, 64)
	if errInt == nil {
		*t = UnixTime{time.Unix(baseInt, 0)}
		return nil
	}
	return errors.New(fmt.Sprint(err, " --> ", errInt))
}
