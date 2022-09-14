package api

import (
	"fmt"
	"testing"
)

func Test_GetToken(t *testing.T) {
	res,_ := GetToken("10551911@qq.com", "1236")
	fmt.Println(res.Token)
}


func Test_GetAPIKey(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ3YW54aWFuZyIsIiIxMDU1NzgxOTExQHFxLmNvbSIsInB3ZCI6IjhEOTY5RUVGNkVDQUQzQzI5QTNBNjI5MjgwRTY4NkNGMEMzRjVENUE4NkFGRjNDQTEyMDIwQzkyM0FEQzZDOTIiLCJyYW5kb20iOiJVRjE5YTBuRSJ9.rGliLbot22uwUOJaDWcoQ7uh5fjulHAEpUWc9P1WYjQ"
	res,_ := GetAPIKey(token)
	fmt.Println(res.Key)
	fmt.Println(res.Msg)
}

func Test_Upload(t *testing.T) {
	apiKey := "7UX9Q0ic7NlmOiGQd4jkFYPtn3rVZ2EHX3TmneBRKYaG40svjf8nj8lC"
	to := "ipfs"
	filepath := "/home/lucas/girl1.png"
	res,_ := Upload(apiKey, to, filepath)
	fmt.Println(res.DecentralizedUrl)
	fmt.Println(res.Msg)
}
