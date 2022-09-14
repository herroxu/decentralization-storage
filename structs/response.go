package structs

type LoginRes struct {
	Code int					`json:"code"`
	Msg string 					`json:"msg"`
	Token string 				`json:"token"`
}

type ApiKeyRes struct {
	Code int					`json:"code"`
	Msg string 					`json:"msg"`
	Key string 					`json:"key"`
}

type UploadRes struct {
	Code int					`json:"code"`
	Msg string 					`json:"msg"`
	DecentralizedUrl string 	`json:"decentralized_url"`
	OssUrl string 				`json:"oss_url"`
}
