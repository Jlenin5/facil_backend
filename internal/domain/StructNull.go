package domain

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// NullTime sin dependencia de database/sql
type NullTime struct {
	Time  *time.Time `json:"time,omitempty"`
	Valid bool       `json:"valid"`
}

// MarshalJSON para personalizar la salida de JSON
func (nt NullTime) MarshalJSON() ([]byte, error) {
	if !nt.Valid || nt.Time == nil {
		return json.Marshal(nil)
	}
	return json.Marshal(nt.Time)
}

// UnmarshalJSON para deserializar NullTime desde JSON
func (nt *NullTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		nt.Valid = false
		nt.Time = nil
		return nil
	}
	var t time.Time
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	nt.Time = &t
	nt.Valid = true
	return nil
}

func (nt *NullTime) Scan(value interface{}) error {
	if value == nil {
		nt.Valid = false
		nt.Time = nil
		return nil
	}

	switch v := value.(type) {
	case time.Time:
		nt.Time = &v
		nt.Valid = true
	default:
		return fmt.Errorf("tipo no soportado para NullTime: %T", value)
	}
	return nil
}

func (nt NullTime) Value() (driver.Value, error) {
	if !nt.Valid || nt.Time == nil {
		return nil, nil
	}
	return *nt.Time, nil
}

// NullString sin dependencia de database/sql
type NullString struct {
	String *string `json:"string,omitempty"`
	Valid  bool    `json:"valid"`
}

// MarshalJSON para NullString
func (ns NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid || ns.String == nil {
		return json.Marshal(nil)
	}
	return json.Marshal(ns.String)
}

// UnmarshalJSON para NullString
func (ns *NullString) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		ns.Valid = false
		ns.String = nil
		return nil
	}
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	ns.String = &s
	ns.Valid = true
	return nil
}

func (ns *NullString) Scan(value interface{}) error {
	if value == nil {
		ns.Valid = false
		ns.String = nil
		return nil
	}

	switch v := value.(type) {
	case string:
		ns.String = &v
		ns.Valid = true
	case []uint8:
		str := string(v)
		ns.String = &str
		ns.Valid = true
	default:
		return fmt.Errorf("tipo no soportado para NullString: %T", value)
	}
	return nil
}

// Implementación de driver.Valuer para que sqlx lo pueda escribir en la BD
func (ns NullString) Value() (driver.Value, error) {
	if !ns.Valid || ns.String == nil {
		return nil, nil
	}
	return *ns.String, nil
}


// NullFloat sin dependencia de database/sql
type NullFloat struct {
	Float *float64 `json:"float,omitempty"`
	Valid bool     `json:"valid"`
}

// MarshalJSON para NullFloat
func (nf NullFloat) MarshalJSON() ([]byte, error) {
	if !nf.Valid || nf.Float == nil {
		return json.Marshal(nil)
	}
	return json.Marshal(nf.Float)
}

// UnmarshalJSON para NullFloat
func (nf *NullFloat) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		nf.Valid = false
		nf.Float = nil
		return nil
	}
	var f float64
	if err := json.Unmarshal(data, &f); err != nil {
		return err
	}
	nf.Float = &f
	nf.Valid = true
	return nil
}

// Scan para asignar valores desde la base de datos
func (nf *NullFloat) Scan(value interface{}) error {
	if value == nil {
		nf.Valid = false
		nf.Float = nil
		return nil
	}

	switch v := value.(type) {
	case float64:
		nf.Float = &v
		nf.Valid = true
	case []uint8:
		str := string(v)
		f, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return err
		}
		nf.Float = &f
		nf.Valid = true
	case string:
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return err
		}
		nf.Float = &f
		nf.Valid = true
	default:
		return fmt.Errorf("tipo no soportado para NullFloat: %T", value)
	}
	return nil
}

// Implementación de driver.Valuer para que sqlx lo pueda escribir en la BD
func (nf NullFloat) Value() (driver.Value, error) {
	if !nf.Valid || nf.Float == nil {
		return nil, nil
	}
	return *nf.Float, nil
}

// NullInt sin dependencia de database/sql
type NullInt struct {
	Int   *int64 `json:"int,omitempty"`
	Valid bool   `json:"valid"`
}

// MarshalJSON para NullInt
func (ni NullInt) MarshalJSON() ([]byte, error) {
	if !ni.Valid || ni.Int == nil {
		return json.Marshal(nil)
	}
	return json.Marshal(ni.Int)
}

// UnmarshalJSON para NullInt
func (ni *NullInt) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		ni.Valid = false
		ni.Int = nil
		return nil
	}
	var i int64
	if err := json.Unmarshal(data, &i); err != nil {
		return err
	}
	ni.Int = &i
	ni.Valid = true
	return nil
}

// Scan para asignar valores desde la base de datos
func (ni *NullInt) Scan(value interface{}) error {
	if value == nil {
		ni.Valid = false
		ni.Int = nil
		return nil
	}

	switch v := value.(type) {
	case int64:
		ni.Int = &v
		ni.Valid = true
	case []uint8:
		str := string(v)
		i, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return err
		}
		ni.Int = &i
		ni.Valid = true
	case string:
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return err
		}
		ni.Int = &i
		ni.Valid = true
	default:
		return fmt.Errorf("tipo no soportado para NullInt: %T", value)
	}
	return nil
}

// Implementación de driver.Valuer para que sqlx lo pueda escribir en la BD
func (ni NullInt) Value() (driver.Value, error) {
	if !ni.Valid || ni.Int == nil {
		return nil, nil
	}
	return *ni.Int, nil
}