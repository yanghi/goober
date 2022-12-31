package auth

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"strconv"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func Salt() string {
	salt := RandStringRunes(10)
	m5 := md5.New()

	m5.Write([]byte(salt))
	m5.Write([]byte(strconv.Itoa(int(time.Now().Unix()))))
	st := m5.Sum(nil)

	return hex.EncodeToString(st)
}

func JoinSalt(s, salt string) string {
	m5 := md5.New()

	m5.Write([]byte(s))
	m5.Write([]byte(salt))
	st := m5.Sum(nil)

	return hex.EncodeToString(st)
}

func Validate(s, salt, strMd5 string) bool {

	// m5 := md5.New()

	// m5.Write([]byte(JoinSalt(s, salt)))
	// st := m5.Sum(nil)

	return JoinSalt(s, salt) == strMd5
}
