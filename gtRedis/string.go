package gtRedis

import (
	"github.com/gomodule/redigo/redis"
)

/*
* 设置string值
* @expire 过期时间,单位: 秒; 永不过期填负数即可(一般填-1)
**/
func StringSet(conn redis.Conn, key string, value string, expire int) error {
	args := redis.Args{}
	args = args.Add(key)
	args = args.Add(value)
	if expire >= 0 {
		args = args.Add("EX")
		args = args.Add(expire)
	}
	_, err := conn.Do("SET", args...)
	return formatError(err, "SET failed. key: %s value: %s expire: %d", key, value, expire)
}

/*
* 获取string值
* @note 获取不存在的key，会报错ErrKeyNotExist
**/
func StringGet(conn redis.Conn, key string) (string, error) {
	value, err := redis.String(conn.Do("GET", key))
	if err == redis.ErrNil {
		return "", ErrKeyNotExist
	}
	return value, formatError(err, "GET failed. key: %s", key)
}
