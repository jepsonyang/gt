package gt

import (
	"encoding/json"
	"fmt"
)

func MapStr2Struct(src map[string]interface{}, dst interface{}) error {
	dataBytes, err := json.Marshal(src)
	if err != nil {
		return fmt.Errorf("Marshal failed. err: %v ", err)
	}
	err = json.Unmarshal(dataBytes, dst)
	if err != nil {
		return fmt.Errorf("Unmarshal failed. err: %v ", err)
	}
	return nil
}
