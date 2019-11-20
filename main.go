package main

import (
	"fmt"

	"github.com/soveran/redisurl"
)

func main() {
	conn, err := redisurl.ConnectToURL("telnet://redis:6379")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	userName := "testUser"
	userKey := "online." + userName

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
}
