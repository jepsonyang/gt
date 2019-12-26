package gt

/*
* @purpose 定义用于JSON解析的基本类型
* @note 前缀V表示valid
**/

import (
	"encoding/json"
	"time"
)

//VInt
type VInt struct {
	Int   int
	Valid bool
}

func (m VInt) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.Int)
}

func (m *VInt) UnmarshalJSON(data []byte) error {
	err := json.Unmarshal(data, &m.Int)
	m.Valid = err == nil
	return err
}

//VInt64
type VInt64 struct {
	Int64 int64
	Valid bool
}

func (m VInt64) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.Int64)
}

func (m *VInt64) UnmarshalJSON(data []byte) error {
	err := json.Unmarshal(data, &m.Int64)
	m.Valid = err == nil
	return err
}

//VFloat32
type VFloat struct {
	Float32 float32
	Valid   bool
}

func (m VFloat) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.Float32)
}

func (m *VFloat) UnmarshalJSON(data []byte) error {
	err := json.Unmarshal(data, &m.Float32)
	m.Valid = err == nil
	return err
}

//VFloat64
type VFloat64 struct {
	Float64 float64
	Valid   bool
}

func (m VFloat64) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.Float64)
}

func (m *VFloat64) UnmarshalJSON(data []byte) error {
	err := json.Unmarshal(data, &m.Float64)
	m.Valid = err == nil
	return err
}

//VString
type VString struct {
	String string
	Valid  bool
}

func (m VString) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.String)
}

func (m *VString) UnmarshalJSON(data []byte) error {
	err := json.Unmarshal(data, &m.String)
	m.Valid = err == nil
	return err
}

//VTime
type VTime struct {
	Time  time.Time
	Valid bool
}

func (m VTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.Time)
}

func (m *VTime) UnmarshalJSON(data []byte) error {
	err := json.Unmarshal(data, &m.Time)
	m.Valid = err == nil
	return err
}
