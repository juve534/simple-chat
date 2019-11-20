// Package user チャットを行うユーザを操作する
package user

import (
	"fmt"

	"github.com/soveran/redisurl"
)

// SaveUser ユーザを保存する
func SaveUser(userName string, userKey string) interface{} {
	conn, err := redisurl.ConnectToURL("telnet://redis:6379")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	val, err := conn.Do("SET", userKey, userName, "NX", "EX", "120")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// 存在判定
	if val == nil {
		fmt.Println("既にオンラインです。")
		panic(err)
	}

	return val
}