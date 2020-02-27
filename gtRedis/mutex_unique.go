package gtRedis

import "github.com/gomodule/redigo/redis"

const srcUniqueLockScript = `
local result = redis.pcall('SET', KEYS[1], ARGV[1], "NX", "EX", ARGV[2])
if result then
    return 1
else
    if redis.pcall('GET', KEYS[1]) == ARGV[1] then
		redis.pcall('EXPIRE', KEYS[1], ARGV[2])
		return 1
	else
		return 0
	end
end
`

var uniqueLockScript *redis.Script

func init() {
	uniqueLockScript = redis.NewScript(1, srcUniqueLockScript)
}

type UniMutex struct {
	key         string
	uniqueValue string
	expire      int
}

func (param *UniMutex) New(key string, uniqueValue string, expire int) {
	param.key = "gtRedisUniMutex:" + key
	param.uniqueValue = uniqueValue
	param.expire = expire
}

func (param *UniMutex) Lock(conn redis.Conn) error {
	args := redis.Args{}
	args = args.Add(param.key)
	args = args.Add(param.uniqueValue)
	args = args.Add(param.expire)
	result, err := redis.Int(uniqueLockScript.Do(conn, args...))
	if err != nil {
		return formatError(err, "do script failed. key: %s uniqueValue: %s expire: %d", param.key, param.uniqueValue, param.expire)
	}
	if result != 1 {
		return ErrLockExist
	}
	return nil
}
