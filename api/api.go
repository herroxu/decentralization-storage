package api

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"

	"github.com/kingdemoking/decentralization-storage/settings"
	"github.com/kingdemoking/decentralization-storage/structs"
	"github.com/kingdemoking/decentralization-storage/utils"
)

const (
	loginUrl = settings.BaseURL + "/v1/login-by-email"
	apiKeyUrl = settings.BaseURL + "/v1/apply-key"
	uploadUrl = settings.BaseURL + "/v1/upload-to-arweave-or-ipfs"
)

func GetToken(email, pwd string)(response structs.LoginRes, err error) {
	v := make(url.Values)
	v.Add("email", email)
	v.Add("pwd", pwd)

	res, err := http.PostForm(loginUrl, v)
	if err != nil {
		return
   	} 
	rawData,err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
   	} 
	err = json.Unmarshal(rawData, &response)
	return response,err
}


func GetAPIKey(token string)(response structs.ApiKeyRes, err error) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", apiKeyUrl, nil)
	if err != nil {
		return
	}

	req.Header.Add("token", token)
	res, err := client.Do(req)
	if err != nil {
		return
   	} 
	rawData,err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
   	} 
	err = json.Unmarshal(rawData, &response)
	return response,err
}


func Upload(apiKey, to, filePath string)(response structs.UploadRes, err error) {
	client := &http.Client{}
	fileName := utils.GetFileName(filePath)
	
	buf := new(bytes.Buffer)
    writer := multipart.NewWriter(buf)
	formFile, err := writer.CreateFormFile("file", fileName)
	// formFile.Write(fileBytes)
	// 从文件读取数据，写入表单
	srcFile, err := os.Open(filePath)
	if err != nil {
		 return
	}
	defer srcFile.Close()

	_, err = io.Copy(formFile, srcFile)
	if err != nil {
		 return
	} 

	contentType := writer.FormDataContentType()
	writer.Close() 
	req, err := http.NewRequest("POST", uploadUrl, buf)
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", contentType)
	req.Header.Set("to", to)
	req.Header.Set("AuthKey", apiKey)

	res, err := client.Do(req)
	if err != nil {
		return
	}
	rawData,err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 
	}
	err = json.Unmarshal(rawData, &response)
	return response,err
}























