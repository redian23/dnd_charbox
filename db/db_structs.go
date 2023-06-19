package db

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	_ "github.com/lib/pq"
)

type Table struct {
	ID   string
	Data Data
}

type Data map[string]interface{}

func (d Data) Value() (driver.Value, error) {
	return json.Marshal(d)
}

func (d *Data) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &d)
}
