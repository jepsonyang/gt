package gt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"runtime"
	"strconv"
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

/*
* 获取协程id
* @note go从1.4版本开始，去掉了获取协程id的接口;此方法通过获取调用堆栈，从而提取协程id，仅可用于对性能不敏感的场景
**/
func GetGID() uint64 {
	buffer := make([]byte, 64)
	buffer = buffer[:runtime.Stack(buffer, false)]
	buffer = bytes.TrimPrefix(buffer, []byte("goroutine "))
	buffer = buffer[:bytes.IndexByte(buffer, ' ')]
	gid, _ := strconv.ParseUint(string(buffer), 10, 64)
	return gid
}
