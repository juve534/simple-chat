package main

import (
	"fmt"

	"github.com/juve534/simple-chat/src"
)

func main() {
	userName := "testUser"
	userKey := "online." + userName

	res := user.SaveUser(userName, userKey)
	fmt.Println(res)

	ress := user.AddUserList(userName)
	fmt.Println(ress)
}
