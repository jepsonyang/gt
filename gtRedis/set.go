package gtRedis

import (
	"github.com/gomodule/redigo/redis"
)

/*
* 添加单个元素到集合
* @return 是否执行了添加操作(已经存在的元素会被忽略，即此时返回值为false)
* @note 已经存在的值，再次添加，不会报错，会被忽略;
**/
func SetAdd(conn redis.Conn, key string, value string) (bool, error) {
	if len(value) <= 0 {
		return false, nil
	}
	count, err := SetAddArray(conn, key, []string{value})
	return count == 1, err
}

/*
* 添加多个元素到集合
* @return 返回被添加到集合的新元素数量,不包括被忽略的元素;
* @note 已经存在的元素，再次添加，不会报错，会被忽略;
**/
func SetAddArray(conn redis.Conn, key string, values []string) (int, error) {
	if len(values) <= 0 {
		return 0, nil
	}
	args := redis.Args{}
	args = args.Add(key)
	for _, v := range values {
		args = args.Add(v)
	}
	count, err := redis.Int(conn.Do("SADD", args...))
	return count, formatError(err, "SADD failed. key: %s values: %+v", key, values)
}

/*
* 删除集合中的单个元素
* @return 是否执行了删除操作(不再集合中的元素会被忽略，即此时返回值为false);
* @note 删除集合中不存在的元素，不会报错，会被忽略;
**/
func SetRemove(conn redis.Conn, key string, value string) (bool, error) {
	if len(value) <= 0 {
		return false, nil
	}
	count, err := SetRemoveArray(conn, key, []string{value})
	return count==1, err
}

/*
* 删除集合中的多个元素
* @return 返回被删除的元素数量,不包括被忽略的元素;
* @note 删除集合中不存在的值，不会报错，会被忽略;
**/
func SetRemoveArray(conn redis.Conn, key string, values []string) (int, error) {
	args := redis.Args{}
	args = args.Add(key)
	for _, v := range values {
		args = args.Add(v)
	}
	count, err := redis.Int(conn.Do("SREM", args...))
	return count, formatError(err, "SREM failed. key: %s values: %+v", key, values)
}

/*
* 判断member是否为集合key的成员
* @return 如果member元素是集合的成员,返回true;如果member不是集合的成员,或key不存在,返回false
**/
func SetIsMember(conn redis.Conn, key string, member string) (bool, error) {
	result, err := redis.Int(conn.Do("SISMEMBER", key, member))
	if err != nil {
		return false, formatError(err, "SISMEMBER failed. key: %s member: %s", key, member)
	}
	return result == 1, nil
}

/*
* 获取集合中的所有成员
**/
func SetMembers(conn redis.Conn, key string) ([]string, error) {
	members, err := redis.Strings(conn.Do("SMEMBERS", key))
	return members, formatError(err, "SMEMBERS failed. key: %s", key)
}

/*
* 获取集合的成员总数
**/
func SetMemberCount(conn redis.Conn, key string) (int, error) {
	count, err := redis.Int(conn.Do("SCARD", key))
	return count, formatError(err, "SCARD failed. key: %s", key)
}
