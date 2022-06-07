package api

import (
	"fmt"
	"testing"
)

func Test_GetToken(t *testing.T) {
	res := GetToken("1055781911@qq.com", "123456")
	fmt.Println(res.Token)
}


func Test_GetAPIKey(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ3YW54aWFuZyIsInVzZXIiOiIxMDU1NzgxOTExQHFxLmNvbSIsInB3ZCI6IjhEOTY5RUVGNkVDQUQzQzI5QTNBNjI5MjgwRTY4NkNGMEMzRjVENUE4NkFGRjNDQTEyMDIwQzkyM0FEQzZDOTIiLCJyYW5kb20iOiJVRjE5YTBuRSJ9.rGliLbot22uwUOJaDWcoQ7uh5fjulHAEpUWc9P1WYjQ"
	res := GetAPIKey(token)
	fmt.Println(res.Key)
	fmt.Println(res.Msg)
}

func Test_Upload(t *testing.T) {
	apiKey := "GYQ6iysjaTCo0eyvPZx3W3ggva1Ev30EHztLER8Yn2whkw7E7TdG8BNe1oZxv3U2"
	to := "ipfs"
	filepath := "/home/lucas/girl1.png"
	res := Upload(apiKey, to, filepath)
	fmt.Println(res.Url)
	fmt.Println(res.Msg)
}



