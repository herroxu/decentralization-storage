package utils

import "strings"


func GetFileName(filePath string)(name string){
	dataArr := strings.Split(filePath, "/")
	length := len(dataArr)
	if (length <= 0){
		return
	}
	name = dataArr[length-1]
	return
}



