package util

import (
	"encoding/json"
	"time"
)

type CreatedAt time.Time
type UpdatedAt time.Time

func (c *CreatedAt) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(*c).Format(time.DateTime))
}

func (c *CreatedAt) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	t, err := time.Parse(time.DateTime, s)
	if err != nil {
		return err
	}
	*c = CreatedAt(t)
	return nil
}

func (u *UpdatedAt) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(*u).Format(time.DateTime))
}

func (u *UpdatedAt) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	t, err := time.Parse(time.DateTime, s)
	if err != nil {
		return err
	}
	*u = UpdatedAt(t)
	return nil
}
