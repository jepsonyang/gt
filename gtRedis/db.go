package gtRedis

import (
	"github.com/gomodule/redigo/redis"
)

/*
* 设置过期时间
* @return 设置成功返回true;当key不存在或者不能为key设置生存时间时(比如在低于2.1.3版本的Redis中你尝试更新key的生存时间),返回false;
**/
func Expire(conn redis.Conn, key string, expire int) (bool, error) {
	result, err := redis.Int(conn.Do("EXPIRE", key, expire))
	return result == 1, formatError(err, "EXPIRE failed. key: %s expire: %d", key, expire)
}

/*
* 获取key的剩余生存时间(TTL: time to live; 单位: 秒)
* @note key不存在返回-2; key存在，但未设置过期时间,返回-1;
**/
func TTL(conn redis.Conn, key string) (int, error) {
	return redis.Int(conn.Do("TTL", key))
}

/*
* 是否存在
* @return 存在返回true，不存在返回false
**/
func Exist(conn redis.Conn, key string) (bool, error) {
	result, err := redis.Int(conn.Do("EXISTS", key))
	return result == 1, formatError(err, "EXISTS failed. key: %s", key)
}

/*
* 删除单个key
* @note 删除不存在的key，不会报错
**/
func Delete(conn redis.Conn, key string) error {
	if len(key) <= 0 {
		return nil
	}
	_, err := DeleteArray(conn, []string{key})
	return err
}

/*
* 删除多个key
* @return 被删除的key个数
* @note keys中有不存在的key，不会报错
**/
func DeleteArray(conn redis.Conn, keys []string) (int, error) {
	if len(keys) <= 0 {
		return 0, nil
	}

	args := redis.Args{}
	for _, v := range keys {
		args = args.Add(v)
	}

	count, err := redis.Int(conn.Do("DEL", args...))
	return count, formatError(err, "DEL failed. keys: %+v", keys)
}

/*
* 遍历
* @cursor 遍历使用的游标，开始遍历填0
* @pattern 指定一个glob风格的模式参数,只返回和给定模式相匹配的元素,如果不需要现在，直接填""即可;
* @count 提示redis每次迭代返回多少个元素,默认为10; 注: 这只是一个提示,redis返回个数,不一定和指定值完全一样,但是，大部分情况是有效的;
* @return nextCursor 下次迭代使用的游标值,0表示遍历完成; keys 本次遍历到的key列表
**/
func Scan(conn redis.Conn, cursor int, pattern string, count int) (nextCursor int, keys []string, err error) {
	args := redis.Args{}
	args = args.Add(cursor)
	if len(pattern) > 0 {
		args = args.Add("MATCH")
		args = args.Add(pattern)
	}
	if count > 0 {
		args = args.Add("COUNT")
		args = args.Add(count)
	}

	var arrValue []interface{}
	arrValue, err = redis.Values(conn.Do("SCAN", args...))
	if err != nil {
		err = formatError(err, "SCAN failed. cursor: %d pattern: %s count: %d", cursor, pattern, count)
		return
	}

	nextCursor, _ = redis.Int(arrValue[0], nil)
	keys, _ = redis.Strings(arrValue[1], nil)

	return
}
