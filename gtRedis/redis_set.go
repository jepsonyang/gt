package gtRedis

import (
	"github.com/gomodule/redigo/redis"
)

/*
* 添加成员到集合
* @return 返回被添加到集合的新元素数量,不包括被忽略的元素;
**/
func SetAdd(conn redis.Conn, key string, values []string) (int, error) {
	if len(values) <= 0 {
		return 0, nil
	}
	args := redis.Args{}
	args = args.Add(key)
	for _, v := range values {
		args = args.Add(v)
	}
	return redis.Int(conn.Do("SADD", args...))
}

/*
* 删除集合中的成员
* @return 返回被删除的元素数量,不包括被忽略的元素;
**/
func SetRemove(conn redis.Conn, key string, values []string) (int, error) {
	args := redis.Args{}
	args = args.Add(key)
	for _, v := range values {
		args = args.Add(v)
	}
	return redis.Int(conn.Do("SREM", args...))
}

/*
* 判断member是否为集合key的成员
* @return 如果member元素是集合的成员,返回true;如果member不是集合的成员,或key不存在,返回false
**/
func SetIsMember(conn redis.Conn, key string, member string) (bool, error) {
	result, err := redis.Int(conn.Do("SISMEMBER", key, member))
	if err != nil {
		return false, err
	}
	return result == 1, nil
}

/*
* 获取集合中的所有成员
**/
func SetMembers(conn redis.Conn, key string) ([]string, error) {
	return redis.Strings(conn.Do("SMEMBERS", key))
}

/*
* 获取集合的成员总数
**/
func SetMemberCount(conn redis.Conn, key string) (int, error) {
	return redis.Int(conn.Do("SCARD", key))
}