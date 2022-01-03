package genid

import (
	"github.com/go-redis/redis/v8"
)

// lua script atom operation
var atomGetSequences = redis.NewScript(`
local key = KEYS[1]
local num = ARGV[1]
local expire = ARGV[2]
local db = ARGV[3]

redis.call("SELECT", db)

local kv = redis.call("GET", key)
local res = 0
if not kv then
	redis.call("SET", key, num)
  	redis.call("EXPIRE", key, expire)
	res = tonumber(num)
else 
	redis.call("INCRBY", key, num)
	res = kv + num
end
return res
`)
