package structs


type LoginRes struct {
	Code string	`json:"code"`
	Msg string 	`json:"msg"`
	Token string `json:"token"`
}

type ApiKeyRes struct {
	Code string	`json:"code"`
	Msg string 	`json:"msg"`
	Key string `json:"key"`
}

type UploadRes struct {
	Code string	`json:"code"`
	Msg string 	`json:"msg"`
	Url string `json:"url"`
}
