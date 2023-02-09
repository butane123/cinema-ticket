-- 1.参数列表
-- 1.1.电影argv
local filmKey = ARGV[1]
-- 2.脚本业务
-- 2.1.判断该key是否存在
if(redis.call('exists', filmKey) == 0) then
    -- 2.2.不存在该电影，则返回1
    return 1
end
-- 2.2.判断库存是否大于0
if(tonumber(redis.call('get', filmKey)) <= 0) then
    -- 2.2.库存不足，则返回2
    return 2
end
-- 2.3.库存减1
redis.call('incrby', filmKey, -1)
return 0
