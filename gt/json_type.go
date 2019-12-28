package gt

/*
* @purpose 定义用于JSON解析的基本类型
* @note (1)前缀V表示valid; (2)MarshalJSON时，不考虑Valid字段
**/

import (
	"encoding/json"
	"time"
)

//VInt
type VInt struct {
	Data  int
	Valid bool
}

func (m VInt) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.Data)
}

func (m *VInt) UnmarshalJSON(data []byte) error {
	err := json.Unmarshal(data, &m.Data)
	m.Valid = err == nil
	return err
}

//VInt64
type VInt64 struct {
	Data  int64
	Valid bool
}

func (m VInt64) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.Data)
}

func (m *VInt64) UnmarshalJSON(data []byte) error {
	err := json.Unmarshal(data, &m.Data)
	m.Valid = err == nil
	return err
}

//VFloat32
type VFloat struct {
	Data  float32
	Valid bool
}

func (m VFloat) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.Data)
}

func (m *VFloat) UnmarshalJSON(data []byte) error {
	err := json.Unmarshal(data, &m.Data)
	m.Valid = err == nil
	return err
}

//VFloat64
type VFloat64 struct {
	Data  float64
	Valid bool
}

func (m VFloat64) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.Data)
}

func (m *VFloat64) UnmarshalJSON(data []byte) error {
	err := json.Unmarshal(data, &m.Data)
	m.Valid = err == nil
	return err
}

//VString
type VString struct {
	Data  string
	Valid bool
}

func (m VString) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.Data)
}

func (m *VString) UnmarshalJSON(data []byte) error {
	err := json.Unmarshal(data, &m.Data)
	m.Valid = err == nil
	return err
}

//VTime
type VTime struct {
	Data  time.Time
	Valid bool
}

func (m VTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.Data)
}

func (m *VTime) UnmarshalJSON(data []byte) error {
	err := json.Unmarshal(data, &m.Data)
	m.Valid = err == nil
	return err
}
