package gtRedis

import (
	"github.com/gomodule/redigo/redis"
)

/*
* 在list的左侧添加单个元素
* @return 添加元素后，list的长度
**/
func ListLPush(conn redis.Conn, key string, value string) (int64, error) {
	return ListLPushArray(conn, key, []string{value})
}

/*
* 在list的左侧添加单个元素
* @return 添加元素后，list的长度
**/
func ListLPushArray(conn redis.Conn, key string, values []string) (int64, error) {
	if len(values) <= 0 {
		return 0, nil
	}
	args := redis.Args{}
	args = args.Add(key)
	for _, v := range values {
		if len(v) <= 0 {
			continue
		}
		args = args.Add(v)
	}
	count, err := redis.Int64(conn.Do("LPUSH", args...))
	return count, formatError(err, "LPUSH failed. key: %s values: %+v", key, values)
}

/*
* 在list的右侧添加单个元素
* @return 添加元素后，list的长度
**/
func ListRPush(conn redis.Conn, key string, value string) (int64, error) {
	return ListRPushArray(conn, key, []string{value})
}

/*
* 在list的右侧添加多个元素
* @return 添加元素后，list的长度
**/
func ListRPushArray(conn redis.Conn, key string, values []string) (int64, error) {
	if len(values) <= 0 {
		return 0, nil
	}
	args := redis.Args{}
	args = args.Add(key)
	for _, v := range values {
		if len(v) <= 0 {
			continue
		}
		args = args.Add(v)
	}
	count, err := redis.Int64(conn.Do("RPUSH", args...))
	return count, formatError(err, "RPUSH failed. key: %s values: %+v", key, values)
}

/*
* 从list左侧弹出一个元素
* @return 弹出的元素值; 如果list不存在，报错ErrKeyNotExist;
**/
func ListLPop(conn redis.Conn, key string) (string, error) {
	args := redis.Args{}
	args = args.Add(key)
	value, err := redis.String(conn.Do("LPOP", args...))
	if err == redis.ErrNil {
		return "", ErrKeyNotExist
	}
	return value, formatError(err, "LPOP failed. key: %s", key)
}

/*
* 从list右侧弹出一个元素
* @return 弹出的元素值; 如果list不存在，报错ErrKeyNotExist;
**/
func ListRPop(conn redis.Conn, key string) (string, error) {
	args := redis.Args{}
	args = args.Add(key)
	value, err := redis.String(conn.Do("RPOP", args...))
	if err == redis.ErrNil {
		return "", ErrKeyNotExist
	}
	return value, formatError(err, "RPOP failed. key: %s", key)
}

/*
* 阻塞模式下，从list左侧弹出一个元素
* @timeout 等待超时时间(单位: 秒)，0表示无限等待;
* @return 弹出的元素值; 如果list不存在，会阻塞等待指定时间;如果等待超时，返回redis.ErrNil;
**/
func ListBLPop(conn redis.Conn, key string, timeout int) (string, error) {
	args := redis.Args{}
	args = args.Add(key)
	args = args.Add(timeout)
	mapResult, err := redis.StringMap(conn.Do("BLPOP", args...))
	if err != nil {
		if err == redis.ErrNil {
			return "", err
		}
		return "", formatError(err, "BLPOP failed. key: %s timeout: %d", key, timeout)
	}
	value, ok := mapResult[key]
	if !ok {
		return "", formatError(err, "BLPOP result not contain the key. key: %s timeout: %d mapResult: %+v", key, timeout, mapResult)
	}
	return value, nil
}

/*
* 阻塞模式下，从list右侧弹出一个元素
* @timeout 等待超时时间(单位: 秒)，0表示无限等待;
* @return 弹出的元素值; 如果list不存在，会阻塞等待指定时间;如果等待超时，返回redis.ErrNil;
**/
func ListBRPop(conn redis.Conn, key string, timeout int) (string, error) {
	args := redis.Args{}
	args = args.Add(key)
	args = args.Add(timeout)
	mapResult, err := redis.StringMap(conn.Do("BRPOP", args...))
	if err != nil {
		if err == redis.ErrNil {
			return "", err
		}
		return "", formatError(err, "BRPOP failed. key: %s timeout: %d", key, timeout)
	}
	value, ok := mapResult[key]
	if !ok {
		return "", formatError(err, "BRPOP result not contain the key. key: %s timeout: %d mapResult: %+v", key, timeout, mapResult)
	}
	return value, nil
}

/*
* 获取list指定区间内的元素
* @startIndex 区间开始下标(从0开始)
* @endIndex 区间结束下标(从0开始)
* @return 指定区间的元素列表; 如果list不存在或指定区间没有值，不会报错，返回列表为空;
* @note 区间下标可以使用负数，如-1表示最后一个元素，-2表示倒数第2个元素; (0, -1)表示整个list的元素;下标超出范围，不会报错;
**/
func ListRange(conn redis.Conn, key string, startIndex int64, endIndex int64) ([]string, error) {
	args := redis.Args{}
	args = args.Add(key)
	args = args.Add(startIndex)
	args = args.Add(endIndex)
	values, err := redis.Strings(conn.Do("LRANGE", args...))
	if err != nil {
		return values, formatError(err, "LRANGE failed. key: %s start: %d end: %d", key, startIndex, endIndex)
	}
	return values, err
}

/*
* 获取list元素数量
* @return 元素数量; 如果key不存在，返回0，且不报错;
**/
func ListLen(conn redis.Conn, key string) (int64, error) {
	args := redis.Args{}
	args = args.Add(key)
	count, err := redis.Int64(conn.Do("LLEN", args...))
	if err != nil {
		return count, formatError(err, "LLEN failed. key: %s", key)
	}
	return count, err
}
