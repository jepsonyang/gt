package gt

import (
	"encoding/json"
	"fmt"
)

/*
* 将map[string]interface{}转成struct
* @dst 必须传入结构体对象指针
**/
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
