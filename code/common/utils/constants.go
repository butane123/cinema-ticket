package utils

// 发送验证码主邮箱
const ServerEmail = ""

// 发送邮箱授权码
const EmailAuthCode = ""

// 邮箱Smtp地址
const EmailSmtpAddr = ""

// 邮箱SmtpHost
const EmailSmtpHost = ""

const VerificationCodeLength = 6

const DefaultPageSize = 8

// 当前时间戳
const BeginTimeStamp = 1675580392

// Id序列号部分的位长
const IdCountBit = 32

// 缓存的统一前缀
const CacheAdvertKey = "cache:advert:"
const CacheCommentKey = "cache:comment:"
const CacheFilmKey = "cache:film:"
const CacheOrderKey = "cache:order:"
const CachePayKey = "cache:pay:"
const CacheUserKey = "cache:user:"
const CacheStockKey = "cache:stock:"
const CacheEmailCodeKey = "cache:email:code:"

// 缓存过期
const RedisLockExpireSeconds = 10

const EmailCodeExpireSeconds = 300
