package dto

import (
	"encoding/json"
	"time"
)

type ByIDRequest struct {
	ID string `param:"id" validate:"required"`
}

type cTime struct {
	time.Time
}

func (c *cTime) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if s == "" {
		c.Time = time.Time{}
		return nil
	}
	var format = "2006-01-02"
	t, err := time.Parse(format, s)
	if err != nil {
		return err
	}
	c.Time = t
	return nil
}

func (c *cTime) ToTime() time.Time {
	return c.Time
}
