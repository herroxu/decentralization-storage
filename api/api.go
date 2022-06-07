package api

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"

	"github.com/kingdemoking/decentralization-storage/settings"
	"github.com/kingdemoking/decentralization-storage/structs"
)

const (
	loginUrl = settings.BaseURL + "/v1/login-by-email"
	apiKeyUrl = settings.BaseURL + "/v1/apply-key"
	uploadUrl = settings.BaseURL + "/v1/upload-to-arweave-or-ipfs"
)

func GetToken(email, pwd string)(response structs.LoginRes) {
	v := make(url.Values)
	v.Add("email", email)
	v.Add("pwd", pwd)

	res, _ := http.PostForm(loginUrl, v)
	rawData,_ := ioutil.ReadAll(res.Body)
	json.Unmarshal(rawData, &response)
	return response
}


func GetAPIKey(token string)(response structs.ApiKeyRes) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", apiKeyUrl, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("token", token)
	res, err := client.Do(req)
	rawData,_ := ioutil.ReadAll(res.Body)
	json.Unmarshal(rawData, &response)
	return response
}


func Upload(apiKey,to,filePath string)(response structs.UploadRes) {
	client := &http.Client{
	}
	
	buf := new(bytes.Buffer)
   writer := multipart.NewWriter(buf)
	formFile, _ := writer.CreateFormFile("file", "girl1.png")
	// formFile.Write(fileBytes)

	// 从文件读取数据，写入表单
	srcFile, err := os.Open(filePath)
	if err != nil {
		 log.Fatalf("%Open source file failed: s\n", err)
	}
	defer srcFile.Close()

	_, err = io.Copy(formFile, srcFile)
	if err != nil {
		 log.Fatalf("Write to form file falied: %s\n", err)
	} 

	contentType := writer.FormDataContentType()
	writer.Close() 
	req, err := http.NewRequest("POST", uploadUrl, buf)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", contentType)
	req.Header.Set("to", to)
	req.Header.Set("AuthKey", apiKey)

	res, err := client.Do(req)
	rawData,_ := ioutil.ReadAll(res.Body)
	json.Unmarshal(rawData, &response)

	return response
}























