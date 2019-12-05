package gtRedis

import (
	"github.com/gomodule/redigo/redis"
)

/*
* 设置hash值
* @value 可以传入map[string]interface{}或者结构体对象
* @note 传入map[string]interface{}时，只会更新(或者覆盖)map包含的字段，其他字段的值不会被改变;
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
**/
func HashGet(conn redis.Conn, key string, dst interface{}) error {
	arrReply, err := redis.Values(conn.Do("HGETALL", key))
	if err != nil {
		return formatError(err, "HGETALL failed. key: %s dst: %+v", key, dst)
	}
	err = redis.ScanStruct(arrReply, dst)
	if err != nil {
		return formatError(err, "redis.ScanStruct() failed. key: %s arrReply: %+v dst: %+v", key, arrReply, dst)
	}
	return nil
}
