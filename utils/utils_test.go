package utils

import (
	"fmt"
	"testing"
)


func Test_GetFileName(t *testing.T) {
	filepath := "/home/lucas/girl1.png"
	fmt.Println(GetFileName(filepath))
}
