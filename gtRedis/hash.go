package gtRedis

import (
	"github.com/gomodule/redigo/redis"
)

/*
* 设置hash值
* @value 可以传入map[string]interface{}或者结构体对象
* @note 传入map[string]interface{}时，只会更新(或者覆盖)map包含的字段，其他字段的值不会被改变;也可以通过指定结构体字段tag标记"omitempty",来忽略空值的字段;
**/
func HashSet(conn redis.Conn, key string, value interface{}) error {
	args := redis.Args{}
	args = args.Add(key)
	args = args.AddFlat(value)
	_, err := conn.Do("HMSET", args...)
	return formatError(err, "HMSET failed. key: %s value: %+v", key, value)
}

/*
* 获取hash值
* @dst 必须传入结构体对象指针
* @note 如果key不存在,报错ErrKeyNotExist;如果key存在,redis hash中有的字段,会对应更新到dst结构体的相应字段,hash中不存在的字段，dst中对应字段值保持不变;
**/
func HashGet(conn redis.Conn, key string, dst interface{}) error {
	arrReply, err := redis.Values(conn.Do("HGETALL", key))
	if err != nil {
		return formatError(err, "HGETALL failed. key: %s dst: %+v", key, dst)
	}
	if len(arrReply) <= 0 {
		return ErrKeyNotExist
	}
	err = redis.ScanStruct(arrReply, dst)
	if err != nil {
		return formatError(err, "redis.ScanStruct() failed. key: %s arrReply: %+v dst: %+v", key, arrReply, dst)
	}
	return nil
}

/*
* 获取hash多个指定字段的值
* @fields 只有此处指定的字段，才会获取其值
* @dst 必须传入结构体对象指针;获取成功后，fields指定的字段，结构体对应的值会被修改，其他字段值保持不变;
* @note 如果key不存在,报错ErrKeyNotExist;如果key存在,redis hash中有的字段,会对应更新到dst结构体的相应字段,hash中不存在的字段，dst中对应字段值保持不变;
**/
func HashGetFields(conn redis.Conn, key string, fields []string, dst interface{}) error {
	if len(fields) <= 0 {
		return nil
	}

	//fields去重
	mapTemp := make(map[string]int)
	tempFields := make([]string, 0)
	for _, v := range fields {
		if _, ok := mapTemp[v]; !ok {
			tempFields = append(tempFields, v)
			mapTemp[v] = 1
		}
	}
	fields = tempFields

	args := redis.Args{}
	args = args.Add(key)
	for _, v := range fields {
		args = args.Add(v)
	}
	_ = conn.Send("MULTI")
	_ = conn.Send("EXISTS", key)
	_ = conn.Send("HMGET", args...)
	arrValues, err := redis.Values(conn.Do("EXEC"))
	if err != nil {
		return formatError(err, "EXEC failed. key: %s fields: %+v", key, fields)
	}

	if arrValues[0].(int64) == 0 {
		return ErrKeyNotExist
	}

	//填充可用于ScanStruct函数的输入参数
	arrReply := make([]interface{}, 0)
	for i, field := range fields {
		arrReply = append(arrReply, []byte(field))
		arrReply = append(arrReply, arrValues[1].([]interface{})[i])
	}

	err = redis.ScanStruct(arrReply, dst)
	if err != nil {
		return formatError(err, "redis.ScanStruct() failed. key: %s arrReply: %+v dst: %+v", key, arrReply, dst)
	}

	return nil
}

/*
* 获取hash指定字段的值
* @field 要获取值的字段名
* @dst 必须传入结构体对象指针;获取成功后，field指定的字段，结构体对应的值会被修改，其他字段值保持不变;
* @note 如果key不存在,报错ErrKeyNotExist;如果key存在,redis hash中有此字段,会对应更新到dst结构体的相应字段,hash中不存在此字段，dst中此字段值保持不变;
**/
func HashGetField(conn redis.Conn, key string, field string, dst interface{}) error {
	return HashGetFields(conn, key, []string{field}, dst)
}
