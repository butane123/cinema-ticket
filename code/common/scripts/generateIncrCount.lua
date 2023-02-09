-- 1.参数列表
-- 1.1.当天key
local dateKey = getKeyString()
-- 2.脚本业务
-- 2.1.判断该key是否存在 get dateKey
if(redis.call('exists', dateKey) == 0) then
    -- 2.2.不存在，则设置该key为0
    redis.call('set', dateKey, 0)
end
-- 2.2.设置该key值加1 incrby dateKey 1
redis.call('incrby', dateKey, 1)
-- 2.3.获取该key值 get dateKey
return tonumber(redis.call('get', dateKey))
