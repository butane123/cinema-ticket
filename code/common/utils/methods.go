package utils

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/zeromicro/go-zero/core/stores/redis"

	"github.com/golang-jwt/jwt/v4"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/scrypt"
)

func GenerateVerificationCode() string {
	str, verificationCode := "0123456789", ""
	rand.Seed(time.Now().Unix())
	for i := 0; i < VerificationCodeLength; i++ {
		verificationCode += fmt.Sprintf("%c", str[rand.Intn(10)])
	}
	return verificationCode
}

func GenerateUUID() string {
	return uuid.NewV4().String()
}

func GenerateJwtToken(secreKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secreKey))
}

func PasswordEncrypt(salt, password string) string {
	dk, _ := scrypt.Key([]byte(password), []byte(salt), 32768, 8, 1, 32)
	return fmt.Sprintf("%x", string(dk))
}

func GenerateNewId(redis *redis.Redis, keyPrefix string) int64 {
	//获取当前时间戳
	nowStamp := time.Now().Unix() - BeginTimeStamp
	//调用lua脚本，获取当天累计序列号
	nowDate := time.Now().Format("2006:01:02")
	newKeyString := "cache:icr:" + keyPrefix + ":" + nowDate
	//L := lua.NewState()
	//defer L.Close()
	//L.SetGlobal("getKeyString", L.NewFunction(func(L *lua.LState) int {
	//	L.Push(lua.LString(newKeyString))
	//	return 1
	//}))
	//if err := L.DoFile("common/scripts/generateIncrCount.lua"); err != nil {
	//	panic(err)
	//}
	//res := L.Get(1)
	//count, err := strconv.ParseInt(res.String(), 10, 64)
	//if err != nil {
	//	fmt.Println("调用lua脚本错误！")
	//	return 0
	//}
	count, err := redis.Incr(newKeyString)
	if err != nil {
		fmt.Println("生成id错误！")
		return 0
	}
	//拼接结果
	return nowStamp<<IdCountBit | count
}
